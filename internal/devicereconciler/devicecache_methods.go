/*
Copyright 2021 NDD.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package devicereconciler

import (
	"encoding/json"

	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/openconfig/ygot/ygot"
	"github.com/yndd/nddp-srl3/internal/cache"
	"github.com/yndd/nddp-srl3/internal/shared"
	"github.com/yndd/nddp-system/pkg/ygotnddp"
)

func (r *reconciler) getSpecdata(resource *ygotnddp.NddpSystem_Gvk) (interface{}, error) {
	spec := resource.Spec
	//r.log.Debug("getSpecdata", "specdata", *spec)
	var x1 interface{}
	if err := json.Unmarshal([]byte(*spec), &x1); err != nil {
		r.log.Debug("getSpecdata", "error", err)
		return nil, err
	}
	return x1, nil
}

// validateCreate takes the current config/goStruct and merge it with the newGoStruct
// validate the new GoStruct against the deviceSchema and return the newGoStruct
func (r *reconciler) validateCreate(resource *ygotnddp.NddpSystem_Gvk) (ygot.ValidatedGoStruct, error) {
	crDeviceName := shared.GetCrDeviceName(r.namespace, r.target.Config.Name)

	x, err := r.getSpecdata(resource)
	if err != nil {
		return nil, err
	}

	//r.log.Debug("validateCreate", "x   ", x)
	//r.log.Debug("validateCreate", "spec", *resource.Spec)

	return r.cache.ValidateCreate(crDeviceName, x)
}

// validateDelete deletes the paths from the current config/goStruct and validates the result
func (r *reconciler) validateDelete(paths []*gnmi.Path) error {
	crDeviceName := shared.GetCrDeviceName(r.namespace, r.target.Config.Name)

	return r.cache.ValidateDelete(crDeviceName, paths, cache.Origin_Reconciler)
}

// validateUpdate updates the current config/goStruct and validates the result
func (r *reconciler) validateUpdate(updates []*gnmi.Update, jsonietf bool) error {
	crDeviceName := shared.GetCrDeviceName(r.namespace, r.target.Config.Name)

	return r.cache.ValidateUpdate(crDeviceName, updates, false, jsonietf, cache.Origin_Reconciler)
}

/* we dont want to do a double diff, once in the controller and 2nd in the reconciler
func (r *reconciler) processUpdate(resource *ygotnddp.NddpSystem_Gvk) ([]*gnmi.Notification, error) {
	crDeviceName := shared.GetCrDeviceName(r.namespace, r.target.Config.Name)
	crSystemDeviceName := shared.GetCrSystemDeviceName(crDeviceName)

	ns := make([]*gnmi.Notification, 0)
	// get spec data
	x1, err := r.getSpecdata(resource)
	if err != nil {
		return nil, err
	}

	// get resource list from the cache
	resourceList, err := r.cache.GetSystemResourceList(crSystemDeviceName)
	if err != nil {
		return nil, err
	}

	// get paths related to the spec
	crPaths, err := r.cache.FindPaths(crDeviceName, x1)
	if err != nil {
		return nil, err
	}

	// get hier paths based on the resourceList
	hierPaths, err := getHierPaths(resource, crPaths, resourceList)
	if err != nil {
		return nil, err
	}

	curGoStruct, err := ygot.DeepCopy(r.cache.GetValidatedGoStruct(crDeviceName))
	if err != nil {
		return nil, err
	}

	x, err := ygot.ConstructIETFJSON(curGoStruct, &ygot.RFC7951JSONConfig{})
	if err != nil {
		return nil, err
	}

	for _, crPath := range crPaths {
		crXpath := yparser.GnmiPath2XPath(crPath, true)

		// deepcopy the spec data to avoid data manipulation of the spec
		j, err := DeepCopy(x1)
		if err != nil {
			return nil, errors.Wrap(err, "error processObserve processSpecData")
		}
		// spec Data pre-processing
		// remove all non relevant data from the spec based on the crPath
		specGoStruct, err := r.cache.GetGoStructFromPath(crDeviceName, crPath, j)
		if err != nil {
			return nil, errors.Wrap(err, "error processObserve getSpecDataFromPath")
		}

		x2, err := DeepCopy(x)
		if err != nil {
			return nil, errors.Wrap(err, "error Deepcopy x")
		}
		// remove hierarchical paths from response for this particular path
		switch x := x2.(type) {
		case map[string]interface{}:
			if hPaths, ok := hierPaths[crXpath]; ok {
				// remove hierarchical paths
				for _, hPath := range hPaths {
					x2 = removeHierarchicalResourceData(x, hPath)
				}
			}
		}

		// remove non default data from the response since it is managed by a different resource
		x2, err = r.cache.RemoveNonDefaultDataFromPath(crDeviceName, crPath, x2)
		if err != nil {
			return nil, errors.Wrap(err, "error removeNonDefaultDataFromPath")
		}

		respGoStruct, err := r.cache.GetGoStructFromPath(crDeviceName, crPath, x2)
		if err != nil {
			return nil, errors.Wrap(err, "error getSpecDataFromPath")
		}
		x2, err = ygot.EmitJSON(respGoStruct, &ygot.EmitJSONConfig{
			Format: ygot.RFC7951,
		})
		if err != nil {
			return nil, errors.Wrap(err, "error ygot EmitJSON x2")
		}

		n, err := ygot.Diff(respGoStruct, specGoStruct, &ygot.DiffPathOpt{MapToSinglePath: true})
		if err != nil {
			return nil, errors.Wrap(err, "error ygot diff")
		}
		if n != nil {
			ns = append(ns, n)
		}
	}
	return ns, nil
}

func getHierPaths(r *ygotnddp.NddpSystem_Gvk, crPaths []*gnmi.Path, resourceList map[string]*ygotnddp.NddpSystem_Gvk) (map[string][]*gnmi.Path, error) {
	hierPaths := make(map[string][]*gnmi.Path, 0)
	for _, crPath := range crPaths {
		crXpath := yparser.GnmiPath2XPath(crPath, true)
		if resourceList != nil {
			for resourceName, resource := range resourceList {
				if resourceName != *r.Name {
					for _, resourcePath := range resource.Path {
						if strings.Contains(resourcePath, crXpath) {
							if _, ok := hierPaths[crXpath]; !ok {
								hierPaths[crXpath] = make([]*gnmi.Path, 0)
							}
							//hPath, err := xpath.ToGNMIPath(strings.TrimPrefix(resourcePath, crXpath))
							hPath, err := xpath.ToGNMIPath(resourcePath)
							if err != nil {
								return nil, err
							}
							hierPaths[crXpath] = append(hierPaths[crXpath], hPath)
						}
					}
				}
			}
		}
	}
	return hierPaths, nil
}

func removeHierarchicalResourceData(x map[string]interface{}, hierPath *gnmi.Path) interface{} {
	// this is the last pathElem of the hierarchical path, which is to be deleted
	if len(hierPath.GetElem()) == 1 {
		delete(x, hierPath.GetElem()[0].GetName())
	} else {
		// there is more pathElem in the hierachical Path
		if xx, ok := x[hierPath.GetElem()[0].GetName()]; ok {
			switch x1 := xx.(type) {
			case map[string]interface{}:
				removeHierarchicalResourceData(x1, &gnmi.Path{Elem: hierPath.GetElem()[1:]})
			case []interface{}:
				for _, xxx := range x1 {
					switch x2 := xxx.(type) {
					case map[string]interface{}:
						removeHierarchicalResourceData(x2, &gnmi.Path{Elem: hierPath.GetElem()[1:]})
					}
				}
			default:
				// it can be that no data is present, so we ignore this
			}
		}
	}

	return x
}

// Make a deep copy from in into out object.
func DeepCopy(in interface{}) (interface{}, error) {
	if in == nil {
		return nil, errors.New("in cannot be nil")
	}
	//fmt.Printf("json copy input %v\n", in)
	bytes, err := json.Marshal(in)
	if err != nil {
		return nil, errors.Wrap(err, "unable to marshal input data")
	}
	var out interface{}
	err = json.Unmarshal(bytes, &out)
	if err != nil {
		return nil, errors.Wrap(err, "unable to unmarshal to output data")
	}
	//fmt.Printf("json copy output %v\n", out)
	return out, nil
}
*/
