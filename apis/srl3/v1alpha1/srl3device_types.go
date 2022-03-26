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
	"reflect"

	nddv1 "github.com/yndd/ndd-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// A DeviceSpec defines the desired state of a Device.
type DeviceSpec struct {
	nddv1.ResourceSpec `json:",inline"`
	//Device             *Device `json:"device,omitempty"`
	// Contains all fields for the specific resource being addressed
	//+kubebuilder:pruning:PreserveUnknownFields
	//+kubebuilder:validation:Required
	Properties runtime.RawExtension `json:"properties,omitempty"`
}

// A DeviceStatus represents the observed state of a Device.
type DeviceStatus struct {
	nddv1.ResourceStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// Srl3Device is the Schema for the Device API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="TARGET",type="string",JSONPath=".status.conditions[?(@.kind=='TargetFound')].status"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.conditions[?(@.kind=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNC",type="string",JSONPath=".status.conditions[?(@.kind=='Synced')].status"
// +kubebuilder:printcolumn:name="ROOTPATH",type="string",JSONPath=".status.conditions[?(@.kind=='RootPathValidationSuccess')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:categories={ndd,srl3}
type Srl3Device struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DeviceSpec   `json:"spec,omitempty"`
	Status DeviceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Srl3DeviceList contains a list of Devices
type Srl3DeviceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Srl3Device `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Srl3Device{}, &Srl3DeviceList{})
}

// Device type metadata.
var (
	DeviceKindKind         = reflect.TypeOf(Srl3Device{}).Name()
	DeviceGroupKind        = schema.GroupKind{Group: Group, Kind: DeviceKindKind}.String()
	DeviceKindAPIVersion   = DeviceKindKind + "." + GroupVersion.String()
	DeviceGroupVersionKind = GroupVersion.WithKind(DeviceKindKind)
)
