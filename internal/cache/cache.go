package cache

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/openconfig/gnmi/value"
	"github.com/openconfig/goyang/pkg/yang"
	"github.com/openconfig/ygot/ygot"
	"github.com/openconfig/ygot/ytypes"
	"github.com/pkg/errors"
	yangcache "github.com/yndd/ndd-yang/pkg/cache"
	"github.com/yndd/ndd-yang/pkg/yparser"
	"github.com/yndd/nddp-srl3/internal/model"
	"github.com/yndd/nddp-system/pkg/ygotnddp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Cache interface {
	HasTarget(target string) bool
	GetValidatedGoStruct(target string) ygot.ValidatedGoStruct
	UpdateValidatedGoStruct(target string, s ygot.ValidatedGoStruct)
	GetCache() *yangcache.Cache
	SetModel(target string, m *model.Model)
	GetModel(target string) *model.Model

	// data manipulation methods
	ValidateCreate(target string, x interface{}) (ygot.ValidatedGoStruct, error)
	ValidateUpdate(target string, updates []*gnmi.Update, replace, jsonietf bool) (ygot.ValidatedGoStruct, error)
	ValidateDelete(target string, paths []*gnmi.Path) (ygot.ValidatedGoStruct, error)

	// system cache methods
	GetSystemExhausted(target string) (*uint32, error)
	SetSystemExhausted(target string, e uint32) error
	SetSystemCacheStatus(target string, status bool) error
	GetSystemResourceList(target string) (map[string]*ygotnddp.NddpSystem_Gvk, error)
	GetSystemResource(target, gvkName string) (*ygotnddp.NddpSystem_Gvk, error)
	UpdateSystemResourceStatus(target, resourceName, reason string, status ygotnddp.E_NddpSystem_ResourceStatus) error
	DeleteSystemResource(target, resourceName string) error

	// observe methods
	//FindPaths(target string, x interface{}) ([]*gnmi.Path, error)
	//GetGoStructFromPath(target string, rootPath *gnmi.Path, x interface{}) (ygot.ValidatedGoStruct, error)
	//RemoveNonDefaultDataFromPath(target string, rootPath *gnmi.Path, x interface{}) (interface{}, error)
}

type cache struct {
	m         sync.RWMutex
	validated map[string]ygot.ValidatedGoStruct
	model     map[string]*model.Model
	c         *yangcache.Cache
}

func New() Cache {
	return &cache{
		validated: make(map[string]ygot.ValidatedGoStruct),
		model:     make(map[string]*model.Model),
		c:         yangcache.New([]string{}),
	}
}

func (c *cache) HasTarget(target string) bool {
	if _, ok := c.validated[target]; ok {
		return true
	}
	return false
}

func (c *cache) GetValidatedGoStruct(target string) ygot.ValidatedGoStruct {
	defer c.m.Unlock()
	c.m.Lock()
	if s, ok := c.validated[target]; ok {
		if s == nil {
			fmt.Println("GetValidatedGoStruct is nil")
		}
		return s
	}
	return nil
}

func (c *cache) UpdateValidatedGoStruct(target string, s ygot.ValidatedGoStruct) {
	defer c.m.Unlock()
	c.m.Lock()
	c.validated[target] = s

	if s == nil {
		fmt.Printf("UpdateValidatedGoStruct, deviceName: %s, goStruct: %v\n", target, s)
	}

}

func (c *cache) GetCache() *yangcache.Cache {
	return c.c
}

func (c *cache) SetModel(target string, m *model.Model) {
	defer c.m.Unlock()
	c.m.Lock()
	c.model[target] = m
}

func (c *cache) GetModel(target string) *model.Model {
	defer c.m.Unlock()
	c.m.Lock()
	if m, ok := c.model[target]; ok {
		return m
	}
	return nil
}

func (c *cache) ValidateCreate(target string, x interface{}) (ygot.ValidatedGoStruct, error) {
	m := c.GetModel(target)

	var curGoStruct ygot.GoStruct
	g := c.GetValidatedGoStruct(target)
	if g == nil {
		curGoStruct = g
	} else {
		var err error
		curGoStruct, err = ygot.DeepCopy(g)
		if err != nil {
			return nil, err
		}
	}

	fmt.Printf("ValidateCreate target: %s data: %v\n", target, x)
	newConfig, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(newConfig))

	newGoStruct, err := m.NewConfigStruct(newConfig, true)
	if err != nil {
		return nil, err
	}

	if err := ygot.MergeStructInto(curGoStruct.(ygot.ValidatedGoStruct), newGoStruct); err != nil {
		return nil, err
	}

	if err := curGoStruct.(ygot.ValidatedGoStruct).Validate(); err != nil {
		return nil, err
	}

	return newGoStruct, nil
}

