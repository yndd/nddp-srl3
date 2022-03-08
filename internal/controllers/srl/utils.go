package srl

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/gnxi/utils/xpath"
	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/openconfig/goyang/pkg/yang"
	"github.com/openconfig/ygot/ygot"
	"github.com/pkg/errors"
	"github.com/yndd/ndd-runtime/pkg/resource"
	"github.com/yndd/ndd-yang/pkg/yparser"

	//srlv1alpha1 "github.com/yndd/nddp-srl3/apis/srl3/v1alpha1"
	"github.com/yndd/nddp-srl3/internal/shared"
	//systemv1alpha1 "github.com/yndd/nddp-system/apis/system/v1alpha1"
	"github.com/yndd/nddp-system/pkg/gvkresource"
	"github.com/yndd/nddp-system/pkg/ygotnddp"
)

func (e *externalDevice) getPaths(x interface{}) ([]*gnmi.Path, error) {
	jsonTree, err := e.processData(x)
	if err != nil {
		e.log.Debug("error in constructing IETF JSON tree from config struct", "error", err)
		return nil, errors.Wrap(err, "error in constructing IETF JSON tree from config struct")
	}

	//e.log.Debug("jsonTree", "jsonTree", jsonTree)

	schema := e.deviceModel.SchemaTreeRoot
	paths := make([]*gnmi.Path, 0)
	return getChildNode(paths, jsonTree, schema, 0, true)
}

//process Spec data marshals the data and remove the prent hierarchical keys
func (e *externalDevice) processData(x interface{}) (map[string]interface{}, error) {
	config, err := json.Marshal(x)
	if err != nil {
		return nil, errors.Wrap(err, errJSONMarshal)
	}

	//e.log.Debug("config", "config", string(config))

	rootSpecStruct, err := e.deviceModel.NewConfigStruct(config, false)
	if err != nil {
		e.log.Debug("processDataRFC7951 newConfigStruct error", "error", err)
		return nil, err
	}

	return ygot.ConstructIETFJSON(rootSpecStruct, &ygot.RFC7951JSONConfig{})
}



func getChildNode(paths []*gnmi.Path, curNode map[string]interface{}, schema *yang.Entry, pathDepth int, startEntry bool) ([]*gnmi.Path, error) {
	//fmt.Printf("getChildNode: %v\n", curNode)
	for k, node := range curNode {
		var nextSchema *yang.Entry
		var ok bool
		if nextSchema, ok = schema.Dir[k]; !ok {
			return nil, fmt.Errorf("wrong schema entry: %s", k)
		}
		//fmt.Printf("schema info name %s key: %s\n", nextSchema.Name, nextSchema.Key)
		if nextSchema.Key == "" {
			if nextSchema.Kind.String() != "Leaf" {
				//fmt.Println("getContainerEntry")
				var err error
				paths, err = getContainerEntry(paths, node, k, nextSchema, pathDepth, startEntry)
				if err != nil {
					return nil, fmt.Errorf("getKeyedListEntry: %v", err)
				}
			}
		} else {
			var err error
			paths, err = getKeyedListEntry(paths, node, k, nextSchema, pathDepth, startEntry)
			if err != nil {
				return nil, fmt.Errorf("getKeyedListEntry: %v", err)
			}
		}
	}
	return paths, nil
}

