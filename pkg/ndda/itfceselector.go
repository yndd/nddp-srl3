package ndda

import (
	"fmt"
	"strings"

	networkv1alpha1 "github.com/yndd/ndda-network/apis/network/v1alpha1"
	"github.com/yndd/ndda-network/pkg/ndda/itfceinfo"
	nddov1 "github.com/yndd/nddo-runtime/apis/common/v1"
	"github.com/yndd/nddo-runtime/pkg/odns"
	"github.com/yndd/nddo-runtime/pkg/resource"
	srlv1alpha1 "github.com/yndd/nddp-srl3/apis/srl3/v1alpha1"
)

func (r *handler) GetSelectedNodeItfces(mg resource.Managed, epgSelectors []*nddov1.EpgInfo, nodeItfceSelectors map[string]*nddov1.ItfceInfo) (map[string][]itfceinfo.ItfceInfo, error) {
	// get all ndda interfaces within the oda scope
	// oda is organization, deployement, availability zone
	opts := odns.GetClientListOptionFromResourceName(mg.GetName())
	fmt.Printf("opts: %v\n", opts)
	nddaDevices := r.newSrlDeviceList()
	if err := r.client.List(r.ctx, nddaDevices); err != nil {
		//if err := r.client.List(r.ctx, nddaItfces, opts...); err != nil {
		return nil, err
	}

	sel := NewNodeItfceSelection()
	sel.GetNodeItfcesByEpgSelector(epgSelectors, nddaDevices)
	sel.GetNodeItfcesByNodeItfceSelector(nodeItfceSelectors, nddaDevices)
	return sel.GetSelectedNodeItfces(), nil

}

/*
func (r *handler) GetSelectedNodeItfcesIrb(mg resource.Managed, s srlschema.Schema, niName string) (map[string][]itfceinfo.ItfceInfo, error) {
	// get all ndda interfaces within the oda scope
	// oda is organization, deployement, availability zone
	opts := odns.GetClientListOptionFromResourceName(mg.GetName())
	nddaItfces := r.newItfceList()
	if err := r.client.List(r.ctx, nddaItfces, opts...); err != nil {
		return nil, err
	}

	sel := NewNodeItfceSelection()
	sel.GetIrbNodeItfces(niName, s, nddaItfces)
	return sel.GetSelectedNodeItfces(), nil
}
*/

/*
func (r *handler) GetSelectedNodeItfcesVxlan(mg resource.Managed, s srlschema.Schema, niName string) (map[string][]itfceinfo.ItfceInfo, error) {
	// get all ndda interfaces within the oda scope
	// oda is organization, deployement, availability zone
	opts := odns.GetClientListOptionFromResourceName(mg.GetName())
	nddaItfces := r.newItfceList()
	if err := r.client.List(r.ctx, nddaItfces, opts...); err != nil {
		return nil, err
	}

	sel := NewNodeItfceSelection()
	sel.GetVxlanNodeItfces(niName, s, nddaItfces)
	return sel.GetSelectedNodeItfces(), nil
}
*/

type NodeItfceSelection interface {
	GetSelectedNodeItfces() map[string][]itfceinfo.ItfceInfo
	GetNodeItfcesByEpgSelector(epgSelectors []*nddov1.EpgInfo, nddaDeviceList srlv1alpha1.IFSrl3DeviceList)
	GetNodeItfcesByNodeItfceSelector(nodeItfceSelectors map[string]*nddov1.ItfceInfo, nddaDeviceList srlv1alpha1.IFSrl3DeviceList)
	//GetVxlanNodeItfces(string, srlschema.Schema, srlv1alpha1.IFSrlInterfaceList)
	//GetIrbNodeItfces(string, srlschema.Schema, srlv1alpha1.IFSrlInterfaceList)
}

func NewNodeItfceSelection() NodeItfceSelection {
	return &selectedNodeItfces{
		nodes: make(map[string][]itfceinfo.ItfceInfo),
	}
}

type selectedNodeItfces struct {
	nodes map[string][]itfceinfo.ItfceInfo
}

func (x *selectedNodeItfces) GetSelectedNodeItfces() map[string][]itfceinfo.ItfceInfo {
	return x.nodes
}

func getEndpointGroup(d srlv1alpha1.Srl3Device) string {
	if s, ok := d.Labels[srlv1alpha1.LabelNddaEndpointGroup]; !ok {
		return ""
	} else {
		return s
	}
}

func getDeviceName(d srlv1alpha1.Srl3Device) string {
	if s, ok := d.Labels[srlv1alpha1.LabelNddaDevice]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *selectedNodeItfces) GetNodeItfcesByEpgSelector(epgSelectors []*nddov1.EpgInfo, nddaDeviceList srlv1alpha1.IFSrl3DeviceList) {
	for _, d := range nddaDeviceList.GetDevices() {
		if d.GetSpec().Device.Interface != nil {
			for _, i := range d.GetSpec().Device.Interface {
				fmt.Printf("getNodeItfcesByEpgSelector: itfceepg: %s, nodename: %s, itfcename: %s, lagmember: %v\n", d.GetEndpointGroup(), d.GetDeviceName(), *i.Name, *i.Lag)
				for _, epgSelector := range epgSelectors {
					if epgSelector.EpgName != "" && epgSelector.EpgName == d.GetEndpointGroup() {
						fmt.Printf("getNodeItfcesByEpgSelector: %s\n", d.GetName())
						// avoid selecting lag members
						if !(i.Ethernet != nil && i.Ethernet.Aggregateid != nil) {
							x.addNodeItfce(d.GetDeviceName(), *i.Name, itfceinfo.NewItfceInfo(
								itfceinfo.WithInnerVlanId(epgSelector.InnerVlanId),
								itfceinfo.WithOuterVlanId(epgSelector.OuterVlanId),
								itfceinfo.WithItfceKind(networkv1alpha1.E_InterfaceKind_INTERFACE),
								itfceinfo.WithIpv4Prefixes(epgSelector.Ipv4Prefixes),
								itfceinfo.WithIpv6Prefixes(epgSelector.Ipv6Prefixes),
							))
						}
					}
				}
			}
		}
	}
}

