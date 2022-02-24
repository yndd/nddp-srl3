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

func initSystemP4rtserver(p *yentry.Entry, opts ...yentry.EntryOption) *yentry.Entry {
	children := map[string]yentry.EntryInitFunc{
		"network-instance": initSystemP4rtserverNetworkinstance,
		"trace-options":    initSystemP4rtserverTraceoptions,
		"unix-socket":      initSystemP4rtserverUnixsocket,
	}
	e := &yentry.Entry{
		Name:             "p4rt-server",
		Key:              []string{},
		Module:           "srl_nokia-p4rt-server",
		Namespace:        "urn:srl_nokia/p4rt-server",
		Prefix:           "srl-p4rt-server",
		Parent:           p,
		Children:         make(map[string]*yentry.Entry),
		ResourceBoundary: false,
		LeafRefs:         []*leafref.LeafRef{},
		Defaults: map[string]string{
			"admin-state":   "disable",
			"rate-limit":    "60",
			"session-limit": "20",
			"timeout":       "7200",
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
