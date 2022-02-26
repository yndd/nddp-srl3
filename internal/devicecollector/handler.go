package devicecollector

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/openconfig/gnmi/value"
	"github.com/openconfig/ygot/ygot"
	"github.com/openconfig/ygot/ytypes"
	"github.com/pkg/errors"
	"github.com/yndd/ndd-yang/pkg/yparser"
	srlv1alpha1 "github.com/yndd/nddp-srl3/apis/srl3/v1alpha1"
	"github.com/yndd/nddp-srl3/internal/shared"
	systemv1alpha1 "github.com/yndd/nddp-system/apis/system/v1alpha1"
	"github.com/yndd/nddp-system/pkg/gvkresource"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
)

const (
	unmanagedResource = "Unmanaged resource"
)

var (
	pbRootPath     = &gnmi.Path{}
	pbDevicePrefix = &gnmi.Path{Elem: []*gnmi.PathElem{{Name: "device"}}}
)

func (c *collector) handleSubscription(resp *gnmi.SubscribeResponse) error {
	log := c.log.WithValues("target", c.target.Config.Name, "address", c.target.Config.Address)
	//log.Debug("handle target update from device")

	switch resp.GetResponse().(type) {
	case *gnmi.SubscribeResponse_Update:
		//log.Debug("handle target update from device", "Prefix", resp.GetUpdate().GetPrefix())

		// check if the target cache exists
		crDeviceName := shared.GetCrDeviceName(c.namespace, c.target.Config.Name)

		if !c.cache.HasTarget(crDeviceName) {
			log.Debug("handle target update target not found in ygot schema cache")
			return errors.New("target cache does not exist")
		}

		if !c.cache.GetCache().GetCache().HasTarget(crDeviceName) {
			log.Debug("handle target update target not found in cache")
			return errors.New("target cache does not exist")
		}

		resourceList, err := c.getResourceList(crDeviceName)
		if err != nil {
			return err
		}
		//log.Debug("resourceList", "list", resourceList)

		jsonTree, err := ygot.ConstructIETFJSON(c.cache.GetValidatedGoStruct(crDeviceName), &ygot.RFC7951JSONConfig{})
		if err != nil {
			log.Debug("error in constructing IETF JSON tree from config struct", "error", err)
			return errors.Wrap(err, "error in constructing IETF JSON tree from config struct")
		}

		// handle deletes
		if err := c.handleDeletes(crDeviceName, resourceList, jsonTree, resp.GetUpdate().Delete); err != nil {
			return err
		}

		if err := c.handleUpdates(crDeviceName, resourceList, jsonTree, resp.GetUpdate().Update); err != nil {
			return err
		}

	case *gnmi.SubscribeResponse_SyncResponse:
		//log.Debug("SyncResponse")
	}

	return nil
}

