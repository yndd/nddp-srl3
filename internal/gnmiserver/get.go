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
	"github.com/openconfig/gnmi/proto/gnmi_ext"
	"github.com/openconfig/gnmi/value"
	"github.com/openconfig/ygot/util"
	"github.com/openconfig/ygot/ygot"
	"github.com/openconfig/ygot/ytypes"
	"github.com/yndd/ndd-yang/pkg/yparser"
	"github.com/yndd/nddp-srl3/internal/cache"
	"github.com/yndd/nddp-srl3/internal/model"
	"github.com/yndd/nddp-srl3/internal/shared"

	//systemv1alpha1 "github.com/yndd/nddp-system/apis/system/v1alpha1"
	"github.com/yndd/nddp-system/pkg/failedmsg"
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
	prefix := req.GetPrefix()
	target := prefix.GetTarget()

	if !strings.HasPrefix(target, "ygot.") {
		return nil, status.Error(codes.Unimplemented, "")
	}
	// we make sure that ygot procedures are used right now ...
	// we will need to get rid of this later
	target = strings.ReplaceAll(target, "ygot.", "")

	if !s.cache.GetCache().GetCache().HasTarget(target) {
		return nil, status.Errorf(codes.Unavailable, "cache not ready")
	}

	// if the request is for the system resources per leaf we take the target/crDeviceName iso
	// adding the system part
	systemTarget := shared.GetCrSystemDeviceName(target)

	extensions := req.GetExtension()
	if len(extensions) > 0 {
		// if the extension is set we check the resourcelist
		// this is needed for the device driver to know when a create should be triggered
		return nil, processExtension(systemTarget, s.cache, extensions)
	}

	var goStruct ygot.GoStruct
	goStruct = s.cache.GetValidatedGoStruct(target) // no DeepCopy required, since we get a deepcopy already

	notifications := make([]*gnmi.Notification, len(req.GetPath()))
	ts := time.Now().UnixNano()

	model := s.cache.GetModel(target)
	// process all the paths from the given request
	for i, path := range req.GetPath() {
		fullPath := path
		if fullPath.GetElem() == nil && fullPath.GetElem() != nil {
			return nil, status.Error(codes.Unimplemented, "deprecated path element type is unsupported")
		}

		nodes, err := ytypes.GetNode(model.SchemaTreeRoot, goStruct, fullPath,
			&ytypes.GetPartialKeyMatch{},
			&ytypes.GetHandleWildcards{},
		)
		if len(nodes) == 0 || err != nil || util.IsValueNil(nodes[0].Data) {
			return nil, status.Errorf(codes.NotFound, "path %v not found: %v", fullPath, err)
		}

		// with wildcards allowed, we might get multiple entries per path query
		// hence the updates are stored in a slice.
		updates := []*gnmi.Update{}

		// generate updates for all the retrieved nodes
		for _, entry := range nodes {
			var update *gnmi.Update
			node := entry.Data
			nodeStruct, isYgotStruct := node.(ygot.GoStruct)
			// if not ok, this mut be a leafnode, meaning a scalar value instead of any ygot struct
			if !isYgotStruct {
				update, err = createLeafNodeUpdate(node, fullPath, model)
			} else {
				// process ygot structs
				update, err = createYgotStructNodeUpdate(nodeStruct, path, req.GetEncoding())
			}
			if err != nil {
				return nil, err
			}
			updates = append(updates, update)
		}
		// generate the notification on a per path basis, as defined by the RFC
		notifications[i] = &gnmi.Notification{
			Timestamp: ts,
			Prefix:    prefix,
			Update:    updates,
		}
	}
	return notifications, nil
}

// processExtension the cache returned result might have gnmi extensions attached. Here these extension are being processsed
func processExtension(systemTarget string, cache cache.Cache, extensions []*gnmi_ext.Extension) error {
	// if the cache is exhausted we need to backoff
	exhausted, err := cache.GetSystemExhausted(systemTarget)
	if err != nil {
		return err
	}
	if *exhausted != 0 {
		return status.Errorf(codes.ResourceExhausted, "device exhausted")
	}
	gvkName := extensions[0].GetRegisteredExt().GetMsg()
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
		return nil
	}

	resource, err := cache.GetSystemResource(systemTarget, string(gvkName))
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	if resource == nil {
		return status.Error(codes.NotFound, "resource does not exist")
	}
	switch resource.Status {
	case ygotnddp.NddpSystem_ResourceStatus_PENDING:
		return status.Error(codes.AlreadyExists, "")
	case ygotnddp.NddpSystem_ResourceStatus_FAILED:
		m := &failedmsg.Message{
			Spec: *resource.Spec,
			Msg:  *resource.Reason,
		}
		errMsgSpec, err := m.Marshal()
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}
		return status.Error(codes.FailedPrecondition, string(errMsgSpec))
	}
	return nil
}

// createLeafNodeUpdate processes the list of returned nodes from the cache, which are Leaf Nodes, carrying terminal values
// rather then ygot struct kind of data. From these, the resulting gnmi.update stucts are being populated.
func createLeafNodeUpdate(node interface{}, path *gnmi.Path, model *model.Model) (*gnmi.Update, error) {
	var val *gnmi.TypedValue
	switch kind := reflect.ValueOf(node).Kind(); kind {
	case reflect.Ptr, reflect.Interface:
		var err error
		val, err = value.FromScalar(reflect.ValueOf(node).Elem().Interface())
		if err != nil {
			msg := fmt.Sprintf("leaf node %v does not contain a scalar type value: %v", path, err)
			return nil, status.Error(codes.Internal, msg)
		}
	case reflect.Int64:
		enumMap, ok := model.EnumData[reflect.TypeOf(node).Name()]
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
	return update, nil
}

// createYgotStructNodeUpdate generated update messages from ygot structs.
func createYgotStructNodeUpdate(nodeStruct ygot.GoStruct, path *gnmi.Path, requestedEncoding gnmi.Encoding) (*gnmi.Update, error) {
	// take care of encoding
	// default to JSON IETF, other option is plain JSON
	var encoder Encoder
	switch requestedEncoding {
	case gnmi.Encoding_JSON:
		encoder = &JSONEncoder{}
	default:
		encoder = &JSONIETFEncoder{}
	}

	jsonTree, err := encoder.Encode(nodeStruct)
	if err != nil {
		msg := fmt.Sprintf("error in constructing JSON tree from requested node: %v", err)
		return nil, status.Error(codes.Internal, msg)
	}

	jsonDump, err := json.Marshal(jsonTree)
	if err != nil {
		msg := fmt.Sprintf("error in marshaling JSON tree to bytes: %v", err)
		return nil, status.Error(codes.Internal, msg)
	}

	return encoder.BuildUpdate(path, jsonDump), nil
}
