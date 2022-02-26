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
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// Device struct
type Device struct {
	//RootAcl
	Acl *DeviceAcl `json:"acl,omitempty"`
	//RootBfd
	Bfd *DeviceBfd `json:"bfd,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Interface []*DeviceInterface `json:"interface,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Networkinstance []*DeviceNetworkinstance `json:"network-instance,omitempty"`
	//RootOam
	Oam *DeviceOam `json:"oam,omitempty"`
	//RootPlatform
	Platform *DevicePlatform `json:"platform,omitempty"`
	//RootQos
	Qos *DeviceQos `json:"qos,omitempty"`
	//RootRoutingpolicy
	Routingpolicy *DeviceRoutingpolicy `json:"routing-policy,omitempty"`
	//RootSystem
	System *DeviceSystem `json:"system,omitempty"`
	//RootTunnel
	Tunnel *DeviceTunnel `json:"tunnel,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Tunnelinterface []*DeviceTunnelinterface `json:"tunnel-interface,omitempty"`
}

// DeviceAcl struct
type DeviceAcl struct {
	//RootAclCapturefilter
	Capturefilter *DeviceAclCapturefilter `json:"capture-filter,omitempty"`
	//RootAclCpmfilter
	Cpmfilter *DeviceAclCpmfilter `json:"cpm-filter,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Ipv4filter []*DeviceAclIpv4filter `json:"ipv4-filter,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Ipv6filter []*DeviceAclIpv6filter `json:"ipv6-filter,omitempty"`
	//RootAclPolicers
	Policers *DeviceAclPolicers `json:"policers,omitempty"`
	//RootAclSystemfilter
	Systemfilter *DeviceAclSystemfilter `json:"system-filter,omitempty"`
	// +kubebuilder:validation:Enum=`default`;`ipv4-egress-scaled`
	Tcamprofile E_DeviceAclTcamprofile `json:"tcam-profile,omitempty"`
	//Tcamprofile *string `json:"tcam-profile,omitempty"`
}

// DeviceAclCapturefilter struct
type DeviceAclCapturefilter struct {
	//RootAclCapturefilterIpv4filter
	Ipv4filter *DeviceAclCapturefilterIpv4filter `json:"ipv4-filter,omitempty"`
	//RootAclCapturefilterIpv6filter
	Ipv6filter *DeviceAclCapturefilterIpv6filter `json:"ipv6-filter,omitempty"`
}

// DeviceAclCapturefilterIpv4filter struct
type DeviceAclCapturefilterIpv4filter struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Entry []*DeviceAclCapturefilterIpv4filterEntry `json:"entry,omitempty"`
}

// DeviceAclCapturefilterIpv4filterEntry struct
type DeviceAclCapturefilterIpv4filterEntry struct {
	//RootAclCapturefilterIpv4filterEntryAction
	Action *DeviceAclCapturefilterIpv4filterEntryAction `json:"action,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Description *string `json:"description,omitempty"`
	//RootAclCapturefilterIpv4filterEntryMatch
	Match *DeviceAclCapturefilterIpv4filterEntryMatch `json:"match,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=65535
	Sequenceid *uint32 `json:"sequence-id"`
}

// DeviceAclCapturefilterIpv4filterEntryAction struct
type DeviceAclCapturefilterIpv4filterEntryAction struct {
	//RootAclCapturefilterIpv4filterEntryActionAccept
	Accept *DeviceAclCapturefilterIpv4filterEntryActionAccept `json:"accept,omitempty"`
	//RootAclCapturefilterIpv4filterEntryActionCopy
	Copy *DeviceAclCapturefilterIpv4filterEntryActionCopy `json:"copy,omitempty"`
}

// DeviceAclCapturefilterIpv4filterEntryActionAccept struct
type DeviceAclCapturefilterIpv4filterEntryActionAccept struct {
}

// DeviceAclCapturefilterIpv4filterEntryActionCopy struct
type DeviceAclCapturefilterIpv4filterEntryActionCopy struct {
}

// DeviceAclCapturefilterIpv4filterEntryMatch struct
type DeviceAclCapturefilterIpv4filterEntryMatch struct {
	//RootAclCapturefilterIpv4filterEntryMatchDestinationip
	Destinationip *DeviceAclCapturefilterIpv4filterEntryMatchDestinationip `json:"destination-ip,omitempty"`
	//RootAclCapturefilterIpv4filterEntryMatchDestinationport
	Destinationport *DeviceAclCapturefilterIpv4filterEntryMatchDestinationport `json:"destination-port,omitempty"`
	Firstfragment   *bool                                                      `json:"first-fragment,omitempty"`
	Fragment        *bool                                                      `json:"fragment,omitempty"`
	//RootAclCapturefilterIpv4filterEntryMatchIcmp
	Icmp     *DeviceAclCapturefilterIpv4filterEntryMatchIcmp `json:"icmp,omitempty"`
	Protocol *string                                         `json:"protocol,omitempty"`
	//RootAclCapturefilterIpv4filterEntryMatchSourceip
	Sourceip *DeviceAclCapturefilterIpv4filterEntryMatchSourceip `json:"source-ip,omitempty"`
	//RootAclCapturefilterIpv4filterEntryMatchSourceport
	Sourceport *DeviceAclCapturefilterIpv4filterEntryMatchSourceport `json:"source-port,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(\(|\)|&|\||!|ack|rst|syn)+`
	Tcpflags *string `json:"tcp-flags,omitempty"`
}

// DeviceAclCapturefilterIpv4filterEntryMatchDestinationip struct
type DeviceAclCapturefilterIpv4filterEntryMatchDestinationip struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Address *string `json:"address,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Mask *string `json:"mask,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))`
	Prefix *string `json:"prefix,omitempty"`
}

// DeviceAclCapturefilterIpv4filterEntryMatchDestinationport struct
type DeviceAclCapturefilterIpv4filterEntryMatchDestinationport struct {
	// +kubebuilder:validation:Enum=`eq`;`ge`;`le`
	Operator E_DeviceAclCapturefilterIpv4filterEntryMatchDestinationportOperator `json:"operator,omitempty"`
	//Operator *string `json:"operator,omitempty"`
	//RootAclCapturefilterIpv4filterEntryMatchDestinationportRange
	Range *DeviceAclCapturefilterIpv4filterEntryMatchDestinationportRange `json:"range,omitempty"`
	Value *string                                                         `json:"value,omitempty"`
}

// DeviceAclCapturefilterIpv4filterEntryMatchDestinationportRange struct
type DeviceAclCapturefilterIpv4filterEntryMatchDestinationportRange struct {
	End   *string `json:"end,omitempty"`
	Start *string `json:"start,omitempty"`
}

// DeviceAclCapturefilterIpv4filterEntryMatchIcmp struct
type DeviceAclCapturefilterIpv4filterEntryMatchIcmp struct {
	//RootAclCapturefilterIpv4filterEntryMatchIcmpCode
	Code *DeviceAclCapturefilterIpv4filterEntryMatchIcmpCode `json:"code,omitempty"`
	Type *string                                             `json:"type,omitempty"`
}

// DeviceAclCapturefilterIpv4filterEntryMatchIcmpCode struct
type DeviceAclCapturefilterIpv4filterEntryMatchIcmpCode struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	Code *uint8 `json:"code,omitempty"`
}

// DeviceAclCapturefilterIpv4filterEntryMatchSourceip struct
type DeviceAclCapturefilterIpv4filterEntryMatchSourceip struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Address *string `json:"address,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Mask *string `json:"mask,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))`
	Prefix *string `json:"prefix,omitempty"`
}

// DeviceAclCapturefilterIpv4filterEntryMatchSourceport struct
type DeviceAclCapturefilterIpv4filterEntryMatchSourceport struct {
	// +kubebuilder:validation:Enum=`eq`;`ge`;`le`
	Operator E_DeviceAclCapturefilterIpv4filterEntryMatchSourceportOperator `json:"operator,omitempty"`
	//Operator *string `json:"operator,omitempty"`
	//RootAclCapturefilterIpv4filterEntryMatchSourceportRange
	Range *DeviceAclCapturefilterIpv4filterEntryMatchSourceportRange `json:"range,omitempty"`
	Value *string                                                    `json:"value,omitempty"`
}

// DeviceAclCapturefilterIpv4filterEntryMatchSourceportRange struct
type DeviceAclCapturefilterIpv4filterEntryMatchSourceportRange struct {
	End   *string `json:"end,omitempty"`
	Start *string `json:"start,omitempty"`
}

// DeviceAclCapturefilterIpv6filter struct
type DeviceAclCapturefilterIpv6filter struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Entry []*DeviceAclCapturefilterIpv6filterEntry `json:"entry,omitempty"`
}

// DeviceAclCapturefilterIpv6filterEntry struct
type DeviceAclCapturefilterIpv6filterEntry struct {
	//RootAclCapturefilterIpv6filterEntryAction
	Action *DeviceAclCapturefilterIpv6filterEntryAction `json:"action,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Description *string `json:"description,omitempty"`
	//RootAclCapturefilterIpv6filterEntryMatch
	Match *DeviceAclCapturefilterIpv6filterEntryMatch `json:"match,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=65535
	Sequenceid *uint32 `json:"sequence-id"`
}

// DeviceAclCapturefilterIpv6filterEntryAction struct
type DeviceAclCapturefilterIpv6filterEntryAction struct {
	//RootAclCapturefilterIpv6filterEntryActionAccept
	Accept *DeviceAclCapturefilterIpv6filterEntryActionAccept `json:"accept,omitempty"`
	//RootAclCapturefilterIpv6filterEntryActionCopy
	Copy *DeviceAclCapturefilterIpv6filterEntryActionCopy `json:"copy,omitempty"`
}

// DeviceAclCapturefilterIpv6filterEntryActionAccept struct
type DeviceAclCapturefilterIpv6filterEntryActionAccept struct {
}

// DeviceAclCapturefilterIpv6filterEntryActionCopy struct
type DeviceAclCapturefilterIpv6filterEntryActionCopy struct {
}

// DeviceAclCapturefilterIpv6filterEntryMatch struct
type DeviceAclCapturefilterIpv6filterEntryMatch struct {
	//RootAclCapturefilterIpv6filterEntryMatchDestinationip
	Destinationip *DeviceAclCapturefilterIpv6filterEntryMatchDestinationip `json:"destination-ip,omitempty"`
	//RootAclCapturefilterIpv6filterEntryMatchDestinationport
	Destinationport *DeviceAclCapturefilterIpv6filterEntryMatchDestinationport `json:"destination-port,omitempty"`
	//RootAclCapturefilterIpv6filterEntryMatchIcmp6
	Icmp6      *DeviceAclCapturefilterIpv6filterEntryMatchIcmp6 `json:"icmp6,omitempty"`
	Nextheader *string                                          `json:"next-header,omitempty"`
	//RootAclCapturefilterIpv6filterEntryMatchSourceip
	Sourceip *DeviceAclCapturefilterIpv6filterEntryMatchSourceip `json:"source-ip,omitempty"`
	//RootAclCapturefilterIpv6filterEntryMatchSourceport
	Sourceport *DeviceAclCapturefilterIpv6filterEntryMatchSourceport `json:"source-port,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(\(|\)|&|\||!|ack|rst|syn)+`
	Tcpflags *string `json:"tcp-flags,omitempty"`
}

// DeviceAclCapturefilterIpv6filterEntryMatchDestinationip struct
type DeviceAclCapturefilterIpv6filterEntryMatchDestinationip struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Address *string `json:"address,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Mask *string `json:"mask,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	Prefix *string `json:"prefix,omitempty"`
}

// DeviceAclCapturefilterIpv6filterEntryMatchDestinationport struct
type DeviceAclCapturefilterIpv6filterEntryMatchDestinationport struct {
	// +kubebuilder:validation:Enum=`eq`;`ge`;`le`
	Operator E_DeviceAclCapturefilterIpv6filterEntryMatchDestinationportOperator `json:"operator,omitempty"`
	//Operator *string `json:"operator,omitempty"`
	//RootAclCapturefilterIpv6filterEntryMatchDestinationportRange
	Range *DeviceAclCapturefilterIpv6filterEntryMatchDestinationportRange `json:"range,omitempty"`
	Value *string                                                         `json:"value,omitempty"`
}

// DeviceAclCapturefilterIpv6filterEntryMatchDestinationportRange struct
type DeviceAclCapturefilterIpv6filterEntryMatchDestinationportRange struct {
	End   *string `json:"end,omitempty"`
	Start *string `json:"start,omitempty"`
}

// DeviceAclCapturefilterIpv6filterEntryMatchIcmp6 struct
type DeviceAclCapturefilterIpv6filterEntryMatchIcmp6 struct {
	//RootAclCapturefilterIpv6filterEntryMatchIcmp6Code
	Code *DeviceAclCapturefilterIpv6filterEntryMatchIcmp6Code `json:"code,omitempty"`
	Type *string                                              `json:"type,omitempty"`
}

// DeviceAclCapturefilterIpv6filterEntryMatchIcmp6Code struct
type DeviceAclCapturefilterIpv6filterEntryMatchIcmp6Code struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	Code *uint8 `json:"code,omitempty"`
}

// DeviceAclCapturefilterIpv6filterEntryMatchSourceip struct
type DeviceAclCapturefilterIpv6filterEntryMatchSourceip struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Address *string `json:"address,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Mask *string `json:"mask,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	Prefix *string `json:"prefix,omitempty"`
}

// DeviceAclCapturefilterIpv6filterEntryMatchSourceport struct
type DeviceAclCapturefilterIpv6filterEntryMatchSourceport struct {
	// +kubebuilder:validation:Enum=`eq`;`ge`;`le`
	Operator E_DeviceAclCapturefilterIpv6filterEntryMatchSourceportOperator `json:"operator,omitempty"`
	//Operator *string `json:"operator,omitempty"`
	//RootAclCapturefilterIpv6filterEntryMatchSourceportRange
	Range *DeviceAclCapturefilterIpv6filterEntryMatchSourceportRange `json:"range,omitempty"`
	Value *string                                                    `json:"value,omitempty"`
}

// DeviceAclCapturefilterIpv6filterEntryMatchSourceportRange struct
type DeviceAclCapturefilterIpv6filterEntryMatchSourceportRange struct {
	End   *string `json:"end,omitempty"`
	Start *string `json:"start,omitempty"`
}

// DeviceAclCpmfilter struct
type DeviceAclCpmfilter struct {
	//RootAclCpmfilterIpv4filter
	Ipv4filter *DeviceAclCpmfilterIpv4filter `json:"ipv4-filter,omitempty"`
	//RootAclCpmfilterIpv6filter
	Ipv6filter *DeviceAclCpmfilterIpv6filter `json:"ipv6-filter,omitempty"`
}

// DeviceAclCpmfilterIpv4filter struct
type DeviceAclCpmfilterIpv4filter struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Entry              []*DeviceAclCpmfilterIpv4filterEntry `json:"entry,omitempty"`
	Statisticsperentry *bool                                `json:"statistics-per-entry,omitempty"`
}

// DeviceAclCpmfilterIpv4filterEntry struct
type DeviceAclCpmfilterIpv4filterEntry struct {
	//RootAclCpmfilterIpv4filterEntryAction
	Action *DeviceAclCpmfilterIpv4filterEntryAction `json:"action,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Description *string `json:"description,omitempty"`
	//RootAclCpmfilterIpv4filterEntryMatch
	Match *DeviceAclCpmfilterIpv4filterEntryMatch `json:"match,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=65535
	Sequenceid *uint32 `json:"sequence-id"`
}

// DeviceAclCpmfilterIpv4filterEntryAction struct
type DeviceAclCpmfilterIpv4filterEntryAction struct {
	//RootAclCpmfilterIpv4filterEntryActionAccept
	Accept *DeviceAclCpmfilterIpv4filterEntryActionAccept `json:"accept,omitempty"`
	//RootAclCpmfilterIpv4filterEntryActionDrop
	Drop *DeviceAclCpmfilterIpv4filterEntryActionDrop `json:"drop,omitempty"`
}

// DeviceAclCpmfilterIpv4filterEntryActionAccept struct
type DeviceAclCpmfilterIpv4filterEntryActionAccept struct {
	// +kubebuilder:default:=false
	Log *bool `json:"log,omitempty"`
	//RootAclCpmfilterIpv4filterEntryActionAcceptRatelimit
	Ratelimit *DeviceAclCpmfilterIpv4filterEntryActionAcceptRatelimit `json:"rate-limit,omitempty"`
}

// DeviceAclCpmfilterIpv4filterEntryActionAcceptRatelimit struct
type DeviceAclCpmfilterIpv4filterEntryActionAcceptRatelimit struct {
	Distributedpolicer *string `json:"distributed-policer,omitempty"`
	Systemcpupolicer   *string `json:"system-cpu-policer,omitempty"`
}

// DeviceAclCpmfilterIpv4filterEntryActionDrop struct
type DeviceAclCpmfilterIpv4filterEntryActionDrop struct {
	// +kubebuilder:default:=false
	Log *bool `json:"log,omitempty"`
}

// DeviceAclCpmfilterIpv4filterEntryMatch struct
type DeviceAclCpmfilterIpv4filterEntryMatch struct {
	//RootAclCpmfilterIpv4filterEntryMatchDestinationip
	Destinationip *DeviceAclCpmfilterIpv4filterEntryMatchDestinationip `json:"destination-ip,omitempty"`
	//RootAclCpmfilterIpv4filterEntryMatchDestinationport
	Destinationport *DeviceAclCpmfilterIpv4filterEntryMatchDestinationport `json:"destination-port,omitempty"`
	Firstfragment   *bool                                                  `json:"first-fragment,omitempty"`
	Fragment        *bool                                                  `json:"fragment,omitempty"`
	//RootAclCpmfilterIpv4filterEntryMatchIcmp
	Icmp     *DeviceAclCpmfilterIpv4filterEntryMatchIcmp `json:"icmp,omitempty"`
	Protocol *string                                     `json:"protocol,omitempty"`
	//RootAclCpmfilterIpv4filterEntryMatchSourceip
	Sourceip *DeviceAclCpmfilterIpv4filterEntryMatchSourceip `json:"source-ip,omitempty"`
	//RootAclCpmfilterIpv4filterEntryMatchSourceport
	Sourceport *DeviceAclCpmfilterIpv4filterEntryMatchSourceport `json:"source-port,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(\(|\)|&|\||!|ack|rst|syn)+`
	Tcpflags *string `json:"tcp-flags,omitempty"`
}

// DeviceAclCpmfilterIpv4filterEntryMatchDestinationip struct
type DeviceAclCpmfilterIpv4filterEntryMatchDestinationip struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Address *string `json:"address,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Mask *string `json:"mask,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))`
	Prefix *string `json:"prefix,omitempty"`
}

// DeviceAclCpmfilterIpv4filterEntryMatchDestinationport struct
type DeviceAclCpmfilterIpv4filterEntryMatchDestinationport struct {
	// +kubebuilder:validation:Enum=`eq`;`ge`;`le`
	Operator E_DeviceAclCpmfilterIpv4filterEntryMatchDestinationportOperator `json:"operator,omitempty"`
	//Operator *string `json:"operator,omitempty"`
	//RootAclCpmfilterIpv4filterEntryMatchDestinationportRange
	Range *DeviceAclCpmfilterIpv4filterEntryMatchDestinationportRange `json:"range,omitempty"`
	Value *string                                                     `json:"value,omitempty"`
}

// DeviceAclCpmfilterIpv4filterEntryMatchDestinationportRange struct
type DeviceAclCpmfilterIpv4filterEntryMatchDestinationportRange struct {
	End   *string `json:"end,omitempty"`
	Start *string `json:"start,omitempty"`
}

// DeviceAclCpmfilterIpv4filterEntryMatchIcmp struct
type DeviceAclCpmfilterIpv4filterEntryMatchIcmp struct {
	//RootAclCpmfilterIpv4filterEntryMatchIcmpCode
	Code *DeviceAclCpmfilterIpv4filterEntryMatchIcmpCode `json:"code,omitempty"`
	Type *string                                         `json:"type,omitempty"`
}

// DeviceAclCpmfilterIpv4filterEntryMatchIcmpCode struct
type DeviceAclCpmfilterIpv4filterEntryMatchIcmpCode struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	Code *uint8 `json:"code,omitempty"`
}

// DeviceAclCpmfilterIpv4filterEntryMatchSourceip struct
type DeviceAclCpmfilterIpv4filterEntryMatchSourceip struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Address *string `json:"address,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Mask *string `json:"mask,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))`
	Prefix *string `json:"prefix,omitempty"`
}

// DeviceAclCpmfilterIpv4filterEntryMatchSourceport struct
type DeviceAclCpmfilterIpv4filterEntryMatchSourceport struct {
	// +kubebuilder:validation:Enum=`eq`;`ge`;`le`
	Operator E_DeviceAclCpmfilterIpv4filterEntryMatchSourceportOperator `json:"operator,omitempty"`
	//Operator *string `json:"operator,omitempty"`
	//RootAclCpmfilterIpv4filterEntryMatchSourceportRange
	Range *DeviceAclCpmfilterIpv4filterEntryMatchSourceportRange `json:"range,omitempty"`
	Value *string                                                `json:"value,omitempty"`
}

// DeviceAclCpmfilterIpv4filterEntryMatchSourceportRange struct
type DeviceAclCpmfilterIpv4filterEntryMatchSourceportRange struct {
	End   *string `json:"end,omitempty"`
	Start *string `json:"start,omitempty"`
}

// DeviceAclCpmfilterIpv6filter struct
type DeviceAclCpmfilterIpv6filter struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Entry              []*DeviceAclCpmfilterIpv6filterEntry `json:"entry,omitempty"`
	Statisticsperentry *bool                                `json:"statistics-per-entry,omitempty"`
}

// DeviceAclCpmfilterIpv6filterEntry struct
type DeviceAclCpmfilterIpv6filterEntry struct {
	//RootAclCpmfilterIpv6filterEntryAction
	Action *DeviceAclCpmfilterIpv6filterEntryAction `json:"action,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Description *string `json:"description,omitempty"`
	//RootAclCpmfilterIpv6filterEntryMatch
	Match *DeviceAclCpmfilterIpv6filterEntryMatch `json:"match,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=65535
	Sequenceid *uint32 `json:"sequence-id"`
}

// DeviceAclCpmfilterIpv6filterEntryAction struct
type DeviceAclCpmfilterIpv6filterEntryAction struct {
	//RootAclCpmfilterIpv6filterEntryActionAccept
	Accept *DeviceAclCpmfilterIpv6filterEntryActionAccept `json:"accept,omitempty"`
	//RootAclCpmfilterIpv6filterEntryActionDrop
	Drop *DeviceAclCpmfilterIpv6filterEntryActionDrop `json:"drop,omitempty"`
}

// DeviceAclCpmfilterIpv6filterEntryActionAccept struct
type DeviceAclCpmfilterIpv6filterEntryActionAccept struct {
	// +kubebuilder:default:=false
	Log *bool `json:"log,omitempty"`
	//RootAclCpmfilterIpv6filterEntryActionAcceptRatelimit
	Ratelimit *DeviceAclCpmfilterIpv6filterEntryActionAcceptRatelimit `json:"rate-limit,omitempty"`
}

// DeviceAclCpmfilterIpv6filterEntryActionAcceptRatelimit struct
type DeviceAclCpmfilterIpv6filterEntryActionAcceptRatelimit struct {
	Distributedpolicer *string `json:"distributed-policer,omitempty"`
	Systemcpupolicer   *string `json:"system-cpu-policer,omitempty"`
}

// DeviceAclCpmfilterIpv6filterEntryActionDrop struct
type DeviceAclCpmfilterIpv6filterEntryActionDrop struct {
	// +kubebuilder:default:=false
	Log *bool `json:"log,omitempty"`
}

// DeviceAclCpmfilterIpv6filterEntryMatch struct
type DeviceAclCpmfilterIpv6filterEntryMatch struct {
	//RootAclCpmfilterIpv6filterEntryMatchDestinationip
	Destinationip *DeviceAclCpmfilterIpv6filterEntryMatchDestinationip `json:"destination-ip,omitempty"`
	//RootAclCpmfilterIpv6filterEntryMatchDestinationport
	Destinationport *DeviceAclCpmfilterIpv6filterEntryMatchDestinationport `json:"destination-port,omitempty"`
	//RootAclCpmfilterIpv6filterEntryMatchIcmp6
	Icmp6      *DeviceAclCpmfilterIpv6filterEntryMatchIcmp6 `json:"icmp6,omitempty"`
	Nextheader *string                                      `json:"next-header,omitempty"`
	//RootAclCpmfilterIpv6filterEntryMatchSourceip
	Sourceip *DeviceAclCpmfilterIpv6filterEntryMatchSourceip `json:"source-ip,omitempty"`
	//RootAclCpmfilterIpv6filterEntryMatchSourceport
	Sourceport *DeviceAclCpmfilterIpv6filterEntryMatchSourceport `json:"source-port,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(\(|\)|&|\||!|ack|rst|syn)+`
	Tcpflags *string `json:"tcp-flags,omitempty"`
}

// DeviceAclCpmfilterIpv6filterEntryMatchDestinationip struct
type DeviceAclCpmfilterIpv6filterEntryMatchDestinationip struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Address *string `json:"address,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Mask *string `json:"mask,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	Prefix *string `json:"prefix,omitempty"`
}

// DeviceAclCpmfilterIpv6filterEntryMatchDestinationport struct
type DeviceAclCpmfilterIpv6filterEntryMatchDestinationport struct {
	// +kubebuilder:validation:Enum=`eq`;`ge`;`le`
	Operator E_DeviceAclCpmfilterIpv6filterEntryMatchDestinationportOperator `json:"operator,omitempty"`
	//Operator *string `json:"operator,omitempty"`
	//RootAclCpmfilterIpv6filterEntryMatchDestinationportRange
	Range *DeviceAclCpmfilterIpv6filterEntryMatchDestinationportRange `json:"range,omitempty"`
	Value *string                                                     `json:"value,omitempty"`
}

// DeviceAclCpmfilterIpv6filterEntryMatchDestinationportRange struct
type DeviceAclCpmfilterIpv6filterEntryMatchDestinationportRange struct {
	End   *string `json:"end,omitempty"`
	Start *string `json:"start,omitempty"`
}

// DeviceAclCpmfilterIpv6filterEntryMatchIcmp6 struct
type DeviceAclCpmfilterIpv6filterEntryMatchIcmp6 struct {
	//RootAclCpmfilterIpv6filterEntryMatchIcmp6Code
	Code *DeviceAclCpmfilterIpv6filterEntryMatchIcmp6Code `json:"code,omitempty"`
	Type *string                                          `json:"type,omitempty"`
}

// DeviceAclCpmfilterIpv6filterEntryMatchIcmp6Code struct
type DeviceAclCpmfilterIpv6filterEntryMatchIcmp6Code struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	Code *uint8 `json:"code,omitempty"`
}

// DeviceAclCpmfilterIpv6filterEntryMatchSourceip struct
type DeviceAclCpmfilterIpv6filterEntryMatchSourceip struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Address *string `json:"address,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Mask *string `json:"mask,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	Prefix *string `json:"prefix,omitempty"`
}

// DeviceAclCpmfilterIpv6filterEntryMatchSourceport struct
type DeviceAclCpmfilterIpv6filterEntryMatchSourceport struct {
	// +kubebuilder:validation:Enum=`eq`;`ge`;`le`
	Operator E_DeviceAclCpmfilterIpv6filterEntryMatchSourceportOperator `json:"operator,omitempty"`
	//Operator *string `json:"operator,omitempty"`
	//RootAclCpmfilterIpv6filterEntryMatchSourceportRange
	Range *DeviceAclCpmfilterIpv6filterEntryMatchSourceportRange `json:"range,omitempty"`
	Value *string                                                `json:"value,omitempty"`
}

// DeviceAclCpmfilterIpv6filterEntryMatchSourceportRange struct
type DeviceAclCpmfilterIpv6filterEntryMatchSourceportRange struct {
	End   *string `json:"end,omitempty"`
	Start *string `json:"start,omitempty"`
}

// DeviceAclIpv4filter struct
type DeviceAclIpv4filter struct {
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Description *string `json:"description,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Entry []*DeviceAclIpv4filterEntry `json:"entry,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name               *string `json:"name"`
	Statisticsperentry *bool   `json:"statistics-per-entry,omitempty"`
	// +kubebuilder:validation:Enum=`disabled`;`input-and-output`;`input-only`;`output-only`
	// +kubebuilder:default:="disabled"
	Subinterfacespecific E_DeviceAclIpv4filterSubinterfacespecific `json:"subinterface-specific,omitempty"`
	//Subinterfacespecific *string `json:"subinterface-specific,omitempty"`
}

// DeviceAclIpv4filterEntry struct
type DeviceAclIpv4filterEntry struct {
	//RootAclIpv4filterEntryAction
	Action *DeviceAclIpv4filterEntryAction `json:"action,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Description *string `json:"description,omitempty"`
	//RootAclIpv4filterEntryMatch
	Match *DeviceAclIpv4filterEntryMatch `json:"match,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=65535
	Sequenceid *uint32 `json:"sequence-id"`
}

// DeviceAclIpv4filterEntryAction struct
type DeviceAclIpv4filterEntryAction struct {
	//RootAclIpv4filterEntryActionAccept
	Accept *DeviceAclIpv4filterEntryActionAccept `json:"accept,omitempty"`
	//RootAclIpv4filterEntryActionDrop
	Drop *DeviceAclIpv4filterEntryActionDrop `json:"drop,omitempty"`
}

// DeviceAclIpv4filterEntryActionAccept struct
type DeviceAclIpv4filterEntryActionAccept struct {
	// +kubebuilder:default:=false
	Log *bool `json:"log,omitempty"`
}

// DeviceAclIpv4filterEntryActionDrop struct
type DeviceAclIpv4filterEntryActionDrop struct {
	// +kubebuilder:default:=false
	Log *bool `json:"log,omitempty"`
}

// DeviceAclIpv4filterEntryMatch struct
type DeviceAclIpv4filterEntryMatch struct {
	//RootAclIpv4filterEntryMatchDestinationip
	Destinationip *DeviceAclIpv4filterEntryMatchDestinationip `json:"destination-ip,omitempty"`
	//RootAclIpv4filterEntryMatchDestinationport
	Destinationport *DeviceAclIpv4filterEntryMatchDestinationport `json:"destination-port,omitempty"`
	Firstfragment   *bool                                         `json:"first-fragment,omitempty"`
	Fragment        *bool                                         `json:"fragment,omitempty"`
	//RootAclIpv4filterEntryMatchIcmp
	Icmp     *DeviceAclIpv4filterEntryMatchIcmp `json:"icmp,omitempty"`
	Protocol *string                            `json:"protocol,omitempty"`
	//RootAclIpv4filterEntryMatchSourceip
	Sourceip *DeviceAclIpv4filterEntryMatchSourceip `json:"source-ip,omitempty"`
	//RootAclIpv4filterEntryMatchSourceport
	Sourceport *DeviceAclIpv4filterEntryMatchSourceport `json:"source-port,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(\(|\)|&|\||!|ack|rst|syn)+`
	Tcpflags *string `json:"tcp-flags,omitempty"`
}

// DeviceAclIpv4filterEntryMatchDestinationip struct
type DeviceAclIpv4filterEntryMatchDestinationip struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Address *string `json:"address,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Mask *string `json:"mask,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))`
	Prefix *string `json:"prefix,omitempty"`
}

// DeviceAclIpv4filterEntryMatchDestinationport struct
type DeviceAclIpv4filterEntryMatchDestinationport struct {
	// +kubebuilder:validation:Enum=`eq`;`ge`;`le`
	Operator E_DeviceAclIpv4filterEntryMatchDestinationportOperator `json:"operator,omitempty"`
	//Operator *string `json:"operator,omitempty"`
	//RootAclIpv4filterEntryMatchDestinationportRange
	Range *DeviceAclIpv4filterEntryMatchDestinationportRange `json:"range,omitempty"`
	Value *string                                            `json:"value,omitempty"`
}

// DeviceAclIpv4filterEntryMatchDestinationportRange struct
type DeviceAclIpv4filterEntryMatchDestinationportRange struct {
	End   *string `json:"end,omitempty"`
	Start *string `json:"start,omitempty"`
}

// DeviceAclIpv4filterEntryMatchIcmp struct
type DeviceAclIpv4filterEntryMatchIcmp struct {
	//RootAclIpv4filterEntryMatchIcmpCode
	Code *DeviceAclIpv4filterEntryMatchIcmpCode `json:"code,omitempty"`
	Type *string                                `json:"type,omitempty"`
}

// DeviceAclIpv4filterEntryMatchIcmpCode struct
type DeviceAclIpv4filterEntryMatchIcmpCode struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	Code *uint8 `json:"code,omitempty"`
}

// DeviceAclIpv4filterEntryMatchSourceip struct
type DeviceAclIpv4filterEntryMatchSourceip struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Address *string `json:"address,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Mask *string `json:"mask,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))`
	Prefix *string `json:"prefix,omitempty"`
}

// DeviceAclIpv4filterEntryMatchSourceport struct
type DeviceAclIpv4filterEntryMatchSourceport struct {
	// +kubebuilder:validation:Enum=`eq`;`ge`;`le`
	Operator E_DeviceAclIpv4filterEntryMatchSourceportOperator `json:"operator,omitempty"`
	//Operator *string `json:"operator,omitempty"`
	//RootAclIpv4filterEntryMatchSourceportRange
	Range *DeviceAclIpv4filterEntryMatchSourceportRange `json:"range,omitempty"`
	Value *string                                       `json:"value,omitempty"`
}

// DeviceAclIpv4filterEntryMatchSourceportRange struct
type DeviceAclIpv4filterEntryMatchSourceportRange struct {
	End   *string `json:"end,omitempty"`
	Start *string `json:"start,omitempty"`
}

// DeviceAclIpv6filter struct
type DeviceAclIpv6filter struct {
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Description *string `json:"description,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Entry []*DeviceAclIpv6filterEntry `json:"entry,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name               *string `json:"name"`
	Statisticsperentry *bool   `json:"statistics-per-entry,omitempty"`
	// +kubebuilder:validation:Enum=`disabled`;`input-and-output`;`input-only`;`output-only`
	// +kubebuilder:default:="disabled"
	Subinterfacespecific E_DeviceAclIpv6filterSubinterfacespecific `json:"subinterface-specific,omitempty"`
	//Subinterfacespecific *string `json:"subinterface-specific,omitempty"`
}

// DeviceAclIpv6filterEntry struct
type DeviceAclIpv6filterEntry struct {
	//RootAclIpv6filterEntryAction
	Action *DeviceAclIpv6filterEntryAction `json:"action,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Description *string `json:"description,omitempty"`
	//RootAclIpv6filterEntryMatch
	Match *DeviceAclIpv6filterEntryMatch `json:"match,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=65535
	Sequenceid *uint32 `json:"sequence-id"`
}

// DeviceAclIpv6filterEntryAction struct
type DeviceAclIpv6filterEntryAction struct {
	//RootAclIpv6filterEntryActionAccept
	Accept *DeviceAclIpv6filterEntryActionAccept `json:"accept,omitempty"`
	//RootAclIpv6filterEntryActionDrop
	Drop *DeviceAclIpv6filterEntryActionDrop `json:"drop,omitempty"`
}

// DeviceAclIpv6filterEntryActionAccept struct
type DeviceAclIpv6filterEntryActionAccept struct {
	// +kubebuilder:default:=false
	Log *bool `json:"log,omitempty"`
}

// DeviceAclIpv6filterEntryActionDrop struct
type DeviceAclIpv6filterEntryActionDrop struct {
	// +kubebuilder:default:=false
	Log *bool `json:"log,omitempty"`
}

// DeviceAclIpv6filterEntryMatch struct
type DeviceAclIpv6filterEntryMatch struct {
	//RootAclIpv6filterEntryMatchDestinationip
	Destinationip *DeviceAclIpv6filterEntryMatchDestinationip `json:"destination-ip,omitempty"`
	//RootAclIpv6filterEntryMatchDestinationport
	Destinationport *DeviceAclIpv6filterEntryMatchDestinationport `json:"destination-port,omitempty"`
	//RootAclIpv6filterEntryMatchIcmp6
	Icmp6      *DeviceAclIpv6filterEntryMatchIcmp6 `json:"icmp6,omitempty"`
	Nextheader *string                             `json:"next-header,omitempty"`
	//RootAclIpv6filterEntryMatchSourceip
	Sourceip *DeviceAclIpv6filterEntryMatchSourceip `json:"source-ip,omitempty"`
	//RootAclIpv6filterEntryMatchSourceport
	Sourceport *DeviceAclIpv6filterEntryMatchSourceport `json:"source-port,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(\(|\)|&|\||!|ack|rst|syn)+`
	Tcpflags *string `json:"tcp-flags,omitempty"`
}

// DeviceAclIpv6filterEntryMatchDestinationip struct
type DeviceAclIpv6filterEntryMatchDestinationip struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Address *string `json:"address,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Mask *string `json:"mask,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	Prefix *string `json:"prefix,omitempty"`
}

// DeviceAclIpv6filterEntryMatchDestinationport struct
type DeviceAclIpv6filterEntryMatchDestinationport struct {
	// +kubebuilder:validation:Enum=`eq`;`ge`;`le`
	Operator E_DeviceAclIpv6filterEntryMatchDestinationportOperator `json:"operator,omitempty"`
	//Operator *string `json:"operator,omitempty"`
	//RootAclIpv6filterEntryMatchDestinationportRange
	Range *DeviceAclIpv6filterEntryMatchDestinationportRange `json:"range,omitempty"`
	Value *string                                            `json:"value,omitempty"`
}

// DeviceAclIpv6filterEntryMatchDestinationportRange struct
type DeviceAclIpv6filterEntryMatchDestinationportRange struct {
	End   *string `json:"end,omitempty"`
	Start *string `json:"start,omitempty"`
}

// DeviceAclIpv6filterEntryMatchIcmp6 struct
type DeviceAclIpv6filterEntryMatchIcmp6 struct {
	//RootAclIpv6filterEntryMatchIcmp6Code
	Code *DeviceAclIpv6filterEntryMatchIcmp6Code `json:"code,omitempty"`
	Type *string                                 `json:"type,omitempty"`
}

// DeviceAclIpv6filterEntryMatchIcmp6Code struct
type DeviceAclIpv6filterEntryMatchIcmp6Code struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	Code *uint8 `json:"code,omitempty"`
}

// DeviceAclIpv6filterEntryMatchSourceip struct
type DeviceAclIpv6filterEntryMatchSourceip struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Address *string `json:"address,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Mask *string `json:"mask,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	Prefix *string `json:"prefix,omitempty"`
}

// DeviceAclIpv6filterEntryMatchSourceport struct
type DeviceAclIpv6filterEntryMatchSourceport struct {
	// +kubebuilder:validation:Enum=`eq`;`ge`;`le`
	Operator E_DeviceAclIpv6filterEntryMatchSourceportOperator `json:"operator,omitempty"`
	//Operator *string `json:"operator,omitempty"`
	//RootAclIpv6filterEntryMatchSourceportRange
	Range *DeviceAclIpv6filterEntryMatchSourceportRange `json:"range,omitempty"`
	Value *string                                       `json:"value,omitempty"`
}

// DeviceAclIpv6filterEntryMatchSourceportRange struct
type DeviceAclIpv6filterEntryMatchSourceportRange struct {
	End   *string `json:"end,omitempty"`
	Start *string `json:"start,omitempty"`
}

// DeviceAclPolicers struct
type DeviceAclPolicers struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Policer []*DeviceAclPolicersPolicer `json:"policer,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Systemcpupolicer []*DeviceAclPolicersSystemcpupolicer `json:"system-cpu-policer,omitempty"`
}

// DeviceAclPolicersPolicer struct
type DeviceAclPolicersPolicer struct {
	// +kubebuilder:default:=false
	Entryspecific *bool `json:"entry-specific,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=125000000
	Maxburst *uint32 `json:"max-burst,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=1000000
	Peakrate *uint32 `json:"peak-rate,omitempty"`
}

// DeviceAclPolicersSystemcpupolicer struct
type DeviceAclPolicersSystemcpupolicer struct {
	// +kubebuilder:default:=false
	Entryspecific *bool `json:"entry-specific,omitempty"`
	// kubebuilder:validation:Minimum=16
	// kubebuilder:validation:Maximum=4000000
	// +kubebuilder:default:=16
	Maxpacketburst *uint32 `json:"max-packet-burst,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4000000
	Peakpacketrate *uint32 `json:"peak-packet-rate,omitempty"`
}

// DeviceAclSystemfilter struct
type DeviceAclSystemfilter struct {
	//RootAclSystemfilterIpv4filter
	Ipv4filter *DeviceAclSystemfilterIpv4filter `json:"ipv4-filter,omitempty"`
	//RootAclSystemfilterIpv6filter
	Ipv6filter *DeviceAclSystemfilterIpv6filter `json:"ipv6-filter,omitempty"`
}

// DeviceAclSystemfilterIpv4filter struct
type DeviceAclSystemfilterIpv4filter struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Entry []*DeviceAclSystemfilterIpv4filterEntry `json:"entry,omitempty"`
}

// DeviceAclSystemfilterIpv4filterEntry struct
type DeviceAclSystemfilterIpv4filterEntry struct {
	//RootAclSystemfilterIpv4filterEntryAction
	Action *DeviceAclSystemfilterIpv4filterEntryAction `json:"action,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Description *string `json:"description,omitempty"`
	//RootAclSystemfilterIpv4filterEntryMatch
	Match *DeviceAclSystemfilterIpv4filterEntryMatch `json:"match,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=256
	Sequenceid *uint32 `json:"sequence-id"`
}

// DeviceAclSystemfilterIpv4filterEntryAction struct
type DeviceAclSystemfilterIpv4filterEntryAction struct {
	//RootAclSystemfilterIpv4filterEntryActionAccept
	Accept *DeviceAclSystemfilterIpv4filterEntryActionAccept `json:"accept,omitempty"`
	//RootAclSystemfilterIpv4filterEntryActionDrop
	Drop *DeviceAclSystemfilterIpv4filterEntryActionDrop `json:"drop,omitempty"`
}

// DeviceAclSystemfilterIpv4filterEntryActionAccept struct
type DeviceAclSystemfilterIpv4filterEntryActionAccept struct {
}

// DeviceAclSystemfilterIpv4filterEntryActionDrop struct
type DeviceAclSystemfilterIpv4filterEntryActionDrop struct {
	// +kubebuilder:default:=false
	Log *bool `json:"log,omitempty"`
}

// DeviceAclSystemfilterIpv4filterEntryMatch struct
type DeviceAclSystemfilterIpv4filterEntryMatch struct {
	//RootAclSystemfilterIpv4filterEntryMatchDestinationip
	Destinationip *DeviceAclSystemfilterIpv4filterEntryMatchDestinationip `json:"destination-ip,omitempty"`
	//RootAclSystemfilterIpv4filterEntryMatchDestinationport
	Destinationport *DeviceAclSystemfilterIpv4filterEntryMatchDestinationport `json:"destination-port,omitempty"`
	Firstfragment   *bool                                                     `json:"first-fragment,omitempty"`
	Fragment        *bool                                                     `json:"fragment,omitempty"`
	//RootAclSystemfilterIpv4filterEntryMatchIcmp
	Icmp     *DeviceAclSystemfilterIpv4filterEntryMatchIcmp `json:"icmp,omitempty"`
	Protocol *string                                        `json:"protocol,omitempty"`
	//RootAclSystemfilterIpv4filterEntryMatchSourceip
	Sourceip *DeviceAclSystemfilterIpv4filterEntryMatchSourceip `json:"source-ip,omitempty"`
	//RootAclSystemfilterIpv4filterEntryMatchSourceport
	Sourceport *DeviceAclSystemfilterIpv4filterEntryMatchSourceport `json:"source-port,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(\(|\)|&|\||!|ack|rst|syn)+`
	Tcpflags *string `json:"tcp-flags,omitempty"`
}

// DeviceAclSystemfilterIpv4filterEntryMatchDestinationip struct
type DeviceAclSystemfilterIpv4filterEntryMatchDestinationip struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Address *string `json:"address,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Mask *string `json:"mask,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))`
	Prefix *string `json:"prefix,omitempty"`
}

// DeviceAclSystemfilterIpv4filterEntryMatchDestinationport struct
type DeviceAclSystemfilterIpv4filterEntryMatchDestinationport struct {
	// +kubebuilder:validation:Enum=`eq`;`ge`;`le`
	Operator E_DeviceAclSystemfilterIpv4filterEntryMatchDestinationportOperator `json:"operator,omitempty"`
	//Operator *string `json:"operator,omitempty"`
	//RootAclSystemfilterIpv4filterEntryMatchDestinationportRange
	Range *DeviceAclSystemfilterIpv4filterEntryMatchDestinationportRange `json:"range,omitempty"`
	Value *string                                                        `json:"value,omitempty"`
}

// DeviceAclSystemfilterIpv4filterEntryMatchDestinationportRange struct
type DeviceAclSystemfilterIpv4filterEntryMatchDestinationportRange struct {
	End   *string `json:"end,omitempty"`
	Start *string `json:"start,omitempty"`
}

// DeviceAclSystemfilterIpv4filterEntryMatchIcmp struct
type DeviceAclSystemfilterIpv4filterEntryMatchIcmp struct {
	//RootAclSystemfilterIpv4filterEntryMatchIcmpCode
	Code *DeviceAclSystemfilterIpv4filterEntryMatchIcmpCode `json:"code,omitempty"`
	Type *string                                            `json:"type,omitempty"`
}

// DeviceAclSystemfilterIpv4filterEntryMatchIcmpCode struct
type DeviceAclSystemfilterIpv4filterEntryMatchIcmpCode struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	Code *uint8 `json:"code,omitempty"`
}

// DeviceAclSystemfilterIpv4filterEntryMatchSourceip struct
type DeviceAclSystemfilterIpv4filterEntryMatchSourceip struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Address *string `json:"address,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Mask *string `json:"mask,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))`
	Prefix *string `json:"prefix,omitempty"`
}

// DeviceAclSystemfilterIpv4filterEntryMatchSourceport struct
type DeviceAclSystemfilterIpv4filterEntryMatchSourceport struct {
	// +kubebuilder:validation:Enum=`eq`;`ge`;`le`
	Operator E_DeviceAclSystemfilterIpv4filterEntryMatchSourceportOperator `json:"operator,omitempty"`
	//Operator *string `json:"operator,omitempty"`
	//RootAclSystemfilterIpv4filterEntryMatchSourceportRange
	Range *DeviceAclSystemfilterIpv4filterEntryMatchSourceportRange `json:"range,omitempty"`
	Value *string                                                   `json:"value,omitempty"`
}

// DeviceAclSystemfilterIpv4filterEntryMatchSourceportRange struct
type DeviceAclSystemfilterIpv4filterEntryMatchSourceportRange struct {
	End   *string `json:"end,omitempty"`
	Start *string `json:"start,omitempty"`
}

// DeviceAclSystemfilterIpv6filter struct
type DeviceAclSystemfilterIpv6filter struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Entry []*DeviceAclSystemfilterIpv6filterEntry `json:"entry,omitempty"`
}

// DeviceAclSystemfilterIpv6filterEntry struct
type DeviceAclSystemfilterIpv6filterEntry struct {
	//RootAclSystemfilterIpv6filterEntryAction
	Action *DeviceAclSystemfilterIpv6filterEntryAction `json:"action,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Description *string `json:"description,omitempty"`
	//RootAclSystemfilterIpv6filterEntryMatch
	Match *DeviceAclSystemfilterIpv6filterEntryMatch `json:"match,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=128
	Sequenceid *uint32 `json:"sequence-id"`
}

// DeviceAclSystemfilterIpv6filterEntryAction struct
type DeviceAclSystemfilterIpv6filterEntryAction struct {
	//RootAclSystemfilterIpv6filterEntryActionAccept
	Accept *DeviceAclSystemfilterIpv6filterEntryActionAccept `json:"accept,omitempty"`
	//RootAclSystemfilterIpv6filterEntryActionDrop
	Drop *DeviceAclSystemfilterIpv6filterEntryActionDrop `json:"drop,omitempty"`
}

// DeviceAclSystemfilterIpv6filterEntryActionAccept struct
type DeviceAclSystemfilterIpv6filterEntryActionAccept struct {
}

// DeviceAclSystemfilterIpv6filterEntryActionDrop struct
type DeviceAclSystemfilterIpv6filterEntryActionDrop struct {
	// +kubebuilder:default:=false
	Log *bool `json:"log,omitempty"`
}

// DeviceAclSystemfilterIpv6filterEntryMatch struct
type DeviceAclSystemfilterIpv6filterEntryMatch struct {
	//RootAclSystemfilterIpv6filterEntryMatchDestinationip
	Destinationip *DeviceAclSystemfilterIpv6filterEntryMatchDestinationip `json:"destination-ip,omitempty"`
	//RootAclSystemfilterIpv6filterEntryMatchDestinationport
	Destinationport *DeviceAclSystemfilterIpv6filterEntryMatchDestinationport `json:"destination-port,omitempty"`
	//RootAclSystemfilterIpv6filterEntryMatchIcmp6
	Icmp6      *DeviceAclSystemfilterIpv6filterEntryMatchIcmp6 `json:"icmp6,omitempty"`
	Nextheader *string                                         `json:"next-header,omitempty"`
	//RootAclSystemfilterIpv6filterEntryMatchSourceip
	Sourceip *DeviceAclSystemfilterIpv6filterEntryMatchSourceip `json:"source-ip,omitempty"`
	//RootAclSystemfilterIpv6filterEntryMatchSourceport
	Sourceport *DeviceAclSystemfilterIpv6filterEntryMatchSourceport `json:"source-port,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(\(|\)|&|\||!|ack|rst|syn)+`
	Tcpflags *string `json:"tcp-flags,omitempty"`
}

// DeviceAclSystemfilterIpv6filterEntryMatchDestinationip struct
type DeviceAclSystemfilterIpv6filterEntryMatchDestinationip struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Address *string `json:"address,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Mask *string `json:"mask,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	Prefix *string `json:"prefix,omitempty"`
}

// DeviceAclSystemfilterIpv6filterEntryMatchDestinationport struct
type DeviceAclSystemfilterIpv6filterEntryMatchDestinationport struct {
	// +kubebuilder:validation:Enum=`eq`;`ge`;`le`
	Operator E_DeviceAclSystemfilterIpv6filterEntryMatchDestinationportOperator `json:"operator,omitempty"`
	//Operator *string `json:"operator,omitempty"`
	//RootAclSystemfilterIpv6filterEntryMatchDestinationportRange
	Range *DeviceAclSystemfilterIpv6filterEntryMatchDestinationportRange `json:"range,omitempty"`
	Value *string                                                        `json:"value,omitempty"`
}

// DeviceAclSystemfilterIpv6filterEntryMatchDestinationportRange struct
type DeviceAclSystemfilterIpv6filterEntryMatchDestinationportRange struct {
	End   *string `json:"end,omitempty"`
	Start *string `json:"start,omitempty"`
}

// DeviceAclSystemfilterIpv6filterEntryMatchIcmp6 struct
type DeviceAclSystemfilterIpv6filterEntryMatchIcmp6 struct {
	//RootAclSystemfilterIpv6filterEntryMatchIcmp6Code
	Code *DeviceAclSystemfilterIpv6filterEntryMatchIcmp6Code `json:"code,omitempty"`
	Type *string                                             `json:"type,omitempty"`
}

// DeviceAclSystemfilterIpv6filterEntryMatchIcmp6Code struct
type DeviceAclSystemfilterIpv6filterEntryMatchIcmp6Code struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	Code *uint8 `json:"code,omitempty"`
}

// DeviceAclSystemfilterIpv6filterEntryMatchSourceip struct
type DeviceAclSystemfilterIpv6filterEntryMatchSourceip struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Address *string `json:"address,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Mask *string `json:"mask,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	Prefix *string `json:"prefix,omitempty"`
}

// DeviceAclSystemfilterIpv6filterEntryMatchSourceport struct
type DeviceAclSystemfilterIpv6filterEntryMatchSourceport struct {
	// +kubebuilder:validation:Enum=`eq`;`ge`;`le`
	Operator E_DeviceAclSystemfilterIpv6filterEntryMatchSourceportOperator `json:"operator,omitempty"`
	//Operator *string `json:"operator,omitempty"`
	//RootAclSystemfilterIpv6filterEntryMatchSourceportRange
	Range *DeviceAclSystemfilterIpv6filterEntryMatchSourceportRange `json:"range,omitempty"`
	Value *string                                                   `json:"value,omitempty"`
}

// DeviceAclSystemfilterIpv6filterEntryMatchSourceportRange struct
type DeviceAclSystemfilterIpv6filterEntryMatchSourceportRange struct {
	End   *string `json:"end,omitempty"`
	Start *string `json:"start,omitempty"`
}

// DeviceBfd struct
type DeviceBfd struct {
	//RootBfdMicrobfdsessions
	Microbfdsessions *DeviceBfdMicrobfdsessions `json:"micro-bfd-sessions,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Subinterface []*DeviceBfdSubinterface `json:"subinterface,omitempty"`
}

// DeviceBfdMicrobfdsessions struct
type DeviceBfdMicrobfdsessions struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Laginterface []*DeviceBfdMicrobfdsessionsLaginterface `json:"lag-interface,omitempty"`
}

// DeviceBfdMicrobfdsessionsLaginterface struct
type DeviceBfdMicrobfdsessionsLaginterface struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_DeviceBfdMicrobfdsessionsLaginterfaceAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=10000
	// kubebuilder:validation:Maximum=100000000
	// +kubebuilder:default:=1000000
	Desiredminimumtransmitinterval *uint32 `json:"desired-minimum-transmit-interval,omitempty"`
	// kubebuilder:validation:Minimum=3
	// kubebuilder:validation:Maximum=20
	// +kubebuilder:default:=3
	Detectionmultiplier *uint8 `json:"detection-multiplier,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Localaddress *string `json:"local-address,omitempty"`
	Name         *string `json:"name"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Remoteaddress *string `json:"remote-address,omitempty"`
	// kubebuilder:validation:Minimum=10000
	// kubebuilder:validation:Maximum=100000000
	// +kubebuilder:default:=1000000
	Requiredminimumreceive *uint32 `json:"required-minimum-receive,omitempty"`
}

// DeviceBfdSubinterface struct
type DeviceBfdSubinterface struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_DeviceBfdSubinterfaceAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=10000
	// kubebuilder:validation:Maximum=100000000
	// +kubebuilder:default:=1000000
	Desiredminimumtransmitinterval *uint32 `json:"desired-minimum-transmit-interval,omitempty"`
	// kubebuilder:validation:Minimum=3
	// kubebuilder:validation:Maximum=20
	// +kubebuilder:default:=3
	Detectionmultiplier *uint8 `json:"detection-multiplier,omitempty"`
	// kubebuilder:validation:MinLength=5
	// kubebuilder:validation:MaxLength=25
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(system0\.0|lo(0|1[0-9][0-9]|2([0-4][0-9]|5[0-5])|[1-9][0-9]|[1-9])\.(0|[1-9](\d){0,3})|ethernet-([1-9](\d){0,1}(/[abcd])?(/[1-9](\d){0,1})?/(([1-9](\d){0,1})|(1[0-1]\d)|(12[0-8])))\.([0]|[1-9](\d){0,3})|irb(0|1[0-9][0-9]|2([0-4][0-9]|5[0-5])|[1-9][0-9]|[1-9])\.(0|[1-9](\d){0,3})|lag(([1-9](\d){0,1})|(1[0-1]\d)|(12[0-8]))\.(0|[1-9](\d){0,3}))`
	Id *string `json:"id"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=0
	// +kubebuilder:default:=0
	Minimumechoreceiveinterval *uint32 `json:"minimum-echo-receive-interval,omitempty"`
	// kubebuilder:validation:Minimum=10000
	// kubebuilder:validation:Maximum=100000000
	// +kubebuilder:default:=1000000
	Requiredminimumreceive *uint32 `json:"required-minimum-receive,omitempty"`
}

// DeviceInterface struct
type DeviceInterface struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_DeviceInterfaceAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	//RootInterfaceBreakoutmode
	Breakoutmode *DeviceInterfaceBreakoutmode `json:"breakout-mode,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Description *string `json:"description,omitempty"`
	//RootInterfaceEthernet
	Ethernet *DeviceInterfaceEthernet `json:"ethernet,omitempty"`
	//RootInterfaceLag
	Lag *DeviceInterfaceLag `json:"lag,omitempty"`
	//RootInterfaceLinux
	Linux        *DeviceInterfaceLinux `json:"linux,omitempty"`
	Loopbackmode *bool                 `json:"loopback-mode,omitempty"`
	// kubebuilder:validation:Minimum=1500
	// kubebuilder:validation:Maximum=9500
	Mtu *uint16 `json:"mtu,omitempty"`
	// kubebuilder:validation:MinLength=3
	// kubebuilder:validation:MaxLength=20
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(mgmt0|mgmt0-standby|system0|lo(0|1[0-9][0-9]|2([0-4][0-9]|5[0-5])|[1-9][0-9]|[1-9])|ethernet-([1-9](\d){0,1}(/[abcd])?(/[1-9](\d){0,1})?/(([1-9](\d){0,1})|(1[0-1]\d)|(12[0-8])))|irb(0|1[0-9][0-9]|2([0-4][0-9]|5[0-5])|[1-9][0-9]|[1-9])|lag(([1-9](\d){0,1})|(1[0-1]\d)|(12[0-8])))`
	Name *string `json:"name"`
	//RootInterfaceP4rt
	P4rt *DeviceInterfaceP4rt `json:"p4rt,omitempty"`
	//RootInterfaceQos
	Qos *DeviceInterfaceQos `json:"qos,omitempty"`
	//RootInterfaceRadio
	Radio *DeviceInterfaceRadio `json:"radio,omitempty"`
	//RootInterfaceSflow
	Sflow *DeviceInterfaceSflow `json:"sflow,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=4095
	Subinterface []*DeviceInterfaceSubinterface `json:"subinterface,omitempty"`
	//RootInterfaceTransceiver
	Transceiver *DeviceInterfaceTransceiver `json:"transceiver,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`
	Uuid *string `json:"uuid,omitempty"`
	//RootInterfaceVhost
	Vhost       *DeviceInterfaceVhost `json:"vhost,omitempty"`
	Vlantagging *bool                 `json:"vlan-tagging,omitempty"`
}

// DeviceInterfaceBreakoutmode struct
type DeviceInterfaceBreakoutmode struct {
	// +kubebuilder:validation:Enum=`10G`;`25G`
	Channelspeed E_DeviceInterfaceBreakoutmodeChannelspeed `json:"channel-speed,omitempty"`
	// +kubebuilder:validation:Enum=`4`
	Numchannels E_DeviceInterfaceBreakoutmodeNumchannels `json:"num-channels,omitempty"`
}

// DeviceInterfaceEthernet struct
type DeviceInterfaceEthernet struct {
	Aggregateid   *string `json:"aggregate-id,omitempty"`
	Autonegotiate *bool   `json:"auto-negotiate,omitempty"`
	// +kubebuilder:validation:Enum=`full`;`half`
	Duplexmode E_DeviceInterfaceEthernetDuplexmode `json:"duplex-mode,omitempty"`
	//Duplexmode *string `json:"duplex-mode,omitempty"`
	//RootInterfaceEthernetFlowcontrol
	Flowcontrol      *DeviceInterfaceEthernetFlowcontrol `json:"flow-control,omitempty"`
	Forwardingviable *bool                               `json:"forwarding-viable,omitempty"`
	//RootInterfaceEthernetHoldtime
	Holdtime *DeviceInterfaceEthernetHoldtime `json:"hold-time,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	Lacpportpriority *uint16 `json:"lacp-port-priority,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}`
	Macaddress *string `json:"mac-address,omitempty"`
	// +kubebuilder:validation:Enum=`100G`;`100M`;`10G`;`10M`;`1G`;`1T`;`200G`;`25G`;`400G`;`40G`;`50G`
	Portspeed E_DeviceInterfaceEthernetPortspeed `json:"port-speed,omitempty"`
	//Portspeed *string `json:"port-speed,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=86400
	Reloaddelay *uint32 `json:"reload-delay,omitempty"`
	// +kubebuilder:validation:Enum=`lacp`;`power-off`
	Standbysignaling E_DeviceInterfaceEthernetStandbysignaling `json:"standby-signaling,omitempty"`
	//Standbysignaling *string `json:"standby-signaling,omitempty"`
	//RootInterfaceEthernetStormcontrol
	Stormcontrol *DeviceInterfaceEthernetStormcontrol `json:"storm-control,omitempty"`
	//RootInterfaceEthernetSynce
	Synce *DeviceInterfaceEthernetSynce `json:"synce,omitempty"`
}

// DeviceInterfaceEthernetFlowcontrol struct
type DeviceInterfaceEthernetFlowcontrol struct {
	Receive  *bool `json:"receive,omitempty"`
	Transmit *bool `json:"transmit,omitempty"`
}

// DeviceInterfaceEthernetHoldtime struct
type DeviceInterfaceEthernetHoldtime struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=86400
	Down *uint32 `json:"down,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=86400
	Up *uint32 `json:"up,omitempty"`
}

// DeviceInterfaceEthernetStormcontrol struct
type DeviceInterfaceEthernetStormcontrol struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100000000
	Broadcastrate *uint32 `json:"broadcast-rate,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100000000
	Multicastrate *uint32 `json:"multicast-rate,omitempty"`
	// +kubebuilder:validation:Enum=`kbps`;`percentage`
	// +kubebuilder:default:="percentage"
	Units E_DeviceInterfaceEthernetStormcontrolUnits `json:"units,omitempty"`
	//Units *string `json:"units,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100000000
	Unknownunicastrate *uint32 `json:"unknown-unicast-rate,omitempty"`
}

// DeviceInterfaceEthernetSynce struct
type DeviceInterfaceEthernetSynce struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_DeviceInterfaceEthernetSynceAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
}

// DeviceInterfaceLag struct
type DeviceInterfaceLag struct {
	//RootInterfaceLagLacp
	Lacp *DeviceInterfaceLagLacp `json:"lacp,omitempty"`
	// +kubebuilder:validation:Enum=`static`
	Lacpfallbackmode E_DeviceInterfaceLagLacpfallbackmode `json:"lacp-fallback-mode,omitempty"`
	//Lacpfallbackmode *string `json:"lacp-fallback-mode,omitempty"`
	// kubebuilder:validation:Minimum=4
	// kubebuilder:validation:Maximum=3600
	Lacpfallbacktimeout *uint16 `json:"lacp-fallback-timeout,omitempty"`
	// +kubebuilder:validation:Enum=`lacp`;`static`
	// +kubebuilder:default:="static"
	Lagtype E_DeviceInterfaceLagLagtype `json:"lag-type,omitempty"`
	//Lagtype *string `json:"lag-type,omitempty"`
	// +kubebuilder:validation:Enum=`100G`;`100M`;`10G`;`10M`;`1G`;`25G`;`400G`;`40G`
	Memberspeed E_DeviceInterfaceLagMemberspeed `json:"member-speed,omitempty"`
	//Memberspeed *string `json:"member-speed,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=64
	// +kubebuilder:default:=1
	Minlinks *uint16 `json:"min-links,omitempty"`
}

// DeviceInterfaceLagLacp struct
type DeviceInterfaceLagLacp struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=65535
	Adminkey *uint16 `json:"admin-key,omitempty"`
	// +kubebuilder:validation:Enum=`FAST`;`SLOW`
	// +kubebuilder:default:="SLOW"
	Interval E_DeviceInterfaceLagLacpInterval `json:"interval,omitempty"`
	//Interval *string `json:"interval,omitempty"`
	// +kubebuilder:validation:Enum=`ACTIVE`;`PASSIVE`
	// +kubebuilder:default:="ACTIVE"
	Lacpmode E_DeviceInterfaceLagLacpLacpmode `json:"lacp-mode,omitempty"`
	//Lacpmode *string `json:"lacp-mode,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}`
	Systemidmac *string `json:"system-id-mac,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	Systempriority *uint16 `json:"system-priority,omitempty"`
}

// DeviceInterfaceLinux struct
type DeviceInterfaceLinux struct {
	// kubebuilder:validation:MinLength=0
	// kubebuilder:validation:MaxLength=10
	Devicename *string `json:"device-name,omitempty"`
}

// DeviceInterfaceP4rt struct
type DeviceInterfaceP4rt struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	Id *uint32 `json:"id,omitempty"`
}

// DeviceInterfaceQos struct
type DeviceInterfaceQos struct {
	//RootInterfaceQosOutput
	Output *DeviceInterfaceQosOutput `json:"output,omitempty"`
}

// DeviceInterfaceQosOutput struct
type DeviceInterfaceQosOutput struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Multicastqueue []*DeviceInterfaceQosOutputMulticastqueue `json:"multicast-queue,omitempty"`
	//RootInterfaceQosOutputScheduler
	Scheduler *DeviceInterfaceQosOutputScheduler `json:"scheduler,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Unicastqueue []*DeviceInterfaceQosOutputUnicastqueue `json:"unicast-queue,omitempty"`
}

// DeviceInterfaceQosOutputMulticastqueue struct
type DeviceInterfaceQosOutputMulticastqueue struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=7
	Queueid *uint8 `json:"queue-id"`
	//RootInterfaceQosOutputMulticastqueueScheduling
	Scheduling *DeviceInterfaceQosOutputMulticastqueueScheduling `json:"scheduling,omitempty"`
	Template   *string                                           `json:"template,omitempty"`
}

// DeviceInterfaceQosOutputMulticastqueueScheduling struct
type DeviceInterfaceQosOutputMulticastqueueScheduling struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=100
	Peakratepercent *uint8 `json:"peak-rate-percent,omitempty"`
}

// DeviceInterfaceQosOutputScheduler struct
type DeviceInterfaceQosOutputScheduler struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1
	Tier []*DeviceInterfaceQosOutputSchedulerTier `json:"tier,omitempty"`
}

// DeviceInterfaceQosOutputSchedulerTier struct
type DeviceInterfaceQosOutputSchedulerTier struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4
	Level *uint8 `json:"level"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=12
	Node []*DeviceInterfaceQosOutputSchedulerTierNode `json:"node,omitempty"`
}

// DeviceInterfaceQosOutputSchedulerTierNode struct
type DeviceInterfaceQosOutputSchedulerTierNode struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=11
	Nodenumber     *uint8 `json:"node-number"`
	Strictpriority *bool  `json:"strict-priority,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=127
	// +kubebuilder:default:=1
	Weight *uint8 `json:"weight,omitempty"`
}

// DeviceInterfaceQosOutputUnicastqueue struct
type DeviceInterfaceQosOutputUnicastqueue struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=7
	Queueid *uint8 `json:"queue-id"`
	//RootInterfaceQosOutputUnicastqueueScheduling
	Scheduling  *DeviceInterfaceQosOutputUnicastqueueScheduling `json:"scheduling,omitempty"`
	Template    *string                                         `json:"template,omitempty"`
	Voqtemplate *string                                         `json:"voq-template,omitempty"`
}

// DeviceInterfaceQosOutputUnicastqueueScheduling struct
type DeviceInterfaceQosOutputUnicastqueueScheduling struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=100
	Peakratepercent *uint8 `json:"peak-rate-percent,omitempty"`
	// +kubebuilder:default:=true
	Strictpriority *bool `json:"strict-priority,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=1
	Weight *uint8 `json:"weight,omitempty"`
}

// DeviceInterfaceRadio struct
type DeviceInterfaceRadio struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_DeviceInterfaceRadioAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	//RootInterfaceRadioRoe
	Roe *DeviceInterfaceRadioRoe `json:"roe,omitempty"`
}

// DeviceInterfaceRadioRoe struct
type DeviceInterfaceRadioRoe struct {
	//RootInterfaceRadioRoeCpriport
	Cpriport *DeviceInterfaceRadioRoeCpriport `json:"cpri-port,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Ethernetlink []*DeviceInterfaceRadioRoeEthernetlink `json:"ethernet-link,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Mapperdemapper []*DeviceInterfaceRadioRoeMapperdemapper `json:"mapper-demapper,omitempty"`
	//RootInterfaceRadioRoeObsaiport
	Obsaiport *DeviceInterfaceRadioRoeObsaiport `json:"obsai-port,omitempty"`
	// +kubebuilder:validation:Enum=`follower`;`leader`
	// +kubebuilder:default:="follower"
	Portrole E_DeviceInterfaceRadioRoePortrole `json:"port-role,omitempty"`
	//Portrole *string `json:"port-role,omitempty"`
	//RootInterfaceRadioRoePrestimeoffset
	Prestimeoffset *DeviceInterfaceRadioRoePrestimeoffset `json:"pres-time-offset,omitempty"`
	//RootInterfaceRadioRoeTargetoffset
	Targetoffset *DeviceInterfaceRadioRoeTargetoffset `json:"target-offset,omitempty"`
}

// DeviceInterfaceRadioRoeCpriport struct
type DeviceInterfaceRadioRoeCpriport struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=0
	Cprieth *uint8 `json:"cpri-eth,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=7
	Cprihdlc *uint8 `json:"cpri-hdlc,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	Cpriver *uint8 `json:"cpri-ver,omitempty"`
	//RootInterfaceRadioRoeCpriportLpt
	Lpt *DeviceInterfaceRadioRoeCpriportLpt `json:"lpt,omitempty"`
	// +kubebuilder:default:=true
	Portparamauto *bool `json:"port-param-auto,omitempty"`
	// +kubebuilder:validation:Enum=`cpri-10`;`cpri-3`;`cpri-5`;`cpri-7`;`cpri-8`
	Portspeed E_DeviceInterfaceRadioRoeCpriportPortspeed `json:"port-speed,omitempty"`
	//Portspeed *string `json:"port-speed,omitempty"`
}

// DeviceInterfaceRadioRoeCpriportLpt struct
type DeviceInterfaceRadioRoeCpriportLpt struct {
	// +kubebuilder:validation:Enum=`cpritxdis`;`disable`;`lasershut`
	Consaction E_DeviceInterfaceRadioRoeCpriportLptConsaction `json:"cons-action,omitempty"`
	//Consaction *string `json:"cons-action,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=300000
	// +kubebuilder:default:=0
	Extensiontime *uint32 `json:"extension-time,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=300000
	// +kubebuilder:default:=0
	Lasershutdelaytime *uint32 `json:"laser-shut-delay-time,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=300000
	// +kubebuilder:default:=0
	Txdisdelaytime *uint32 `json:"tx-dis-delay-time,omitempty"`
}

// DeviceInterfaceRadioRoeEthernetlink struct
type DeviceInterfaceRadioRoeEthernetlink struct {
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Description *string `json:"description,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=7
	// +kubebuilder:default:=7
	Dot1p *uint8 `json:"dot1p,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=48
	Ethernetid *uint32 `json:"ethernet-id"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}`
	Localmacaddress *string `json:"local-mac-address,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}`
	Remotemacaddress *string `json:"remote-mac-address,omitempty"`
	Subinterface     *string `json:"subinterface,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4094
	Vlanid *uint16 `json:"vlan-id,omitempty"`
}

// DeviceInterfaceRadioRoeMapperdemapper struct
type DeviceInterfaceRadioRoeMapperdemapper struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	Axcposition *uint16 `json:"axc-position,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=6
	Bandwidth *uint8 `json:"bandwidth,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1
	Ctrlwordconts []*DeviceInterfaceRadioRoeMapperdemapperCtrlwordconts `json:"ctrl-word-conts,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Description *string `json:"description,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	Ethlinkid *uint32 `json:"ethlink-id,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=254
	// +kubebuilder:default:=0
	Flowid *uint8 `json:"flow-id,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=128
	Mapperdemapperid *uint32 `json:"mapper-demapper-id"`
	// +kubebuilder:validation:Enum=`demapper`;`mapper`
	Mapperdemappertype E_DeviceInterfaceRadioRoeMapperdemapperMapperdemappertype `json:"mapper-demapper-type,omitempty"`
	//Mapperdemappertype *string `json:"mapper-demapper-type,omitempty"`
	// +kubebuilder:validation:Enum=`NATIVE-FREQUENCY-DOMAIN`;`NATIVE-TIME-DOMAIN`;`STRUCTURE-AGNOSTIC-LINE-CODING-AWARE`;`STRUCTURE-AGNOSTIC-TUNNELING`;`STRUCTURE-AWARE-FREQUENCY-DOMAIN-CPRI`;`STRUCTURE-AWARE-TIME-DOMAIN-CPRI`
	Mappertype E_DeviceInterfaceRadioRoeMapperdemapperMappertype `json:"mapper-type,omitempty"`
	//Mappertype *string `json:"mapper-type,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	Payloadlen *uint16 `json:"payload-len,omitempty"`
	// kubebuilder:validation:Minimum=4
	// kubebuilder:validation:Maximum=16
	// +kubebuilder:default:=15
	Samplewidth *uint16 `json:"sample-width,omitempty"`
	// +kubebuilder:validation:Enum=`AXC`;`DATA`;`FAST`;`SLOW`;`VSD`
	// +kubebuilder:default:="DATA"
	Structureawaremappertype E_DeviceInterfaceRadioRoeMapperdemapperStructureawaremappertype `json:"structure-aware-mapper-type,omitempty"`
	//Structureawaremappertype *string `json:"structure-aware-mapper-type,omitempty"`
}

// DeviceInterfaceRadioRoeMapperdemapperCtrlwordconts struct
type DeviceInterfaceRadioRoeMapperdemapperCtrlwordconts struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	Ctrlwordid *uint8 `json:"ctrl-word-id"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=63
	Cwsize *uint8 `json:"cw-size,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=63
	Cwstart *uint8 `json:"cw-start,omitempty"`
}

// DeviceInterfaceRadioRoeObsaiport struct
type DeviceInterfaceRadioRoeObsaiport struct {
	//RootInterfaceRadioRoeObsaiportLpt
	Lpt *DeviceInterfaceRadioRoeObsaiportLpt `json:"lpt,omitempty"`
	// +kubebuilder:validation:Enum=`obsai-4`;`obsai-8`
	Portspeed E_DeviceInterfaceRadioRoeObsaiportPortspeed `json:"port-speed,omitempty"`
	//Portspeed *string `json:"port-speed,omitempty"`
}

// DeviceInterfaceRadioRoeObsaiportLpt struct
type DeviceInterfaceRadioRoeObsaiportLpt struct {
	// +kubebuilder:validation:Enum=`disable`;`lasershut`;`obsaitxdis`
	Consaction E_DeviceInterfaceRadioRoeObsaiportLptConsaction `json:"cons-action,omitempty"`
	//Consaction *string `json:"cons-action,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=300000
	// +kubebuilder:default:=0
	Extensiontime *uint32 `json:"extension-time,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=300000
	// +kubebuilder:default:=0
	Lasershutdelaytime *uint32 `json:"laser-shut-delay-time,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=300000
	// +kubebuilder:default:=0
	Txdisdelaytime *uint32 `json:"tx-dis-delay-time,omitempty"`
}

// DeviceInterfaceRadioRoePrestimeoffset struct
type DeviceInterfaceRadioRoePrestimeoffset struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=31
	// +kubebuilder:default:=0
	Fractionalnanosecond *uint8 `json:"fractional-nanosecond,omitempty"`
	// kubebuilder:validation:Minimum=5000
	// kubebuilder:validation:Maximum=16777216
	// +kubebuilder:default:=100000
	Integernanosecond *uint32 `json:"integer-nanosecond,omitempty"`
}

// DeviceInterfaceRadioRoeTargetoffset struct
type DeviceInterfaceRadioRoeTargetoffset struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=31
	// +kubebuilder:default:=0
	Fractionalnanosecond *uint8 `json:"fractional-nanosecond,omitempty"`
	// kubebuilder:validation:Minimum=5000
	// kubebuilder:validation:Maximum=16777216
	// +kubebuilder:default:=100000
	Integernanosecond *uint32 `json:"integer-nanosecond,omitempty"`
}

// DeviceInterfaceSflow struct
type DeviceInterfaceSflow struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	Adminstate E_DeviceInterfaceSflowAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
}

// DeviceInterfaceSubinterface struct
type DeviceInterfaceSubinterface struct {
	//RootInterfaceSubinterfaceAcl
	Acl *DeviceInterfaceSubinterfaceAcl `json:"acl,omitempty"`
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_DeviceInterfaceSubinterfaceAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	//RootInterfaceSubinterfaceAnycastgw
	Anycastgw *DeviceInterfaceSubinterfaceAnycastgw `json:"anycast-gw,omitempty"`
	//RootInterfaceSubinterfaceBridgetable
	Bridgetable *DeviceInterfaceSubinterfaceBridgetable `json:"bridge-table,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Description *string `json:"description,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=9999
	Index *uint32 `json:"index"`
	// kubebuilder:validation:Minimum=1280
	// kubebuilder:validation:Maximum=9486
	Ipmtu *uint16 `json:"ip-mtu,omitempty"`
	//RootInterfaceSubinterfaceIpv4
	Ipv4 *DeviceInterfaceSubinterfaceIpv4 `json:"ipv4,omitempty"`
	//RootInterfaceSubinterfaceIpv6
	Ipv6 *DeviceInterfaceSubinterfaceIpv6 `json:"ipv6,omitempty"`
	// kubebuilder:validation:Minimum=1500
	// kubebuilder:validation:Maximum=9500
	L2mtu *uint16 `json:"l2-mtu,omitempty"`
	//RootInterfaceSubinterfaceLocalmirrordestination
	Localmirrordestination *DeviceInterfaceSubinterfaceLocalmirrordestination `json:"local-mirror-destination,omitempty"`
	// kubebuilder:validation:Minimum=1284
	// kubebuilder:validation:Maximum=9496
	Mplsmtu *uint16 `json:"mpls-mtu,omitempty"`
	//RootInterfaceSubinterfaceQos
	Qos *DeviceInterfaceSubinterfaceQos `json:"qos,omitempty"`
	//RootInterfaceSubinterfaceRaguard
	Raguard *DeviceInterfaceSubinterfaceRaguard `json:"ra-guard,omitempty"`
	Type    *string                             `json:"type,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`
	Uuid *string `json:"uuid,omitempty"`
	//RootInterfaceSubinterfaceVlan
	Vlan *DeviceInterfaceSubinterfaceVlan `json:"vlan,omitempty"`
}

// DeviceInterfaceSubinterfaceAcl struct
type DeviceInterfaceSubinterfaceAcl struct {
	//RootInterfaceSubinterfaceAclInput
	Input *DeviceInterfaceSubinterfaceAclInput `json:"input,omitempty"`
	//RootInterfaceSubinterfaceAclOutput
	Output *DeviceInterfaceSubinterfaceAclOutput `json:"output,omitempty"`
}

// DeviceInterfaceSubinterfaceAclInput struct
type DeviceInterfaceSubinterfaceAclInput struct {
	Ipv4filter *string `json:"ipv4-filter,omitempty"`
	Ipv6filter *string `json:"ipv6-filter,omitempty"`
}

// DeviceInterfaceSubinterfaceAclOutput struct
type DeviceInterfaceSubinterfaceAclOutput struct {
	Ipv4filter *string `json:"ipv4-filter,omitempty"`
	Ipv6filter *string `json:"ipv6-filter,omitempty"`
}

// DeviceInterfaceSubinterfaceAnycastgw struct
type DeviceInterfaceSubinterfaceAnycastgw struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}`
	Anycastgwmac *string `json:"anycast-gw-mac,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=1
	Virtualrouterid *uint8 `json:"virtual-router-id,omitempty"`
}

// DeviceInterfaceSubinterfaceBridgetable struct
type DeviceInterfaceSubinterfaceBridgetable struct {
	// +kubebuilder:default:=false
	Discardunknownsrcmac *bool `json:"discard-unknown-src-mac,omitempty"`
	//RootInterfaceSubinterfaceBridgetableMacduplication
	Macduplication *DeviceInterfaceSubinterfaceBridgetableMacduplication `json:"mac-duplication,omitempty"`
	//RootInterfaceSubinterfaceBridgetableMaclearning
	Maclearning *DeviceInterfaceSubinterfaceBridgetableMaclearning `json:"mac-learning,omitempty"`
	//RootInterfaceSubinterfaceBridgetableMaclimit
	Maclimit *DeviceInterfaceSubinterfaceBridgetableMaclimit `json:"mac-limit,omitempty"`
}

// DeviceInterfaceSubinterfaceBridgetableMacduplication struct
type DeviceInterfaceSubinterfaceBridgetableMacduplication struct {
	// +kubebuilder:validation:Enum=`blackhole`;`oper-down`;`stop-learning`;`use-net-instance-action`
	// +kubebuilder:default:="use-net-instance-action"
	Action E_DeviceInterfaceSubinterfaceBridgetableMacduplicationAction `json:"action,omitempty"`
	//Action *string `json:"action,omitempty"`
}

// DeviceInterfaceSubinterfaceBridgetableMaclearning struct
type DeviceInterfaceSubinterfaceBridgetableMaclearning struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_DeviceInterfaceSubinterfaceBridgetableMaclearningAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	//RootInterfaceSubinterfaceBridgetableMaclearningAging
	Aging *DeviceInterfaceSubinterfaceBridgetableMaclearningAging `json:"aging,omitempty"`
}

// DeviceInterfaceSubinterfaceBridgetableMaclearningAging struct
type DeviceInterfaceSubinterfaceBridgetableMaclearningAging struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_DeviceInterfaceSubinterfaceBridgetableMaclearningAgingAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
}

// DeviceInterfaceSubinterfaceBridgetableMaclimit struct
type DeviceInterfaceSubinterfaceBridgetableMaclimit struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=8192
	// +kubebuilder:default:=250
	Maximumentries *int32 `json:"maximum-entries,omitempty"`
	// kubebuilder:validation:Minimum=6
	// kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=95
	Warningthresholdpct *int32 `json:"warning-threshold-pct,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv4 struct
type DeviceInterfaceSubinterfaceIpv4 struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=64
	Address []*DeviceInterfaceSubinterfaceIpv4Address `json:"address,omitempty"`
	// +kubebuilder:default:=false
	Allowdirectedbroadcast *bool `json:"allow-directed-broadcast,omitempty"`
	//RootInterfaceSubinterfaceIpv4Arp
	Arp *DeviceInterfaceSubinterfaceIpv4Arp `json:"arp,omitempty"`
	//RootInterfaceSubinterfaceIpv4Dhcpclient
	Dhcpclient *DeviceInterfaceSubinterfaceIpv4Dhcpclient `json:"dhcp-client,omitempty"`
	//RootInterfaceSubinterfaceIpv4Dhcprelay
	Dhcprelay *DeviceInterfaceSubinterfaceIpv4Dhcprelay `json:"dhcp-relay,omitempty"`
	//RootInterfaceSubinterfaceIpv4Dhcpserver
	Dhcpserver *DeviceInterfaceSubinterfaceIpv4Dhcpserver `json:"dhcp-server,omitempty"`
	//RootInterfaceSubinterfaceIpv4Vrrp
	Vrrp *DeviceInterfaceSubinterfaceIpv4Vrrp `json:"vrrp,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv4Address struct
type DeviceInterfaceSubinterfaceIpv4Address struct {
	Anycastgw *bool `json:"anycast-gw,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))`
	Ipprefix *string `json:"ip-prefix"`
	Primary  *string `json:"primary,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv4Arp struct
type DeviceInterfaceSubinterfaceIpv4Arp struct {
	//RootInterfaceSubinterfaceIpv4ArpDebug
	Debug *DeviceInterfaceSubinterfaceIpv4ArpDebug `json:"debug,omitempty"`
	// +kubebuilder:default:=true
	Duplicateaddressdetection *bool `json:"duplicate-address-detection,omitempty"`
	//RootInterfaceSubinterfaceIpv4ArpEvpn
	Evpn *DeviceInterfaceSubinterfaceIpv4ArpEvpn `json:"evpn,omitempty"`
	//RootInterfaceSubinterfaceIpv4ArpHostroute
	Hostroute *DeviceInterfaceSubinterfaceIpv4ArpHostroute `json:"host-route,omitempty"`
	// +kubebuilder:default:=false
	Learnunsolicited *bool `json:"learn-unsolicited,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Neighbor []*DeviceInterfaceSubinterfaceIpv4ArpNeighbor `json:"neighbor,omitempty"`
	// kubebuilder:validation:Minimum=60
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=14400
	Timeout *uint16 `json:"timeout,omitempty"`
	//RootInterfaceSubinterfaceIpv4ArpVirtualipv4discovery
	Virtualipv4discovery *DeviceInterfaceSubinterfaceIpv4ArpVirtualipv4discovery `json:"virtual-ipv4-discovery,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv4ArpDebug struct
type DeviceInterfaceSubinterfaceIpv4ArpDebug struct {
	// +kubebuilder:validation:Enum=`messages`
	Debug E_DeviceInterfaceSubinterfaceIpv4ArpDebugDebug `json:"debug,omitempty"`
	//Debug *string `json:"debug,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv4ArpEvpn struct
type DeviceInterfaceSubinterfaceIpv4ArpEvpn struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Advertise []*DeviceInterfaceSubinterfaceIpv4ArpEvpnAdvertise `json:"advertise,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv4ArpEvpnAdvertise struct
type DeviceInterfaceSubinterfaceIpv4ArpEvpnAdvertise struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=0
	Admintag *uint32 `json:"admin-tag,omitempty"`
	// +kubebuilder:validation:Enum=`dynamic`;`static`
	Routetype E_DeviceInterfaceSubinterfaceIpv4ArpEvpnAdvertiseRoutetype `json:"route-type,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv4ArpHostroute struct
type DeviceInterfaceSubinterfaceIpv4ArpHostroute struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Populate []*DeviceInterfaceSubinterfaceIpv4ArpHostroutePopulate `json:"populate,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv4ArpHostroutePopulate struct
type DeviceInterfaceSubinterfaceIpv4ArpHostroutePopulate struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=255
	Admintag *uint32 `json:"admin-tag,omitempty"`
	// +kubebuilder:validation:Enum=`dynamic`;`evpn`;`static`
	Routetype E_DeviceInterfaceSubinterfaceIpv4ArpHostroutePopulateRoutetype `json:"route-type,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv4ArpNeighbor struct
type DeviceInterfaceSubinterfaceIpv4ArpNeighbor struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Ipv4address *string `json:"ipv4-address"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}`
	Linklayeraddress *string `json:"link-layer-address"`
}

// DeviceInterfaceSubinterfaceIpv4ArpVirtualipv4discovery struct
type DeviceInterfaceSubinterfaceIpv4ArpVirtualipv4discovery struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=640
	Address []*DeviceInterfaceSubinterfaceIpv4ArpVirtualipv4discoveryAddress `json:"address,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv4ArpVirtualipv4discoveryAddress struct
type DeviceInterfaceSubinterfaceIpv4ArpVirtualipv4discoveryAddress struct {
	//RootInterfaceSubinterfaceIpv4ArpVirtualipv4discoveryAddressAllowedmacs
	Allowedmacs *DeviceInterfaceSubinterfaceIpv4ArpVirtualipv4discoveryAddressAllowedmacs `json:"allowed-macs,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Ipv4address *string `json:"ipv4-address"`
	//RootInterfaceSubinterfaceIpv4ArpVirtualipv4discoveryAddressProbebridgedsubinterfaces
	Probebridgedsubinterfaces *DeviceInterfaceSubinterfaceIpv4ArpVirtualipv4discoveryAddressProbebridgedsubinterfaces `json:"probe-bridged-subinterfaces,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=0
	// +kubebuilder:default:=0
	Probeinterval *uint32 `json:"probe-interval,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv4ArpVirtualipv4discoveryAddressAllowedmacs struct
type DeviceInterfaceSubinterfaceIpv4ArpVirtualipv4discoveryAddressAllowedmacs struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`([0-9a-fA-F][02468aceACE])(:[0-9a-fA-F]{2}){5}|.*[1-9a-fA-F].*`
	Allowedmacs *string `json:"allowed-macs,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv4ArpVirtualipv4discoveryAddressProbebridgedsubinterfaces struct
type DeviceInterfaceSubinterfaceIpv4ArpVirtualipv4discoveryAddressProbebridgedsubinterfaces struct {
	Probebridgedsubinterfaces *string `json:"probe-bridged-subinterfaces,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv4Dhcpclient struct
type DeviceInterfaceSubinterfaceIpv4Dhcpclient struct {
	//RootInterfaceSubinterfaceIpv4DhcpclientTraceoptions
	Traceoptions *DeviceInterfaceSubinterfaceIpv4DhcpclientTraceoptions `json:"trace-options,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv4DhcpclientTraceoptions struct
type DeviceInterfaceSubinterfaceIpv4DhcpclientTraceoptions struct {
	//RootInterfaceSubinterfaceIpv4DhcpclientTraceoptionsTrace
	Trace *DeviceInterfaceSubinterfaceIpv4DhcpclientTraceoptionsTrace `json:"trace,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv4DhcpclientTraceoptionsTrace struct
type DeviceInterfaceSubinterfaceIpv4DhcpclientTraceoptionsTrace struct {
	// +kubebuilder:validation:Enum=`messages`
	Trace E_DeviceInterfaceSubinterfaceIpv4DhcpclientTraceoptionsTraceTrace `json:"trace,omitempty"`
	//Trace *string `json:"trace,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv4Dhcprelay struct
type DeviceInterfaceSubinterfaceIpv4Dhcprelay struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_DeviceInterfaceSubinterfaceIpv4DhcprelayAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Giaddress       *string `json:"gi-address,omitempty"`
	Networkinstance *string `json:"network-instance,omitempty"`
	//RootInterfaceSubinterfaceIpv4DhcprelayOption
	Option *DeviceInterfaceSubinterfaceIpv4DhcprelayOption `json:"option,omitempty"`
	//RootInterfaceSubinterfaceIpv4DhcprelayServer
	Server *DeviceInterfaceSubinterfaceIpv4DhcprelayServer `json:"server,omitempty"`
	//RootInterfaceSubinterfaceIpv4DhcprelayTraceoptions
	Traceoptions *DeviceInterfaceSubinterfaceIpv4DhcprelayTraceoptions `json:"trace-options,omitempty"`
	// +kubebuilder:default:=false
	Usegiaddrassrcipaddr *bool `json:"use-gi-addr-as-src-ip-addr,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv4DhcprelayOption struct
type DeviceInterfaceSubinterfaceIpv4DhcprelayOption struct {
	// +kubebuilder:validation:Enum=`circuit-id`;`remote-id`
	Option E_DeviceInterfaceSubinterfaceIpv4DhcprelayOptionOption `json:"option,omitempty"`
	//Option *string `json:"option,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv4DhcprelayServer struct
type DeviceInterfaceSubinterfaceIpv4DhcprelayServer struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((([a-zA-Z0-9_]([a-zA-Z0-9\-_]){0,61})?[a-zA-Z0-9]\.)*([a-zA-Z0-9_]([a-zA-Z0-9\-_]){0,61})?[a-zA-Z0-9]\.?)|\.`
	Server *string `json:"server,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv4DhcprelayTraceoptions struct
type DeviceInterfaceSubinterfaceIpv4DhcprelayTraceoptions struct {
	//RootInterfaceSubinterfaceIpv4DhcprelayTraceoptionsTrace
	Trace *DeviceInterfaceSubinterfaceIpv4DhcprelayTraceoptionsTrace `json:"trace,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv4DhcprelayTraceoptionsTrace struct
type DeviceInterfaceSubinterfaceIpv4DhcprelayTraceoptionsTrace struct {
	// +kubebuilder:validation:Enum=`messages`
	Trace E_DeviceInterfaceSubinterfaceIpv4DhcprelayTraceoptionsTraceTrace `json:"trace,omitempty"`
	//Trace *string `json:"trace,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv4Dhcpserver struct
type DeviceInterfaceSubinterfaceIpv4Dhcpserver struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_DeviceInterfaceSubinterfaceIpv4DhcpserverAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv4Vrrp struct
type DeviceInterfaceSubinterfaceIpv4Vrrp struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Vrrpgroup []*DeviceInterfaceSubinterfaceIpv4VrrpVrrpgroup `json:"vrrp-group,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv4VrrpVrrpgroup struct
type DeviceInterfaceSubinterfaceIpv4VrrpVrrpgroup struct {
	Acceptmode *bool `json:"accept-mode,omitempty"`
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_DeviceInterfaceSubinterfaceIpv4VrrpVrrpgroupAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=1000
	Advertiseinterval *uint16 `json:"advertise-interval,omitempty"`
	//RootInterfaceSubinterfaceIpv4VrrpVrrpgroupAuthentication
	Authentication *DeviceInterfaceSubinterfaceIpv4VrrpVrrpgroupAuthentication `json:"authentication,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	Initdelay *uint16 `json:"init-delay,omitempty"`
	//RootInterfaceSubinterfaceIpv4VrrpVrrpgroupInterfacetracking
	Interfacetracking *DeviceInterfaceSubinterfaceIpv4VrrpVrrpgroupInterfacetracking `json:"interface-tracking,omitempty"`
	// +kubebuilder:default:=false
	Masterinheritinterval *bool `json:"master-inherit-interval,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	Operinterval *uint16 `json:"oper-interval,omitempty"`
	Preempt      *bool   `json:"preempt,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	Preemptdelay *uint16 `json:"preempt-delay,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=100
	Priority *uint8 `json:"priority,omitempty"`
	//RootInterfaceSubinterfaceIpv4VrrpVrrpgroupStatistics
	Statistics *DeviceInterfaceSubinterfaceIpv4VrrpVrrpgroupStatistics `json:"statistics,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=2
	Version *uint8 `json:"version,omitempty"`
	//RootInterfaceSubinterfaceIpv4VrrpVrrpgroupVirtualaddress
	Virtualaddress *DeviceInterfaceSubinterfaceIpv4VrrpVrrpgroupVirtualaddress `json:"virtual-address,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	Virtualrouterid *uint8 `json:"virtual-router-id"`
}

// DeviceInterfaceSubinterfaceIpv4VrrpVrrpgroupAuthentication struct
type DeviceInterfaceSubinterfaceIpv4VrrpVrrpgroupAuthentication struct {
	Keychain *string `json:"keychain,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv4VrrpVrrpgroupInterfacetracking struct
type DeviceInterfaceSubinterfaceIpv4VrrpVrrpgroupInterfacetracking struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Trackinterface []*DeviceInterfaceSubinterfaceIpv4VrrpVrrpgroupInterfacetrackingTrackinterface `json:"track-interface,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv4VrrpVrrpgroupInterfacetrackingTrackinterface struct
type DeviceInterfaceSubinterfaceIpv4VrrpVrrpgroupInterfacetrackingTrackinterface struct {
	Interface *string `json:"interface"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	Prioritydecrement *uint8 `json:"priority-decrement,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv4VrrpVrrpgroupStatistics struct
type DeviceInterfaceSubinterfaceIpv4VrrpVrrpgroupStatistics struct {
}

// DeviceInterfaceSubinterfaceIpv4VrrpVrrpgroupVirtualaddress struct
type DeviceInterfaceSubinterfaceIpv4VrrpVrrpgroupVirtualaddress struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Virtualaddress *string `json:"virtual-address,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv6 struct
type DeviceInterfaceSubinterfaceIpv6 struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=18
	Address []*DeviceInterfaceSubinterfaceIpv6Address `json:"address,omitempty"`
	//RootInterfaceSubinterfaceIpv6Dhcpclient
	Dhcpclient *DeviceInterfaceSubinterfaceIpv6Dhcpclient `json:"dhcp-client,omitempty"`
	//RootInterfaceSubinterfaceIpv6Dhcprelay
	Dhcprelay *DeviceInterfaceSubinterfaceIpv6Dhcprelay `json:"dhcp-relay,omitempty"`
	//RootInterfaceSubinterfaceIpv6Dhcpv6server
	Dhcpv6server *DeviceInterfaceSubinterfaceIpv6Dhcpv6server `json:"dhcpv6-server,omitempty"`
	//RootInterfaceSubinterfaceIpv6Neighbordiscovery
	Neighbordiscovery *DeviceInterfaceSubinterfaceIpv6Neighbordiscovery `json:"neighbor-discovery,omitempty"`
	//RootInterfaceSubinterfaceIpv6Routeradvertisement
	Routeradvertisement *DeviceInterfaceSubinterfaceIpv6Routeradvertisement `json:"router-advertisement,omitempty"`
	//RootInterfaceSubinterfaceIpv6Vrrp
	Vrrp *DeviceInterfaceSubinterfaceIpv6Vrrp `json:"vrrp,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv6Address struct
type DeviceInterfaceSubinterfaceIpv6Address struct {
	Anycastgw *bool `json:"anycast-gw,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	Ipprefix *string `json:"ip-prefix"`
	Primary  *string `json:"primary,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv6Dhcpclient struct
type DeviceInterfaceSubinterfaceIpv6Dhcpclient struct {
	//RootInterfaceSubinterfaceIpv6DhcpclientTraceoptions
	Traceoptions *DeviceInterfaceSubinterfaceIpv6DhcpclientTraceoptions `json:"trace-options,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv6DhcpclientTraceoptions struct
type DeviceInterfaceSubinterfaceIpv6DhcpclientTraceoptions struct {
	//RootInterfaceSubinterfaceIpv6DhcpclientTraceoptionsTrace
	Trace *DeviceInterfaceSubinterfaceIpv6DhcpclientTraceoptionsTrace `json:"trace,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv6DhcpclientTraceoptionsTrace struct
type DeviceInterfaceSubinterfaceIpv6DhcpclientTraceoptionsTrace struct {
	// +kubebuilder:validation:Enum=`messages`
	Trace E_DeviceInterfaceSubinterfaceIpv6DhcpclientTraceoptionsTraceTrace `json:"trace,omitempty"`
	//Trace *string `json:"trace,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv6Dhcprelay struct
type DeviceInterfaceSubinterfaceIpv6Dhcprelay struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_DeviceInterfaceSubinterfaceIpv6DhcprelayAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	Networkinstance *string `json:"network-instance,omitempty"`
	//RootInterfaceSubinterfaceIpv6DhcprelayOption
	Option *DeviceInterfaceSubinterfaceIpv6DhcprelayOption `json:"option,omitempty"`
	//RootInterfaceSubinterfaceIpv6DhcprelayServer
	Server *DeviceInterfaceSubinterfaceIpv6DhcprelayServer `json:"server,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Sourceaddress *string `json:"source-address,omitempty"`
	//RootInterfaceSubinterfaceIpv6DhcprelayTraceoptions
	Traceoptions *DeviceInterfaceSubinterfaceIpv6DhcprelayTraceoptions `json:"trace-options,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv6DhcprelayOption struct
type DeviceInterfaceSubinterfaceIpv6DhcprelayOption struct {
	// +kubebuilder:validation:Enum=`interface-id`;`remote-id`
	Option E_DeviceInterfaceSubinterfaceIpv6DhcprelayOptionOption `json:"option,omitempty"`
	//Option *string `json:"option,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv6DhcprelayServer struct
type DeviceInterfaceSubinterfaceIpv6DhcprelayServer struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))|((([a-zA-Z0-9_]([a-zA-Z0-9\-_]){0,61})?[a-zA-Z0-9]\.)*([a-zA-Z0-9_]([a-zA-Z0-9\-_]){0,61})?[a-zA-Z0-9]\.?)|\.`
	Server *string `json:"server,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv6DhcprelayTraceoptions struct
type DeviceInterfaceSubinterfaceIpv6DhcprelayTraceoptions struct {
	//RootInterfaceSubinterfaceIpv6DhcprelayTraceoptionsTrace
	Trace *DeviceInterfaceSubinterfaceIpv6DhcprelayTraceoptionsTrace `json:"trace,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv6DhcprelayTraceoptionsTrace struct
type DeviceInterfaceSubinterfaceIpv6DhcprelayTraceoptionsTrace struct {
	// +kubebuilder:validation:Enum=`messages`
	Trace E_DeviceInterfaceSubinterfaceIpv6DhcprelayTraceoptionsTraceTrace `json:"trace,omitempty"`
	//Trace *string `json:"trace,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv6Dhcpv6server struct
type DeviceInterfaceSubinterfaceIpv6Dhcpv6server struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_DeviceInterfaceSubinterfaceIpv6Dhcpv6serverAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv6Neighbordiscovery struct
type DeviceInterfaceSubinterfaceIpv6Neighbordiscovery struct {
	//RootInterfaceSubinterfaceIpv6NeighbordiscoveryDebug
	Debug *DeviceInterfaceSubinterfaceIpv6NeighbordiscoveryDebug `json:"debug,omitempty"`
	// +kubebuilder:default:=true
	Duplicateaddressdetection *bool `json:"duplicate-address-detection,omitempty"`
	//RootInterfaceSubinterfaceIpv6NeighbordiscoveryEvpn
	Evpn *DeviceInterfaceSubinterfaceIpv6NeighbordiscoveryEvpn `json:"evpn,omitempty"`
	//RootInterfaceSubinterfaceIpv6NeighbordiscoveryHostroute
	Hostroute *DeviceInterfaceSubinterfaceIpv6NeighbordiscoveryHostroute `json:"host-route,omitempty"`
	// +kubebuilder:validation:Enum=`both`;`global`;`link-local`;`none`
	// +kubebuilder:default:="none"
	Learnunsolicited E_DeviceInterfaceSubinterfaceIpv6NeighbordiscoveryLearnunsolicited `json:"learn-unsolicited,omitempty"`
	//Learnunsolicited *string `json:"learn-unsolicited,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Neighbor []*DeviceInterfaceSubinterfaceIpv6NeighbordiscoveryNeighbor `json:"neighbor,omitempty"`
	// kubebuilder:validation:Minimum=30
	// kubebuilder:validation:Maximum=3600
	// +kubebuilder:default:=30
	Reachabletime *uint32 `json:"reachable-time,omitempty"`
	// kubebuilder:validation:Minimum=60
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=14400
	Staletime *uint32 `json:"stale-time,omitempty"`
	//RootInterfaceSubinterfaceIpv6NeighbordiscoveryVirtualipv6discovery
	Virtualipv6discovery *DeviceInterfaceSubinterfaceIpv6NeighbordiscoveryVirtualipv6discovery `json:"virtual-ipv6-discovery,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv6NeighbordiscoveryDebug struct
type DeviceInterfaceSubinterfaceIpv6NeighbordiscoveryDebug struct {
	// +kubebuilder:validation:Enum=`messages`
	Debug E_DeviceInterfaceSubinterfaceIpv6NeighbordiscoveryDebugDebug `json:"debug,omitempty"`
	//Debug *string `json:"debug,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv6NeighbordiscoveryEvpn struct
type DeviceInterfaceSubinterfaceIpv6NeighbordiscoveryEvpn struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Advertise []*DeviceInterfaceSubinterfaceIpv6NeighbordiscoveryEvpnAdvertise `json:"advertise,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv6NeighbordiscoveryEvpnAdvertise struct
type DeviceInterfaceSubinterfaceIpv6NeighbordiscoveryEvpnAdvertise struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=0
	Admintag *uint32 `json:"admin-tag,omitempty"`
	// +kubebuilder:validation:Enum=`dynamic`;`static`
	Routetype E_DeviceInterfaceSubinterfaceIpv6NeighbordiscoveryEvpnAdvertiseRoutetype `json:"route-type,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv6NeighbordiscoveryHostroute struct
type DeviceInterfaceSubinterfaceIpv6NeighbordiscoveryHostroute struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Populate []*DeviceInterfaceSubinterfaceIpv6NeighbordiscoveryHostroutePopulate `json:"populate,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv6NeighbordiscoveryHostroutePopulate struct
type DeviceInterfaceSubinterfaceIpv6NeighbordiscoveryHostroutePopulate struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=255
	Admintag *uint32 `json:"admin-tag,omitempty"`
	// +kubebuilder:validation:Enum=`dynamic`;`evpn`;`static`
	Routetype E_DeviceInterfaceSubinterfaceIpv6NeighbordiscoveryHostroutePopulateRoutetype `json:"route-type,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv6NeighbordiscoveryNeighbor struct
type DeviceInterfaceSubinterfaceIpv6NeighbordiscoveryNeighbor struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Ipv6address *string `json:"ipv6-address"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}`
	Linklayeraddress *string `json:"link-layer-address"`
}

// DeviceInterfaceSubinterfaceIpv6NeighbordiscoveryVirtualipv6discovery struct
type DeviceInterfaceSubinterfaceIpv6NeighbordiscoveryVirtualipv6discovery struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=640
	Address []*DeviceInterfaceSubinterfaceIpv6NeighbordiscoveryVirtualipv6discoveryAddress `json:"address,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv6NeighbordiscoveryVirtualipv6discoveryAddress struct
type DeviceInterfaceSubinterfaceIpv6NeighbordiscoveryVirtualipv6discoveryAddress struct {
	//RootInterfaceSubinterfaceIpv6NeighbordiscoveryVirtualipv6discoveryAddressAllowedmacs
	Allowedmacs *DeviceInterfaceSubinterfaceIpv6NeighbordiscoveryVirtualipv6discoveryAddressAllowedmacs `json:"allowed-macs,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Ipv6address *string `json:"ipv6-address"`
	//RootInterfaceSubinterfaceIpv6NeighbordiscoveryVirtualipv6discoveryAddressProbebridgedsubinterfaces
	Probebridgedsubinterfaces *DeviceInterfaceSubinterfaceIpv6NeighbordiscoveryVirtualipv6discoveryAddressProbebridgedsubinterfaces `json:"probe-bridged-subinterfaces,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=0
	// +kubebuilder:default:=0
	Probeinterval *uint32 `json:"probe-interval,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv6NeighbordiscoveryVirtualipv6discoveryAddressAllowedmacs struct
type DeviceInterfaceSubinterfaceIpv6NeighbordiscoveryVirtualipv6discoveryAddressAllowedmacs struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`([0-9a-fA-F][02468aceACE])(:[0-9a-fA-F]{2}){5}|.*[1-9a-fA-F].*`
	Allowedmacs *string `json:"allowed-macs,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv6NeighbordiscoveryVirtualipv6discoveryAddressProbebridgedsubinterfaces struct
type DeviceInterfaceSubinterfaceIpv6NeighbordiscoveryVirtualipv6discoveryAddressProbebridgedsubinterfaces struct {
	Probebridgedsubinterfaces *string `json:"probe-bridged-subinterfaces,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv6Routeradvertisement struct
type DeviceInterfaceSubinterfaceIpv6Routeradvertisement struct {
	//RootInterfaceSubinterfaceIpv6RouteradvertisementDebug
	Debug *DeviceInterfaceSubinterfaceIpv6RouteradvertisementDebug `json:"debug,omitempty"`
	//RootInterfaceSubinterfaceIpv6RouteradvertisementRouterrole
	Routerrole *DeviceInterfaceSubinterfaceIpv6RouteradvertisementRouterrole `json:"router-role,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv6RouteradvertisementDebug struct
type DeviceInterfaceSubinterfaceIpv6RouteradvertisementDebug struct {
	// +kubebuilder:validation:Enum=`messages`
	Debug E_DeviceInterfaceSubinterfaceIpv6RouteradvertisementDebugDebug `json:"debug,omitempty"`
	//Debug *string `json:"debug,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv6RouteradvertisementRouterrole struct
type DeviceInterfaceSubinterfaceIpv6RouteradvertisementRouterrole struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_DeviceInterfaceSubinterfaceIpv6RouteradvertisementRouterroleAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=64
	Currenthoplimit *uint8 `json:"current-hop-limit,omitempty"`
	// kubebuilder:validation:Minimum=1280
	// kubebuilder:validation:Maximum=9486
	Ipmtu *uint16 `json:"ip-mtu,omitempty"`
	// +kubebuilder:default:=false
	Managedconfigurationflag *bool `json:"managed-configuration-flag,omitempty"`
	// kubebuilder:validation:Minimum=4
	// kubebuilder:validation:Maximum=1800
	// +kubebuilder:default:=600
	Maxadvertisementinterval *uint16 `json:"max-advertisement-interval,omitempty"`
	// kubebuilder:validation:Minimum=3
	// kubebuilder:validation:Maximum=1350
	// +kubebuilder:default:=200
	Minadvertisementinterval *uint16 `json:"min-advertisement-interval,omitempty"`
	// +kubebuilder:default:=false
	Otherconfigurationflag *bool `json:"other-configuration-flag,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=16
	Prefix []*DeviceInterfaceSubinterfaceIpv6RouteradvertisementRouterrolePrefix `json:"prefix,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=3600000
	// +kubebuilder:default:=0
	Reachabletime *uint32 `json:"reachable-time,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=1800000
	// +kubebuilder:default:=0
	Retransmittime *uint32 `json:"retransmit-time,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=9000
	// +kubebuilder:default:=1800
	Routerlifetime *uint16 `json:"router-lifetime,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv6RouteradvertisementRouterrolePrefix struct
type DeviceInterfaceSubinterfaceIpv6RouteradvertisementRouterrolePrefix struct {
	// +kubebuilder:default:=true
	Autonomousflag *bool `json:"autonomous-flag,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	Ipv6prefix *string `json:"ipv6-prefix"`
	// +kubebuilder:default:=true
	Onlinkflag *bool `json:"on-link-flag,omitempty"`
	// +kubebuilder:default:=604800
	Preferredlifetime *uint32 `json:"preferred-lifetime,omitempty"`
	// +kubebuilder:default:=2592000
	Validlifetime *uint32 `json:"valid-lifetime,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv6Vrrp struct
type DeviceInterfaceSubinterfaceIpv6Vrrp struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Vrrpgroup []*DeviceInterfaceSubinterfaceIpv6VrrpVrrpgroup `json:"vrrp-group,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv6VrrpVrrpgroup struct
type DeviceInterfaceSubinterfaceIpv6VrrpVrrpgroup struct {
	Acceptmode *bool `json:"accept-mode,omitempty"`
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_DeviceInterfaceSubinterfaceIpv6VrrpVrrpgroupAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=1000
	Advertiseinterval *uint16 `json:"advertise-interval,omitempty"`
	//RootInterfaceSubinterfaceIpv6VrrpVrrpgroupAuthentication
	Authentication *DeviceInterfaceSubinterfaceIpv6VrrpVrrpgroupAuthentication `json:"authentication,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	Initdelay *uint16 `json:"init-delay,omitempty"`
	//RootInterfaceSubinterfaceIpv6VrrpVrrpgroupInterfacetracking
	Interfacetracking *DeviceInterfaceSubinterfaceIpv6VrrpVrrpgroupInterfacetracking `json:"interface-tracking,omitempty"`
	// +kubebuilder:default:=false
	Masterinheritinterval *bool `json:"master-inherit-interval,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	Operinterval *uint16 `json:"oper-interval,omitempty"`
	Preempt      *bool   `json:"preempt,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	Preemptdelay *uint16 `json:"preempt-delay,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=100
	Priority *uint8 `json:"priority,omitempty"`
	//RootInterfaceSubinterfaceIpv6VrrpVrrpgroupStatistics
	Statistics *DeviceInterfaceSubinterfaceIpv6VrrpVrrpgroupStatistics `json:"statistics,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=3
	Version *uint8 `json:"version,omitempty"`
	//RootInterfaceSubinterfaceIpv6VrrpVrrpgroupVirtualaddress
	Virtualaddress *DeviceInterfaceSubinterfaceIpv6VrrpVrrpgroupVirtualaddress `json:"virtual-address,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	Virtualrouterid *uint8 `json:"virtual-router-id"`
}

// DeviceInterfaceSubinterfaceIpv6VrrpVrrpgroupAuthentication struct
type DeviceInterfaceSubinterfaceIpv6VrrpVrrpgroupAuthentication struct {
	Keychain *string `json:"keychain,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv6VrrpVrrpgroupInterfacetracking struct
type DeviceInterfaceSubinterfaceIpv6VrrpVrrpgroupInterfacetracking struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Trackinterface []*DeviceInterfaceSubinterfaceIpv6VrrpVrrpgroupInterfacetrackingTrackinterface `json:"track-interface,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv6VrrpVrrpgroupInterfacetrackingTrackinterface struct
type DeviceInterfaceSubinterfaceIpv6VrrpVrrpgroupInterfacetrackingTrackinterface struct {
	Interface *string `json:"interface"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	Prioritydecrement *uint8 `json:"priority-decrement,omitempty"`
}

// DeviceInterfaceSubinterfaceIpv6VrrpVrrpgroupStatistics struct
type DeviceInterfaceSubinterfaceIpv6VrrpVrrpgroupStatistics struct {
}

// DeviceInterfaceSubinterfaceIpv6VrrpVrrpgroupVirtualaddress struct
type DeviceInterfaceSubinterfaceIpv6VrrpVrrpgroupVirtualaddress struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Virtualaddress *string `json:"virtual-address,omitempty"`
}

// DeviceInterfaceSubinterfaceLocalmirrordestination struct
type DeviceInterfaceSubinterfaceLocalmirrordestination struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_DeviceInterfaceSubinterfaceLocalmirrordestinationAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
}

// DeviceInterfaceSubinterfaceQos struct
type DeviceInterfaceSubinterfaceQos struct {
	//RootInterfaceSubinterfaceQosInput
	Input *DeviceInterfaceSubinterfaceQosInput `json:"input,omitempty"`
	//RootInterfaceSubinterfaceQosOutput
	Output *DeviceInterfaceSubinterfaceQosOutput `json:"output,omitempty"`
}

// DeviceInterfaceSubinterfaceQosInput struct
type DeviceInterfaceSubinterfaceQosInput struct {
	//RootInterfaceSubinterfaceQosInputClassifiers
	Classifiers *DeviceInterfaceSubinterfaceQosInputClassifiers `json:"classifiers,omitempty"`
}

// DeviceInterfaceSubinterfaceQosInputClassifiers struct
type DeviceInterfaceSubinterfaceQosInputClassifiers struct {
	Dscp             *string `json:"dscp,omitempty"`
	Ipv4dscp         *string `json:"ipv4-dscp,omitempty"`
	Ipv6dscp         *string `json:"ipv6-dscp,omitempty"`
	Mplstrafficclass *string `json:"mpls-traffic-class,omitempty"`
}

// DeviceInterfaceSubinterfaceQosOutput struct
type DeviceInterfaceSubinterfaceQosOutput struct {
	//RootInterfaceSubinterfaceQosOutputRewriterules
	Rewriterules *DeviceInterfaceSubinterfaceQosOutputRewriterules `json:"rewrite-rules,omitempty"`
}

// DeviceInterfaceSubinterfaceQosOutputRewriterules struct
type DeviceInterfaceSubinterfaceQosOutputRewriterules struct {
	Dscp             *string `json:"dscp,omitempty"`
	Ipv4dscp         *string `json:"ipv4-dscp,omitempty"`
	Ipv6dscp         *string `json:"ipv6-dscp,omitempty"`
	Mplstrafficclass *string `json:"mpls-traffic-class,omitempty"`
}

// DeviceInterfaceSubinterfaceRaguard struct
type DeviceInterfaceSubinterfaceRaguard struct {
	Policy *string `json:"policy,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Vlanlist []*DeviceInterfaceSubinterfaceRaguardVlanlist `json:"vlan-list,omitempty"`
}

// DeviceInterfaceSubinterfaceRaguardVlanlist struct
type DeviceInterfaceSubinterfaceRaguardVlanlist struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4095
	Vlanid *uint16 `json:"vlan-id"`
}

// DeviceInterfaceSubinterfaceVlan struct
type DeviceInterfaceSubinterfaceVlan struct {
	//RootInterfaceSubinterfaceVlanEncap
	Encap *DeviceInterfaceSubinterfaceVlanEncap `json:"encap,omitempty"`
	//RootInterfaceSubinterfaceVlanVlandiscovery
	Vlandiscovery *DeviceInterfaceSubinterfaceVlanVlandiscovery `json:"vlan-discovery,omitempty"`
}

// DeviceInterfaceSubinterfaceVlanEncap struct
type DeviceInterfaceSubinterfaceVlanEncap struct {
	//RootInterfaceSubinterfaceVlanEncapSingletagged
	Singletagged *DeviceInterfaceSubinterfaceVlanEncapSingletagged `json:"single-tagged,omitempty"`
	//RootInterfaceSubinterfaceVlanEncapUntagged
	Untagged *DeviceInterfaceSubinterfaceVlanEncapUntagged `json:"untagged,omitempty"`
}

// DeviceInterfaceSubinterfaceVlanEncapSingletagged struct
type DeviceInterfaceSubinterfaceVlanEncapSingletagged struct {
	Vlanid *string `json:"vlan-id,omitempty"`
}

// DeviceInterfaceSubinterfaceVlanEncapUntagged struct
type DeviceInterfaceSubinterfaceVlanEncapUntagged struct {
}

// DeviceInterfaceSubinterfaceVlanVlandiscovery struct
type DeviceInterfaceSubinterfaceVlanVlandiscovery struct {
	// +kubebuilder:validation:Enum=`IPv4`;`IPv4v6`;`IPv6`
	// +kubebuilder:default:="IPv4v6"
	Type E_DeviceInterfaceSubinterfaceVlanVlandiscoveryType `json:"type,omitempty"`
	//Type *string `json:"type,omitempty"`
}

// DeviceInterfaceTransceiver struct
type DeviceInterfaceTransceiver struct {
	Ddmevents *bool `json:"ddm-events,omitempty"`
	// +kubebuilder:validation:Enum=`base-r`;`disabled`;`rs-108`;`rs-528`;`rs-544`
	Forwarderrorcorrection E_DeviceInterfaceTransceiverForwarderrorcorrection `json:"forward-error-correction,omitempty"`
	//Forwarderrorcorrection *string `json:"forward-error-correction,omitempty"`
	Txlaser *bool `json:"tx-laser,omitempty"`
}

// DeviceInterfaceVhost struct
type DeviceInterfaceVhost struct {
	// +kubebuilder:validation:Enum=`client`;`server`
	// +kubebuilder:default:="client"
	Vhostsocketmode E_DeviceInterfaceVhostVhostsocketmode `json:"vhost-socket-mode,omitempty"`
	//Vhostsocketmode *string `json:"vhost-socket-mode,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(/[0-9A-Za-z_\-\.]+)+`
	Vhostsocketpath *string `json:"vhost-socket-path"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=1024
	Vhostsocketqueues *uint16 `json:"vhost-socket-queues,omitempty"`
}

// DeviceNetworkinstance struct
type DeviceNetworkinstance struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_DeviceNetworkinstanceAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	//RootNetworkinstanceAggregateroutes
	Aggregateroutes *DeviceNetworkinstanceAggregateroutes `json:"aggregate-routes,omitempty"`
	//RootNetworkinstanceBridgetable
	Bridgetable *DeviceNetworkinstanceBridgetable `json:"bridge-table,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Description *string `json:"description,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Interface []*DeviceNetworkinstanceInterface `json:"interface,omitempty"`
	//RootNetworkinstanceIpforwarding
	Ipforwarding *DeviceNetworkinstanceIpforwarding `json:"ip-forwarding,omitempty"`
	//RootNetworkinstanceIploadbalancing
	Iploadbalancing *DeviceNetworkinstanceIploadbalancing `json:"ip-load-balancing,omitempty"`
	//RootNetworkinstanceMpls
	Mpls *DeviceNetworkinstanceMpls `json:"mpls,omitempty"`
	//RootNetworkinstanceMplsforwarding
	Mplsforwarding *DeviceNetworkinstanceMplsforwarding `json:"mpls-forwarding,omitempty"`
	//RootNetworkinstanceMtu
	Mtu *DeviceNetworkinstanceMtu `json:"mtu,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
	//RootNetworkinstanceNexthopgroups
	Nexthopgroups *DeviceNetworkinstanceNexthopgroups `json:"next-hop-groups,omitempty"`
	//RootNetworkinstancePolicyforwarding
	Policyforwarding *DeviceNetworkinstancePolicyforwarding `json:"policy-forwarding,omitempty"`
	//RootNetworkinstanceProtocols
	Protocols *DeviceNetworkinstanceProtocols `json:"protocols,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Routerid *string `json:"router-id,omitempty"`
	//RootNetworkinstanceSegmentrouting
	Segmentrouting *DeviceNetworkinstanceSegmentrouting `json:"segment-routing,omitempty"`
	//RootNetworkinstanceStaticroutes
	Staticroutes *DeviceNetworkinstanceStaticroutes `json:"static-routes,omitempty"`
	//RootNetworkinstanceTrafficengineering
	Trafficengineering *DeviceNetworkinstanceTrafficengineering `json:"traffic-engineering,omitempty"`
	// +kubebuilder:default:="default"
	Type *string `json:"type,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1
	Vxlaninterface []*DeviceNetworkinstanceVxlaninterface `json:"vxlan-interface,omitempty"`
}

// DeviceNetworkinstanceAggregateroutes struct
type DeviceNetworkinstanceAggregateroutes struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=16384
	Route []*DeviceNetworkinstanceAggregateroutesRoute `json:"route,omitempty"`
}

// DeviceNetworkinstanceAggregateroutesRoute struct
type DeviceNetworkinstanceAggregateroutesRoute struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_DeviceNetworkinstanceAggregateroutesRouteAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	//RootNetworkinstanceAggregateroutesRouteAggregator
	Aggregator *DeviceNetworkinstanceAggregateroutesRouteAggregator `json:"aggregator,omitempty"`
	//RootNetworkinstanceAggregateroutesRouteCommunities
	Communities  *DeviceNetworkinstanceAggregateroutesRouteCommunities `json:"communities,omitempty"`
	Generateicmp *bool                                                 `json:"generate-icmp,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	Prefix *string `json:"prefix"`
	// +kubebuilder:default:=false
	Summaryonly *bool `json:"summary-only,omitempty"`
}

// DeviceNetworkinstanceAggregateroutesRouteAggregator struct
type DeviceNetworkinstanceAggregateroutesRouteAggregator struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Address *string `json:"address,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	Asnumber *uint32 `json:"as-number,omitempty"`
}

// DeviceNetworkinstanceAggregateroutesRouteCommunities struct
type DeviceNetworkinstanceAggregateroutesRouteCommunities struct {
	//RootNetworkinstanceAggregateroutesRouteCommunitiesAdd
	Add *DeviceNetworkinstanceAggregateroutesRouteCommunitiesAdd `json:"add,omitempty"`
}

// DeviceNetworkinstanceAggregateroutesRouteCommunitiesAdd struct
type DeviceNetworkinstanceAggregateroutesRouteCommunitiesAdd struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])|.*:.*|([1-9][0-9]{0,9}):([1-9][0-9]{0,9}):([1-9][0-9]{0,9})|.*:.*:.*`
	Add *string `json:"add,omitempty"`
}

// DeviceNetworkinstanceBridgetable struct
type DeviceNetworkinstanceBridgetable struct {
	// +kubebuilder:default:=false
	Discardunknowndestmac *bool `json:"discard-unknown-dest-mac,omitempty"`
	//RootNetworkinstanceBridgetableMacduplication
	Macduplication *DeviceNetworkinstanceBridgetableMacduplication `json:"mac-duplication,omitempty"`
	//RootNetworkinstanceBridgetableMaclearning
	Maclearning *DeviceNetworkinstanceBridgetableMaclearning `json:"mac-learning,omitempty"`
	//RootNetworkinstanceBridgetableMaclimit
	Maclimit *DeviceNetworkinstanceBridgetableMaclimit `json:"mac-limit,omitempty"`
	// +kubebuilder:default:=false
	Protectanycastgwmac *bool `json:"protect-anycast-gw-mac,omitempty"`
	//RootNetworkinstanceBridgetableProxyarp
	Proxyarp *DeviceNetworkinstanceBridgetableProxyarp `json:"proxy-arp,omitempty"`
	//RootNetworkinstanceBridgetableStaticmac
	Staticmac *DeviceNetworkinstanceBridgetableStaticmac `json:"static-mac,omitempty"`
}

// DeviceNetworkinstanceBridgetableMacduplication struct
type DeviceNetworkinstanceBridgetableMacduplication struct {
	// +kubebuilder:validation:Enum=`blackhole`;`oper-down`;`stop-learning`
	// +kubebuilder:default:="stop-learning"
	Action E_DeviceNetworkinstanceBridgetableMacduplicationAction `json:"action,omitempty"`
	//Action *string `json:"action,omitempty"`
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_DeviceNetworkinstanceBridgetableMacduplicationAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// +kubebuilder:default:=9
	Holddowntime *uint32 `json:"hold-down-time,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=15
	// +kubebuilder:default:=3
	Monitoringwindow *uint32 `json:"monitoring-window,omitempty"`
	// kubebuilder:validation:Minimum=3
	// kubebuilder:validation:Maximum=10
	// +kubebuilder:default:=5
	Nummoves *uint32 `json:"num-moves,omitempty"`
}

// DeviceNetworkinstanceBridgetableMaclearning struct
type DeviceNetworkinstanceBridgetableMaclearning struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_DeviceNetworkinstanceBridgetableMaclearningAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	//RootNetworkinstanceBridgetableMaclearningAging
	Aging *DeviceNetworkinstanceBridgetableMaclearningAging `json:"aging,omitempty"`
}

// DeviceNetworkinstanceBridgetableMaclearningAging struct
type DeviceNetworkinstanceBridgetableMaclearningAging struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_DeviceNetworkinstanceBridgetableMaclearningAgingAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=60
	// kubebuilder:validation:Maximum=86400
	// +kubebuilder:default:=300
	Agetime *int32 `json:"age-time,omitempty"`
}

// DeviceNetworkinstanceBridgetableMaclimit struct
type DeviceNetworkinstanceBridgetableMaclimit struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=8192
	// +kubebuilder:default:=250
	Maximumentries *int32 `json:"maximum-entries,omitempty"`
	// kubebuilder:validation:Minimum=6
	// kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=95
	Warningthresholdpct *int32 `json:"warning-threshold-pct,omitempty"`
}

// DeviceNetworkinstanceBridgetableProxyarp struct
type DeviceNetworkinstanceBridgetableProxyarp struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_DeviceNetworkinstanceBridgetableProxyarpAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	Agetime *uint32 `json:"age-time,omitempty"`
	//RootNetworkinstanceBridgetableProxyarpDuplicatedetect
	Duplicatedetect *DeviceNetworkinstanceBridgetableProxyarpDuplicatedetect `json:"duplicate-detect,omitempty"`
	// +kubebuilder:default:=false
	Dynamicpopulate *bool `json:"dynamic-populate,omitempty"`
	//RootNetworkinstanceBridgetableProxyarpEvpn
	Evpn *DeviceNetworkinstanceBridgetableProxyarpEvpn `json:"evpn,omitempty"`
	// +kubebuilder:default:="never"
	Sendrefresh *string `json:"send-refresh,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=16383
	// +kubebuilder:default:=250
	Tablesize *uint32 `json:"table-size,omitempty"`
}

// DeviceNetworkinstanceBridgetableProxyarpDuplicatedetect struct
type DeviceNetworkinstanceBridgetableProxyarpDuplicatedetect struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}`
	Antispoofmac *string `json:"anti-spoof-mac,omitempty"`
	// +kubebuilder:default:="9"
	Holddowntime *string `json:"hold-down-time,omitempty"`
	// kubebuilder:validation:Minimum=3
	// kubebuilder:validation:Maximum=10
	// +kubebuilder:default:=5
	Nummoves *uint32 `json:"num-moves,omitempty"`
	// +kubebuilder:default:=false
	Staticblackhole *bool `json:"static-blackhole,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=15
	// +kubebuilder:default:=3
	Window *uint32 `json:"window,omitempty"`
}

// DeviceNetworkinstanceBridgetableProxyarpEvpn struct
type DeviceNetworkinstanceBridgetableProxyarpEvpn struct {
	//RootNetworkinstanceBridgetableProxyarpEvpnFlood
	Flood *DeviceNetworkinstanceBridgetableProxyarpEvpnFlood `json:"flood,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=0
	Routetag *uint32 `json:"route-tag,omitempty"`
}

// DeviceNetworkinstanceBridgetableProxyarpEvpnFlood struct
type DeviceNetworkinstanceBridgetableProxyarpEvpnFlood struct {
	// +kubebuilder:default:=true
	Gratuitousarp *bool `json:"gratuitous-arp,omitempty"`
	// +kubebuilder:default:=true
	Unknownarpreq *bool `json:"unknown-arp-req,omitempty"`
}

// DeviceNetworkinstanceBridgetableStaticmac struct
type DeviceNetworkinstanceBridgetableStaticmac struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Mac []*DeviceNetworkinstanceBridgetableStaticmacMac `json:"mac,omitempty"`
}

// DeviceNetworkinstanceBridgetableStaticmacMac struct
type DeviceNetworkinstanceBridgetableStaticmacMac struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}`
	Address     *string `json:"address"`
	Destination *string `json:"destination"`
}

// DeviceNetworkinstanceInterface struct
type DeviceNetworkinstanceInterface struct {
	// kubebuilder:validation:MinLength=5
	// kubebuilder:validation:MaxLength=25
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(mgmt0\.0|system0\.0|lo(0|1[0-9][0-9]|2([0-4][0-9]|5[0-5])|[1-9][0-9]|[1-9])\.(0|[1-9](\d){0,3})|ethernet-([1-9](\d){0,1}(/[abcd])?(/[1-9](\d){0,1})?/(([1-9](\d){0,1})|(1[0-1]\d)|(12[0-8])))\.([0]|[1-9](\d){0,3})|irb(0|1[0-9][0-9]|2([0-4][0-9]|5[0-5])|[1-9][0-9]|[1-9])\.(0|[1-9](\d){0,3})|lag(([1-9](\d){0,1})|(1[0-1]\d)|(12[0-8]))\.(0|[1-9](\d){0,3}))`
	Name *string `json:"name"`
}

// DeviceNetworkinstanceIpforwarding struct
type DeviceNetworkinstanceIpforwarding struct {
	Receiveipv4check *bool `json:"receive-ipv4-check,omitempty"`
	Receiveipv6check *bool `json:"receive-ipv6-check,omitempty"`
}

// DeviceNetworkinstanceIploadbalancing struct
type DeviceNetworkinstanceIploadbalancing struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Resilienthashprefix []*DeviceNetworkinstanceIploadbalancingResilienthashprefix `json:"resilient-hash-prefix,omitempty"`
}

// DeviceNetworkinstanceIploadbalancingResilienthashprefix struct
type DeviceNetworkinstanceIploadbalancingResilienthashprefix struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=32
	// +kubebuilder:default:=1
	Hashbucketsperpath *uint8 `json:"hash-buckets-per-path,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	Ipprefix *string `json:"ip-prefix"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=64
	// +kubebuilder:default:=1
	Maxpaths *uint8 `json:"max-paths,omitempty"`
}

// DeviceNetworkinstanceMpls struct
type DeviceNetworkinstanceMpls struct {
	Icmptunneling *bool `json:"icmp-tunneling,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Ndklabelblock []*DeviceNetworkinstanceMplsNdklabelblock `json:"ndk-label-block,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Staticentry      []*DeviceNetworkinstanceMplsStaticentry `json:"static-entry,omitempty"`
	Staticlabelblock *string                                 `json:"static-label-block,omitempty"`
}

// DeviceNetworkinstanceMplsNdklabelblock struct
type DeviceNetworkinstanceMplsNdklabelblock struct {
	Applicationname  *string `json:"application-name"`
	Staticlabelblock *string `json:"static-label-block,omitempty"`
}

// DeviceNetworkinstanceMplsStaticentry struct
type DeviceNetworkinstanceMplsStaticentry struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_DeviceNetworkinstanceMplsStaticentryAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// +kubebuilder:default:=false
	Collectstats *bool   `json:"collect-stats,omitempty"`
	Nexthopgroup *string `json:"next-hop-group,omitempty"`
	// +kubebuilder:validation:Enum=`pop`;`swap`
	// +kubebuilder:default:="swap"
	Operation E_DeviceNetworkinstanceMplsStaticentryOperation `json:"operation,omitempty"`
	//Operation *string `json:"operation,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	Preference *uint8 `json:"preference"`
	// kubebuilder:validation:Minimum=16
	// kubebuilder:validation:Maximum=1048575
	Toplabel *uint32 `json:"top-label"`
}

// DeviceNetworkinstanceMplsforwarding struct
type DeviceNetworkinstanceMplsforwarding struct {
	Forwardreceivedpackets *bool `json:"forward-received-packets,omitempty"`
}

// DeviceNetworkinstanceMtu struct
type DeviceNetworkinstanceMtu struct {
	// +kubebuilder:default:=true
	Pathmtudiscovery *bool `json:"path-mtu-discovery,omitempty"`
}

// DeviceNetworkinstanceNexthopgroups struct
type DeviceNetworkinstanceNexthopgroups struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Group []*DeviceNetworkinstanceNexthopgroupsGroup `json:"group,omitempty"`
}

// DeviceNetworkinstanceNexthopgroupsGroup struct
type DeviceNetworkinstanceNexthopgroupsGroup struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_DeviceNetworkinstanceNexthopgroupsGroupAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	//RootNetworkinstanceNexthopgroupsGroupBlackhole
	Blackhole *DeviceNetworkinstanceNexthopgroupsGroupBlackhole `json:"blackhole,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=128
	Nexthop []*DeviceNetworkinstanceNexthopgroupsGroupNexthop `json:"nexthop,omitempty"`
}

// DeviceNetworkinstanceNexthopgroupsGroupBlackhole struct
type DeviceNetworkinstanceNexthopgroupsGroupBlackhole struct {
	// +kubebuilder:default:=false
	Generateicmp *bool `json:"generate-icmp,omitempty"`
}

// DeviceNetworkinstanceNexthopgroupsGroupNexthop struct
type DeviceNetworkinstanceNexthopgroupsGroupNexthop struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_DeviceNetworkinstanceNexthopgroupsGroupNexthopAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	//RootNetworkinstanceNexthopgroupsGroupNexthopFailuredetection
	Failuredetection *DeviceNetworkinstanceNexthopgroupsGroupNexthopFailuredetection `json:"failure-detection,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	Index *uint16 `json:"index"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Ipaddress *string `json:"ip-address,omitempty"`
	//RootNetworkinstanceNexthopgroupsGroupNexthopPushedmplslabelstack
	Pushedmplslabelstack *DeviceNetworkinstanceNexthopgroupsGroupNexthopPushedmplslabelstack `json:"pushed-mpls-label-stack,omitempty"`
	// +kubebuilder:default:=true
	Resolve *bool `json:"resolve,omitempty"`
}

// DeviceNetworkinstanceNexthopgroupsGroupNexthopFailuredetection struct
type DeviceNetworkinstanceNexthopgroupsGroupNexthopFailuredetection struct {
	//RootNetworkinstanceNexthopgroupsGroupNexthopFailuredetectionEnablebfd
	Enablebfd *DeviceNetworkinstanceNexthopgroupsGroupNexthopFailuredetectionEnablebfd `json:"enable-bfd,omitempty"`
}

// DeviceNetworkinstanceNexthopgroupsGroupNexthopFailuredetectionEnablebfd struct
type DeviceNetworkinstanceNexthopgroupsGroupNexthopFailuredetectionEnablebfd struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Localaddress *string `json:"local-address"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=16384
	Localdiscriminator *uint32 `json:"local-discriminator,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=16384
	Remotediscriminator *uint32 `json:"remote-discriminator,omitempty"`
}

// DeviceNetworkinstanceNexthopgroupsGroupNexthopPushedmplslabelstack struct
type DeviceNetworkinstanceNexthopgroupsGroupNexthopPushedmplslabelstack struct {
	Pushedmplslabelstack *string `json:"pushed-mpls-label-stack,omitempty"`
}

// DeviceNetworkinstancePolicyforwarding struct
type DeviceNetworkinstancePolicyforwarding struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Interface []*DeviceNetworkinstancePolicyforwardingInterface `json:"interface,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=4
	Policy []*DeviceNetworkinstancePolicyforwardingPolicy `json:"policy,omitempty"`
}

// DeviceNetworkinstancePolicyforwardingInterface struct
type DeviceNetworkinstancePolicyforwardingInterface struct {
	Applyforwardingpolicy *string `json:"apply-forwarding-policy"`
	Subinterface          *string `json:"subinterface"`
}

// DeviceNetworkinstancePolicyforwardingPolicy struct
type DeviceNetworkinstancePolicyforwardingPolicy struct {
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Description *string `json:"description,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Policyid *string `json:"policy-id"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Rule []*DeviceNetworkinstancePolicyforwardingPolicyRule `json:"rule,omitempty"`
}

// DeviceNetworkinstancePolicyforwardingPolicyRule struct
type DeviceNetworkinstancePolicyforwardingPolicyRule struct {
	//RootNetworkinstancePolicyforwardingPolicyRuleAction
	Action *DeviceNetworkinstancePolicyforwardingPolicyRuleAction `json:"action,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Description *string `json:"description,omitempty"`
	//RootNetworkinstancePolicyforwardingPolicyRuleMatch
	Match *DeviceNetworkinstancePolicyforwardingPolicyRuleMatch `json:"match,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=128
	Sequenceid *uint32 `json:"sequence-id"`
}

// DeviceNetworkinstancePolicyforwardingPolicyRuleAction struct
type DeviceNetworkinstancePolicyforwardingPolicyRuleAction struct {
	Networkinstance *string `json:"network-instance,omitempty"`
}

// DeviceNetworkinstancePolicyforwardingPolicyRuleMatch struct
type DeviceNetworkinstancePolicyforwardingPolicyRuleMatch struct {
	//RootNetworkinstancePolicyforwardingPolicyRuleMatchIpv4
	Ipv4 *DeviceNetworkinstancePolicyforwardingPolicyRuleMatchIpv4 `json:"ipv4,omitempty"`
}

// DeviceNetworkinstancePolicyforwardingPolicyRuleMatchIpv4 struct
type DeviceNetworkinstancePolicyforwardingPolicyRuleMatchIpv4 struct {
	//RootNetworkinstancePolicyforwardingPolicyRuleMatchIpv4Dscpset
	Dscpset  *DeviceNetworkinstancePolicyforwardingPolicyRuleMatchIpv4Dscpset `json:"dscp-set,omitempty"`
	Protocol *string                                                          `json:"protocol,omitempty"`
}

// DeviceNetworkinstancePolicyforwardingPolicyRuleMatchIpv4Dscpset struct
type DeviceNetworkinstancePolicyforwardingPolicyRuleMatchIpv4Dscpset struct {
	Dscpset *string `json:"dscp-set,omitempty"`
}

// DeviceNetworkinstanceProtocols struct
type DeviceNetworkinstanceProtocols struct {
	//RootNetworkinstanceProtocolsBgp
	Bgp *DeviceNetworkinstanceProtocolsBgp `json:"bgp,omitempty"`
	//RootNetworkinstanceProtocolsBgpevpn
	Bgpevpn *DeviceNetworkinstanceProtocolsBgpevpn `json:"bgp-evpn,omitempty"`
	//RootNetworkinstanceProtocolsBgpvpn
	Bgpvpn *DeviceNetworkinstanceProtocolsBgpvpn `json:"bgp-vpn,omitempty"`
	//RootNetworkinstanceProtocolsDirectlyconnected
	Directlyconnected *DeviceNetworkinstanceProtocolsDirectlyconnected `json:"directly-connected,omitempty"`
	//RootNetworkinstanceProtocolsGribi
	Gribi *DeviceNetworkinstanceProtocolsGribi `json:"gribi,omitempty"`
	//RootNetworkinstanceProtocolsIgmp
	Igmp *DeviceNetworkinstanceProtocolsIgmp `json:"igmp,omitempty"`
	//RootNetworkinstanceProtocolsIsis
	Isis *DeviceNetworkinstanceProtocolsIsis `json:"isis,omitempty"`
	//RootNetworkinstanceProtocolsLdp
	Ldp *DeviceNetworkinstanceProtocolsLdp `json:"ldp,omitempty"`
	//RootNetworkinstanceProtocolsLinux
	Linux *DeviceNetworkinstanceProtocolsLinux `json:"linux,omitempty"`
	//RootNetworkinstanceProtocolsMld
	Mld *DeviceNetworkinstanceProtocolsMld `json:"mld,omitempty"`
	//RootNetworkinstanceProtocolsOspf
	Ospf *DeviceNetworkinstanceProtocolsOspf `json:"ospf,omitempty"`
	//RootNetworkinstanceProtocolsPim
	Pim *DeviceNetworkinstanceProtocolsPim `json:"pim,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgp struct
type DeviceNetworkinstanceProtocolsBgp struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_DeviceNetworkinstanceProtocolsBgpAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	//RootNetworkinstanceProtocolsBgpAspathoptions
	Aspathoptions *DeviceNetworkinstanceProtocolsBgpAspathoptions `json:"as-path-options,omitempty"`
	//RootNetworkinstanceProtocolsBgpAuthentication
	Authentication *DeviceNetworkinstanceProtocolsBgpAuthentication `json:"authentication,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	Autonomoussystem *uint32 `json:"autonomous-system"`
	//RootNetworkinstanceProtocolsBgpConvergence
	Convergence *DeviceNetworkinstanceProtocolsBgpConvergence `json:"convergence,omitempty"`
	//RootNetworkinstanceProtocolsBgpDynamicneighbors
	Dynamicneighbors *DeviceNetworkinstanceProtocolsBgpDynamicneighbors `json:"dynamic-neighbors,omitempty"`
	//RootNetworkinstanceProtocolsBgpEbgpdefaultpolicy
	Ebgpdefaultpolicy *DeviceNetworkinstanceProtocolsBgpEbgpdefaultpolicy `json:"ebgp-default-policy,omitempty"`
	//RootNetworkinstanceProtocolsBgpEvpn
	Evpn         *DeviceNetworkinstanceProtocolsBgpEvpn `json:"evpn,omitempty"`
	Exportpolicy *string                                `json:"export-policy,omitempty"`
	//RootNetworkinstanceProtocolsBgpFailuredetection
	Failuredetection *DeviceNetworkinstanceProtocolsBgpFailuredetection `json:"failure-detection,omitempty"`
	//RootNetworkinstanceProtocolsBgpGracefulrestart
	Gracefulrestart *DeviceNetworkinstanceProtocolsBgpGracefulrestart `json:"graceful-restart,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Group        []*DeviceNetworkinstanceProtocolsBgpGroup `json:"group,omitempty"`
	Importpolicy *string                                   `json:"import-policy,omitempty"`
	//RootNetworkinstanceProtocolsBgpIpv4unicast
	Ipv4unicast *DeviceNetworkinstanceProtocolsBgpIpv4unicast `json:"ipv4-unicast,omitempty"`
	//RootNetworkinstanceProtocolsBgpIpv6unicast
	Ipv6unicast *DeviceNetworkinstanceProtocolsBgpIpv6unicast `json:"ipv6-unicast,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	// +kubebuilder:default:=100
	Localpreference *uint32 `json:"local-preference,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Neighbor []*DeviceNetworkinstanceProtocolsBgpNeighbor `json:"neighbor,omitempty"`
	//RootNetworkinstanceProtocolsBgpPreference
	Preference *DeviceNetworkinstanceProtocolsBgpPreference `json:"preference,omitempty"`
	//RootNetworkinstanceProtocolsBgpRouteadvertisement
	Routeadvertisement *DeviceNetworkinstanceProtocolsBgpRouteadvertisement `json:"route-advertisement,omitempty"`
	//RootNetworkinstanceProtocolsBgpRoutereflector
	Routereflector *DeviceNetworkinstanceProtocolsBgpRoutereflector `json:"route-reflector,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Routerid *string `json:"router-id"`
	//RootNetworkinstanceProtocolsBgpSendcommunity
	Sendcommunity *DeviceNetworkinstanceProtocolsBgpSendcommunity `json:"send-community,omitempty"`
	//RootNetworkinstanceProtocolsBgpTraceoptions
	Traceoptions *DeviceNetworkinstanceProtocolsBgpTraceoptions `json:"trace-options,omitempty"`
	//RootNetworkinstanceProtocolsBgpTransport
	Transport *DeviceNetworkinstanceProtocolsBgpTransport `json:"transport,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpAspathoptions struct
type DeviceNetworkinstanceProtocolsBgpAspathoptions struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=0
	Allowownas *uint8 `json:"allow-own-as,omitempty"`
	//RootNetworkinstanceProtocolsBgpAspathoptionsRemoveprivateas
	Removeprivateas *DeviceNetworkinstanceProtocolsBgpAspathoptionsRemoveprivateas `json:"remove-private-as,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpAspathoptionsRemoveprivateas struct
type DeviceNetworkinstanceProtocolsBgpAspathoptionsRemoveprivateas struct {
	// +kubebuilder:default:=false
	Ignorepeeras *bool `json:"ignore-peer-as,omitempty"`
	// +kubebuilder:default:=false
	Leadingonly *bool `json:"leading-only,omitempty"`
	// +kubebuilder:validation:Enum=`delete`;`disabled`;`replace`
	// +kubebuilder:default:="disabled"
	Mode E_DeviceNetworkinstanceProtocolsBgpAspathoptionsRemoveprivateasMode `json:"mode,omitempty"`
	//Mode *string `json:"mode,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpAuthentication struct
type DeviceNetworkinstanceProtocolsBgpAuthentication struct {
	Keychain *string `json:"keychain,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpConvergence struct
type DeviceNetworkinstanceProtocolsBgpConvergence struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=3600
	// +kubebuilder:default:=0
	Minwaittoadvertise *uint16 `json:"min-wait-to-advertise,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpDynamicneighbors struct
type DeviceNetworkinstanceProtocolsBgpDynamicneighbors struct {
	//RootNetworkinstanceProtocolsBgpDynamicneighborsAccept
	Accept *DeviceNetworkinstanceProtocolsBgpDynamicneighborsAccept `json:"accept,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpDynamicneighborsAccept struct
type DeviceNetworkinstanceProtocolsBgpDynamicneighborsAccept struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Match []*DeviceNetworkinstanceProtocolsBgpDynamicneighborsAcceptMatch `json:"match,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=0
	Maxsessions *uint16 `json:"max-sessions,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpDynamicneighborsAcceptMatch struct
type DeviceNetworkinstanceProtocolsBgpDynamicneighborsAcceptMatch struct {
	//RootNetworkinstanceProtocolsBgpDynamicneighborsAcceptMatchAllowedpeeras
	Allowedpeeras *DeviceNetworkinstanceProtocolsBgpDynamicneighborsAcceptMatchAllowedpeeras `json:"allowed-peer-as,omitempty"`
	Peergroup     *string                                                                    `json:"peer-group"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	Prefix *string `json:"prefix"`
}

// DeviceNetworkinstanceProtocolsBgpDynamicneighborsAcceptMatchAllowedpeeras struct
type DeviceNetworkinstanceProtocolsBgpDynamicneighborsAcceptMatchAllowedpeeras struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`([1-9][0-9]*)|([1-9][0-9]*)\.\.([1-9][0-9]*)`
	Allowedpeeras *string `json:"allowed-peer-as,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpEbgpdefaultpolicy struct
type DeviceNetworkinstanceProtocolsBgpEbgpdefaultpolicy struct {
	// +kubebuilder:default:=true
	Exportrejectall *bool `json:"export-reject-all,omitempty"`
	// +kubebuilder:default:=true
	Importrejectall *bool `json:"import-reject-all,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpEvpn struct
type DeviceNetworkinstanceProtocolsBgpEvpn struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_DeviceNetworkinstanceProtocolsBgpEvpnAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// +kubebuilder:default:=false
	Advertiseipv6nexthops *bool `json:"advertise-ipv6-next-hops,omitempty"`
	// +kubebuilder:default:=false
	Interasvpn *bool `json:"inter-as-vpn,omitempty"`
	// +kubebuilder:default:=false
	Keepallroutes *bool `json:"keep-all-routes,omitempty"`
	// +kubebuilder:default:=false
	Rapidupdate *bool `json:"rapid-update,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpFailuredetection struct
type DeviceNetworkinstanceProtocolsBgpFailuredetection struct {
	// +kubebuilder:default:=false
	Enablebfd *bool `json:"enable-bfd,omitempty"`
	// +kubebuilder:default:=true
	Fastfailover *bool `json:"fast-failover,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpGracefulrestart struct
type DeviceNetworkinstanceProtocolsBgpGracefulrestart struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_DeviceNetworkinstanceProtocolsBgpGracefulrestartAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=3600
	// +kubebuilder:default:=360
	Staleroutestime *uint16 `json:"stale-routes-time,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpGroup struct
type DeviceNetworkinstanceProtocolsBgpGroup struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_DeviceNetworkinstanceProtocolsBgpGroupAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	//RootNetworkinstanceProtocolsBgpGroupAspathoptions
	Aspathoptions *DeviceNetworkinstanceProtocolsBgpGroupAspathoptions `json:"as-path-options,omitempty"`
	//RootNetworkinstanceProtocolsBgpGroupAuthentication
	Authentication *DeviceNetworkinstanceProtocolsBgpGroupAuthentication `json:"authentication,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Description *string `json:"description,omitempty"`
	//RootNetworkinstanceProtocolsBgpGroupEvpn
	Evpn         *DeviceNetworkinstanceProtocolsBgpGroupEvpn `json:"evpn,omitempty"`
	Exportpolicy *string                                     `json:"export-policy,omitempty"`
	//RootNetworkinstanceProtocolsBgpGroupFailuredetection
	Failuredetection *DeviceNetworkinstanceProtocolsBgpGroupFailuredetection `json:"failure-detection,omitempty"`
	//RootNetworkinstanceProtocolsBgpGroupGracefulrestart
	Gracefulrestart *DeviceNetworkinstanceProtocolsBgpGroupGracefulrestart `json:"graceful-restart,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=64
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Groupname    *string `json:"group-name"`
	Importpolicy *string `json:"import-policy,omitempty"`
	//RootNetworkinstanceProtocolsBgpGroupIpv4unicast
	Ipv4unicast *DeviceNetworkinstanceProtocolsBgpGroupIpv4unicast `json:"ipv4-unicast,omitempty"`
	//RootNetworkinstanceProtocolsBgpGroupIpv6unicast
	Ipv6unicast *DeviceNetworkinstanceProtocolsBgpGroupIpv6unicast `json:"ipv6-unicast,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1
	Localas []*DeviceNetworkinstanceProtocolsBgpGroupLocalas `json:"local-as,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	Localpreference *uint32 `json:"local-preference,omitempty"`
	//RootNetworkinstanceProtocolsBgpGroupMultihop
	Multihop *DeviceNetworkinstanceProtocolsBgpGroupMultihop `json:"multihop,omitempty"`
	// +kubebuilder:default:=false
	Nexthopself *bool `json:"next-hop-self,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	Peeras *uint32 `json:"peer-as,omitempty"`
	//RootNetworkinstanceProtocolsBgpGroupRoutereflector
	Routereflector *DeviceNetworkinstanceProtocolsBgpGroupRoutereflector `json:"route-reflector,omitempty"`
	//RootNetworkinstanceProtocolsBgpGroupSendcommunity
	Sendcommunity *DeviceNetworkinstanceProtocolsBgpGroupSendcommunity `json:"send-community,omitempty"`
	//RootNetworkinstanceProtocolsBgpGroupSenddefaultroute
	Senddefaultroute *DeviceNetworkinstanceProtocolsBgpGroupSenddefaultroute `json:"send-default-route,omitempty"`
	//RootNetworkinstanceProtocolsBgpGroupTimers
	Timers *DeviceNetworkinstanceProtocolsBgpGroupTimers `json:"timers,omitempty"`
	//RootNetworkinstanceProtocolsBgpGroupTraceoptions
	Traceoptions *DeviceNetworkinstanceProtocolsBgpGroupTraceoptions `json:"trace-options,omitempty"`
	//RootNetworkinstanceProtocolsBgpGroupTransport
	Transport *DeviceNetworkinstanceProtocolsBgpGroupTransport `json:"transport,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpGroupAspathoptions struct
type DeviceNetworkinstanceProtocolsBgpGroupAspathoptions struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	Allowownas *uint8 `json:"allow-own-as,omitempty"`
	//RootNetworkinstanceProtocolsBgpGroupAspathoptionsRemoveprivateas
	Removeprivateas *DeviceNetworkinstanceProtocolsBgpGroupAspathoptionsRemoveprivateas `json:"remove-private-as,omitempty"`
	Replacepeeras   *bool                                                               `json:"replace-peer-as,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpGroupAspathoptionsRemoveprivateas struct
type DeviceNetworkinstanceProtocolsBgpGroupAspathoptionsRemoveprivateas struct {
	// +kubebuilder:default:=false
	Ignorepeeras *bool `json:"ignore-peer-as,omitempty"`
	// +kubebuilder:default:=false
	Leadingonly *bool `json:"leading-only,omitempty"`
	// +kubebuilder:validation:Enum=`delete`;`disabled`;`replace`
	Mode E_DeviceNetworkinstanceProtocolsBgpGroupAspathoptionsRemoveprivateasMode `json:"mode,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpGroupAuthentication struct
type DeviceNetworkinstanceProtocolsBgpGroupAuthentication struct {
	Keychain *string `json:"keychain,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpGroupEvpn struct
type DeviceNetworkinstanceProtocolsBgpGroupEvpn struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	Adminstate E_DeviceNetworkinstanceProtocolsBgpGroupEvpnAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	Advertiseipv6nexthops *bool `json:"advertise-ipv6-next-hops,omitempty"`
	//RootNetworkinstanceProtocolsBgpGroupEvpnPrefixlimit
	Prefixlimit *DeviceNetworkinstanceProtocolsBgpGroupEvpnPrefixlimit `json:"prefix-limit,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpGroupEvpnPrefixlimit struct
type DeviceNetworkinstanceProtocolsBgpGroupEvpnPrefixlimit struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	// +kubebuilder:default:=4294967295
	Maxreceivedroutes *uint32 `json:"max-received-routes,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=90
	Warningthresholdpct *uint8 `json:"warning-threshold-pct,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpGroupFailuredetection struct
type DeviceNetworkinstanceProtocolsBgpGroupFailuredetection struct {
	Enablebfd    *bool `json:"enable-bfd,omitempty"`
	Fastfailover *bool `json:"fast-failover,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpGroupGracefulrestart struct
type DeviceNetworkinstanceProtocolsBgpGroupGracefulrestart struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	Adminstate E_DeviceNetworkinstanceProtocolsBgpGroupGracefulrestartAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=3600
	Staleroutestime *uint16 `json:"stale-routes-time,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpGroupIpv4unicast struct
type DeviceNetworkinstanceProtocolsBgpGroupIpv4unicast struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	Adminstate E_DeviceNetworkinstanceProtocolsBgpGroupIpv4unicastAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	Advertiseipv6nexthops *bool `json:"advertise-ipv6-next-hops,omitempty"`
	//RootNetworkinstanceProtocolsBgpGroupIpv4unicastPrefixlimit
	Prefixlimit         *DeviceNetworkinstanceProtocolsBgpGroupIpv4unicastPrefixlimit `json:"prefix-limit,omitempty"`
	Receiveipv6nexthops *bool                                                         `json:"receive-ipv6-next-hops,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpGroupIpv4unicastPrefixlimit struct
type DeviceNetworkinstanceProtocolsBgpGroupIpv4unicastPrefixlimit struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	// +kubebuilder:default:=4294967295
	Maxreceivedroutes *uint32 `json:"max-received-routes,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=90
	Warningthresholdpct *uint8 `json:"warning-threshold-pct,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpGroupIpv6unicast struct
type DeviceNetworkinstanceProtocolsBgpGroupIpv6unicast struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	Adminstate E_DeviceNetworkinstanceProtocolsBgpGroupIpv6unicastAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	//RootNetworkinstanceProtocolsBgpGroupIpv6unicastPrefixlimit
	Prefixlimit *DeviceNetworkinstanceProtocolsBgpGroupIpv6unicastPrefixlimit `json:"prefix-limit,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpGroupIpv6unicastPrefixlimit struct
type DeviceNetworkinstanceProtocolsBgpGroupIpv6unicastPrefixlimit struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	// +kubebuilder:default:=4294967295
	Maxreceivedroutes *uint32 `json:"max-received-routes,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=90
	Warningthresholdpct *uint8 `json:"warning-threshold-pct,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpGroupLocalas struct
type DeviceNetworkinstanceProtocolsBgpGroupLocalas struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	Asnumber *uint32 `json:"as-number"`
	// +kubebuilder:default:=true
	Prependglobalas *bool `json:"prepend-global-as,omitempty"`
	// +kubebuilder:default:=true
	Prependlocalas *bool `json:"prepend-local-as,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpGroupMultihop struct
type DeviceNetworkinstanceProtocolsBgpGroupMultihop struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	Adminstate E_DeviceNetworkinstanceProtocolsBgpGroupMultihopAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=255
	Maximumhops *uint8 `json:"maximum-hops,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpGroupRoutereflector struct
type DeviceNetworkinstanceProtocolsBgpGroupRoutereflector struct {
	Client *bool `json:"client,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Clusterid *string `json:"cluster-id,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpGroupSendcommunity struct
type DeviceNetworkinstanceProtocolsBgpGroupSendcommunity struct {
	Large    *bool `json:"large,omitempty"`
	Standard *bool `json:"standard,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpGroupSenddefaultroute struct
type DeviceNetworkinstanceProtocolsBgpGroupSenddefaultroute struct {
	Exportpolicy *string `json:"export-policy,omitempty"`
	// +kubebuilder:default:=false
	Ipv4unicast *bool `json:"ipv4-unicast,omitempty"`
	// +kubebuilder:default:=false
	Ipv6unicast *bool `json:"ipv6-unicast,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpGroupTimers struct
type DeviceNetworkinstanceProtocolsBgpGroupTimers struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=120
	Connectretry *uint16 `json:"connect-retry,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=0
	// +kubebuilder:default:=90
	Holdtime *uint16 `json:"hold-time,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=21845
	Keepaliveinterval *uint16 `json:"keepalive-interval,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=5
	Minimumadvertisementinterval *uint16 `json:"minimum-advertisement-interval,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpGroupTraceoptions struct
type DeviceNetworkinstanceProtocolsBgpGroupTraceoptions struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Flag []*DeviceNetworkinstanceProtocolsBgpGroupTraceoptionsFlag `json:"flag,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpGroupTraceoptionsFlag struct
type DeviceNetworkinstanceProtocolsBgpGroupTraceoptionsFlag struct {
	// +kubebuilder:validation:Enum=`detail`;`receive`;`send`
	Modifier E_DeviceNetworkinstanceProtocolsBgpGroupTraceoptionsFlagModifier `json:"modifier,omitempty"`
	//Modifier *string `json:"modifier,omitempty"`
	// +kubebuilder:validation:Enum=`events`;`graceful-restart`;`keepalive`;`notification`;`open`;`packets`;`route`;`socket`;`timers`;`update`
	Name E_DeviceNetworkinstanceProtocolsBgpGroupTraceoptionsFlagName `json:"name,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpGroupTransport struct
type DeviceNetworkinstanceProtocolsBgpGroupTransport struct {
	Localaddress *string `json:"local-address,omitempty"`
	// +kubebuilder:default:=false
	Passivemode *bool `json:"passive-mode,omitempty"`
	// kubebuilder:validation:Minimum=536
	// kubebuilder:validation:Maximum=9446
	Tcpmss *uint16 `json:"tcp-mss,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpIpv4unicast struct
type DeviceNetworkinstanceProtocolsBgpIpv4unicast struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_DeviceNetworkinstanceProtocolsBgpIpv4unicastAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// +kubebuilder:default:=false
	Advertiseipv6nexthops *bool `json:"advertise-ipv6-next-hops,omitempty"`
	//RootNetworkinstanceProtocolsBgpIpv4unicastConvergence
	Convergence *DeviceNetworkinstanceProtocolsBgpIpv4unicastConvergence `json:"convergence,omitempty"`
	//RootNetworkinstanceProtocolsBgpIpv4unicastMultipath
	Multipath *DeviceNetworkinstanceProtocolsBgpIpv4unicastMultipath `json:"multipath,omitempty"`
	//RootNetworkinstanceProtocolsBgpIpv4unicastNexthopresolution
	Nexthopresolution *DeviceNetworkinstanceProtocolsBgpIpv4unicastNexthopresolution `json:"next-hop-resolution,omitempty"`
	// +kubebuilder:default:=false
	Receiveipv6nexthops *bool `json:"receive-ipv6-next-hops,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpIpv4unicastConvergence struct
type DeviceNetworkinstanceProtocolsBgpIpv4unicastConvergence struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=3600
	// +kubebuilder:default:=0
	Maxwaittoadvertise *uint16 `json:"max-wait-to-advertise,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpIpv4unicastMultipath struct
type DeviceNetworkinstanceProtocolsBgpIpv4unicastMultipath struct {
	// +kubebuilder:default:=true
	Allowmultipleas *bool `json:"allow-multiple-as,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=64
	// +kubebuilder:default:=1
	Maxpathslevel1 *uint32 `json:"max-paths-level-1,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=64
	// +kubebuilder:default:=1
	Maxpathslevel2 *uint32 `json:"max-paths-level-2,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpIpv4unicastNexthopresolution struct
type DeviceNetworkinstanceProtocolsBgpIpv4unicastNexthopresolution struct {
	//RootNetworkinstanceProtocolsBgpIpv4unicastNexthopresolutionIpv4nexthops
	Ipv4nexthops *DeviceNetworkinstanceProtocolsBgpIpv4unicastNexthopresolutionIpv4nexthops `json:"ipv4-next-hops,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpIpv4unicastNexthopresolutionIpv4nexthops struct
type DeviceNetworkinstanceProtocolsBgpIpv4unicastNexthopresolutionIpv4nexthops struct {
	//RootNetworkinstanceProtocolsBgpIpv4unicastNexthopresolutionIpv4nexthopsTunnelresolution
	Tunnelresolution *DeviceNetworkinstanceProtocolsBgpIpv4unicastNexthopresolutionIpv4nexthopsTunnelresolution `json:"tunnel-resolution,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpIpv4unicastNexthopresolutionIpv4nexthopsTunnelresolution struct
type DeviceNetworkinstanceProtocolsBgpIpv4unicastNexthopresolutionIpv4nexthopsTunnelresolution struct {
	//RootNetworkinstanceProtocolsBgpIpv4unicastNexthopresolutionIpv4nexthopsTunnelresolutionAllowedtunneltypes
	Allowedtunneltypes *DeviceNetworkinstanceProtocolsBgpIpv4unicastNexthopresolutionIpv4nexthopsTunnelresolutionAllowedtunneltypes `json:"allowed-tunnel-types,omitempty"`
	// +kubebuilder:validation:Enum=`disabled`;`prefer`;`require`
	// +kubebuilder:default:="disabled"
	Mode E_DeviceNetworkinstanceProtocolsBgpIpv4unicastNexthopresolutionIpv4nexthopsTunnelresolutionMode `json:"mode,omitempty"`
	//Mode *string `json:"mode,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpIpv4unicastNexthopresolutionIpv4nexthopsTunnelresolutionAllowedtunneltypes struct
type DeviceNetworkinstanceProtocolsBgpIpv4unicastNexthopresolutionIpv4nexthopsTunnelresolutionAllowedtunneltypes struct {
	Allowedtunneltypes *string `json:"allowed-tunnel-types,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpIpv6unicast struct
type DeviceNetworkinstanceProtocolsBgpIpv6unicast struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_DeviceNetworkinstanceProtocolsBgpIpv6unicastAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	//RootNetworkinstanceProtocolsBgpIpv6unicastConvergence
	Convergence *DeviceNetworkinstanceProtocolsBgpIpv6unicastConvergence `json:"convergence,omitempty"`
	//RootNetworkinstanceProtocolsBgpIpv6unicastMultipath
	Multipath *DeviceNetworkinstanceProtocolsBgpIpv6unicastMultipath `json:"multipath,omitempty"`
	//RootNetworkinstanceProtocolsBgpIpv6unicastNexthopresolution
	Nexthopresolution *DeviceNetworkinstanceProtocolsBgpIpv6unicastNexthopresolution `json:"next-hop-resolution,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpIpv6unicastConvergence struct
type DeviceNetworkinstanceProtocolsBgpIpv6unicastConvergence struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=3600
	// +kubebuilder:default:=0
	Maxwaittoadvertise *uint16 `json:"max-wait-to-advertise,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpIpv6unicastMultipath struct
type DeviceNetworkinstanceProtocolsBgpIpv6unicastMultipath struct {
	// +kubebuilder:default:=true
	Allowmultipleas *bool `json:"allow-multiple-as,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=64
	// +kubebuilder:default:=1
	Maxpathslevel1 *uint32 `json:"max-paths-level-1,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=64
	// +kubebuilder:default:=1
	Maxpathslevel2 *uint32 `json:"max-paths-level-2,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpIpv6unicastNexthopresolution struct
type DeviceNetworkinstanceProtocolsBgpIpv6unicastNexthopresolution struct {
	//RootNetworkinstanceProtocolsBgpIpv6unicastNexthopresolutionIpv4nexthops
	Ipv4nexthops *DeviceNetworkinstanceProtocolsBgpIpv6unicastNexthopresolutionIpv4nexthops `json:"ipv4-next-hops,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpIpv6unicastNexthopresolutionIpv4nexthops struct
type DeviceNetworkinstanceProtocolsBgpIpv6unicastNexthopresolutionIpv4nexthops struct {
	//RootNetworkinstanceProtocolsBgpIpv6unicastNexthopresolutionIpv4nexthopsTunnelresolution
	Tunnelresolution *DeviceNetworkinstanceProtocolsBgpIpv6unicastNexthopresolutionIpv4nexthopsTunnelresolution `json:"tunnel-resolution,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpIpv6unicastNexthopresolutionIpv4nexthopsTunnelresolution struct
type DeviceNetworkinstanceProtocolsBgpIpv6unicastNexthopresolutionIpv4nexthopsTunnelresolution struct {
	//RootNetworkinstanceProtocolsBgpIpv6unicastNexthopresolutionIpv4nexthopsTunnelresolutionAllowedtunneltypes
	Allowedtunneltypes *DeviceNetworkinstanceProtocolsBgpIpv6unicastNexthopresolutionIpv4nexthopsTunnelresolutionAllowedtunneltypes `json:"allowed-tunnel-types,omitempty"`
	// +kubebuilder:validation:Enum=`disabled`;`prefer`;`require`
	// +kubebuilder:default:="disabled"
	Mode E_DeviceNetworkinstanceProtocolsBgpIpv6unicastNexthopresolutionIpv4nexthopsTunnelresolutionMode `json:"mode,omitempty"`
	//Mode *string `json:"mode,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpIpv6unicastNexthopresolutionIpv4nexthopsTunnelresolutionAllowedtunneltypes struct
type DeviceNetworkinstanceProtocolsBgpIpv6unicastNexthopresolutionIpv4nexthopsTunnelresolutionAllowedtunneltypes struct {
	Allowedtunneltypes *string `json:"allowed-tunnel-types,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpNeighbor struct
type DeviceNetworkinstanceProtocolsBgpNeighbor struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_DeviceNetworkinstanceProtocolsBgpNeighborAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	//RootNetworkinstanceProtocolsBgpNeighborAspathoptions
	Aspathoptions *DeviceNetworkinstanceProtocolsBgpNeighborAspathoptions `json:"as-path-options,omitempty"`
	//RootNetworkinstanceProtocolsBgpNeighborAuthentication
	Authentication *DeviceNetworkinstanceProtocolsBgpNeighborAuthentication `json:"authentication,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Description *string `json:"description,omitempty"`
	//RootNetworkinstanceProtocolsBgpNeighborEvpn
	Evpn         *DeviceNetworkinstanceProtocolsBgpNeighborEvpn `json:"evpn,omitempty"`
	Exportpolicy *string                                        `json:"export-policy,omitempty"`
	//RootNetworkinstanceProtocolsBgpNeighborFailuredetection
	Failuredetection *DeviceNetworkinstanceProtocolsBgpNeighborFailuredetection `json:"failure-detection,omitempty"`
	//RootNetworkinstanceProtocolsBgpNeighborGracefulrestart
	Gracefulrestart *DeviceNetworkinstanceProtocolsBgpNeighborGracefulrestart `json:"graceful-restart,omitempty"`
	Importpolicy    *string                                                   `json:"import-policy,omitempty"`
	//RootNetworkinstanceProtocolsBgpNeighborIpv4unicast
	Ipv4unicast *DeviceNetworkinstanceProtocolsBgpNeighborIpv4unicast `json:"ipv4-unicast,omitempty"`
	//RootNetworkinstanceProtocolsBgpNeighborIpv6unicast
	Ipv6unicast *DeviceNetworkinstanceProtocolsBgpNeighborIpv6unicast `json:"ipv6-unicast,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1
	Localas []*DeviceNetworkinstanceProtocolsBgpNeighborLocalas `json:"local-as,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	Localpreference *uint32 `json:"local-preference,omitempty"`
	//RootNetworkinstanceProtocolsBgpNeighborMultihop
	Multihop    *DeviceNetworkinstanceProtocolsBgpNeighborMultihop `json:"multihop,omitempty"`
	Nexthopself *bool                                              `json:"next-hop-self,omitempty"`
	// +kubebuilder:validation:Optional
	Peeraddress *string `json:"peer-address"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	Peeras    *uint32 `json:"peer-as,omitempty"`
	Peergroup *string `json:"peer-group"`
	//RootNetworkinstanceProtocolsBgpNeighborRoutereflector
	Routereflector *DeviceNetworkinstanceProtocolsBgpNeighborRoutereflector `json:"route-reflector,omitempty"`
	//RootNetworkinstanceProtocolsBgpNeighborSendcommunity
	Sendcommunity *DeviceNetworkinstanceProtocolsBgpNeighborSendcommunity `json:"send-community,omitempty"`
	//RootNetworkinstanceProtocolsBgpNeighborSenddefaultroute
	Senddefaultroute *DeviceNetworkinstanceProtocolsBgpNeighborSenddefaultroute `json:"send-default-route,omitempty"`
	//RootNetworkinstanceProtocolsBgpNeighborTimers
	Timers *DeviceNetworkinstanceProtocolsBgpNeighborTimers `json:"timers,omitempty"`
	//RootNetworkinstanceProtocolsBgpNeighborTraceoptions
	Traceoptions *DeviceNetworkinstanceProtocolsBgpNeighborTraceoptions `json:"trace-options,omitempty"`
	//RootNetworkinstanceProtocolsBgpNeighborTransport
	Transport *DeviceNetworkinstanceProtocolsBgpNeighborTransport `json:"transport,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpNeighborAspathoptions struct
type DeviceNetworkinstanceProtocolsBgpNeighborAspathoptions struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	Allowownas *uint8 `json:"allow-own-as,omitempty"`
	//RootNetworkinstanceProtocolsBgpNeighborAspathoptionsRemoveprivateas
	Removeprivateas *DeviceNetworkinstanceProtocolsBgpNeighborAspathoptionsRemoveprivateas `json:"remove-private-as,omitempty"`
	Replacepeeras   *bool                                                                  `json:"replace-peer-as,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpNeighborAspathoptionsRemoveprivateas struct
type DeviceNetworkinstanceProtocolsBgpNeighborAspathoptionsRemoveprivateas struct {
	// +kubebuilder:default:=false
	Ignorepeeras *bool `json:"ignore-peer-as,omitempty"`
	// +kubebuilder:default:=false
	Leadingonly *bool `json:"leading-only,omitempty"`
	// +kubebuilder:validation:Enum=`delete`;`disabled`;`replace`
	Mode E_DeviceNetworkinstanceProtocolsBgpNeighborAspathoptionsRemoveprivateasMode `json:"mode,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpNeighborAuthentication struct
type DeviceNetworkinstanceProtocolsBgpNeighborAuthentication struct {
	Keychain *string `json:"keychain,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpNeighborEvpn struct
type DeviceNetworkinstanceProtocolsBgpNeighborEvpn struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	Adminstate E_DeviceNetworkinstanceProtocolsBgpNeighborEvpnAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	Advertiseipv6nexthops *bool `json:"advertise-ipv6-next-hops,omitempty"`
	//RootNetworkinstanceProtocolsBgpNeighborEvpnPrefixlimit
	Prefixlimit *DeviceNetworkinstanceProtocolsBgpNeighborEvpnPrefixlimit `json:"prefix-limit,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpNeighborEvpnPrefixlimit struct
type DeviceNetworkinstanceProtocolsBgpNeighborEvpnPrefixlimit struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	Maxreceivedroutes *uint32 `json:"max-received-routes,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100
	Warningthresholdpct *uint8 `json:"warning-threshold-pct,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpNeighborFailuredetection struct
type DeviceNetworkinstanceProtocolsBgpNeighborFailuredetection struct {
	Enablebfd    *bool `json:"enable-bfd,omitempty"`
	Fastfailover *bool `json:"fast-failover,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpNeighborGracefulrestart struct
type DeviceNetworkinstanceProtocolsBgpNeighborGracefulrestart struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	Adminstate E_DeviceNetworkinstanceProtocolsBgpNeighborGracefulrestartAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=3600
	Staleroutestime *uint16 `json:"stale-routes-time,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpNeighborIpv4unicast struct
type DeviceNetworkinstanceProtocolsBgpNeighborIpv4unicast struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	Adminstate E_DeviceNetworkinstanceProtocolsBgpNeighborIpv4unicastAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	Advertiseipv6nexthops *bool `json:"advertise-ipv6-next-hops,omitempty"`
	//RootNetworkinstanceProtocolsBgpNeighborIpv4unicastPrefixlimit
	Prefixlimit         *DeviceNetworkinstanceProtocolsBgpNeighborIpv4unicastPrefixlimit `json:"prefix-limit,omitempty"`
	Receiveipv6nexthops *bool                                                            `json:"receive-ipv6-next-hops,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpNeighborIpv4unicastPrefixlimit struct
type DeviceNetworkinstanceProtocolsBgpNeighborIpv4unicastPrefixlimit struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	Maxreceivedroutes *uint32 `json:"max-received-routes,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100
	Warningthresholdpct *uint8 `json:"warning-threshold-pct,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpNeighborIpv6unicast struct
type DeviceNetworkinstanceProtocolsBgpNeighborIpv6unicast struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	Adminstate E_DeviceNetworkinstanceProtocolsBgpNeighborIpv6unicastAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	//RootNetworkinstanceProtocolsBgpNeighborIpv6unicastPrefixlimit
	Prefixlimit *DeviceNetworkinstanceProtocolsBgpNeighborIpv6unicastPrefixlimit `json:"prefix-limit,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpNeighborIpv6unicastPrefixlimit struct
type DeviceNetworkinstanceProtocolsBgpNeighborIpv6unicastPrefixlimit struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	Maxreceivedroutes *uint32 `json:"max-received-routes,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100
	Warningthresholdpct *uint8 `json:"warning-threshold-pct,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpNeighborLocalas struct
type DeviceNetworkinstanceProtocolsBgpNeighborLocalas struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	Asnumber        *uint32 `json:"as-number"`
	Prependglobalas *bool   `json:"prepend-global-as,omitempty"`
	Prependlocalas  *bool   `json:"prepend-local-as,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpNeighborMultihop struct
type DeviceNetworkinstanceProtocolsBgpNeighborMultihop struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	Adminstate E_DeviceNetworkinstanceProtocolsBgpNeighborMultihopAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=255
	Maximumhops *uint8 `json:"maximum-hops,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpNeighborRoutereflector struct
type DeviceNetworkinstanceProtocolsBgpNeighborRoutereflector struct {
	Client *bool `json:"client,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Clusterid *string `json:"cluster-id,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpNeighborSendcommunity struct
type DeviceNetworkinstanceProtocolsBgpNeighborSendcommunity struct {
	Large    *bool `json:"large,omitempty"`
	Standard *bool `json:"standard,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpNeighborSenddefaultroute struct
type DeviceNetworkinstanceProtocolsBgpNeighborSenddefaultroute struct {
	Exportpolicy *string `json:"export-policy,omitempty"`
	Ipv4unicast  *bool   `json:"ipv4-unicast,omitempty"`
	Ipv6unicast  *bool   `json:"ipv6-unicast,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpNeighborTimers struct
type DeviceNetworkinstanceProtocolsBgpNeighborTimers struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=65535
	Connectretry *uint16 `json:"connect-retry,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=0
	Holdtime *uint16 `json:"hold-time,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=21845
	Keepaliveinterval *uint16 `json:"keepalive-interval,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=255
	Minimumadvertisementinterval *uint16 `json:"minimum-advertisement-interval,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpNeighborTraceoptions struct
type DeviceNetworkinstanceProtocolsBgpNeighborTraceoptions struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Flag []*DeviceNetworkinstanceProtocolsBgpNeighborTraceoptionsFlag `json:"flag,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpNeighborTraceoptionsFlag struct
type DeviceNetworkinstanceProtocolsBgpNeighborTraceoptionsFlag struct {
	// +kubebuilder:validation:Enum=`detail`;`receive`;`send`
	Modifier E_DeviceNetworkinstanceProtocolsBgpNeighborTraceoptionsFlagModifier `json:"modifier,omitempty"`
	//Modifier *string `json:"modifier,omitempty"`
	// +kubebuilder:validation:Enum=`events`;`graceful-restart`;`keepalive`;`notification`;`open`;`packets`;`route`;`socket`;`timers`;`update`
	Name E_DeviceNetworkinstanceProtocolsBgpNeighborTraceoptionsFlagName `json:"name,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpNeighborTransport struct
type DeviceNetworkinstanceProtocolsBgpNeighborTransport struct {
	Localaddress *string `json:"local-address,omitempty"`
	Passivemode  *bool   `json:"passive-mode,omitempty"`
	// kubebuilder:validation:Minimum=536
	// kubebuilder:validation:Maximum=9446
	Tcpmss *uint16 `json:"tcp-mss,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpPreference struct
type DeviceNetworkinstanceProtocolsBgpPreference struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=170
	Ebgp *uint8 `json:"ebgp,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=170
	Ibgp *uint8 `json:"ibgp,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpRouteadvertisement struct
type DeviceNetworkinstanceProtocolsBgpRouteadvertisement struct {
	// +kubebuilder:default:=false
	Rapidwithdrawal *bool `json:"rapid-withdrawal,omitempty"`
	// +kubebuilder:default:=true
	Waitforfibinstall *bool `json:"wait-for-fib-install,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpRoutereflector struct
type DeviceNetworkinstanceProtocolsBgpRoutereflector struct {
	// +kubebuilder:default:=false
	Client *bool `json:"client,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Clusterid *string `json:"cluster-id,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpSendcommunity struct
type DeviceNetworkinstanceProtocolsBgpSendcommunity struct {
	// +kubebuilder:default:=true
	Large *bool `json:"large,omitempty"`
	// +kubebuilder:default:=true
	Standard *bool `json:"standard,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpTraceoptions struct
type DeviceNetworkinstanceProtocolsBgpTraceoptions struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Flag []*DeviceNetworkinstanceProtocolsBgpTraceoptionsFlag `json:"flag,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpTraceoptionsFlag struct
type DeviceNetworkinstanceProtocolsBgpTraceoptionsFlag struct {
	// +kubebuilder:validation:Enum=`detail`;`receive`;`send`
	Modifier E_DeviceNetworkinstanceProtocolsBgpTraceoptionsFlagModifier `json:"modifier,omitempty"`
	//Modifier *string `json:"modifier,omitempty"`
	// +kubebuilder:validation:Enum=`events`;`graceful-restart`;`keepalive`;`notification`;`open`;`packets`;`route`;`socket`;`timers`;`update`
	Name E_DeviceNetworkinstanceProtocolsBgpTraceoptionsFlagName `json:"name,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpTransport struct
type DeviceNetworkinstanceProtocolsBgpTransport struct {
	// kubebuilder:validation:Minimum=536
	// kubebuilder:validation:Maximum=9446
	// +kubebuilder:default:=1024
	Tcpmss *uint16 `json:"tcp-mss,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpevpn struct
type DeviceNetworkinstanceProtocolsBgpevpn struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1
	Bgpinstance []*DeviceNetworkinstanceProtocolsBgpevpnBgpinstance `json:"bgp-instance,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpevpnBgpinstance struct
type DeviceNetworkinstanceProtocolsBgpevpnBgpinstance struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_DeviceNetworkinstanceProtocolsBgpevpnBgpinstanceAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	// +kubebuilder:default:=0
	Defaultadmintag *uint32 `json:"default-admin-tag,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=8
	// +kubebuilder:default:=1
	Ecmp *uint8 `json:"ecmp,omitempty"`
	// +kubebuilder:validation:Enum=`vxlan`
	// +kubebuilder:default:="vxlan"
	Encapsulationtype E_DeviceNetworkinstanceProtocolsBgpevpnBgpinstanceEncapsulationtype `json:"encapsulation-type,omitempty"`
	//Encapsulationtype *string `json:"encapsulation-type,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=65535
	Evi *uint32 `json:"evi"`
	Id  *string `json:"id"`
	//RootNetworkinstanceProtocolsBgpevpnBgpinstanceRoutes
	Routes         *DeviceNetworkinstanceProtocolsBgpevpnBgpinstanceRoutes `json:"routes,omitempty"`
	Vxlaninterface *string                                                 `json:"vxlan-interface,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpevpnBgpinstanceRoutes struct
type DeviceNetworkinstanceProtocolsBgpevpnBgpinstanceRoutes struct {
	//RootNetworkinstanceProtocolsBgpevpnBgpinstanceRoutesBridgetable
	Bridgetable *DeviceNetworkinstanceProtocolsBgpevpnBgpinstanceRoutesBridgetable `json:"bridge-table,omitempty"`
	//RootNetworkinstanceProtocolsBgpevpnBgpinstanceRoutesRoutetable
	Routetable *DeviceNetworkinstanceProtocolsBgpevpnBgpinstanceRoutesRoutetable `json:"route-table,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpevpnBgpinstanceRoutesBridgetable struct
type DeviceNetworkinstanceProtocolsBgpevpnBgpinstanceRoutesBridgetable struct {
	//RootNetworkinstanceProtocolsBgpevpnBgpinstanceRoutesBridgetableInclusivemcast
	Inclusivemcast *DeviceNetworkinstanceProtocolsBgpevpnBgpinstanceRoutesBridgetableInclusivemcast `json:"inclusive-mcast,omitempty"`
	//RootNetworkinstanceProtocolsBgpevpnBgpinstanceRoutesBridgetableMacip
	Macip *DeviceNetworkinstanceProtocolsBgpevpnBgpinstanceRoutesBridgetableMacip `json:"mac-ip,omitempty"`
	// +kubebuilder:default:="use-system-ipv4-address"
	Nexthop *string `json:"next-hop,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=16777215
	// +kubebuilder:default:=0
	Vlanawarebundleethtag *uint32 `json:"vlan-aware-bundle-eth-tag,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpevpnBgpinstanceRoutesBridgetableInclusivemcast struct
type DeviceNetworkinstanceProtocolsBgpevpnBgpinstanceRoutesBridgetableInclusivemcast struct {
	// +kubebuilder:default:=true
	Advertise *bool `json:"advertise,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Originatingip *string `json:"originating-ip,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpevpnBgpinstanceRoutesBridgetableMacip struct
type DeviceNetworkinstanceProtocolsBgpevpnBgpinstanceRoutesBridgetableMacip struct {
	// +kubebuilder:default:=true
	Advertise *bool `json:"advertise,omitempty"`
	// +kubebuilder:default:=false
	Advertisearpndonlywithmactableentry *bool `json:"advertise-arp-nd-only-with-mac-table-entry,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpevpnBgpinstanceRoutesRoutetable struct
type DeviceNetworkinstanceProtocolsBgpevpnBgpinstanceRoutesRoutetable struct {
	//RootNetworkinstanceProtocolsBgpevpnBgpinstanceRoutesRoutetableMacip
	Macip *DeviceNetworkinstanceProtocolsBgpevpnBgpinstanceRoutesRoutetableMacip `json:"mac-ip,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpevpnBgpinstanceRoutesRoutetableMacip struct
type DeviceNetworkinstanceProtocolsBgpevpnBgpinstanceRoutesRoutetableMacip struct {
	// +kubebuilder:default:=false
	Advertisegatewaymac *bool `json:"advertise-gateway-mac,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpvpn struct
type DeviceNetworkinstanceProtocolsBgpvpn struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1
	Bgpinstance []*DeviceNetworkinstanceProtocolsBgpvpnBgpinstance `json:"bgp-instance,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpvpnBgpinstance struct
type DeviceNetworkinstanceProtocolsBgpvpnBgpinstance struct {
	Exportpolicy *string `json:"export-policy,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=2
	Id           *uint8  `json:"id"`
	Importpolicy *string `json:"import-policy,omitempty"`
	//RootNetworkinstanceProtocolsBgpvpnBgpinstanceRoutedistinguisher
	Routedistinguisher *DeviceNetworkinstanceProtocolsBgpvpnBgpinstanceRoutedistinguisher `json:"route-distinguisher,omitempty"`
	//RootNetworkinstanceProtocolsBgpvpnBgpinstanceRoutetarget
	Routetarget *DeviceNetworkinstanceProtocolsBgpvpnBgpinstanceRoutetarget `json:"route-target,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpvpnBgpinstanceRoutedistinguisher struct
type DeviceNetworkinstanceProtocolsBgpvpnBgpinstanceRoutedistinguisher struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|42[0-8][0-9]{7}|4[0-1][0-9]{8}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9])|(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])|(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|42[0-8][0-9]{7}|4[0-1][0-9]{8}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])|(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]).(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])`
	Rd *string `json:"rd,omitempty"`
}

// DeviceNetworkinstanceProtocolsBgpvpnBgpinstanceRoutetarget struct
type DeviceNetworkinstanceProtocolsBgpvpnBgpinstanceRoutetarget struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|42[0-8][0-9]{7}|4[0-1][0-9]{8}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9])|(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])|(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|42[0-8][0-9]{7}|4[0-1][0-9]{8}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])|target:(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|42[0-8][0-9]{7}|4[0-1][0-9]{8}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9])|target:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])|target:(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|42[0-8][0-9]{7}|4[0-1][0-9]{8}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])|origin:(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|42[0-8][0-9]{7}|4[0-1][0-9]{8}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9])|origin:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])|origin:(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|42[0-8][0-9]{7}|4[0-1][0-9]{8}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])|color:[0-1]{2}:(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9])`
	Exportrt *string `json:"export-rt,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|42[0-8][0-9]{7}|4[0-1][0-9]{8}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9])|(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])|(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|42[0-8][0-9]{7}|4[0-1][0-9]{8}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])|target:(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|42[0-8][0-9]{7}|4[0-1][0-9]{8}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9])|target:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])|target:(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|42[0-8][0-9]{7}|4[0-1][0-9]{8}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])|origin:(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|42[0-8][0-9]{7}|4[0-1][0-9]{8}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9])|origin:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])|origin:(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|42[0-8][0-9]{7}|4[0-1][0-9]{8}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])|color:[0-1]{2}:(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9])`
	Importrt *string `json:"import-rt,omitempty"`
}

// DeviceNetworkinstanceProtocolsDirectlyconnected struct
type DeviceNetworkinstanceProtocolsDirectlyconnected struct {
	//RootNetworkinstanceProtocolsDirectlyconnectedTedatabaseinstall
	Tedatabaseinstall *DeviceNetworkinstanceProtocolsDirectlyconnectedTedatabaseinstall `json:"te-database-install,omitempty"`
}

// DeviceNetworkinstanceProtocolsDirectlyconnectedTedatabaseinstall struct
type DeviceNetworkinstanceProtocolsDirectlyconnectedTedatabaseinstall struct {
	//RootNetworkinstanceProtocolsDirectlyconnectedTedatabaseinstallBgpls
	Bgpls *DeviceNetworkinstanceProtocolsDirectlyconnectedTedatabaseinstallBgpls `json:"bgp-ls,omitempty"`
}

// DeviceNetworkinstanceProtocolsDirectlyconnectedTedatabaseinstallBgpls struct
type DeviceNetworkinstanceProtocolsDirectlyconnectedTedatabaseinstallBgpls struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	Bgplsidentifier *uint32 `json:"bgp-ls-identifier,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=-1
	Igpidentifier *uint64 `json:"igp-identifier,omitempty"`
}

// DeviceNetworkinstanceProtocolsGribi struct
type DeviceNetworkinstanceProtocolsGribi struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_DeviceNetworkinstanceProtocolsGribiAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	// +kubebuilder:default:=1
	Defaultmetric *uint32 `json:"default-metric,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=4
	Defaultpreference *uint8 `json:"default-preference,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=128
	// +kubebuilder:default:=64
	Maxecmppaths *uint8 `json:"max-ecmp-paths,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	// +kubebuilder:default:=0
	Maximumroutes *uint32 `json:"maximum-routes,omitempty"`
}

// DeviceNetworkinstanceProtocolsIgmp struct
type DeviceNetworkinstanceProtocolsIgmp struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_DeviceNetworkinstanceProtocolsIgmpAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Interface []*DeviceNetworkinstanceProtocolsIgmpInterface `json:"interface,omitempty"`
	// kubebuilder:validation:Minimum=2
	// kubebuilder:validation:Maximum=1024
	// +kubebuilder:default:=125
	Queryinterval *uint32 `json:"query-interval,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=1023
	// +kubebuilder:default:=1
	Querylastmemberinterval *uint32 `json:"query-last-member-interval,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=1023
	// +kubebuilder:default:=10
	Queryresponseinterval *uint32 `json:"query-response-interval,omitempty"`
	// kubebuilder:validation:Minimum=2
	// kubebuilder:validation:Maximum=10
	// +kubebuilder:default:=2
	Robustcount *uint32 `json:"robust-count,omitempty"`
	//RootNetworkinstanceProtocolsIgmpSsm
	Ssm *DeviceNetworkinstanceProtocolsIgmpSsm `json:"ssm,omitempty"`
	//RootNetworkinstanceProtocolsIgmpTraceoptions
	Traceoptions *DeviceNetworkinstanceProtocolsIgmpTraceoptions `json:"trace-options,omitempty"`
}

// DeviceNetworkinstanceProtocolsIgmpInterface struct
type DeviceNetworkinstanceProtocolsIgmpInterface struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_DeviceNetworkinstanceProtocolsIgmpInterfaceAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	Importpolicy  *string `json:"import-policy,omitempty"`
	Interfacename *string `json:"interface-name"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4096
	// +kubebuilder:default:=0
	Maxgroupsources *uint32 `json:"max-group-sources,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4096
	// +kubebuilder:default:=0
	Maxgroups *uint32 `json:"max-groups,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=512
	// +kubebuilder:default:=0
	Maxsources *uint32 `json:"max-sources,omitempty"`
	// kubebuilder:validation:Minimum=2
	// kubebuilder:validation:Maximum=1024
	Queryinterval *uint32 `json:"query-interval,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=1023
	Querylastmemberinterval *uint32 `json:"query-last-member-interval,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=1023
	Queryresponseinterval *uint32 `json:"query-response-interval,omitempty"`
	// +kubebuilder:default:=true
	Routeralertcheck *bool `json:"router-alert-check,omitempty"`
	//RootNetworkinstanceProtocolsIgmpInterfaceSsm
	Ssm *DeviceNetworkinstanceProtocolsIgmpInterfaceSsm `json:"ssm,omitempty"`
	//RootNetworkinstanceProtocolsIgmpInterfaceStatic
	Static *DeviceNetworkinstanceProtocolsIgmpInterfaceStatic `json:"static,omitempty"`
	// +kubebuilder:default:=true
	Subnetcheck *bool `json:"subnet-check,omitempty"`
	//RootNetworkinstanceProtocolsIgmpInterfaceTraceoptions
	Traceoptions *DeviceNetworkinstanceProtocolsIgmpInterfaceTraceoptions `json:"trace-options,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=3
	// +kubebuilder:default:=3
	Version *uint8 `json:"version,omitempty"`
}

// DeviceNetworkinstanceProtocolsIgmpInterfaceSsm struct
type DeviceNetworkinstanceProtocolsIgmpInterfaceSsm struct {
	//RootNetworkinstanceProtocolsIgmpInterfaceSsmMappings
	Mappings *DeviceNetworkinstanceProtocolsIgmpInterfaceSsmMappings `json:"mappings,omitempty"`
}

// DeviceNetworkinstanceProtocolsIgmpInterfaceSsmMappings struct
type DeviceNetworkinstanceProtocolsIgmpInterfaceSsmMappings struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Grouprange []*DeviceNetworkinstanceProtocolsIgmpInterfaceSsmMappingsGrouprange `json:"group-range,omitempty"`
}

// DeviceNetworkinstanceProtocolsIgmpInterfaceSsmMappingsGrouprange struct
type DeviceNetworkinstanceProtocolsIgmpInterfaceSsmMappingsGrouprange struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	End *string `json:"end"`
	//+kubebuilder:validation:MinItems=1
	//+kubebuilder:validation:MaxItems=1024
	Source []*DeviceNetworkinstanceProtocolsIgmpInterfaceSsmMappingsGrouprangeSource `json:"source,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Start *string `json:"start"`
}

// DeviceNetworkinstanceProtocolsIgmpInterfaceSsmMappingsGrouprangeSource struct
type DeviceNetworkinstanceProtocolsIgmpInterfaceSsmMappingsGrouprangeSource struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Source *string `json:"source"`
}

// DeviceNetworkinstanceProtocolsIgmpInterfaceStatic struct
type DeviceNetworkinstanceProtocolsIgmpInterfaceStatic struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Grouprange []*DeviceNetworkinstanceProtocolsIgmpInterfaceStaticGrouprange `json:"group-range,omitempty"`
}

// DeviceNetworkinstanceProtocolsIgmpInterfaceStaticGrouprange struct
type DeviceNetworkinstanceProtocolsIgmpInterfaceStaticGrouprange struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	End *string `json:"end"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Source []*DeviceNetworkinstanceProtocolsIgmpInterfaceStaticGrouprangeSource `json:"source,omitempty"`
	Starg  *string                                                              `json:"starg,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Start *string `json:"start"`
}

// DeviceNetworkinstanceProtocolsIgmpInterfaceStaticGrouprangeSource struct
type DeviceNetworkinstanceProtocolsIgmpInterfaceStaticGrouprangeSource struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Source *string `json:"source"`
}

// DeviceNetworkinstanceProtocolsIgmpInterfaceTraceoptions struct
type DeviceNetworkinstanceProtocolsIgmpInterfaceTraceoptions struct {
	//RootNetworkinstanceProtocolsIgmpInterfaceTraceoptionsTrace
	Trace *DeviceNetworkinstanceProtocolsIgmpInterfaceTraceoptionsTrace `json:"trace,omitempty"`
}

// DeviceNetworkinstanceProtocolsIgmpInterfaceTraceoptionsTrace struct
type DeviceNetworkinstanceProtocolsIgmpInterfaceTraceoptionsTrace struct {
	Interfaces *string `json:"interfaces,omitempty"`
	//RootNetworkinstanceProtocolsIgmpInterfaceTraceoptionsTracePacket
	Packet *DeviceNetworkinstanceProtocolsIgmpInterfaceTraceoptionsTracePacket `json:"packet,omitempty"`
}

// DeviceNetworkinstanceProtocolsIgmpInterfaceTraceoptionsTracePacket struct
type DeviceNetworkinstanceProtocolsIgmpInterfaceTraceoptionsTracePacket struct {
	// +kubebuilder:validation:Enum=`dropped`;`egress-ingress-and-dropped`;`ingress-and-dropped`
	Modifier E_DeviceNetworkinstanceProtocolsIgmpInterfaceTraceoptionsTracePacketModifier `json:"modifier,omitempty"`
	//Modifier *string `json:"modifier,omitempty"`
	// +kubebuilder:validation:Enum=`all`;`query`;`v1-report`;`v2-leave`;`v2-report`;`v3-report`
	Type E_DeviceNetworkinstanceProtocolsIgmpInterfaceTraceoptionsTracePacketType `json:"type,omitempty"`
	//Type *string `json:"type,omitempty"`
}

// DeviceNetworkinstanceProtocolsIgmpSsm struct
type DeviceNetworkinstanceProtocolsIgmpSsm struct {
	//RootNetworkinstanceProtocolsIgmpSsmMappings
	Mappings *DeviceNetworkinstanceProtocolsIgmpSsmMappings `json:"mappings,omitempty"`
}

// DeviceNetworkinstanceProtocolsIgmpSsmMappings struct
type DeviceNetworkinstanceProtocolsIgmpSsmMappings struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Grouprange []*DeviceNetworkinstanceProtocolsIgmpSsmMappingsGrouprange `json:"group-range,omitempty"`
}

// DeviceNetworkinstanceProtocolsIgmpSsmMappingsGrouprange struct
type DeviceNetworkinstanceProtocolsIgmpSsmMappingsGrouprange struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	End *string `json:"end"`
	//+kubebuilder:validation:MinItems=1
	//+kubebuilder:validation:MaxItems=1024
	Source []*DeviceNetworkinstanceProtocolsIgmpSsmMappingsGrouprangeSource `json:"source,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Start *string `json:"start"`
}

// DeviceNetworkinstanceProtocolsIgmpSsmMappingsGrouprangeSource struct
type DeviceNetworkinstanceProtocolsIgmpSsmMappingsGrouprangeSource struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Source *string `json:"source"`
}

// DeviceNetworkinstanceProtocolsIgmpTraceoptions struct
type DeviceNetworkinstanceProtocolsIgmpTraceoptions struct {
	//RootNetworkinstanceProtocolsIgmpTraceoptionsTrace
	Trace *DeviceNetworkinstanceProtocolsIgmpTraceoptionsTrace `json:"trace,omitempty"`
}

// DeviceNetworkinstanceProtocolsIgmpTraceoptionsTrace struct
type DeviceNetworkinstanceProtocolsIgmpTraceoptionsTrace struct {
	Interfaces *string `json:"interfaces,omitempty"`
	//RootNetworkinstanceProtocolsIgmpTraceoptionsTracePacket
	Packet *DeviceNetworkinstanceProtocolsIgmpTraceoptionsTracePacket `json:"packet,omitempty"`
}

// DeviceNetworkinstanceProtocolsIgmpTraceoptionsTracePacket struct
type DeviceNetworkinstanceProtocolsIgmpTraceoptionsTracePacket struct {
	// +kubebuilder:validation:Enum=`dropped`;`egress-ingress-and-dropped`;`ingress-and-dropped`
	Modifier E_DeviceNetworkinstanceProtocolsIgmpTraceoptionsTracePacketModifier `json:"modifier,omitempty"`
	//Modifier *string `json:"modifier,omitempty"`
	// +kubebuilder:validation:Enum=`all`;`query`;`v1-report`;`v2-leave`;`v2-report`;`v3-report`
	Type E_DeviceNetworkinstanceProtocolsIgmpTraceoptionsTracePacketType `json:"type,omitempty"`
	//Type *string `json:"type,omitempty"`
}

// DeviceNetworkinstanceProtocolsIsis struct
type DeviceNetworkinstanceProtocolsIsis struct {
	Dynamiclabelblock *string `json:"dynamic-label-block,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1
	Instance []*DeviceNetworkinstanceProtocolsIsisInstance `json:"instance,omitempty"`
}

// DeviceNetworkinstanceProtocolsIsisInstance struct
type DeviceNetworkinstanceProtocolsIsisInstance struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_DeviceNetworkinstanceProtocolsIsisInstanceAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	//RootNetworkinstanceProtocolsIsisInstanceAttachedbit
	Attachedbit *DeviceNetworkinstanceProtocolsIsisInstanceAttachedbit `json:"attached-bit,omitempty"`
	//RootNetworkinstanceProtocolsIsisInstanceAuthentication
	Authentication *DeviceNetworkinstanceProtocolsIsisInstanceAuthentication `json:"authentication,omitempty"`
	//RootNetworkinstanceProtocolsIsisInstanceAutocost
	Autocost     *DeviceNetworkinstanceProtocolsIsisInstanceAutocost `json:"auto-cost,omitempty"`
	Exportpolicy *string                                             `json:"export-policy,omitempty"`
	//RootNetworkinstanceProtocolsIsisInstanceGracefulrestart
	Gracefulrestart *DeviceNetworkinstanceProtocolsIsisInstanceGracefulrestart `json:"graceful-restart,omitempty"`
	//RootNetworkinstanceProtocolsIsisInstanceInterlevelpropagationpolicies
	Interlevelpropagationpolicies *DeviceNetworkinstanceProtocolsIsisInstanceInterlevelpropagationpolicies `json:"inter-level-propagation-policies,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Interface []*DeviceNetworkinstanceProtocolsIsisInstanceInterface `json:"interface,omitempty"`
	//RootNetworkinstanceProtocolsIsisInstanceIpv4unicast
	Ipv4unicast *DeviceNetworkinstanceProtocolsIsisInstanceIpv4unicast `json:"ipv4-unicast,omitempty"`
	//RootNetworkinstanceProtocolsIsisInstanceIpv6unicast
	Ipv6unicast *DeviceNetworkinstanceProtocolsIsisInstanceIpv6unicast `json:"ipv6-unicast,omitempty"`
	//RootNetworkinstanceProtocolsIsisInstanceLdpsynchronization
	Ldpsynchronization *DeviceNetworkinstanceProtocolsIsisInstanceLdpsynchronization `json:"ldp-synchronization,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=2
	Level []*DeviceNetworkinstanceProtocolsIsisInstanceLevel `json:"level,omitempty"`
	// +kubebuilder:validation:Enum=`L1`;`L1L2`;`L2`
	// +kubebuilder:default:="L2"
	Levelcapability E_DeviceNetworkinstanceProtocolsIsisInstanceLevelcapability `json:"level-capability,omitempty"`
	//Levelcapability *string `json:"level-capability,omitempty"`
	Lspauthentication *bool `json:"lsp-authentication,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=64
	// +kubebuilder:default:=1
	Maxecmppaths *uint8 `json:"max-ecmp-paths,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
	//RootNetworkinstanceProtocolsIsisInstanceNet
	Net *DeviceNetworkinstanceProtocolsIsisInstanceNet `json:"net,omitempty"`
	//RootNetworkinstanceProtocolsIsisInstanceOverload
	Overload *DeviceNetworkinstanceProtocolsIsisInstanceOverload `json:"overload,omitempty"`
	// +kubebuilder:default:=false
	Poitlv *bool `json:"poi-tlv,omitempty"`
	//RootNetworkinstanceProtocolsIsisInstanceSegmentrouting
	Segmentrouting *DeviceNetworkinstanceProtocolsIsisInstanceSegmentrouting `json:"segment-routing,omitempty"`
	//RootNetworkinstanceProtocolsIsisInstanceTedatabaseinstall
	Tedatabaseinstall *DeviceNetworkinstanceProtocolsIsisInstanceTedatabaseinstall `json:"te-database-install,omitempty"`
	//RootNetworkinstanceProtocolsIsisInstanceTimers
	Timers *DeviceNetworkinstanceProtocolsIsisInstanceTimers `json:"timers,omitempty"`
	//RootNetworkinstanceProtocolsIsisInstanceTraceoptions
	Traceoptions *DeviceNetworkinstanceProtocolsIsisInstanceTraceoptions `json:"trace-options,omitempty"`
	//RootNetworkinstanceProtocolsIsisInstanceTrafficengineering
	Trafficengineering *DeviceNetworkinstanceProtocolsIsisInstanceTrafficengineering `json:"traffic-engineering,omitempty"`
	//RootNetworkinstanceProtocolsIsisInstanceTransport
	Transport *DeviceNetworkinstanceProtocolsIsisInstanceTransport `json:"transport,omitempty"`
}

// DeviceNetworkinstanceProtocolsIsisInstanceAttachedbit struct
type DeviceNetworkinstanceProtocolsIsisInstanceAttachedbit struct {
	// +kubebuilder:default:=false
	Ignore *bool `json:"ignore,omitempty"`
	// +kubebuilder:default:=false
	Suppress *bool `json:"suppress,omitempty"`
}

// DeviceNetworkinstanceProtocolsIsisInstanceAuthentication struct
type DeviceNetworkinstanceProtocolsIsisInstanceAuthentication struct {
	// +kubebuilder:validation:Enum=`cleartext`;`hmac-md5`
	Algorithm E_DeviceNetworkinstanceProtocolsIsisInstanceAuthenticationAlgorithm `json:"algorithm,omitempty"`
	//Algorithm *string `json:"algorithm,omitempty"`
	// +kubebuilder:validation:Enum=`disable`;`loose`;`strict`
	Checkreceived E_DeviceNetworkinstanceProtocolsIsisInstanceAuthenticationCheckreceived `json:"check-received,omitempty"`
	//Checkreceived *string `json:"check-received,omitempty"`
	Csnpauthentication  *bool `json:"csnp-authentication,omitempty"`
	Helloauthentication *bool `json:"hello-authentication,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=25
	Key                *string `json:"key,omitempty"`
	Keychain           *string `json:"keychain,omitempty"`
	Psnpauthentication *bool   `json:"psnp-authentication,omitempty"`
}

// DeviceNetworkinstanceProtocolsIsisInstanceAutocost struct
type DeviceNetworkinstanceProtocolsIsisInstanceAutocost struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=8000000000
	Referencebandwidth *uint64 `json:"reference-bandwidth,omitempty"`
}

// DeviceNetworkinstanceProtocolsIsisInstanceGracefulrestart struct
type DeviceNetworkinstanceProtocolsIsisInstanceGracefulrestart struct {
	// +kubebuilder:default:=false
	Helpermode *bool `json:"helper-mode,omitempty"`
}

// DeviceNetworkinstanceProtocolsIsisInstanceInterlevelpropagationpolicies struct
type DeviceNetworkinstanceProtocolsIsisInstanceInterlevelpropagationpolicies struct {
	//RootNetworkinstanceProtocolsIsisInstanceInterlevelpropagationpoliciesLevel1tolevel2
	Level1tolevel2 *DeviceNetworkinstanceProtocolsIsisInstanceInterlevelpropagationpoliciesLevel1tolevel2 `json:"level1-to-level2,omitempty"`
}

// DeviceNetworkinstanceProtocolsIsisInstanceInterlevelpropagationpoliciesLevel1tolevel2 struct
type DeviceNetworkinstanceProtocolsIsisInstanceInterlevelpropagationpoliciesLevel1tolevel2 struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Summaryaddress []*DeviceNetworkinstanceProtocolsIsisInstanceInterlevelpropagationpoliciesLevel1tolevel2Summaryaddress `json:"summary-address,omitempty"`
}

// DeviceNetworkinstanceProtocolsIsisInstanceInterlevelpropagationpoliciesLevel1tolevel2Summaryaddress struct
type DeviceNetworkinstanceProtocolsIsisInstanceInterlevelpropagationpoliciesLevel1tolevel2Summaryaddress struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	Ipprefix *string `json:"ip-prefix"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	Routetag *uint32 `json:"route-tag,omitempty"`
}

// DeviceNetworkinstanceProtocolsIsisInstanceInterface struct
type DeviceNetworkinstanceProtocolsIsisInstanceInterface struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_DeviceNetworkinstanceProtocolsIsisInstanceInterfaceAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	//RootNetworkinstanceProtocolsIsisInstanceInterfaceAuthentication
	Authentication *DeviceNetworkinstanceProtocolsIsisInstanceInterfaceAuthentication `json:"authentication,omitempty"`
	// +kubebuilder:validation:Enum=`broadcast`;`point-to-point`
	Circuittype E_DeviceNetworkinstanceProtocolsIsisInstanceInterfaceCircuittype `json:"circuit-type,omitempty"`
	//Circuittype *string `json:"circuit-type,omitempty"`
	// +kubebuilder:validation:Enum=`adaptive`;`disable`;`loose`;`strict`
	// +kubebuilder:default:="disable"
	Hellopadding E_DeviceNetworkinstanceProtocolsIsisInstanceInterfaceHellopadding `json:"hello-padding,omitempty"`
	//Hellopadding *string `json:"hello-padding,omitempty"`
	Interfacename *string `json:"interface-name"`
	//RootNetworkinstanceProtocolsIsisInstanceInterfaceIpv4unicast
	Ipv4unicast *DeviceNetworkinstanceProtocolsIsisInstanceInterfaceIpv4unicast `json:"ipv4-unicast,omitempty"`
	//RootNetworkinstanceProtocolsIsisInstanceInterfaceIpv6unicast
	Ipv6unicast *DeviceNetworkinstanceProtocolsIsisInstanceInterfaceIpv6unicast `json:"ipv6-unicast,omitempty"`
	//RootNetworkinstanceProtocolsIsisInstanceInterfaceLdpsynchronization
	Ldpsynchronization *DeviceNetworkinstanceProtocolsIsisInstanceInterfaceLdpsynchronization `json:"ldp-synchronization,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=2
	Level []*DeviceNetworkinstanceProtocolsIsisInstanceInterfaceLevel `json:"level,omitempty"`
	// +kubebuilder:default:=false
	Passive *bool `json:"passive,omitempty"`
	//RootNetworkinstanceProtocolsIsisInstanceInterfaceSegmentrouting
	Segmentrouting *DeviceNetworkinstanceProtocolsIsisInstanceInterfaceSegmentrouting `json:"segment-routing,omitempty"`
	//RootNetworkinstanceProtocolsIsisInstanceInterfaceTimers
	Timers *DeviceNetworkinstanceProtocolsIsisInstanceInterfaceTimers `json:"timers,omitempty"`
	//RootNetworkinstanceProtocolsIsisInstanceInterfaceTraceoptions
	Traceoptions *DeviceNetworkinstanceProtocolsIsisInstanceInterfaceTraceoptions `json:"trace-options,omitempty"`
}

// DeviceNetworkinstanceProtocolsIsisInstanceInterfaceAuthentication struct
type DeviceNetworkinstanceProtocolsIsisInstanceInterfaceAuthentication struct {
	// +kubebuilder:validation:Enum=`cleartext`;`hmac-md5`
	Algorithm E_DeviceNetworkinstanceProtocolsIsisInstanceInterfaceAuthenticationAlgorithm `json:"algorithm,omitempty"`
	//Algorithm *string `json:"algorithm,omitempty"`
	Helloauthentication *bool `json:"hello-authentication,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=25
	Key      *string `json:"key,omitempty"`
	Keychain *string `json:"keychain,omitempty"`
}

// DeviceNetworkinstanceProtocolsIsisInstanceInterfaceIpv4unicast struct
type DeviceNetworkinstanceProtocolsIsisInstanceInterfaceIpv4unicast struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_DeviceNetworkinstanceProtocolsIsisInstanceInterfaceIpv4unicastAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// +kubebuilder:default:=false
	Enablebfd *bool `json:"enable-bfd,omitempty"`
	// +kubebuilder:default:=false
	Includebfdtlv *bool `json:"include-bfd-tlv,omitempty"`
}

// DeviceNetworkinstanceProtocolsIsisInstanceInterfaceIpv6unicast struct
type DeviceNetworkinstanceProtocolsIsisInstanceInterfaceIpv6unicast struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_DeviceNetworkinstanceProtocolsIsisInstanceInterfaceIpv6unicastAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// +kubebuilder:default:=false
	Enablebfd *bool `json:"enable-bfd,omitempty"`
	// +kubebuilder:default:=false
	Includebfdtlv *bool `json:"include-bfd-tlv,omitempty"`
}

// DeviceNetworkinstanceProtocolsIsisInstanceInterfaceLdpsynchronization struct
type DeviceNetworkinstanceProtocolsIsisInstanceInterfaceLdpsynchronization struct {
	Disable  *string `json:"disable,omitempty"`
	Endoflib *bool   `json:"end-of-lib,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=1800
	Holddowntimer *uint16 `json:"hold-down-timer,omitempty"`
}

// DeviceNetworkinstanceProtocolsIsisInstanceInterfaceLevel struct
type DeviceNetworkinstanceProtocolsIsisInstanceInterfaceLevel struct {
	//RootNetworkinstanceProtocolsIsisInstanceInterfaceLevelAuthentication
	Authentication *DeviceNetworkinstanceProtocolsIsisInstanceInterfaceLevelAuthentication `json:"authentication,omitempty"`
	// +kubebuilder:default:=false
	Disable *bool `json:"disable,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=16777215
	Ipv6unicastmetric *uint32 `json:"ipv6-unicast-metric,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=2
	Levelnumber *uint8 `json:"level-number"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=16777215
	Metric *uint32 `json:"metric,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=127
	// +kubebuilder:default:=64
	Priority *uint8 `json:"priority,omitempty"`
	//RootNetworkinstanceProtocolsIsisInstanceInterfaceLevelTimers
	Timers *DeviceNetworkinstanceProtocolsIsisInstanceInterfaceLevelTimers `json:"timers,omitempty"`
}

// DeviceNetworkinstanceProtocolsIsisInstanceInterfaceLevelAuthentication struct
type DeviceNetworkinstanceProtocolsIsisInstanceInterfaceLevelAuthentication struct {
	// +kubebuilder:validation:Enum=`cleartext`;`hmac-md5`
	Algorithm E_DeviceNetworkinstanceProtocolsIsisInstanceInterfaceLevelAuthenticationAlgorithm `json:"algorithm,omitempty"`
	//Algorithm *string `json:"algorithm,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=25
	Key      *string `json:"key,omitempty"`
	Keychain *string `json:"keychain,omitempty"`
}

// DeviceNetworkinstanceProtocolsIsisInstanceInterfaceLevelTimers struct
type DeviceNetworkinstanceProtocolsIsisInstanceInterfaceLevelTimers struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=20000
	// +kubebuilder:default:=9
	Hellointerval *uint32 `json:"hello-interval,omitempty"`
	// kubebuilder:validation:Minimum=2
	// kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=3
	Hellomultiplier *uint8 `json:"hello-multiplier,omitempty"`
}

// DeviceNetworkinstanceProtocolsIsisInstanceInterfaceSegmentrouting struct
type DeviceNetworkinstanceProtocolsIsisInstanceInterfaceSegmentrouting struct {
	//RootNetworkinstanceProtocolsIsisInstanceInterfaceSegmentroutingMpls
	Mpls *DeviceNetworkinstanceProtocolsIsisInstanceInterfaceSegmentroutingMpls `json:"mpls,omitempty"`
}

// DeviceNetworkinstanceProtocolsIsisInstanceInterfaceSegmentroutingMpls struct
type DeviceNetworkinstanceProtocolsIsisInstanceInterfaceSegmentroutingMpls struct {
	//RootNetworkinstanceProtocolsIsisInstanceInterfaceSegmentroutingMplsIpv4adjacencysid
	Ipv4adjacencysid *DeviceNetworkinstanceProtocolsIsisInstanceInterfaceSegmentroutingMplsIpv4adjacencysid `json:"ipv4-adjacency-sid,omitempty"`
	//RootNetworkinstanceProtocolsIsisInstanceInterfaceSegmentroutingMplsIpv4nodesid
	Ipv4nodesid *DeviceNetworkinstanceProtocolsIsisInstanceInterfaceSegmentroutingMplsIpv4nodesid `json:"ipv4-node-sid,omitempty"`
	//RootNetworkinstanceProtocolsIsisInstanceInterfaceSegmentroutingMplsIpv6adjacencysid
	Ipv6adjacencysid *DeviceNetworkinstanceProtocolsIsisInstanceInterfaceSegmentroutingMplsIpv6adjacencysid `json:"ipv6-adjacency-sid,omitempty"`
	//RootNetworkinstanceProtocolsIsisInstanceInterfaceSegmentroutingMplsIpv6nodesid
	Ipv6nodesid *DeviceNetworkinstanceProtocolsIsisInstanceInterfaceSegmentroutingMplsIpv6nodesid `json:"ipv6-node-sid,omitempty"`
}

// DeviceNetworkinstanceProtocolsIsisInstanceInterfaceSegmentroutingMplsIpv4adjacencysid struct
type DeviceNetworkinstanceProtocolsIsisInstanceInterfaceSegmentroutingMplsIpv4adjacencysid struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=1048575
	Labelvalue *uint32 `json:"label-value,omitempty"`
}

// DeviceNetworkinstanceProtocolsIsisInstanceInterfaceSegmentroutingMplsIpv4nodesid struct
type DeviceNetworkinstanceProtocolsIsisInstanceInterfaceSegmentroutingMplsIpv4nodesid struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=1048575
	Index *uint32 `json:"index,omitempty"`
}

// DeviceNetworkinstanceProtocolsIsisInstanceInterfaceSegmentroutingMplsIpv6adjacencysid struct
type DeviceNetworkinstanceProtocolsIsisInstanceInterfaceSegmentroutingMplsIpv6adjacencysid struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=1048575
	Labelvalue *uint32 `json:"label-value,omitempty"`
}

// DeviceNetworkinstanceProtocolsIsisInstanceInterfaceSegmentroutingMplsIpv6nodesid struct
type DeviceNetworkinstanceProtocolsIsisInstanceInterfaceSegmentroutingMplsIpv6nodesid struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=1048575
	Index *uint32 `json:"index,omitempty"`
}

// DeviceNetworkinstanceProtocolsIsisInstanceInterfaceTimers struct
type DeviceNetworkinstanceProtocolsIsisInstanceInterfaceTimers struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=10
	Csnpinterval *uint16 `json:"csnp-interval,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100000
	// +kubebuilder:default:=100
	Lsppacinginterval *uint64 `json:"lsp-pacing-interval,omitempty"`
}

// DeviceNetworkinstanceProtocolsIsisInstanceInterfaceTraceoptions struct
type DeviceNetworkinstanceProtocolsIsisInstanceInterfaceTraceoptions struct {
	//RootNetworkinstanceProtocolsIsisInstanceInterfaceTraceoptionsTrace
	Trace *DeviceNetworkinstanceProtocolsIsisInstanceInterfaceTraceoptionsTrace `json:"trace,omitempty"`
}

// DeviceNetworkinstanceProtocolsIsisInstanceInterfaceTraceoptionsTrace struct
type DeviceNetworkinstanceProtocolsIsisInstanceInterfaceTraceoptionsTrace struct {
	// +kubebuilder:validation:Enum=`adjacencies`;`packets-all`;`packets-l1-csnp`;`packets-l1-hello`;`packets-l1-lsp`;`packets-l1-psnp`;`packets-l2-csnp`;`packets-l2-hello`;`packets-l2-lsp`;`packets-l2-psnp`;`packets-p2p-hello`
	Trace E_DeviceNetworkinstanceProtocolsIsisInstanceInterfaceTraceoptionsTraceTrace `json:"trace,omitempty"`
	//Trace *string `json:"trace,omitempty"`
}

// DeviceNetworkinstanceProtocolsIsisInstanceIpv4unicast struct
type DeviceNetworkinstanceProtocolsIsisInstanceIpv4unicast struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_DeviceNetworkinstanceProtocolsIsisInstanceIpv4unicastAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
}

// DeviceNetworkinstanceProtocolsIsisInstanceIpv6unicast struct
type DeviceNetworkinstanceProtocolsIsisInstanceIpv6unicast struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_DeviceNetworkinstanceProtocolsIsisInstanceIpv6unicastAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// +kubebuilder:default:=false
	Multitopology *bool `json:"multi-topology,omitempty"`
}

// DeviceNetworkinstanceProtocolsIsisInstanceLdpsynchronization struct
type DeviceNetworkinstanceProtocolsIsisInstanceLdpsynchronization struct {
	// +kubebuilder:default:=false
	Endoflib *bool `json:"end-of-lib,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=1800
	// +kubebuilder:default:=60
	Holddowntimer *uint16 `json:"hold-down-timer,omitempty"`
}

// DeviceNetworkinstanceProtocolsIsisInstanceLevel struct
type DeviceNetworkinstanceProtocolsIsisInstanceLevel struct {
	//RootNetworkinstanceProtocolsIsisInstanceLevelAuthentication
	Authentication *DeviceNetworkinstanceProtocolsIsisInstanceLevelAuthentication `json:"authentication,omitempty"`
	// +kubebuilder:default:=false
	Bgplsexclude *bool `json:"bgp-ls-exclude,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=2
	Levelnumber *uint8 `json:"level-number"`
	// +kubebuilder:validation:Enum=`narrow`;`wide`
	// +kubebuilder:default:="wide"
	Metricstyle E_DeviceNetworkinstanceProtocolsIsisInstanceLevelMetricstyle `json:"metric-style,omitempty"`
	//Metricstyle *string `json:"metric-style,omitempty"`
	//RootNetworkinstanceProtocolsIsisInstanceLevelRoutepreference
	Routepreference *DeviceNetworkinstanceProtocolsIsisInstanceLevelRoutepreference `json:"route-preference,omitempty"`
	//RootNetworkinstanceProtocolsIsisInstanceLevelTraceoptions
	Traceoptions *DeviceNetworkinstanceProtocolsIsisInstanceLevelTraceoptions `json:"trace-options,omitempty"`
}

// DeviceNetworkinstanceProtocolsIsisInstanceLevelAuthentication struct
type DeviceNetworkinstanceProtocolsIsisInstanceLevelAuthentication struct {
	// +kubebuilder:validation:Enum=`cleartext`;`hmac-md5`
	Algorithm E_DeviceNetworkinstanceProtocolsIsisInstanceLevelAuthenticationAlgorithm `json:"algorithm,omitempty"`
	//Algorithm *string `json:"algorithm,omitempty"`
	// +kubebuilder:validation:Enum=`disable`;`loose`;`strict`
	Checkreceived E_DeviceNetworkinstanceProtocolsIsisInstanceLevelAuthenticationCheckreceived `json:"check-received,omitempty"`
	//Checkreceived *string `json:"check-received,omitempty"`
	Csnpauthentication  *bool `json:"csnp-authentication,omitempty"`
	Helloauthentication *bool `json:"hello-authentication,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=25
	Key                *string `json:"key,omitempty"`
	Keychain           *string `json:"keychain,omitempty"`
	Psnpauthentication *bool   `json:"psnp-authentication,omitempty"`
}

// DeviceNetworkinstanceProtocolsIsisInstanceLevelRoutepreference struct
type DeviceNetworkinstanceProtocolsIsisInstanceLevelRoutepreference struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=255
	External *uint8 `json:"external,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=255
	Internal *uint8 `json:"internal,omitempty"`
}

// DeviceNetworkinstanceProtocolsIsisInstanceLevelTraceoptions struct
type DeviceNetworkinstanceProtocolsIsisInstanceLevelTraceoptions struct {
	//RootNetworkinstanceProtocolsIsisInstanceLevelTraceoptionsTrace
	Trace *DeviceNetworkinstanceProtocolsIsisInstanceLevelTraceoptionsTrace `json:"trace,omitempty"`
}

// DeviceNetworkinstanceProtocolsIsisInstanceLevelTraceoptionsTrace struct
type DeviceNetworkinstanceProtocolsIsisInstanceLevelTraceoptionsTrace struct {
	// +kubebuilder:validation:Enum=`adjacencies`;`lsdb`;`routes`;`spf`
	Trace E_DeviceNetworkinstanceProtocolsIsisInstanceLevelTraceoptionsTraceTrace `json:"trace,omitempty"`
	//Trace *string `json:"trace,omitempty"`
}

// DeviceNetworkinstanceProtocolsIsisInstanceNet struct
type DeviceNetworkinstanceProtocolsIsisInstanceNet struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[a-fA-F0-9]{2}(\.[a-fA-F0-9]{4}){3,9}\.[0]{2}`
	Net *string `json:"net,omitempty"`
}

// DeviceNetworkinstanceProtocolsIsisInstanceOverload struct
type DeviceNetworkinstanceProtocolsIsisInstanceOverload struct {
	// +kubebuilder:default:=false
	Advertiseexternal *bool `json:"advertise-external,omitempty"`
	// +kubebuilder:default:=false
	Advertiseinterlevel *bool `json:"advertise-interlevel,omitempty"`
	//RootNetworkinstanceProtocolsIsisInstanceOverloadImmediate
	Immediate *DeviceNetworkinstanceProtocolsIsisInstanceOverloadImmediate `json:"immediate,omitempty"`
	//RootNetworkinstanceProtocolsIsisInstanceOverloadOnboot
	Onboot *DeviceNetworkinstanceProtocolsIsisInstanceOverloadOnboot `json:"on-boot,omitempty"`
}

// DeviceNetworkinstanceProtocolsIsisInstanceOverloadImmediate struct
type DeviceNetworkinstanceProtocolsIsisInstanceOverloadImmediate struct {
	// +kubebuilder:default:=false
	Maxmetric *bool `json:"max-metric,omitempty"`
	// +kubebuilder:default:=false
	Setbit *bool `json:"set-bit,omitempty"`
}

// DeviceNetworkinstanceProtocolsIsisInstanceOverloadOnboot struct
type DeviceNetworkinstanceProtocolsIsisInstanceOverloadOnboot struct {
	Maxmetric *bool `json:"max-metric,omitempty"`
	Setbit    *bool `json:"set-bit,omitempty"`
	// kubebuilder:validation:Minimum=60
	// kubebuilder:validation:Maximum=1800
	Timeout *uint16 `json:"timeout,omitempty"`
}

// DeviceNetworkinstanceProtocolsIsisInstanceSegmentrouting struct
type DeviceNetworkinstanceProtocolsIsisInstanceSegmentrouting struct {
	//RootNetworkinstanceProtocolsIsisInstanceSegmentroutingMpls
	Mpls *DeviceNetworkinstanceProtocolsIsisInstanceSegmentroutingMpls `json:"mpls,omitempty"`
}

// DeviceNetworkinstanceProtocolsIsisInstanceSegmentroutingMpls struct
type DeviceNetworkinstanceProtocolsIsisInstanceSegmentroutingMpls struct {
	// +kubebuilder:default:=15
	Adjacencysidholdtime *uint16 `json:"adjacency-sid-hold-time,omitempty"`
	Staticlabelblock     *string `json:"static-label-block,omitempty"`
}

// DeviceNetworkinstanceProtocolsIsisInstanceTedatabaseinstall struct
type DeviceNetworkinstanceProtocolsIsisInstanceTedatabaseinstall struct {
	//RootNetworkinstanceProtocolsIsisInstanceTedatabaseinstallBgpls
	Bgpls *DeviceNetworkinstanceProtocolsIsisInstanceTedatabaseinstallBgpls `json:"bgp-ls,omitempty"`
}

// DeviceNetworkinstanceProtocolsIsisInstanceTedatabaseinstallBgpls struct
type DeviceNetworkinstanceProtocolsIsisInstanceTedatabaseinstallBgpls struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	Bgplsidentifier *uint32 `json:"bgp-ls-identifier,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=-1
	Igpidentifier *uint64 `json:"igp-identifier,omitempty"`
}

// DeviceNetworkinstanceProtocolsIsisInstanceTimers struct
type DeviceNetworkinstanceProtocolsIsisInstanceTimers struct {
	//RootNetworkinstanceProtocolsIsisInstanceTimersLspgeneration
	Lspgeneration *DeviceNetworkinstanceProtocolsIsisInstanceTimersLspgeneration `json:"lsp-generation,omitempty"`
	// kubebuilder:validation:Minimum=350
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=1200
	Lsplifetime *uint16 `json:"lsp-lifetime,omitempty"`
	//RootNetworkinstanceProtocolsIsisInstanceTimersLsprefresh
	Lsprefresh *DeviceNetworkinstanceProtocolsIsisInstanceTimersLsprefresh `json:"lsp-refresh,omitempty"`
	//RootNetworkinstanceProtocolsIsisInstanceTimersSpf
	Spf *DeviceNetworkinstanceProtocolsIsisInstanceTimersSpf `json:"spf,omitempty"`
}

// DeviceNetworkinstanceProtocolsIsisInstanceTimersLspgeneration struct
type DeviceNetworkinstanceProtocolsIsisInstanceTimersLspgeneration struct {
	// kubebuilder:validation:Minimum=10
	// kubebuilder:validation:Maximum=100000
	// +kubebuilder:default:=10
	Initialwait *uint64 `json:"initial-wait,omitempty"`
	// kubebuilder:validation:Minimum=10
	// kubebuilder:validation:Maximum=120000
	// +kubebuilder:default:=5000
	Maxwait *uint64 `json:"max-wait,omitempty"`
	// kubebuilder:validation:Minimum=10
	// kubebuilder:validation:Maximum=100000
	// +kubebuilder:default:=1000
	Secondwait *uint64 `json:"second-wait,omitempty"`
}

// DeviceNetworkinstanceProtocolsIsisInstanceTimersLsprefresh struct
type DeviceNetworkinstanceProtocolsIsisInstanceTimersLsprefresh struct {
	// +kubebuilder:default:=true
	Halflifetime *bool `json:"half-lifetime,omitempty"`
	// kubebuilder:validation:Minimum=150
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=600
	Interval *uint16 `json:"interval,omitempty"`
}

// DeviceNetworkinstanceProtocolsIsisInstanceTimersSpf struct
type DeviceNetworkinstanceProtocolsIsisInstanceTimersSpf struct {
	// kubebuilder:validation:Minimum=10
	// kubebuilder:validation:Maximum=100000
	// +kubebuilder:default:=1000
	Initialwait *uint64 `json:"initial-wait,omitempty"`
	// kubebuilder:validation:Minimum=10
	// kubebuilder:validation:Maximum=120000
	// +kubebuilder:default:=10000
	Maxwait *uint64 `json:"max-wait,omitempty"`
	// kubebuilder:validation:Minimum=10
	// kubebuilder:validation:Maximum=100000
	// +kubebuilder:default:=1000
	Secondwait *uint64 `json:"second-wait,omitempty"`
}

// DeviceNetworkinstanceProtocolsIsisInstanceTraceoptions struct
type DeviceNetworkinstanceProtocolsIsisInstanceTraceoptions struct {
	//RootNetworkinstanceProtocolsIsisInstanceTraceoptionsTrace
	Trace *DeviceNetworkinstanceProtocolsIsisInstanceTraceoptionsTrace `json:"trace,omitempty"`
}

// DeviceNetworkinstanceProtocolsIsisInstanceTraceoptionsTrace struct
type DeviceNetworkinstanceProtocolsIsisInstanceTraceoptionsTrace struct {
	// +kubebuilder:validation:Enum=`adjacencies`;`graceful-restart`;`interfaces`;`packets-all`;`packets-l1-csnp`;`packets-l1-hello`;`packets-l1-lsp`;`packets-l1-psnp`;`packets-l2-csnp`;`packets-l2-hello`;`packets-l2-lsp`;`packets-l2-psnp`;`packets-p2p-hello`;`routes`;`summary-addresses`
	Trace E_DeviceNetworkinstanceProtocolsIsisInstanceTraceoptionsTraceTrace `json:"trace,omitempty"`
	//Trace *string `json:"trace,omitempty"`
}

// DeviceNetworkinstanceProtocolsIsisInstanceTrafficengineering struct
type DeviceNetworkinstanceProtocolsIsisInstanceTrafficengineering struct {
	// +kubebuilder:default:=false
	Advertisement *bool `json:"advertisement,omitempty"`
	// +kubebuilder:default:=true
	Legacylinkattributeadvertisement *bool `json:"legacy-link-attribute-advertisement,omitempty"`
}

// DeviceNetworkinstanceProtocolsIsisInstanceTransport struct
type DeviceNetworkinstanceProtocolsIsisInstanceTransport struct {
	// kubebuilder:validation:Minimum=490
	// kubebuilder:validation:Maximum=9490
	// +kubebuilder:default:=1492
	Lspmtusize *uint16 `json:"lsp-mtu-size,omitempty"`
}

// DeviceNetworkinstanceProtocolsLdp struct
type DeviceNetworkinstanceProtocolsLdp struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_DeviceNetworkinstanceProtocolsLdpAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	//RootNetworkinstanceProtocolsLdpDiscovery
	Discovery         *DeviceNetworkinstanceProtocolsLdpDiscovery `json:"discovery,omitempty"`
	Dynamiclabelblock *string                                     `json:"dynamic-label-block"`
	//RootNetworkinstanceProtocolsLdpGracefulrestart
	Gracefulrestart *DeviceNetworkinstanceProtocolsLdpGracefulrestart `json:"graceful-restart,omitempty"`
	//RootNetworkinstanceProtocolsLdpIpv4
	Ipv4 *DeviceNetworkinstanceProtocolsLdpIpv4 `json:"ipv4,omitempty"`
	//RootNetworkinstanceProtocolsLdpMultipath
	Multipath *DeviceNetworkinstanceProtocolsLdpMultipath `json:"multipath,omitempty"`
	//RootNetworkinstanceProtocolsLdpPeers
	Peers *DeviceNetworkinstanceProtocolsLdpPeers `json:"peers,omitempty"`
}

// DeviceNetworkinstanceProtocolsLdpDiscovery struct
type DeviceNetworkinstanceProtocolsLdpDiscovery struct {
	//RootNetworkinstanceProtocolsLdpDiscoveryInterfaces
	Interfaces *DeviceNetworkinstanceProtocolsLdpDiscoveryInterfaces `json:"interfaces,omitempty"`
}

// DeviceNetworkinstanceProtocolsLdpDiscoveryInterfaces struct
type DeviceNetworkinstanceProtocolsLdpDiscoveryInterfaces struct {
	// kubebuilder:validation:Minimum=15
	// kubebuilder:validation:Maximum=3600
	// +kubebuilder:default:=15
	Helloholdtime *uint16 `json:"hello-holdtime,omitempty"`
	// kubebuilder:validation:Minimum=5
	// kubebuilder:validation:Maximum=1200
	// +kubebuilder:default:=5
	Hellointerval *uint16 `json:"hello-interval,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Interface []*DeviceNetworkinstanceProtocolsLdpDiscoveryInterfacesInterface `json:"interface,omitempty"`
	//RootNetworkinstanceProtocolsLdpDiscoveryInterfacesTraceoptions
	Traceoptions *DeviceNetworkinstanceProtocolsLdpDiscoveryInterfacesTraceoptions `json:"trace-options,omitempty"`
}

// DeviceNetworkinstanceProtocolsLdpDiscoveryInterfacesInterface struct
type DeviceNetworkinstanceProtocolsLdpDiscoveryInterfacesInterface struct {
	// kubebuilder:validation:Minimum=15
	// kubebuilder:validation:Maximum=3600
	// +kubebuilder:default:=15
	Helloholdtime *uint16 `json:"hello-holdtime,omitempty"`
	// kubebuilder:validation:Minimum=5
	// kubebuilder:validation:Maximum=1200
	// +kubebuilder:default:=5
	Hellointerval *uint16 `json:"hello-interval,omitempty"`
	//RootNetworkinstanceProtocolsLdpDiscoveryInterfacesInterfaceIpv4
	Ipv4 *DeviceNetworkinstanceProtocolsLdpDiscoveryInterfacesInterfaceIpv4 `json:"ipv4,omitempty"`
	Name *string                                                            `json:"name"`
}

// DeviceNetworkinstanceProtocolsLdpDiscoveryInterfacesInterfaceIpv4 struct
type DeviceNetworkinstanceProtocolsLdpDiscoveryInterfacesInterfaceIpv4 struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_DeviceNetworkinstanceProtocolsLdpDiscoveryInterfacesInterfaceIpv4Adminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	//RootNetworkinstanceProtocolsLdpDiscoveryInterfacesInterfaceIpv4Traceoptions
	Traceoptions *DeviceNetworkinstanceProtocolsLdpDiscoveryInterfacesInterfaceIpv4Traceoptions `json:"trace-options,omitempty"`
}

// DeviceNetworkinstanceProtocolsLdpDiscoveryInterfacesInterfaceIpv4Traceoptions struct
type DeviceNetworkinstanceProtocolsLdpDiscoveryInterfacesInterfaceIpv4Traceoptions struct {
	//RootNetworkinstanceProtocolsLdpDiscoveryInterfacesInterfaceIpv4TraceoptionsTrace
	Trace *DeviceNetworkinstanceProtocolsLdpDiscoveryInterfacesInterfaceIpv4TraceoptionsTrace `json:"trace,omitempty"`
}

// DeviceNetworkinstanceProtocolsLdpDiscoveryInterfacesInterfaceIpv4TraceoptionsTrace struct
type DeviceNetworkinstanceProtocolsLdpDiscoveryInterfacesInterfaceIpv4TraceoptionsTrace struct {
	// +kubebuilder:validation:Enum=`all`;`events-discovery`;`messages-hello`;`messages-hello-detail`
	Trace E_DeviceNetworkinstanceProtocolsLdpDiscoveryInterfacesInterfaceIpv4TraceoptionsTraceTrace `json:"trace,omitempty"`
	//Trace *string `json:"trace,omitempty"`
}

// DeviceNetworkinstanceProtocolsLdpDiscoveryInterfacesTraceoptions struct
type DeviceNetworkinstanceProtocolsLdpDiscoveryInterfacesTraceoptions struct {
	//RootNetworkinstanceProtocolsLdpDiscoveryInterfacesTraceoptionsTrace
	Trace *DeviceNetworkinstanceProtocolsLdpDiscoveryInterfacesTraceoptionsTrace `json:"trace,omitempty"`
}

// DeviceNetworkinstanceProtocolsLdpDiscoveryInterfacesTraceoptionsTrace struct
type DeviceNetworkinstanceProtocolsLdpDiscoveryInterfacesTraceoptionsTrace struct {
	// +kubebuilder:validation:Enum=`all`;`events-discovery`;`messages-hello`;`messages-hello-detail`
	Trace E_DeviceNetworkinstanceProtocolsLdpDiscoveryInterfacesTraceoptionsTraceTrace `json:"trace,omitempty"`
	//Trace *string `json:"trace,omitempty"`
}

// DeviceNetworkinstanceProtocolsLdpGracefulrestart struct
type DeviceNetworkinstanceProtocolsLdpGracefulrestart struct {
	// +kubebuilder:default:=false
	Helperenable *bool `json:"helper-enable,omitempty"`
	// kubebuilder:validation:Minimum=10
	// kubebuilder:validation:Maximum=1800
	// +kubebuilder:default:=120
	Maxreconnecttime *uint16 `json:"max-reconnect-time,omitempty"`
	// kubebuilder:validation:Minimum=30
	// kubebuilder:validation:Maximum=3600
	// +kubebuilder:default:=120
	Maxrecoverytime *uint16 `json:"max-recovery-time,omitempty"`
}

// DeviceNetworkinstanceProtocolsLdpIpv4 struct
type DeviceNetworkinstanceProtocolsLdpIpv4 struct {
	//RootNetworkinstanceProtocolsLdpIpv4Fecresolution
	Fecresolution *DeviceNetworkinstanceProtocolsLdpIpv4Fecresolution `json:"fec-resolution,omitempty"`
}

// DeviceNetworkinstanceProtocolsLdpIpv4Fecresolution struct
type DeviceNetworkinstanceProtocolsLdpIpv4Fecresolution struct {
	// +kubebuilder:default:=false
	Longestprefix *bool `json:"longest-prefix,omitempty"`
}

// DeviceNetworkinstanceProtocolsLdpMultipath struct
type DeviceNetworkinstanceProtocolsLdpMultipath struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=64
	// +kubebuilder:default:=1
	Maxpaths *uint8 `json:"max-paths,omitempty"`
}

// DeviceNetworkinstanceProtocolsLdpPeers struct
type DeviceNetworkinstanceProtocolsLdpPeers struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Peer []*DeviceNetworkinstanceProtocolsLdpPeersPeer `json:"peer,omitempty"`
	// kubebuilder:validation:Minimum=45
	// kubebuilder:validation:Maximum=3600
	// +kubebuilder:default:=180
	Sessionkeepaliveholdtime *uint16 `json:"session-keepalive-holdtime,omitempty"`
	// kubebuilder:validation:Minimum=15
	// kubebuilder:validation:Maximum=1200
	// +kubebuilder:default:=60
	Sessionkeepaliveinterval *uint16 `json:"session-keepalive-interval,omitempty"`
	//RootNetworkinstanceProtocolsLdpPeersTraceoptions
	Traceoptions *DeviceNetworkinstanceProtocolsLdpPeersTraceoptions `json:"trace-options,omitempty"`
}

// DeviceNetworkinstanceProtocolsLdpPeersPeer struct
type DeviceNetworkinstanceProtocolsLdpPeersPeer struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	// +kubebuilder:default:=0
	Feclimit *uint32 `json:"fec-limit,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	Labelspaceid *uint16 `json:"label-space-id"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Lsrid *string `json:"lsr-id"`
	//RootNetworkinstanceProtocolsLdpPeersPeerTcptransport
	Tcptransport *DeviceNetworkinstanceProtocolsLdpPeersPeerTcptransport `json:"tcp-transport,omitempty"`
	//RootNetworkinstanceProtocolsLdpPeersPeerTraceoptions
	Traceoptions *DeviceNetworkinstanceProtocolsLdpPeersPeerTraceoptions `json:"trace-options,omitempty"`
}

// DeviceNetworkinstanceProtocolsLdpPeersPeerTcptransport struct
type DeviceNetworkinstanceProtocolsLdpPeersPeerTcptransport struct {
}

// DeviceNetworkinstanceProtocolsLdpPeersPeerTraceoptions struct
type DeviceNetworkinstanceProtocolsLdpPeersPeerTraceoptions struct {
	//RootNetworkinstanceProtocolsLdpPeersPeerTraceoptionsTrace
	Trace *DeviceNetworkinstanceProtocolsLdpPeersPeerTraceoptionsTrace `json:"trace,omitempty"`
}

// DeviceNetworkinstanceProtocolsLdpPeersPeerTraceoptionsTrace struct
type DeviceNetworkinstanceProtocolsLdpPeersPeerTraceoptionsTrace struct {
	// +kubebuilder:validation:Enum=`all`;`events-all`;`events-binding`;`events-session`;`messages-all`;`messages-all-detail`;`messages-initialization`;`messages-initialization-detail`;`messages-keepalive`;`messages-label`;`messages-label-detail`
	Trace E_DeviceNetworkinstanceProtocolsLdpPeersPeerTraceoptionsTraceTrace `json:"trace,omitempty"`
	//Trace *string `json:"trace,omitempty"`
}

// DeviceNetworkinstanceProtocolsLdpPeersTraceoptions struct
type DeviceNetworkinstanceProtocolsLdpPeersTraceoptions struct {
	//RootNetworkinstanceProtocolsLdpPeersTraceoptionsTrace
	Trace *DeviceNetworkinstanceProtocolsLdpPeersTraceoptionsTrace `json:"trace,omitempty"`
}

// DeviceNetworkinstanceProtocolsLdpPeersTraceoptionsTrace struct
type DeviceNetworkinstanceProtocolsLdpPeersTraceoptionsTrace struct {
	// +kubebuilder:validation:Enum=`all`;`events-all`;`events-binding`;`events-session`;`messages-all`;`messages-all-detail`;`messages-initialization`;`messages-initialization-detail`;`messages-keepalive`;`messages-label`;`messages-label-detail`
	Trace E_DeviceNetworkinstanceProtocolsLdpPeersTraceoptionsTraceTrace `json:"trace,omitempty"`
	//Trace *string `json:"trace,omitempty"`
}

// DeviceNetworkinstanceProtocolsLinux struct
type DeviceNetworkinstanceProtocolsLinux struct {
	// +kubebuilder:default:=true
	Exportneighbors *bool `json:"export-neighbors,omitempty"`
	// +kubebuilder:default:=false
	Exportroutes *bool `json:"export-routes,omitempty"`
	// +kubebuilder:default:=false
	Importroutes *bool `json:"import-routes,omitempty"`
}

// DeviceNetworkinstanceProtocolsMld struct
type DeviceNetworkinstanceProtocolsMld struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_DeviceNetworkinstanceProtocolsMldAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Interface []*DeviceNetworkinstanceProtocolsMldInterface `json:"interface,omitempty"`
	// kubebuilder:validation:Minimum=2
	// kubebuilder:validation:Maximum=1024
	// +kubebuilder:default:=125
	Queryinterval *uint32 `json:"query-interval,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=1023
	// +kubebuilder:default:=1
	Querylastmemberinterval *uint32 `json:"query-last-member-interval,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=1023
	// +kubebuilder:default:=10
	Queryresponseinterval *uint32 `json:"query-response-interval,omitempty"`
	// kubebuilder:validation:Minimum=2
	// kubebuilder:validation:Maximum=10
	// +kubebuilder:default:=2
	Robustcount *uint32 `json:"robust-count,omitempty"`
	//RootNetworkinstanceProtocolsMldSsm
	Ssm *DeviceNetworkinstanceProtocolsMldSsm `json:"ssm,omitempty"`
	//RootNetworkinstanceProtocolsMldTraceoptions
	Traceoptions *DeviceNetworkinstanceProtocolsMldTraceoptions `json:"trace-options,omitempty"`
}

// DeviceNetworkinstanceProtocolsMldInterface struct
type DeviceNetworkinstanceProtocolsMldInterface struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_DeviceNetworkinstanceProtocolsMldInterfaceAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	Importpolicy  *string `json:"import-policy,omitempty"`
	Interfacename *string `json:"interface-name"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4096
	// +kubebuilder:default:=0
	Maxgroupsources *uint32 `json:"max-group-sources,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4096
	// +kubebuilder:default:=0
	Maxgroups *uint32 `json:"max-groups,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=512
	// +kubebuilder:default:=0
	Maxsources *uint32 `json:"max-sources,omitempty"`
	// kubebuilder:validation:Minimum=2
	// kubebuilder:validation:Maximum=1024
	Queryinterval *uint32 `json:"query-interval,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=1023
	Querylastmemberinterval *uint32 `json:"query-last-member-interval,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=1023
	Queryresponseinterval *uint32 `json:"query-response-interval,omitempty"`
	// +kubebuilder:default:=true
	Routeralertcheck *bool `json:"router-alert-check,omitempty"`
	//RootNetworkinstanceProtocolsMldInterfaceSsm
	Ssm *DeviceNetworkinstanceProtocolsMldInterfaceSsm `json:"ssm,omitempty"`
	//RootNetworkinstanceProtocolsMldInterfaceStatic
	Static *DeviceNetworkinstanceProtocolsMldInterfaceStatic `json:"static,omitempty"`
	//RootNetworkinstanceProtocolsMldInterfaceTraceoptions
	Traceoptions *DeviceNetworkinstanceProtocolsMldInterfaceTraceoptions `json:"trace-options,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=2
	// +kubebuilder:default:=2
	Version *uint8 `json:"version,omitempty"`
}

// DeviceNetworkinstanceProtocolsMldInterfaceSsm struct
type DeviceNetworkinstanceProtocolsMldInterfaceSsm struct {
	//RootNetworkinstanceProtocolsMldInterfaceSsmMappings
	Mappings *DeviceNetworkinstanceProtocolsMldInterfaceSsmMappings `json:"mappings,omitempty"`
}

// DeviceNetworkinstanceProtocolsMldInterfaceSsmMappings struct
type DeviceNetworkinstanceProtocolsMldInterfaceSsmMappings struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Grouprange []*DeviceNetworkinstanceProtocolsMldInterfaceSsmMappingsGrouprange `json:"group-range,omitempty"`
}

// DeviceNetworkinstanceProtocolsMldInterfaceSsmMappingsGrouprange struct
type DeviceNetworkinstanceProtocolsMldInterfaceSsmMappingsGrouprange struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	End *string `json:"end"`
	//+kubebuilder:validation:MinItems=1
	//+kubebuilder:validation:MaxItems=1024
	Source []*DeviceNetworkinstanceProtocolsMldInterfaceSsmMappingsGrouprangeSource `json:"source,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Start *string `json:"start"`
}

// DeviceNetworkinstanceProtocolsMldInterfaceSsmMappingsGrouprangeSource struct
type DeviceNetworkinstanceProtocolsMldInterfaceSsmMappingsGrouprangeSource struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Source *string `json:"source"`
}

// DeviceNetworkinstanceProtocolsMldInterfaceStatic struct
type DeviceNetworkinstanceProtocolsMldInterfaceStatic struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Grouprange []*DeviceNetworkinstanceProtocolsMldInterfaceStaticGrouprange `json:"group-range,omitempty"`
}

// DeviceNetworkinstanceProtocolsMldInterfaceStaticGrouprange struct
type DeviceNetworkinstanceProtocolsMldInterfaceStaticGrouprange struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	End *string `json:"end"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Source []*DeviceNetworkinstanceProtocolsMldInterfaceStaticGrouprangeSource `json:"source,omitempty"`
	Starg  *string                                                             `json:"starg,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Start *string `json:"start"`
}

// DeviceNetworkinstanceProtocolsMldInterfaceStaticGrouprangeSource struct
type DeviceNetworkinstanceProtocolsMldInterfaceStaticGrouprangeSource struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Source *string `json:"source"`
}

// DeviceNetworkinstanceProtocolsMldInterfaceTraceoptions struct
type DeviceNetworkinstanceProtocolsMldInterfaceTraceoptions struct {
	//RootNetworkinstanceProtocolsMldInterfaceTraceoptionsTrace
	Trace *DeviceNetworkinstanceProtocolsMldInterfaceTraceoptionsTrace `json:"trace,omitempty"`
}

// DeviceNetworkinstanceProtocolsMldInterfaceTraceoptionsTrace struct
type DeviceNetworkinstanceProtocolsMldInterfaceTraceoptionsTrace struct {
	Interfaces *string `json:"interfaces,omitempty"`
	//RootNetworkinstanceProtocolsMldInterfaceTraceoptionsTracePacket
	Packet *DeviceNetworkinstanceProtocolsMldInterfaceTraceoptionsTracePacket `json:"packet,omitempty"`
}

// DeviceNetworkinstanceProtocolsMldInterfaceTraceoptionsTracePacket struct
type DeviceNetworkinstanceProtocolsMldInterfaceTraceoptionsTracePacket struct {
	// +kubebuilder:validation:Enum=`dropped`;`egress-ingress-and-dropped`;`ingress-and-dropped`
	Modifier E_DeviceNetworkinstanceProtocolsMldInterfaceTraceoptionsTracePacketModifier `json:"modifier,omitempty"`
	//Modifier *string `json:"modifier,omitempty"`
	// +kubebuilder:validation:Enum=`all`;`query`;`v1-done`;`v1-report`;`v2-report`
	Type E_DeviceNetworkinstanceProtocolsMldInterfaceTraceoptionsTracePacketType `json:"type,omitempty"`
	//Type *string `json:"type,omitempty"`
}

// DeviceNetworkinstanceProtocolsMldSsm struct
type DeviceNetworkinstanceProtocolsMldSsm struct {
	//RootNetworkinstanceProtocolsMldSsmMappings
	Mappings *DeviceNetworkinstanceProtocolsMldSsmMappings `json:"mappings,omitempty"`
}

// DeviceNetworkinstanceProtocolsMldSsmMappings struct
type DeviceNetworkinstanceProtocolsMldSsmMappings struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Grouprange []*DeviceNetworkinstanceProtocolsMldSsmMappingsGrouprange `json:"group-range,omitempty"`
}

// DeviceNetworkinstanceProtocolsMldSsmMappingsGrouprange struct
type DeviceNetworkinstanceProtocolsMldSsmMappingsGrouprange struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	End *string `json:"end"`
	//+kubebuilder:validation:MinItems=1
	//+kubebuilder:validation:MaxItems=1024
	Source []*DeviceNetworkinstanceProtocolsMldSsmMappingsGrouprangeSource `json:"source,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Start *string `json:"start"`
}

// DeviceNetworkinstanceProtocolsMldSsmMappingsGrouprangeSource struct
type DeviceNetworkinstanceProtocolsMldSsmMappingsGrouprangeSource struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Source *string `json:"source"`
}

// DeviceNetworkinstanceProtocolsMldTraceoptions struct
type DeviceNetworkinstanceProtocolsMldTraceoptions struct {
	//RootNetworkinstanceProtocolsMldTraceoptionsTrace
	Trace *DeviceNetworkinstanceProtocolsMldTraceoptionsTrace `json:"trace,omitempty"`
}

// DeviceNetworkinstanceProtocolsMldTraceoptionsTrace struct
type DeviceNetworkinstanceProtocolsMldTraceoptionsTrace struct {
	Interfaces *string `json:"interfaces,omitempty"`
	//RootNetworkinstanceProtocolsMldTraceoptionsTracePacket
	Packet *DeviceNetworkinstanceProtocolsMldTraceoptionsTracePacket `json:"packet,omitempty"`
}

// DeviceNetworkinstanceProtocolsMldTraceoptionsTracePacket struct
type DeviceNetworkinstanceProtocolsMldTraceoptionsTracePacket struct {
	// +kubebuilder:validation:Enum=`dropped`;`egress-ingress-and-dropped`;`ingress-and-dropped`
	Modifier E_DeviceNetworkinstanceProtocolsMldTraceoptionsTracePacketModifier `json:"modifier,omitempty"`
	//Modifier *string `json:"modifier,omitempty"`
	// +kubebuilder:validation:Enum=`all`;`query`;`v1-done`;`v1-report`;`v2-report`
	Type E_DeviceNetworkinstanceProtocolsMldTraceoptionsTracePacketType `json:"type,omitempty"`
	//Type *string `json:"type,omitempty"`
}

// DeviceNetworkinstanceProtocolsOspf struct
type DeviceNetworkinstanceProtocolsOspf struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=3
	Instance []*DeviceNetworkinstanceProtocolsOspfInstance `json:"instance,omitempty"`
}

// DeviceNetworkinstanceProtocolsOspfInstance struct
type DeviceNetworkinstanceProtocolsOspfInstance struct {
	Addressfamily *string `json:"address-family,omitempty"`
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_DeviceNetworkinstanceProtocolsOspfInstanceAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// +kubebuilder:validation:Enum=`area`;`as`;`false`;`link`
	Advertiseroutercapability E_DeviceNetworkinstanceProtocolsOspfInstanceAdvertiseroutercapability `json:"advertise-router-capability,omitempty"`
	//Advertiseroutercapability *string `json:"advertise-router-capability,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Area []*DeviceNetworkinstanceProtocolsOspfInstanceArea `json:"area,omitempty"`
	//RootNetworkinstanceProtocolsOspfInstanceAsbr
	Asbr *DeviceNetworkinstanceProtocolsOspfInstanceAsbr `json:"asbr,omitempty"`
	//RootNetworkinstanceProtocolsOspfInstanceExportlimit
	Exportlimit  *DeviceNetworkinstanceProtocolsOspfInstanceExportlimit `json:"export-limit,omitempty"`
	Exportpolicy *string                                                `json:"export-policy,omitempty"`
	//RootNetworkinstanceProtocolsOspfInstanceExternaldboverflow
	Externaldboverflow *DeviceNetworkinstanceProtocolsOspfInstanceExternaldboverflow `json:"external-db-overflow,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=150
	Externalpreference *uint8 `json:"external-preference,omitempty"`
	//RootNetworkinstanceProtocolsOspfInstanceGracefulrestart
	Gracefulrestart *DeviceNetworkinstanceProtocolsOspfInstanceGracefulrestart `json:"graceful-restart,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	Instanceid *uint32 `json:"instance-id,omitempty"`
	//RootNetworkinstanceProtocolsOspfInstanceLdpsynchronization
	Ldpsynchronization *DeviceNetworkinstanceProtocolsOspfInstanceLdpsynchronization `json:"ldp-synchronization,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=64
	// +kubebuilder:default:=1
	Maxecmppaths *uint8 `json:"max-ecmp-paths,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
	//RootNetworkinstanceProtocolsOspfInstanceOverload
	Overload *DeviceNetworkinstanceProtocolsOspfInstanceOverload `json:"overload,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=10
	Preference *uint8 `json:"preference,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=8000000000
	// +kubebuilder:default:=400000000
	Referencebandwidth *uint64 `json:"reference-bandwidth,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Routerid *string `json:"router-id,omitempty"`
	//RootNetworkinstanceProtocolsOspfInstanceTedatabaseinstall
	Tedatabaseinstall *DeviceNetworkinstanceProtocolsOspfInstanceTedatabaseinstall `json:"te-database-install,omitempty"`
	//RootNetworkinstanceProtocolsOspfInstanceTimers
	Timers *DeviceNetworkinstanceProtocolsOspfInstanceTimers `json:"timers,omitempty"`
	//RootNetworkinstanceProtocolsOspfInstanceTraceoptions
	Traceoptions *DeviceNetworkinstanceProtocolsOspfInstanceTraceoptions `json:"trace-options,omitempty"`
	//RootNetworkinstanceProtocolsOspfInstanceTrafficengineering
	Trafficengineering *DeviceNetworkinstanceProtocolsOspfInstanceTrafficengineering `json:"traffic-engineering,omitempty"`
	Version            *string                                                       `json:"version"`
}

// DeviceNetworkinstanceProtocolsOspfInstanceArea struct
type DeviceNetworkinstanceProtocolsOspfInstanceArea struct {
	// +kubebuilder:default:=true
	Advertiseroutercapability *bool `json:"advertise-router-capability,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|[0-9\.]*|(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])([\p{N}\p{L}]+)?`
	Areaid *string `json:"area-id"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Arearange []*DeviceNetworkinstanceProtocolsOspfInstanceAreaArearange `json:"area-range,omitempty"`
	// +kubebuilder:default:=false
	Bgplsexclude *bool `json:"bgp-ls-exclude,omitempty"`
	// +kubebuilder:default:=true
	Blackholeaggregate *bool   `json:"blackhole-aggregate,omitempty"`
	Exportpolicy       *string `json:"export-policy,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Interface []*DeviceNetworkinstanceProtocolsOspfInstanceAreaInterface `json:"interface,omitempty"`
	//RootNetworkinstanceProtocolsOspfInstanceAreaNssa
	Nssa *DeviceNetworkinstanceProtocolsOspfInstanceAreaNssa `json:"nssa,omitempty"`
	//RootNetworkinstanceProtocolsOspfInstanceAreaStub
	Stub *DeviceNetworkinstanceProtocolsOspfInstanceAreaStub `json:"stub,omitempty"`
}

// DeviceNetworkinstanceProtocolsOspfInstanceAreaArearange struct
type DeviceNetworkinstanceProtocolsOspfInstanceAreaArearange struct {
	// +kubebuilder:default:=true
	Advertise *bool `json:"advertise,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	Ipprefixmask *string `json:"ip-prefix-mask"`
}

// DeviceNetworkinstanceProtocolsOspfInstanceAreaInterface struct
type DeviceNetworkinstanceProtocolsOspfInstanceAreaInterface struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_DeviceNetworkinstanceProtocolsOspfInstanceAreaInterfaceAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// +kubebuilder:default:=true
	Advertiseroutercapability *bool `json:"advertise-router-capability,omitempty"`
	// +kubebuilder:default:=true
	Advertisesubnet *bool `json:"advertise-subnet,omitempty"`
	//RootNetworkinstanceProtocolsOspfInstanceAreaInterfaceAuthentication
	Authentication *DeviceNetworkinstanceProtocolsOspfInstanceAreaInterfaceAuthentication `json:"authentication,omitempty"`
	// kubebuilder:validation:Minimum=2
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=40
	Deadinterval *uint32 `json:"dead-interval,omitempty"`
	//RootNetworkinstanceProtocolsOspfInstanceAreaInterfaceFailuredetection
	Failuredetection *DeviceNetworkinstanceProtocolsOspfInstanceAreaInterfaceFailuredetection `json:"failure-detection,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=10
	Hellointerval *uint32 `json:"hello-interval,omitempty"`
	Interfacename *string `json:"interface-name"`
	// +kubebuilder:validation:Enum=`broadcast`;`point-to-point`
	Interfacetype E_DeviceNetworkinstanceProtocolsOspfInstanceAreaInterfaceInterfacetype `json:"interface-type,omitempty"`
	//Interfacetype *string `json:"interface-type,omitempty"`
	//RootNetworkinstanceProtocolsOspfInstanceAreaInterfaceLdpsynchronization
	Ldpsynchronization *DeviceNetworkinstanceProtocolsOspfInstanceAreaInterfaceLdpsynchronization `json:"ldp-synchronization,omitempty"`
	// +kubebuilder:validation:Enum=`all`;`except-own-rtrlsa`;`except-own-rtrlsa-and-defaults`;`none`
	// +kubebuilder:default:="none"
	Lsafilterout E_DeviceNetworkinstanceProtocolsOspfInstanceAreaInterfaceLsafilterout `json:"lsa-filter-out,omitempty"`
	//Lsafilterout *string `json:"lsa-filter-out,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	Metric *uint16 `json:"metric,omitempty"`
	// kubebuilder:validation:Minimum=512
	// kubebuilder:validation:Maximum=9486
	Mtu     *uint32 `json:"mtu,omitempty"`
	Passive *bool   `json:"passive,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=1
	Priority *uint16 `json:"priority,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=1800
	// +kubebuilder:default:=5
	Retransmitinterval *uint32 `json:"retransmit-interval,omitempty"`
	//RootNetworkinstanceProtocolsOspfInstanceAreaInterfaceTraceoptions
	Traceoptions *DeviceNetworkinstanceProtocolsOspfInstanceAreaInterfaceTraceoptions `json:"trace-options,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=1800
	// +kubebuilder:default:=1
	Transitdelay *uint32 `json:"transit-delay,omitempty"`
}

// DeviceNetworkinstanceProtocolsOspfInstanceAreaInterfaceAuthentication struct
type DeviceNetworkinstanceProtocolsOspfInstanceAreaInterfaceAuthentication struct {
	Keychain *string `json:"keychain,omitempty"`
}

// DeviceNetworkinstanceProtocolsOspfInstanceAreaInterfaceFailuredetection struct
type DeviceNetworkinstanceProtocolsOspfInstanceAreaInterfaceFailuredetection struct {
	// +kubebuilder:default:=false
	Enablebfd *bool `json:"enable-bfd,omitempty"`
}

// DeviceNetworkinstanceProtocolsOspfInstanceAreaInterfaceLdpsynchronization struct
type DeviceNetworkinstanceProtocolsOspfInstanceAreaInterfaceLdpsynchronization struct {
	Disable  *string `json:"disable,omitempty"`
	Endoflib *bool   `json:"end-of-lib,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=1800
	Holddowntimer *uint16 `json:"hold-down-timer,omitempty"`
}

// DeviceNetworkinstanceProtocolsOspfInstanceAreaInterfaceTraceoptions struct
type DeviceNetworkinstanceProtocolsOspfInstanceAreaInterfaceTraceoptions struct {
	//RootNetworkinstanceProtocolsOspfInstanceAreaInterfaceTraceoptionsTrace
	Trace *DeviceNetworkinstanceProtocolsOspfInstanceAreaInterfaceTraceoptionsTrace `json:"trace,omitempty"`
}

// DeviceNetworkinstanceProtocolsOspfInstanceAreaInterfaceTraceoptionsTrace struct
type DeviceNetworkinstanceProtocolsOspfInstanceAreaInterfaceTraceoptionsTrace struct {
	Adjacencies *string `json:"adjacencies,omitempty"`
	Interfaces  *string `json:"interfaces,omitempty"`
	//RootNetworkinstanceProtocolsOspfInstanceAreaInterfaceTraceoptionsTracePacket
	Packet *DeviceNetworkinstanceProtocolsOspfInstanceAreaInterfaceTraceoptionsTracePacket `json:"packet,omitempty"`
}

// DeviceNetworkinstanceProtocolsOspfInstanceAreaInterfaceTraceoptionsTracePacket struct
type DeviceNetworkinstanceProtocolsOspfInstanceAreaInterfaceTraceoptionsTracePacket struct {
	Detail *string `json:"detail,omitempty"`
	// +kubebuilder:validation:Enum=`drop`;`egress`;`in-and-egress`;`ingress`
	Modifier E_DeviceNetworkinstanceProtocolsOspfInstanceAreaInterfaceTraceoptionsTracePacketModifier `json:"modifier,omitempty"`
	//Modifier *string `json:"modifier,omitempty"`
	// +kubebuilder:validation:Enum=`all`;`dbdescr`;`hello`;`ls-ack`;`ls-request`;`ls-update`
	Type E_DeviceNetworkinstanceProtocolsOspfInstanceAreaInterfaceTraceoptionsTracePacketType `json:"type,omitempty"`
	//Type *string `json:"type,omitempty"`
}

// DeviceNetworkinstanceProtocolsOspfInstanceAreaNssa struct
type DeviceNetworkinstanceProtocolsOspfInstanceAreaNssa struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Arearange []*DeviceNetworkinstanceProtocolsOspfInstanceAreaNssaArearange `json:"area-range,omitempty"`
	//RootNetworkinstanceProtocolsOspfInstanceAreaNssaOriginatedefaultroute
	Originatedefaultroute *DeviceNetworkinstanceProtocolsOspfInstanceAreaNssaOriginatedefaultroute `json:"originate-default-route,omitempty"`
	// +kubebuilder:default:=true
	Redistributeexternal *bool `json:"redistribute-external,omitempty"`
	// +kubebuilder:default:=true
	Summaries *bool `json:"summaries,omitempty"`
}

// DeviceNetworkinstanceProtocolsOspfInstanceAreaNssaArearange struct
type DeviceNetworkinstanceProtocolsOspfInstanceAreaNssaArearange struct {
	// +kubebuilder:default:=true
	Advertise *bool `json:"advertise,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	Ipprefixmask *string `json:"ip-prefix-mask"`
}

// DeviceNetworkinstanceProtocolsOspfInstanceAreaNssaOriginatedefaultroute struct
type DeviceNetworkinstanceProtocolsOspfInstanceAreaNssaOriginatedefaultroute struct {
	// +kubebuilder:default:=true
	Adjacencycheck *bool `json:"adjacency-check,omitempty"`
	// +kubebuilder:default:=false
	Typenssa *bool `json:"type-nssa,omitempty"`
}

// DeviceNetworkinstanceProtocolsOspfInstanceAreaStub struct
type DeviceNetworkinstanceProtocolsOspfInstanceAreaStub struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=1
	Defaultmetric *uint16 `json:"default-metric,omitempty"`
	// +kubebuilder:default:=true
	Summaries *bool `json:"summaries,omitempty"`
}

// DeviceNetworkinstanceProtocolsOspfInstanceAsbr struct
type DeviceNetworkinstanceProtocolsOspfInstanceAsbr struct {
	// +kubebuilder:default:="none"
	Tracepath *string `json:"trace-path,omitempty"`
}

// DeviceNetworkinstanceProtocolsOspfInstanceExportlimit struct
type DeviceNetworkinstanceProtocolsOspfInstanceExportlimit struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=100
	Logpercent *uint32 `json:"log-percent,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	Number *uint32 `json:"number"`
}

// DeviceNetworkinstanceProtocolsOspfInstanceExternaldboverflow struct
type DeviceNetworkinstanceProtocolsOspfInstanceExternaldboverflow struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=2147483647
	// +kubebuilder:default:=0
	Interval *uint32 `json:"interval,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=2147483647
	// +kubebuilder:default:=0
	Limit *uint32 `json:"limit,omitempty"`
}

// DeviceNetworkinstanceProtocolsOspfInstanceGracefulrestart struct
type DeviceNetworkinstanceProtocolsOspfInstanceGracefulrestart struct {
	// +kubebuilder:default:=false
	Helpermode *bool `json:"helper-mode,omitempty"`
	// +kubebuilder:default:=false
	Strictlsachecking *bool `json:"strict-lsa-checking,omitempty"`
}

// DeviceNetworkinstanceProtocolsOspfInstanceLdpsynchronization struct
type DeviceNetworkinstanceProtocolsOspfInstanceLdpsynchronization struct {
	// +kubebuilder:default:=false
	Endoflib *bool `json:"end-of-lib,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=1800
	// +kubebuilder:default:=60
	Holddowntimer *uint16 `json:"hold-down-timer,omitempty"`
}

// DeviceNetworkinstanceProtocolsOspfInstanceOverload struct
type DeviceNetworkinstanceProtocolsOspfInstanceOverload struct {
	// +kubebuilder:default:=false
	Active *bool `json:"active,omitempty"`
	// +kubebuilder:default:=false
	Overloadincludeext1 *bool `json:"overload-include-ext-1,omitempty"`
	// +kubebuilder:default:=false
	Overloadincludeext2 *bool `json:"overload-include-ext-2,omitempty"`
	// +kubebuilder:default:=false
	Overloadincludestub *bool `json:"overload-include-stub,omitempty"`
	//RootNetworkinstanceProtocolsOspfInstanceOverloadOverloadonboot
	Overloadonboot *DeviceNetworkinstanceProtocolsOspfInstanceOverloadOverloadonboot `json:"overload-on-boot,omitempty"`
	//RootNetworkinstanceProtocolsOspfInstanceOverloadRtradvlsalimit
	Rtradvlsalimit *DeviceNetworkinstanceProtocolsOspfInstanceOverloadRtradvlsalimit `json:"rtr-adv-lsa-limit,omitempty"`
}

// DeviceNetworkinstanceProtocolsOspfInstanceOverloadOverloadonboot struct
type DeviceNetworkinstanceProtocolsOspfInstanceOverloadOverloadonboot struct {
	// kubebuilder:validation:Minimum=60
	// kubebuilder:validation:Maximum=1800
	// +kubebuilder:default:=60
	Timeout *uint32 `json:"timeout,omitempty"`
}

// DeviceNetworkinstanceProtocolsOspfInstanceOverloadRtradvlsalimit struct
type DeviceNetworkinstanceProtocolsOspfInstanceOverloadRtradvlsalimit struct {
	Logonly *bool `json:"log-only,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	Maxlsacount *uint32 `json:"max-lsa-count,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=1800
	Overloadtimeout *uint16 `json:"overload-timeout,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=0
	Warningthreshold *uint8 `json:"warning-threshold,omitempty"`
}

// DeviceNetworkinstanceProtocolsOspfInstanceTedatabaseinstall struct
type DeviceNetworkinstanceProtocolsOspfInstanceTedatabaseinstall struct {
	//RootNetworkinstanceProtocolsOspfInstanceTedatabaseinstallBgpls
	Bgpls *DeviceNetworkinstanceProtocolsOspfInstanceTedatabaseinstallBgpls `json:"bgp-ls,omitempty"`
}

// DeviceNetworkinstanceProtocolsOspfInstanceTedatabaseinstallBgpls struct
type DeviceNetworkinstanceProtocolsOspfInstanceTedatabaseinstallBgpls struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	Bgplsidentifier *uint32 `json:"bgp-ls-identifier,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=-1
	Igpidentifier *uint64 `json:"igp-identifier,omitempty"`
}

// DeviceNetworkinstanceProtocolsOspfInstanceTimers struct
type DeviceNetworkinstanceProtocolsOspfInstanceTimers struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=1000
	// +kubebuilder:default:=1000
	Incrementalspfwait *uint32 `json:"incremental-spf-wait,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=1000
	// +kubebuilder:default:=1000
	Lsaaccumulate *uint32 `json:"lsa-accumulate,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=600000
	// +kubebuilder:default:=1000
	Lsaarrival *uint32 `json:"lsa-arrival,omitempty"`
	//RootNetworkinstanceProtocolsOspfInstanceTimersLsagenerate
	Lsagenerate *DeviceNetworkinstanceProtocolsOspfInstanceTimersLsagenerate `json:"lsa-generate,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=1000
	// +kubebuilder:default:=1000
	Redistributedelay *uint32 `json:"redistribute-delay,omitempty"`
	//RootNetworkinstanceProtocolsOspfInstanceTimersSpfwait
	Spfwait *DeviceNetworkinstanceProtocolsOspfInstanceTimersSpfwait `json:"spf-wait,omitempty"`
}

// DeviceNetworkinstanceProtocolsOspfInstanceTimersLsagenerate struct
type DeviceNetworkinstanceProtocolsOspfInstanceTimersLsagenerate struct {
	// kubebuilder:validation:Minimum=10
	// kubebuilder:validation:Maximum=600000
	// +kubebuilder:default:=5000
	Lsainitialwait *uint32 `json:"lsa-initial-wait,omitempty"`
	// kubebuilder:validation:Minimum=10
	// kubebuilder:validation:Maximum=600000
	// +kubebuilder:default:=5000
	Lsasecondwait *uint32 `json:"lsa-second-wait,omitempty"`
	// kubebuilder:validation:Minimum=10
	// kubebuilder:validation:Maximum=600000
	// +kubebuilder:default:=5000
	Maxlsawait *uint32 `json:"max-lsa-wait,omitempty"`
}

// DeviceNetworkinstanceProtocolsOspfInstanceTimersSpfwait struct
type DeviceNetworkinstanceProtocolsOspfInstanceTimersSpfwait struct {
	// kubebuilder:validation:Minimum=10
	// kubebuilder:validation:Maximum=100000
	// +kubebuilder:default:=1000
	Spfinitialwait *uint32 `json:"spf-initial-wait,omitempty"`
	// kubebuilder:validation:Minimum=10
	// kubebuilder:validation:Maximum=120000
	// +kubebuilder:default:=10000
	Spfmaxwait *uint32 `json:"spf-max-wait,omitempty"`
	// kubebuilder:validation:Minimum=10
	// kubebuilder:validation:Maximum=100000
	// +kubebuilder:default:=1000
	Spfsecondwait *uint32 `json:"spf-second-wait,omitempty"`
}

// DeviceNetworkinstanceProtocolsOspfInstanceTraceoptions struct
type DeviceNetworkinstanceProtocolsOspfInstanceTraceoptions struct {
	//RootNetworkinstanceProtocolsOspfInstanceTraceoptionsTrace
	Trace *DeviceNetworkinstanceProtocolsOspfInstanceTraceoptionsTrace `json:"trace,omitempty"`
}

// DeviceNetworkinstanceProtocolsOspfInstanceTraceoptionsTrace struct
type DeviceNetworkinstanceProtocolsOspfInstanceTraceoptionsTrace struct {
	Adjacencies     *string `json:"adjacencies,omitempty"`
	Gracefulrestart *string `json:"graceful-restart,omitempty"`
	Interfaces      *string `json:"interfaces,omitempty"`
	//RootNetworkinstanceProtocolsOspfInstanceTraceoptionsTraceLsdb
	Lsdb *DeviceNetworkinstanceProtocolsOspfInstanceTraceoptionsTraceLsdb `json:"lsdb,omitempty"`
	Misc *string                                                          `json:"misc,omitempty"`
	//RootNetworkinstanceProtocolsOspfInstanceTraceoptionsTracePacket
	Packet *DeviceNetworkinstanceProtocolsOspfInstanceTraceoptionsTracePacket `json:"packet,omitempty"`
	//RootNetworkinstanceProtocolsOspfInstanceTraceoptionsTraceRoutes
	Routes *DeviceNetworkinstanceProtocolsOspfInstanceTraceoptionsTraceRoutes `json:"routes,omitempty"`
	//RootNetworkinstanceProtocolsOspfInstanceTraceoptionsTraceSpf
	Spf *DeviceNetworkinstanceProtocolsOspfInstanceTraceoptionsTraceSpf `json:"spf,omitempty"`
}

// DeviceNetworkinstanceProtocolsOspfInstanceTraceoptionsTraceLsdb struct
type DeviceNetworkinstanceProtocolsOspfInstanceTraceoptionsTraceLsdb struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Linkstateid *string `json:"link-state-id,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Routerid *string `json:"router-id,omitempty"`
	// +kubebuilder:validation:Enum=`all`;`external`;`inter-area-prefix`;`inter-area-router`;`intra-area-prefix`;`network`;`nssa`;`opaque`;`router`;`summary`
	Type E_DeviceNetworkinstanceProtocolsOspfInstanceTraceoptionsTraceLsdbType `json:"type,omitempty"`
	//Type *string `json:"type,omitempty"`
}

// DeviceNetworkinstanceProtocolsOspfInstanceTraceoptionsTracePacket struct
type DeviceNetworkinstanceProtocolsOspfInstanceTraceoptionsTracePacket struct {
	Detail *string `json:"detail,omitempty"`
	// +kubebuilder:validation:Enum=`drop`;`egress`;`in-and-egress`;`ingress`
	Modifier E_DeviceNetworkinstanceProtocolsOspfInstanceTraceoptionsTracePacketModifier `json:"modifier,omitempty"`
	//Modifier *string `json:"modifier,omitempty"`
	// +kubebuilder:validation:Enum=`all`;`dbdescr`;`hello`;`ls-ack`;`ls-request`;`ls-update`
	Type E_DeviceNetworkinstanceProtocolsOspfInstanceTraceoptionsTracePacketType `json:"type,omitempty"`
	//Type *string `json:"type,omitempty"`
}

// DeviceNetworkinstanceProtocolsOspfInstanceTraceoptionsTraceRoutes struct
type DeviceNetworkinstanceProtocolsOspfInstanceTraceoptionsTraceRoutes struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Destaddress *string `json:"dest-address,omitempty"`
}

// DeviceNetworkinstanceProtocolsOspfInstanceTraceoptionsTraceSpf struct
type DeviceNetworkinstanceProtocolsOspfInstanceTraceoptionsTraceSpf struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Destaddress *string `json:"dest-address,omitempty"`
}

// DeviceNetworkinstanceProtocolsOspfInstanceTrafficengineering struct
type DeviceNetworkinstanceProtocolsOspfInstanceTrafficengineering struct {
	// +kubebuilder:default:=false
	Advertisement *bool `json:"advertisement,omitempty"`
	// +kubebuilder:default:=true
	Legacylinkattributeadvertisement *bool `json:"legacy-link-attribute-advertisement,omitempty"`
}

// DeviceNetworkinstanceProtocolsPim struct
type DeviceNetworkinstanceProtocolsPim struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_DeviceNetworkinstanceProtocolsPimAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// +kubebuilder:default:=false
	Ecmpbalance *bool `json:"ecmp-balance,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=600
	// +kubebuilder:default:=1
	Ecmpbalancehold *uint32 `json:"ecmp-balance-hold,omitempty"`
	// +kubebuilder:default:=false
	Ecmphashing  *bool   `json:"ecmp-hashing,omitempty"`
	Importpolicy *string `json:"import-policy,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Interface []*DeviceNetworkinstanceProtocolsPimInterface `json:"interface,omitempty"`
	//RootNetworkinstanceProtocolsPimIpv4
	Ipv4 *DeviceNetworkinstanceProtocolsPimIpv4 `json:"ipv4,omitempty"`
	//RootNetworkinstanceProtocolsPimIpv6
	Ipv6 *DeviceNetworkinstanceProtocolsPimIpv6 `json:"ipv6,omitempty"`
	//RootNetworkinstanceProtocolsPimSsm
	Ssm *DeviceNetworkinstanceProtocolsPimSsm `json:"ssm,omitempty"`
}

// DeviceNetworkinstanceProtocolsPimInterface struct
type DeviceNetworkinstanceProtocolsPimInterface struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_DeviceNetworkinstanceProtocolsPimInterfaceAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=300
	// +kubebuilder:default:=60
	Assertinterval *uint32 `json:"assert-interval,omitempty"`
	// +kubebuilder:default:=false
	Bfdipv4 *bool `json:"bfd-ipv4,omitempty"`
	// +kubebuilder:default:=false
	Bfdipv6 *bool `json:"bfd-ipv6,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=30
	Hellointerval *uint32 `json:"hello-interval,omitempty"`
	// kubebuilder:validation:Minimum=20
	// kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=35
	Hellomultiplier *uint32 `json:"hello-multiplier,omitempty"`
	Interfacename   *string `json:"interface-name"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	// +kubebuilder:default:=1
	Priority *uint32 `json:"priority,omitempty"`
}

// DeviceNetworkinstanceProtocolsPimIpv4 struct
type DeviceNetworkinstanceProtocolsPimIpv4 struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_DeviceNetworkinstanceProtocolsPimIpv4Adminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
}

// DeviceNetworkinstanceProtocolsPimIpv6 struct
type DeviceNetworkinstanceProtocolsPimIpv6 struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_DeviceNetworkinstanceProtocolsPimIpv6Adminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
}

// DeviceNetworkinstanceProtocolsPimSsm struct
type DeviceNetworkinstanceProtocolsPimSsm struct {
	//RootNetworkinstanceProtocolsPimSsmSsmranges
	Ssmranges *DeviceNetworkinstanceProtocolsPimSsmSsmranges `json:"ssm-ranges,omitempty"`
}

// DeviceNetworkinstanceProtocolsPimSsmSsmranges struct
type DeviceNetworkinstanceProtocolsPimSsmSsmranges struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Grouprange []*DeviceNetworkinstanceProtocolsPimSsmSsmrangesGrouprange `json:"group-range,omitempty"`
}

// DeviceNetworkinstanceProtocolsPimSsmSsmrangesGrouprange struct
type DeviceNetworkinstanceProtocolsPimSsmSsmrangesGrouprange struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	Ipprefix *string `json:"ip-prefix"`
}

// DeviceNetworkinstanceSegmentrouting struct
type DeviceNetworkinstanceSegmentrouting struct {
	//RootNetworkinstanceSegmentroutingMpls
	Mpls *DeviceNetworkinstanceSegmentroutingMpls `json:"mpls,omitempty"`
	//RootNetworkinstanceSegmentroutingSegmentroutingpolicies
	Segmentroutingpolicies *DeviceNetworkinstanceSegmentroutingSegmentroutingpolicies `json:"segment-routing-policies,omitempty"`
}

// DeviceNetworkinstanceSegmentroutingMpls struct
type DeviceNetworkinstanceSegmentroutingMpls struct {
	//RootNetworkinstanceSegmentroutingMplsGlobalblock
	Globalblock *DeviceNetworkinstanceSegmentroutingMplsGlobalblock `json:"global-block,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=4
	Localprefixsid []*DeviceNetworkinstanceSegmentroutingMplsLocalprefixsid `json:"local-prefix-sid,omitempty"`
}

// DeviceNetworkinstanceSegmentroutingMplsGlobalblock struct
type DeviceNetworkinstanceSegmentroutingMplsGlobalblock struct {
	Labelrange *string `json:"label-range"`
}

// DeviceNetworkinstanceSegmentroutingMplsLocalprefixsid struct
type DeviceNetworkinstanceSegmentroutingMplsLocalprefixsid struct {
	Interface *string `json:"interface"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=1048575
	Ipv4labelindex *uint32 `json:"ipv4-label-index,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=1048575
	Ipv6labelindex *uint32 `json:"ipv6-label-index,omitempty"`
	// +kubebuilder:default:=true
	Nodesid *bool `json:"node-sid,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4
	Prefixsidindex *uint8 `json:"prefix-sid-index"`
}

// DeviceNetworkinstanceSegmentroutingSegmentroutingpolicies struct
type DeviceNetworkinstanceSegmentroutingSegmentroutingpolicies struct {
	//RootNetworkinstanceSegmentroutingSegmentroutingpoliciesNamedpaths
	Namedpaths *DeviceNetworkinstanceSegmentroutingSegmentroutingpoliciesNamedpaths `json:"named-paths,omitempty"`
	//RootNetworkinstanceSegmentroutingSegmentroutingpoliciesProtectionpolicies
	Protectionpolicies *DeviceNetworkinstanceSegmentroutingSegmentroutingpoliciesProtectionpolicies `json:"protection-policies,omitempty"`
	//RootNetworkinstanceSegmentroutingSegmentroutingpoliciesStaticpolicies
	Staticpolicies *DeviceNetworkinstanceSegmentroutingSegmentroutingpoliciesStaticpolicies `json:"static-policies,omitempty"`
}

// DeviceNetworkinstanceSegmentroutingSegmentroutingpoliciesNamedpaths struct
type DeviceNetworkinstanceSegmentroutingSegmentroutingpoliciesNamedpaths struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Path []*DeviceNetworkinstanceSegmentroutingSegmentroutingpoliciesNamedpathsPath `json:"path,omitempty"`
}

// DeviceNetworkinstanceSegmentroutingSegmentroutingpoliciesNamedpathsPath struct
type DeviceNetworkinstanceSegmentroutingSegmentroutingpoliciesNamedpathsPath struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Hop []*DeviceNetworkinstanceSegmentroutingSegmentroutingpoliciesNamedpathsPathHop `json:"hop,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Namedpathname *string `json:"named-path-name"`
}

// DeviceNetworkinstanceSegmentroutingSegmentroutingpoliciesNamedpathsPathHop struct
type DeviceNetworkinstanceSegmentroutingSegmentroutingpoliciesNamedpathsPathHop struct {
	// +kubebuilder:validation:Enum=`loose`;`strict`
	Hoptype E_DeviceNetworkinstanceSegmentroutingSegmentroutingpoliciesNamedpathsPathHopHoptype `json:"hop-type,omitempty"`
	//Hoptype *string `json:"hop-type,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=15
	Index *uint8 `json:"index"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Ipaddress *string `json:"ip-address,omitempty"`
}

// DeviceNetworkinstanceSegmentroutingSegmentroutingpoliciesProtectionpolicies struct
type DeviceNetworkinstanceSegmentroutingSegmentroutingpoliciesProtectionpolicies struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Policy []*DeviceNetworkinstanceSegmentroutingSegmentroutingpoliciesProtectionpoliciesPolicy `json:"policy,omitempty"`
}

// DeviceNetworkinstanceSegmentroutingSegmentroutingpoliciesProtectionpoliciesPolicy struct
type DeviceNetworkinstanceSegmentroutingSegmentroutingpoliciesProtectionpoliciesPolicy struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=3
	Holddowntimer *uint16 `json:"hold-down-timer,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=32
	// +kubebuilder:default:=1
	Minsegmentlistthreshold *uint8 `json:"min-segment-list-threshold,omitempty"`
	// +kubebuilder:validation:Enum=`active-standby`;`ecmp`
	// +kubebuilder:default:="ecmp"
	Mode E_DeviceNetworkinstanceSegmentroutingSegmentroutingpoliciesProtectionpoliciesPolicyMode `json:"mode,omitempty"`
	//Mode *string `json:"mode,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Protectionpolicyname *string `json:"protection-policy-name"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=10
	Reverttimer *uint16 `json:"revert-timer,omitempty"`
	// +kubebuilder:default:=true
	Seamlessbfd *bool `json:"seamless-bfd,omitempty"`
}

// DeviceNetworkinstanceSegmentroutingSegmentroutingpoliciesStaticpolicies struct
type DeviceNetworkinstanceSegmentroutingSegmentroutingpoliciesStaticpolicies struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Policy []*DeviceNetworkinstanceSegmentroutingSegmentroutingpoliciesStaticpoliciesPolicy `json:"policy,omitempty"`
}

// DeviceNetworkinstanceSegmentroutingSegmentroutingpoliciesStaticpoliciesPolicy struct
type DeviceNetworkinstanceSegmentroutingSegmentroutingpoliciesStaticpoliciesPolicy struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_DeviceNetworkinstanceSegmentroutingSegmentroutingpoliciesStaticpoliciesPolicyAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	Color *uint32 `json:"color,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Description *string `json:"description,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Endpoint *string `json:"endpoint"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=100
	Preference       *uint32 `json:"preference,omitempty"`
	Protectionpolicy *string `json:"protection-policy,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=0
	Reoptimizationtimer *uint16 `json:"re-optimization-timer,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Segmentlist []*DeviceNetworkinstanceSegmentroutingSegmentroutingpoliciesStaticpoliciesPolicySegmentlist `json:"segment-list,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Staticpolicyname *string `json:"static-policy-name"`
}

// DeviceNetworkinstanceSegmentroutingSegmentroutingpoliciesStaticpoliciesPolicySegmentlist struct
type DeviceNetworkinstanceSegmentroutingSegmentroutingpoliciesStaticpoliciesPolicySegmentlist struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_DeviceNetworkinstanceSegmentroutingSegmentroutingpoliciesStaticpoliciesPolicySegmentlistAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=32
	Segmentlistindex *uint8  `json:"segment-list-index"`
	Namedpath        *string `json:"named-path,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Segment []*DeviceNetworkinstanceSegmentroutingSegmentroutingpoliciesStaticpoliciesPolicySegmentlistSegment `json:"segment,omitempty"`
}

// DeviceNetworkinstanceSegmentroutingSegmentroutingpoliciesStaticpoliciesPolicySegmentlistSegment struct
type DeviceNetworkinstanceSegmentroutingSegmentroutingpoliciesStaticpoliciesPolicySegmentlistSegment struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	Segmentindex *uint8 `json:"segment-index"`
	//RootNetworkinstanceSegmentroutingSegmentroutingpoliciesStaticpoliciesPolicySegmentlistSegmentSegmenttypea
	Segmenttypea *DeviceNetworkinstanceSegmentroutingSegmentroutingpoliciesStaticpoliciesPolicySegmentlistSegmentSegmenttypea `json:"segment-type-a,omitempty"`
}

// DeviceNetworkinstanceSegmentroutingSegmentroutingpoliciesStaticpoliciesPolicySegmentlistSegmentSegmenttypea struct
type DeviceNetworkinstanceSegmentroutingSegmentroutingpoliciesStaticpoliciesPolicySegmentlistSegmentSegmenttypea struct {
	// kubebuilder:validation:Minimum=16
	// kubebuilder:validation:Maximum=1048575
	Sidvalue *uint32 `json:"sid-value,omitempty"`
}

// DeviceNetworkinstanceStaticroutes struct
type DeviceNetworkinstanceStaticroutes struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=16384
	Route []*DeviceNetworkinstanceStaticroutesRoute `json:"route,omitempty"`
}

// DeviceNetworkinstanceStaticroutesRoute struct
type DeviceNetworkinstanceStaticroutesRoute struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_DeviceNetworkinstanceStaticroutesRouteAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	// +kubebuilder:default:=1
	Metric       *uint32 `json:"metric,omitempty"`
	Nexthopgroup *string `json:"next-hop-group,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=5
	Preference *uint8 `json:"preference,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	Prefix *string `json:"prefix"`
}

// DeviceNetworkinstanceTrafficengineering struct
type DeviceNetworkinstanceTrafficengineering struct {
	//RootNetworkinstanceTrafficengineeringAdmingroups
	Admingroups *DeviceNetworkinstanceTrafficengineeringAdmingroups `json:"admin-groups,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	Autonomoussystem *uint32 `json:"autonomous-system,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Interface []*DeviceNetworkinstanceTrafficengineeringInterface `json:"interface,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Ipv4terouterid *string `json:"ipv4-te-router-id,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Ipv6terouterid *string `json:"ipv6-te-router-id,omitempty"`
	//RootNetworkinstanceTrafficengineeringSharedrisklinkgroups
	Sharedrisklinkgroups *DeviceNetworkinstanceTrafficengineeringSharedrisklinkgroups `json:"shared-risk-link-groups,omitempty"`
}

// DeviceNetworkinstanceTrafficengineeringAdmingroups struct
type DeviceNetworkinstanceTrafficengineeringAdmingroups struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Group []*DeviceNetworkinstanceTrafficengineeringAdmingroupsGroup `json:"group,omitempty"`
}

// DeviceNetworkinstanceTrafficengineeringAdmingroupsGroup struct
type DeviceNetworkinstanceTrafficengineeringAdmingroupsGroup struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=31
	Bitposition *uint32 `json:"bit-position"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
}

// DeviceNetworkinstanceTrafficengineeringInterface struct
type DeviceNetworkinstanceTrafficengineeringInterface struct {
	//RootNetworkinstanceTrafficengineeringInterfaceAdmingroup
	Admingroup *DeviceNetworkinstanceTrafficengineeringInterfaceAdmingroup `json:"admin-group,omitempty"`
	//RootNetworkinstanceTrafficengineeringInterfaceDelay
	Delay         *DeviceNetworkinstanceTrafficengineeringInterfaceDelay `json:"delay,omitempty"`
	Interfacename *string                                                `json:"interface-name"`
	//RootNetworkinstanceTrafficengineeringInterfaceSrlgmembership
	Srlgmembership *DeviceNetworkinstanceTrafficengineeringInterfaceSrlgmembership `json:"srlg-membership,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=16777215
	Temetric *uint32 `json:"te-metric,omitempty"`
}

// DeviceNetworkinstanceTrafficengineeringInterfaceAdmingroup struct
type DeviceNetworkinstanceTrafficengineeringInterfaceAdmingroup struct {
	Admingroup *string `json:"admin-group,omitempty"`
}

// DeviceNetworkinstanceTrafficengineeringInterfaceDelay struct
type DeviceNetworkinstanceTrafficengineeringInterfaceDelay struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	Static *uint32 `json:"static,omitempty"`
}

// DeviceNetworkinstanceTrafficengineeringInterfaceSrlgmembership struct
type DeviceNetworkinstanceTrafficengineeringInterfaceSrlgmembership struct {
	Srlgmembership *string `json:"srlg-membership,omitempty"`
}

// DeviceNetworkinstanceTrafficengineeringSharedrisklinkgroups struct
type DeviceNetworkinstanceTrafficengineeringSharedrisklinkgroups struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Group []*DeviceNetworkinstanceTrafficengineeringSharedrisklinkgroupsGroup `json:"group,omitempty"`
}

// DeviceNetworkinstanceTrafficengineeringSharedrisklinkgroupsGroup struct
type DeviceNetworkinstanceTrafficengineeringSharedrisklinkgroupsGroup struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	Cost *uint32 `json:"cost,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Staticmember []*DeviceNetworkinstanceTrafficengineeringSharedrisklinkgroupsGroupStaticmember `json:"static-member,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	Value *uint32 `json:"value"`
}

// DeviceNetworkinstanceTrafficengineeringSharedrisklinkgroupsGroupStaticmember struct
type DeviceNetworkinstanceTrafficengineeringSharedrisklinkgroupsGroupStaticmember struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Fromaddress *string `json:"from-address"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Toaddress *string `json:"to-address,omitempty"`
}

// DeviceNetworkinstanceVxlaninterface struct
type DeviceNetworkinstanceVxlaninterface struct {
	// kubebuilder:validation:MinLength=8
	// kubebuilder:validation:MaxLength=17
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(vxlan(0|1[0-9][0-9]|2([0-4][0-9]|5[0-5])|[1-9][0-9]|[1-9])\.(0|[1-9](\d){0,8}))`
	Name *string `json:"name"`
}

// DeviceOam struct
type DeviceOam struct {
	//RootOamEthcfm
	Ethcfm *DeviceOamEthcfm `json:"ethcfm,omitempty"`
	//RootOamTwamp
	Twamp *DeviceOamTwamp `json:"twamp,omitempty"`
}

// DeviceOamEthcfm struct
type DeviceOamEthcfm struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=255
	Domain []*DeviceOamEthcfmDomain `json:"domain,omitempty"`
	//RootOamEthcfmPmontemplate
	Pmontemplate *DeviceOamEthcfmPmontemplate `json:"pmon-template,omitempty"`
}

// DeviceOamEthcfmDomain struct
type DeviceOamEthcfmDomain struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Association []*DeviceOamEthcfmDomainAssociation `json:"association,omitempty"`
	// +kubebuilder:validation:Enum=`char-string`;`none`
	// +kubebuilder:default:="none"
	Format E_DeviceOamEthcfmDomainFormat `json:"format,omitempty"`
	//Format *string `json:"format,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=64
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(.*\S.*)|()|([1-9]|[1-9][0-9]{1,8}|[1-3][0-9]{9}|4[0-1][0-9]{8}|42[0-8][0-9]{7}|429[0-3][0-9]{6}|4294[0-8][0-9]{5}|42949[0-5][0-9]{4}|429496[0-6][0-9]{3}|4294967[0-1][0-9]{2}|42949672[0-8][0-9]|429496729[0-5])|([^0-9_ +]\P{C}*[^ ])|([^0-9_ +])`
	Id *string `json:"id"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=7
	// +kubebuilder:default:=0
	Level *uint8 `json:"level,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=43
	// +kubebuilder:default:="DEFAULT"
	Name *string `json:"name,omitempty"`
}

// DeviceOamEthcfmDomainAssociation struct
type DeviceOamEthcfmDomainAssociation struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1
	Bridge []*DeviceOamEthcfmDomainAssociationBridge `json:"bridge,omitempty"`
	// +kubebuilder:validation:Enum=`100ms`;`10ms`;`10s`;`1s`;`300hz`;`600s`;`60s`
	// +kubebuilder:default:="1s"
	Ccminterval E_DeviceOamEthcfmDomainAssociationCcminterval `json:"ccm-interval,omitempty"`
	//Ccminterval *string `json:"ccm-interval,omitempty"`
	// +kubebuilder:validation:Enum=`char-string`;`icc-based`
	Format E_DeviceOamEthcfmDomainAssociationFormat `json:"format,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=64
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(.*\S.*)|()|([1-9]|[1-9][0-9]{1,8}|[1-3][0-9]{9}|4[0-1][0-9]{8}|42[0-8][0-9]{7}|429[0-3][0-9]{6}|4294[0-8][0-9]{5}|42949[0-5][0-9]{4}|429496[0-6][0-9]{3}|4294967[0-1][0-9]{2}|42949672[0-8][0-9]|429496729[0-5])|([^0-9_ +]\P{C}*[^ ])|([^0-9_ +])`
	Id *string `json:"id"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Mep []*DeviceOamEthcfmDomainAssociationMep `json:"mep,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=45
	Name *string `json:"name"`
	//RootOamEthcfmDomainAssociationRemotemeps
	Remotemeps *DeviceOamEthcfmDomainAssociationRemotemeps `json:"remote-meps,omitempty"`
}

// DeviceOamEthcfmDomainAssociationBridge struct
type DeviceOamEthcfmDomainAssociationBridge struct {
	Id *string `json:"id"`
	// +kubebuilder:validation:Enum=`default`;`defer`;`explicit`;`none`
	// +kubebuilder:default:="defer"
	Mhfcreation E_DeviceOamEthcfmDomainAssociationBridgeMhfcreation `json:"mhf-creation,omitempty"`
	//Mhfcreation *string `json:"mhf-creation,omitempty"`
}

// DeviceOamEthcfmDomainAssociationMep struct
type DeviceOamEthcfmDomainAssociationMep struct {
	// +kubebuilder:default:=false
	Adminstate *bool `json:"admin-state,omitempty"`
	// +kubebuilder:default:=false
	Ccmenabled *bool `json:"ccm-enabled,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=7
	// +kubebuilder:default:=7
	Ccmltmpriority *uint8 `json:"ccm-ltm-priority,omitempty"`
	// +kubebuilder:validation:Enum=`down`;`up`
	Direction E_DeviceOamEthcfmDomainAssociationMepDirection `json:"direction,omitempty"`
	// +kubebuilder:validation:Enum=`all-def`;`err-xcon`;`mac-rem-err-xcon`;`no-xcon`;`rem-err-xcon`;`xcon`
	// +kubebuilder:default:="mac-rem-err-xcon"
	Lowestfaultprioritydefect E_DeviceOamEthcfmDomainAssociationMepLowestfaultprioritydefect `json:"lowest-fault-priority-defect,omitempty"`
	//Lowestfaultprioritydefect *string `json:"lowest-fault-priority-defect,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}`
	// +kubebuilder:default:="00:00:00:00:00:00"
	Macaddress *string `json:"mac-address,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=8191
	Mepid *uint16 `json:"mep-id"`
	//RootOamEthcfmDomainAssociationMepPmon
	Pmon         *DeviceOamEthcfmDomainAssociationMepPmon `json:"pmon,omitempty"`
	Interface    *string                                  `json:"interface,omitempty"`
	Subinterface *string                                  `json:"subinterface,omitempty"`
}

// DeviceOamEthcfmDomainAssociationMepPmon struct
type DeviceOamEthcfmDomainAssociationMepPmon struct {
	//RootOamEthcfmDomainAssociationMepPmonOnewaydm
	Onewaydm *DeviceOamEthcfmDomainAssociationMepPmonOnewaydm `json:"one-way-dm,omitempty"`
	//RootOamEthcfmDomainAssociationMepPmonTwowaydm
	Twowaydm *DeviceOamEthcfmDomainAssociationMepPmonTwowaydm `json:"two-way-dm,omitempty"`
	//RootOamEthcfmDomainAssociationMepPmonTwowayslm
	Twowayslm *DeviceOamEthcfmDomainAssociationMepPmonTwowayslm `json:"two-way-slm,omitempty"`
}

// DeviceOamEthcfmDomainAssociationMepPmonOnewaydm struct
type DeviceOamEthcfmDomainAssociationMepPmonOnewaydm struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=8
	Session []*DeviceOamEthcfmDomainAssociationMepPmonOnewaydmSession `json:"session,omitempty"`
}

// DeviceOamEthcfmDomainAssociationMepPmonOnewaydmSession struct
type DeviceOamEthcfmDomainAssociationMepPmonOnewaydmSession struct {
	// +kubebuilder:default:=false
	Adminstate *bool `json:"admin-state,omitempty"`
	// +kubebuilder:default:="1"
	Bingroup *string `json:"bin-group,omitempty"`
	// kubebuilder:validation:Minimum=64
	// kubebuilder:validation:Maximum=9600
	// +kubebuilder:default:=64
	Framesize *uint32 `json:"frame-size,omitempty"`
	// kubebuilder:validation:Minimum=3
	// kubebuilder:validation:Maximum=3600000
	// +kubebuilder:default:=100
	Interval *uint32 `json:"interval,omitempty"`
	// +kubebuilder:default:="1"
	Measurementinterval *string `json:"measurement-interval,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=7
	// +kubebuilder:default:=7
	Priority *uint8 `json:"priority,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=64
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(.*\S.*)|()|([1-9]|[1-9][0-9]{1,8}|[1-3][0-9]{9}|4[0-1][0-9]{8}|42[0-8][0-9]{7}|429[0-3][0-9]{6}|4294[0-8][0-9]{5}|42949[0-5][0-9]{4}|429496[0-6][0-9]{3}|4294967[0-1][0-9]{2}|42949672[0-8][0-9]|429496729[0-5])|([^0-9_ +]\P{C}*[^ ])|([^0-9_ +])`
	Sessionid *string `json:"session-id"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}`
	Target *string `json:"target"`
}

// DeviceOamEthcfmDomainAssociationMepPmonTwowaydm struct
type DeviceOamEthcfmDomainAssociationMepPmonTwowaydm struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=8
	Session []*DeviceOamEthcfmDomainAssociationMepPmonTwowaydmSession `json:"session,omitempty"`
}

// DeviceOamEthcfmDomainAssociationMepPmonTwowaydmSession struct
type DeviceOamEthcfmDomainAssociationMepPmonTwowaydmSession struct {
	// +kubebuilder:default:=false
	Adminstate *bool `json:"admin-state,omitempty"`
	// +kubebuilder:default:="1"
	Bingroup *string `json:"bin-group,omitempty"`
	// kubebuilder:validation:Minimum=64
	// kubebuilder:validation:Maximum=9600
	// +kubebuilder:default:=64
	Framesize *uint32 `json:"frame-size,omitempty"`
	// kubebuilder:validation:Minimum=3
	// kubebuilder:validation:Maximum=3600000
	// +kubebuilder:default:=100
	Interval *uint32 `json:"interval,omitempty"`
	// +kubebuilder:default:="1"
	Measurementinterval *string `json:"measurement-interval,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=7
	// +kubebuilder:default:=7
	Priority *uint8 `json:"priority,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=64
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(.*\S.*)|()|([1-9]|[1-9][0-9]{1,8}|[1-3][0-9]{9}|4[0-1][0-9]{8}|42[0-8][0-9]{7}|429[0-3][0-9]{6}|4294[0-8][0-9]{5}|42949[0-5][0-9]{4}|429496[0-6][0-9]{3}|4294967[0-1][0-9]{2}|42949672[0-8][0-9]|429496729[0-5])|([^0-9_ +]\P{C}*[^ ])|([^0-9_ +])`
	Sessionid *string `json:"session-id"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}`
	Target *string `json:"target"`
}

// DeviceOamEthcfmDomainAssociationMepPmonTwowayslm struct
type DeviceOamEthcfmDomainAssociationMepPmonTwowayslm struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=8
	Session []*DeviceOamEthcfmDomainAssociationMepPmonTwowayslmSession `json:"session,omitempty"`
}

// DeviceOamEthcfmDomainAssociationMepPmonTwowayslmSession struct
type DeviceOamEthcfmDomainAssociationMepPmonTwowayslmSession struct {
	// +kubebuilder:default:=false
	Adminstate *bool `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=1000
	// +kubebuilder:default:=10
	Consecutiveintervals *uint32 `json:"consecutive-intervals,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100000
	// +kubebuilder:default:=50000
	Flrthreshold *uint32 `json:"flr-threshold,omitempty"`
	// kubebuilder:validation:Minimum=64
	// kubebuilder:validation:Maximum=9600
	// +kubebuilder:default:=64
	Framesize *uint32 `json:"frame-size,omitempty"`
	// kubebuilder:validation:Minimum=3
	// kubebuilder:validation:Maximum=3600000
	// +kubebuilder:default:=100
	Interval *uint32 `json:"interval,omitempty"`
	//RootOamEthcfmDomainAssociationMepPmonTwowayslmSessionLossevent
	Lossevent *DeviceOamEthcfmDomainAssociationMepPmonTwowayslmSessionLossevent `json:"loss-event,omitempty"`
	// +kubebuilder:default:="1"
	Measurementinterval *string `json:"measurement-interval,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=7
	// +kubebuilder:default:=7
	Priority *uint8 `json:"priority,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=64
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(.*\S.*)|()|([1-9]|[1-9][0-9]{1,8}|[1-3][0-9]{9}|4[0-1][0-9]{8}|42[0-8][0-9]{7}|429[0-3][0-9]{6}|4294[0-8][0-9]{5}|42949[0-5][0-9]{4}|429496[0-6][0-9]{3}|4294967[0-1][0-9]{2}|42949672[0-8][0-9]|429496729[0-5])|([^0-9_ +]\P{C}*[^ ])|([^0-9_ +])`
	Sessionid *string `json:"session-id"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}`
	Target *string `json:"target"`
}

// DeviceOamEthcfmDomainAssociationMepPmonTwowayslmSessionLossevent struct
type DeviceOamEthcfmDomainAssociationMepPmonTwowayslmSessionLossevent struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Avgflr []*DeviceOamEthcfmDomainAssociationMepPmonTwowayslmSessionLosseventAvgflr `json:"avg-flr,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Chli []*DeviceOamEthcfmDomainAssociationMepPmonTwowayslmSessionLosseventChli `json:"chli,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Hli []*DeviceOamEthcfmDomainAssociationMepPmonTwowayslmSessionLosseventHli `json:"hli,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Unavail []*DeviceOamEthcfmDomainAssociationMepPmonTwowayslmSessionLosseventUnavail `json:"unavail,omitempty"`
}

// DeviceOamEthcfmDomainAssociationMepPmonTwowayslmSessionLosseventAvgflr struct
type DeviceOamEthcfmDomainAssociationMepPmonTwowayslmSessionLosseventAvgflr struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100000
	// +kubebuilder:default:=0
	Clearthreshold *uint32 `json:"clear-threshold,omitempty"`
	// +kubebuilder:validation:Enum=`backward`;`forward`
	Direction E_DeviceOamEthcfmDomainAssociationMepPmonTwowayslmSessionLosseventAvgflrDirection `json:"direction,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100000
	// +kubebuilder:default:=0
	Raisethreshold *uint32 `json:"raise-threshold,omitempty"`
}

// DeviceOamEthcfmDomainAssociationMepPmonTwowayslmSessionLosseventChli struct
type DeviceOamEthcfmDomainAssociationMepPmonTwowayslmSessionLosseventChli struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=863999
	// +kubebuilder:default:=0
	Clearthreshold *uint32 `json:"clear-threshold,omitempty"`
	// +kubebuilder:validation:Enum=`backward`;`forward`
	Direction E_DeviceOamEthcfmDomainAssociationMepPmonTwowayslmSessionLosseventChliDirection `json:"direction,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=864000
	// +kubebuilder:default:=0
	Raisethreshold *uint32 `json:"raise-threshold,omitempty"`
}

// DeviceOamEthcfmDomainAssociationMepPmonTwowayslmSessionLosseventHli struct
type DeviceOamEthcfmDomainAssociationMepPmonTwowayslmSessionLosseventHli struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=863999
	// +kubebuilder:default:=0
	Clearthreshold *uint32 `json:"clear-threshold,omitempty"`
	// +kubebuilder:validation:Enum=`backward`;`forward`
	Direction E_DeviceOamEthcfmDomainAssociationMepPmonTwowayslmSessionLosseventHliDirection `json:"direction,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=864000
	// +kubebuilder:default:=0
	Raisethreshold *uint32 `json:"raise-threshold,omitempty"`
}

// DeviceOamEthcfmDomainAssociationMepPmonTwowayslmSessionLosseventUnavail struct
type DeviceOamEthcfmDomainAssociationMepPmonTwowayslmSessionLosseventUnavail struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=863999
	// +kubebuilder:default:=0
	Clearthreshold *uint32 `json:"clear-threshold,omitempty"`
	// +kubebuilder:validation:Enum=`backward`;`forward`
	Direction E_DeviceOamEthcfmDomainAssociationMepPmonTwowayslmSessionLosseventUnavailDirection `json:"direction,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=864000
	// +kubebuilder:default:=0
	Raisethreshold *uint32 `json:"raise-threshold,omitempty"`
}

// DeviceOamEthcfmDomainAssociationRemotemeps struct
type DeviceOamEthcfmDomainAssociationRemotemeps struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=8191
	Remotemeps *uint16 `json:"remote-meps,omitempty"`
}

// DeviceOamEthcfmPmontemplate struct
type DeviceOamEthcfmPmontemplate struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=512
	Bingroup []*DeviceOamEthcfmPmontemplateBingroup `json:"bin-group,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=512
	Measurementinterval []*DeviceOamEthcfmPmontemplateMeasurementinterval `json:"measurement-interval,omitempty"`
}

// DeviceOamEthcfmPmontemplateBingroup struct
type DeviceOamEthcfmPmontemplateBingroup struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=512
	Bgid *uint32 `json:"bg-id"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Bintype []*DeviceOamEthcfmPmontemplateBingroupBintype `json:"bin-type,omitempty"`
}

// DeviceOamEthcfmPmontemplateBingroupBintype struct
type DeviceOamEthcfmPmontemplateBingroupBintype struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=10
	Bin []*DeviceOamEthcfmPmontemplateBingroupBintypeBin `json:"bin,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=10
	// +kubebuilder:default:=2
	Bincount *uint8 `json:"bin-count,omitempty"`
	// +kubebuilder:validation:Enum=`fd`;`ifdv`
	Btid E_DeviceOamEthcfmPmontemplateBingroupBintypeBtid `json:"bt-id,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Delayevent []*DeviceOamEthcfmPmontemplateBingroupBintypeDelayevent `json:"delay-event,omitempty"`
}

// DeviceOamEthcfmPmontemplateBingroupBintypeBin struct
type DeviceOamEthcfmPmontemplateBingroupBintypeBin struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=9
	Binid *uint8 `json:"bin-id"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=5000
	Lowerbound *uint32 `json:"lower-bound,omitempty"`
}

// DeviceOamEthcfmPmontemplateBingroupBintypeDelayevent struct
type DeviceOamEthcfmPmontemplateBingroupBintypeDelayevent struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=863999
	// +kubebuilder:default:=0
	Clearthreshold *uint32 `json:"clear-threshold,omitempty"`
	// +kubebuilder:validation:Enum=`backward`;`forward`;`round-trip`
	Direction E_DeviceOamEthcfmPmontemplateBingroupBintypeDelayeventDirection `json:"direction,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=9
	// +kubebuilder:default:=0
	Lowestbin *uint8 `json:"lowest-bin,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=864000
	// +kubebuilder:default:=0
	Raisethreshold *uint32 `json:"raise-threshold,omitempty"`
}

// DeviceOamEthcfmPmontemplateMeasurementinterval struct
type DeviceOamEthcfmPmontemplateMeasurementinterval struct {
	// +kubebuilder:validation:Enum=`clock-aligned`;`test-relative`
	// +kubebuilder:default:="clock-aligned"
	Boundarytype E_DeviceOamEthcfmPmontemplateMeasurementintervalBoundarytype `json:"boundary-type,omitempty"`
	//Boundarytype *string `json:"boundary-type,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=4
	Duration []*DeviceOamEthcfmPmontemplateMeasurementintervalDuration `json:"duration,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=512
	Miid *uint32 `json:"mi-id"`
}

// DeviceOamEthcfmPmontemplateMeasurementintervalDuration struct
type DeviceOamEthcfmPmontemplateMeasurementintervalDuration struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=96
	Intervalcount *uint8 `json:"interval-count,omitempty"`
	// +kubebuilder:validation:Enum=`1-day`;`1-hour`;`15-minutes`;`5-minutes`
	Mitype E_DeviceOamEthcfmPmontemplateMeasurementintervalDurationMitype `json:"mi-type,omitempty"`
}

// DeviceOamTwamp struct
type DeviceOamTwamp struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1
	Server []*DeviceOamTwampServer `json:"server,omitempty"`
}

// DeviceOamTwampServer struct
type DeviceOamTwampServer struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_DeviceOamTwampServerAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Clientconnection []*DeviceOamTwampServerClientconnection `json:"client-connection,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=63
	Controlpacketdscp *uint8 `json:"control-packet-dscp,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=256
	Maxconnserver *uint32 `json:"max-conn-server,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=1024
	Maxsessserver *uint32 `json:"max-sess-server,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Serverinstancename *string `json:"server-instance-name"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=604800
	// +kubebuilder:default:=900
	Servwait *uint32 `json:"servwait,omitempty"`
	//RootOamTwampServerSessionreflector
	Sessionreflector *DeviceOamTwampServerSessionreflector `json:"session-reflector,omitempty"`
}

// DeviceOamTwampServerClientconnection struct
type DeviceOamTwampServerClientconnection struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	Clientip *string `json:"client-ip"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=256
	Maxconnserver *uint32 `json:"max-conn-server,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=1024
	Maxsessserver *uint32 `json:"max-sess-server,omitempty"`
}

// DeviceOamTwampServerSessionreflector struct
type DeviceOamTwampServerSessionreflector struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_DeviceOamTwampServerSessionreflectorAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=604800
	// +kubebuilder:default:=900
	Refwait *uint32 `json:"refwait,omitempty"`
}

// DevicePlatform struct
type DevicePlatform struct {
	//RootPlatformChassis
	Chassis *DevicePlatformChassis `json:"chassis,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Fabric []*DevicePlatformFabric `json:"fabric,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Linecard []*DevicePlatformLinecard `json:"linecard,omitempty"`
	//RootPlatformRedundancy
	Redundancy *DevicePlatformRedundancy `json:"redundancy,omitempty"`
	//RootPlatformResourcemanagement
	Resourcemanagement *DevicePlatformResourcemanagement `json:"resource-management,omitempty"`
	//RootPlatformResourcemonitoring
	Resourcemonitoring *DevicePlatformResourcemonitoring `json:"resource-monitoring,omitempty"`
	//RootPlatformVxdp
	Vxdp *DevicePlatformVxdp `json:"vxdp,omitempty"`
}

// DevicePlatformChassis struct
type DevicePlatformChassis struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}`
	Macaddress *string `json:"mac-address,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=8192
	// +kubebuilder:default:=1024
	Macaddressallocation *uint32 `json:"mac-address-allocation,omitempty"`
	//RootPlatformChassisPower
	Power *DevicePlatformChassisPower `json:"power,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`
	Uuid *string `json:"uuid,omitempty"`
}

// DevicePlatformChassisPower struct
type DevicePlatformChassisPower struct {
}

// DevicePlatformFabric struct
type DevicePlatformFabric struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_DevicePlatformFabricAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=255
	Slot *uint8 `json:"slot"`
}

// DevicePlatformLinecard struct
type DevicePlatformLinecard struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_DevicePlatformLinecardAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Forwardingcomplex []*DevicePlatformLinecardForwardingcomplex `json:"forwarding-complex,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=8
	Slot *uint8 `json:"slot"`
}

// DevicePlatformLinecardForwardingcomplex struct
type DevicePlatformLinecardForwardingcomplex struct {
	//RootPlatformLinecardForwardingcomplexBuffermemory
	Buffermemory *DevicePlatformLinecardForwardingcomplexBuffermemory `json:"buffer-memory,omitempty"`
	//RootPlatformLinecardForwardingcomplexFabric
	Fabric *DevicePlatformLinecardForwardingcomplexFabric `json:"fabric,omitempty"`
	// +kubebuilder:validation:Enum=`0`;`1`
	Name E_DevicePlatformLinecardForwardingcomplexName `json:"name,omitempty"`
	//RootPlatformLinecardForwardingcomplexP4rt
	P4rt *DevicePlatformLinecardForwardingcomplexP4rt `json:"p4rt,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Pipeline []*DevicePlatformLinecardForwardingcomplexPipeline `json:"pipeline,omitempty"`
}

// DevicePlatformLinecardForwardingcomplexBuffermemory struct
type DevicePlatformLinecardForwardingcomplexBuffermemory struct {
}

// DevicePlatformLinecardForwardingcomplexFabric struct
type DevicePlatformLinecardForwardingcomplexFabric struct {
}

// DevicePlatformLinecardForwardingcomplexP4rt struct
type DevicePlatformLinecardForwardingcomplexP4rt struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=-1
	Id *uint64 `json:"id,omitempty"`
}

// DevicePlatformLinecardForwardingcomplexPipeline struct
type DevicePlatformLinecardForwardingcomplexPipeline struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=7
	Index *uint8 `json:"index"`
}

// DevicePlatformRedundancy struct
type DevicePlatformRedundancy struct {
	//RootPlatformRedundancySynchronization
	Synchronization *DevicePlatformRedundancySynchronization `json:"synchronization,omitempty"`
}

// DevicePlatformRedundancySynchronization struct
type DevicePlatformRedundancySynchronization struct {
	//RootPlatformRedundancySynchronizationOverlay
	Overlay *DevicePlatformRedundancySynchronizationOverlay `json:"overlay,omitempty"`
}

// DevicePlatformRedundancySynchronizationOverlay struct
type DevicePlatformRedundancySynchronizationOverlay struct {
	// kubebuilder:validation:Minimum=30
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=60
	Synchronizationfrequency *uint32 `json:"synchronization-frequency,omitempty"`
}

// DevicePlatformResourcemanagement struct
type DevicePlatformResourcemanagement struct {
	//RootPlatformResourcemanagementTcam
	Tcam *DevicePlatformResourcemanagementTcam `json:"tcam,omitempty"`
	//RootPlatformResourcemanagementUnifiedforwardingresources
	Unifiedforwardingresources *DevicePlatformResourcemanagementUnifiedforwardingresources `json:"unified-forwarding-resources,omitempty"`
}

// DevicePlatformResourcemanagementTcam struct
type DevicePlatformResourcemanagementTcam struct {
}

// DevicePlatformResourcemanagementUnifiedforwardingresources struct
type DevicePlatformResourcemanagementUnifiedforwardingresources struct {
	// +kubebuilder:validation:Enum=`disabled`;`enabled`;`high-scale`
	// +kubebuilder:default:="disabled"
	Alpm E_DevicePlatformResourcemanagementUnifiedforwardingresourcesAlpm `json:"alpm,omitempty"`
	//Alpm *string `json:"alpm,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=8192
	Ipv6128bitlpmentries *uint16 `json:"ipv6-128bit-lpm-entries,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=262144
	Requestedextraiphostentries *uint32 `json:"requested-extra-ip-host-entries,omitempty"`
}

// DevicePlatformResourcemonitoring struct
type DevicePlatformResourcemonitoring struct {
	//RootPlatformResourcemonitoringAcl
	Acl *DevicePlatformResourcemonitoringAcl `json:"acl,omitempty"`
	//RootPlatformResourcemonitoringDatapath
	Datapath *DevicePlatformResourcemonitoringDatapath `json:"datapath,omitempty"`
	//RootPlatformResourcemonitoringMtu
	Mtu *DevicePlatformResourcemonitoringMtu `json:"mtu,omitempty"`
	//RootPlatformResourcemonitoringQos
	Qos *DevicePlatformResourcemonitoringQos `json:"qos,omitempty"`
	//RootPlatformResourcemonitoringTcam
	Tcam *DevicePlatformResourcemonitoringTcam `json:"tcam,omitempty"`
}

// DevicePlatformResourcemonitoringAcl struct
type DevicePlatformResourcemonitoringAcl struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Resource []*DevicePlatformResourcemonitoringAclResource `json:"resource,omitempty"`
}

// DevicePlatformResourcemonitoringAclResource struct
type DevicePlatformResourcemonitoringAclResource struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=70
	Fallingthresholdlog *uint8  `json:"falling-threshold-log,omitempty"`
	Name                *string `json:"name"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=90
	Risingthresholdlog *uint8 `json:"rising-threshold-log,omitempty"`
}

// DevicePlatformResourcemonitoringDatapath struct
type DevicePlatformResourcemonitoringDatapath struct {
	//RootPlatformResourcemonitoringDatapathAsic
	Asic *DevicePlatformResourcemonitoringDatapathAsic `json:"asic,omitempty"`
	//RootPlatformResourcemonitoringDatapathXdp
	Xdp *DevicePlatformResourcemonitoringDatapathXdp `json:"xdp,omitempty"`
}

// DevicePlatformResourcemonitoringDatapathAsic struct
type DevicePlatformResourcemonitoringDatapathAsic struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Resource []*DevicePlatformResourcemonitoringDatapathAsicResource `json:"resource,omitempty"`
}

// DevicePlatformResourcemonitoringDatapathAsicResource struct
type DevicePlatformResourcemonitoringDatapathAsicResource struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=70
	Fallingthresholdlog *uint8  `json:"falling-threshold-log,omitempty"`
	Name                *string `json:"name"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=90
	Risingthresholdlog *uint8 `json:"rising-threshold-log,omitempty"`
}

// DevicePlatformResourcemonitoringDatapathXdp struct
type DevicePlatformResourcemonitoringDatapathXdp struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Resource []*DevicePlatformResourcemonitoringDatapathXdpResource `json:"resource,omitempty"`
}

// DevicePlatformResourcemonitoringDatapathXdpResource struct
type DevicePlatformResourcemonitoringDatapathXdpResource struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=70
	Fallingthresholdlog *uint8  `json:"falling-threshold-log,omitempty"`
	Name                *string `json:"name"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=90
	Risingthresholdlog *uint8 `json:"rising-threshold-log,omitempty"`
}

// DevicePlatformResourcemonitoringMtu struct
type DevicePlatformResourcemonitoringMtu struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Resource []*DevicePlatformResourcemonitoringMtuResource `json:"resource,omitempty"`
}

// DevicePlatformResourcemonitoringMtuResource struct
type DevicePlatformResourcemonitoringMtuResource struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=70
	Fallingthresholdlog *uint8  `json:"falling-threshold-log,omitempty"`
	Name                *string `json:"name"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=90
	Risingthresholdlog *uint8 `json:"rising-threshold-log,omitempty"`
}

// DevicePlatformResourcemonitoringQos struct
type DevicePlatformResourcemonitoringQos struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Resource []*DevicePlatformResourcemonitoringQosResource `json:"resource,omitempty"`
}

// DevicePlatformResourcemonitoringQosResource struct
type DevicePlatformResourcemonitoringQosResource struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=70
	Fallingthresholdlog *uint8  `json:"falling-threshold-log,omitempty"`
	Name                *string `json:"name"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=90
	Risingthresholdlog *uint8 `json:"rising-threshold-log,omitempty"`
}

// DevicePlatformResourcemonitoringTcam struct
type DevicePlatformResourcemonitoringTcam struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Resource []*DevicePlatformResourcemonitoringTcamResource `json:"resource,omitempty"`
}

// DevicePlatformResourcemonitoringTcamResource struct
type DevicePlatformResourcemonitoringTcamResource struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=70
	Fallingthresholdlog *uint8  `json:"falling-threshold-log,omitempty"`
	Name                *string `json:"name"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=90
	Risingthresholdlog *uint8 `json:"rising-threshold-log,omitempty"`
}

// DevicePlatformVxdp struct
type DevicePlatformVxdp struct {
	//RootPlatformVxdpCpuset
	Cpuset *DevicePlatformVxdpCpuset `json:"cpu-set,omitempty"`
}

// DevicePlatformVxdpCpuset struct
type DevicePlatformVxdpCpuset struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	Cpuset *uint8 `json:"cpu-set,omitempty"`
}

// DeviceQos struct
type DeviceQos struct {
	//RootQosClassifiers
	Classifiers *DeviceQosClassifiers `json:"classifiers,omitempty"`
	//RootQosExplicitcongestionnotification
	Explicitcongestionnotification *DeviceQosExplicitcongestionnotification `json:"explicit-congestion-notification,omitempty"`
	//RootQosQueuetemplates
	Queuetemplates *DeviceQosQueuetemplates `json:"queue-templates,omitempty"`
	//RootQosRewriterules
	Rewriterules *DeviceQosRewriterules `json:"rewrite-rules,omitempty"`
}

// DeviceQosClassifiers struct
type DeviceQosClassifiers struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Dscppolicy []*DeviceQosClassifiersDscppolicy `json:"dscp-policy,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Mplstrafficclasspolicy []*DeviceQosClassifiersMplstrafficclasspolicy `json:"mpls-traffic-class-policy,omitempty"`
	Vxlandefault           *string                                       `json:"vxlan-default,omitempty"`
}

// DeviceQosClassifiersDscppolicy struct
type DeviceQosClassifiersDscppolicy struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Dscp []*DeviceQosClassifiersDscppolicyDscp `json:"dscp,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
}

// DeviceQosClassifiersDscppolicyDscp struct
type DeviceQosClassifiersDscppolicyDscp struct {
	// +kubebuilder:validation:Enum=`high`;`low`;`medium`
	Dropprobability E_DeviceQosClassifiersDscppolicyDscpDropprobability `json:"drop-probability,omitempty"`
	//Dropprobability *string `json:"drop-probability,omitempty"`
	// +kubebuilder:validation:Enum=`fc0`;`fc1`;`fc2`;`fc3`;`fc4`;`fc5`;`fc6`;`fc7`
	Forwardingclass E_DeviceQosClassifiersDscppolicyDscpForwardingclass `json:"forwarding-class,omitempty"`
	//Forwardingclass *string `json:"forwarding-class,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=63
	Value *uint8 `json:"value"`
}

// DeviceQosClassifiersMplstrafficclasspolicy struct
type DeviceQosClassifiersMplstrafficclasspolicy struct {
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Trafficclass []*DeviceQosClassifiersMplstrafficclasspolicyTrafficclass `json:"traffic-class,omitempty"`
}

// DeviceQosClassifiersMplstrafficclasspolicyTrafficclass struct
type DeviceQosClassifiersMplstrafficclasspolicyTrafficclass struct {
	// +kubebuilder:validation:Enum=`high`;`low`;`medium`
	Dropprobability E_DeviceQosClassifiersMplstrafficclasspolicyTrafficclassDropprobability `json:"drop-probability,omitempty"`
	//Dropprobability *string `json:"drop-probability,omitempty"`
	// +kubebuilder:validation:Enum=`fc0`;`fc1`;`fc2`;`fc3`;`fc4`;`fc5`;`fc6`;`fc7`
	Forwardingclass E_DeviceQosClassifiersMplstrafficclasspolicyTrafficclassForwardingclass `json:"forwarding-class,omitempty"`
	//Forwardingclass *string `json:"forwarding-class,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=7
	Value *uint8 `json:"value"`
}

// DeviceQosExplicitcongestionnotification struct
type DeviceQosExplicitcongestionnotification struct {
	Ecndscppolicy *string `json:"ecn-dscp-policy"`
}

// DeviceQosQueuetemplates struct
type DeviceQosQueuetemplates struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=64
	Queuetemplate []*DeviceQosQueuetemplatesQueuetemplate `json:"queue-template,omitempty"`
}

// DeviceQosQueuetemplatesQueuetemplate struct
type DeviceQosQueuetemplatesQueuetemplate struct {
	//RootQosQueuetemplatesQueuetemplateActivequeuemanagement
	Activequeuemanagement *DeviceQosQueuetemplatesQueuetemplateActivequeuemanagement `json:"active-queue-management,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
	//RootQosQueuetemplatesQueuetemplateQueuedepth
	Queuedepth *DeviceQosQueuetemplatesQueuetemplateQueuedepth `json:"queue-depth,omitempty"`
}

// DeviceQosQueuetemplatesQueuetemplateActivequeuemanagement struct
type DeviceQosQueuetemplatesQueuetemplateActivequeuemanagement struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Ecnslope []*DeviceQosQueuetemplatesQueuetemplateActivequeuemanagementEcnslope `json:"ecn-slope,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=15
	// +kubebuilder:default:=0
	Weightfactor *uint8 `json:"weight-factor,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Wredslope []*DeviceQosQueuetemplatesQueuetemplateActivequeuemanagementWredslope `json:"wred-slope,omitempty"`
}

// DeviceQosQueuetemplatesQueuetemplateActivequeuemanagementEcnslope struct
type DeviceQosQueuetemplatesQueuetemplateActivequeuemanagementEcnslope struct {
	// +kubebuilder:validation:Enum=`all`;`high`;`low`;`medium`
	Ecndropprobability E_DeviceQosQueuetemplatesQueuetemplateActivequeuemanagementEcnslopeEcndropprobability `json:"ecn-drop-probability,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=0
	Maxprobability *uint8 `json:"max-probability,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=100
	Maxthresholdpercent *uint8 `json:"max-threshold-percent,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=100
	Minthresholdpercent *uint8 `json:"min-threshold-percent,omitempty"`
}

// DeviceQosQueuetemplatesQueuetemplateActivequeuemanagementWredslope struct
type DeviceQosQueuetemplatesQueuetemplateActivequeuemanagementWredslope struct {
	// +kubebuilder:validation:Enum=`high`;`low`;`medium`
	Dropprobability E_DeviceQosQueuetemplatesQueuetemplateActivequeuemanagementWredslopeDropprobability `json:"drop-probability,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=0
	Maxprobability *uint8 `json:"max-probability,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=100
	Maxthresholdpercent *uint8 `json:"max-threshold-percent,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=100
	Minthresholdpercent *uint8 `json:"min-threshold-percent,omitempty"`
	// +kubebuilder:validation:Enum=`all`;`non-tcp`;`tcp`
	Traffictype E_DeviceQosQueuetemplatesQueuetemplateActivequeuemanagementWredslopeTraffictype `json:"traffic-type,omitempty"`
}

// DeviceQosQueuetemplatesQueuetemplateQueuedepth struct
type DeviceQosQueuetemplatesQueuetemplateQueuedepth struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	// +kubebuilder:default:=0
	Highthresholdbytes *uint32 `json:"high-threshold-bytes,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	// +kubebuilder:default:=0
	Maximumburstsize *uint32 `json:"maximum-burst-size,omitempty"`
}

// DeviceQosRewriterules struct
type DeviceQosRewriterules struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Dscppolicy []*DeviceQosRewriterulesDscppolicy `json:"dscp-policy,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Mplstrafficclasspolicy []*DeviceQosRewriterulesMplstrafficclasspolicy `json:"mpls-traffic-class-policy,omitempty"`
}

// DeviceQosRewriterulesDscppolicy struct
type DeviceQosRewriterulesDscppolicy struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Map []*DeviceQosRewriterulesDscppolicyMap `json:"map,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
}

// DeviceQosRewriterulesDscppolicyMap struct
type DeviceQosRewriterulesDscppolicyMap struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Dropprobability []*DeviceQosRewriterulesDscppolicyMapDropprobability `json:"drop-probability,omitempty"`
	Dscp            *string                                              `json:"dscp"`
	// +kubebuilder:validation:Enum=`fc0`;`fc1`;`fc2`;`fc3`;`fc4`;`fc5`;`fc6`;`fc7`
	Forwardingclass E_DeviceQosRewriterulesDscppolicyMapForwardingclass `json:"forwarding-class,omitempty"`
}

// DeviceQosRewriterulesDscppolicyMapDropprobability struct
type DeviceQosRewriterulesDscppolicyMapDropprobability struct {
	// +kubebuilder:validation:Enum=`high`;`low`;`medium`
	Dropprobability E_DeviceQosRewriterulesDscppolicyMapDropprobabilityDropprobability `json:"drop-probability,omitempty"`
	Dscp            *string                                                            `json:"dscp"`
}

// DeviceQosRewriterulesMplstrafficclasspolicy struct
type DeviceQosRewriterulesMplstrafficclasspolicy struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Map []*DeviceQosRewriterulesMplstrafficclasspolicyMap `json:"map,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
}

// DeviceQosRewriterulesMplstrafficclasspolicyMap struct
type DeviceQosRewriterulesMplstrafficclasspolicyMap struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Dropprobability []*DeviceQosRewriterulesMplstrafficclasspolicyMapDropprobability `json:"drop-probability,omitempty"`
	// +kubebuilder:validation:Enum=`fc0`;`fc1`;`fc2`;`fc3`;`fc4`;`fc5`;`fc6`;`fc7`
	Forwardingclass E_DeviceQosRewriterulesMplstrafficclasspolicyMapForwardingclass `json:"forwarding-class,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=7
	Trafficclass *uint8 `json:"traffic-class"`
}

// DeviceQosRewriterulesMplstrafficclasspolicyMapDropprobability struct
type DeviceQosRewriterulesMplstrafficclasspolicyMapDropprobability struct {
	// +kubebuilder:validation:Enum=`high`;`low`;`medium`
	Dropprobability E_DeviceQosRewriterulesMplstrafficclasspolicyMapDropprobabilityDropprobability `json:"drop-probability,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=7
	Trafficclass *uint8 `json:"traffic-class"`
}

// DeviceRoutingpolicy struct
type DeviceRoutingpolicy struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Aspathset []*DeviceRoutingpolicyAspathset `json:"as-path-set,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Communityset []*DeviceRoutingpolicyCommunityset `json:"community-set,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Policy []*DeviceRoutingpolicyPolicy `json:"policy,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Prefixset []*DeviceRoutingpolicyPrefixset `json:"prefix-set,omitempty"`
}

// DeviceRoutingpolicyAspathset struct
type DeviceRoutingpolicyAspathset struct {
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=65535
	Expression *string `json:"expression,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
}

// DeviceRoutingpolicyCommunityset struct
type DeviceRoutingpolicyCommunityset struct {
	//RootRoutingpolicyCommunitysetMember
	Member *DeviceRoutingpolicyCommunitysetMember `json:"member,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
}

// DeviceRoutingpolicyCommunitysetMember struct
type DeviceRoutingpolicyCommunitysetMember struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])|.*:.*|([1-9][0-9]{0,9}):([1-9][0-9]{0,9}):([1-9][0-9]{0,9})|.*:.*:.*`
	Member *string `json:"member,omitempty"`
}

// DeviceRoutingpolicyPolicy struct
type DeviceRoutingpolicyPolicy struct {
	//RootRoutingpolicyPolicyDefaultaction
	Defaultaction *DeviceRoutingpolicyPolicyDefaultaction `json:"default-action,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Statement []*DeviceRoutingpolicyPolicyStatement `json:"statement,omitempty"`
}

// DeviceRoutingpolicyPolicyDefaultaction struct
type DeviceRoutingpolicyPolicyDefaultaction struct {
	//RootRoutingpolicyPolicyDefaultactionAccept
	Accept *DeviceRoutingpolicyPolicyDefaultactionAccept `json:"accept,omitempty"`
	//RootRoutingpolicyPolicyDefaultactionNextentry
	Nextentry *DeviceRoutingpolicyPolicyDefaultactionNextentry `json:"next-entry,omitempty"`
	//RootRoutingpolicyPolicyDefaultactionNextpolicy
	Nextpolicy *DeviceRoutingpolicyPolicyDefaultactionNextpolicy `json:"next-policy,omitempty"`
	//RootRoutingpolicyPolicyDefaultactionReject
	Reject *DeviceRoutingpolicyPolicyDefaultactionReject `json:"reject,omitempty"`
}

// DeviceRoutingpolicyPolicyDefaultactionAccept struct
type DeviceRoutingpolicyPolicyDefaultactionAccept struct {
	//RootRoutingpolicyPolicyDefaultactionAcceptBgp
	Bgp *DeviceRoutingpolicyPolicyDefaultactionAcceptBgp `json:"bgp,omitempty"`
}

// DeviceRoutingpolicyPolicyDefaultactionAcceptBgp struct
type DeviceRoutingpolicyPolicyDefaultactionAcceptBgp struct {
	//RootRoutingpolicyPolicyDefaultactionAcceptBgpAspath
	Aspath *DeviceRoutingpolicyPolicyDefaultactionAcceptBgpAspath `json:"as-path,omitempty"`
	//RootRoutingpolicyPolicyDefaultactionAcceptBgpCommunities
	Communities *DeviceRoutingpolicyPolicyDefaultactionAcceptBgpCommunities `json:"communities,omitempty"`
	//RootRoutingpolicyPolicyDefaultactionAcceptBgpLocalpreference
	Localpreference *DeviceRoutingpolicyPolicyDefaultactionAcceptBgpLocalpreference `json:"local-preference,omitempty"`
	//RootRoutingpolicyPolicyDefaultactionAcceptBgpOrigin
	Origin *DeviceRoutingpolicyPolicyDefaultactionAcceptBgpOrigin `json:"origin,omitempty"`
}

// DeviceRoutingpolicyPolicyDefaultactionAcceptBgpAspath struct
type DeviceRoutingpolicyPolicyDefaultactionAcceptBgpAspath struct {
	//RootRoutingpolicyPolicyDefaultactionAcceptBgpAspathPrepend
	Prepend *DeviceRoutingpolicyPolicyDefaultactionAcceptBgpAspathPrepend `json:"prepend,omitempty"`
	Remove  *bool                                                         `json:"remove,omitempty"`
	//RootRoutingpolicyPolicyDefaultactionAcceptBgpAspathReplace
	Replace *DeviceRoutingpolicyPolicyDefaultactionAcceptBgpAspathReplace `json:"replace,omitempty"`
}

// DeviceRoutingpolicyPolicyDefaultactionAcceptBgpAspathPrepend struct
type DeviceRoutingpolicyPolicyDefaultactionAcceptBgpAspathPrepend struct {
	Asnumber *string `json:"as-number,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=50
	// +kubebuilder:default:=1
	Repeatn *uint8 `json:"repeat-n,omitempty"`
}

// DeviceRoutingpolicyPolicyDefaultactionAcceptBgpAspathReplace struct
type DeviceRoutingpolicyPolicyDefaultactionAcceptBgpAspathReplace struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	Replace *uint32 `json:"replace,omitempty"`
}

// DeviceRoutingpolicyPolicyDefaultactionAcceptBgpCommunities struct
type DeviceRoutingpolicyPolicyDefaultactionAcceptBgpCommunities struct {
	Add     *string `json:"add,omitempty"`
	Remove  *string `json:"remove,omitempty"`
	Replace *string `json:"replace,omitempty"`
}

// DeviceRoutingpolicyPolicyDefaultactionAcceptBgpLocalpreference struct
type DeviceRoutingpolicyPolicyDefaultactionAcceptBgpLocalpreference struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	Set *uint32 `json:"set,omitempty"`
}

// DeviceRoutingpolicyPolicyDefaultactionAcceptBgpOrigin struct
type DeviceRoutingpolicyPolicyDefaultactionAcceptBgpOrigin struct {
	// +kubebuilder:validation:Enum=`egp`;`igp`;`incomplete`
	Set E_DeviceRoutingpolicyPolicyDefaultactionAcceptBgpOriginSet `json:"set,omitempty"`
	//Set *string `json:"set,omitempty"`
}

// DeviceRoutingpolicyPolicyDefaultactionNextentry struct
type DeviceRoutingpolicyPolicyDefaultactionNextentry struct {
}

// DeviceRoutingpolicyPolicyDefaultactionNextpolicy struct
type DeviceRoutingpolicyPolicyDefaultactionNextpolicy struct {
}

// DeviceRoutingpolicyPolicyDefaultactionReject struct
type DeviceRoutingpolicyPolicyDefaultactionReject struct {
}

// DeviceRoutingpolicyPolicyStatement struct
type DeviceRoutingpolicyPolicyStatement struct {
	//RootRoutingpolicyPolicyStatementAction
	Action *DeviceRoutingpolicyPolicyStatementAction `json:"action,omitempty"`
	//RootRoutingpolicyPolicyStatementMatch
	Match *DeviceRoutingpolicyPolicyStatementMatch `json:"match,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	Sequenceid *uint32 `json:"sequence-id"`
}

// DeviceRoutingpolicyPolicyStatementAction struct
type DeviceRoutingpolicyPolicyStatementAction struct {
	//RootRoutingpolicyPolicyStatementActionAccept
	Accept *DeviceRoutingpolicyPolicyStatementActionAccept `json:"accept,omitempty"`
	//RootRoutingpolicyPolicyStatementActionNextentry
	Nextentry *DeviceRoutingpolicyPolicyStatementActionNextentry `json:"next-entry,omitempty"`
	//RootRoutingpolicyPolicyStatementActionNextpolicy
	Nextpolicy *DeviceRoutingpolicyPolicyStatementActionNextpolicy `json:"next-policy,omitempty"`
	//RootRoutingpolicyPolicyStatementActionReject
	Reject *DeviceRoutingpolicyPolicyStatementActionReject `json:"reject,omitempty"`
}

// DeviceRoutingpolicyPolicyStatementActionAccept struct
type DeviceRoutingpolicyPolicyStatementActionAccept struct {
	//RootRoutingpolicyPolicyStatementActionAcceptBgp
	Bgp *DeviceRoutingpolicyPolicyStatementActionAcceptBgp `json:"bgp,omitempty"`
}

// DeviceRoutingpolicyPolicyStatementActionAcceptBgp struct
type DeviceRoutingpolicyPolicyStatementActionAcceptBgp struct {
	//RootRoutingpolicyPolicyStatementActionAcceptBgpAspath
	Aspath *DeviceRoutingpolicyPolicyStatementActionAcceptBgpAspath `json:"as-path,omitempty"`
	//RootRoutingpolicyPolicyStatementActionAcceptBgpCommunities
	Communities *DeviceRoutingpolicyPolicyStatementActionAcceptBgpCommunities `json:"communities,omitempty"`
	//RootRoutingpolicyPolicyStatementActionAcceptBgpLocalpreference
	Localpreference *DeviceRoutingpolicyPolicyStatementActionAcceptBgpLocalpreference `json:"local-preference,omitempty"`
	//RootRoutingpolicyPolicyStatementActionAcceptBgpOrigin
	Origin *DeviceRoutingpolicyPolicyStatementActionAcceptBgpOrigin `json:"origin,omitempty"`
}

// DeviceRoutingpolicyPolicyStatementActionAcceptBgpAspath struct
type DeviceRoutingpolicyPolicyStatementActionAcceptBgpAspath struct {
	//RootRoutingpolicyPolicyStatementActionAcceptBgpAspathPrepend
	Prepend *DeviceRoutingpolicyPolicyStatementActionAcceptBgpAspathPrepend `json:"prepend,omitempty"`
	Remove  *bool                                                           `json:"remove,omitempty"`
	//RootRoutingpolicyPolicyStatementActionAcceptBgpAspathReplace
	Replace *DeviceRoutingpolicyPolicyStatementActionAcceptBgpAspathReplace `json:"replace,omitempty"`
}

// DeviceRoutingpolicyPolicyStatementActionAcceptBgpAspathPrepend struct
type DeviceRoutingpolicyPolicyStatementActionAcceptBgpAspathPrepend struct {
	Asnumber *string `json:"as-number,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=50
	// +kubebuilder:default:=1
	Repeatn *uint8 `json:"repeat-n,omitempty"`
}

// DeviceRoutingpolicyPolicyStatementActionAcceptBgpAspathReplace struct
type DeviceRoutingpolicyPolicyStatementActionAcceptBgpAspathReplace struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	Replace *uint32 `json:"replace,omitempty"`
}

// DeviceRoutingpolicyPolicyStatementActionAcceptBgpCommunities struct
type DeviceRoutingpolicyPolicyStatementActionAcceptBgpCommunities struct {
	Add     *string `json:"add,omitempty"`
	Remove  *string `json:"remove,omitempty"`
	Replace *string `json:"replace,omitempty"`
}

// DeviceRoutingpolicyPolicyStatementActionAcceptBgpLocalpreference struct
type DeviceRoutingpolicyPolicyStatementActionAcceptBgpLocalpreference struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	Set *uint32 `json:"set,omitempty"`
}

// DeviceRoutingpolicyPolicyStatementActionAcceptBgpOrigin struct
type DeviceRoutingpolicyPolicyStatementActionAcceptBgpOrigin struct {
	// +kubebuilder:validation:Enum=`egp`;`igp`;`incomplete`
	Set E_DeviceRoutingpolicyPolicyStatementActionAcceptBgpOriginSet `json:"set,omitempty"`
	//Set *string `json:"set,omitempty"`
}

// DeviceRoutingpolicyPolicyStatementActionNextentry struct
type DeviceRoutingpolicyPolicyStatementActionNextentry struct {
}

// DeviceRoutingpolicyPolicyStatementActionNextpolicy struct
type DeviceRoutingpolicyPolicyStatementActionNextpolicy struct {
}

// DeviceRoutingpolicyPolicyStatementActionReject struct
type DeviceRoutingpolicyPolicyStatementActionReject struct {
}

// DeviceRoutingpolicyPolicyStatementMatch struct
type DeviceRoutingpolicyPolicyStatementMatch struct {
	//RootRoutingpolicyPolicyStatementMatchBgp
	Bgp    *DeviceRoutingpolicyPolicyStatementMatchBgp `json:"bgp,omitempty"`
	Family *string                                     `json:"family,omitempty"`
	//RootRoutingpolicyPolicyStatementMatchIsis
	Isis *DeviceRoutingpolicyPolicyStatementMatchIsis `json:"isis,omitempty"`
	//RootRoutingpolicyPolicyStatementMatchOspf
	Ospf      *DeviceRoutingpolicyPolicyStatementMatchOspf `json:"ospf,omitempty"`
	Prefixset *string                                      `json:"prefix-set,omitempty"`
	Protocol  *string                                      `json:"protocol,omitempty"`
}

// DeviceRoutingpolicyPolicyStatementMatchBgp struct
type DeviceRoutingpolicyPolicyStatementMatchBgp struct {
	//RootRoutingpolicyPolicyStatementMatchBgpAspathlength
	Aspathlength *DeviceRoutingpolicyPolicyStatementMatchBgpAspathlength `json:"as-path-length,omitempty"`
	Aspathset    *string                                                 `json:"as-path-set,omitempty"`
	Communityset *string                                                 `json:"community-set,omitempty"`
	//RootRoutingpolicyPolicyStatementMatchBgpEvpn
	Evpn *DeviceRoutingpolicyPolicyStatementMatchBgpEvpn `json:"evpn,omitempty"`
}

// DeviceRoutingpolicyPolicyStatementMatchBgpAspathlength struct
type DeviceRoutingpolicyPolicyStatementMatchBgpAspathlength struct {
	// +kubebuilder:validation:Enum=`eq`;`ge`;`le`
	// +kubebuilder:default:="eq"
	Operator E_DeviceRoutingpolicyPolicyStatementMatchBgpAspathlengthOperator `json:"operator,omitempty"`
	//Operator *string `json:"operator,omitempty"`
	// +kubebuilder:default:=false
	Unique *bool `json:"unique,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	Value *uint8 `json:"value"`
}

// DeviceRoutingpolicyPolicyStatementMatchBgpEvpn struct
type DeviceRoutingpolicyPolicyStatementMatchBgpEvpn struct {
	//RootRoutingpolicyPolicyStatementMatchBgpEvpnRoutetype
	Routetype *DeviceRoutingpolicyPolicyStatementMatchBgpEvpnRoutetype `json:"route-type,omitempty"`
}

// DeviceRoutingpolicyPolicyStatementMatchBgpEvpnRoutetype struct
type DeviceRoutingpolicyPolicyStatementMatchBgpEvpnRoutetype struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=5
	Routetype *uint8 `json:"route-type,omitempty"`
}

// DeviceRoutingpolicyPolicyStatementMatchIsis struct
type DeviceRoutingpolicyPolicyStatementMatchIsis struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=2
	Level *uint8 `json:"level,omitempty"`
	// +kubebuilder:validation:Enum=`external`;`internal`
	Routetype E_DeviceRoutingpolicyPolicyStatementMatchIsisRoutetype `json:"route-type,omitempty"`
	//Routetype *string `json:"route-type,omitempty"`
}

// DeviceRoutingpolicyPolicyStatementMatchOspf struct
type DeviceRoutingpolicyPolicyStatementMatchOspf struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|[0-9\.]*|(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])([\p{N}\p{L}]+)?`
	Areaid *string `json:"area-id,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	Instanceid *uint32 `json:"instance-id,omitempty"`
	Routetype  *string `json:"route-type,omitempty"`
}

// DeviceRoutingpolicyPrefixset struct
type DeviceRoutingpolicyPrefixset struct {
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Prefix []*DeviceRoutingpolicyPrefixsetPrefix `json:"prefix,omitempty"`
}

// DeviceRoutingpolicyPrefixsetPrefix struct
type DeviceRoutingpolicyPrefixsetPrefix struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	Ipprefix *string `json:"ip-prefix"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`([0-9]+\.\.[0-9]+)|exact`
	Masklengthrange *string `json:"mask-length-range"`
}

// DeviceSystem struct
type DeviceSystem struct {
	//RootSystemAaa
	Aaa *DeviceSystemAaa `json:"aaa,omitempty"`
	//RootSystemAuthentication
	Authentication *DeviceSystemAuthentication `json:"authentication,omitempty"`
	//RootSystemBanner
	Banner *DeviceSystemBanner `json:"banner,omitempty"`
	//RootSystemBoot
	Boot *DeviceSystemBoot `json:"boot,omitempty"`
	//RootSystemBridgetable
	Bridgetable *DeviceSystemBridgetable `json:"bridge-table,omitempty"`
	//RootSystemClock
	Clock *DeviceSystemClock `json:"clock,omitempty"`
	//RootSystemConfiguration
	Configuration *DeviceSystemConfiguration `json:"configuration,omitempty"`
	//RootSystemDhcpserver
	Dhcpserver *DeviceSystemDhcpserver `json:"dhcp-server,omitempty"`
	//RootSystemDns
	Dns *DeviceSystemDns `json:"dns,omitempty"`
	//RootSystemFtpserver
	Ftpserver *DeviceSystemFtpserver `json:"ftp-server,omitempty"`
	//RootSystemGnmiserver
	Gnmiserver *DeviceSystemGnmiserver `json:"gnmi-server,omitempty"`
	//RootSystemGribiserver
	Gribiserver *DeviceSystemGribiserver `json:"gribi-server,omitempty"`
	//RootSystemInformation
	Information *DeviceSystemInformation `json:"information,omitempty"`
	//RootSystemJsonrpcserver
	Jsonrpcserver *DeviceSystemJsonrpcserver `json:"json-rpc-server,omitempty"`
	//RootSystemLacp
	Lacp *DeviceSystemLacp `json:"lacp,omitempty"`
	//RootSystemLldp
	Lldp *DeviceSystemLldp `json:"lldp,omitempty"`
	//RootSystemLoadbalancing
	Loadbalancing *DeviceSystemLoadbalancing `json:"load-balancing,omitempty"`
	//RootSystemLogging
	Logging *DeviceSystemLogging `json:"logging,omitempty"`
	//RootSystemMaintenance
	Maintenance *DeviceSystemMaintenance `json:"maintenance,omitempty"`
	//RootSystemMirroring
	Mirroring *DeviceSystemMirroring `json:"mirroring,omitempty"`
	//RootSystemMpls
	Mpls *DeviceSystemMpls `json:"mpls,omitempty"`
	//RootSystemMtu
	Mtu *DeviceSystemMtu `json:"mtu,omitempty"`
	//RootSystemName
	Name *DeviceSystemName `json:"name,omitempty"`
	//RootSystemNetworkinstance
	Networkinstance *DeviceSystemNetworkinstance `json:"network-instance,omitempty"`
	//RootSystemNtp
	Ntp *DeviceSystemNtp `json:"ntp,omitempty"`
	//RootSystemP4rtserver
	P4rtserver *DeviceSystemP4rtserver `json:"p4rt-server,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=64
	Raguardpolicy []*DeviceSystemRaguardpolicy `json:"ra-guard-policy,omitempty"`
	//RootSystemSflow
	Sflow *DeviceSystemSflow `json:"sflow,omitempty"`
	//RootSystemSnmp
	Snmp *DeviceSystemSnmp `json:"snmp,omitempty"`
	//RootSystemSshserver
	Sshserver *DeviceSystemSshserver `json:"ssh-server,omitempty"`
	//RootSystemSync
	Sync *DeviceSystemSync `json:"sync,omitempty"`
	//RootSystemTls
	Tls *DeviceSystemTls `json:"tls,omitempty"`
	//RootSystemTraceoptions
	Traceoptions *DeviceSystemTraceoptions `json:"trace-options,omitempty"`
	//RootSystemWarmreboot
	Warmreboot *DeviceSystemWarmreboot `json:"warm-reboot,omitempty"`
}

// DeviceSystemAaa struct
type DeviceSystemAaa struct {
	//RootSystemAaaAccounting
	Accounting *DeviceSystemAaaAccounting `json:"accounting,omitempty"`
	//RootSystemAaaAuthentication
	Authentication *DeviceSystemAaaAuthentication `json:"authentication,omitempty"`
	//RootSystemAaaAuthorization
	Authorization *DeviceSystemAaaAuthorization `json:"authorization,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=2
	Servergroup []*DeviceSystemAaaServergroup `json:"server-group,omitempty"`
}

// DeviceSystemAaaAccounting struct
type DeviceSystemAaaAccounting struct {
	//RootSystemAaaAccountingAccountingmethod
	Accountingmethod *DeviceSystemAaaAccountingAccountingmethod `json:"accounting-method,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Event []*DeviceSystemAaaAccountingEvent `json:"event,omitempty"`
}

// DeviceSystemAaaAccountingAccountingmethod struct
type DeviceSystemAaaAccountingAccountingmethod struct {
	Accountingmethod *string `json:"accounting-method,omitempty"`
}

// DeviceSystemAaaAccountingEvent struct
type DeviceSystemAaaAccountingEvent struct {
	Eventtype *string `json:"event-type"`
	Record    *string `json:"record,omitempty"`
}

// DeviceSystemAaaAuthentication struct
type DeviceSystemAaaAuthentication struct {
	//RootSystemAaaAuthenticationAdminuser
	Adminuser *DeviceSystemAaaAuthenticationAdminuser `json:"admin-user,omitempty"`
	//RootSystemAaaAuthenticationAuthenticationmethod
	Authenticationmethod *DeviceSystemAaaAuthenticationAuthenticationmethod `json:"authentication-method,omitempty"`
	// +kubebuilder:default:=false
	Exitonreject *bool `json:"exit-on-reject,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	// +kubebuilder:default:=600
	Idletimeout *uint32 `json:"idle-timeout,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=128
	User []*DeviceSystemAaaAuthenticationUser `json:"user,omitempty"`
}

// DeviceSystemAaaAuthenticationAdminuser struct
type DeviceSystemAaaAuthenticationAdminuser struct {
	Password *string `json:"password,omitempty"`
}

// DeviceSystemAaaAuthenticationAuthenticationmethod struct
type DeviceSystemAaaAuthenticationAuthenticationmethod struct {
	Authenticationmethod *string `json:"authentication-method,omitempty"`
}

// DeviceSystemAaaAuthenticationUser struct
type DeviceSystemAaaAuthenticationUser struct {
	Password *string `json:"password,omitempty"`
	//RootSystemAaaAuthenticationUserRole
	Role *DeviceSystemAaaAuthenticationUserRole `json:"role,omitempty"`
	//RootSystemAaaAuthenticationUserSshkey
	Sshkey *DeviceSystemAaaAuthenticationUserSshkey `json:"ssh-key,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=32
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Username *string `json:"username"`
}

// DeviceSystemAaaAuthenticationUserRole struct
type DeviceSystemAaaAuthenticationUserRole struct {
	Role *string `json:"role,omitempty"`
}

// DeviceSystemAaaAuthenticationUserSshkey struct
type DeviceSystemAaaAuthenticationUserSshkey struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`ssh-rsa .*`
	Sshkey *string `json:"ssh-key,omitempty"`
}

// DeviceSystemAaaAuthorization struct
type DeviceSystemAaaAuthorization struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Role []*DeviceSystemAaaAuthorizationRole `json:"role,omitempty"`
}

// DeviceSystemAaaAuthorizationRole struct
type DeviceSystemAaaAuthorizationRole struct {
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=32
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Rolename *string `json:"rolename"`
	//RootSystemAaaAuthorizationRoleServices
	Services *DeviceSystemAaaAuthorizationRoleServices `json:"services,omitempty"`
	//RootSystemAaaAuthorizationRoleTacacs
	Tacacs *DeviceSystemAaaAuthorizationRoleTacacs `json:"tacacs,omitempty"`
}

// DeviceSystemAaaAuthorizationRoleServices struct
type DeviceSystemAaaAuthorizationRoleServices struct {
	// +kubebuilder:validation:Enum=`cli`;`ftp`;`gnmi`;`gribi`;`json-rpc`;`p4rt`
	Services E_DeviceSystemAaaAuthorizationRoleServicesServices `json:"services,omitempty"`
	//Services *string `json:"services,omitempty"`
}

// DeviceSystemAaaAuthorizationRoleTacacs struct
type DeviceSystemAaaAuthorizationRoleTacacs struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=15
	Privlvl *uint8 `json:"priv-lvl,omitempty"`
}

// DeviceSystemAaaServergroup struct
type DeviceSystemAaaServergroup struct {
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
	// +kubebuilder:default:=false
	Privlvlauthorization *bool `json:"priv-lvl-authorization,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=5
	Server []*DeviceSystemAaaServergroupServer `json:"server,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=3600
	// +kubebuilder:default:=10
	Timeout *uint16 `json:"timeout,omitempty"`
	Type    *string `json:"type"`
}

// DeviceSystemAaaServergroupServer struct
type DeviceSystemAaaServergroupServer struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Address *string `json:"address"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name            *string `json:"name,omitempty"`
	Networkinstance *string `json:"network-instance,omitempty"`
	//RootSystemAaaServergroupServerTacacs
	Tacacs *DeviceSystemAaaServergroupServerTacacs `json:"tacacs,omitempty"`
}

// DeviceSystemAaaServergroupServerTacacs struct
type DeviceSystemAaaServergroupServerTacacs struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=49
	Port      *uint16 `json:"port,omitempty"`
	Secretkey *string `json:"secret-key,omitempty"`
}

// DeviceSystemAuthentication struct
type DeviceSystemAuthentication struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Keychain []*DeviceSystemAuthenticationKeychain `json:"keychain,omitempty"`
}

// DeviceSystemAuthenticationKeychain struct
type DeviceSystemAuthenticationKeychain struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_DeviceSystemAuthenticationKeychainAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Description *string `json:"description,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1
	Key []*DeviceSystemAuthenticationKeychainKey `json:"key,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
	// +kubebuilder:validation:Enum=`isis`;`ospf`;`tcp-ao`;`tcp-md5`;`vrrp`
	Type E_DeviceSystemAuthenticationKeychainType `json:"type,omitempty"`
	//Type *string `json:"type,omitempty"`
}

// DeviceSystemAuthenticationKeychainKey struct
type DeviceSystemAuthenticationKeychainKey struct {
	// +kubebuilder:validation:Enum=`aes-128-cmac`;`cleartext`;`hmac-md5`;`hmac-sha-1`;`hmac-sha-256`;`md5`
	Algorithm E_DeviceSystemAuthenticationKeychainKeyAlgorithm `json:"algorithm,omitempty"`
	//Algorithm *string `json:"algorithm,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=25
	Authenticationkey *string `json:"authentication-key,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	Index *uint8 `json:"index"`
}

// DeviceSystemBanner struct
type DeviceSystemBanner struct {
	Loginbanner *string `json:"login-banner,omitempty"`
	Motdbanner  *string `json:"motd-banner,omitempty"`
}

// DeviceSystemBoot struct
type DeviceSystemBoot struct {
	//RootSystemBootAutoboot
	Autoboot *DeviceSystemBootAutoboot `json:"autoboot,omitempty"`
}

// DeviceSystemBootAutoboot struct
type DeviceSystemBootAutoboot struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	Adminstate E_DeviceSystemBootAutobootAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=10
	Attempts *uint8 `json:"attempts,omitempty"`
	// +kubebuilder:validation:Enum=`serial`
	Clientid E_DeviceSystemBootAutobootClientid `json:"client-id,omitempty"`
	//Clientid *string `json:"client-id,omitempty"`
	// +kubebuilder:default:="mgmt0"
	Interface *string `json:"interface,omitempty"`
	// kubebuilder:validation:Minimum=200
	// kubebuilder:validation:Maximum=3600
	Timeout *uint32 `json:"timeout,omitempty"`
}

// DeviceSystemBridgetable struct
type DeviceSystemBridgetable struct {
	//RootSystemBridgetableMaclearning
	Maclearning *DeviceSystemBridgetableMaclearning `json:"mac-learning,omitempty"`
	//RootSystemBridgetableMaclimit
	Maclimit *DeviceSystemBridgetableMaclimit `json:"mac-limit,omitempty"`
	//RootSystemBridgetableProxyarp
	Proxyarp *DeviceSystemBridgetableProxyarp `json:"proxy-arp,omitempty"`
}

// DeviceSystemBridgetableMaclearning struct
type DeviceSystemBridgetableMaclearning struct {
}

// DeviceSystemBridgetableMaclimit struct
type DeviceSystemBridgetableMaclimit struct {
}

// DeviceSystemBridgetableProxyarp struct
type DeviceSystemBridgetableProxyarp struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=156000
	// +kubebuilder:default:=156000
	Tablesize *int32 `json:"table-size,omitempty"`
}

// DeviceSystemClock struct
type DeviceSystemClock struct {
	// +kubebuilder:validation:Enum=`Africa/Abidjan`;`Africa/Accra`;`Africa/Addis_Ababa`;`Africa/Algiers`;`Africa/Asmara`;`Africa/Bamako`;`Africa/Bangui`;`Africa/Banjul`;`Africa/Bissau`;`Africa/Blantyre`;`Africa/Brazzaville`;`Africa/Bujumbura`;`Africa/Cairo`;`Africa/Casablanca`;`Africa/Ceuta`;`Africa/Conakry`;`Africa/Dakar`;`Africa/Dar_es_Salaam`;`Africa/Djibouti`;`Africa/Douala`;`Africa/El_Aaiun`;`Africa/Freetown`;`Africa/Gaborone`;`Africa/Harare`;`Africa/Johannesburg`;`Africa/Juba`;`Africa/Kampala`;`Africa/Khartoum`;`Africa/Kigali`;`Africa/Kinshasa`;`Africa/Lagos`;`Africa/Libreville`;`Africa/Lome`;`Africa/Luanda`;`Africa/Lubumbashi`;`Africa/Lusaka`;`Africa/Malabo`;`Africa/Maputo`;`Africa/Maseru`;`Africa/Mbabane`;`Africa/Mogadishu`;`Africa/Monrovia`;`Africa/Nairobi`;`Africa/Ndjamena`;`Africa/Niamey`;`Africa/Nouakchott`;`Africa/Ouagadougou`;`Africa/Porto-Novo`;`Africa/Sao_Tome`;`Africa/Tripoli`;`Africa/Tunis`;`Africa/Windhoek`;`America/Adak`;`America/Anchorage`;`America/Anguilla`;`America/Antigua`;`America/Araguaina`;`America/Argentina/Buenos_Aires`;`America/Argentina/Catamarca`;`America/Argentina/Cordoba`;`America/Argentina/Jujuy`;`America/Argentina/La_Rioja`;`America/Argentina/Mendoza`;`America/Argentina/Rio_Gallegos`;`America/Argentina/Salta`;`America/Argentina/San_Juan`;`America/Argentina/San_Luis`;`America/Argentina/Tucuman`;`America/Argentina/Ushuaia`;`America/Aruba`;`America/Asuncion`;`America/Atikokan`;`America/Bahia`;`America/Bahia_Banderas`;`America/Barbados`;`America/Belem`;`America/Belize`;`America/Blanc-Sablon`;`America/Boa_Vista`;`America/Bogota`;`America/Boise`;`America/Cambridge_Bay`;`America/Campo_Grande`;`America/Cancun`;`America/Caracas`;`America/Cayenne`;`America/Cayman`;`America/Chicago`;`America/Chihuahua`;`America/Costa_Rica`;`America/Creston`;`America/Cuiaba`;`America/Curacao`;`America/Danmarkshavn`;`America/Dawson`;`America/Dawson_Creek`;`America/Denver`;`America/Detroit`;`America/Dominica`;`America/Edmonton`;`America/Eirunepe`;`America/El_Salvador`;`America/Fort_Nelson`;`America/Fortaleza`;`America/Glace_Bay`;`America/Godthab`;`America/Goose_Bay`;`America/Grand_Turk`;`America/Grenada`;`America/Guadeloupe`;`America/Guatemala`;`America/Guayaquil`;`America/Guyana`;`America/Halifax`;`America/Havana`;`America/Hermosillo`;`America/Indiana/Indianapolis`;`America/Indiana/Knox`;`America/Indiana/Marengo`;`America/Indiana/Petersburg`;`America/Indiana/Tell_City`;`America/Indiana/Vevay`;`America/Indiana/Vincennes`;`America/Indiana/Winamac`;`America/Inuvik`;`America/Iqaluit`;`America/Jamaica`;`America/Juneau`;`America/Kentucky/Louisville`;`America/Kentucky/Monticello`;`America/Kralendijk`;`America/La_Paz`;`America/Lima`;`America/Los_Angeles`;`America/Lower_Princes`;`America/Maceio`;`America/Managua`;`America/Manaus`;`America/Marigot`;`America/Martinique`;`America/Matamoros`;`America/Mazatlan`;`America/Menominee`;`America/Merida`;`America/Metlakatla`;`America/Mexico_City`;`America/Miquelon`;`America/Moncton`;`America/Monterrey`;`America/Montevideo`;`America/Montserrat`;`America/Nassau`;`America/New_York`;`America/Nipigon`;`America/Nome`;`America/Noronha`;`America/North_Dakota/Beulah`;`America/North_Dakota/Center`;`America/North_Dakota/New_Salem`;`America/Ojinaga`;`America/Panama`;`America/Pangnirtung`;`America/Paramaribo`;`America/Phoenix`;`America/Port-au-Prince`;`America/Port_of_Spain`;`America/Porto_Velho`;`America/Puerto_Rico`;`America/Punta_Arenas`;`America/Rainy_River`;`America/Rankin_Inlet`;`America/Recife`;`America/Regina`;`America/Resolute`;`America/Rio_Branco`;`America/Santarem`;`America/Santiago`;`America/Santo_Domingo`;`America/Sao_Paulo`;`America/Scoresbysund`;`America/Sitka`;`America/St_Barthelemy`;`America/St_Johns`;`America/St_Kitts`;`America/St_Lucia`;`America/St_Thomas`;`America/St_Vincent`;`America/Swift_Current`;`America/Tegucigalpa`;`America/Thule`;`America/Thunder_Bay`;`America/Tijuana`;`America/Toronto`;`America/Tortola`;`America/Vancouver`;`America/Whitehorse`;`America/Winnipeg`;`America/Yakutat`;`America/Yellowknife`;`Antarctica/Casey`;`Antarctica/Davis`;`Antarctica/DumontDUrville`;`Antarctica/Macquarie`;`Antarctica/Mawson`;`Antarctica/McMurdo`;`Antarctica/Palmer`;`Antarctica/Rothera`;`Antarctica/Syowa`;`Antarctica/Troll`;`Antarctica/Vostok`;`Arctic/Longyearbyen`;`Asia/Aden`;`Asia/Almaty`;`Asia/Amman`;`Asia/Anadyr`;`Asia/Aqtau`;`Asia/Aqtobe`;`Asia/Ashgabat`;`Asia/Atyrau`;`Asia/Baghdad`;`Asia/Bahrain`;`Asia/Baku`;`Asia/Bangkok`;`Asia/Barnaul`;`Asia/Beirut`;`Asia/Bishkek`;`Asia/Brunei`;`Asia/Chita`;`Asia/Choibalsan`;`Asia/Colombo`;`Asia/Damascus`;`Asia/Dhaka`;`Asia/Dili`;`Asia/Dubai`;`Asia/Dushanbe`;`Asia/Famagusta`;`Asia/Gaza`;`Asia/Hebron`;`Asia/Ho_Chi_Minh`;`Asia/Hong_Kong`;`Asia/Hovd`;`Asia/Irkutsk`;`Asia/Jakarta`;`Asia/Jayapura`;`Asia/Jerusalem`;`Asia/Kabul`;`Asia/Kamchatka`;`Asia/Karachi`;`Asia/Kathmandu`;`Asia/Khandyga`;`Asia/Kolkata`;`Asia/Krasnoyarsk`;`Asia/Kuala_Lumpur`;`Asia/Kuching`;`Asia/Kuwait`;`Asia/Macau`;`Asia/Magadan`;`Asia/Makassar`;`Asia/Manila`;`Asia/Muscat`;`Asia/Nicosia`;`Asia/Novokuznetsk`;`Asia/Novosibirsk`;`Asia/Omsk`;`Asia/Oral`;`Asia/Phnom_Penh`;`Asia/Pontianak`;`Asia/Pyongyang`;`Asia/Qatar`;`Asia/Qostanay`;`Asia/Qyzylorda`;`Asia/Riyadh`;`Asia/Sakhalin`;`Asia/Samarkand`;`Asia/Seoul`;`Asia/Shanghai`;`Asia/Singapore`;`Asia/Srednekolymsk`;`Asia/Taipei`;`Asia/Tashkent`;`Asia/Tbilisi`;`Asia/Tehran`;`Asia/Thimphu`;`Asia/Tokyo`;`Asia/Tomsk`;`Asia/Ulaanbaatar`;`Asia/Urumqi`;`Asia/Ust-Nera`;`Asia/Vientiane`;`Asia/Vladivostok`;`Asia/Yakutsk`;`Asia/Yangon`;`Asia/Yekaterinburg`;`Asia/Yerevan`;`Atlantic/Azores`;`Atlantic/Bermuda`;`Atlantic/Canary`;`Atlantic/Cape_Verde`;`Atlantic/Faroe`;`Atlantic/Madeira`;`Atlantic/Reykjavik`;`Atlantic/South_Georgia`;`Atlantic/St_Helena`;`Atlantic/Stanley`;`Australia/Adelaide`;`Australia/Brisbane`;`Australia/Broken_Hill`;`Australia/Currie`;`Australia/Darwin`;`Australia/Eucla`;`Australia/Hobart`;`Australia/Lindeman`;`Australia/Lord_Howe`;`Australia/Melbourne`;`Australia/Perth`;`Australia/Sydney`;`Europe/Amsterdam`;`Europe/Andorra`;`Europe/Astrakhan`;`Europe/Athens`;`Europe/Belgrade`;`Europe/Berlin`;`Europe/Bratislava`;`Europe/Brussels`;`Europe/Bucharest`;`Europe/Budapest`;`Europe/Busingen`;`Europe/Chisinau`;`Europe/Copenhagen`;`Europe/Dublin`;`Europe/Gibraltar`;`Europe/Guernsey`;`Europe/Helsinki`;`Europe/Isle_of_Man`;`Europe/Istanbul`;`Europe/Jersey`;`Europe/Kaliningrad`;`Europe/Kiev`;`Europe/Kirov`;`Europe/Lisbon`;`Europe/Ljubljana`;`Europe/London`;`Europe/Luxembourg`;`Europe/Madrid`;`Europe/Malta`;`Europe/Mariehamn`;`Europe/Minsk`;`Europe/Monaco`;`Europe/Moscow`;`Europe/Oslo`;`Europe/Paris`;`Europe/Podgorica`;`Europe/Prague`;`Europe/Riga`;`Europe/Rome`;`Europe/Samara`;`Europe/San_Marino`;`Europe/Sarajevo`;`Europe/Saratov`;`Europe/Simferopol`;`Europe/Skopje`;`Europe/Sofia`;`Europe/Stockholm`;`Europe/Tallinn`;`Europe/Tirane`;`Europe/Ulyanovsk`;`Europe/Uzhgorod`;`Europe/Vaduz`;`Europe/Vatican`;`Europe/Vienna`;`Europe/Vilnius`;`Europe/Volgograd`;`Europe/Warsaw`;`Europe/Zagreb`;`Europe/Zaporozhye`;`Europe/Zurich`;`Indian/Antananarivo`;`Indian/Chagos`;`Indian/Christmas`;`Indian/Cocos`;`Indian/Comoro`;`Indian/Kerguelen`;`Indian/Mahe`;`Indian/Maldives`;`Indian/Mauritius`;`Indian/Mayotte`;`Indian/Reunion`;`Pacific/Apia`;`Pacific/Auckland`;`Pacific/Bougainville`;`Pacific/Chatham`;`Pacific/Chuuk`;`Pacific/Easter`;`Pacific/Efate`;`Pacific/Enderbury`;`Pacific/Fakaofo`;`Pacific/Fiji`;`Pacific/Funafuti`;`Pacific/Galapagos`;`Pacific/Gambier`;`Pacific/Guadalcanal`;`Pacific/Guam`;`Pacific/Honolulu`;`Pacific/Kiritimati`;`Pacific/Kosrae`;`Pacific/Kwajalein`;`Pacific/Majuro`;`Pacific/Marquesas`;`Pacific/Midway`;`Pacific/Nauru`;`Pacific/Niue`;`Pacific/Norfolk`;`Pacific/Noumea`;`Pacific/Pago_Pago`;`Pacific/Palau`;`Pacific/Pitcairn`;`Pacific/Pohnpei`;`Pacific/Port_Moresby`;`Pacific/Rarotonga`;`Pacific/Saipan`;`Pacific/Tahiti`;`Pacific/Tarawa`;`Pacific/Tongatapu`;`Pacific/Wake`;`Pacific/Wallis`;`UTC`
	Timezone E_DeviceSystemClockTimezone `json:"timezone,omitempty"`
	//Timezone *string `json:"timezone,omitempty"`
}

// DeviceSystemConfiguration struct
type DeviceSystemConfiguration struct {
	// +kubebuilder:default:=false
	Autocheckpoint *bool `json:"auto-checkpoint,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=10080
	Idletimeout *uint16 `json:"idle-timeout,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=10
	Maxcandidates *uint8 `json:"max-candidates,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=10
	Maxcheckpoints *uint8 `json:"max-checkpoints,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=32
	Role []*DeviceSystemConfigurationRole `json:"role,omitempty"`
}

// DeviceSystemConfigurationRole struct
type DeviceSystemConfigurationRole struct {
	Name *string `json:"name"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=256
	Rule []*DeviceSystemConfigurationRoleRule `json:"rule,omitempty"`
}

// DeviceSystemConfigurationRoleRule struct
type DeviceSystemConfigurationRoleRule struct {
	// +kubebuilder:validation:Enum=`deny`;`read`;`write`
	Action        E_DeviceSystemConfigurationRoleRuleAction `json:"action,omitempty"`
	Pathreference *string                                   `json:"path-reference"`
}

// DeviceSystemDhcpserver struct
type DeviceSystemDhcpserver struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_DeviceSystemDhcpserverAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Networkinstance []*DeviceSystemDhcpserverNetworkinstance `json:"network-instance,omitempty"`
}

// DeviceSystemDhcpserverNetworkinstance struct
type DeviceSystemDhcpserverNetworkinstance struct {
	//RootSystemDhcpserverNetworkinstanceDhcpv4
	Dhcpv4 *DeviceSystemDhcpserverNetworkinstanceDhcpv4 `json:"dhcpv4,omitempty"`
	//RootSystemDhcpserverNetworkinstanceDhcpv6
	Dhcpv6 *DeviceSystemDhcpserverNetworkinstanceDhcpv6 `json:"dhcpv6,omitempty"`
	Name   *string                                      `json:"name"`
}

// DeviceSystemDhcpserverNetworkinstanceDhcpv4 struct
type DeviceSystemDhcpserverNetworkinstanceDhcpv4 struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_DeviceSystemDhcpserverNetworkinstanceDhcpv4Adminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	//RootSystemDhcpserverNetworkinstanceDhcpv4Options
	Options *DeviceSystemDhcpserverNetworkinstanceDhcpv4Options `json:"options,omitempty"`
	//RootSystemDhcpserverNetworkinstanceDhcpv4Staticallocation
	Staticallocation *DeviceSystemDhcpserverNetworkinstanceDhcpv4Staticallocation `json:"static-allocation,omitempty"`
	//RootSystemDhcpserverNetworkinstanceDhcpv4Traceoptions
	Traceoptions *DeviceSystemDhcpserverNetworkinstanceDhcpv4Traceoptions `json:"trace-options,omitempty"`
}

// DeviceSystemDhcpserverNetworkinstanceDhcpv4Options struct
type DeviceSystemDhcpserverNetworkinstanceDhcpv4Options struct {
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=128
	Bootfilename *string `json:"bootfile-name,omitempty"`
	//RootSystemDhcpserverNetworkinstanceDhcpv4OptionsDnsserver
	Dnsserver *DeviceSystemDhcpserverNetworkinstanceDhcpv4OptionsDnsserver `json:"dns-server,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=253
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((([a-zA-Z0-9_]([a-zA-Z0-9\-_]){0,61})?[a-zA-Z0-9]\.)*([a-zA-Z0-9_]([a-zA-Z0-9\-_]){0,61})?[a-zA-Z0-9]\.?)|\.`
	Domainname *string `json:"domain-name,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=63
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([a-zA-Z]|[a-zA-Z][a-zA-Z0-9._-]*[a-zA-Z0-9]))*([A-Za-z]|[A-Za-z][A-Za-z0-9._-]*[A-Za-z0-9])`
	Hostname *string `json:"hostname,omitempty"`
	//RootSystemDhcpserverNetworkinstanceDhcpv4OptionsNtpserver
	Ntpserver *DeviceSystemDhcpserverNetworkinstanceDhcpv4OptionsNtpserver `json:"ntp-server,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Router *string `json:"router,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Serverid *string `json:"server-id,omitempty"`
}

// DeviceSystemDhcpserverNetworkinstanceDhcpv4OptionsDnsserver struct
type DeviceSystemDhcpserverNetworkinstanceDhcpv4OptionsDnsserver struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Dnsserver *string `json:"dns-server,omitempty"`
}

// DeviceSystemDhcpserverNetworkinstanceDhcpv4OptionsNtpserver struct
type DeviceSystemDhcpserverNetworkinstanceDhcpv4OptionsNtpserver struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Ntpserver *string `json:"ntp-server,omitempty"`
}

// DeviceSystemDhcpserverNetworkinstanceDhcpv4Staticallocation struct
type DeviceSystemDhcpserverNetworkinstanceDhcpv4Staticallocation struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Host []*DeviceSystemDhcpserverNetworkinstanceDhcpv4StaticallocationHost `json:"host,omitempty"`
}

// DeviceSystemDhcpserverNetworkinstanceDhcpv4StaticallocationHost struct
type DeviceSystemDhcpserverNetworkinstanceDhcpv4StaticallocationHost struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))`
	Ipaddress *string `json:"ip-address,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}`
	Mac *string `json:"mac"`
	//RootSystemDhcpserverNetworkinstanceDhcpv4StaticallocationHostOptions
	Options *DeviceSystemDhcpserverNetworkinstanceDhcpv4StaticallocationHostOptions `json:"options,omitempty"`
}

// DeviceSystemDhcpserverNetworkinstanceDhcpv4StaticallocationHostOptions struct
type DeviceSystemDhcpserverNetworkinstanceDhcpv4StaticallocationHostOptions struct {
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=128
	Bootfilename *string `json:"bootfile-name,omitempty"`
	//RootSystemDhcpserverNetworkinstanceDhcpv4StaticallocationHostOptionsDnsserver
	Dnsserver *DeviceSystemDhcpserverNetworkinstanceDhcpv4StaticallocationHostOptionsDnsserver `json:"dns-server,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=253
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((([a-zA-Z0-9_]([a-zA-Z0-9\-_]){0,61})?[a-zA-Z0-9]\.)*([a-zA-Z0-9_]([a-zA-Z0-9\-_]){0,61})?[a-zA-Z0-9]\.?)|\.`
	Domainname *string `json:"domain-name,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=63
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([a-zA-Z]|[a-zA-Z][a-zA-Z0-9._-]*[a-zA-Z0-9]))*([A-Za-z]|[A-Za-z][A-Za-z0-9._-]*[A-Za-z0-9])`
	Hostname *string `json:"hostname,omitempty"`
	//RootSystemDhcpserverNetworkinstanceDhcpv4StaticallocationHostOptionsNtpserver
	Ntpserver *DeviceSystemDhcpserverNetworkinstanceDhcpv4StaticallocationHostOptionsNtpserver `json:"ntp-server,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Router *string `json:"router,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Serverid *string `json:"server-id,omitempty"`
}

// DeviceSystemDhcpserverNetworkinstanceDhcpv4StaticallocationHostOptionsDnsserver struct
type DeviceSystemDhcpserverNetworkinstanceDhcpv4StaticallocationHostOptionsDnsserver struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Dnsserver *string `json:"dns-server,omitempty"`
}

// DeviceSystemDhcpserverNetworkinstanceDhcpv4StaticallocationHostOptionsNtpserver struct
type DeviceSystemDhcpserverNetworkinstanceDhcpv4StaticallocationHostOptionsNtpserver struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Ntpserver *string `json:"ntp-server,omitempty"`
}

// DeviceSystemDhcpserverNetworkinstanceDhcpv4Traceoptions struct
type DeviceSystemDhcpserverNetworkinstanceDhcpv4Traceoptions struct {
	//RootSystemDhcpserverNetworkinstanceDhcpv4TraceoptionsTrace
	Trace *DeviceSystemDhcpserverNetworkinstanceDhcpv4TraceoptionsTrace `json:"trace,omitempty"`
}

// DeviceSystemDhcpserverNetworkinstanceDhcpv4TraceoptionsTrace struct
type DeviceSystemDhcpserverNetworkinstanceDhcpv4TraceoptionsTrace struct {
	// +kubebuilder:validation:Enum=`messages`
	Trace E_DeviceSystemDhcpserverNetworkinstanceDhcpv4TraceoptionsTraceTrace `json:"trace,omitempty"`
	//Trace *string `json:"trace,omitempty"`
}

// DeviceSystemDhcpserverNetworkinstanceDhcpv6 struct
type DeviceSystemDhcpserverNetworkinstanceDhcpv6 struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_DeviceSystemDhcpserverNetworkinstanceDhcpv6Adminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	//RootSystemDhcpserverNetworkinstanceDhcpv6Options
	Options *DeviceSystemDhcpserverNetworkinstanceDhcpv6Options `json:"options,omitempty"`
	//RootSystemDhcpserverNetworkinstanceDhcpv6Staticallocation
	Staticallocation *DeviceSystemDhcpserverNetworkinstanceDhcpv6Staticallocation `json:"static-allocation,omitempty"`
	//RootSystemDhcpserverNetworkinstanceDhcpv6Traceoptions
	Traceoptions *DeviceSystemDhcpserverNetworkinstanceDhcpv6Traceoptions `json:"trace-options,omitempty"`
}

// DeviceSystemDhcpserverNetworkinstanceDhcpv6Options struct
type DeviceSystemDhcpserverNetworkinstanceDhcpv6Options struct {
	//RootSystemDhcpserverNetworkinstanceDhcpv6OptionsDnsserver
	Dnsserver *DeviceSystemDhcpserverNetworkinstanceDhcpv6OptionsDnsserver `json:"dns-server,omitempty"`
}

// DeviceSystemDhcpserverNetworkinstanceDhcpv6OptionsDnsserver struct
type DeviceSystemDhcpserverNetworkinstanceDhcpv6OptionsDnsserver struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Dnsserver *string `json:"dns-server,omitempty"`
}

// DeviceSystemDhcpserverNetworkinstanceDhcpv6Staticallocation struct
type DeviceSystemDhcpserverNetworkinstanceDhcpv6Staticallocation struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Host []*DeviceSystemDhcpserverNetworkinstanceDhcpv6StaticallocationHost `json:"host,omitempty"`
}

// DeviceSystemDhcpserverNetworkinstanceDhcpv6StaticallocationHost struct
type DeviceSystemDhcpserverNetworkinstanceDhcpv6StaticallocationHost struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	Ipaddress *string `json:"ip-address,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}`
	Mac *string `json:"mac"`
	//RootSystemDhcpserverNetworkinstanceDhcpv6StaticallocationHostOptions
	Options *DeviceSystemDhcpserverNetworkinstanceDhcpv6StaticallocationHostOptions `json:"options,omitempty"`
}

// DeviceSystemDhcpserverNetworkinstanceDhcpv6StaticallocationHostOptions struct
type DeviceSystemDhcpserverNetworkinstanceDhcpv6StaticallocationHostOptions struct {
	//RootSystemDhcpserverNetworkinstanceDhcpv6StaticallocationHostOptionsDnsserver
	Dnsserver *DeviceSystemDhcpserverNetworkinstanceDhcpv6StaticallocationHostOptionsDnsserver `json:"dns-server,omitempty"`
}

// DeviceSystemDhcpserverNetworkinstanceDhcpv6StaticallocationHostOptionsDnsserver struct
type DeviceSystemDhcpserverNetworkinstanceDhcpv6StaticallocationHostOptionsDnsserver struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Dnsserver *string `json:"dns-server,omitempty"`
}

// DeviceSystemDhcpserverNetworkinstanceDhcpv6Traceoptions struct
type DeviceSystemDhcpserverNetworkinstanceDhcpv6Traceoptions struct {
	//RootSystemDhcpserverNetworkinstanceDhcpv6TraceoptionsTrace
	Trace *DeviceSystemDhcpserverNetworkinstanceDhcpv6TraceoptionsTrace `json:"trace,omitempty"`
}

// DeviceSystemDhcpserverNetworkinstanceDhcpv6TraceoptionsTrace struct
type DeviceSystemDhcpserverNetworkinstanceDhcpv6TraceoptionsTrace struct {
	// +kubebuilder:validation:Enum=`messages`
	Trace E_DeviceSystemDhcpserverNetworkinstanceDhcpv6TraceoptionsTraceTrace `json:"trace,omitempty"`
	//Trace *string `json:"trace,omitempty"`
}

// DeviceSystemDns struct
type DeviceSystemDns struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Hostentry       []*DeviceSystemDnsHostentry `json:"host-entry,omitempty"`
	Networkinstance *string                     `json:"network-instance"`
	//RootSystemDnsSearchlist
	Searchlist *DeviceSystemDnsSearchlist `json:"search-list,omitempty"`
	//RootSystemDnsServerlist
	Serverlist *DeviceSystemDnsServerlist `json:"server-list,omitempty"`
}

// DeviceSystemDnsHostentry struct
type DeviceSystemDnsHostentry struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Ipv4address *string `json:"ipv4-address,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Ipv6address *string `json:"ipv6-address,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=253
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((([a-zA-Z0-9_]([a-zA-Z0-9\-_]){0,61})?[a-zA-Z0-9]\.)*([a-zA-Z0-9_]([a-zA-Z0-9\-_]){0,61})?[a-zA-Z0-9]\.?)|\.`
	Name *string `json:"name"`
}

// DeviceSystemDnsSearchlist struct
type DeviceSystemDnsSearchlist struct {
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=253
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((([a-zA-Z0-9_]([a-zA-Z0-9\-_]){0,61})?[a-zA-Z0-9]\.)*([a-zA-Z0-9_]([a-zA-Z0-9\-_]){0,61})?[a-zA-Z0-9]\.?)|\.`
	Searchlist *string `json:"search-list,omitempty"`
}

// DeviceSystemDnsServerlist struct
type DeviceSystemDnsServerlist struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Serverlist *string `json:"server-list,omitempty"`
}

// DeviceSystemFtpserver struct
type DeviceSystemFtpserver struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Networkinstance []*DeviceSystemFtpserverNetworkinstance `json:"network-instance,omitempty"`
}

// DeviceSystemFtpserverNetworkinstance struct
type DeviceSystemFtpserverNetworkinstance struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_DeviceSystemFtpserverNetworkinstanceAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	Name *string `json:"name"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=20
	Sessionlimit *uint8 `json:"session-limit,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	// +kubebuilder:default:="::"
	Sourceaddress *string `json:"source-address,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=300
	Timeout *uint16 `json:"timeout,omitempty"`
}

// DeviceSystemGnmiserver struct
type DeviceSystemGnmiserver struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_DeviceSystemGnmiserverAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=86400
	// +kubebuilder:default:=0
	Commitconfirmedtimeout *uint32 `json:"commit-confirmed-timeout,omitempty"`
	// +kubebuilder:default:=false
	Commitsave *bool `json:"commit-save,omitempty"`
	// +kubebuilder:default:=false
	Includedefaultsinconfigonlyresponses *bool `json:"include-defaults-in-config-only-responses,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Networkinstance []*DeviceSystemGnmiserverNetworkinstance `json:"network-instance,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=60
	Ratelimit *uint16 `json:"rate-limit,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=20
	Sessionlimit *uint16 `json:"session-limit,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=7200
	Timeout *uint16 `json:"timeout,omitempty"`
	//RootSystemGnmiserverTraceoptions
	Traceoptions *DeviceSystemGnmiserverTraceoptions `json:"trace-options,omitempty"`
	//RootSystemGnmiserverUnixsocket
	Unixsocket *DeviceSystemGnmiserverUnixsocket `json:"unix-socket,omitempty"`
}

// DeviceSystemGnmiserverNetworkinstance struct
type DeviceSystemGnmiserverNetworkinstance struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_DeviceSystemGnmiserverNetworkinstanceAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	Name *string `json:"name"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=57400
	Port *uint16 `json:"port,omitempty"`
	//RootSystemGnmiserverNetworkinstanceSourceaddress
	Sourceaddress *DeviceSystemGnmiserverNetworkinstanceSourceaddress `json:"source-address,omitempty"`
	Tlsprofile    *string                                             `json:"tls-profile"`
	// +kubebuilder:default:=true
	Useauthentication *bool `json:"use-authentication,omitempty"`
}

// DeviceSystemGnmiserverNetworkinstanceSourceaddress struct
type DeviceSystemGnmiserverNetworkinstanceSourceaddress struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	// +kubebuilder:default:="::"
	Sourceaddress *string `json:"source-address,omitempty"`
}

// DeviceSystemGnmiserverTraceoptions struct
type DeviceSystemGnmiserverTraceoptions struct {
	// +kubebuilder:validation:Enum=`common`;`request`;`response`
	Traceoptions E_DeviceSystemGnmiserverTraceoptionsTraceoptions `json:"trace-options,omitempty"`
	//Traceoptions *string `json:"trace-options,omitempty"`
}

// DeviceSystemGnmiserverUnixsocket struct
type DeviceSystemGnmiserverUnixsocket struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_DeviceSystemGnmiserverUnixsocketAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	Tlsprofile *string `json:"tls-profile,omitempty"`
	// +kubebuilder:default:=true
	Useauthentication *bool `json:"use-authentication,omitempty"`
}

// DeviceSystemGribiserver struct
type DeviceSystemGribiserver struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_DeviceSystemGribiserverAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Networkinstance []*DeviceSystemGribiserverNetworkinstance `json:"network-instance,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=60
	Ratelimit *uint16 `json:"rate-limit,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=20
	Sessionlimit *uint16 `json:"session-limit,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=7200
	Timeout *uint16 `json:"timeout,omitempty"`
	//RootSystemGribiserverTraceoptions
	Traceoptions *DeviceSystemGribiserverTraceoptions `json:"trace-options,omitempty"`
	//RootSystemGribiserverUnixsocket
	Unixsocket *DeviceSystemGribiserverUnixsocket `json:"unix-socket,omitempty"`
}

// DeviceSystemGribiserverNetworkinstance struct
type DeviceSystemGribiserverNetworkinstance struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_DeviceSystemGribiserverNetworkinstanceAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	Name *string `json:"name"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=57401
	Port *uint16 `json:"port,omitempty"`
	//RootSystemGribiserverNetworkinstanceSourceaddress
	Sourceaddress *DeviceSystemGribiserverNetworkinstanceSourceaddress `json:"source-address,omitempty"`
	Tlsprofile    *string                                              `json:"tls-profile"`
	// +kubebuilder:default:=true
	Useauthentication *bool `json:"use-authentication,omitempty"`
}

// DeviceSystemGribiserverNetworkinstanceSourceaddress struct
type DeviceSystemGribiserverNetworkinstanceSourceaddress struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Sourceaddress *string `json:"source-address,omitempty"`
}

// DeviceSystemGribiserverTraceoptions struct
type DeviceSystemGribiserverTraceoptions struct {
	// +kubebuilder:validation:Enum=`common`;`request`;`response`
	Traceoptions E_DeviceSystemGribiserverTraceoptionsTraceoptions `json:"trace-options,omitempty"`
	//Traceoptions *string `json:"trace-options,omitempty"`
}

// DeviceSystemGribiserverUnixsocket struct
type DeviceSystemGribiserverUnixsocket struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_DeviceSystemGribiserverUnixsocketAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	Tlsprofile *string `json:"tls-profile,omitempty"`
	// +kubebuilder:default:=true
	Useauthentication *bool `json:"use-authentication,omitempty"`
}

// DeviceSystemInformation struct
type DeviceSystemInformation struct {
	Contact  *string `json:"contact,omitempty"`
	Location *string `json:"location,omitempty"`
}

// DeviceSystemJsonrpcserver struct
type DeviceSystemJsonrpcserver struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_DeviceSystemJsonrpcserverAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=86400
	// +kubebuilder:default:=0
	Commitconfirmedtimeout *uint32 `json:"commit-confirmed-timeout,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Networkinstance []*DeviceSystemJsonrpcserverNetworkinstance `json:"network-instance,omitempty"`
	//RootSystemJsonrpcserverTraceoptions
	Traceoptions *DeviceSystemJsonrpcserverTraceoptions `json:"trace-options,omitempty"`
	//RootSystemJsonrpcserverUnixsocket
	Unixsocket *DeviceSystemJsonrpcserverUnixsocket `json:"unix-socket,omitempty"`
}

// DeviceSystemJsonrpcserverNetworkinstance struct
type DeviceSystemJsonrpcserverNetworkinstance struct {
	//RootSystemJsonrpcserverNetworkinstanceHttp
	Http *DeviceSystemJsonrpcserverNetworkinstanceHttp `json:"http,omitempty"`
	//RootSystemJsonrpcserverNetworkinstanceHttps
	Https *DeviceSystemJsonrpcserverNetworkinstanceHttps `json:"https,omitempty"`
	Name  *string                                        `json:"name"`
}

// DeviceSystemJsonrpcserverNetworkinstanceHttp struct
type DeviceSystemJsonrpcserverNetworkinstanceHttp struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_DeviceSystemJsonrpcserverNetworkinstanceHttpAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=80
	Port *uint16 `json:"port,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=10
	Sessionlimit *uint16 `json:"session-limit,omitempty"`
	//RootSystemJsonrpcserverNetworkinstanceHttpSourceaddress
	Sourceaddress *DeviceSystemJsonrpcserverNetworkinstanceHttpSourceaddress `json:"source-address,omitempty"`
	// +kubebuilder:default:=true
	Useauthentication *bool `json:"use-authentication,omitempty"`
}

// DeviceSystemJsonrpcserverNetworkinstanceHttpSourceaddress struct
type DeviceSystemJsonrpcserverNetworkinstanceHttpSourceaddress struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	// +kubebuilder:default:="::"
	Sourceaddress *string `json:"source-address,omitempty"`
}

// DeviceSystemJsonrpcserverNetworkinstanceHttps struct
type DeviceSystemJsonrpcserverNetworkinstanceHttps struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_DeviceSystemJsonrpcserverNetworkinstanceHttpsAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=443
	Port *uint16 `json:"port,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=10
	Sessionlimit *uint16 `json:"session-limit,omitempty"`
	//RootSystemJsonrpcserverNetworkinstanceHttpsSourceaddress
	Sourceaddress *DeviceSystemJsonrpcserverNetworkinstanceHttpsSourceaddress `json:"source-address,omitempty"`
	Tlsprofile    *string                                                     `json:"tls-profile"`
	// +kubebuilder:default:=true
	Useauthentication *bool `json:"use-authentication,omitempty"`
}

// DeviceSystemJsonrpcserverNetworkinstanceHttpsSourceaddress struct
type DeviceSystemJsonrpcserverNetworkinstanceHttpsSourceaddress struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	// +kubebuilder:default:="::"
	Sourceaddress *string `json:"source-address,omitempty"`
}

// DeviceSystemJsonrpcserverTraceoptions struct
type DeviceSystemJsonrpcserverTraceoptions struct {
	// +kubebuilder:validation:Enum=`common`;`request`;`response`
	Traceoptions E_DeviceSystemJsonrpcserverTraceoptionsTraceoptions `json:"trace-options,omitempty"`
	//Traceoptions *string `json:"trace-options,omitempty"`
}

// DeviceSystemJsonrpcserverUnixsocket struct
type DeviceSystemJsonrpcserverUnixsocket struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_DeviceSystemJsonrpcserverUnixsocketAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	Tlsprofile *string `json:"tls-profile,omitempty"`
	// +kubebuilder:default:=true
	Useauthentication *bool `json:"use-authentication,omitempty"`
}

// DeviceSystemLacp struct
type DeviceSystemLacp struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}`
	Systemid *string `json:"system-id,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	Systempriority *uint16 `json:"system-priority,omitempty"`
}

// DeviceSystemLldp struct
type DeviceSystemLldp struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_DeviceSystemLldpAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	//RootSystemLldpBgpautodiscovery
	Bgpautodiscovery *DeviceSystemLldpBgpautodiscovery `json:"bgp-auto-discovery,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=-1
	// +kubebuilder:default:=30
	Hellotimer *uint64 `json:"hello-timer,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=4
	Holdmultiplier *uint8 `json:"hold-multiplier,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Interface []*DeviceSystemLldpInterface `json:"interface,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Managementaddress []*DeviceSystemLldpManagementaddress `json:"management-address,omitempty"`
	//RootSystemLldpTraceoptions
	Traceoptions *DeviceSystemLldpTraceoptions `json:"trace-options,omitempty"`
}

// DeviceSystemLldpBgpautodiscovery struct
type DeviceSystemLldpBgpautodiscovery struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_DeviceSystemLldpBgpautodiscoveryAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	Groupid *uint32 `json:"group-id,omitempty"`
	//RootSystemLldpBgpautodiscoveryNetworkinstance
	Networkinstance *DeviceSystemLldpBgpautodiscoveryNetworkinstance `json:"network-instance,omitempty"`
}

// DeviceSystemLldpBgpautodiscoveryNetworkinstance struct
type DeviceSystemLldpBgpautodiscoveryNetworkinstance struct {
	Networkinstance *string `json:"network-instance,omitempty"`
}

// DeviceSystemLldpInterface struct
type DeviceSystemLldpInterface struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_DeviceSystemLldpInterfaceAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	//RootSystemLldpInterfaceBgpautodiscovery
	Bgpautodiscovery *DeviceSystemLldpInterfaceBgpautodiscovery `json:"bgp-auto-discovery,omitempty"`
	Name             *string                                    `json:"name"`
}

// DeviceSystemLldpInterfaceBgpautodiscovery struct
type DeviceSystemLldpInterfaceBgpautodiscovery struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_DeviceSystemLldpInterfaceBgpautodiscoveryAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	Groupid *uint32 `json:"group-id,omitempty"`
	//RootSystemLldpInterfaceBgpautodiscoveryPeeringaddress
	Peeringaddress *DeviceSystemLldpInterfaceBgpautodiscoveryPeeringaddress `json:"peering-address,omitempty"`
}

// DeviceSystemLldpInterfaceBgpautodiscoveryPeeringaddress struct
type DeviceSystemLldpInterfaceBgpautodiscoveryPeeringaddress struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Peeringaddress *string `json:"peering-address,omitempty"`
}

// DeviceSystemLldpManagementaddress struct
type DeviceSystemLldpManagementaddress struct {
	// kubebuilder:validation:MinLength=5
	// kubebuilder:validation:MaxLength=25
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(mgmt0\.0|system0\.0|lo(0|1[0-9][0-9]|2([0-4][0-9]|5[0-5])|[1-9][0-9]|[1-9])\.(0|[1-9](\d){0,3})|ethernet-([1-9](\d){0,1}(/[abcd])?(/[1-9](\d){0,1})?/(([1-9](\d){0,1})|(1[0-1]\d)|(12[0-8])))\.([0]|[1-9](\d){0,3})|irb(0|1[0-9][0-9]|2([0-4][0-9]|5[0-5])|[1-9][0-9]|[1-9])\.(0|[1-9](\d){0,3})|lag(([1-9](\d){0,1})|(1[0-1]\d)|(12[0-8]))\.(0|[1-9](\d){0,3}))`
	Subinterface *string `json:"subinterface"`
	//RootSystemLldpManagementaddressType
	Type *DeviceSystemLldpManagementaddressType `json:"type,omitempty"`
}

// DeviceSystemLldpManagementaddressType struct
type DeviceSystemLldpManagementaddressType struct {
	// +kubebuilder:validation:Enum=`IPv4`;`IPv6`
	Type E_DeviceSystemLldpManagementaddressTypeType `json:"type,omitempty"`
	//Type *string `json:"type,omitempty"`
}

// DeviceSystemLldpTraceoptions struct
type DeviceSystemLldpTraceoptions struct {
	// +kubebuilder:validation:Enum=`common`;`received`;`transmitted`
	Traceoptions E_DeviceSystemLldpTraceoptionsTraceoptions `json:"trace-options,omitempty"`
	//Traceoptions *string `json:"trace-options,omitempty"`
}

// DeviceSystemLoadbalancing struct
type DeviceSystemLoadbalancing struct {
	//RootSystemLoadbalancingHashoptions
	Hashoptions *DeviceSystemLoadbalancingHashoptions `json:"hash-options,omitempty"`
}

// DeviceSystemLoadbalancingHashoptions struct
type DeviceSystemLoadbalancingHashoptions struct {
	// +kubebuilder:default:=true
	Destinationaddress *bool `json:"destination-address,omitempty"`
	// +kubebuilder:default:=true
	Destinationport *bool `json:"destination-port,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=0
	Hashseed *uint16 `json:"hash-seed,omitempty"`
	// +kubebuilder:default:=false
	Ipv6flowlabel *bool `json:"ipv6-flow-label,omitempty"`
	// +kubebuilder:default:=false
	Mplslabelstack *bool `json:"mpls-label-stack,omitempty"`
	// +kubebuilder:default:=true
	Protocol *bool `json:"protocol,omitempty"`
	// +kubebuilder:default:=true
	Sourceaddress *bool `json:"source-address,omitempty"`
	// +kubebuilder:default:=true
	Sourceport *bool `json:"source-port,omitempty"`
	// +kubebuilder:default:=true
	Vlan *bool `json:"vlan,omitempty"`
}

// DeviceSystemLogging struct
type DeviceSystemLogging struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Buffer []*DeviceSystemLoggingBuffer `json:"buffer,omitempty"`
	//RootSystemLoggingConsole
	Console *DeviceSystemLoggingConsole `json:"console,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	File []*DeviceSystemLoggingFile `json:"file,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Filter          []*DeviceSystemLoggingFilter `json:"filter,omitempty"`
	Networkinstance *string                      `json:"network-instance,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Remoteserver []*DeviceSystemLoggingRemoteserver `json:"remote-server,omitempty"`
	// +kubebuilder:validation:Enum=`auth`;`authpriv`;`cron`;`daemon`;`ftp`;`kern`;`local0`;`local1`;`local2`;`local3`;`local4`;`local5`;`local6`;`local7`;`lpr`;`mail`;`news`;`syslog`;`user`;`uucp`
	// +kubebuilder:default:="local6"
	Subsystemfacility E_DeviceSystemLoggingSubsystemfacility `json:"subsystem-facility,omitempty"`
	//Subsystemfacility *string `json:"subsystem-facility,omitempty"`
}

// DeviceSystemLoggingBuffer struct
type DeviceSystemLoggingBuffer struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`([^./][^/]*)|(\.[^\./]+)|(\.\.[^/])+`
	Buffername *string `json:"buffer-name"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Facility []*DeviceSystemLoggingBufferFacility `json:"facility,omitempty"`
	//RootSystemLoggingBufferFilter
	Filter *DeviceSystemLoggingBufferFilter `json:"filter,omitempty"`
	// +kubebuilder:default:="%TIMEGENERATED:::date-rfc3339% %HOSTNAME% %SYSLOGTAG%%MSG:::sp-if-no-1st-sp%%MSG:::drop-last-lf%\n"
	Format *string `json:"format,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=0
	// +kubebuilder:default:=0
	Persist *uint32 `json:"persist,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=4
	Rotate *uint16 `json:"rotate,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[1-9][0-9]{0,15}(K|M|G){0,1}`
	// +kubebuilder:default:="10M"
	Size *string `json:"size,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Subsystem []*DeviceSystemLoggingBufferSubsystem `json:"subsystem,omitempty"`
}

// DeviceSystemLoggingBufferFacility struct
type DeviceSystemLoggingBufferFacility struct {
	// +kubebuilder:validation:Enum=`auth`;`authpriv`;`cron`;`daemon`;`ftp`;`kern`;`local0`;`local1`;`local2`;`local3`;`local4`;`local5`;`local6`;`local7`;`lpr`;`mail`;`news`;`syslog`;`user`;`uucp`
	Facilityname E_DeviceSystemLoggingBufferFacilityFacilityname `json:"facility-name,omitempty"`
	//RootSystemLoggingBufferFacilityPriority
	Priority *DeviceSystemLoggingBufferFacilityPriority `json:"priority,omitempty"`
}

// DeviceSystemLoggingBufferFacilityPriority struct
type DeviceSystemLoggingBufferFacilityPriority struct {
	// +kubebuilder:validation:Enum=`alert`;`critical`;`debug`;`emergency`;`error`;`informational`;`notice`;`warning`
	Matchabove E_DeviceSystemLoggingBufferFacilityPriorityMatchabove `json:"match-above,omitempty"`
	//Matchabove *string `json:"match-above,omitempty"`
	//RootSystemLoggingBufferFacilityPriorityMatchexact
	Matchexact *DeviceSystemLoggingBufferFacilityPriorityMatchexact `json:"match-exact,omitempty"`
}

// DeviceSystemLoggingBufferFacilityPriorityMatchexact struct
type DeviceSystemLoggingBufferFacilityPriorityMatchexact struct {
	// +kubebuilder:validation:Enum=`alert`;`critical`;`debug`;`emergency`;`error`;`informational`;`notice`;`warning`
	Matchexact E_DeviceSystemLoggingBufferFacilityPriorityMatchexactMatchexact `json:"match-exact,omitempty"`
	//Matchexact *string `json:"match-exact,omitempty"`
}

// DeviceSystemLoggingBufferFilter struct
type DeviceSystemLoggingBufferFilter struct {
	Filter *string `json:"filter,omitempty"`
}

// DeviceSystemLoggingBufferSubsystem struct
type DeviceSystemLoggingBufferSubsystem struct {
	//RootSystemLoggingBufferSubsystemPriority
	Priority *DeviceSystemLoggingBufferSubsystemPriority `json:"priority,omitempty"`
	// +kubebuilder:validation:Enum=`aaa`;`accounting`;`acl`;`app`;`arpnd`;`bfd`;`bgp`;`bridgetable`;`chassis`;`debug`;`dhcp`;`ethcfm`;`evpn`;`fib`;`gnmi`;`gribi`;`igmp`;`isis`;`json`;`lag`;`ldp`;`linux`;`lldp`;`log`;`mgmt`;`mld`;`mpls`;`netinst`;`ospf`;`p4rt`;`pim`;`platform`;`policy`;`qos`;`radio`;`sdk`;`sflow`;`staticroute`;`twamp`;`vxlan`;`xdp`
	Subsystemname E_DeviceSystemLoggingBufferSubsystemSubsystemname `json:"subsystem-name,omitempty"`
}

// DeviceSystemLoggingBufferSubsystemPriority struct
type DeviceSystemLoggingBufferSubsystemPriority struct {
	// +kubebuilder:validation:Enum=`alert`;`critical`;`debug`;`emergency`;`error`;`informational`;`notice`;`warning`
	Matchabove E_DeviceSystemLoggingBufferSubsystemPriorityMatchabove `json:"match-above,omitempty"`
	//Matchabove *string `json:"match-above,omitempty"`
	//RootSystemLoggingBufferSubsystemPriorityMatchexact
	Matchexact *DeviceSystemLoggingBufferSubsystemPriorityMatchexact `json:"match-exact,omitempty"`
}

// DeviceSystemLoggingBufferSubsystemPriorityMatchexact struct
type DeviceSystemLoggingBufferSubsystemPriorityMatchexact struct {
	// +kubebuilder:validation:Enum=`alert`;`critical`;`debug`;`emergency`;`error`;`informational`;`notice`;`warning`
	Matchexact E_DeviceSystemLoggingBufferSubsystemPriorityMatchexactMatchexact `json:"match-exact,omitempty"`
	//Matchexact *string `json:"match-exact,omitempty"`
}

// DeviceSystemLoggingConsole struct
type DeviceSystemLoggingConsole struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Facility []*DeviceSystemLoggingConsoleFacility `json:"facility,omitempty"`
	//RootSystemLoggingConsoleFilter
	Filter *DeviceSystemLoggingConsoleFilter `json:"filter,omitempty"`
	// +kubebuilder:default:="%TIMEGENERATED:::date-rfc3339% %HOSTNAME% %SYSLOGTAG%%MSG:::sp-if-no-1st-sp%%MSG:::drop-last-lf%\n"
	Format *string `json:"format,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Subsystem []*DeviceSystemLoggingConsoleSubsystem `json:"subsystem,omitempty"`
}

// DeviceSystemLoggingConsoleFacility struct
type DeviceSystemLoggingConsoleFacility struct {
	// +kubebuilder:validation:Enum=`auth`;`authpriv`;`cron`;`daemon`;`ftp`;`kern`;`local0`;`local1`;`local2`;`local3`;`local4`;`local5`;`local6`;`local7`;`lpr`;`mail`;`news`;`syslog`;`user`;`uucp`
	Facilityname E_DeviceSystemLoggingConsoleFacilityFacilityname `json:"facility-name,omitempty"`
	//RootSystemLoggingConsoleFacilityPriority
	Priority *DeviceSystemLoggingConsoleFacilityPriority `json:"priority,omitempty"`
}

// DeviceSystemLoggingConsoleFacilityPriority struct
type DeviceSystemLoggingConsoleFacilityPriority struct {
	// +kubebuilder:validation:Enum=`alert`;`critical`;`debug`;`emergency`;`error`;`informational`;`notice`;`warning`
	Matchabove E_DeviceSystemLoggingConsoleFacilityPriorityMatchabove `json:"match-above,omitempty"`
	//Matchabove *string `json:"match-above,omitempty"`
	//RootSystemLoggingConsoleFacilityPriorityMatchexact
	Matchexact *DeviceSystemLoggingConsoleFacilityPriorityMatchexact `json:"match-exact,omitempty"`
}

// DeviceSystemLoggingConsoleFacilityPriorityMatchexact struct
type DeviceSystemLoggingConsoleFacilityPriorityMatchexact struct {
	// +kubebuilder:validation:Enum=`alert`;`critical`;`debug`;`emergency`;`error`;`informational`;`notice`;`warning`
	Matchexact E_DeviceSystemLoggingConsoleFacilityPriorityMatchexactMatchexact `json:"match-exact,omitempty"`
	//Matchexact *string `json:"match-exact,omitempty"`
}

// DeviceSystemLoggingConsoleFilter struct
type DeviceSystemLoggingConsoleFilter struct {
	Filter *string `json:"filter,omitempty"`
}

// DeviceSystemLoggingConsoleSubsystem struct
type DeviceSystemLoggingConsoleSubsystem struct {
	//RootSystemLoggingConsoleSubsystemPriority
	Priority *DeviceSystemLoggingConsoleSubsystemPriority `json:"priority,omitempty"`
	// +kubebuilder:validation:Enum=`aaa`;`accounting`;`acl`;`app`;`arpnd`;`bfd`;`bgp`;`bridgetable`;`chassis`;`debug`;`dhcp`;`ethcfm`;`evpn`;`fib`;`gnmi`;`gribi`;`igmp`;`isis`;`json`;`lag`;`ldp`;`linux`;`lldp`;`log`;`mgmt`;`mld`;`mpls`;`netinst`;`ospf`;`p4rt`;`pim`;`platform`;`policy`;`qos`;`radio`;`sdk`;`sflow`;`staticroute`;`twamp`;`vxlan`;`xdp`
	Subsystemname E_DeviceSystemLoggingConsoleSubsystemSubsystemname `json:"subsystem-name,omitempty"`
}

// DeviceSystemLoggingConsoleSubsystemPriority struct
type DeviceSystemLoggingConsoleSubsystemPriority struct {
	// +kubebuilder:validation:Enum=`alert`;`critical`;`debug`;`emergency`;`error`;`informational`;`notice`;`warning`
	Matchabove E_DeviceSystemLoggingConsoleSubsystemPriorityMatchabove `json:"match-above,omitempty"`
	//Matchabove *string `json:"match-above,omitempty"`
	//RootSystemLoggingConsoleSubsystemPriorityMatchexact
	Matchexact *DeviceSystemLoggingConsoleSubsystemPriorityMatchexact `json:"match-exact,omitempty"`
}

// DeviceSystemLoggingConsoleSubsystemPriorityMatchexact struct
type DeviceSystemLoggingConsoleSubsystemPriorityMatchexact struct {
	// +kubebuilder:validation:Enum=`alert`;`critical`;`debug`;`emergency`;`error`;`informational`;`notice`;`warning`
	Matchexact E_DeviceSystemLoggingConsoleSubsystemPriorityMatchexactMatchexact `json:"match-exact,omitempty"`
	//Matchexact *string `json:"match-exact,omitempty"`
}

// DeviceSystemLoggingFile struct
type DeviceSystemLoggingFile struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`/(.)*`
	// +kubebuilder:default:="/var/log/srlinux/file"
	Directory *string `json:"directory,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Facility []*DeviceSystemLoggingFileFacility `json:"facility,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`([^./][^/]*)|(\.[^\./]+)|(\.\.[^/])+`
	Filename *string `json:"file-name"`
	//RootSystemLoggingFileFilter
	Filter *DeviceSystemLoggingFileFilter `json:"filter,omitempty"`
	// +kubebuilder:default:="%TIMEGENERATED:::date-rfc3339% %HOSTNAME% %SYSLOGTAG%%MSG:::sp-if-no-1st-sp%%MSG:::drop-last-lf%\n"
	Format *string `json:"format,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=4
	Rotate *uint16 `json:"rotate,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[1-9][0-9]{0,15}(K|M|G){0,1}`
	// +kubebuilder:default:="10M"
	Size *string `json:"size,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Subsystem []*DeviceSystemLoggingFileSubsystem `json:"subsystem,omitempty"`
}

// DeviceSystemLoggingFileFacility struct
type DeviceSystemLoggingFileFacility struct {
	// +kubebuilder:validation:Enum=`auth`;`authpriv`;`cron`;`daemon`;`ftp`;`kern`;`local0`;`local1`;`local2`;`local3`;`local4`;`local5`;`local6`;`local7`;`lpr`;`mail`;`news`;`syslog`;`user`;`uucp`
	Facilityname E_DeviceSystemLoggingFileFacilityFacilityname `json:"facility-name,omitempty"`
	//RootSystemLoggingFileFacilityPriority
	Priority *DeviceSystemLoggingFileFacilityPriority `json:"priority,omitempty"`
}

// DeviceSystemLoggingFileFacilityPriority struct
type DeviceSystemLoggingFileFacilityPriority struct {
	// +kubebuilder:validation:Enum=`alert`;`critical`;`debug`;`emergency`;`error`;`informational`;`notice`;`warning`
	Matchabove E_DeviceSystemLoggingFileFacilityPriorityMatchabove `json:"match-above,omitempty"`
	//Matchabove *string `json:"match-above,omitempty"`
	//RootSystemLoggingFileFacilityPriorityMatchexact
	Matchexact *DeviceSystemLoggingFileFacilityPriorityMatchexact `json:"match-exact,omitempty"`
}

// DeviceSystemLoggingFileFacilityPriorityMatchexact struct
type DeviceSystemLoggingFileFacilityPriorityMatchexact struct {
	// +kubebuilder:validation:Enum=`alert`;`critical`;`debug`;`emergency`;`error`;`informational`;`notice`;`warning`
	Matchexact E_DeviceSystemLoggingFileFacilityPriorityMatchexactMatchexact `json:"match-exact,omitempty"`
	//Matchexact *string `json:"match-exact,omitempty"`
}

// DeviceSystemLoggingFileFilter struct
type DeviceSystemLoggingFileFilter struct {
	Filter *string `json:"filter,omitempty"`
}

// DeviceSystemLoggingFileSubsystem struct
type DeviceSystemLoggingFileSubsystem struct {
	//RootSystemLoggingFileSubsystemPriority
	Priority *DeviceSystemLoggingFileSubsystemPriority `json:"priority,omitempty"`
	// +kubebuilder:validation:Enum=`aaa`;`accounting`;`acl`;`app`;`arpnd`;`bfd`;`bgp`;`bridgetable`;`chassis`;`debug`;`dhcp`;`ethcfm`;`evpn`;`fib`;`gnmi`;`gribi`;`igmp`;`isis`;`json`;`lag`;`ldp`;`linux`;`lldp`;`log`;`mgmt`;`mld`;`mpls`;`netinst`;`ospf`;`p4rt`;`pim`;`platform`;`policy`;`qos`;`radio`;`sdk`;`sflow`;`staticroute`;`twamp`;`vxlan`;`xdp`
	Subsystemname E_DeviceSystemLoggingFileSubsystemSubsystemname `json:"subsystem-name,omitempty"`
}

// DeviceSystemLoggingFileSubsystemPriority struct
type DeviceSystemLoggingFileSubsystemPriority struct {
	// +kubebuilder:validation:Enum=`alert`;`critical`;`debug`;`emergency`;`error`;`informational`;`notice`;`warning`
	Matchabove E_DeviceSystemLoggingFileSubsystemPriorityMatchabove `json:"match-above,omitempty"`
	//Matchabove *string `json:"match-above,omitempty"`
	//RootSystemLoggingFileSubsystemPriorityMatchexact
	Matchexact *DeviceSystemLoggingFileSubsystemPriorityMatchexact `json:"match-exact,omitempty"`
}

// DeviceSystemLoggingFileSubsystemPriorityMatchexact struct
type DeviceSystemLoggingFileSubsystemPriorityMatchexact struct {
	// +kubebuilder:validation:Enum=`alert`;`critical`;`debug`;`emergency`;`error`;`informational`;`notice`;`warning`
	Matchexact E_DeviceSystemLoggingFileSubsystemPriorityMatchexactMatchexact `json:"match-exact,omitempty"`
	//Matchexact *string `json:"match-exact,omitempty"`
}

// DeviceSystemLoggingFilter struct
type DeviceSystemLoggingFilter struct {
	Contains *string `json:"contains,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Facility []*DeviceSystemLoggingFilterFacility `json:"facility,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`([0-9a-zA-Z\-_.@!^*()\[\]{}|\\/<>,;])+`
	Filtername *string `json:"filter-name"`
	Prefix     *string `json:"prefix,omitempty"`
	Regex      *string `json:"regex,omitempty"`
	Tag        *string `json:"tag,omitempty"`
}

// DeviceSystemLoggingFilterFacility struct
type DeviceSystemLoggingFilterFacility struct {
	// +kubebuilder:validation:Enum=`auth`;`authpriv`;`cron`;`daemon`;`ftp`;`kern`;`local0`;`local1`;`local2`;`local3`;`local4`;`local5`;`local6`;`local7`;`lpr`;`mail`;`news`;`syslog`;`user`;`uucp`
	Facilityname E_DeviceSystemLoggingFilterFacilityFacilityname `json:"facility-name,omitempty"`
	//RootSystemLoggingFilterFacilityPriority
	Priority *DeviceSystemLoggingFilterFacilityPriority `json:"priority,omitempty"`
}

// DeviceSystemLoggingFilterFacilityPriority struct
type DeviceSystemLoggingFilterFacilityPriority struct {
	// +kubebuilder:validation:Enum=`alert`;`critical`;`debug`;`emergency`;`error`;`informational`;`notice`;`warning`
	Matchabove E_DeviceSystemLoggingFilterFacilityPriorityMatchabove `json:"match-above,omitempty"`
	//Matchabove *string `json:"match-above,omitempty"`
	//RootSystemLoggingFilterFacilityPriorityMatchexact
	Matchexact *DeviceSystemLoggingFilterFacilityPriorityMatchexact `json:"match-exact,omitempty"`
}

// DeviceSystemLoggingFilterFacilityPriorityMatchexact struct
type DeviceSystemLoggingFilterFacilityPriorityMatchexact struct {
	// +kubebuilder:validation:Enum=`alert`;`critical`;`debug`;`emergency`;`error`;`informational`;`notice`;`warning`
	Matchexact E_DeviceSystemLoggingFilterFacilityPriorityMatchexactMatchexact `json:"match-exact,omitempty"`
	//Matchexact *string `json:"match-exact,omitempty"`
}

// DeviceSystemLoggingRemoteserver struct
type DeviceSystemLoggingRemoteserver struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Facility []*DeviceSystemLoggingRemoteserverFacility `json:"facility,omitempty"`
	//RootSystemLoggingRemoteserverFilter
	Filter *DeviceSystemLoggingRemoteserverFilter `json:"filter,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`((([a-zA-Z0-9_]([a-zA-Z0-9\-_]){0,61})?[a-zA-Z0-9]\.)*([a-zA-Z0-9_]([a-zA-Z0-9\-_]){0,61})?[a-zA-Z0-9]\.?)|\.`
	Host *string `json:"host"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	// +kubebuilder:default:=514
	Remoteport *uint32 `json:"remote-port,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Subsystem []*DeviceSystemLoggingRemoteserverSubsystem `json:"subsystem,omitempty"`
	// +kubebuilder:validation:Enum=`tcp`;`udp`
	// +kubebuilder:default:="udp"
	Transport E_DeviceSystemLoggingRemoteserverTransport `json:"transport,omitempty"`
	//Transport *string `json:"transport,omitempty"`
}

// DeviceSystemLoggingRemoteserverFacility struct
type DeviceSystemLoggingRemoteserverFacility struct {
	// +kubebuilder:validation:Enum=`auth`;`authpriv`;`cron`;`daemon`;`ftp`;`kern`;`local0`;`local1`;`local2`;`local3`;`local4`;`local5`;`local6`;`local7`;`lpr`;`mail`;`news`;`syslog`;`user`;`uucp`
	Facilityname E_DeviceSystemLoggingRemoteserverFacilityFacilityname `json:"facility-name,omitempty"`
	//RootSystemLoggingRemoteserverFacilityPriority
	Priority *DeviceSystemLoggingRemoteserverFacilityPriority `json:"priority,omitempty"`
}

// DeviceSystemLoggingRemoteserverFacilityPriority struct
type DeviceSystemLoggingRemoteserverFacilityPriority struct {
	// +kubebuilder:validation:Enum=`alert`;`critical`;`debug`;`emergency`;`error`;`informational`;`notice`;`warning`
	Matchabove E_DeviceSystemLoggingRemoteserverFacilityPriorityMatchabove `json:"match-above,omitempty"`
	//Matchabove *string `json:"match-above,omitempty"`
	//RootSystemLoggingRemoteserverFacilityPriorityMatchexact
	Matchexact *DeviceSystemLoggingRemoteserverFacilityPriorityMatchexact `json:"match-exact,omitempty"`
}

// DeviceSystemLoggingRemoteserverFacilityPriorityMatchexact struct
type DeviceSystemLoggingRemoteserverFacilityPriorityMatchexact struct {
	// +kubebuilder:validation:Enum=`alert`;`critical`;`debug`;`emergency`;`error`;`informational`;`notice`;`warning`
	Matchexact E_DeviceSystemLoggingRemoteserverFacilityPriorityMatchexactMatchexact `json:"match-exact,omitempty"`
	//Matchexact *string `json:"match-exact,omitempty"`
}

// DeviceSystemLoggingRemoteserverFilter struct
type DeviceSystemLoggingRemoteserverFilter struct {
	Filter *string `json:"filter,omitempty"`
}

// DeviceSystemLoggingRemoteserverSubsystem struct
type DeviceSystemLoggingRemoteserverSubsystem struct {
	//RootSystemLoggingRemoteserverSubsystemPriority
	Priority *DeviceSystemLoggingRemoteserverSubsystemPriority `json:"priority,omitempty"`
	// +kubebuilder:validation:Enum=`aaa`;`accounting`;`acl`;`app`;`arpnd`;`bfd`;`bgp`;`bridgetable`;`chassis`;`debug`;`dhcp`;`ethcfm`;`evpn`;`fib`;`gnmi`;`gribi`;`igmp`;`isis`;`json`;`lag`;`ldp`;`linux`;`lldp`;`log`;`mgmt`;`mld`;`mpls`;`netinst`;`ospf`;`p4rt`;`pim`;`platform`;`policy`;`qos`;`radio`;`sdk`;`sflow`;`staticroute`;`twamp`;`vxlan`;`xdp`
	Subsystemname E_DeviceSystemLoggingRemoteserverSubsystemSubsystemname `json:"subsystem-name,omitempty"`
}

// DeviceSystemLoggingRemoteserverSubsystemPriority struct
type DeviceSystemLoggingRemoteserverSubsystemPriority struct {
	// +kubebuilder:validation:Enum=`alert`;`critical`;`debug`;`emergency`;`error`;`informational`;`notice`;`warning`
	Matchabove E_DeviceSystemLoggingRemoteserverSubsystemPriorityMatchabove `json:"match-above,omitempty"`
	//Matchabove *string `json:"match-above,omitempty"`
	//RootSystemLoggingRemoteserverSubsystemPriorityMatchexact
	Matchexact *DeviceSystemLoggingRemoteserverSubsystemPriorityMatchexact `json:"match-exact,omitempty"`
}

// DeviceSystemLoggingRemoteserverSubsystemPriorityMatchexact struct
type DeviceSystemLoggingRemoteserverSubsystemPriorityMatchexact struct {
	// +kubebuilder:validation:Enum=`alert`;`critical`;`debug`;`emergency`;`error`;`informational`;`notice`;`warning`
	Matchexact E_DeviceSystemLoggingRemoteserverSubsystemPriorityMatchexactMatchexact `json:"match-exact,omitempty"`
	//Matchexact *string `json:"match-exact,omitempty"`
}

// DeviceSystemMaintenance struct
type DeviceSystemMaintenance struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Group []*DeviceSystemMaintenanceGroup `json:"group,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Profile []*DeviceSystemMaintenanceProfile `json:"profile,omitempty"`
}

// DeviceSystemMaintenanceGroup struct
type DeviceSystemMaintenanceGroup struct {
	//RootSystemMaintenanceGroupMaintenancemode
	Maintenancemode    *DeviceSystemMaintenanceGroupMaintenancemode `json:"maintenance-mode,omitempty"`
	Maintenanceprofile *string                                      `json:"maintenance-profile,omitempty"`
	//RootSystemMaintenanceGroupMembers
	Members *DeviceSystemMaintenanceGroupMembers `json:"members,omitempty"`
	Name    *string                              `json:"name"`
}

// DeviceSystemMaintenanceGroupMaintenancemode struct
type DeviceSystemMaintenanceGroupMaintenancemode struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_DeviceSystemMaintenanceGroupMaintenancemodeAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
}

// DeviceSystemMaintenanceGroupMembers struct
type DeviceSystemMaintenanceGroupMembers struct {
	//RootSystemMaintenanceGroupMembersBgp
	Bgp *DeviceSystemMaintenanceGroupMembersBgp `json:"bgp,omitempty"`
	//RootSystemMaintenanceGroupMembersIsis
	Isis *DeviceSystemMaintenanceGroupMembersIsis `json:"isis,omitempty"`
}

// DeviceSystemMaintenanceGroupMembersBgp struct
type DeviceSystemMaintenanceGroupMembersBgp struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Networkinstance []*DeviceSystemMaintenanceGroupMembersBgpNetworkinstance `json:"network-instance,omitempty"`
}

// DeviceSystemMaintenanceGroupMembersBgpNetworkinstance struct
type DeviceSystemMaintenanceGroupMembersBgpNetworkinstance struct {
	Name *string `json:"name"`
	//RootSystemMaintenanceGroupMembersBgpNetworkinstanceNeighbor
	Neighbor *DeviceSystemMaintenanceGroupMembersBgpNetworkinstanceNeighbor `json:"neighbor,omitempty"`
	//RootSystemMaintenanceGroupMembersBgpNetworkinstancePeergroup
	Peergroup *DeviceSystemMaintenanceGroupMembersBgpNetworkinstancePeergroup `json:"peer-group,omitempty"`
}

// DeviceSystemMaintenanceGroupMembersBgpNetworkinstanceNeighbor struct
type DeviceSystemMaintenanceGroupMembersBgpNetworkinstanceNeighbor struct {
	Neighbor *string `json:"neighbor,omitempty"`
}

// DeviceSystemMaintenanceGroupMembersBgpNetworkinstancePeergroup struct
type DeviceSystemMaintenanceGroupMembersBgpNetworkinstancePeergroup struct {
	Peergroup *string `json:"peer-group,omitempty"`
}

// DeviceSystemMaintenanceGroupMembersIsis struct
type DeviceSystemMaintenanceGroupMembersIsis struct {
	//RootSystemMaintenanceGroupMembersIsisNetworkinstances
	Networkinstances *DeviceSystemMaintenanceGroupMembersIsisNetworkinstances `json:"network-instances,omitempty"`
}

// DeviceSystemMaintenanceGroupMembersIsisNetworkinstances struct
type DeviceSystemMaintenanceGroupMembersIsisNetworkinstances struct {
	Networkinstances *string `json:"network-instances,omitempty"`
}

// DeviceSystemMaintenanceProfile struct
type DeviceSystemMaintenanceProfile struct {
	//RootSystemMaintenanceProfileBgp
	Bgp *DeviceSystemMaintenanceProfileBgp `json:"bgp,omitempty"`
	//RootSystemMaintenanceProfileIsis
	Isis *DeviceSystemMaintenanceProfileIsis `json:"isis,omitempty"`
	Name *string                             `json:"name"`
}

// DeviceSystemMaintenanceProfileBgp struct
type DeviceSystemMaintenanceProfileBgp struct {
	Exportpolicy *string `json:"export-policy,omitempty"`
	Importpolicy *string `json:"import-policy,omitempty"`
}

// DeviceSystemMaintenanceProfileIsis struct
type DeviceSystemMaintenanceProfileIsis struct {
	//RootSystemMaintenanceProfileIsisOverload
	Overload *DeviceSystemMaintenanceProfileIsisOverload `json:"overload,omitempty"`
}

// DeviceSystemMaintenanceProfileIsisOverload struct
type DeviceSystemMaintenanceProfileIsisOverload struct {
	// +kubebuilder:default:=false
	Maxmetric *bool `json:"max-metric,omitempty"`
	// +kubebuilder:default:=false
	Setbit *bool `json:"set-bit,omitempty"`
}

// DeviceSystemMirroring struct
type DeviceSystemMirroring struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=4
	Mirroringinstance []*DeviceSystemMirroringMirroringinstance `json:"mirroring-instance,omitempty"`
}

// DeviceSystemMirroringMirroringinstance struct
type DeviceSystemMirroringMirroringinstance struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_DeviceSystemMirroringMirroringinstanceAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Description *string `json:"description,omitempty"`
	//RootSystemMirroringMirroringinstanceMirrordestination
	Mirrordestination *DeviceSystemMirroringMirroringinstanceMirrordestination `json:"mirror-destination,omitempty"`
	//RootSystemMirroringMirroringinstanceMirrorsource
	Mirrorsource *DeviceSystemMirroringMirroringinstanceMirrorsource `json:"mirror-source,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
}

// DeviceSystemMirroringMirroringinstanceMirrordestination struct
type DeviceSystemMirroringMirroringinstanceMirrordestination struct {
	// kubebuilder:validation:MinLength=5
	// kubebuilder:validation:MaxLength=25
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(ethernet-([1-9](\d){0,1}(/[abcd])?(/[1-9](\d){0,1})?/(([1-9](\d){0,1})|(1[0-1]\d)|(12[0-8])))\.([0]|[1-9](\d){0,3})|lag(([1-9](\d){0,1})|(1[0-1]\d)|(12[0-8]))\.(0|[1-9](\d){0,3}))`
	Local *string `json:"local,omitempty"`
	//RootSystemMirroringMirroringinstanceMirrordestinationRemote
	Remote *DeviceSystemMirroringMirroringinstanceMirrordestinationRemote `json:"remote,omitempty"`
}

// DeviceSystemMirroringMirroringinstanceMirrordestinationRemote struct
type DeviceSystemMirroringMirroringinstanceMirrordestinationRemote struct {
	// +kubebuilder:validation:Enum=`l2ogre`
	Encap           E_DeviceSystemMirroringMirroringinstanceMirrordestinationRemoteEncap `json:"encap,omitempty"`
	Networkinstance *string                                                              `json:"network-instance"`
	//RootSystemMirroringMirroringinstanceMirrordestinationRemoteTunnelendpoints
	Tunnelendpoints *DeviceSystemMirroringMirroringinstanceMirrordestinationRemoteTunnelendpoints `json:"tunnel-end-points,omitempty"`
}

// DeviceSystemMirroringMirroringinstanceMirrordestinationRemoteTunnelendpoints struct
type DeviceSystemMirroringMirroringinstanceMirrordestinationRemoteTunnelendpoints struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_DeviceSystemMirroringMirroringinstanceMirrordestinationRemoteTunnelendpointsAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Dstipv4 *string `json:"dst-ipv4"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Srcipv4 *string `json:"src-ipv4"`
}

// DeviceSystemMirroringMirroringinstanceMirrorsource struct
type DeviceSystemMirroringMirroringinstanceMirrorsource struct {
	//RootSystemMirroringMirroringinstanceMirrorsourceAcl
	Acl *DeviceSystemMirroringMirroringinstanceMirrorsourceAcl `json:"acl,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=128
	Interface []*DeviceSystemMirroringMirroringinstanceMirrorsourceInterface `json:"interface,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=128
	Subinterface []*DeviceSystemMirroringMirroringinstanceMirrorsourceSubinterface `json:"subinterface,omitempty"`
}

// DeviceSystemMirroringMirroringinstanceMirrorsourceAcl struct
type DeviceSystemMirroringMirroringinstanceMirrorsourceAcl struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Ipv4filter []*DeviceSystemMirroringMirroringinstanceMirrorsourceAclIpv4filter `json:"ipv4-filter,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Ipv6filter []*DeviceSystemMirroringMirroringinstanceMirrorsourceAclIpv6filter `json:"ipv6-filter,omitempty"`
}

// DeviceSystemMirroringMirroringinstanceMirrorsourceAclIpv4filter struct
type DeviceSystemMirroringMirroringinstanceMirrorsourceAclIpv4filter struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Entry []*DeviceSystemMirroringMirroringinstanceMirrorsourceAclIpv4filterEntry `json:"entry,omitempty"`
	Name  *string                                                                 `json:"name"`
}

// DeviceSystemMirroringMirroringinstanceMirrorsourceAclIpv4filterEntry struct
type DeviceSystemMirroringMirroringinstanceMirrorsourceAclIpv4filterEntry struct {
	Sequenceid *string `json:"sequence-id"`
}

// DeviceSystemMirroringMirroringinstanceMirrorsourceAclIpv6filter struct
type DeviceSystemMirroringMirroringinstanceMirrorsourceAclIpv6filter struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Entry []*DeviceSystemMirroringMirroringinstanceMirrorsourceAclIpv6filterEntry `json:"entry,omitempty"`
	Name  *string                                                                 `json:"name"`
}

// DeviceSystemMirroringMirroringinstanceMirrorsourceAclIpv6filterEntry struct
type DeviceSystemMirroringMirroringinstanceMirrorsourceAclIpv6filterEntry struct {
	Sequenceid *string `json:"sequence-id"`
}

// DeviceSystemMirroringMirroringinstanceMirrorsourceInterface struct
type DeviceSystemMirroringMirroringinstanceMirrorsourceInterface struct {
	// +kubebuilder:validation:Enum=`egress-only`;`ingress-egress`;`ingress-only`
	// +kubebuilder:default:="egress-only"
	Direction E_DeviceSystemMirroringMirroringinstanceMirrorsourceInterfaceDirection `json:"direction,omitempty"`
	//Direction *string `json:"direction,omitempty"`
	// kubebuilder:validation:MinLength=3
	// kubebuilder:validation:MaxLength=20
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(ethernet-([1-9](\d){0,1}(/[abcd])?(/[1-9](\d){0,1})?/(([1-9](\d){0,1})|(1[0-1]\d)|(12[0-8])))|lag(([1-9](\d){0,1})|(1[0-1]\d)|(12[0-8])))`
	Name *string `json:"name"`
}

// DeviceSystemMirroringMirroringinstanceMirrorsourceSubinterface struct
type DeviceSystemMirroringMirroringinstanceMirrorsourceSubinterface struct {
	// +kubebuilder:validation:Enum=`egress-only`;`ingress-egress`;`ingress-only`
	// +kubebuilder:default:="egress-only"
	Direction E_DeviceSystemMirroringMirroringinstanceMirrorsourceSubinterfaceDirection `json:"direction,omitempty"`
	//Direction *string `json:"direction,omitempty"`
	// kubebuilder:validation:MinLength=5
	// kubebuilder:validation:MaxLength=25
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(ethernet-([1-9](\d){0,1}(/[abcd])?(/[1-9](\d){0,1})?/(([1-9](\d){0,1})|(1[0-1]\d)|(12[0-8])))\.([0]|[1-9](\d){0,3})|lag(([1-9](\d){0,1})|(1[0-1]\d)|(12[0-8]))\.(0|[1-9](\d){0,3}))`
	Name *string `json:"name"`
}

// DeviceSystemMpls struct
type DeviceSystemMpls struct {
	//RootSystemMplsLabelranges
	Labelranges *DeviceSystemMplsLabelranges `json:"label-ranges,omitempty"`
}

// DeviceSystemMplsLabelranges struct
type DeviceSystemMplsLabelranges struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Dynamic []*DeviceSystemMplsLabelrangesDynamic `json:"dynamic,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Static []*DeviceSystemMplsLabelrangesStatic `json:"static,omitempty"`
}

// DeviceSystemMplsLabelrangesDynamic struct
type DeviceSystemMplsLabelrangesDynamic struct {
	// kubebuilder:validation:Minimum=16
	// kubebuilder:validation:Maximum=1048575
	Endlabel *uint32 `json:"end-label"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
	// kubebuilder:validation:Minimum=16
	// kubebuilder:validation:Maximum=1048575
	Startlabel *uint32 `json:"start-label"`
}

// DeviceSystemMplsLabelrangesStatic struct
type DeviceSystemMplsLabelrangesStatic struct {
	// kubebuilder:validation:Minimum=16
	// kubebuilder:validation:Maximum=1048575
	Endlabel *uint32 `json:"end-label"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
	// +kubebuilder:default:=true
	Shared *bool `json:"shared,omitempty"`
	// kubebuilder:validation:Minimum=16
	// kubebuilder:validation:Maximum=1048575
	Startlabel *uint32 `json:"start-label"`
}

// DeviceSystemMtu struct
type DeviceSystemMtu struct {
	// kubebuilder:validation:Minimum=1280
	// kubebuilder:validation:Maximum=9486
	// +kubebuilder:default:=1500
	Defaultipmtu *uint16 `json:"default-ip-mtu,omitempty"`
	// kubebuilder:validation:Minimum=1500
	// kubebuilder:validation:Maximum=9500
	// +kubebuilder:default:=9232
	Defaultl2mtu *uint16 `json:"default-l2-mtu,omitempty"`
	// kubebuilder:validation:Minimum=1284
	// kubebuilder:validation:Maximum=9496
	// +kubebuilder:default:=1508
	Defaultmplsmtu *uint16 `json:"default-mpls-mtu,omitempty"`
	// kubebuilder:validation:Minimum=1500
	// kubebuilder:validation:Maximum=9500
	// +kubebuilder:default:=9232
	Defaultportmtu *uint16 `json:"default-port-mtu,omitempty"`
	// kubebuilder:validation:Minimum=552
	// kubebuilder:validation:Maximum=9232
	// +kubebuilder:default:=552
	Minpathmtu *uint16 `json:"min-path-mtu,omitempty"`
}

// DeviceSystemName struct
type DeviceSystemName struct {
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=253
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((([a-zA-Z0-9_]([a-zA-Z0-9\-_]){0,61})?[a-zA-Z0-9]\.)*([a-zA-Z0-9_]([a-zA-Z0-9\-_]){0,61})?[a-zA-Z0-9]\.?)|\.`
	Domainname *string `json:"domain-name,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=63
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\-]*[A-Za-z0-9])`
	Hostname *string `json:"host-name,omitempty"`
}

// DeviceSystemNetworkinstance struct
type DeviceSystemNetworkinstance struct {
	//RootSystemNetworkinstanceProtocols
	Protocols *DeviceSystemNetworkinstanceProtocols `json:"protocols,omitempty"`
}

// DeviceSystemNetworkinstanceProtocols struct
type DeviceSystemNetworkinstanceProtocols struct {
	//RootSystemNetworkinstanceProtocolsBgpvpn
	Bgpvpn *DeviceSystemNetworkinstanceProtocolsBgpvpn `json:"bgp-vpn,omitempty"`
	//RootSystemNetworkinstanceProtocolsEvpn
	Evpn *DeviceSystemNetworkinstanceProtocolsEvpn `json:"evpn,omitempty"`
}

// DeviceSystemNetworkinstanceProtocolsBgpvpn struct
type DeviceSystemNetworkinstanceProtocolsBgpvpn struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1
	Bgpinstance []*DeviceSystemNetworkinstanceProtocolsBgpvpnBgpinstance `json:"bgp-instance,omitempty"`
}

// DeviceSystemNetworkinstanceProtocolsBgpvpnBgpinstance struct
type DeviceSystemNetworkinstanceProtocolsBgpvpnBgpinstance struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=2
	Id *uint8 `json:"id"`
	//RootSystemNetworkinstanceProtocolsBgpvpnBgpinstanceRoutedistinguisher
	Routedistinguisher *DeviceSystemNetworkinstanceProtocolsBgpvpnBgpinstanceRoutedistinguisher `json:"route-distinguisher,omitempty"`
	//RootSystemNetworkinstanceProtocolsBgpvpnBgpinstanceRoutetarget
	Routetarget *DeviceSystemNetworkinstanceProtocolsBgpvpnBgpinstanceRoutetarget `json:"route-target,omitempty"`
}

// DeviceSystemNetworkinstanceProtocolsBgpvpnBgpinstanceRoutedistinguisher struct
type DeviceSystemNetworkinstanceProtocolsBgpvpnBgpinstanceRoutedistinguisher struct {
}

// DeviceSystemNetworkinstanceProtocolsBgpvpnBgpinstanceRoutetarget struct
type DeviceSystemNetworkinstanceProtocolsBgpvpnBgpinstanceRoutetarget struct {
}

// DeviceSystemNetworkinstanceProtocolsEvpn struct
type DeviceSystemNetworkinstanceProtocolsEvpn struct {
	//RootSystemNetworkinstanceProtocolsEvpnEthernetsegments
	Ethernetsegments *DeviceSystemNetworkinstanceProtocolsEvpnEthernetsegments `json:"ethernet-segments,omitempty"`
}

// DeviceSystemNetworkinstanceProtocolsEvpnEthernetsegments struct
type DeviceSystemNetworkinstanceProtocolsEvpnEthernetsegments struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1
	Bgpinstance []*DeviceSystemNetworkinstanceProtocolsEvpnEthernetsegmentsBgpinstance `json:"bgp-instance,omitempty"`
	//RootSystemNetworkinstanceProtocolsEvpnEthernetsegmentsTimers
	Timers *DeviceSystemNetworkinstanceProtocolsEvpnEthernetsegmentsTimers `json:"timers,omitempty"`
}

// DeviceSystemNetworkinstanceProtocolsEvpnEthernetsegmentsBgpinstance struct
type DeviceSystemNetworkinstanceProtocolsEvpnEthernetsegmentsBgpinstance struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=128
	Esi []*DeviceSystemNetworkinstanceProtocolsEvpnEthernetsegmentsBgpinstanceEthernetsegment `json:"esi,omitempty"`
	Id  *string                                                                               `json:"id"`
}

// DeviceSystemNetworkinstanceProtocolsEvpnEthernetsegmentsBgpinstanceEthernetsegment struct
type DeviceSystemNetworkinstanceProtocolsEvpnEthernetsegmentsBgpinstanceEthernetsegment struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_DeviceSystemNetworkinstanceProtocolsEvpnEthernetsegmentsBgpinstanceEthernetsegmentAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	//RootSystemNetworkinstanceProtocolsEvpnEthernetsegmentsBgpinstanceEthernetsegmentDfelection
	Dfelection *DeviceSystemNetworkinstanceProtocolsEvpnEthernetsegmentsBgpinstanceEthernetsegmentDfelection `json:"df-election,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){9}`
	Esi       *string `json:"esi,omitempty"`
	Interface *string `json:"interface,omitempty"`
	// +kubebuilder:validation:Enum=`all-active`;`single-active`
	// +kubebuilder:default:="all-active"
	Multihomingmode E_DeviceSystemNetworkinstanceProtocolsEvpnEthernetsegmentsBgpinstanceEthernetsegmentMultihomingmode `json:"multi-homing-mode,omitempty"`
	//Multihomingmode *string `json:"multi-homing-mode,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=32
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
	//RootSystemNetworkinstanceProtocolsEvpnEthernetsegmentsBgpinstanceEthernetsegmentRoutes
	Routes *DeviceSystemNetworkinstanceProtocolsEvpnEthernetsegmentsBgpinstanceEthernetsegmentRoutes `json:"routes,omitempty"`
}

// DeviceSystemNetworkinstanceProtocolsEvpnEthernetsegmentsBgpinstanceEthernetsegmentDfelection struct
type DeviceSystemNetworkinstanceProtocolsEvpnEthernetsegmentsBgpinstanceEthernetsegmentDfelection struct {
	//RootSystemNetworkinstanceProtocolsEvpnEthernetsegmentsBgpinstanceEthernetsegmentDfelectionAlgorithm
	Algorithm *DeviceSystemNetworkinstanceProtocolsEvpnEthernetsegmentsBgpinstanceEthernetsegmentDfelectionAlgorithm `json:"algorithm,omitempty"`
	//RootSystemNetworkinstanceProtocolsEvpnEthernetsegmentsBgpinstanceEthernetsegmentDfelectionInterfacestandbysignalingonnondf
	Interfacestandbysignalingonnondf *DeviceSystemNetworkinstanceProtocolsEvpnEthernetsegmentsBgpinstanceEthernetsegmentDfelectionInterfacestandbysignalingonnondf `json:"interface-standby-signaling-on-non-df,omitempty"`
	//RootSystemNetworkinstanceProtocolsEvpnEthernetsegmentsBgpinstanceEthernetsegmentDfelectionTimers
	Timers *DeviceSystemNetworkinstanceProtocolsEvpnEthernetsegmentsBgpinstanceEthernetsegmentDfelectionTimers `json:"timers,omitempty"`
}

// DeviceSystemNetworkinstanceProtocolsEvpnEthernetsegmentsBgpinstanceEthernetsegmentDfelectionAlgorithm struct
type DeviceSystemNetworkinstanceProtocolsEvpnEthernetsegmentsBgpinstanceEthernetsegmentDfelectionAlgorithm struct {
	//RootSystemNetworkinstanceProtocolsEvpnEthernetsegmentsBgpinstanceEthernetsegmentDfelectionAlgorithmPreferencealg
	Preferencealg *DeviceSystemNetworkinstanceProtocolsEvpnEthernetsegmentsBgpinstanceEthernetsegmentDfelectionAlgorithmPreferencealg `json:"preference-alg,omitempty"`
	// +kubebuilder:validation:Enum=`default`;`preference`
	// +kubebuilder:default:="default"
	Type E_DeviceSystemNetworkinstanceProtocolsEvpnEthernetsegmentsBgpinstanceEthernetsegmentDfelectionAlgorithmType `json:"type,omitempty"`
	//Type *string `json:"type,omitempty"`
}

// DeviceSystemNetworkinstanceProtocolsEvpnEthernetsegmentsBgpinstanceEthernetsegmentDfelectionAlgorithmPreferencealg struct
type DeviceSystemNetworkinstanceProtocolsEvpnEthernetsegmentsBgpinstanceEthernetsegmentDfelectionAlgorithmPreferencealg struct {
	//RootSystemNetworkinstanceProtocolsEvpnEthernetsegmentsBgpinstanceEthernetsegmentDfelectionAlgorithmPreferencealgCapabilities
	Capabilities *DeviceSystemNetworkinstanceProtocolsEvpnEthernetsegmentsBgpinstanceEthernetsegmentDfelectionAlgorithmPreferencealgCapabilities `json:"capabilities,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=32767
	Preferencevalue *uint32 `json:"preference-value,omitempty"`
}

// DeviceSystemNetworkinstanceProtocolsEvpnEthernetsegmentsBgpinstanceEthernetsegmentDfelectionAlgorithmPreferencealgCapabilities struct
type DeviceSystemNetworkinstanceProtocolsEvpnEthernetsegmentsBgpinstanceEthernetsegmentDfelectionAlgorithmPreferencealgCapabilities struct {
	// +kubebuilder:validation:Enum=`exclude`;`include`
	// +kubebuilder:default:="include"
	Acdf E_DeviceSystemNetworkinstanceProtocolsEvpnEthernetsegmentsBgpinstanceEthernetsegmentDfelectionAlgorithmPreferencealgCapabilitiesAcdf `json:"ac-df,omitempty"`
	//Acdf *string `json:"ac-df,omitempty"`
	// +kubebuilder:default:=false
	Nonrevertive *bool `json:"non-revertive,omitempty"`
}

// DeviceSystemNetworkinstanceProtocolsEvpnEthernetsegmentsBgpinstanceEthernetsegmentDfelectionInterfacestandbysignalingonnondf struct
type DeviceSystemNetworkinstanceProtocolsEvpnEthernetsegmentsBgpinstanceEthernetsegmentDfelectionInterfacestandbysignalingonnondf struct {
}

// DeviceSystemNetworkinstanceProtocolsEvpnEthernetsegmentsBgpinstanceEthernetsegmentDfelectionTimers struct
type DeviceSystemNetworkinstanceProtocolsEvpnEthernetsegmentsBgpinstanceEthernetsegmentDfelectionTimers struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100
	Activationtimer *uint32 `json:"activation-timer,omitempty"`
}

// DeviceSystemNetworkinstanceProtocolsEvpnEthernetsegmentsBgpinstanceEthernetsegmentRoutes struct
type DeviceSystemNetworkinstanceProtocolsEvpnEthernetsegmentsBgpinstanceEthernetsegmentRoutes struct {
	//RootSystemNetworkinstanceProtocolsEvpnEthernetsegmentsBgpinstanceEthernetsegmentRoutesEthernetsegment
	Esi *DeviceSystemNetworkinstanceProtocolsEvpnEthernetsegmentsBgpinstanceEthernetsegmentRoutesEthernetsegment `json:"esi,omitempty"`
	// +kubebuilder:validation:Enum=`use-system-ipv4-address`
	// +kubebuilder:default:="use-system-ipv4-address"
	Nexthop E_DeviceSystemNetworkinstanceProtocolsEvpnEthernetsegmentsBgpinstanceEthernetsegmentRoutesNexthop `json:"next-hop,omitempty"`
	//Nexthop *string `json:"next-hop,omitempty"`
}

// DeviceSystemNetworkinstanceProtocolsEvpnEthernetsegmentsBgpinstanceEthernetsegmentRoutesEthernetsegment struct
type DeviceSystemNetworkinstanceProtocolsEvpnEthernetsegmentsBgpinstanceEthernetsegmentRoutesEthernetsegment struct {
	// +kubebuilder:validation:Enum=`use-system-ipv4-address`
	// +kubebuilder:default:="use-system-ipv4-address"
	Originatingip E_DeviceSystemNetworkinstanceProtocolsEvpnEthernetsegmentsBgpinstanceEthernetsegmentRoutesEthernetsegmentOriginatingip `json:"originating-ip,omitempty"`
	//Originatingip *string `json:"originating-ip,omitempty"`
}

// DeviceSystemNetworkinstanceProtocolsEvpnEthernetsegmentsTimers struct
type DeviceSystemNetworkinstanceProtocolsEvpnEthernetsegmentsTimers struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=3
	Activationtimer *uint32 `json:"activation-timer,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=6000
	// +kubebuilder:default:=10
	Boottimer *uint32 `json:"boot-timer,omitempty"`
}

// DeviceSystemNtp struct
type DeviceSystemNtp struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	Adminstate E_DeviceSystemNtpAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	Networkinstance *string `json:"network-instance"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Server []*DeviceSystemNtpServer `json:"server,omitempty"`
}

// DeviceSystemNtpServer struct
type DeviceSystemNtpServer struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Address *string `json:"address"`
	// +kubebuilder:default:=false
	Iburst *bool `json:"iburst,omitempty"`
	// +kubebuilder:default:=false
	Prefer *bool `json:"prefer,omitempty"`
}

// DeviceSystemP4rtserver struct
type DeviceSystemP4rtserver struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_DeviceSystemP4rtserverAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Networkinstance []*DeviceSystemP4rtserverNetworkinstance `json:"network-instance,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=60
	Ratelimit *uint16 `json:"rate-limit,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=20
	Sessionlimit *uint16 `json:"session-limit,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=7200
	Timeout *uint16 `json:"timeout,omitempty"`
	//RootSystemP4rtserverTraceoptions
	Traceoptions *DeviceSystemP4rtserverTraceoptions `json:"trace-options,omitempty"`
	//RootSystemP4rtserverUnixsocket
	Unixsocket *DeviceSystemP4rtserverUnixsocket `json:"unix-socket,omitempty"`
}

// DeviceSystemP4rtserverNetworkinstance struct
type DeviceSystemP4rtserverNetworkinstance struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_DeviceSystemP4rtserverNetworkinstanceAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	Name *string `json:"name"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=9559
	Port *uint16 `json:"port,omitempty"`
	//RootSystemP4rtserverNetworkinstanceSourceaddress
	Sourceaddress *DeviceSystemP4rtserverNetworkinstanceSourceaddress `json:"source-address,omitempty"`
	Tlsprofile    *string                                             `json:"tls-profile"`
	// +kubebuilder:default:=true
	Useauthentication *bool `json:"use-authentication,omitempty"`
}

// DeviceSystemP4rtserverNetworkinstanceSourceaddress struct
type DeviceSystemP4rtserverNetworkinstanceSourceaddress struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Sourceaddress *string `json:"source-address,omitempty"`
}

// DeviceSystemP4rtserverTraceoptions struct
type DeviceSystemP4rtserverTraceoptions struct {
	// +kubebuilder:validation:Enum=`common`;`request`;`response`
	Traceoptions E_DeviceSystemP4rtserverTraceoptionsTraceoptions `json:"trace-options,omitempty"`
	//Traceoptions *string `json:"trace-options,omitempty"`
}

// DeviceSystemP4rtserverUnixsocket struct
type DeviceSystemP4rtserverUnixsocket struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_DeviceSystemP4rtserverUnixsocketAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	Tlsprofile *string `json:"tls-profile,omitempty"`
	// +kubebuilder:default:=true
	Useauthentication *bool `json:"use-authentication,omitempty"`
}

// DeviceSystemRaguardpolicy struct
type DeviceSystemRaguardpolicy struct {
	// +kubebuilder:validation:Enum=`accept`;`discard`
	// +kubebuilder:default:="discard"
	Action E_DeviceSystemRaguardpolicyAction `json:"action,omitempty"`
	//Action *string `json:"action,omitempty"`
	Advertiseprefixset *string `json:"advertise-prefix-set,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=255
	Hoplimit          *uint8 `json:"hop-limit,omitempty"`
	Managedconfigflag *bool  `json:"managed-config-flag,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name            *string `json:"name"`
	Otherconfigflag *bool   `json:"other-config-flag,omitempty"`
	// +kubebuilder:validation:Enum=`high`;`low`;`medium`
	Routerpreference E_DeviceSystemRaguardpolicyRouterpreference `json:"router-preference,omitempty"`
	//Routerpreference *string `json:"router-preference,omitempty"`
	Sourceprefixset *string `json:"source-prefix-set,omitempty"`
}

// DeviceSystemSflow struct
type DeviceSystemSflow struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_DeviceSystemSflowAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=8
	Collector []*DeviceSystemSflowCollector `json:"collector,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=2000000
	// +kubebuilder:default:=10000
	Samplerate *uint32 `json:"sample-rate,omitempty"`
	// kubebuilder:validation:Minimum=256
	// kubebuilder:validation:Maximum=256
	// +kubebuilder:default:=256
	Samplesize *uint16 `json:"sample-size,omitempty"`
}

// DeviceSystemSflowCollector struct
type DeviceSystemSflowCollector struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Collectoraddress *string `json:"collector-address,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=8
	Collectorid     *uint16 `json:"collector-id"`
	Networkinstance *string `json:"network-instance"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=6343
	Port *uint16 `json:"port,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Sourceaddress *string `json:"source-address"`
}

// DeviceSystemSnmp struct
type DeviceSystemSnmp struct {
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Community *string `json:"community"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Networkinstance []*DeviceSystemSnmpNetworkinstance `json:"network-instance,omitempty"`
}

// DeviceSystemSnmpNetworkinstance struct
type DeviceSystemSnmpNetworkinstance struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	Adminstate E_DeviceSystemSnmpNetworkinstanceAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	Name *string `json:"name"`
	//RootSystemSnmpNetworkinstanceSourceaddress
	Sourceaddress *DeviceSystemSnmpNetworkinstanceSourceaddress `json:"source-address,omitempty"`
}

// DeviceSystemSnmpNetworkinstanceSourceaddress struct
type DeviceSystemSnmpNetworkinstanceSourceaddress struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	// +kubebuilder:default:="::"
	Sourceaddress *string `json:"source-address,omitempty"`
}

// DeviceSystemSshserver struct
type DeviceSystemSshserver struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Networkinstance []*DeviceSystemSshserverNetworkinstance `json:"network-instance,omitempty"`
}

// DeviceSystemSshserverNetworkinstance struct
type DeviceSystemSshserverNetworkinstance struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	Adminstate E_DeviceSystemSshserverNetworkinstanceAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	Name *string `json:"name"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=20
	Ratelimit *uint8 `json:"rate-limit,omitempty"`
	//RootSystemSshserverNetworkinstanceSourceaddress
	Sourceaddress *DeviceSystemSshserverNetworkinstanceSourceaddress `json:"source-address,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=0
	Timeout *uint16 `json:"timeout,omitempty"`
}

// DeviceSystemSshserverNetworkinstanceSourceaddress struct
type DeviceSystemSshserverNetworkinstanceSourceaddress struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Sourceaddress *string `json:"source-address,omitempty"`
}

// DeviceSystemSync struct
type DeviceSystemSync struct {
	//RootSystemSyncBits
	Bits *DeviceSystemSyncBits `json:"bits,omitempty"`
	//RootSystemSyncFreqclock
	Freqclock *DeviceSystemSyncFreqclock `json:"freq-clock,omitempty"`
	//RootSystemSyncFreqreferences
	Freqreferences *DeviceSystemSyncFreqreferences `json:"freq-references,omitempty"`
	//RootSystemSyncPpstest
	Ppstest *DeviceSystemSyncPpstest `json:"ppstest,omitempty"`
	//RootSystemSyncPtp
	Ptp *DeviceSystemSyncPtp `json:"ptp,omitempty"`
	//RootSystemSyncTod
	Tod *DeviceSystemSyncTod `json:"tod,omitempty"`
}

// DeviceSystemSyncBits struct
type DeviceSystemSyncBits struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_DeviceSystemSyncBitsAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=40
	Alarmprofile *string `json:"alarm-profile,omitempty"`
	// +kubebuilder:validation:Enum=`input`;`output`
	// +kubebuilder:default:="input"
	Direction E_DeviceSystemSyncBitsDirection `json:"direction,omitempty"`
	//Direction *string `json:"direction,omitempty"`
	// +kubebuilder:validation:Enum=`output-selection`;`setg`
	// +kubebuilder:default:="setg"
	Freqoutputselect E_DeviceSystemSyncBitsFreqoutputselect `json:"freq-output-select,omitempty"`
	//Freqoutputselect *string `json:"freq-output-select,omitempty"`
	// kubebuilder:validation:Minimum=75
	// kubebuilder:validation:Maximum=75
	// +kubebuilder:default:=120
	Impedance *uint8 `json:"impedance,omitempty"`
	// +kubebuilder:validation:Enum=`133`;`266`;`399`;`533`;`655`;`none`
	// +kubebuilder:default:="133"
	Lbo E_DeviceSystemSyncBitsLbo `json:"lbo,omitempty"`
	//Lbo *string `json:"lbo,omitempty"`
	// +kubebuilder:validation:Enum=`ami`;`b8zs`
	// +kubebuilder:default:="b8zs"
	Linecode E_DeviceSystemSyncBitsLinecode `json:"line-code,omitempty"`
	//Linecode *string `json:"line-code,omitempty"`
	// +kubebuilder:validation:Enum=`auto`;`prc`;`prs`;`sec`;`ssu-a`;`ssu-b`;`st2`;`st3`;`st3e`;`stu`;`tnc`
	Qloutputoverride E_DeviceSystemSyncBitsQloutputoverride `json:"ql-output-override,omitempty"`
	//Qloutputoverride *string `json:"ql-output-override,omitempty"`
	// +kubebuilder:validation:Enum=`dnu`;`not-applicable`;`pno`;`prc`;`prs`;`sec`;`smc`;`ssu-a`;`ssu-b`;`st2`;`st3`;`st3e`;`st4`;`stu`;`tnc`;`unknown`
	// +kubebuilder:default:="st3"
	Qlsquelchthreshold E_DeviceSystemSyncBitsQlsquelchthreshold `json:"ql-squelch-threshold,omitempty"`
	//Qlsquelchthreshold *string `json:"ql-squelch-threshold,omitempty"`
	// kubebuilder:validation:Minimum=4
	// kubebuilder:validation:Maximum=8
	// +kubebuilder:default:=4
	Sabit *uint8 `json:"sa-bit,omitempty"`
	// +kubebuilder:validation:Enum=`2mhz`;`ds1-esf`;`ds1-sf`;`framed-e1`;`framed-e1-ssm`;`unframed-e1`
	// +kubebuilder:default:="ds1-esf"
	Signaltype E_DeviceSystemSyncBitsSignaltype `json:"signal-type,omitempty"`
	//Signaltype *string `json:"signal-type,omitempty"`
	// +kubebuilder:validation:Enum=`ais`;`ql`
	// +kubebuilder:default:="ql"
	Squelchmode E_DeviceSystemSyncBitsSquelchmode `json:"squelch-mode,omitempty"`
	//Squelchmode *string `json:"squelch-mode,omitempty"`
}

// DeviceSystemSyncFreqclock struct
type DeviceSystemSyncFreqclock struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_DeviceSystemSyncFreqclockAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=40
	Alarmprofile *string `json:"alarm-profile,omitempty"`
	// +kubebuilder:validation:Enum=`dnu`;`not-applicable`;`pno`;`prc`;`prs`;`sec`;`smc`;`ssu-a`;`ssu-b`;`st2`;`st3`;`st3e`;`st4`;`stu`;`tnc`;`unknown`
	// +kubebuilder:default:="st3"
	Qlinputthreshold E_DeviceSystemSyncFreqclockQlinputthreshold `json:"ql-input-threshold,omitempty"`
	//Qlinputthreshold *string `json:"ql-input-threshold,omitempty"`
	// +kubebuilder:validation:Enum=`ansi`;`etsi`
	// +kubebuilder:default:="ansi"
	Syncemode E_DeviceSystemSyncFreqclockSyncemode `json:"synce-mode,omitempty"`
	//Syncemode *string `json:"synce-mode,omitempty"`
	// +kubebuilder:default:=true
	Systemqlenable *bool `json:"system-ql-enable,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=12
	// +kubebuilder:default:=5
	Waittorestore *int8 `json:"wait-to-restore,omitempty"`
}

// DeviceSystemSyncFreqreferences struct
type DeviceSystemSyncFreqreferences struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Instancelist []*DeviceSystemSyncFreqreferencesInstancelist `json:"instance-list,omitempty"`
}

// DeviceSystemSyncFreqreferencesInstancelist struct
type DeviceSystemSyncFreqreferencesInstancelist struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_DeviceSystemSyncFreqreferencesInstancelistAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=40
	Alarmprofile *string `json:"alarm-profile,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=6
	Instancenumber *uint8 `json:"instance-number"`
	// +kubebuilder:default:=false
	Lockoutref *bool `json:"lockout-ref,omitempty"`
	// +kubebuilder:default:=false
	Lockoutrefbitsout *bool `json:"lockout-ref-bits-out,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=7
	// +kubebuilder:default:=0
	Priority *uint8 `json:"priority,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=7
	// +kubebuilder:default:=0
	Priorityoutput *uint8 `json:"priority-output,omitempty"`
	// +kubebuilder:validation:Enum=`auto`;`prc`;`prs`;`sec`;`ssu-a`;`ssu-b`;`st2`;`st3`;`st3e`;`stu`;`tnc`
	// +kubebuilder:default:="auto"
	Qloverride E_DeviceSystemSyncFreqreferencesInstancelistQloverride `json:"ql-override,omitempty"`
	//Qloverride *string `json:"ql-override,omitempty"`
	// +kubebuilder:default:=true
	Refqlenable *bool `json:"ref-ql-enable,omitempty"`
	// +kubebuilder:validation:Enum=`force`;`manual`;`none`
	// +kubebuilder:default:="none"
	Switchrefrequest E_DeviceSystemSyncFreqreferencesInstancelistSwitchrefrequest `json:"switch-ref-request,omitempty"`
	//Switchrefrequest *string `json:"switch-ref-request,omitempty"`
	// +kubebuilder:validation:Enum=`force`;`manual`;`none`
	// +kubebuilder:default:="none"
	Switchrefrequestbitsout E_DeviceSystemSyncFreqreferencesInstancelistSwitchrefrequestbitsout `json:"switch-ref-request-bits-out,omitempty"`
	//Switchrefrequestbitsout *string `json:"switch-ref-request-bits-out,omitempty"`
	// kubebuilder:validation:MinLength=3
	// kubebuilder:validation:MaxLength=20
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(ethernet-([1-9](\d){0,1}(/[abcd])?(/[1-9](\d){0,1})?/(([1-9](\d){0,1})|(1[0-1]\d)|(12[0-8])))|otc-1/([1-8])|BITS|ethernet-1/x([1-4]))`
	Underlyinginterface *string `json:"underlying-interface,omitempty"`
}

// DeviceSystemSyncPpstest struct
type DeviceSystemSyncPpstest struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_DeviceSystemSyncPpstestAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=127
	// kubebuilder:validation:Maximum=127
	// +kubebuilder:default:=0
	Outputcompensation *int8 `json:"output-compensation,omitempty"`
}

// DeviceSystemSyncPtp struct
type DeviceSystemSyncPtp struct {
	// +kubebuilder:default:="layer1a"
	Frequencyref *string `json:"frequency-ref,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Instancelist []*DeviceSystemSyncPtpInstancelist `json:"instance-list,omitempty"`
	// +kubebuilder:validation:Enum=`itug8275dot1`
	// +kubebuilder:default:="itug8275dot1"
	Ptpprofile E_DeviceSystemSyncPtpPtpprofile `json:"ptp-profile,omitempty"`
	//Ptpprofile *string `json:"ptp-profile,omitempty"`
	// +kubebuilder:validation:Enum=`extTod`;`ptp`
	// +kubebuilder:default:="ptp"
	Timeref E_DeviceSystemSyncPtpTimeref `json:"time-ref,omitempty"`
	//Timeref *string `json:"time-ref,omitempty"`
	// +kubebuilder:validation:Enum=`ccsa`;`itu`
	// +kubebuilder:default:="itu"
	Todmsgtype E_DeviceSystemSyncPtpTodmsgtype `json:"tod-msg-type,omitempty"`
	//Todmsgtype *string `json:"tod-msg-type,omitempty"`
}

// DeviceSystemSyncPtpInstancelist struct
type DeviceSystemSyncPtpInstancelist struct {
	// +kubebuilder:validation:Enum=`ptpdisabled`;`tbc`;`ttsc`
	// +kubebuilder:default:="ptpdisabled"
	Clockmode E_DeviceSystemSyncPtpInstancelistClockmode `json:"clock-mode,omitempty"`
	//Clockmode *string `json:"clock-mode,omitempty"`
	//RootSystemSyncPtpInstancelistDefaultds
	Defaultds *DeviceSystemSyncPtpInstancelistDefaultds `json:"default-ds,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=10
	Instancenumber *uint32 `json:"instance-number"`
	//RootSystemSyncPtpInstancelistParentds
	Parentds *DeviceSystemSyncPtpInstancelistParentds `json:"parent-ds,omitempty"`
	//RootSystemSyncPtpInstancelistPtpportdataset
	Ptpportdataset *DeviceSystemSyncPtpInstancelistPtpportdataset `json:"ptp-port-data-set,omitempty"`
	//RootSystemSyncPtpInstancelistTimepropertiesds
	Timepropertiesds *DeviceSystemSyncPtpInstancelistTimepropertiesds `json:"time-properties-ds,omitempty"`
}

// DeviceSystemSyncPtpInstancelistDefaultds struct
type DeviceSystemSyncPtpInstancelistDefaultds struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_DeviceSystemSyncPtpInstancelistDefaultdsAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=24
	// kubebuilder:validation:Maximum=43
	// +kubebuilder:default:=24
	Domainnumber *uint8 `json:"domain-number,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=128
	Localpriority *uint8 `json:"local-priority,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=255
	Maxstepsremoved *uint8 `json:"max-steps-removed,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=128
	Priority1 *uint8 `json:"priority1,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=128
	Priority2 *uint8 `json:"priority2,omitempty"`
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Syncuncertain E_DeviceSystemSyncPtpInstancelistDefaultdsSyncuncertain `json:"sync-uncertain,omitempty"`
	//Syncuncertain *string `json:"sync-uncertain,omitempty"`
}

// DeviceSystemSyncPtpInstancelistParentds struct
type DeviceSystemSyncPtpInstancelistParentds struct {
	// +kubebuilder:default:=false
	Parentstats *bool `json:"parent-stats,omitempty"`
}

// DeviceSystemSyncPtpInstancelistPtpportdataset struct
type DeviceSystemSyncPtpInstancelistPtpportdataset struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Portdslist []*DeviceSystemSyncPtpInstancelistPtpportdatasetPortdslist `json:"port-ds-list,omitempty"`
}

// DeviceSystemSyncPtpInstancelistPtpportdatasetPortdslist struct
type DeviceSystemSyncPtpInstancelistPtpportdatasetPortdslist struct {
	// +kubebuilder:validation:Enum=`multicast`
	// +kubebuilder:default:="multicast"
	Addressingmode E_DeviceSystemSyncPtpInstancelistPtpportdatasetPortdslistAddressingmode `json:"addressing-mode,omitempty"`
	//Addressingmode *string `json:"addressing-mode,omitempty"`
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_DeviceSystemSyncPtpInstancelistPtpportdatasetPortdslistAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=40
	Alarmprofile *string `json:"alarm-profile,omitempty"`
	// kubebuilder:validation:Minimum=2
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=3
	Announcereceipttimeout *uint8 `json:"announce-receipt-timeout,omitempty"`
	// kubebuilder:validation:Minimum=100000000
	// kubebuilder:validation:Maximum=100000000
	// +kubebuilder:default:=0
	Asymcorrection *int32 `json:"asym-correction,omitempty"`
	// +kubebuilder:validation:Enum=`forwardable`;`non-forwardable`
	// +kubebuilder:default:="forwardable"
	Destmac E_DeviceSystemSyncPtpInstancelistPtpportdatasetPortdslistDestmac `json:"dest-mac,omitempty"`
	//Destmac *string `json:"dest-mac,omitempty"`
	// +kubebuilder:validation:Enum=`ethernet`
	// +kubebuilder:default:="ethernet"
	Encaptype E_DeviceSystemSyncPtpInstancelistPtpportdatasetPortdslistEncaptype `json:"encap-type,omitempty"`
	//Encaptype *string `json:"encap-type,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=128
	Localpriority *uint8 `json:"local-priority,omitempty"`
	// kubebuilder:validation:Minimum=4
	// kubebuilder:validation:Maximum=4
	// +kubebuilder:default:=-3
	Logannounceinterval *int8 `json:"log-announce-interval,omitempty"`
	// kubebuilder:validation:Minimum=7
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=-4
	Logmindelayreqinterval *int8 `json:"log-min-delay-req-interval,omitempty"`
	// kubebuilder:validation:Minimum=7
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=-4
	Logsyncinterval *int8 `json:"log-sync-interval,omitempty"`
	// +kubebuilder:default:=true
	Masteronly *bool `json:"master-only,omitempty"`
	// kubebuilder:validation:Minimum=-9223372036854775808
	// kubebuilder:validation:Maximum=4294967295
	// +kubebuilder:default:=0
	Peermeanpathdelay *int64 `json:"peer-mean-path-delay,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	Portnumber *uint16 `json:"port-number"`
	// +kubebuilder:validation:Enum=`auto`
	// +kubebuilder:default:="auto"
	Portrole E_DeviceSystemSyncPtpInstancelistPtpportdatasetPortdslistPortrole `json:"port-role,omitempty"`
	//Portrole *string `json:"port-role,omitempty"`
	// +kubebuilder:validation:Enum=`disabled`;`faulty`;`initializing`;`listening`;`master`;`passive`;`pre-master`;`slave`;`uncalibrated`
	// +kubebuilder:default:="initializing"
	Portstate E_DeviceSystemSyncPtpInstancelistPtpportdatasetPortdslistPortstate `json:"port-state,omitempty"`
	//Portstate *string `json:"port-state,omitempty"`
	// +kubebuilder:validation:Enum=`disabled`;`ptp-port-mode`
	// +kubebuilder:default:="ptp-port-mode"
	Ptpentitytype E_DeviceSystemSyncPtpInstancelistPtpportdatasetPortdslistPtpentitytype `json:"ptp-entity-type,omitempty"`
	//Ptpentitytype *string `json:"ptp-entity-type,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=40
	Ptpportalarmprofile *string `json:"ptp-port-alarm-profile,omitempty"`
	//RootSystemSyncPtpInstancelistPtpportdatasetPortdslistStatistics
	Statistics *DeviceSystemSyncPtpInstancelistPtpportdatasetPortdslistStatistics `json:"statistics,omitempty"`
	// kubebuilder:validation:MinLength=3
	// kubebuilder:validation:MaxLength=20
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(ethernet-([1-9](\d){0,1}(/[abcd])?(/[1-9](\d){0,1})?/(([1-9](\d){0,1})|(1[0-1]\d)|(12[0-8])))|otc-1/([1-8])|ethernet-1/x([1-4]))`
	Underlyinginterface *string `json:"underlying-interface,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4094
	// +kubebuilder:default:=0
	Vlanid *uint16 `json:"vlan-id,omitempty"`
}

// DeviceSystemSyncPtpInstancelistPtpportdatasetPortdslistStatistics struct
type DeviceSystemSyncPtpInstancelistPtpportdatasetPortdslistStatistics struct {
}

// DeviceSystemSyncPtpInstancelistTimepropertiesds struct
type DeviceSystemSyncPtpInstancelistTimepropertiesds struct {
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=40
	Alarmprofile *string `json:"alarm-profile,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	Clockclassinputthreshold *uint8 `json:"clockclass-input-threshold,omitempty"`
	// kubebuilder:validation:Minimum=2
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=3
	Delayresplossdetect *uint8 `json:"delay-resp-loss-detect,omitempty"`
	//RootSystemSyncPtpInstancelistTimepropertiesdsPathtracedataset
	Pathtracedataset *DeviceSystemSyncPtpInstancelistTimepropertiesdsPathtracedataset `json:"pathtrace-data-set,omitempty"`
	// kubebuilder:validation:Minimum=2
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=3
	Synclossdetect *uint8 `json:"sync-loss-detect,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=15
	// +kubebuilder:default:=1
	Timeerrormonduration *uint8 `json:"time-error-mon-duration,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=5
	// +kubebuilder:default:=5
	Timeerrormoninterval *uint8 `json:"time-error-mon-interval,omitempty"`
	// +kubebuilder:validation:Enum=`none`;`tod`
	// +kubebuilder:default:="none"
	Timeerrormonref E_DeviceSystemSyncPtpInstancelistTimepropertiesdsTimeerrormonref `json:"time-error-mon-ref,omitempty"`
	//Timeerrormonref *string `json:"time-error-mon-ref,omitempty"`
	// +kubebuilder:validation:Enum=`240`;`241`;`242`;`243`;`244`;`245`;`246`;`247`;`248`;`249`;`250`;`251`;`252`;`253`;`254`;`atomic-clock`;`gps`;`hand-set`;`internal-oscillator`;`ntp`;`other`;`ptp`;`reserved`;`terrestrial-radio`
	Timesource E_DeviceSystemSyncPtpInstancelistTimepropertiesdsTimesource `json:"time-source,omitempty"`
	//Timesource *string `json:"time-source,omitempty"`
}

// DeviceSystemSyncPtpInstancelistTimepropertiesdsPathtracedataset struct
type DeviceSystemSyncPtpInstancelistTimepropertiesdsPathtracedataset struct {
	Pathtraceenable *bool `json:"path-trace-enable,omitempty"`
}

// DeviceSystemSyncTod struct
type DeviceSystemSyncTod struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_DeviceSystemSyncTodAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=40
	Alarmprofile *string `json:"alarm-profile,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100000000
	// +kubebuilder:default:=0
	Cablecorrection *uint32 `json:"cable-correction,omitempty"`
	// +kubebuilder:validation:Enum=`input`;`output`
	// +kubebuilder:default:="input"
	Direction E_DeviceSystemSyncTodDirection `json:"direction,omitempty"`
	//Direction *string `json:"direction,omitempty"`
	// kubebuilder:validation:Minimum=127
	// kubebuilder:validation:Maximum=127
	// +kubebuilder:default:=0
	Inputcompensation *int8 `json:"input-compensation,omitempty"`
	// kubebuilder:validation:Minimum=127
	// kubebuilder:validation:Maximum=127
	// +kubebuilder:default:=0
	Outputcompensation *int8 `json:"output-compensation,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=7
	// +kubebuilder:default:=6
	Todinputthreshold *uint32 `json:"tod-input-threshold,omitempty"`
}

// DeviceSystemTls struct
type DeviceSystemTls struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Serverprofile []*DeviceSystemTlsServerprofile `json:"server-profile,omitempty"`
}

// DeviceSystemTlsServerprofile struct
type DeviceSystemTlsServerprofile struct {
	// +kubebuilder:default:=false
	Authenticateclient *bool   `json:"authenticate-client,omitempty"`
	Certificate        *string `json:"certificate"`
	//RootSystemTlsServerprofileCipherlist
	Cipherlist *DeviceSystemTlsServerprofileCipherlist `json:"cipher-list,omitempty"`
	Key        *string                                 `json:"key"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name        *string `json:"name"`
	Trustanchor *string `json:"trust-anchor,omitempty"`
}

// DeviceSystemTlsServerprofileCipherlist struct
type DeviceSystemTlsServerprofileCipherlist struct {
	Cipherlist *string `json:"cipher-list,omitempty"`
}

// DeviceSystemTraceoptions struct
type DeviceSystemTraceoptions struct {
	// +kubebuilder:validation:Enum=`common`;`request`;`response`
	Traceoptions E_DeviceSystemTraceoptionsTraceoptions `json:"trace-options,omitempty"`
	//Traceoptions *string `json:"trace-options,omitempty"`
}

// DeviceSystemWarmreboot struct
type DeviceSystemWarmreboot struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=3600
	// +kubebuilder:default:=600
	Bgpmaxwait *uint16 `json:"bgp-max-wait,omitempty"`
}

// DeviceTunnel struct
type DeviceTunnel struct {
	//RootTunnelVxlantunnel
	Vxlantunnel *DeviceTunnelVxlantunnel `json:"vxlan-tunnel,omitempty"`
}

// DeviceTunnelVxlantunnel struct
type DeviceTunnelVxlantunnel struct {
}

// DeviceTunnelinterface struct
type DeviceTunnelinterface struct {
	// kubebuilder:validation:MinLength=6
	// kubebuilder:validation:MaxLength=8
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(vxlan(0|1[0-9][0-9]|2([0-4][0-9]|5[0-5])|[1-9][0-9]|[1-9]))`
	Name *string `json:"name"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=16384
	Vxlaninterface []*DeviceTunnelinterfaceVxlaninterface `json:"vxlan-interface,omitempty"`
}

// DeviceTunnelinterfaceVxlaninterface struct
type DeviceTunnelinterfaceVxlaninterface struct {
	//RootTunnelinterfaceVxlaninterfaceBridgetable
	Bridgetable *DeviceTunnelinterfaceVxlaninterfaceBridgetable `json:"bridge-table,omitempty"`
	//RootTunnelinterfaceVxlaninterfaceEgress
	Egress *DeviceTunnelinterfaceVxlaninterfaceEgress `json:"egress,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=99999999
	Index *uint32 `json:"index"`
	//RootTunnelinterfaceVxlaninterfaceIngress
	Ingress *DeviceTunnelinterfaceVxlaninterfaceIngress `json:"ingress,omitempty"`
	Type    *string                                     `json:"type"`
}

// DeviceTunnelinterfaceVxlaninterfaceBridgetable struct
type DeviceTunnelinterfaceVxlaninterfaceBridgetable struct {
}

// DeviceTunnelinterfaceVxlaninterfaceEgress struct
type DeviceTunnelinterfaceVxlaninterfaceEgress struct {
	//RootTunnelinterfaceVxlaninterfaceEgressDestinationgroups
	Destinationgroups *DeviceTunnelinterfaceVxlaninterfaceEgressDestinationgroups `json:"destination-groups,omitempty"`
	//RootTunnelinterfaceVxlaninterfaceEgressInnerethernetheader
	Innerethernetheader *DeviceTunnelinterfaceVxlaninterfaceEgressInnerethernetheader `json:"inner-ethernet-header,omitempty"`
	// +kubebuilder:default:="use-system-ipv4-address"
	Sourceip *string `json:"source-ip,omitempty"`
}

// DeviceTunnelinterfaceVxlaninterfaceEgressDestinationgroups struct
type DeviceTunnelinterfaceVxlaninterfaceEgressDestinationgroups struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Group []*DeviceTunnelinterfaceVxlaninterfaceEgressDestinationgroupsGroup `json:"group,omitempty"`
}

// DeviceTunnelinterfaceVxlaninterfaceEgressDestinationgroupsGroup struct
type DeviceTunnelinterfaceVxlaninterfaceEgressDestinationgroupsGroup struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_DeviceTunnelinterfaceVxlaninterfaceEgressDestinationgroupsGroupAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=128
	Destination []*DeviceTunnelinterfaceVxlaninterfaceEgressDestinationgroupsGroupDestination `json:"destination,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){9}`
	Esi *string `json:"esi,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9!@#$^&()|+=`~.,'/_:;?-][A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
}

// DeviceTunnelinterfaceVxlaninterfaceEgressDestinationgroupsGroupDestination struct
type DeviceTunnelinterfaceVxlaninterfaceEgressDestinationgroupsGroupDestination struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_DeviceTunnelinterfaceVxlaninterfaceEgressDestinationgroupsGroupDestinationAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	Index *uint16 `json:"index"`
	//RootTunnelinterfaceVxlaninterfaceEgressDestinationgroupsGroupDestinationInnerethernetheader
	Innerethernetheader *DeviceTunnelinterfaceVxlaninterfaceEgressDestinationgroupsGroupDestinationInnerethernetheader `json:"inner-ethernet-header,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=16777215
	Vni *uint32 `json:"vni,omitempty"`
}

// DeviceTunnelinterfaceVxlaninterfaceEgressDestinationgroupsGroupDestinationInnerethernetheader struct
type DeviceTunnelinterfaceVxlaninterfaceEgressDestinationgroupsGroupDestinationInnerethernetheader struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}`
	Destinationmac *string `json:"destination-mac,omitempty"`
}

// DeviceTunnelinterfaceVxlaninterfaceEgressInnerethernetheader struct
type DeviceTunnelinterfaceVxlaninterfaceEgressInnerethernetheader struct {
	// +kubebuilder:default:="use-system-mac"
	Sourcemac *string `json:"source-mac,omitempty"`
}

// DeviceTunnelinterfaceVxlaninterfaceIngress struct
type DeviceTunnelinterfaceVxlaninterfaceIngress struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=16777215
	Vni *uint32 `json:"vni"`
}

// A DeviceSpec defines the desired state of a Device.
type DeviceSpec struct {
	nddv1.ResourceSpec `json:",inline"`
	Device             *Device `json:"device,omitempty"`
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
// +kubebuilder:printcolumn:name="LEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='LeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="PARENTDEP",type="string",JSONPath=".status.conditions[?(@.kind=='ParentValidationSuccess')].status"
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
