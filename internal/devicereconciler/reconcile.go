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

package devicereconciler

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/gnxi/utils/xpath"
	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/openconfig/ygot/ygot"
	"github.com/yndd/ndd-yang/pkg/yparser"
	"github.com/yndd/nddp-srl3/internal/shared"
	"github.com/yndd/nddp-system/pkg/ygotnddp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (r *reconciler) handlePendingResources() error {
	crDeviceName := shared.GetCrDeviceName(r.namespace, r.target.Config.Name)
	crSystemDeviceName := shared.GetCrSystemDeviceName(crDeviceName)

	// get the list of Managed Resources (MR)
	resourceList, err := r.cache.GetSystemResourceList(crSystemDeviceName)
	if err != nil {
		return err
	}
	pendingResources := map[ygotnddp.E_NddpSystem_ResourceAction]*ygotnddp.NddpSystem_Gvk{}
	// loop over all resource and check if work is required on them
	for _, resource := range resourceList {
		if resource.Status == ygotnddp.NddpSystem_ResourceStatus_PENDING {
			if pendingResource, ok := pendingResources[resource.Action]; !ok {
				// no resource is pending for this action so far
				pendingResources[resource.Action] = resource
			} else {
				// check if the new resource is a better fit to be scheduled first or not
				// check path dependency
				for _, rootPath := range resource.Path {
					for _, pendingRootPath := range pendingResource.Path {
						if strings.Contains(pendingRootPath, rootPath) {
							// we need to switch the current pending resource with the new pending resource
							pendingResources[resource.Action] = resource
							continue
						}
					}
				}
				// check if the failure attempts are smaller
				if *resource.Attempt < *pendingResource.Attempt {
					// we need to switch the current pending resource with the new pending resource
					pendingResources[resource.Action] = resource
				}
			}
		}

	}
	// execute the action to the device and update the status
	// first do update, after delete and lastly create
	if resource, ok := pendingResources[ygotnddp.NddpSystem_ResourceAction_UPDATE]; ok {

		if r != nil {
			reconcileErr := r.reconcileUpdate(r.ctx, resource)
			if err := r.updateResourceStatus(resource, reconcileErr); err != nil {
				return err
			}
		}
	}
	if resource, ok := pendingResources[ygotnddp.NddpSystem_ResourceAction_DELETE]; ok {
		if r != nil {
			reconcileErr := r.reconcileDelete(r.ctx, resource)
			if err := r.updateResourceStatus(resource, reconcileErr); err != nil {
				return err
			}
			// delete the resource from the system cache if all succeeded
			if reconcileErr == nil {
				r.cache.DeleteSystemResource(crSystemDeviceName, *resource.Name)
			}
		}

	}
	if resource, ok := pendingResources[ygotnddp.NddpSystem_ResourceAction_CREATE]; ok {
		if r != nil {
			reconcileErr := r.reconcileCreate(r.ctx, resource)
			if err := r.updateResourceStatus(resource, reconcileErr); err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *reconciler) updateResourceStatus(resource *ygotnddp.NddpSystem_Gvk, reconcileErr error) error {
	log := r.log.WithValues("target", r.target.Config.Name, "address", r.target.Config.Address)
	crDeviceName := shared.GetCrDeviceName(r.namespace, r.target.Config.Name)
	crSystemDeviceName := shared.GetCrSystemDeviceName(crDeviceName)
	log.Debug("reconcile updateResourceStatus", "resource", *resource.Name, "err", reconcileErr)
	if reconcileErr != nil {
		// transaction failed
		if err := r.cache.UpdateSystemResourceStatus(crSystemDeviceName, *resource.Name, reconcileErr.Error(), ygotnddp.NddpSystem_ResourceStatus_FAILED); err != nil {
			return err
		}
		log.Debug("reconciler error", "error", reconcileErr)
		return nil
	}
	// transaction succeeded
	if err := r.cache.UpdateSystemResourceStatus(crSystemDeviceName, *resource.Name, "", ygotnddp.NddpSystem_ResourceStatus_SUCCESS); err != nil {
		return err
	}
	if resource, err := r.cache.GetSystemResource(crSystemDeviceName, *resource.Name); err == nil {
		log.Debug("reconciled resource config end", "resourceName", *resource.Name, "action", resource.Action, "status", resource.Status, "attempts", *resource.Attempt)
	}

	return nil
}

func (r *reconciler) reconcileCreate(ctx context.Context, resource *ygotnddp.NddpSystem_Gvk) error {
	log := r.log.WithValues("target", r.target.Config.Name, "address", r.target.Config.Address,
		"resourceName", *resource.Name, "action", resource.Action, "status", resource.Status, "attempts", *resource.Attempt)
	log.Debug("reconciled resource config start")

	crDeviceName := shared.GetCrDeviceName(r.namespace, r.target.Config.Name)
	crSystemDeviceName := shared.GetCrSystemDeviceName(crDeviceName)

	// validateResource, merges the latest device config with the new config
	// and validates it against the device Yang schema
	newGoStruct, err := r.validateCreate(resource)
	if err != nil {
		log.Debug("validation failed", "error", err)
		return err
	}
	j, err := ygot.EmitJSON(newGoStruct, &ygot.EmitJSONConfig{
		Format:         ygot.RFC7951,
		RFC7951Config:  &ygot.RFC7951JSONConfig{},
		SkipValidation: true,
	})
	if err != nil {
		return err
	}

	log.Debug("json update", "json", j)

	updates := []*gnmi.Update{
		{
			Path: &gnmi.Path{},
			Val:  &gnmi.TypedValue{Value: &gnmi.TypedValue_JsonIetfVal{JsonIetfVal: []byte(j)}},
		},
	}
	if _, err := r.device.GNMISet(ctx, updates, nil); err != nil {
		// update failed
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.ResourceExhausted:
				log.Debug("gnmi update failed exhausted", "error", err)
				if err := r.cache.SetSystemExhausted(crSystemDeviceName, 60); err != nil {
					return err
				}
			}
		}
		// the status will be set in the reconciler
		return err
	}
	// the status will be set in the reconciler result
	return nil
}

