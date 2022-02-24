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

func initInterfaceSubinterfaceIpv6RouteradvertisementRouterrole(p *yentry.Entry, opts ...yentry.EntryOption) *yentry.Entry {
	children := map[string]yentry.EntryInitFunc{
		"prefix": initInterfaceSubinterfaceIpv6RouteradvertisementRouterrolePrefix,
	}
	e := &yentry.Entry{
		Name:             "router-role",
		Key:              []string{},
		Module:           "",
		Namespace:        "",
		Prefix:           "srl_nokia-if-ip-ra",
		Parent:           p,
		Children:         make(map[string]*yentry.Entry),
		ResourceBoundary: false,
		LeafRefs:         []*leafref.LeafRef{},
		Defaults: map[string]string{
			"admin-state":                "disable",
			"current-hop-limit":          "64",
			"managed-configuration-flag": "false",
			"max-advertisement-interval": "600",
			"min-advertisement-interval": "200",
			"other-configuration-flag":   "false",
			"reachable-time":             "0",
			"retransmit-time":            "0",
			"router-lifetime":            "1800",
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