func getContainerEntry(paths []*gnmi.Path, curNode interface{}, elemName string, schema *yang.Entry, pathDepth int, startEntry bool) ([]*gnmi.Path, error) {
	container, ok := curNode.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("wrong node type: %v", curNode)
	}

	// take the latest path
	var path *gnmi.Path
	if startEntry {
		path = &gnmi.Path{}
		paths = append(paths, path)
	} else {
		path = paths[len(paths)-1]
	}

	// create or update the new path elem
	newPathElem := &gnmi.PathElem{
		Name: schema.Name,
	}
	//fmt.Printf("getContainerEntry pathElem: %s, len path elem: %d, pathDepth: %d\n", schema.Name, len(path.Elem), pathDepth)
	if len(path.Elem) == 0 || len(path.Elem) <= pathDepth {
		path.Elem = append(path.Elem, newPathElem)
	} else {
		path.Elem[pathDepth] = newPathElem
	}

	//fmt.Printf("getContainerEntry container: %v\n", container)

	hasData, hasContainer, err := hasDataAndOrContainer(container, schema)
	if err != nil {
		//fmt.Printf("getContainerEntry hasDataAndOrContainer: Error: %v\n", err)
		return nil, err
	}
	if hasData {
		//fmt.Printf("getContainerEntry hasData: ElemName: %s\n", schema.Name)
	}
	if hasContainer {
		//fmt.Printf("getContainerEntry hasContainer: ElemName: %s\n", schema.Name)
		getChildNode(paths, container, schema, pathDepth+1, false)
	}

	//fmt.Printf("paths: %v\n", paths)
	return paths, nil
}

func getKeyedListEntry(paths []*gnmi.Path, curNode interface{}, elemName string, schema *yang.Entry, pathDepth int, startEntry bool) ([]*gnmi.Path, error) {
	keyedList, ok := curNode.([]interface{})
	if !ok {
		return nil, fmt.Errorf("wrong node type: %v", curNode)
	}
	keyedListIdx := 0
	for _, n := range keyedList {
		nextNode, ok := n.(map[string]interface{})
		if !ok {
			return paths, fmt.Errorf("wrong keyed list entry type: %T", n)
		}
		//fmt.Printf("getKeyedListEntry schema elem: %s, keyname: %s dir: %v\n", schema.Name, schema.Key, schema.Dir)
		//fmt.Printf("getKeyedListEntry pathElem: %s, key: %s len paths: %d, startEntry: %t\n", schema.Name, key, len(paths), startEntry)
		// take the latest path
		var path *gnmi.Path
		if startEntry {
			path = &gnmi.Path{}
			paths = append(paths, path)
		} else {
			path = paths[len(paths)-1]
		}
		//fmt.Printf("keyedListIdx: %d\n", keyedListIdx)
		if keyedListIdx > 0 {
			path = yparser.DeepCopyGnmiPath(path)
			path.Elem = path.Elem[:pathDepth]
			paths = append(paths, path)

		}
		// get the keys from the schema and node
		schemaKeys := strings.Split(schema.Key, " ")
		keys := make(map[string]string)
		for _, k := range schemaKeys {
			keyValue, ok := nextNode[k]
			if !ok {
				return paths, fmt.Errorf("key not found in list: %v", k)
			}
			keys[k] = fmt.Sprintf("%v", keyValue)
		}

		// create or update the new path elem
		newPathElem := &gnmi.PathElem{
			Name: schema.Name,
			Key:  keys,
		}
		//fmt.Printf("getKeyedListEntry pathElem: %s, key: %s len path elem: %d, pathDepth: %d\n", schema.Name, key, len(path.Elem), pathDepth)
		if len(path.Elem) == 0 || len(path.Elem) <= pathDepth {
			path.Elem = append(path.Elem, newPathElem)
		} else {
			path.Elem[pathDepth] = newPathElem
		}

		//fmt.Printf("getKeyedListEntry pathElem: %s, key: %s paths: %v\n", schema.Name, key, paths)
		//fmt.Printf("getKeyedListEntry pathElem: %s, key: %s path: %v\n", schema.Name, key, path)

		keyedListIdx++
		startEntry = false
		hasData, hasContainer, err := hasDataAndOrContainer(nextNode, schema)
		if err != nil {
			//fmt.Printf("hasDataAndOrContainer: Error: %v\n", err)
			return nil, err
		}
		if hasData {
			//fmt.Printf("hasData: Key: %v\n", keys)
			continue
		}
		if hasContainer {
			//fmt.Printf("hasContainer: Key: %v\n", keys)
			getChildNode(paths, nextNode, schema, pathDepth+1, false)
		}

	}
	//fmt.Printf("paths: %v\n", paths)
	return paths, nil
}

