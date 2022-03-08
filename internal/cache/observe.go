package cache

/*
import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/openconfig/goyang/pkg/yang"
	"github.com/openconfig/ygot/ygot"
	"github.com/pkg/errors"
	"github.com/yndd/ndd-yang/pkg/yparser"
)

func (c *cache) processData(crDeviceName string, x interface{}, internal bool) (map[string]interface{}, error) {
	config, err := json.Marshal(x)
	if err != nil {
		return nil, errors.Wrap(err, "cache FindPaths error marshal json")
	}

	//e.log.Debug("config", "config", string(config))
	m := c.GetModel(crDeviceName)
	rootSpecStruct, err := m.NewConfigStruct(config, false)
	if err != nil {
		return nil, err
	}

	if internal {
		return ygot.ConstructInternalJSON(rootSpecStruct)
	}
	return ygot.ConstructIETFJSON(rootSpecStruct, &ygot.RFC7951JSONConfig{})
}

func (c *cache) FindPaths(crDeviceName string, x interface{}) ([]*gnmi.Path, error) {
	jsonTree, err := c.processData(crDeviceName, x, true)
	if err != nil {
		return nil, errors.Wrap(err, "error in constructing IETF JSON tree from config struct")
	}
	m := c.GetModel(crDeviceName)
	schema := m.SchemaTreeRoot
	paths := make([]*gnmi.Path, 0)
	return getFindPathsChildNode(paths, jsonTree, schema, 0, true)
}

func getFindPathsChildNode(paths []*gnmi.Path, curNode map[string]interface{}, schema *yang.Entry, pathDepth int, startEntry bool) ([]*gnmi.Path, error) {
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
				paths, err = getFindPathsContainerEntry(paths, node, k, nextSchema, pathDepth, startEntry)
				if err != nil {
					return nil, fmt.Errorf("getKeyedListEntry: %v", err)
				}
			}
		} else {
			var err error
			paths, err = getFindPathsKeyedListEntry(paths, node, k, nextSchema, pathDepth, startEntry)
			if err != nil {
				return nil, fmt.Errorf("getKeyedListEntry: %v", err)
			}
		}
	}
	return paths, nil
}

func getFindPathsContainerEntry(paths []*gnmi.Path, curNode interface{}, elemName string, schema *yang.Entry, pathDepth int, startEntry bool) ([]*gnmi.Path, error) {
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
	_, hasContainer, err := hasFindPathsDataAndOrContainer(container, schema)
	if err != nil {
		//fmt.Printf("getContainerEntry hasDataAndOrContainer: Error: %v\n", err)
		return nil, err
	}
	if hasContainer {
		//fmt.Printf("getContainerEntry hasContainer: ElemName: %s\n", schema.Name)
		getFindPathsChildNode(paths, container, schema, pathDepth+1, false)
	}

	//fmt.Printf("paths: %v\n", paths)
	return paths, nil
}

func getFindPathsKeyedListEntry(paths []*gnmi.Path, curNode interface{}, elemName string, schema *yang.Entry, pathDepth int, startEntry bool) ([]*gnmi.Path, error) {
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
		hasData, hasContainer, err := hasFindPathsDataAndOrContainer(nextNode, schema)
		if err != nil {
			//fmt.Printf("hasDataAndOrContainer: Error: %v\n", err)
			return nil, err
		}
		if hasData {
			//fmt.Printf("hasData: Key: %s\n", key)
			continue
		}
		if hasContainer {
			//fmt.Printf("hasContainer: Key: %s\n", key)
			getFindPathsChildNode(paths, nextNode, schema, pathDepth+1, false)
		}

	}
	//fmt.Printf("paths: %v\n", paths)
	return paths, nil
}

func hasFindPathsDataAndOrContainer(curNode map[string]interface{}, schema *yang.Entry) (bool, bool, error) {
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

func (c *cache) GetGoStructFromPath(crDeviceName string, rootPath *gnmi.Path, x interface{}) (ygot.ValidatedGoStruct, error) {
	jsonIETFTree, err := c.processData(crDeviceName, x, false)
	if err != nil {
		return nil, errors.Wrap(err, "error in constructing IETF JSON tree from config struct")
	}

	//e.log.Debug("jsonIETFTree", "jsonIETFTree", jsonIETFTree)
	m := c.GetModel(crDeviceName)
	schema := m.SchemaTreeRoot
	if err := getData(rootPath, jsonIETFTree, schema); err != nil {
		return nil, errors.Wrap(err, "error getting spec Data from config")
	}

	jsonDump, err := json.Marshal(jsonIETFTree)
	if err != nil {
		return nil, errors.Wrap(err, "error in marshaling IETF JSON tree to bytes")
	}
	//fmt.Println(string(jsonDump))
	goStruct, err := m.NewConfigStruct(jsonDump, false)
	if err != nil {
		return nil, errors.Wrap(err, "error in creating config struct from IETF JSON data")
	}
	return goStruct, nil
}

func getData(rootPath *gnmi.Path, curNode map[string]interface{}, schema *yang.Entry) error {
	if len(rootPath.GetElem()) != 0 {
		// delete the data that is not relevant for this spec
		for pathElemName := range curNode {
			// avoid cutting the keys from the json/node struct
			if schema.Key == "" {
				if pathElemName != rootPath.GetElem()[0].GetName() {
					delete(curNode, pathElemName)
				}
			}
		}
		node, ok := curNode[rootPath.GetElem()[0].GetName()]
		if !ok {
			return fmt.Errorf("getSpecData data not found in jsonTree: %s", rootPath.GetElem()[0].GetName())
		}
		nextSchema, ok := schema.Dir[rootPath.GetElem()[0].GetName()]
		if !ok {
			return fmt.Errorf("getSpecData wrong schema entry: %s", rootPath.GetElem()[0].GetName())
		}

		//fmt.Printf("getSpecData schema info name %s key: %s\n", nextSchema.Name, nextSchema.Key)
		if nextSchema.Key == "" {
			if nextSchema.Kind.String() != "Leaf" {
				nextNode, ok := node.(map[string]interface{})
				if !ok {
					return fmt.Errorf("getSpecData wrong node type: %v", curNode)
				}
				return getData(&gnmi.Path{Elem: rootPath.GetElem()[1:]}, nextNode, nextSchema)
			}
		} else {
			var err error
			curNode[rootPath.GetElem()[0].GetName()], err = getNextKeyedListEntry(rootPath, node, nextSchema)
			if err != nil {
				return err
			}
		}

	}
	return nil
}

func getNextKeyedListEntry(rootPath *gnmi.Path, curNode interface{}, schema *yang.Entry) (interface{}, error) {
	//fmt.Printf("getSpecNextKeyedListEntry: curNode %v\n", curNode)

	switch keyedNode := curNode.(type) {
	case []interface{}:
		var n map[string]interface{}
		var found bool
		//var idx int
		for _, node := range keyedNode {
			nextNode, ok := node.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("getSpecNextKeyedListEntry nextNode wrong node type: %v", nextNode)
			}
			found = true
			for keyName, keyValue := range rootPath.GetElem()[0].GetKey() {
				//fmt.Printf("keyName: %s, keyValue: %s \n", keyName, keyValue)
				//fmt.Printf("nextNode:: %s \n", nextNode[keyName])
				var nextNodeKeyValue string
				switch kv := nextNode[keyName].(type) {
				case string:
					nextNodeKeyValue = string(kv)
				case uint32:
					nextNodeKeyValue = strconv.Itoa(int(kv))
				case float64:
					nextNodeKeyValue = fmt.Sprintf("%.0f", kv)
				}
				if nextNodeKeyValue != keyValue {
					found = false
					break
				}
			}
			if found {
				//idx = i
				n = nextNode
				break
			}
		}
		if !found {
			return nil, fmt.Errorf("getSpecNextKeyedListEntry index not found: %v", rootPath.GetElem()[0].GetKey())
		}
		//fmt.Printf("idx: %d, len(keyedNode): %d\n", idx, len(keyedNode))
		//curNode = append(keyedNode[:0], keyedNode[:idx+1]...)

		curNode = make([]interface{}, 0, 1)
		curNode = append(curNode.([]interface{}), n)
		//fmt.Printf("keyedNode: %v\n", keyedNode)
		//fmt.Printf("curNode: %v\n", curNode)

		nextNode, ok := curNode.([]interface{})[0].(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("wrong keyed list entry type: %T", curNode.([]interface{})[0])
		}
		//fmt.Printf("getSpecNextKeyedListEntry schema elem: %s, keyname: %s dir: %v, key: %v\n", schema.Name, schema.Key, schema.Dir, rootPath.GetElem()[0].GetKey())
		//fmt.Printf("getSpecNextKeyedListEntry nextNode: %v \n", nextNode)

		if err := getData(&gnmi.Path{Elem: rootPath.GetElem()[1:]}, nextNode, schema); err != nil {
			return nil, err
		}
		return curNode, nil

	default:
		return nil, fmt.Errorf("getSpecNextKeyedListEntry keyedNode wrong node type: %v", curNode)
	}
}

func (c *cache) RemoveNonDefaultDataFromPath(crDeviceName string, rootPath *gnmi.Path, x interface{}) (interface{}, error) {
	jsonIETFTree, err := c.processData(crDeviceName, x, false)
	if err != nil {
		return nil, errors.Wrap(err, "error in constructing IETF JSON tree from config struct")
	}

	//e.log.Debug("jsonIETFTree", "jsonIETFTree", jsonIETFTree)
	m := c.GetModel(crDeviceName)
	schema := m.SchemaTreeRoot
	if err := removeNonDefaultData(rootPath, jsonIETFTree, schema); err != nil {
		return nil, errors.Wrap(err, "error getting spec Data from config")
	}
	return jsonIETFTree, nil
}

func removeNonDefaultData(rootPath *gnmi.Path, curNode map[string]interface{}, schema *yang.Entry) error {
	if len(rootPath.GetElem()) >= 1 {
		// delete the data that is not relevant for this spec
		for pathElemName, node := range curNode {
			nextSchema, ok := schema.Dir[pathElemName]
			if !ok {
				return fmt.Errorf("removeNonDefaultData wrong pathElemName: %s, schema: %v", pathElemName, schema.Dir)
			}
			// avoid cutting default leafs
			//fmt.Printf("removeNonDefaultData: path: %s, pathElemName: %s\n", yparser.GnmiPath2XPath(rootPath, true), pathElemName)
			if nextSchema.Kind.String() == "Leaf" {
				if nextSchema.Default != nil {
					switch d := node.(type) {
					case string:
						if d != nextSchema.Default[0] {
							delete(curNode, pathElemName)
						}
					}
					continue
				}
				// pathElemName with the keyname should be ignored
				if schema.Key != pathElemName {
					delete(curNode, pathElemName)
				}
			}
		}
		node, ok := curNode[rootPath.GetElem()[0].GetName()]
		if !ok {
			return fmt.Errorf("getSpecData data not found in jsonTree: %s", rootPath.GetElem()[0].GetName())
		}
		nextSchema, ok := schema.Dir[rootPath.GetElem()[0].GetName()]
		if !ok {
			return fmt.Errorf("getSpecData wrong schema entry: %s", rootPath.GetElem()[0].GetName())
		}

		//fmt.Printf("getSpecData schema info name %s key: %s\n", nextSchema.Name, nextSchema.Key)
		if nextSchema.Key == "" {
			if nextSchema.Kind.String() != "Leaf" {
				nextNode, ok := node.(map[string]interface{})
				if !ok {
					return fmt.Errorf("getSpecData wrong node type: %v", curNode)
				}
				return removeNonDefaultData(&gnmi.Path{Elem: rootPath.GetElem()[1:]}, nextNode, nextSchema)
			}
		} else {
			if err := findNextKeyedListEntry(rootPath, node, nextSchema); err != nil {
				return err
			}
		}
	}
	return nil
}

func findNextKeyedListEntry(rootPath *gnmi.Path, curNode interface{}, schema *yang.Entry) error {
	//fmt.Printf("getSpecNextKeyedListEntry: curNode %v\n", curNode)

	switch keyedNode := curNode.(type) {
	case []interface{}:
		//var n map[string]interface{}
		var found bool
		var idx int
		for i, node := range keyedNode {
			nextNode, ok := node.(map[string]interface{})
			if !ok {
				return fmt.Errorf("findNextKeyedListEntry nextNode wrong node type: %v", nextNode)
			}
			found = true
			for keyName, keyValue := range rootPath.GetElem()[0].GetKey() {
				//fmt.Printf("keyName: %s, keyValue: %s \n", keyName, keyValue)
				//fmt.Printf("nextNode:: %s \n", nextNode[keyName])
				var nextNodeKeyValue string
				switch kv := nextNode[keyName].(type) {
				case string:
					nextNodeKeyValue = string(kv)
				case uint32:
					nextNodeKeyValue = strconv.Itoa(int(kv))
				case float64:
					nextNodeKeyValue = fmt.Sprintf("%.0f", kv)
				}
				if nextNodeKeyValue != keyValue {
					found = false
					break
				}
			}
			if found {
				idx = i
				//n = nextNode
				break
			}
		}
		if !found {
			return fmt.Errorf("findNextKeyedListEntry index not found: %v", rootPath.GetElem()[0].GetKey())
		}
		//fmt.Printf("idx: %d, len(keyedNode): %d\n", idx, len(keyedNode))
		//curNode = append(keyedNode[:0], keyedNode[:idx+1]...)

		//curNode = make([]interface{}, 0, 1)
		//curNode = append(curNode.([]interface{}), n)
		//fmt.Printf("keyedNode: %v\n", keyedNode)
		//fmt.Printf("curNode: %v\n", curNode)

		nextNode, ok := curNode.([]interface{})[idx].(map[string]interface{})
		if !ok {
			return fmt.Errorf("wrong keyed list entry type: %T", curNode.([]interface{})[idx])
		}
		//fmt.Printf("getSpecNextKeyedListEntry schema elem: %s, keyname: %s dir: %v, key: %v\n", schema.Name, schema.Key, schema.Dir, rootPath.GetElem()[0].GetKey())
		//fmt.Printf("getSpecNextKeyedListEntry nextNode: %v \n", nextNode)

		if err := removeNonDefaultData(&gnmi.Path{Elem: rootPath.GetElem()[1:]}, nextNode, schema); err != nil {
			return err
		}
		return nil

	default:
		return fmt.Errorf("findNextKeyedListEntry keyedNode wrong node type: %v", curNode)
	}

}
*/