func (r *reconciler) reconcileUpdate(ctx context.Context, resource *ygotnddp.NddpSystem_Gvk) error {
	log := r.log.WithValues("target", r.target.Config.Name, "address", r.target.Config.Address,
		"resourceName", *resource.Name, "action", resource.Action, "status", resource.Status, "attempts", *resource.Attempt)
	log.Debug("reconciled resource config start")

	crDeviceName := shared.GetCrDeviceName(r.namespace, r.target.Config.Name)
	crSystemDeviceName := shared.GetCrSystemDeviceName(crDeviceName)

	// delete and update come from the resource
	deletes := []*gnmi.Path{}
	for _, path := range resource.Delete {
		p, err := xpath.ToGNMIPath(path)
		if err != nil {
			return err
		}
		deletes = append(deletes, p)
	}
	updates := []*gnmi.Update{}
	for _, u := range resource.Update {
		p, err := xpath.ToGNMIPath(*u.Path)
		if err != nil {
			return err
		}
		updates = append(updates, &gnmi.Update{
			Path: p,
			Val:  &gnmi.TypedValue{Value: &gnmi.TypedValue_JsonIetfVal{JsonIetfVal: []byte(*u.Val)}},
		})
	}

	for _, path := range deletes {
		log.Debug("reconciling device config update -> delete paths", "path", yparser.GnmiPath2XPath(path, true))
	}
	for _, u := range updates {
		log.Debug("reconciling device config update -> update info ", "path", yparser.GnmiPath2XPath(u.GetPath(), true), "val", u.GetVal())
	}

	if err := r.validateDelete(deletes); err != nil {
		return err
	}
	if err := r.validateUpdate(updates, true); err != nil {
		return err
	}

	// execute the deletes and updates in the cache and to the device
	_, err := r.device.GNMISet(r.ctx, updates, deletes)
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.ResourceExhausted:
				log.Debug("gnmi update failed exhausted", "error", err)
				if err := r.cache.SetSystemExhausted(crSystemDeviceName, 60); err != nil {
					return err
				}
			}
		}
		// the status will be set in the reconciler
		return err
	}
	// the status will be set in the reconciler
	return nil
}

func (r *reconciler) reconcileDelete(ctx context.Context, resource *ygotnddp.NddpSystem_Gvk) error {
	log := r.log.WithValues("target", r.target.Config.Name, "address", r.target.Config.Address,
		"resourceName", *resource.Name, "action", resource.Action, "status", resource.Status, "attempts", *resource.Attempt)
	log.Debug("reconciled resource config start")

	crDeviceName := shared.GetCrDeviceName(r.namespace, r.target.Config.Name)
	crSystemDeviceName := shared.GetCrSystemDeviceName(crDeviceName)

	delPaths := make([]*gnmi.Path, 0)
	murder := false
	for _, xp := range resource.Path {
		log.Debug("reconciled resource config start", "path", xp)
		path, err := xpath.ToGNMIPath(xp)
		if err != nil {
			return err
		}

		if path == nil || len(path.Elem) == 0 {
			murder = true
		}
		delPaths = append(delPaths, path)
	}
	// if we dont do suicide and len delete paths > 0, perform delete
	if !murder && len(delPaths) > 0 {
		// validate Delete
		if err := r.validateDelete(delPaths); err != nil {
			return err
		}
		// apply deletes on the device
		_, err := r.device.GNMISet(ctx, nil, delPaths)
		if err != nil {
			if e, ok := status.FromError(err); ok {
				switch e.Code() {
				case codes.ResourceExhausted:
					log.Debug("gnmi delete failed exhausted", "error", err)
					if err := r.cache.SetSystemExhausted(crSystemDeviceName, 60); err != nil {
						return err
					}
				}
			}
			log.Debug("gnmi delete failed", "Paths", delPaths, "Error", err)
			// we only process 1 resource at the time
			return err

		}
		// the status will be set in the reconciler
		log.Debug("gnmi delete success", "Paths", delPaths)
		return nil
	}
	// the status will be set in the reconciler result
	return fmt.Errorf("deleting a resource with no valid paths murder: %t, resourcePaths %v, delPaths: %v", murder, resource.Path, delPaths)
}