func (c *cache) ValidateUpdate(target string, updates []*gnmi.Update, replace, jsonietf bool) (ygot.ValidatedGoStruct, error) {
	var curGoStruct ygot.GoStruct
	g := c.GetValidatedGoStruct(target)
	if g == nil {
		curGoStruct = g
	} else {
		var err error
		curGoStruct, err = ygot.DeepCopy(g)
		if err != nil {
			return nil, err
		}
	}

	if curGoStruct == nil {
		return nil, errors.New("ValidateUpdate using empty go struct")
	}
	//fmt.Printf("ValidateUpdate deviceName: %s, goStruct: %v\n", target, curGoStruct)
	jsonTree, err := ygot.ConstructIETFJSON(curGoStruct, &ygot.RFC7951JSONConfig{})
	if err != nil {
		return nil, errors.Wrap(err, "error in constructing IETF JSON tree from config struct")
	}

	m := c.GetModel(target)

	//fmt.Printf("ValidateUpdate deviceName: %s, curGoStruct: %v\n", target, curGoStruct)
	//fmt.Printf("ValidateUpdate deviceName: %s, jsonTree: %v\n", target, jsonTree)
	for _, u := range updates {
		fullPath := cleanPath(u.GetPath())
		val := u.GetVal()
		fmt.Printf("ValidateUpdate path: %s, val: %v\n", yparser.GnmiPath2XPath(fullPath, true), val)
	}

	var goStruct ygot.ValidatedGoStruct
	for _, u := range updates {
		fullPath := cleanPath(u.GetPath())
		val := u.GetVal()

		fmt.Printf("ValidateUpdate path: %s, val: %v\n", yparser.GnmiPath2XPath(fullPath, true), val)
		//	if replace {
		//		fmt.Printf("ValidateUpdate path: %s, val: %v\n", yparser.GnmiPath2XPath(fullPath, true), val)
		//	}

		// we need to return to the schema root for the next update
		schema := m.SchemaTreeRoot

		// Validate the operation
		var emptyNode interface{}
		if replace {
			//fmt.Printf("ValidateUpdate val: %v\n", val)
			// in case of delete we first delete the node and recreate it again
			if err := ytypes.DeleteNode(schema, curGoStruct, fullPath); err != nil {
				return nil, errors.Wrap(err, fmt.Sprintf("path %v cannot delete path", fullPath))
			}
			emptyNode, _, err = ytypes.GetOrCreateNode(schema, curGoStruct, fullPath)
			if err != nil {
				return nil, errors.Wrap(err, fmt.Sprintf("path %v is not found in the config structure", fullPath))
			}
		} else {
			emptyNode, _, err = ytypes.GetOrCreateNode(schema, curGoStruct, fullPath)
			if err != nil {
				return nil, errors.Wrap(err, fmt.Sprintf("path %v is not found in the config structure", fullPath))
			}
		}

		/*
			if replace ||
				yparser.GnmiPath2XPath(fullPath, true) == "/interface[name=ethernet-1/49]/subinterface[index=1]/ipv4/address[ip-prefix=100.64.0.0/31]" {
				nodeStruct, _ := emptyNode.(ygot.ValidatedGoStruct)
				vvvv, err := ygot.ConstructIETFJSON(nodeStruct, &ygot.RFC7951JSONConfig{})
				if err != nil {
					return nil, errors.Wrap(err, "error in constructing IETF JSON tree from config struct:")
				}
				fmt.Printf("ValidateUpdate vvvv: %v\n", vvvv)
				fmt.Printf("ValidateUpdate nodeStruct: %v\n", nodeStruct)
			}
		*/
		var nodeVal interface{}
		nodeStruct, ok := emptyNode.(ygot.ValidatedGoStruct)
		if ok {
			v := val.GetJsonVal()
			if jsonietf {
				v = val.GetJsonIetfVal()
			}
			/*
				if replace {
					fmt.Printf("ValidateUpdate ok nodeStruct: %v\n", nodeStruct)
				}
			*/

			if err := m.JsonUnmarshaler(v, nodeStruct); err != nil {
				fmt.Printf("unmarshal error: nodeStruct: %v\n", nodeStruct)
				fmt.Printf("unmarshal error: path: %s, val: %v\n", fullPath, v)
				vvvv, err := ygot.ConstructIETFJSON(nodeStruct, &ygot.RFC7951JSONConfig{})
				if err != nil {
					return nil, errors.Wrap(err, "error in constructing IETF JSON tree from config struct:")
				}
				fmt.Printf("ValidateUpdate vvvv: %v\n", vvvv)
				return nil, errors.Wrap(err, "unmarshaling json data to config struct fails")
			}
			if err := nodeStruct.Validate(); err != nil {
				return nil, errors.Wrap(err, "config data validation fails")
			}
			var err error
			if nodeVal, err = ygot.ConstructIETFJSON(nodeStruct, &ygot.RFC7951JSONConfig{}); err != nil {
				return nil, errors.Wrap(err, "error in constructing IETF JSON tree from config struct:")
			}
			//if replace {
			//	fmt.Printf("ValidateUpdate ok nodeVal 1: %v\n", nodeVal)

			//	if err := json.Unmarshal(v, &nodeVal); err != nil {
			//		return nil, errors.Wrap(err, "unmarshaling json data failed")
			//	}
			//	fmt.Printf("ValidateUpdate ok nodeVal 2: %v\n", nodeVal)

			//}
		} else {
			/*
				if replace {
					fmt.Printf("ValidateUpdate nok nodeStruct: %v\n", nodeStruct)
				}
			*/
			var err error
			/*
				if nodeVal, err = value.ToScalar(val); err != nil {
					return nil, errors.Wrap(err, "cannot convert leaf node to scalar type")
				}
			*/
			if nodeVal, err = yparser.GetValue(val); err != nil {
				return nil, errors.Wrap(err, "cannot convert leaf node to scalar type")
			}

			if !replace {
				fmt.Printf("ValidateUpdate scalar nodeVal: %v\n", nodeVal)
			}

		}
		// replace or update
		op := gnmi.UpdateResult_UPDATE
		if replace {
			op = gnmi.UpdateResult_REPLACE
		}
		/*
			if op == gnmi.UpdateResult_REPLACE {
				fmt.Printf("ValidateUpdate nodeVal: %v\n", nodeVal)
			}
		*/

		// Update json tree of the device config.
		var curNode interface{} = jsonTree
		schema = m.SchemaTreeRoot
		for i, elem := range fullPath.GetElem() {
			switch node := curNode.(type) {
			case map[string]interface{}:
				// Set node value.
				if i == len(fullPath.GetElem())-1 {
					/*
						if op == gnmi.UpdateResult_REPLACE {
							fmt.Printf("ValidateUpdate: getChildNode last elem, path: %s, index: %d, elem: %v, node: %v, nodeVal: %v\n", yparser.GnmiPath2XPath(fullPath, true), i, elem, node, nodeVal)
						}
					*/

					if len(elem.GetKey()) == 0 {
						// err is grpcstatus error
						fmt.Printf("setPathWithoutAttribute: schemaName: %s, schemaKind: %s\n", schema.Name, schema.Kind.String())
						if newSchema, ok := schema.Dir[elem.Name]; ok {
							fmt.Printf("setPathWithoutAttribute: newSchemaName: %s, newSchemaKind: %s\n", newSchema.Name, newSchema.Kind.String())
						}
						if err := setPathWithoutAttribute(op, node, elem, nodeVal); err != nil {
							fmt.Printf("setPathWithoutAttribute error: %v\n", err)
							return nil, err
						}
						// set defaults

						if err := setDefaults(node, elem, schema); err != nil {
							fmt.Printf("setPathWithoutAttribute setDefaults error: %v\n", err)
							return nil, err
						}

						break
					}
					// err is grpcstatus error
					fmt.Printf("setPathWithAttribute: schemaName: %s, schemaKind: %s\n", schema.Name, schema.Kind.String())
					if err := setPathWithAttribute(op, node, elem, nodeVal, schema); err != nil {
						fmt.Printf("setPathWithAttribute: error: %v\n", err)
						return nil, err
					}
					break
				}
				//fmt.Printf("ValidateUpdate: getChildNode before: %s, index: %d, elem: %v, node: %v\n", yparser.GnmiPath2XPath(fullPath, true), i, elem, node)
				if curNode, schema = getChildNode(node, schema, elem, true); curNode == nil {
					return nil, errors.Wrap(err, fmt.Sprintf("path elem not found: %v", elem))
				}
				fmt.Printf("ValidateUpdate: getChildNode after : %s, index: %d, elem: %v, node: %v\n", yparser.GnmiPath2XPath(fullPath, true), i, elem, curNode)
				fmt.Printf("ValidateUpdate: getChildNode after : %s, schemaDir: %v\n", yparser.GnmiPath2XPath(fullPath, true), schema.Dir)
			case []interface{}:
				return nil, errors.Wrap(err, fmt.Sprintf("incompatible path elem: %v", elem))
			default:
				return nil, errors.Wrap(err, fmt.Sprintf("wrong node type: %T", curNode))
			}
		}
		if strings.Contains(yparser.GnmiPath2XPath(fullPath, true), "/interface[name=ethernet-1/49]") {
			fmt.Printf("ValidateUpdate jsonTree: %v\n", jsonTree["interface"])
		}
		jsonDump, err := json.Marshal(jsonTree)
		if err != nil {
			return nil, fmt.Errorf("error in marshaling IETF JSON tree to bytes: %v", err)
		}

		goStruct, err = m.NewConfigStruct(jsonDump, true)
		if err != nil {
			return nil, fmt.Errorf("error in creating config struct from IETF JSON data: %v", err)
		}
	}
	fmt.Printf("ValidateUpdate jsonTree done\n")
	return goStruct, nil
}

