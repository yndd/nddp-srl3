package srl

import (
	"encoding/json"
	"fmt"

	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/openconfig/goyang/pkg/yang"
	"github.com/openconfig/ygot/ygot"
	"github.com/pkg/errors"
)

func (e *externalDevice) getRootPaths(x interface{}) ([]*gnmi.Path, error) {
	//paths := make([]*gnmi.Path, 0)

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

	jsonTree, err := ygot.ConstructInternalJSON(rootSpecStruct)
	if err != nil {
		e.log.Debug("error in constructing IETF JSON tree from config struct", "error", err)
		return nil, errors.Wrap(err, "error in constructing IETF JSON tree from config struct")
	}

	return e.walkStruct(jsonTree)
}

func (e *externalDevice) walkStruct(jsonTree map[string]interface{}) ([]*gnmi.Path, error) {
	e.log.Debug("walkStruct", "jsonTree", jsonTree)
	paths := make([]*gnmi.Path, 0)

	schema := e.deviceModel.SchemaTreeRoot
	var curNode interface{} = jsonTree
	if err := e.findPaths(paths, curNode, schema); err != nil {
		e.log.Debug("walkStruct", "error", err)
		return nil, err
	}
	return paths, nil
}

func (e *externalDevice) findPaths(paths []*gnmi.Path, curNode interface{}, schema *yang.Entry) error {
	e.log.Debug("walkStruct", "curNode", curNode)
	switch node := curNode.(type) {
	case map[string]interface{}:
		var proceed bool
		var err error
		curNode, schema, proceed, err = e.getChildNode(paths, node, schema)
		if err != nil {
			return err
		}
		if !proceed {
			return nil
		}

	case []interface{}:
	default:
	}
	return nil
}

func (e *externalDevice) getChildNode(paths []*gnmi.Path, curNode map[string]interface{}, schema *yang.Entry) (interface{}, *yang.Entry, bool, error) {
	e.log.Debug("getChildNode", "curNode", curNode)
	for k, node := range curNode {
		var nextSchema *yang.Entry
		var ok bool
		if nextSchema, ok = schema.Dir[k]; !ok {
			return nil, nil, false, fmt.Errorf("wrong schema entry: %s", k)
		}
		e.log.Debug("schema info", "name", nextSchema.Name, "key", nextSchema.Key)
		if nextSchema.Key == "" {
			// todo check the content of the container, to see if data exists
			// if no data exists, proceed
			// if data exists stop
		} else {
			var err error
			nextNode, err = e.getKeyedListEntry(paths, node, k, nextSchema)
		}
		
	}
	return nil, nil, false, nil
}

func (e *externalDevice) getKeyedListEntry(paths []*gnmi.Path, curNode interface{}, elemName string, schema *yang.Entry) (map[string]interface{}, error) {
	keyedList, ok := curNode.([]interface{})
	if !ok {
		return nil, fmt.Errorf("wrong node type: %v", curNode)
	}
	for _, n := range keyedList {
		m, ok := n.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("wrong keyed list entry type: %T", n)
		}
	}
}