func hasDataAndOrContainer(curNode map[string]interface{}, schema *yang.Entry) (bool, bool, error) {
	hasData := false
	hasContainer := false
	for elemName, node := range curNode {
		// check the elemName in the schema
		if nextSchema, ok := schema.Dir[elemName]; ok {
			// we only need to validate leafs
			if nextSchema.Kind.String() == "Leaf" {
				// check default, data present if the value of the elem is not equal to the default value
				if nextSchema.Default != nil {
					switch d := node.(type) {
					case string:
						if d != nextSchema.Default[0] {
							hasData = true
						}
					}
					continue
				}
				// elemName with the keyname should be ignored
				if schema.Key != elemName {
					hasData = true
				}
			} else {
				hasContainer = true
			}

		} else {
			return hasData, hasContainer, fmt.Errorf("wrong entry in schema: %s", elemName)
		}
	}
	return hasData, hasContainer, nil
}

/*
func (e *externalDevice) processK8s(mg resource.Managed, paths []*gnmi.Path, action systemv1alpha1.E_GvkAction) ([]*gnmi.Update, error) {
	gvkName := gvkresource.GetGvkName(mg)

	cr, ok := mg.(*srlv1alpha1.Srl3Device)
	if !ok {
		return nil, errors.New(errUnexpectedDevice)
	}

	config, err := json.Marshal(cr.Spec.Device)
	if err != nil {
		return nil, errors.Wrap(err, errJSONMarshal)
	}

	gnmiUpdate := &gnmi.Update{
		Path: &gnmi.Path{
			Elem: []*gnmi.PathElem{
				{Name: "gvk", Key: map[string]string{"name": gvkName}},
				{Name: "data"},
			},
		},
		Val: &gnmi.TypedValue{Value: &gnmi.TypedValue_JsonVal{JsonVal: config}},
	}

	// prepare gvk resource data
	gvkUpdates, err := e.processK8sResource(mg, paths, action)
	if err != nil {
		return nil, err
	}

	// append gnmi data to the resource data
	gvkUpdates = append(gvkUpdates, gnmiUpdate)
	return gvkUpdates, nil
}
*/
/*
func (e *externalDevice) processK8sResource(mg resource.Managed, paths []*gnmi.Path, action systemv1alpha1.E_GvkAction) ([]*gnmi.Update, error) {
	var gvkData *systemv1alpha1.Gvk
	switch action {
	case systemv1alpha1.E_GvkAction_Create:
		gvkData = gvkresource.GetK8sResourceCreate(mg, paths)
	case systemv1alpha1.E_GvkAction_Update:
		gvkData = gvkresource.GetK8sResourceUpdate(mg, paths)
	case systemv1alpha1.E_GvkAction_Delete:
		gvkData = gvkresource.GetK8sResourceDelete(mg, paths)
	default:
		return nil, errors.New("invalid action")
	}

	gvkPath := &gnmi.Path{
		Elem: []*gnmi.PathElem{
			{Name: "gvk", Key: map[string]string{"name": gvkData.Name}},
		},
	}
	gvkd, err := processGvkData(*gvkData)
	if err != nil {
		return nil, err
	}

	return yparser.GetUpdatesFromJSON(gvkPath, gvkd, e.nddpSchema)
}
*/