func (x *selectedNodeItfces) GetNodeItfcesByNodeItfceSelector(nodeItfceSelectors map[string]*nddov1.ItfceInfo, nddaDeviceList srlv1alpha1.IFSrl3DeviceList) {
	for _, d := range nddaDeviceList.GetDevices() {
		if d.GetSpec().Device.Interface != nil {
			for _, i := range d.GetSpec().Device.Interface {
				for deviceName, itfceInfo := range nodeItfceSelectors {
					fmt.Printf("getNodeItfcesByNodeItfceSelector: nodename: %s, itfcename: %s, lagmember: %v, nodename: %s\n", d.GetDeviceName(), *i.Name, *i.Lag, deviceName)

					var itfceName string
					if strings.Contains(itfceInfo.ItfceName, "lag") {
						itfceName = strings.ReplaceAll(itfceInfo.ItfceName, "-", "")
					}
					if strings.Contains(itfceInfo.ItfceName, "int") {
						itfceName = strings.ReplaceAll(itfceInfo.ItfceName, "int", "ethernet")
						split := strings.Split(itfceName, "/")
						if len(split) > 2 {
							itfceName = "ethernet-" + split[len(split)-2] + "/" + split[len(split)-1]
						}
					}

					// avoid selecting lag members
					if !(i.Ethernet != nil && i.Ethernet.Aggregateid != nil) {
						if deviceName == d.GetDeviceName() &&
							itfceName == *i.Name {
							fmt.Printf("getNodeItfcesByNodeItfceSelector selected: nodename: %s, itfcename: %s, lagmember: %v, nodename: %s\n", d.GetDeviceName(), *i.Name, *i.Lag, deviceName)
							x.addNodeItfce(d.GetDeviceName(), *i.Name, itfceinfo.NewItfceInfo(
								itfceinfo.WithInnerVlanId(itfceInfo.InnerVlanId),
								itfceinfo.WithOuterVlanId(itfceInfo.OuterVlanId),
								itfceinfo.WithItfceKind(networkv1alpha1.E_InterfaceKind_INTERFACE),
								itfceinfo.WithIpv4Prefixes(itfceInfo.Ipv4Prefixes),
								itfceinfo.WithIpv6Prefixes(itfceInfo.Ipv6Prefixes),
							))
						}
					}
				}
			}
		}
	}
}

/*
func (x *selectedNodeItfces) GetVxlanNodeItfces(niName string, s srlschema.Schema, nddaItfceList srlv1alpha1.IFSrlInterfaceList) {
	for _, nddaItfce := range nddaItfceList.GetInterfaces() {
		for deviceName, d := range s.GetDevices() {
			for dniName := range d.GetNetworkinstances() {
				if dniName == niName {
					if nddaItfce.GetDeviceName() == deviceName && strings.Contains(nddaItfce.GetInterfaceName()  {
						x.addNodeItfce(deviceName, nddaItfce.GetInterfaceName(), itfceinfo.NewItfceInfo(
							itfceinfo.WithItfceKind(networkv1alpha1.E_InterfaceKind_VXLAN),
							//WithItfceIndex(ni.GetIndex()), // we use the vxlan
							//WithIpv4Prefixes(make([]*string, 0)),
							//WithIpv6Prefixes(make([]*string, 0)),
						))
					}
				}
			}
		}
	}
}

func (x *selectedNodeItfces) GetIrbNodeItfces(niName string, s networkschema.Schema, nddaItfceList networkv1alpha1.IFNetworkInterfaceList) {
	for _, nddaItfce := range nddaItfceList.GetInterfaces() {
		for deviceName, d := range s.GetDevices() {
			for dniName := range d.GetNetworkInstances() {
				if dniName == niName {
					// we only select the irb interfaces to retain the index
					if nddaItfce.GetDeviceName() == deviceName && nddaItfce.GetInterfaceConfigKind() == networkv1alpha1.E_InterfaceKind_IRB {
						x.addNodeItfce(deviceName, nddaItfce.GetInterfaceName(), itfceinfo.NewItfceInfo(
							itfceinfo.WithItfceKind(networkv1alpha1.E_InterfaceKind_IRB),
							//WithItfceIndex(9999), // dummy
							//WithIpv4Prefixes(ipv4Prefixes),
							//WithIpv6Prefixes(ipv6Prefixes),
						))
					}
				}
			}
		}
	}
}
*/

func (x *selectedNodeItfces) addNodeItfce(nodeName, intName string, ifInfo itfceinfo.ItfceInfo) {
	// check if node exists, if not initialize the node
	if _, ok := x.nodes[nodeName]; !ok {
		x.nodes[nodeName] = make([]itfceinfo.ItfceInfo, 0)
	}
	// check if the interfacename was already present on the node
	// if not add it to the list
	for _, itfceInfo := range x.nodes[nodeName] {
		if itfceInfo.GetItfceName() == intName {
			return
		}
	}
	ifInfo.SetItfceName(intName)
	x.nodes[nodeName] = append(x.nodes[nodeName], ifInfo)
}