func (c *cache) ValidateDelete(target string, paths []*gnmi.Path) (ygot.ValidatedGoStruct, error) {
	var curGoStruct ygot.GoStruct
	g := c.GetValidatedGoStruct(target)
	if g == nil {
		curGoStruct = g
	} else {
		var err error
		curGoStruct, err = ygot.DeepCopy(g)
		if err != nil {
			return nil, err
		}
	}

	if curGoStruct == nil {
		return nil, errors.New("ValidateDelete using empty go struct")
	}

	jsonTree, err := ygot.ConstructIETFJSON(curGoStruct, &ygot.RFC7951JSONConfig{})
	if err != nil {
		return nil, errors.Wrap(err, "error in constructing IETF JSON tree from config struct")
	}
	m := c.GetModel(target)

	var goStruct ygot.ValidatedGoStruct
	for _, path := range paths {
		var curNode interface{} = jsonTree
		pathDeleted := false
		fullPath := cleanPath(path)
		fmt.Printf("ValidateDelete path: %s\n", yparser.GnmiPath2XPath(fullPath, true))

		// we need to go back to the schemaroot for the next delete path
		schema := m.SchemaTreeRoot

		for i, elem := range fullPath.GetElem() { // Delete sub-tree or leaf node.
			node, ok := curNode.(map[string]interface{})
			if !ok {
				// we can break since the element does not exist in the schema
				break
			}

			// Delete node
			if i == len(fullPath.GetElem())-1 {
				//c.log.Debug("getChildNode last elem", "path", yparser.GnmiPath2XPath(fullPath, true), "elem", elem, "node", node)
				//fmt.Printf("ValidateDelete last elem, path: %s, index: %d, elem: %v, node: %v\n", yparser.GnmiPath2XPath(fullPath, true), i, elem, node)
				if len(elem.GetKey()) == 0 {
					// check schema for defaults and fallback to default if required
					if schema, ok = schema.Dir[elem.GetName()]; ok {
						if len(schema.Default) > 0 {
							if err := setPathWithoutAttribute(gnmi.UpdateResult_UPDATE, node, elem, schema.Default[0]); err != nil {
								return nil, err
							}
							// should be update, but we abuse pathDeleted
							pathDeleted = true
							break
						}
					}
					delete(node, elem.GetName())
					pathDeleted = true
					break
				}
				pathDeleted = deleteKeyedListEntry(node, elem)
				break
			}
			//fmt.Printf("ValidateDelete: getChildNode before: %s, index: %d, elem: %v, node: %v\n", yparser.GnmiPath2XPath(fullPath, true), i, elem, node)
			if curNode, schema = getChildNode(node, schema, elem, false); curNode == nil {
				break
			}
			//fmt.Printf("ValidateDelete: getChildNode after: %s, index: %d, elem: %v, node: %v\n", yparser.GnmiPath2XPath(fullPath, true), i, elem, curNode)
		}

		// Validate the new config
		if pathDeleted {
			jsonDump, err := json.Marshal(jsonTree)
			if err != nil {
				return nil, fmt.Errorf("error in marshaling IETF JSON tree to bytes: %v", err)
			}
			goStruct, err = m.NewConfigStruct(jsonDump, true)
			if err != nil {
				return nil, fmt.Errorf("error in creating config struct from IETF JSON data: %v", err)
			}
		}
	}
	return goStruct, nil
}

