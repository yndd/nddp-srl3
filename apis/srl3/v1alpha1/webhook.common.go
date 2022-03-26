/*
Copyright 2022 NDD.

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

package v1alpha1

import (
	"encoding/json"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

func ValidateSpec(properties runtime.RawExtension) *field.Error {
	b, err := json.Marshal(properties)
	if err != nil {
		return field.Invalid(field.NewPath("spec"), string(properties.Raw), err.Error())
	}

	_, err = m.NewConfigStruct(b, false)
	if err != nil {
		return field.Invalid(field.NewPath("spec"), string(properties.Raw), err.Error())
	}

	return nil
}
