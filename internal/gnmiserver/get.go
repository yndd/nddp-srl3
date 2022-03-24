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
	"github.com/yndd/ndd-yang/pkg/yparser"
	"github.com/yndd/nddp-srl3/internal/shared"

	//systemv1alpha1 "github.com/yndd/nddp-system/apis/system/v1alpha1"
	"github.com/yndd/nddp-system/pkg/failedmsg"
	"github.com/yndd/nddp-system/pkg/gvkresource"
	"github.com/yndd/nddp-system/pkg/ygotnddp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *GnmiServerImpl) Get(ctx context.Context, req *gnmi.GetRequest) (*gnmi.GetResponse, error) {
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

func (s *GnmiServerImpl) HandleGet(req *gnmi.GetRequest) ([]*gnmi.Notification, error) {
	prefix := req.GetPrefix()
	//origin := req.GetPrefix().GetOrigin()

	target := req.GetPrefix().GetTarget()
	ygotCache := false
	if strings.HasPrefix(target, "ygot") {
		ygotCache = true
		target = strings.ReplaceAll(target, "ygot.", "")
	}

	if !s.cache.HasTarget(target) {
		return nil, status.Errorf(codes.Unavailable, "cache not ready")
	}

	var systemTarget string
	// if the request is for the system resources per leaf we take the target/crDeviceName iso
	// adding the system part
	if strings.HasPrefix(target, shared.SystemNamespace) {
		systemTarget = target
	} else {
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
				switch resource.Status {
				case ygotnddp.NddpSystem_ResourceStatus_PENDING:
					return nil, status.Error(codes.AlreadyExists, "")
				case ygotnddp.NddpSystem_ResourceStatus_FAILED:
					m := &failedmsg.Message{
						Spec: *resource.Spec,
						Msg:  *resource.Reason,
					}
					errMsgSpec, err := m.Marshal()
					if err != nil {
						return nil, status.Error(codes.Internal, err.Error())
					}
					return nil, status.Error(codes.FailedPrecondition, string(errMsgSpec))
				}
			}
		}
	}

	s.log.Debug("Get Cache", "ygot", ygotCache, "exists", exists)

	for i, path := range req.GetPath() {
		if ygotCache {
			fullPath := path
			if fullPath.GetElem() == nil && fullPath.GetElem() != nil {
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

			/*
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
			*/

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
			return nil, status.Error(codes.Unimplemented, "")
		}

	}
	if exists {
		// resource exists
		return notifications, nil
	}
	// the resource does not exists
	return notifications, status.Error(codes.NotFound, "resource does not exist")
}
