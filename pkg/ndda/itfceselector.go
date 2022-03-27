package ndda

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/openconfig/gnmi/proto/gnmi"
	networkv1alpha1 "github.com/yndd/ndda-network/apis/network/v1alpha1"
	"github.com/yndd/ndda-network/pkg/ndda/itfceinfo"
	nddov1 "github.com/yndd/nddo-runtime/apis/common/v1"
	"github.com/yndd/nddo-runtime/pkg/odns"
	"github.com/yndd/nddo-runtime/pkg/resource"
	srlv1alpha1 "github.com/yndd/nddp-srl3/apis/srl3/v1alpha1"
	"github.com/yndd/nddp-srl3/internal/model"
	"github.com/yndd/nddp-srl3/pkg/ygotsrl"
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

type NodeItfceSelection interface {
	GetSelectedNodeItfces() map[string][]itfceinfo.ItfceInfo
	GetNodeItfcesByEpgSelector(epgSelectors []*nddov1.EpgInfo, nddaDeviceList srlv1alpha1.IFSrl3DeviceList) error
	GetNodeItfcesByNodeItfceSelector(nodeItfceSelectors map[string]*nddov1.ItfceInfo, nddaDeviceList srlv1alpha1.IFSrl3DeviceList) error
	//GetVxlanNodeItfces(string, srlschema.Schema, srlv1alpha1.IFSrlInterfaceList)
	//GetIrbNodeItfces(string, srlschema.Schema, srlv1alpha1.IFSrlInterfaceList)
}

func NewNodeItfceSelection() NodeItfceSelection {
	return &selectedNodeItfces{
		m: &model.Model{
			ModelData:       make([]*gnmi.ModelData, 0),
			StructRootType:  reflect.TypeOf((*ygotsrl.Device)(nil)),
			SchemaTreeRoot:  ygotsrl.SchemaTree["Device"],
			JsonUnmarshaler: ygotsrl.Unmarshal,
			EnumData:        ygotsrl.ΛEnum,
		},
		nodes: make(map[string][]itfceinfo.ItfceInfo),
	}
}

type selectedNodeItfces struct {
	m     *model.Model
	nodes map[string][]itfceinfo.ItfceInfo
}

func (x *selectedNodeItfces) GetSelectedNodeItfces() map[string][]itfceinfo.ItfceInfo {
	return x.nodes
}

func (x *selectedNodeItfces) GetNodeItfcesByEpgSelector(epgSelectors []*nddov1.EpgInfo, nddaDeviceList srlv1alpha1.IFSrl3DeviceList) error {
	for _, d := range nddaDeviceList.GetDevices() {
		deviceConfig, err := x.getDeviceConfig(d.GetSpec().Properties.Raw)
		if err != nil {
			return err
		}
		for _, i := range deviceConfig.Interface {
			fmt.Printf("getNodeItfcesByEpgSelector: itfceepg: %s, nodename: %s, itfcename: %s\n", d.GetEndpointGroup(), d.GetDeviceName(), *i.Name)
			for _, epgSelector := range epgSelectors {
				if epgSelector.EpgName != "" && epgSelector.EpgName == d.GetEndpointGroup() {
					fmt.Printf("getNodeItfcesByEpgSelector: %s\n", d.GetName())
					// avoid selecting lag members
					if !(i.Ethernet != nil && i.Ethernet.AggregateId != nil) {
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
		//TO BE UPDATED
		/*
			if d.GetSpec().Device.Interface != nil {
				for _, i := range d.GetSpec().Properties {
					fmt.Printf("getNodeItfcesByEpgSelector: itfceepg: %s, nodename: %s, itfcename: %s\n", d.GetEndpointGroup(), d.GetDeviceName(), *i.Name)
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
		*/
		//fmt.Printf("d:%v\n", d)
	}
	return nil
}

func (x *selectedNodeItfces) GetNodeItfcesByNodeItfceSelector(nodeItfceSelectors map[string]*nddov1.ItfceInfo, nddaDeviceList srlv1alpha1.IFSrl3DeviceList) error {
	for _, d := range nddaDeviceList.GetDevices() {
		deviceConfig, err := x.getDeviceConfig(d.GetSpec().Properties.Raw)
		if err != nil {
			return err
		}
		for _, i := range deviceConfig.Interface {
			for deviceName, itfceInfo := range nodeItfceSelectors {
				fmt.Printf("getNodeItfcesByNodeItfceSelector: nodename: %s, itfcename: %s, nodename: %s\n", d.GetDeviceName(), *i.Name, deviceName)

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
				if !(i.Ethernet != nil && i.Ethernet.AggregateId != nil) {
					if deviceName == d.GetDeviceName() &&
						itfceName == *i.Name {
						fmt.Printf("getNodeItfcesByNodeItfceSelector selected: nodename: %s, itfcename: %s, nodename: %s\n", d.GetDeviceName(), *i.Name, deviceName)
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
		//TO BE UPDATED
		/*
			if d.GetSpec().Device.Interface != nil {
				for _, i := range d.GetSpec().Device.Interface {
					for deviceName, itfceInfo := range nodeItfceSelectors {
						fmt.Printf("getNodeItfcesByNodeItfceSelector: nodename: %s, itfcename: %s, nodename: %s\n", d.GetDeviceName(), *i.Name, deviceName)

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
								fmt.Printf("getNodeItfcesByNodeItfceSelector selected: nodename: %s, itfcename: %s, nodename: %s\n", d.GetDeviceName(), *i.Name, deviceName)
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
		*/
		//fmt.Printf("d:%v\n", d)
	}
	return nil
}

func (x *selectedNodeItfces) getDeviceConfig(config []byte) (*ygotsrl.Device, error) {
	deviceStruct, err := x.m.NewConfigStruct(config, false)
	if err != nil {
		return nil, err
	}
	deviceConfig, ok := deviceStruct.(*ygotsrl.Device)
	if !ok {
		return nil, fmt.Errorf("wrong device config: %s", string(config))
	}
	return deviceConfig, nil
}

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