func (c *collector) handleDeletes(crDeviceName string, resourceList []*systemv1alpha1.Gvk, jsonTree map[string]interface{}, delPaths []*gnmi.Path) error {
	m := c.cache.GetModel(crDeviceName)
	schema := m.SchemaTreeRoot

	for _, path := range delPaths {
		xpath := yparser.GnmiPath2XPath(path, true)
		resourceName, err := c.findManagedResource(xpath, resourceList)
		if err != nil {
			return err
		}

		var curNode interface{} = jsonTree
		pathDeleted := false
		fullPath := cleanPath(path)

		c.log.Debug("subscription config delete", "path", yparser.GnmiPath2XPath(fullPath, true))
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
					//c.log.Debug("schema", "yangEntry", *schema)
					// check schema for defaults
					if schema, ok = schema.Dir[elem.GetName()]; ok {
						if len(schema.Default) > 0 {
							//c.log.Debug("schema", "yangEntry", *schema, "default", schema.Default)

							if grpcStatusError := setPathWithoutAttribute(gnmi.UpdateResult_UPDATE, node, elem, schema.Default[0]); grpcStatusError != nil {
								c.log.Debug("setPathWithoutAttribute", "path", yparser.GnmiPath2XPath(fullPath, true), "error", grpcStatusError)
								return grpcStatusError
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

			//c.log.Debug("getChildNode before", "path", yparser.GnmiPath2XPath(fullPath, true), "elem", elem, "node", node)
			if curNode, schema = c.getChildNode(node, schema, elem, false); curNode == nil {
				break
			}
			//c.log.Debug("getChildNode after", "path", yparser.GnmiPath2XPath(fullPath, true), "elem", elem, "newNode", curNode)
		}
		if reflect.DeepEqual(fullPath, pbRootPath) { // Delete root
			for k := range jsonTree {
				delete(jsonTree, k)
			}
		}

		// Apply the validated operation to the config tree and device.
		if pathDeleted {
			if err := c.UpdateValidatedConfig(crDeviceName, jsonTree); err != nil {
				return err
			}
			/*
				if s.callback != nil {
					if applyErr := s.callback(newConfig); applyErr != nil {
						if rollbackErr := s.callback(s.config); rollbackErr != nil {
							return nil, status.Errorf(codes.Internal, "error in rollback the failed operation (%v): %v", applyErr, rollbackErr)
						}
						return nil, status.Errorf(codes.Aborted, "error in applying operation to device: %v", applyErr)
					}
				}
			*/
		}

		/*
			// update the cache with the latest config from the device
			if err := c.cache.GetCache().GnmiUpdate(crDeviceName, n); err != nil {
				c.log.Debug("handle target update", "error", err, "Path", xpath)
				return errors.New("cache update failed")
			}
		*/

		if *resourceName != unmanagedResource {
			// TODO Trigger reconcile event
			c.triggerReconcileEvent(resourceName)
		}
	}
	return nil
}

func (c *collector) handleUpdates(crDeviceName string, resourceList []*systemv1alpha1.Gvk, jsonTree map[string]interface{}, upd []*gnmi.Update) error {
	m := c.cache.GetModel(crDeviceName)
	for _, u := range upd {
		xpath := yparser.GnmiPath2XPath(u.GetPath(), true)
		// check if this is a managed resource or unmanged resource
		// name == unmanagedResource is an unmanaged resource
		resourceName, err := c.findManagedResource(xpath, resourceList)
		if err != nil {
			return err
		}
		walkInternalGoStruct := true

		if !walkInternalGoStruct {
			// we create a json blob and merge this in the gostruct
			fullPath := u.GetPath()
			val := u.GetVal()
			json, err := generateJson(u)
			if err != nil {
				c.log.Debug("generate json error", "err", err)
				return err
			}
			c.log.Debug("subscription config update", "path", yparser.GnmiPath2XPath(fullPath, true), "val", val)
			//c.log.Debug("subscription config update", "path", yparser.GnmiPath2XPath(fullPath, true), "json", string(json))

			// create newGostruct which will not be validated
			newGoStruct, err := m.NewConfigStruct(json, false)
			if err != nil {
				c.log.Debug("generate new gostruct", "err", err)
				return err
			}

			// merge the newGostruct with the current config
			currGoStruct := c.cache.GetValidatedGoStruct(crDeviceName)
			if err := ygot.MergeStructInto(currGoStruct, newGoStruct); err != nil {
				c.log.Debug("merge  gostructs", "err", err)
				return err
			}

			// validate the merged config
			if err := currGoStruct.Validate(); err != nil {
				c.log.Debug("validate new gostructs", "error", err)
				return err
			}
			// since validation is successfull we can set the newGostruct as the new valid config
			c.cache.UpdateValidatedGoStruct(crDeviceName, newGoStruct)
		} else {
			fullPath := cleanPath(u.GetPath())
			val := u.GetVal()

			//c.log.Debug("subscription config update", "path", yparser.GnmiPath2XPath(fullPath, true), "val", val)

			// Validate the operation
			config := c.cache.GetValidatedGoStruct(crDeviceName)
			emptyNode, _, err := ytypes.GetOrCreateNode(m.SchemaTreeRoot, config, fullPath)
			if err != nil {
				c.log.Debug("path not found in config structure", "path", fullPath, "error", err)
				return errors.Wrap(err, fmt.Sprintf("path %v is not found in the config structure", fullPath))
			}
			//c.log.Debug("emptyNode", "emptyNode", emptyNode)
			var nodeVal interface{}
			nodeStruct, ok := emptyNode.(ygot.ValidatedGoStruct)
			if ok {
				if err := m.JsonUnmarshaler(val.GetJsonIetfVal(), nodeStruct); err != nil {
					c.log.Debug("unmarshaling json data to config struct fails", "error", err)
					return errors.Wrap(err, "unmarshaling json data to config struct fails")
				}
				if err := nodeStruct.Validate(); err != nil {
					c.log.Debug("config data validation fails", "error", err)
					return errors.Wrap(err, "config data validation fails")
				}
				var err error
				if nodeVal, err = ygot.ConstructIETFJSON(nodeStruct, &ygot.RFC7951JSONConfig{}); err != nil {
					c.log.Debug("error in constructing IETF JSON tree from config struct", "error", err)
					return errors.Wrap(err, "error in constructing IETF JSON tree from config struct:")
				}
			} else {
				var err error
				if nodeVal, err = value.ToScalar(val); err != nil {
					c.log.Debug("cannot convert leaf node to scalar type", "error", err)
					return errors.Wrap(err, "cannot convert leaf node to scalar type")
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
							if grpcStatusError := setPathWithoutAttribute(gnmi.UpdateResult_UPDATE, node, elem, nodeVal); grpcStatusError != nil {
								c.log.Debug("setPathWithoutAttribute", "path", yparser.GnmiPath2XPath(fullPath, true), "error", grpcStatusError)
								return grpcStatusError
							}
							break
						}
						if grpcStatusError := c.setPathWithAttribute(gnmi.UpdateResult_UPDATE, node, elem, nodeVal); grpcStatusError != nil {
							c.log.Debug("setPathWithAttribute", "path", yparser.GnmiPath2XPath(fullPath, true), "error", grpcStatusError)
							return grpcStatusError
						}
						break
					}
					//c.log.Debug("getChildNode before", "path", yparser.GnmiPath2XPath(fullPath, true), "elem", elem, "node", node)
					if curNode, schema = c.getChildNode(node, schema, elem, true); curNode == nil {
						c.log.Debug("path elem not found", "elem", elem)
						return errors.Wrap(err, fmt.Sprintf("path elem not found: %v", elem))
					}
					//c.log.Debug("getChildNode after", "path", yparser.GnmiPath2XPath(fullPath, true), "elem", elem, "newNode", curNode)
				case []interface{}:
					c.log.Debug("wrong type incompatible path elem", "elem", elem)
					return errors.Wrap(err, fmt.Sprintf("incompatible path elem: %v", elem))
				default:
					c.log.Debug("wrong type", "node", curNode)
					return errors.Wrap(err, fmt.Sprintf("wrong node type: %T", curNode))
				}
			}
			if reflect.DeepEqual(fullPath, pbRootPath) { // Replace/Update root.
				//if op == gnmi.UpdateResult_UPDATE {
				return errors.Wrap(err, "update the root of config tree is unsupported")
				//}

				//	nodeValAsTree, ok := nodeVal.(map[string]interface{})
				//	if !ok {
				//		return errors.Wrap(err, fmt.Sprintf("expect a tree to replace the root, got a scalar value: %T", nodeVal))
				//	}
				//	for k := range jsonTree {
				//		delete(jsonTree, k)
				//	}
				//	for k, v := range nodeValAsTree {
				//		jsonTree[k] = v
				//	}

			}

			if err := c.UpdateValidatedConfig(crDeviceName, jsonTree); err != nil {
				return err
			}
		}

		// TO BE ADDED AGAIN

		/*
			// update the cache with the latest config from the device
			if err := c.cache.GetCache().GnmiUpdate(crDeviceName, n); err != nil {
				for _, u := range n.GetUpdate() {
					c.log.Debug("collector config update", "path", yparser.GnmiPath2XPath(u.GetPath(), true), "value", u.GetVal(), "error", err)
				}
				return errors.Wrap(err, "cache update failed")
			}

			for _, u := range n.GetUpdate() {
				c.log.Debug("collector config update", "path", yparser.GnmiPath2XPath(u.GetPath(), true), "value", u.GetVal())
			}
		*/

		if *resourceName != unmanagedResource {
			c.triggerReconcileEvent(resourceName)
		}
	}
	return nil
}

