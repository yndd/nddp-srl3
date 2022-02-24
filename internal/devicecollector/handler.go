package devicecollector

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

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
	pbRootPath = &gnmi.Path{}
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
	for _, path := range delPaths {
		xpath := yparser.GnmiPath2XPath(path, true)
		resourceName, err := c.findManagedResource(xpath, resourceList)
		if err != nil {
			return err
		}

		// clean the path for now to remove the module information from the pathElem
		for _, pe := range path.GetElem() {
			pe.Name = strings.Split(pe.Name, ":")[len(strings.Split(pe.Name, ":"))-1]
		}

		// if a default is enabled on the path we should revert to default
		def := c.deviceSchema.GetPathDefault(path)
		var n *gnmi.Notification
		c.log.Debug("collector config delete", "path", xpath, "default", def)
		if def != "" {
			d, err := json.Marshal(def)
			if err != nil {
				return err
			}
			// if the data is empty, there is no need for an update
			if string(d) == "null" {
				return nil
			}

			n = &gnmi.Notification{
				Timestamp: time.Now().UnixNano(),
				Prefix:    &gnmi.Path{Target: crDeviceName},
				Update: []*gnmi.Update{
					{
						Path: path,
						Val: &gnmi.TypedValue{
							Value: &gnmi.TypedValue_JsonIetfVal{
								JsonIetfVal: bytes.Trim(d, " \r\n\t"),
							},
						},
					},
				},
			}
		} else {
			n = &gnmi.Notification{
				Timestamp: time.Now().UnixNano(),
				Prefix:    &gnmi.Path{Target: crDeviceName},
				Delete:    []*gnmi.Path{path},
			}
		}
		// update the cache with the latest config from the device
		if err := c.cache.GetCache().GnmiUpdate(crDeviceName, n); err != nil {
			c.log.Debug("handle target update", "error", err, "Path", xpath)
			return errors.New("cache update failed")
		}

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

		/*
			// clean the path for now to remove the module information from the pathElem
			for _, pe := range upd.GetPath().GetElem() {
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

			switch upd.GetVal().Value.(type) {
			case *gnmi.TypedValue_JsonIetfVal:
				jsondata := upd.GetVal().GetJsonIetfVal()
				var v interface{}
				if len(jsondata) != 0 {
					err := json.Unmarshal(jsondata, &v)
					if err != nil {
						return err
					}
				}


				//	if strings.Contains(string(jsondata), "routed") {
				//		//fmt.Printf("type of data1: %v %v\n", string(jsondata), reflect.TypeOf(jsondata))
				//		fmt.Printf("type of data2: %v %v\n", string(jsondata), reflect.TypeOf(v))
				//	}
				switch vv := v.(type) {
				case map[string]interface{}:
					vv = yparser.CleanConfig(vv)
					b, err := json.Marshal(vv)
					if err != nil {
						return err
					}
					//fmt.Printf("string cleaned: %s\n", string(b))
					upd.Val = &gnmi.TypedValue{
						Value: &gnmi.TypedValue_JsonIetfVal{
							JsonIetfVal: bytes.Trim(b, " \r\n\t"),
						},
					}
				case string:
					// for string values there can be also a header in the values e.g. type, Value: srl_nokia-network-instance:ip-vrf
					if !strings.Contains(vv, "::") {
						// if there are more ":" in the string it is likely an esi or mac address
						if len(strings.Split(vv, ":")) <= 2 {
							//fmt.Printf("string to be cleaned: %s\n", vv)
							vv = strings.Split(vv, ":")[len(strings.Split(vv, ":"))-1]
							b, err := json.Marshal(vv)
							if err != nil {
								return err
							}
							//fmt.Printf("string cleaned: %s\n", string(b))
							upd.Val = &gnmi.TypedValue{
								Value: &gnmi.TypedValue_JsonIetfVal{
									JsonIetfVal: bytes.Trim(b, " \r\n\t"),
								},
							}
						}
					}
				}
			}
		*/

		/*
			// Validate if the path has a key using the device schema
			// Used to allow insertion of an empty container
			keys := c.deviceSchema.GetKeys(upd.GetPath())
			hashKey := false
			if len(keys) != 0 {
				hashKey = true
			}
			crDeviceName := shared.GetCrDeviceName(c.namespace, c.target.Config.Name)
			n, err := c.cache.GetCache().GetNotificationFromUpdate(&gnmi.Path{Target: crDeviceName}, upd, hashKey)
			if err != nil {
				return err
			}

			// default handling
			defaults := c.deviceSchema.GetPathDefaults(upd.GetPath())
			for pathElemName, defValue := range defaults {
				c.log.Debug("collector config update defaults", "pathElemName", pathElemName, "defValue", defValue, "path", yparser.GnmiPath2XPath(upd.GetPath(), true))

				d, err := json.Marshal(defValue)
				if err != nil {
					return err
				}
				// if the data is empty, there is no need for an update
				if string(d) == "null" {
					return nil
				}

				// check if the element exists in the original notification
				// if not we add the default, if it is there we avoid adding the default
				found := false
				for _, nu := range n.GetUpdate() {
					if nu.GetPath().GetElem()[len(nu.GetPath().GetElem())-1].GetName() == pathElemName {
						found = true
					}
				}
				if !found {
					newPath := yparser.DeepCopyGnmiPath(upd.GetPath())
					newPath.Elem = append(newPath.GetElem(), &gnmi.PathElem{Name: pathElemName})
					u := &gnmi.Update{
						Path: newPath,
						Val: &gnmi.TypedValue{
							Value: &gnmi.TypedValue_JsonIetfVal{
								JsonIetfVal: bytes.Trim(d, " \r\n\t"),
							},
						},
					}
					n.Update = append(n.GetUpdate(), u)
				}

			}
		*/

		//fullPath := gnmiFullPath(prefix, u.GetPath())
		fullPath := u.GetPath()
		val := u.GetVal()

		emptyNode, _, err := ytypes.GetOrCreateNode(m.SchemaTreeRoot, m.NewRootValue(), fullPath)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("path %v is not found in the config structure", fullPath))
		}
		var nodeVal interface{}
		nodeStruct, ok := emptyNode.(ygot.ValidatedGoStruct)
		if ok {
			if err := m.JsonUnmarshaler(val.GetJsonIetfVal(), nodeStruct); err != nil {
				return errors.Wrap(err, "unmarshaling json data to config struct fails")
			}
			if err := nodeStruct.Validate(); err != nil {
				return errors.Wrap(err, "config data validation fails")
			}
			var err error
			if nodeVal, err = ygot.ConstructIETFJSON(nodeStruct, &ygot.RFC7951JSONConfig{}); err != nil {
				return errors.Wrap(err, "error in constructing IETF JSON tree from config struct:")
			}
		} else {
			var err error
			if nodeVal, err = value.ToScalar(val); err != nil {
				return errors.Wrap(err, "cannot convert leaf node to scalar type")
			}
		}

		// Update json tree of the device config.
		var curNode interface{} = jsonTree
		schema := m.SchemaTreeRoot
		for i, elem := range fullPath.Elem {
			switch node := curNode.(type) {
			case map[string]interface{}:
				// Set node value.
				if i == len(fullPath.Elem)-1 {
					if elem.GetKey() == nil {
						if grpcStatusError := setPathWithoutAttribute(gnmi.UpdateResult_UPDATE, node, elem, nodeVal); grpcStatusError != nil {
							return err
						}
						break
					}
					if grpcStatusError := setPathWithAttribute(gnmi.UpdateResult_UPDATE, node, elem, nodeVal); grpcStatusError != nil {
						return err
					}
					break
				}

				if curNode, schema = getChildNode(node, schema, elem, true); curNode == nil {
					return errors.Wrap(err, fmt.Sprintf("path elem not found: %v", elem))
				}
			case []interface{}:
				return errors.Wrap(err, fmt.Sprintf("incompatible path elem: %v", elem))
			default:
				return errors.Wrap(err, fmt.Sprintf("wrong node type: %T", curNode))
			}
		}
		if reflect.DeepEqual(fullPath, pbRootPath) { // Replace/Update root.
			//if op == gnmi.UpdateResult_UPDATE {
			return errors.Wrap(err, "update the root of config tree is unsupported")
			//}
			/*
				nodeValAsTree, ok := nodeVal.(map[string]interface{})
				if !ok {
					return errors.Wrap(err, fmt.Sprintf("expect a tree to replace the root, got a scalar value: %T", nodeVal))
				}
				for k := range jsonTree {
					delete(jsonTree, k)
				}
				for k, v := range nodeValAsTree {
					jsonTree[k] = v
				}
			*/
		}
		if err := c.UpdateValidatedConfig(crDeviceName, jsonTree); err != nil {
			return err
		}

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
	goStruct, err := m.NewConfigStruct(jsonDump)
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