func cleanPath(path *gnmi.Path) *gnmi.Path {
	// clean the path for now to remove the module information from the pathElem
	p := yparser.DeepCopyGnmiPath(path)
	for _, pe := range p.GetElem() {
		pe.Name = strings.Split(pe.Name, ":")[len(strings.Split(pe.Name, ":"))-1]
		keys := make(map[string]string)
		for k, v := range pe.GetKey() {
			if strings.Contains(v, "::") {
				keys[strings.Split(k, ":")[len(strings.Split(k, ":"))-1]] = v
			} else {
				keys[strings.Split(k, ":")[len(strings.Split(k, ":"))-1]] = strings.Split(v, ":")[len(strings.Split(v, ":"))-1]
			}
		}
		pe.Key = keys
	}
	return p
}

// setPathWithoutAttribute replaces or updates a child node of curNode in the
// IETF config tree, where the child node is indexed by pathElem without
// attribute. The function returns grpc status error if unsuccessful.
func setPathWithoutAttribute(op gnmi.UpdateResult_Operation, curNode map[string]interface{}, pathElem *gnmi.PathElem, nodeVal interface{}) error {
	//fmt.Printf("ValidateUpdate: setPathWithoutAttribute, operation: %v, elem: %v, node: %v, nodeVal: %v\n", op, pathElem, curNode, nodeVal)
	target, hasElem := curNode[pathElem.Name]
	nodeValAsTree, nodeValIsTree := nodeVal.(map[string]interface{})
	if op == gnmi.UpdateResult_REPLACE || !hasElem || !nodeValIsTree {
		curNode[pathElem.Name] = nodeVal
		fmt.Printf("ValidateUpdate: curNode: %v, nodeVal: %v\n", curNode, nodeVal)
		return nil
	}
	targetAsTree, ok := target.(map[string]interface{})
	if !ok {
		return status.Errorf(codes.Internal, "error in setting path: expect map[string]interface{} to update, got %T", target)
	}
	for k, v := range nodeValAsTree {
		targetAsTree[k] = v
	}

	return nil
}

