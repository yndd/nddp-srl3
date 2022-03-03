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
	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/openconfig/ygot/ygot"
	"github.com/yndd/nddp-srl3/internal/shared"
	systemv1alpha1 "github.com/yndd/nddp-system/apis/system/v1alpha1"
)

func (r *reconciler) getSpecdata(resource *systemv1alpha1.Gvk) (interface{}, error) {
	crDeviceName := shared.GetCrDeviceName(r.namespace, r.target.Config.Name)
	crSystemDeviceName := shared.GetCrSystemDeviceName(crDeviceName)

	x1, err := r.cache.GetCache().GetJson(
		crSystemDeviceName,
		&gnmi.Path{Target: crSystemDeviceName},
		&gnmi.Path{
			Elem: []*gnmi.PathElem{
				{Name: "gvk", Key: map[string]string{"name": resource.Name}},
			},
		},
		r.nddpSchema)
	if err != nil {
		return nil, err
	}
	switch x := x1.(type) {
	case map[string]interface{}:
		x1 = x["data"]
	}
	return x1, nil
}

// validateCreate takes the current config/goStruct and merge it with the newGoStruct
// validate the new GoStruct against the deviceSchema and return the newGoStruct
func (r *reconciler) validateCreate(resource *systemv1alpha1.Gvk) (ygot.ValidatedGoStruct, error) {
	crDeviceName := shared.GetCrDeviceName(r.namespace, r.target.Config.Name)

	x, err := r.getSpecdata(resource)
	if err != nil {
		return nil, err
	}

	return r.cache.ValidateCreate(crDeviceName, x)
}

// validateDelete deletes the paths from the current config/goStruct and validates the result
func (r *reconciler) validateDelete(paths []*gnmi.Path) error {
	crDeviceName := shared.GetCrDeviceName(r.namespace, r.target.Config.Name)

	_, err := r.cache.ValidateDelete(crDeviceName, paths)

	return err
}

// validateUpdate updates the current config/goStruct and validates the result
func (r *reconciler) validateUpdate(updates []*gnmi.Update) error {
	crDeviceName := shared.GetCrDeviceName(r.namespace, r.target.Config.Name)

	_, err := r.cache.ValidateUpdate(crDeviceName, updates)
	return err
}
