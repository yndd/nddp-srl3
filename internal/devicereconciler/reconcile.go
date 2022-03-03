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
	"context"

	"github.com/google/gnxi/utils/xpath"
	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/openconfig/ygot/ygot"
	systemv1alpha1 "github.com/yndd/nddp-system/apis/system/v1alpha1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (r *reconciler) reconcileCreate(ctx context.Context, resource *systemv1alpha1.Gvk) error {
	log := r.log.WithValues("target", r.target.Config.Name, "address", r.target.Config.Address)
	log.Debug("reconciling device config create")

	// validateResource, merges the latest device config with the new config
	// and validates it against the device Yang schema
	newGoStruct, err := r.validateCreate(resource)
	if err != nil {
		log.Debug("validation failed", "errpr", err)
	}
	j, err := ygot.EmitJSON(newGoStruct, &ygot.EmitJSONConfig{
		Format: ygot.RFC7951,
		//Indent:        "",
		RFC7951Config: &ygot.RFC7951JSONConfig{},
	})

	log.Debug("json update", "json", j)

	updates := []*gnmi.Update{
		{
			Path: &gnmi.Path{},
			Val:  &gnmi.TypedValue{Value: &gnmi.TypedValue_JsonIetfVal{JsonIetfVal: []byte(j)}},
		},
	}
	if _, err := r.device.UpdateGnmi(ctx, updates); err != nil {
		// update failed
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.ResourceExhausted:
				log.Debug("gnmi update failed exhausted", "error", err)
				if err := r.setExhausted(60); err != nil {
					return err
				}
				// we return and keep the status as is since we can retry once the device is back in normal state
				return nil
			}
		}
		// update failed, update resource status in the system cache
		if err := r.updateResourceStatus(resource.Name, systemv1alpha1.E_GvkStatus_Failed); err != nil {
			return err
		}
		return err
	}
	// update succeeded, update resource status in the system cache
	if err := r.updateResourceStatus(resource.Name, systemv1alpha1.E_GvkStatus_Success); err != nil {
		return err
	}
	return nil
}

func (r *reconciler) reconcileUpdate(ctx context.Context, resource *systemv1alpha1.Gvk) error {
	log := r.log.WithValues("target", r.target.Config.Name, "address", r.target.Config.Address)
	log.Debug("reconciling device config update")

	// check deletes, updates
	// TBD HOW DO WE DO THE DIFF
	/*
		deletes, updates, err := r.processUpdates(resource)
		if err != nil {
			return err
		}
	*/
	deletes := []*gnmi.Path{}
	updates := []*gnmi.Update{}

	if err := r.validateDelete(deletes); err != nil {
		return err
	}
	if err := r.validateUpdate(updates); err != nil {
		return err
	}

	// execute the deletes and updates in the cache and to the device
	_, err := r.device.SetGnmi(r.ctx, updates, deletes)
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.ResourceExhausted:
				log.Debug("gnmi update failed exhausted", "error", err)
				if err := r.setExhausted(60); err != nil {
					return err
				}
			}
		}
		// Set status to failed
		if err := r.updateResourceStatus(resource.Name, systemv1alpha1.E_GvkStatus_Failed); err != nil {
			return err
		}
		return err
	}
	if err := r.updateResourceStatus(resource.Name, systemv1alpha1.E_GvkStatus_Success); err != nil {
		return err
	}
	return nil
}

func (r *reconciler) reconcileDelete(ctx context.Context, resource *systemv1alpha1.Gvk) error {
	log := r.log.WithValues("target", r.target.Config.Name, "address", r.target.Config.Address)
	log.Debug("reconciling device config update")

	delPaths := make([]*gnmi.Path, 0)
	murder := false
	for _, xp := range resource.Paths {
		path, err := xpath.ToGNMIPath(xp)
		if err != nil {
			return err
		}

		if path == nil || len(path.Elem) == 0 {
			murder = true
		}
		delPaths = append(delPaths, path)
	}
	// if we dont do suicide and len delete paths > 0, perform delete
	if !murder && len(delPaths) > 0 {
		// validate Delete
		if err := r.validateDelete(delPaths); err != nil {
			return err
		}
		// apply deletes on the device
		_, err := r.device.DeleteGnmi(ctx, delPaths)
		if err != nil {
			if e, ok := status.FromError(err); ok {
				switch e.Code() {
				case codes.ResourceExhausted:
					log.Debug("gnmi delate failed exhausted", "error", err)
					if err := r.setExhausted(60); err != nil {
						return err
					}
					// we return and keep the status as is since we can retry once the device is back in normal state
					return nil
				}
			}
			log.Debug("gnmi delete failed", "Paths", delPaths, "Error", err)
			// update failed, update resource status in the system cache
			if err := r.updateResourceStatus(resource.Name, systemv1alpha1.E_GvkStatus_Failed); err != nil {
				return err
			}
			// we only process 1 resource at the time
			return err

		}
		// delete resources from the system cache
		if err := r.deleteResource(resource.Name); err != nil {
			return err
		}
	}
	return nil
}

func (r *reconciler) PrintResourceList(idx int) error {
	resourceListRaw, err := r.getResourceListRaw()
	if err != nil {
		return err
	}
	r.log.Debug("resourceList", "idx", idx, "raw", resourceListRaw)
	return nil
}