// setPathWithAttribute replaces or updates a child node of curNode in the IETF
// JSON config tree, where the child node is indexed by pathElem with attribute.
// The function returns grpc status error if unsuccessful.
func setPathWithAttribute(op gnmi.UpdateResult_Operation, curNode map[string]interface{}, pathElem *gnmi.PathElem, nodeVal interface{}, schema *yang.Entry) error {
	//fmt.Printf("ValidateUpdate: setPathWithAttribute, operation: %v, elem: %v, node: %v, nodeVal: %v\n", op, pathElem, curNode, nodeVal)
	nodeValAsTree, ok := nodeVal.(map[string]interface{})
	if !ok {
		return status.Errorf(codes.InvalidArgument, "expect nodeVal is a json node of map[string]interface{}, received %T", nodeVal)
	}
	m := getKeyedListEntry(curNode, pathElem, true)
	if m == nil {
		return status.Errorf(codes.NotFound, "path elem not found: %v", pathElem)
	}
	if op == gnmi.UpdateResult_REPLACE {
		for k := range m {
			//fmt.Printf("ValidateUpdate: setPathWithAttribute 1: k: %v, m: %v\n", k, m)
			delete(m, k)
		}
	}
	// Debug to be removed below
	//if op == gnmi.UpdateResult_REPLACE {
	//	fmt.Printf("ValidateUpdate: setPathWithAttribute 2: m: %v\n", m)
	//}
	// Debug to be removed above
	for attrKey, attrVal := range pathElem.GetKey() {
		m[attrKey] = attrVal
		if asNum, err := strconv.ParseFloat(attrVal, 64); err == nil {
			m[attrKey] = asNum
		}
		for k, v := range nodeValAsTree {
			if k == attrKey && fmt.Sprintf("%v", v) != attrVal {
				return status.Errorf(codes.InvalidArgument, "invalid config data: %v is a path attribute", k)
			}
		}
		// Debug to be removed below
		//if op == gnmi.UpdateResult_REPLACE {
		//	fmt.Printf("ValidateUpdate: setPathWithAttribute 3: m: %v\n", m)
		//}
		// Debug to be removed above
	}
	for k, v := range nodeValAsTree {
		m[k] = v
		// Debug to be removed below
		//if op == gnmi.UpdateResult_REPLACE {
		//	fmt.Printf("ValidateUpdate: setPathWithAttribute 4: k: %v, v: %v\n", k, v)
		//}
		// Debug to be removed above
	}
	// Debug to be removed below
	//if op == gnmi.UpdateResult_REPLACE {
	//	fmt.Printf("ValidateUpdate: setPathWithAttribute 5: m: %v\n", m)
	//	fmt.Printf("ValidateUpdate: setPathWithAttribute 5: curNode: %v\n", curNode)
	//}
	// Debug to be removed above

	// set defaults
	/*
		newSchema, ok := schema.Dir[pathElem.GetName()]
		if ok {
			if err := setDefaults(m, newSchema); err != nil {
				fmt.Printf("set default error: %v\n", err)
				return err
			}
		}
	*/
	if err := setDefaults(m, pathElem, schema); err != nil {
		fmt.Printf("set default error: %v\n", err)
		return err
	}

	return nil
}