func (c *collector) UpdateValidatedConfig(crDeviceName string, jsonTree map[string]interface{}) error {
	m := c.cache.GetModel(crDeviceName)
	jsonDump, err := json.Marshal(jsonTree)
	if err != nil {
		return fmt.Errorf("error in marshaling IETF JSON tree to bytes: %v", err)
	}
	goStruct, err := m.NewConfigStruct(jsonDump, true)
	if err != nil {
		return fmt.Errorf("error in creating config struct from IETF JSON data: %v", err)
	}
	c.cache.UpdateValidatedGoStruct(crDeviceName, goStruct)
	return nil
}

func (c *collector) findManagedResource(xpath string, resourceList []*systemv1alpha1.Gvk) (*string, error) {
	matchedResourceName := unmanagedResource
	matchedResourcePath := ""
	for _, r := range resourceList {
		if strings.Contains(xpath, r.Rootpath) {
			// if there is a better match we use the better match
			if len(r.Rootpath) > len(matchedResourcePath) {
				matchedResourcePath = r.Rootpath
				matchedResourceName = r.Name
			}
		}
	}
	return &matchedResourceName, nil
}

func (c *collector) getResourceList(crDeviceName string) ([]*systemv1alpha1.Gvk, error) {
	crSystemDeviceName := shared.GetCrSystemDeviceName(crDeviceName)

	rl, err := c.cache.GetCache().GetJson(crSystemDeviceName,
		&gnmi.Path{Target: crSystemDeviceName},
		&gnmi.Path{Elem: []*gnmi.PathElem{{Name: "gvk"}}},
		c.nddpSchema)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return gvkresource.GetResourceList(rl)
}

