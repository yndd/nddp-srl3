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
	"github.com/yndd/ndd-yang/pkg/leafref"
	"github.com/yndd/ndd-yang/pkg/yentry"
)

func initNetworkinstanceProtocolsIgmp(p *yentry.Entry, opts ...yentry.EntryOption) *yentry.Entry {
	children := map[string]yentry.EntryInitFunc{
		"interface":     initNetworkinstanceProtocolsIgmpInterface,
		"ssm":           initNetworkinstanceProtocolsIgmpSsm,
		"trace-options": initNetworkinstanceProtocolsIgmpTraceoptions,
	}
	e := &yentry.Entry{
		Name:             "igmp",
		Key:              []string{},
		Module:           "srl_nokia-igmp",
		Namespace:        "urn:srl_nokia/igmp",
		Prefix:           "srl_nokia-igmp",
		Parent:           p,
		Children:         make(map[string]*yentry.Entry),
		ResourceBoundary: false,
		LeafRefs:         []*leafref.LeafRef{},
		Defaults: map[string]string{
			"admin-state":                "disable",
			"query-interval":             "125",
			"query-last-member-interval": "1",
			"query-response-interval":    "10",
			"robust-count":               "2",
		},
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
