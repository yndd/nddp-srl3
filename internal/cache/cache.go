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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Cache interface {
	HasTarget(crDeviceName string) bool
	GetValidatedGoStruct(crDeviceName string) ygot.ValidatedGoStruct
	UpdateValidatedGoStruct(crDeviceName string, s ygot.ValidatedGoStruct)
	GetCache() *yangcache.Cache
	SetModel(crDeviceName string, m *model.Model)
	GetModel(crDeviceName string) *model.Model

	ValidateCreate(crDeviceName string, x interface{}) (ygot.ValidatedGoStruct, error)
	ValidateUpdate(crDeviceName string, updates []*gnmi.Update) (ygot.ValidatedGoStruct, error)
	ValidateDelete(crDeviceName string, paths []*gnmi.Path) (ygot.ValidatedGoStruct, error)
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

func (c *cache) HasTarget(crDeviceName string) bool {
	if _, ok := c.validated[crDeviceName]; ok {
		return true
	}
	return false
}

func (c *cache) GetValidatedGoStruct(crDeviceName string) ygot.ValidatedGoStruct {
	defer c.m.Unlock()
	c.m.Lock()
	if s, ok := c.validated[crDeviceName]; ok {
		return s
	}
	return nil
}

func (c *cache) UpdateValidatedGoStruct(crDeviceName string, s ygot.ValidatedGoStruct) {
	defer c.m.Unlock()
	c.m.Lock()
	c.validated[crDeviceName] = s

	//fmt.Printf("UpdateValidatedGoStruct, deviceName: %s, goStruct: %v\n", crDeviceName, s)
}

func (c *cache) GetCache() *yangcache.Cache {
	return c.c
}

func (c *cache) SetModel(crDeviceName string, m *model.Model) {
	defer c.m.Unlock()
	c.m.Lock()
	c.model[crDeviceName] = m
}

func (c *cache) GetModel(crDeviceName string) *model.Model {
	defer c.m.Unlock()
	c.m.Lock()
	if m, ok := c.model[crDeviceName]; ok {
		return m
	}
	return nil
}

func (c *cache) ValidateCreate(crDeviceName string, x interface{}) (ygot.ValidatedGoStruct, error) {
	m := c.GetModel(crDeviceName)

	curGoStruct, err := ygot.DeepCopy(c.GetValidatedGoStruct(crDeviceName))
	if err != nil {
		return nil, err
	}

	newConfig, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

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

func (c *cache) ValidateUpdate(crDeviceName string, updates []*gnmi.Update) (ygot.ValidatedGoStruct, error) {
	//fmt.Printf("ValidateUpdate deviceName: %s, goStruct: %v\n", crDeviceName, c.GetValidatedGoStruct(crDeviceName))
	jsonTree, err := ygot.ConstructIETFJSON(c.GetValidatedGoStruct(crDeviceName), &ygot.RFC7951JSONConfig{})
	if err != nil {
		return nil, errors.Wrap(err, "error in constructing IETF JSON tree from config struct")
	}

	curGoStruct := c.GetValidatedGoStruct(crDeviceName)
	m := c.GetModel(crDeviceName)
	schema := m.SchemaTreeRoot

	//fmt.Printf("ValidateUpdate deviceName: %s, curGoStruct: %v\n", crDeviceName, c.GetValidatedGoStruct(crDeviceName))

	var goStruct ygot.ValidatedGoStruct
	for _, u := range updates {
		fullPath := cleanPath(u.GetPath())
		val := u.GetVal()

		// Validate the operation
		emptyNode, _, err := ytypes.GetOrCreateNode(schema, curGoStruct, fullPath)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("path %v is not found in the config structure", fullPath))
		}
		var nodeVal interface{}
		nodeStruct, ok := emptyNode.(ygot.ValidatedGoStruct)
		if ok {
			if err := m.JsonUnmarshaler(val.GetJsonIetfVal(), nodeStruct); err != nil {
				return nil, errors.Wrap(err, "unmarshaling json data to config struct fails")
			}
			if err := nodeStruct.Validate(); err != nil {
				return nil, errors.Wrap(err, "config data validation fails")
			}
			var err error
			if nodeVal, err = ygot.ConstructIETFJSON(nodeStruct, &ygot.RFC7951JSONConfig{}); err != nil {
				return nil, errors.Wrap(err, "error in constructing IETF JSON tree from config struct:")
			}
		} else {
			var err error
			if nodeVal, err = value.ToScalar(val); err != nil {
				return nil, errors.Wrap(err, "cannot convert leaf node to scalar type")
			}
		}

		// Update json tree of the device config.
		var curNode interface{} = jsonTree
		schema := m.SchemaTreeRoot
		for i, elem := range fullPath.GetElem() {
			switch node := curNode.(type) {
			case map[string]interface{}:
				// Set node value.
				if i == len(fullPath.GetElem())-1 {
					//c.log.Debug("getChildNode last elem", "path", yparser.GnmiPath2XPath(fullPath, true), "elem", elem, "node", node)
					if elem.GetKey() == nil {
						// err is grpcstatus error
						if err := setPathWithoutAttribute(gnmi.UpdateResult_UPDATE, node, elem, nodeVal); err != nil {
							return nil, err
						}
						break
					}
					// err is grpcstatus error
					if err := setPathWithAttribute(gnmi.UpdateResult_UPDATE, node, elem, nodeVal); err != nil {
						return nil, err
					}
					break
				}
				//c.log.Debug("getChildNode before", "path", yparser.GnmiPath2XPath(fullPath, true), "elem", elem, "node", node)
				if curNode, schema = getChildNode(node, schema, elem, true); curNode == nil {
					return nil, errors.Wrap(err, fmt.Sprintf("path elem not found: %v", elem))
				}
				//c.log.Debug("getChildNode after", "path", yparser.GnmiPath2XPath(fullPath, true), "elem", elem, "newNode", curNode)
			case []interface{}:
				return nil, errors.Wrap(err, fmt.Sprintf("incompatible path elem: %v", elem))
			default:
				return nil, errors.Wrap(err, fmt.Sprintf("wrong node type: %T", curNode))
			}
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
	return goStruct, nil
}

func (c *cache) ValidateDelete(crDeviceName string, paths []*gnmi.Path) (ygot.ValidatedGoStruct, error) {
	jsonTree, err := ygot.ConstructIETFJSON(c.GetValidatedGoStruct(crDeviceName), &ygot.RFC7951JSONConfig{})
	if err != nil {
		return nil, errors.Wrap(err, "error in constructing IETF JSON tree from config struct")
	}
	m := c.GetModel(crDeviceName)
	schema := m.SchemaTreeRoot

	var goStruct ygot.ValidatedGoStruct
	for _, path := range paths {
		var curNode interface{} = jsonTree
		pathDeleted := false
		fullPath := cleanPath(path)

		for i, elem := range fullPath.GetElem() { // Delete sub-tree or leaf node.
			node, ok := curNode.(map[string]interface{})
			if !ok {
				// we can break since the element does not exist in the schema
				break
			}

			// Delete node
			if i == len(fullPath.GetElem())-1 {
				//c.log.Debug("getChildNode last elem", "path", yparser.GnmiPath2XPath(fullPath, true), "elem", elem, "node", node)
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
			if curNode, schema = getChildNode(node, schema, elem, false); curNode == nil {
				break
			}
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
	target, hasElem := curNode[pathElem.Name]
	nodeValAsTree, nodeValIsTree := nodeVal.(map[string]interface{})
	if op == gnmi.UpdateResult_REPLACE || !hasElem || !nodeValIsTree {
		curNode[pathElem.Name] = nodeVal
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
func setPathWithAttribute(op gnmi.UpdateResult_Operation, curNode map[string]interface{}, pathElem *gnmi.PathElem, nodeVal interface{}) error {
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
			delete(m, k)
		}
	}
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
	}
	for k, v := range nodeValAsTree {
		m[k] = v
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
		return nil, nil
	}

	var nextNode interface{}
	if len(elem.GetKey()) == 0 {
		//c.log.Debug("getChildNode container", "elem name", elem.GetName(), "elem key", elem.GetKey())
		if nextNode, ok = node[elem.GetName()]; ok {
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