func (e *externalDevice) getResourceList(ctx context.Context, mg resource.Managed) (map[string]*ygotnddp.NddpSystem_Gvk, error) {
	// get system device list
	crSystemDeviceName := "ygot." + shared.GetCrSystemDeviceName(shared.GetCrDeviceName(mg.GetNamespace(), mg.GetNetworkNodeReference().Name))
	// gnmi get request
	reqSystemCache := &gnmi.GetRequest{
		Prefix:   &gnmi.Path{Target: crSystemDeviceName},
		Path:     []*gnmi.Path{{}},
		Encoding: gnmi.Encoding_JSON,
	}

	// gnmi get response
	resp, err := e.client.Get(ctx, reqSystemCache)
	if err != nil {
		return nil, err
	}
	var systemCache interface{}
	if len(resp.GetNotification()) == 0 {
		return nil, nil
	}
	if len(resp.GetNotification()) != 0 && len(resp.GetNotification()[0].GetUpdate()) != 0 {
		// get value from gnmi get response
		systemCache, err = yparser.GetValue(resp.GetNotification()[0].GetUpdate()[0].Val)
		if err != nil {
			return nil, errors.Wrap(err, errJSONMarshal)
		}

		switch systemCache.(type) {
		case nil:
			// resource has no data
			return nil, nil
		}
	}

	systemData, err := json.Marshal(systemCache)
	if err != nil {
		return nil, err
	}
	goStruct, err := e.systemModel.NewConfigStruct(systemData, true)
	if err != nil {
		return nil, err
	}
	nddpDevice, ok := goStruct.(*ygotnddp.Device)
	if !ok {
		return nil, errors.New("wrong object nddp")
	}

	return nddpDevice.Gvk, nil

}