// deleteKeyedListEntry deletes the keyed list entry from node that matches the
// path elem. If the entry is the only one in keyed list, deletes the entire
// list. If the entry is found and deleted, the function returns true. If it is
// not found, the function returns false.
func deleteKeyedListEntry(node map[string]interface{}, elem *gnmi.PathElem) bool {
	curNode, ok := node[elem.Name]
	if !ok {
		return false
	}

	keyedList, ok := curNode.([]interface{})
	if !ok {
		return false
	}
	for i, n := range keyedList {
		m, ok := n.(map[string]interface{})
		if !ok {
			fmt.Printf("expect map[string]interface{} for a keyed list entry, got %T", n)
			return false
		}
		keyMatching := true
		for k, v := range elem.Key {
			attrVal, ok := m[k]
			if !ok {
				return false
			}
			if v != fmt.Sprintf("%v", attrVal) {
				keyMatching = false
				break
			}
		}
		if keyMatching {
			listLen := len(keyedList)
			if listLen == 1 {
				delete(node, elem.Name)
				return true
			}
			keyedList[i] = keyedList[listLen-1]
			node[elem.Name] = keyedList[0 : listLen-1]
			return true
		}
	}
	return false
}

// getChildNode gets a node's child with corresponding schema specified by path
// element. If not found and createIfNotExist is set as true, an empty node is
// created and returned.
func getChildNode(node map[string]interface{}, schema *yang.Entry, elem *gnmi.PathElem, createIfNotExist bool) (interface{}, *yang.Entry) {
	var nextSchema *yang.Entry
	var ok bool

	//fmt.Printf("elem name: %s, key: %v\n", elem.GetName(), elem.GetKey())
	//c.log.Debug("getChildNode", "elem name", elem.GetName(), "elem key", elem.GetKey())

	if nextSchema, ok = schema.Dir[elem.GetName()]; !ok {
		// returning nil will be picked up as an error
		return nil, nil
	}

	var nextNode interface{}
	if len(elem.GetKey()) == 0 {
		//c.log.Debug("getChildNode container", "elem name", elem.GetName(), "elem key", elem.GetKey())
		if nextNode, ok = node[elem.GetName()]; !ok {
			//c.log.Debug("getChildNode new container entry", "elem name", elem.GetName(), "elem key", elem.GetKey())
			if createIfNotExist {
				node[elem.Name] = make(map[string]interface{})
				nextNode = node[elem.GetName()]
			}
		}
		return nextNode, nextSchema
	}

	nextNode = getKeyedListEntry(node, elem, createIfNotExist)
	return nextNode, nextSchema
}

