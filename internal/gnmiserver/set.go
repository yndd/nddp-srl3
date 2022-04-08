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
	"time"

	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/yndd/ndd-yang/pkg/yparser"
	"github.com/yndd/nddp-srl3/internal/device/validator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *GnmiServerImpl) Set(ctx context.Context, req *gnmi.SetRequest) (*gnmi.SetResponse, error) {

	ok := s.unaryRPCsem.TryAcquire(1)
	if !ok {
		return nil, status.Errorf(codes.ResourceExhausted, errMaxNbrOfUnaryRPCReached)
	}
	defer s.unaryRPCsem.Release(1)

	numUpdates := len(req.GetUpdate())
	numReplaces := len(req.GetReplace())
	numDeletes := len(req.GetDelete())
	if numUpdates+numReplaces+numDeletes == 0 {
		return nil, status.Errorf(codes.InvalidArgument, errMissingPathsInGNMISet)
	}

	log := s.log.WithValues("numUpdates", numUpdates, "numReplaces", numReplaces, "numDeletes", numDeletes)
	prefix := req.GetPrefix()
	target := prefix.GetTarget()

	ce, err := s.cache.GetEntry(target)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, errTargetNotFoundInCache)
	}

	if numReplaces > 0 {
		log.Debug("Set Replace", "target", prefix.Target, "Path", yparser.GnmiPath2XPath(req.GetReplace()[0].GetPath(), true))

		// if the cache was deleted we can easily update without history

		if err := validator.ValidateUpdate(ce, req.GetReplace(), true, false, validator.Origin_GnmiServer); err != nil {
			return nil, status.Errorf(codes.Internal, err.Error())
		}
	}

	if numUpdates > 0 {
		log.Debug("Set Update", "target", prefix.Target, "Path", yparser.GnmiPath2XPath(req.GetUpdate()[0].GetPath(), true))
		// check if the update is a transaction or not -> determines if the individual reconciler has to run
		return nil, status.Errorf(codes.Unimplemented, "not implemented")
	}

	if numDeletes > 0 {
		log.Debug("Set Delete", "target", prefix.Target, "Path", yparser.GnmiPath2XPath(req.GetDelete()[0], true))
		// check if the update is a transaction or not -> determines if the individual reconciler has to run
		return nil, status.Errorf(codes.Unimplemented, "not implemented")
	}

	if err := ce.SetSystemCacheStatus(true); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &gnmi.SetResponse{
		Response: []*gnmi.UpdateResult{
			{
				Timestamp: time.Now().UnixNano(),
			},
		}}, nil
}