func getHierPaths(mg resource.Managed, crPaths []*gnmi.Path, resourceList map[string]*ygotnddp.NddpSystem_Gvk) (map[string][]*gnmi.Path, error) {
	hierPaths := make(map[string][]*gnmi.Path, 0)
	for _, crPath := range crPaths {
		crXpath := yparser.GnmiPath2XPath(crPath, true)
		if resourceList != nil {
			for resourceName, resource := range resourceList {
				if resourceName != gvkresource.GetGvkName(mg) {
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

// 1. validate the repsonse to check if it contains the right # elements, data
func (e *externalDevice) processObserve(crPaths []*gnmi.Path, hierPaths map[string][]*gnmi.Path, specData interface{}, resp *gnmi.GetResponse) (*observe, error) {
	// validate gnmi resp information
	if len(resp.GetNotification()) == 0 {
		return &observe{hasData: false}, nil
	}
	if len(resp.GetNotification()) > 1 {
		return &observe{hasData: false}, errors.New("processObserve invalid update length")
	}
	if len(resp.GetNotification()[0].GetUpdate()) == 0 || len(resp.GetNotification()[0].GetUpdate()) > 1 {
		return &observe{hasData: false}, errors.New("processObserve invalid update length")
	}

	// validate the response data from the cache
	x, err := yparser.GetValue(resp.GetNotification()[0].GetUpdate()[0].Val)
	if err != nil {
		return nil, errors.Wrap(err, errJSONMarshal)
	}
	switch x.(type) {
	case nil:
		// resource has no data
		return &observe{hasData: false}, nil
	}

	deletes := []*gnmi.Path{}
	updates := []*gnmi.Update{}
	upToDate := true
	// for each path perform the diff between the spec and resp data
	for _, crPath := range crPaths {
		crXpath := yparser.GnmiPath2XPath(crPath, true)

		// deepcopy the spec data to avoid data manipulation of the spec
		j, err := DeepCopy(specData)
		if err != nil {
			return nil, errors.Wrap(err, "error processObserve processSpecData")
		}

		// spec Data pre-processing
		// remove all non relevant data from the spec based on the crPath
		specGoStruct, err := e.getGoStructFromPath(crPath, j)
		if err != nil {
			return nil, errors.Wrap(err, "error processObserve getSpecDataFromPath")
		}
		x1, err := ygot.ConstructIETFJSON(specGoStruct, &ygot.RFC7951JSONConfig{})
		if err != nil {
			return nil, errors.Wrap(err, "error ConstructIETFJSON x1")
		}
		/*
			x1, err := ygot.EmitJSON(specGoStruct, &ygot.EmitJSONConfig{
				Format: ygot.RFC7951,
			})
			if err != nil {
				return nil, errors.Wrap(err, "error processObserve getSpecDataFromPath x1")
			}
		*/

		// resp Data pre-processing
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
		x2, err = e.removeNonDefaultDataFromPath(crPath, x2)
		if err != nil {
			return nil, errors.Wrap(err, "error removeNonDefaultDataFromPath")
		}

		respGoStruct, err := e.getGoStructFromPath(crPath, x2)
		if err != nil {
			return nil, errors.Wrap(err, "error getSpecDataFromPath")
		}
		x2, err = ygot.ConstructIETFJSON(respGoStruct, &ygot.RFC7951JSONConfig{})
		if err != nil {
			return nil, errors.Wrap(err, "error ConstructIETFJSON x2")
		}
		/*
			x2, err = ygot.EmitJSON(respGoStruct, &ygot.EmitJSONConfig{
				Format: ygot.RFC7951,
			})
			if err != nil {
				return nil, errors.Wrap(err, "error ygot EmitJSON x2")
			}
		*/
		fmt.Printf("processObserve path   : %s  \n", crXpath)
		fmt.Printf("processObserve x1 data: %s\n", x1)
		fmt.Printf("processObserve x2 data: %v\n", x2)

		n, err := ygot.Diff(respGoStruct, specGoStruct, &ygot.DiffPathOpt{MapToSinglePath: true})
		if err != nil {
			return nil, errors.Wrap(err, "error ygot diff")
		}
		if n != nil {
			//fmt.Printf("processObserve len updates: %d\n", len(n.GetUpdate()))
			//fmt.Printf("processObserve len deletes: %d\n", len(n.GetDelete()))
			if len(n.GetUpdate()) != 0 || len(n.GetDelete()) != 0 {
				upToDate = false
			} else {
				fmt.Printf("processObserve: up To date\n")
			}
			for _, u := range n.GetUpdate() {
				fmt.Printf("processObserve: diff update old path: %s, value: %v\n", yparser.GnmiPath2XPath(u.GetPath(), true), u.GetVal())
				// workaround since the diff can return double pathElem
				update := validateUpdate(u)
				fmt.Printf("processObserve: diff update new path: %s, value: %v\n", yparser.GnmiPath2XPath(update.GetPath(), true), update.GetVal())
				updates = append(updates, update)
			}
			for _, path := range n.GetDelete() {
				fmt.Printf("processObserve: diff delete path: %s\n", yparser.GnmiPath2XPath(path, true))
				deletes = append(deletes, path)
			}
		}
	}

	return &observe{
		hasData:  true,
		upToDate: upToDate,
		deletes:  deletes,
		updates:  updates,
	}, err
}

// workaround for the dif handling
func validateUpdate(u *gnmi.Update) *gnmi.Update {
	if len(u.GetPath().GetElem()) <= 1 {
		return u
	}
	// when the 2nd last pathElem has a key and the last PathElem is an entry in the Key we should trim the last entry from the path
	// e.g. /interface[name=ethernet-1/49]/subinterface[index=1]/ipv4/address[ip-prefix=100.64.0.0/31]/ip-prefix, value: string_val:"100.64.0.0/31"
	// e.g. /interface[name=ethernet-1/49]/subinterface[index=1]/ipv4/address[ip-prefix=100.64.0.0/31]/ip-prefix, value: string_val:"100.64.0.0/31"
	if len(u.GetPath().GetElem()[len(u.GetPath().GetElem())-2].GetKey()) > 0 {
		if _, ok := u.GetPath().GetElem()[len(u.GetPath().GetElem())-2].GetKey()[u.GetPath().GetElem()[len(u.GetPath().GetElem())-1].GetName()]; ok {
			u.Path.Elem = u.Path.Elem[:len(u.GetPath().GetElem())-1]
			v := make(map[string]interface{})
			b, _ := json.Marshal(v)
			u.Val = &gnmi.TypedValue{Value: &gnmi.TypedValue_JsonVal{JsonVal: b}}
		}
	}

	return u
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
