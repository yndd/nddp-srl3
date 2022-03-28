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
	"encoding/json"

	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/openconfig/ygot/ygot"
	"github.com/yndd/nddp-srl3/internal/cache"
	"github.com/yndd/nddp-srl3/internal/shared"
	"github.com/yndd/nddp-system/pkg/ygotnddp"
)

func (r *reconciler) getSpecdata(resource *ygotnddp.NddpSystem_Gvk) (interface{}, error) {
	spec := resource.Spec
	//r.log.Debug("getSpecdata", "specdata", *spec)
	var x1 interface{}
	if err := json.Unmarshal([]byte(*spec), &x1); err != nil {
		r.log.Debug("getSpecdata", "error", err)
		return nil, err
	}
	return x1, nil
}

// validateCreate takes the current config/goStruct and merge it with the newGoStruct
// validate the new GoStruct against the deviceSchema and return the newGoStruct
func (r *reconciler) validateCreate(resource *ygotnddp.NddpSystem_Gvk) (ygot.ValidatedGoStruct, error) {
	crDeviceName := shared.GetCrDeviceName(r.namespace, r.target.Config.Name)

	x, err := r.getSpecdata(resource)
	if err != nil {
		return nil, err
	}

	//r.log.Debug("validateCreate", "x   ", x)
	//r.log.Debug("validateCreate", "spec", *resource.Spec)

	return r.cache.ValidateCreate(crDeviceName, x)
}

// validateDelete deletes the paths from the current config/goStruct and validates the result
func (r *reconciler) validateDelete(paths []*gnmi.Path) error {
	crDeviceName := shared.GetCrDeviceName(r.namespace, r.target.Config.Name)

	return r.cache.ValidateDelete(crDeviceName, paths, cache.Origin_Reconciler)
}

// validateUpdate updates the current config/goStruct and validates the result
func (r *reconciler) validateUpdate(updates []*gnmi.Update, jsonietf bool) error {
	crDeviceName := shared.GetCrDeviceName(r.namespace, r.target.Config.Name)

	return r.cache.ValidateUpdate(crDeviceName, updates, false, jsonietf, cache.Origin_Reconciler)
}
