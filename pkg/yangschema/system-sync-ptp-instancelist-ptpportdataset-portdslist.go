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

func initSystemSyncPtpInstancelistPtpportdatasetPortdslist(p *yentry.Entry, opts ...yentry.EntryOption) *yentry.Entry {
	children := map[string]yentry.EntryInitFunc{
		"statistics": initSystemSyncPtpInstancelistPtpportdatasetPortdslistStatistics,
	}
	e := &yentry.Entry{
		Name: "port-ds-list",
		Key: []string{
			"port-number",
		},
		Module:           "",
		Namespace:        "",
		Prefix:           "srl_nokia-sync-ptp",
		Parent:           p,
		Children:         make(map[string]*yentry.Entry),
		ResourceBoundary: false,
		LeafRefs:         []*leafref.LeafRef{},
		Defaults: map[string]string{
			"addressing-mode":            "multicast",
			"admin-state":                "disable",
			"announce-receipt-timeout":   "3",
			"asym-correction":            "0",
			"dest-mac":                   "forwardable",
			"encap-type":                 "ethernet",
			"local-priority":             "128",
			"log-announce-interval":      "-3",
			"log-min-delay-req-interval": "-4",
			"log-sync-interval":          "-4",
			"master-only":                "true",
			"peer-mean-path-delay":       "0",
			"port-role":                  "auto",
			"port-state":                 "initializing",
			"ptp-entity-type":            "ptp-port-mode",
			"vlan-id":                    "0",
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
