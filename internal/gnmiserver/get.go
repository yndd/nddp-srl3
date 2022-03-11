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

package gnmiserver

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/openconfig/gnmi/value"
	"github.com/openconfig/ygot/util"
	"github.com/openconfig/ygot/ygot"
	"github.com/openconfig/ygot/ytypes"
	"github.com/pkg/errors"
	"github.com/yndd/ndd-yang/pkg/yentry"
	"github.com/yndd/ndd-yang/pkg/yparser"
	"github.com/yndd/nddp-srl3/internal/shared"

	//systemv1alpha1 "github.com/yndd/nddp-system/apis/system/v1alpha1"
	"github.com/yndd/nddp-system/pkg/gvkresource"
	"github.com/yndd/nddp-system/pkg/ygotnddp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) Get(ctx context.Context, req *gnmi.GetRequest) (*gnmi.GetResponse, error) {
	ok := s.unaryRPCsem.TryAcquire(1)
	if !ok {
		return nil, status.Errorf(codes.ResourceExhausted, "max number of Unary RPC reached")
	}
	defer s.unaryRPCsem.Release(1)

	prefix := req.GetPrefix()
	target := ""
	if prefix != nil {
		target = prefix.Target
	}
	log := s.log.WithValues("Type", req.GetType())
	if req.GetPath() != nil {
		log.Debug("Get...", "Target", target, "Path", yparser.GnmiPath2XPath(req.GetPath()[0], true))
	} else {
		log.Debug("Get...", "Target", target)
	}

	// We dont act upon the error here, but we pass it on the response with updates
	ns, err := s.HandleGet(req)
	return &gnmi.GetResponse{
		Notification: ns,
	}, err
}

