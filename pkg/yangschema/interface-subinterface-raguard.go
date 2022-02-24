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
package yangschema

import (
	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/yndd/ndd-yang/pkg/leafref"
	"github.com/yndd/ndd-yang/pkg/yentry"
)

func initInterfaceSubinterfaceRaguard(p *yentry.Entry, opts ...yentry.EntryOption) *yentry.Entry {
	children := map[string]yentry.EntryInitFunc{
		"vlan-list": initInterfaceSubinterfaceRaguardVlanlist,
	}
	e := &yentry.Entry{
		Name:             "ra-guard",
		Key:              []string{},
		Module:           "srl_nokia-ra_guard",
		Namespace:        "urn:srl_nokia/ra_guard",
		Prefix:           "srl_nokia-ra_guard",
		Parent:           p,
		Children:         make(map[string]*yentry.Entry),
		ResourceBoundary: false,
		LeafRefs: []*leafref.LeafRef{
			{
				LocalPath: &gnmi.Path{
					Elem: []*gnmi.PathElem{
						{Name: "policy"},
					},
				},
				RemotePath: &gnmi.Path{
					Elem: []*gnmi.PathElem{
						{Name: "system"},
						{Name: "ra-guard-policy", Key: map[string]string{"name": ""}},
					},
				},
			},
		},
		Defaults: map[string]string{},
	}

	for _, opt := range opts {
		opt(e)
	}

	for name, initFunc := range children {
		e.Children[name] = initFunc(e, yentry.WithLogging(e.Log))
	}

	//if e.ResourceBoundary {
	//	e.Register(&gnmi.Path{Elem: []*gnmi.PathElem{}})
	//}

	return e
}
