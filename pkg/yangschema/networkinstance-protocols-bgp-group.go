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

func initNetworkinstanceProtocolsBgpGroup(p *yentry.Entry, opts ...yentry.EntryOption) *yentry.Entry {
	children := map[string]yentry.EntryInitFunc{
		"as-path-options":    initNetworkinstanceProtocolsBgpGroupAspathoptions,
		"authentication":     initNetworkinstanceProtocolsBgpGroupAuthentication,
		"evpn":               initNetworkinstanceProtocolsBgpGroupEvpn,
		"failure-detection":  initNetworkinstanceProtocolsBgpGroupFailuredetection,
		"graceful-restart":   initNetworkinstanceProtocolsBgpGroupGracefulrestart,
		"ipv4-unicast":       initNetworkinstanceProtocolsBgpGroupIpv4unicast,
		"ipv6-unicast":       initNetworkinstanceProtocolsBgpGroupIpv6unicast,
		"local-as":           initNetworkinstanceProtocolsBgpGroupLocalas,
		"multihop":           initNetworkinstanceProtocolsBgpGroupMultihop,
		"route-reflector":    initNetworkinstanceProtocolsBgpGroupRoutereflector,
		"send-community":     initNetworkinstanceProtocolsBgpGroupSendcommunity,
		"send-default-route": initNetworkinstanceProtocolsBgpGroupSenddefaultroute,
		"timers":             initNetworkinstanceProtocolsBgpGroupTimers,
		"trace-options":      initNetworkinstanceProtocolsBgpGroupTraceoptions,
		"transport":          initNetworkinstanceProtocolsBgpGroupTransport,
	}
	e := &yentry.Entry{
		Name: "group",
		Key: []string{
			"group-name",
		},
		Module:           "",
		Namespace:        "",
		Prefix:           "srl_nokia-bgp",
		Parent:           p,
		Children:         make(map[string]*yentry.Entry),
		ResourceBoundary: false,
		LeafRefs: []*leafref.LeafRef{
			{
				LocalPath: &gnmi.Path{
					Elem: []*gnmi.PathElem{
						{Name: "export-policy"},
					},
				},
				RemotePath: &gnmi.Path{
					Elem: []*gnmi.PathElem{
						{Name: "routing-policy"},
						{Name: "policy", Key: map[string]string{"name": ""}},
					},
				},
			},
			{
				LocalPath: &gnmi.Path{
					Elem: []*gnmi.PathElem{
						{Name: "import-policy"},
					},
				},
				RemotePath: &gnmi.Path{
					Elem: []*gnmi.PathElem{
						{Name: "routing-policy"},
						{Name: "policy", Key: map[string]string{"name": ""}},
					},
				},
			},
		},
		Defaults: map[string]string{
			"admin-state":   "enable",
			"next-hop-self": "false",
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