func (s *server) HandleGet(req *gnmi.GetRequest) ([]*gnmi.Notification, error) {
	//var err error
	updates := make([]*gnmi.Update, 0)

	prefix := req.GetPrefix()
	//origin := req.GetPrefix().GetOrigin()

	target := req.GetPrefix().GetTarget()
	ygotCache := false
	if strings.HasPrefix(target, "ygot") {
		ygotCache = true
		target = strings.ReplaceAll(target, "ygot.", "")
	}

	if !s.cache.GetCache().GetCache().HasTarget(target) {
		return nil, status.Errorf(codes.Unavailable, "cache not ready")
	}

	var schema *yentry.Entry
	var systemTarget string
	// if the request is for the system resources per leaf we take the target/crDeviceName iso
	// adding the system part
	if strings.HasPrefix(target, shared.SystemNamespace) {
		schema = s.nddpSchema
		systemTarget = target
	} else {
		schema = s.deviceSchema
		systemTarget = shared.GetCrSystemDeviceName(target)
	}

	var goStruct ygot.GoStruct
	g := s.cache.GetValidatedGoStruct(target)
	if g == nil {
		goStruct = g
	} else {
		var err error
		goStruct, err = ygot.DeepCopy(g)
		if err != nil {
			return nil, err
		}
	}
	m := s.cache.GetModel(target)

	paths := req.GetPath()
	notifications := make([]*gnmi.Notification, len(paths))

	ts := time.Now().UnixNano()

	// if the extension is set we check the resourcelist
	// this is needed for the device driver to know when a create should be triggered
	exists := true
	if len(req.GetExtension()) > 0 {
		// if the cache is exhausted we need to backoff
		exhausted, err := s.cache.GetSystemExhausted(systemTarget)
		if err != nil {
			return nil, err
		}
		if *exhausted != 0 {
			return nil, status.Errorf(codes.ResourceExhausted, "device exhausted")
		}
		gvkName := req.GetExtension()[0].GetRegisteredExt().GetMsg()
		if string(gvkName) == gvkresource.Operation_GetResourceNameFromPath {
			// procedure to get resource name
			/*
				updates, err := s.getResourceName(systemTarget, req.GetPath()[0])
				if err != nil {
					return nil, err
				}
				notifications[0] = &gnmi.Notification{
					Timestamp: ts,
					Prefix:    prefix,
					Update:    updates,
				}
				return notifications, nil
			*/
			return nil, nil
		} else {
			resource, err := s.cache.GetSystemResource(systemTarget, string(gvkName))
			if err != nil {
				return nil, status.Error(codes.Internal, err.Error())
			}
			if resource == nil {
				exists = false
			} else {
				if resource.Action != ygotnddp.NddpSystem_ResourceAction_DELETE {
					switch resource.Status {
					case ygotnddp.NddpSystem_ResourceStatus_PENDING:
						// the action did not complete so far
						return nil, status.Error(codes.AlreadyExists, "")
					}
				} else {
					switch resource.Status {
					case ygotnddp.NddpSystem_ResourceStatus_FAILED:
						// resource exists, but failed, return the spec data
						notifications[0] = &gnmi.Notification{
							Timestamp: ts,
							Prefix:    prefix,
							Update: []*gnmi.Update{
								{
									Path: &gnmi.Path{},
									Val:  &gnmi.TypedValue{Value: &gnmi.TypedValue_JsonVal{JsonVal: []byte(*resource.Spec)}},
								},
							},
						}
						return notifications, status.Error(codes.FailedPrecondition, "resource exist, but failed")
					case ygotnddp.NddpSystem_ResourceStatus_PENDING:
						// the action did not complete so far
						return nil, status.Error(codes.AlreadyExists, "")
					}
				}
			}
		}
	}

	s.log.Debug("Get Cache", "ygot", ygotCache, "exists", exists)

	for i, path := range req.GetPath() {
		if ygotCache {
			fullPath := path
			if fullPath.GetElem() == nil && fullPath.GetElement() != nil {
				return nil, status.Error(codes.Unimplemented, "deprecated path element type is unsupported")
			}

			nodes, err := ytypes.GetNode(m.SchemaTreeRoot, goStruct, fullPath,
				&ytypes.GetPartialKeyMatch{},
				&ytypes.GetHandleWildcards{},
			)
			if len(nodes) == 0 || err != nil || util.IsValueNil(nodes[0].Data) {
				return nil, status.Errorf(codes.NotFound, "path %v not found: %v", fullPath, err)
			}
			node := nodes[0].Data

			for _, node := range nodes {
				nodeStruct, ok := node.Data.(ygot.GoStruct)
				// return a leaf node
				if !ok {
					var val *gnmi.TypedValue
					switch kind := reflect.ValueOf(node).Kind(); kind {
					case reflect.Ptr, reflect.Interface:
						var err error
						val, err = value.FromScalar(reflect.ValueOf(node).Elem().Interface())
						if err != nil {
							msg := fmt.Sprintf("leaf node %v does not contain a scalar type value: %v", path, err)
							s.log.Debug("Error", "msg", msg)
							return nil, status.Error(codes.Internal, msg)
						}
					case reflect.Int64:
						enumMap, ok := m.EnumData[reflect.TypeOf(node).Name()]
						if !ok {
							return nil, status.Error(codes.Internal, "not a GoStruct enumeration type")
						}
						val = &gnmi.TypedValue{
							Value: &gnmi.TypedValue_StringVal{
								StringVal: enumMap[reflect.ValueOf(node).Int()].Name,
							},
						}
					default:
						return nil, status.Errorf(codes.Internal, "unexpected kind of leaf node type: %v %v", node, kind)
					}
					fmt.Println(val)
				}
				// Return IETF JSON by default.
				jsonEncoder := func() (map[string]interface{}, error) {
					return ygot.ConstructIETFJSON(nodeStruct, &ygot.RFC7951JSONConfig{AppendModuleName: true})
				}
				jsonType := "IETF with moduleName"
				//buildUpdate := func(b []byte) *gnmi.Update {
				//	return &gnmi.Update{Path: path, Val: &gnmi.TypedValue{Value: &gnmi.TypedValue_JsonIetfVal{JsonIetfVal: b}}}
				//}

				if req.GetEncoding() == gnmi.Encoding_JSON {
					jsonEncoder = func() (map[string]interface{}, error) {
						return ygot.ConstructIETFJSON(nodeStruct, &ygot.RFC7951JSONConfig{})
					}
					jsonType = "IETF without moduleName"
					//buildUpdate = func(b []byte) *gnmi.Update {
					//	return &gnmi.Update{Path: path, Val: &gnmi.TypedValue{Value: &gnmi.TypedValue_JsonVal{JsonVal: b}}}
					//}
				}

				jsonTree, err := jsonEncoder()
				if err != nil {
					msg := fmt.Sprintf("error in constructing %s JSON tree from requested node: %v", jsonType, err)
					s.log.Debug("Error", "msg", msg)
					return nil, status.Error(codes.Internal, msg)
				}

				jsonDump, err := json.Marshal(jsonTree)
				if err != nil {
					msg := fmt.Sprintf("error in marshaling %s JSON tree to bytes: %v", jsonType, err)
					s.log.Debug("Error", "msg", msg)
					return nil, status.Error(codes.Internal, msg)
				}
				fmt.Println(string(jsonDump))
			}

			nodeStruct, ok := node.(ygot.GoStruct)
			// Return leaf node.
			if !ok {
				var val *gnmi.TypedValue
				switch kind := reflect.ValueOf(node).Kind(); kind {
				case reflect.Ptr, reflect.Interface:
					var err error
					val, err = value.FromScalar(reflect.ValueOf(node).Elem().Interface())
					if err != nil {
						msg := fmt.Sprintf("leaf node %v does not contain a scalar type value: %v", path, err)
						s.log.Debug("Error", "msg", msg)
						return nil, status.Error(codes.Internal, msg)
					}
				case reflect.Int64:
					enumMap, ok := m.EnumData[reflect.TypeOf(node).Name()]
					if !ok {
						return nil, status.Error(codes.Internal, "not a GoStruct enumeration type")
					}
					val = &gnmi.TypedValue{
						Value: &gnmi.TypedValue_StringVal{
							StringVal: enumMap[reflect.ValueOf(node).Int()].Name,
						},
					}
				default:
					return nil, status.Errorf(codes.Internal, "unexpected kind of leaf node type: %v %v", node, kind)
				}

				update := &gnmi.Update{Path: path, Val: val}
				notifications[i] = &gnmi.Notification{
					Timestamp: ts,
					Prefix:    prefix,
					Update:    []*gnmi.Update{update},
				}
				continue
			}
			// Return IETF JSON by default.
			jsonEncoder := func() (map[string]interface{}, error) {
				return ygot.ConstructIETFJSON(nodeStruct, &ygot.RFC7951JSONConfig{AppendModuleName: true})
			}
			jsonType := "IETF with moduleName"
			buildUpdate := func(b []byte) *gnmi.Update {
				return &gnmi.Update{Path: path, Val: &gnmi.TypedValue{Value: &gnmi.TypedValue_JsonIetfVal{JsonIetfVal: b}}}
			}

			if req.GetEncoding() == gnmi.Encoding_JSON {
				jsonEncoder = func() (map[string]interface{}, error) {
					return ygot.ConstructIETFJSON(nodeStruct, &ygot.RFC7951JSONConfig{})
				}
				jsonType = "IETF without moduleName"
				buildUpdate = func(b []byte) *gnmi.Update {
					return &gnmi.Update{Path: path, Val: &gnmi.TypedValue{Value: &gnmi.TypedValue_JsonVal{JsonVal: b}}}
				}
			}

			jsonTree, err := jsonEncoder()
			if err != nil {
				msg := fmt.Sprintf("error in constructing %s JSON tree from requested node: %v", jsonType, err)
				s.log.Debug("Error", "msg", msg)
				return nil, status.Error(codes.Internal, msg)
			}

			jsonDump, err := json.Marshal(jsonTree)
			if err != nil {
				msg := fmt.Sprintf("error in marshaling %s JSON tree to bytes: %v", jsonType, err)
				s.log.Debug("Error", "msg", msg)
				return nil, status.Error(codes.Internal, msg)
			}

			update := buildUpdate(jsonDump)
			notifications[i] = &gnmi.Notification{
				Timestamp: ts,
				Prefix:    prefix,
				Update:    []*gnmi.Update{update},
			}
		} else {
			x, err := s.cache.GetCache().GetJson(prefix.GetTarget(), prefix, path, schema)
			if err != nil {
				return nil, status.Error(codes.Internal, err.Error())
			}
			if updates, err = appendUpdateResponse(x, path, updates); err != nil {
				return nil, status.Error(codes.Internal, err.Error())
			}
			notifications[i] = &gnmi.Notification{
				Timestamp: ts,
				Prefix:    prefix,
				Update:    updates,
			}
		}

	}
	if exists {
		// resource exists
		return notifications, nil
	}
	// the resource does not exists
	return notifications, status.Error(codes.NotFound, "resource does not exist")
}

