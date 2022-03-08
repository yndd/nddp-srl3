package srl

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/openconfig/goyang/pkg/yang"
	"github.com/openconfig/ygot/ygot"
	"github.com/pkg/errors"
)

func (e *externalDevice) removeNonDefaultDataFromPath(rootPath *gnmi.Path, x interface{}) (interface{}, error) {
	jsonIETFTree, err := e.processData(x)
	if err != nil {
		e.log.Debug("error in constructing IETF JSON tree from config struct", "error", err)
		return nil, errors.Wrap(err, "error in constructing IETF JSON tree from config struct")
	}

	//e.log.Debug("jsonIETFTree", "jsxxonIETFTree", jsonIETFTree)

	schema := e.deviceModel.SchemaTreeRoot
	if err := removeNonDefaultData(rootPath, jsonIETFTree, schema); err != nil {
		//e.log.Debug("error getting spec Data from config", "error", err)
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
				return fmt.Errorf("removeNonDefaultData wrong pathElemName: %s, schema: %v\n", pathElemName, schema.Dir)
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

func (e *externalDevice) getGoStructFromPath(rootPath *gnmi.Path, x interface{}) (ygot.ValidatedGoStruct, error) {
	jsonIETFTree, err := e.processData(x)
	if err != nil {
		e.log.Debug("error in constructing IETF JSON tree from config struct", "error", err)
		return nil, errors.Wrap(err, "error in constructing IETF JSON tree from config struct")
	}

	//e.log.Debug("jsonIETFTree", "jsonIETFTree", jsonIETFTree)

	schema := e.deviceModel.SchemaTreeRoot
	if err := getData(rootPath, jsonIETFTree, schema); err != nil {
		//e.log.Debug("error getting spec Data from config", "error", err)
		return nil, errors.Wrap(err, "error getting spec Data from config")
	}

	jsonDump, err := json.Marshal(jsonIETFTree)
	if err != nil {
		e.log.Debug("error in marshaling IETF JSON tree to bytes", "error", err)
		return nil, errors.Wrap(err, "error in marshaling IETF JSON tree to bytes")
	}
	//fmt.Println(string(jsonDump))
	goStruct, err := e.deviceModel.NewConfigStruct(jsonDump, false)
	if err != nil {
		e.log.Debug("error in creating config struct from IETF JSON data", "error", err)
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