// getKeyedListEntry finds the keyed list entry in node by the name and key of
// path elem. If entry is not found and createIfNotExist is true, an empty entry
// will be created (the list will be created if necessary).
func getKeyedListEntry(node map[string]interface{}, elem *gnmi.PathElem, createIfNotExist bool) map[string]interface{} {
	//c.log.Debug("getKeyedListEntry", "elem name", elem.GetName(), "elem key", elem.GetKey())
	curNode, ok := node[elem.GetName()]
	if !ok {
		if !createIfNotExist {
			return nil
		}

		// Create a keyed list as node child and initialize an entry.
		m := make(map[string]interface{})
		for k, v := range elem.GetKey() {
			m[k] = v
			if vAsNum, err := strconv.ParseFloat(v, 64); err == nil {
				m[k] = vAsNum
			}
		}
		node[elem.GetName()] = []interface{}{m}
		return m
	}

	// Search entry in keyed list.
	keyedList, ok := curNode.([]interface{})
	if !ok {
		switch m := curNode.(type) {
		case map[string]interface{}:
			return m
		default:
			return nil

		}

	}
	for _, n := range keyedList {
		m, ok := n.(map[string]interface{})
		if !ok {
			fmt.Printf("wrong keyed list entry type: %T", n)
			return nil
		}
		keyMatching := true
		// must be exactly match
		for k, v := range elem.GetKey() {
			attrVal, ok := m[k]
			if !ok {
				return nil
			}
			if v != fmt.Sprintf("%v", attrVal) {
				keyMatching = false
				break
			}
		}
		if keyMatching {
			return m
		}
	}
	if !createIfNotExist {
		return nil
	}

	// Create an entry in keyed list.
	m := make(map[string]interface{})
	for k, v := range elem.GetKey() {
		m[k] = v
		if vAsNum, err := strconv.ParseFloat(v, 64); err == nil {
			m[k] = vAsNum
		}
	}
	node[elem.GetName()] = append(keyedList, m)
	return m
}

func setDefaults(node map[string]interface{}, pathElem *gnmi.PathElem, schema *yang.Entry) error {
	// check schema for defaults and fallback to default if required
	newSchema, ok := schema.Dir[pathElem.GetName()]
	if !ok {
		return fmt.Errorf("wrong schema in setDefaults elem: %v", pathElem)
	}
	if newSchema.Kind.String() == "Leaf" {
		// this is a leaf
		if len(newSchema.Default) > 0 {
			fmt.Printf("ValidateUpdate: set default value: elem: %s, schema default: %v, node: %v\n", pathElem.GetName(), schema.Default, node)
			setDefaultValue(node, pathElem.GetName(), newSchema)
		}
		return nil
	}
	// this is a directory

	for elem, schema := range newSchema.Dir {
		fmt.Printf("ValidateUpdate: set default: elem: %s,  schema default: %v\n", elem, schema.Default)
		if len(schema.Default) > 0 {
			fmt.Printf("ValidateUpdate: set default value: elem: %s, schema default: %v, node: %v\n", elem, schema.Default, node)
			setDefaultValue(node, elem, schema)
		}
	}
	return nil
}

func setDefaultValue(node map[string]interface{}, elem string, schema *yang.Entry) error {
	fmt.Printf("ValidateUpdate: default elem: %s, val: %v, type: %s\n", elem, schema.Default, schema.Type.Kind.String())
	var nodeVal interface{}
	nodeVal = schema.Default[0]
	switch schema.Type.Kind.String() {
	case "boolean":
		v, err := strconv.ParseBool(schema.Default[0])
		if err != nil {
			return err
		}
		if nodeVal, err = value.ToScalar(&gnmi.TypedValue{Value: &gnmi.TypedValue_BoolVal{BoolVal: v}}); err != nil {
			return errors.Wrap(err, "cannot convert leaf node to scalar type")
		}
		fmt.Printf("ValidateUpdate: elem: %s nodeVal: %#v\n", elem, nodeVal)
	case "uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64":
		d, err := strconv.Atoi(schema.Default[0])
		if err != nil {
			return err
		}
		switch schema.Type.Kind.String() {
		case "uint8":
			nodeVal = uint8(d)
		case "uint16":
			nodeVal = uint16(d)
		case "uint32":
			nodeVal = uint32(d)
		case "uint64":
			nodeVal = uint64(d)
		case "int8":
			nodeVal = int8(d)
		case "int16":
			nodeVal = int16(d)
		case "int32":
			nodeVal = int32(d)
		case "int64":
			nodeVal = int64(d)
		}
	}
	if err := setPathWithoutAttribute(gnmi.UpdateResult_UPDATE, node, &gnmi.PathElem{Name: elem}, nodeVal); err != nil {
		return err
	}
	return nil
}