func appendUpdateResponse(data interface{}, path *gnmi.Path, updates []*gnmi.Update) ([]*gnmi.Update, error) {
	var err error
	var d []byte
	//fmt.Printf("data1: %v\n", data)
	if data != nil {
		d, err = json.Marshal(data)
		if err != nil {
			return nil, err
		}
	}

	//fmt.Printf("data2: %v\n", string(d))

	upd := &gnmi.Update{
		Path: path,
		Val:  &gnmi.TypedValue{Value: &gnmi.TypedValue_JsonVal{JsonVal: d}},
	}
	updates = append(updates, upd)
	return updates, nil
}

/*
func (s *server) getResourceList(crSystemDeviceName string) ([]*systemv1alpha1.Gvk, error) {
	rl, err := s.cache.GetCache().GetJson(crSystemDeviceName,
		&gnmi.Path{Target: crSystemDeviceName},
		&gnmi.Path{Elem: []*gnmi.PathElem{{Name: "gvk"}}},
		s.nddpSchema)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return gvkresource.GetResourceList(rl)
}
*/
/*
func (s *server) getResource(crSystemDeviceName, gvkName string) (*systemv1alpha1.Gvk, error) {
	rl, err := s.cache.GetCache().GetJson(crSystemDeviceName,
		&gnmi.Path{Target: crSystemDeviceName},
		&gnmi.Path{Elem: []*gnmi.PathElem{
			{Name: "gvk", Key: map[string]string{"name": gvkName}},
		}},
		s.nddpSchema)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return gvkresource.GetResource(rl)
}
*/
/*
func (s *server) getResourceName(crSystemDeviceName string, reqpath *gnmi.Path) ([]*gnmi.Update, error) {
	// provide a string from the gnmi Path, we expect a single path in the GetRequest
	reqPath := yparser.GnmiPath2XPath(reqpath, true)
	// initialize the variables which are used to keep track of the matched strings
	matchedResourceName := ""
	matchedResourcePath := ""
	// loop over the resourceList
	resourceList, err := s.getResourceList(crSystemDeviceName)
	if err != nil {
		return nil, err
	}
	for _, resource := range resourceList {
		if resource != nil {
			//s.log.Debug("K8sResource GetResourceName Match", "resource", resource)
			// check if the string is contained in the path
			// TOOOODOOOO UPDATE TO the new Paths procesure
			if strings.Contains(reqPath, resource.Paths[0]) {
				// if there is a better match we use the better match
				if len(resource.Paths) > len(matchedResourcePath) {
					matchedResourcePath = resource.Paths[0]
					matchedResourceName = resource.Name
				}
			}
		}
	}
	//s.log.Debug("K8sResource GetResourceName Match", "ResourceName", matchedResourceName)

	d, err := json.Marshal(nddv1.ResourceName{
		Name: matchedResourceName,
	})
	if err != nil {
		return nil, err
	}
	updates := make([]*gnmi.Update, 0)
	upd := &gnmi.Update{
		Path: reqpath,
		Val:  &gnmi.TypedValue{Value: &gnmi.TypedValue_JsonVal{JsonVal: d}},
	}
	updates = append(updates, upd)
	return updates, nil
}
*/
/*
func (s *server) getSpecdata(crSystemDeviceName string, resource *systemv1alpha1.Gvk) (interface{}, error) {
	x1, err := s.cache.GetCache().GetJson(
		crSystemDeviceName,
		&gnmi.Path{Target: crSystemDeviceName},
		&gnmi.Path{
			Elem: []*gnmi.PathElem{
				{Name: "gvk", Key: map[string]string{"name": resource.Name}},
			},
		},
		s.nddpSchema)
	if err != nil {
		return nil, err
	}
	// remove the rootPath data
	//	rootPath, err := xpath.ToGNMIPath(resource.Rootpath)
	//	if err != nil {
	//		return nil, err
	//	}
	//	switch x := x1.(type) {
	//	case map[string]interface{}:
	//		x1 = x["data"]
	//		x1 = getDataFromRootPath(rootPath, x1)
	//	}
	return x1, nil
}
*/

