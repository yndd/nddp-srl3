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

package devicecollector

import (
	"fmt"
	"strconv"

	"github.com/openconfig/goyang/pkg/yang"

	"github.com/openconfig/gnmi/proto/gnmi"
)

// getChildNode gets a node's child with corresponding schema specified by path
// element. If not found and createIfNotExist is set as true, an empty node is
// created and returned.
func (c *collector) getChildNode(node map[string]interface{}, schema *yang.Entry, elem *gnmi.PathElem, createIfNotExist bool) (interface{}, *yang.Entry) {
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

	nextNode = c.getKeyedListEntry(node, elem, createIfNotExist)
	return nextNode, nextSchema
}

// getKeyedListEntry finds the keyed list entry in node by the name and key of
// path elem. If entry is not found and createIfNotExist is true, an empty entry
// will be created (the list will be created if necessary).
func (c *collector) getKeyedListEntry(node map[string]interface{}, elem *gnmi.PathElem, createIfNotExist bool) map[string]interface{} {
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
