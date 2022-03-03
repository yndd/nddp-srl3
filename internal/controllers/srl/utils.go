package srl

import (
	"encoding/json"
	"fmt"

	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/openconfig/goyang/pkg/yang"
	"github.com/openconfig/ygot/ygot"
	"github.com/pkg/errors"
	"github.com/yndd/ndd-runtime/pkg/resource"
	"github.com/yndd/ndd-yang/pkg/yparser"
	srlv1alpha1 "github.com/yndd/nddp-srl3/apis/srl3/v1alpha1"
	"github.com/yndd/nddp-system/pkg/gvkresource"
)

func (e *externalDevice) processSpecData(x interface{}) (map[string]interface{}, error) {
	config, err := json.Marshal(x)
	if err != nil {
		return nil, errors.Wrap(err, errJSONMarshal)
	}

	e.log.Debug("config", "config", string(config))

	rootSpecStruct, err := e.deviceModel.NewConfigStruct(config, true)
	if err != nil {
		e.log.Debug("NewConfigStruct error", "error", err)
		return nil, err
	}

	return ygot.ConstructInternalJSON(rootSpecStruct)
}

func (e *externalDevice) findPaths(x interface{}) ([]*gnmi.Path, error) {
	//paths := make([]*gnmi.Path, 0)

	jsonTree, err := e.processSpecData(x)
	if err != nil {
		e.log.Debug("error in constructing IETF JSON tree from config struct", "error", err)
		return nil, errors.Wrap(err, "error in constructing IETF JSON tree from config struct")
	}

	e.log.Debug("jsonTree", "jsonTree", jsonTree)

	schema := e.deviceModel.SchemaTreeRoot
	paths := make([]*gnmi.Path, 0)
	return getChildNode(paths, jsonTree, schema, 0, true)
}

func getChildNode(paths []*gnmi.Path, curNode map[string]interface{}, schema *yang.Entry, pathDepth int, startEntry bool) ([]*gnmi.Path, error) {
	fmt.Printf("getChildNode: %v\n", curNode)
	for k, node := range curNode {
		var nextSchema *yang.Entry
		var ok bool
		if nextSchema, ok = schema.Dir[k]; !ok {
			return nil, fmt.Errorf("wrong schema entry: %s", k)
		}
		fmt.Printf("schema info name %s key: %s\n", nextSchema.Name, nextSchema.Key)
		if nextSchema.Key == "" {
			if nextSchema.Kind.String() != "Leaf" {
				fmt.Println("getContainerEntry")
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
	fmt.Printf("getContainerEntry pathElem: %s, len path elem: %d, pathDepth: %d\n", schema.Name, len(path.Elem), pathDepth)
	if len(path.Elem) == 0 || len(path.Elem) <= pathDepth {
		path.Elem = append(path.Elem, newPathElem)
	} else {
		path.Elem[pathDepth] = newPathElem
	}

	fmt.Printf("getContainerEntry container: %v\n", container)

	hasData, hasContainer, err := hasDataAndOrContainer(container, schema)
	if err != nil {
		fmt.Printf("getContainerEntry hasDataAndOrContainer: Error: %v\n", err)
		return nil, err
	}
	if hasData {
		fmt.Printf("getContainerEntry hasData: ElemName: %s\n", schema.Name)
	}
	if hasContainer {
		fmt.Printf("getContainerEntry hasContainer: ElemName: %s\n", schema.Name)
		getChildNode(paths, container, schema, pathDepth+1, false)
	}

	fmt.Printf("paths: %v\n", paths)
	return paths, nil
}

func getKeyedListEntry(paths []*gnmi.Path, curNode interface{}, elemName string, schema *yang.Entry, pathDepth int, startEntry bool) ([]*gnmi.Path, error) {
	keyedList, ok := curNode.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("wrong node type: %v", curNode)
	}
	keyedListIdx := 0
	for key, n := range keyedList {
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
		// create or update the new path elem
		newPathElem := &gnmi.PathElem{
			Name: schema.Name,
			Key: map[string]string{
				schema.Key: key,
			},
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
			fmt.Printf("hasDataAndOrContainer: Error: %v\n", err)
			return nil, err
		}
		if hasData {
			fmt.Printf("hasData: Key: %s\n", key)
			continue
		}
		if hasContainer {
			fmt.Printf("hasContainer: Key: %s\n", key)
			getChildNode(paths, nextNode, schema, pathDepth+1, false)
		}

	}
	fmt.Printf("paths: %v\n", paths)
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

func (e *externalDevice) processCreateK8s(mg resource.Managed, paths []*gnmi.Path) ([]*gnmi.Update, error) {
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
	gvkUpdates, err := e.processCreateK8sResource(mg, paths)
	if err != nil {
		return nil, err
	}

	// append gnmi data to the resource data
	gvkUpdates = append(gvkUpdates, gnmiUpdate)
	return gvkUpdates, nil
}

func (e *externalDevice) processCreateK8sResource(mg resource.Managed, paths []*gnmi.Path) ([]*gnmi.Update, error) {
	gvkData := gvkresource.GetK8sResourceCreate(mg, paths)

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