func getDataFromRootPath(path *gnmi.Path, x1 interface{}) interface{} {
	//fmt.Printf("gnmiserver getDataFromRootPath: %s, data: %v\n", yparser.GnmiPath2XPath(path, true), x1)
	p := yparser.DeepCopyGnmiPath(path)
	if len(p.GetElem()) > 0 {
		hasKey := false
		if len(p.GetElem()[0].GetKey()) > 0 {
			hasKey = true
		}
		switch x := x1.(type) {
		case map[string]interface{}:
			for k, v := range x {
				if k == p.GetElem()[0].GetName() {
					// when the spec data rootpath last element has a key
					// we need to return the first element of the list
					if hasKey {
						switch x := v.(type) {
						case []interface{}:
							x1 = x[0]
						}
					} else {
						x1 = x[p.GetElem()[0].GetName()]
					}
				}
			}

		}
		p.Elem = p.Elem[1:]
		x1 = getDataFromRootPath(p, x1)
	}
	return x1
}

func (s *server) getExhausted(crSystemDeviceName string) (int64, error) {

	path := &gnmi.Path{
		Elem: []*gnmi.PathElem{{Name: "exhausted"}},
	}

	n, err := s.cache.GetCache().Query(crSystemDeviceName, &gnmi.Path{Target: crSystemDeviceName}, path)
	if err != nil {
		return 0, err
	}
	if n != nil {
		for _, u := range n.GetUpdate() {
			val, err := yparser.GetValue(u.GetVal())
			if err != nil {
				return 0, err
			}
			switch v := val.(type) {
			case int64:
				return v, nil
			}
		}
		return 0, errors.New("unknown type in cache")
	}

	return 0, nil
}

/*
func (s *server) getTransaction(crSystemDeviceName, transactionName string) (*systemv1alpha1.Transaction, error) {
	path := &gnmi.Path{
		Elem: []*gnmi.PathElem{{Name: "transaction", Key: map[string]string{"name": transactionName}}},
	}

	n, err := s.cache.GetCache().Query(crSystemDeviceName, &gnmi.Path{Target: crSystemDeviceName}, path)
	if err != nil {
		return nil, err
	}
	resp := &gnmi.GetResponse{
		Notification: []*gnmi.Notification{n},
	}
	return transaction.GetTransactionFromGnmiResponse(resp)
}
*/
/*
func (s *server) deleteResource(crSystemDeviceName, resourceGvkName string) error {
	n := &gnmi.Notification{
		Timestamp: time.Now().UnixNano(),
		Prefix:    &gnmi.Path{Target: crSystemDeviceName},
		Delete: []*gnmi.Path{
			{
				Elem: []*gnmi.PathElem{{Name: "gvk", Key: map[string]string{"name": resourceGvkName}}},
			},
		},
	}

	if err := s.cache.GetCache().GnmiUpdate(crSystemDeviceName, n); err != nil {
		return errors.Wrap(err, "cache delete failed")
	}
	return nil
}
*/