func (c *cache) GetSystemExhausted(target string) (*uint32, error) {
	defer c.m.Unlock()
	c.m.Lock()
	goStruct, ok := c.validated[target]
	if !ok {
		return nil, errors.New("target not ready")
	}

	nddpDevice, ok := goStruct.(*ygotnddp.Device)
	if !ok {
		return nil, errors.New("wrong object nddp")
	}
	return nddpDevice.Cache.Exhausted, nil
}

func (c *cache) SetSystemExhausted(target string, e uint32) error {
	defer c.m.Unlock()
	c.m.Lock()
	goStruct, ok := c.validated[target]
	if !ok {
		return errors.New("target not ready")
	}

	nddpDevice, ok := goStruct.(*ygotnddp.Device)
	if !ok {
		return errors.New("wrong object nddp")
	}
	nddpDevice.Cache.Exhausted = ygot.Uint32(e)
	*nddpDevice.Cache.ExhaustedNbr++

	c.validated[target] = nddpDevice
	return nil
}

func (c *cache) SetSystemCacheStatus(target string, status bool) error {
	defer c.m.Unlock()
	c.m.Lock()
	goStruct, ok := c.validated[target]
	if !ok {
		return errors.New("target not ready")
	}
	nddpDevice, ok := goStruct.(*ygotnddp.Device)
	if !ok {
		return errors.New("wrong object nddp")
	}
	nddpDevice.GetOrCreateCache().Update = ygot.Bool(status)

	c.validated[target] = nddpDevice
	return nil

}

func (c *cache) GetSystemResourceList(target string) (map[string]*ygotnddp.NddpSystem_Gvk, error) {
	defer c.m.Unlock()
	c.m.Lock()
	goStruct, ok := c.validated[target]
	if !ok {
		return nil, errors.New("target not ready")
	}

	nddpDevice, ok := goStruct.(*ygotnddp.Device)
	if !ok {
		return nil, errors.New("wrong object nddp")
	}
	return nddpDevice.Gvk, nil
}

func (c *cache) GetSystemResource(target, resourceName string) (*ygotnddp.NddpSystem_Gvk, error) {
	defer c.m.Unlock()
	c.m.Lock()
	goStruct, ok := c.validated[target]
	if !ok {
		return nil, errors.New("target not ready")
	}

	nddpDevice, ok := goStruct.(*ygotnddp.Device)
	if !ok {
		return nil, errors.New("wrong object nddp")
	}
	r, ok := nddpDevice.Gvk[resourceName]
	if !ok {
		errors.New("resource not found")
	}

	return r, nil
}

func (c *cache) UpdateSystemResourceStatus(target, resourceName, reason string, status ygotnddp.E_NddpSystem_ResourceStatus) error {
	defer c.m.Unlock()
	c.m.Lock()
	goStruct, ok := c.validated[target]
	if !ok {
		return errors.New("target not ready")
	}
	nddpDevice, ok := goStruct.(*ygotnddp.Device)
	if !ok {
		return errors.New("wrong object nddp")
	}

	r, ok := nddpDevice.Gvk[resourceName]
	if !ok {
		errors.New("resource not found")
	}

	if status == ygotnddp.NddpSystem_ResourceStatus_FAILED {
		*r.Attempt++
		// we dont update the status to failed unless we tried 3 times
		if *r.Attempt > 3 {
			r.Status = ygotnddp.NddpSystem_ResourceStatus_FAILED
			r.Reason = ygot.String(reason)
		}
	} else {
		// success
		r.Status = status
		r.Reason = ygot.String(reason)
	}

	c.validated[target] = nddpDevice
	return nil
}

func (c *cache) DeleteSystemResource(target, resourceName string) error {
	defer c.m.Unlock()
	c.m.Lock()
	goStruct, ok := c.validated[target]
	if !ok {
		return errors.New("target not ready")
	}
	nddpDevice, ok := goStruct.(*ygotnddp.Device)
	if !ok {
		return errors.New("wrong object nddp")
	}
	nddpDevice.DeleteGvk(resourceName)

	c.validated[target] = nddpDevice
	return nil
}