func (c *collector) triggerReconcileEvent(resourceName *string) error {

	gvk, err := gvkresource.String2Gvk(*resourceName)
	if err != nil {
		return err
	}
	kindgroup := strings.Join([]string{gvk.Kind, gvk.Group}, ".")

	object := getObject(gvk)

	//c.log.Debug("triggerReconcileEvent", "kindgroup", kindgroup, "gvk", gvk, "object", object)

	if eventCh, ok := c.eventChs[kindgroup]; ok {
		c.log.Debug("triggerReconcileEvent with channel lookup", "kindgroup", kindgroup, "gvk", gvk, "object", object)
		eventCh <- event.GenericEvent{
			Object: object,
		}
	}
	return nil
}

func getObject(gvk *gvkresource.GVK) client.Object {
	switch gvk.Kind {
	case "Srl3Device":
		return &srlv1alpha1.Srl3Device{
			ObjectMeta: metav1.ObjectMeta{Name: gvk.Name, Namespace: gvk.NameSpace},
		}
	default:
		fmt.Printf("getObject not found gvk: %v\n", *gvk)
		return nil
	}
}

// gnmiFullPath builds the full path from the prefix and path.
func gnmiFullPath(prefix, path *gnmi.Path) *gnmi.Path {
	fullPath := &gnmi.Path{Origin: path.Origin}
	if path.GetElement() != nil {
		fullPath.Element = append(prefix.GetElement(), path.GetElement()...)
	}
	if path.GetElem() != nil {
		fullPath.Elem = append(prefix.GetElem(), path.GetElem()...)
	}
	return fullPath
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
func (c *collector) setPathWithAttribute(op gnmi.UpdateResult_Operation, curNode map[string]interface{}, pathElem *gnmi.PathElem, nodeVal interface{}) error {
	nodeValAsTree, ok := nodeVal.(map[string]interface{})
	if !ok {
		return status.Errorf(codes.InvalidArgument, "expect nodeVal is a json node of map[string]interface{}, received %T", nodeVal)
	}
	m := c.getKeyedListEntry(curNode, pathElem, true)
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

func generateJson(u *gnmi.Update) ([]byte, error) {
	var err error
	var d interface{}
	path := yparser.DeepCopyGnmiPath(u.GetPath())
	if d, err = addData(d, path.GetElem(), u.GetVal()); err != nil {
		return nil, err
	}
	return json.Marshal(d)

}

func addData(d interface{}, pe []*gnmi.PathElem, val *gnmi.TypedValue) (interface{}, error) {
	var err error
	if len(pe) == 0 {
		return nil, nil
	}
	e := pe[0].GetName()
	k := pe[0].GetKey()
	//fmt.Printf("addData, Len: %d, Elem: %s, Key: %v, Data: %v\n", len(elems), e, k, d)
	if len(pe)-1 == 0 {
		// last element
		if len(k) == 0 {
			// last element with container
			d, err = addContainerValue(d, e, val)
			return d, err
		} else {
			// last element with list
			// not sure if this will ever exist
			d, err = addListValue(d, e, k, val)
			return d, err
		}
	} else {
		if len(k) == 0 {
			// not last element -> container
			d, err = addContainer(d, e, pe, val)
			return d, err
		} else {
			// not last element -> list + keys
			d, err = addList(d, e, k, pe, val)
			return d, err
		}
	}
}

func addContainer(d interface{}, e string, elems []*gnmi.PathElem, val *gnmi.TypedValue) (interface{}, error) {
	var err error
	// initialize the data
	//fmt.Printf("addContainer: %v pathElem: %s val: %v\n", elems, e, val)
	if reflect.TypeOf((d)) == nil {
		d = make(map[string]interface{})
	}
	switch dd := d.(type) {
	case map[string]interface{}:
		// add the container
		dd[e], err = addData(dd[e], elems[1:], val)
		return d, err
	default:
		return nil, errors.New("addListLastValue JSON unexpected data structure")
	}
	//}

}

func addList(d interface{}, e string, k map[string]string, elems []*gnmi.PathElem, val *gnmi.TypedValue) (interface{}, error) {
	var err error
	//fmt.Printf("addList pathElem: %s, key: %v d: %v\n", e, k, d)
	// lean approach -> since we know the query should return paths that match the original query we can assume we match the path

	// initialize the data
	if reflect.TypeOf((d)) == nil {
		d = make(map[string]interface{})
	}
	switch dd := d.(type) {
	case map[string]interface{}:
		// initialize the data
		if _, ok := dd[e]; !ok {
			dd[e] = make([]interface{}, 0)
		}
		switch l := dd[e].(type) {
		case []interface{}:
			// check if the list entry exists
			for i, le := range l {
				// initialize the data
				if reflect.TypeOf((le)) == nil {
					le = make(map[string]interface{})
				}
				found := true
				switch dd := le.(type) {
				case map[string]interface{}:
					for keyName, keyValue := range k {
						if dd[keyName] != keyValue {
							found = false
						}
					}
					if found {
						// augment the list
						l[i], err = addData(dd, elems[1:], val)
						if err != nil {
							return nil, err
						}
						return d, err
					}
				}
			}
			// list entry not found, add a list entry
			de := make(map[string]interface{})
			for keyName, keyValue := range k {
				de[keyName] = keyValue
			}
			// augment the list
			x, err := addData(de, elems[1:], val)
			if err != nil {
				return nil, err
			}
			// add the list entry to the list
			dd[e] = append(l, x)
			return d, nil
		default:
			return nil, errors.New("list last value JSON unexpected data structure")
		}
	default:
		return nil, errors.New("list last value JSON unexpected data structure")
	}
}

func addContainerValue(d interface{}, e string, val *gnmi.TypedValue) (interface{}, error) {
	var err error
	// check if the data was initialized
	if reflect.TypeOf((d)) == nil {
		d = make(map[string]interface{})
	}
	switch dd := d.(type) {
	case map[string]interface{}:
		// add the value to the element
		dd[e], err = yparser.GetValue(val)
		return d, err
	default:
		// we should never end up here
		return nil, errors.New("container last value JSON unexpected data structure")
	}
}

func addListValue(d interface{}, e string, k map[string]string, val *gnmi.TypedValue) (interface{}, error) {
	// initialize the data
	if reflect.TypeOf((d)) == nil {
		d = make(map[string]interface{})
	}
	switch dd := d.(type) {
	case map[string]interface{}:
		// initialize the data
		if _, ok := dd[e]; !ok {
			dd[e] = make([]interface{}, 0)
		}
		// create a container and initialize with keyNames/keyValues and value
		de := make(map[string]interface{})
		// add value
		v, err := yparser.GetValue(val)
		if err != nil {
			return nil, err
		}
		switch vv := v.(type) {
		case map[string]interface{}:
			for k, v := range vv {
				de[k] = v
			}
		default:
		}

		// add keyNames/keyValues
		for keyName, keyValue := range k {
			de[keyName] = keyValue
		}
		// add the data to the list
		switch l := dd[e].(type) {
		case []interface{}:
			dd[e] = append(l, de)
		default:
			return nil, errors.New("list last value JSON unexpected data structure")
		}
	}
	return d, nil
}
