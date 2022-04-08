package abstract

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/yndd/ndda-network/pkg/abstraction"
	"github.com/yndd/ndda-network/pkg/ndda/itfceinfo"
	"github.com/yndd/ndda-network/pkg/ygotndda"
	nddov1 "github.com/yndd/nddo-runtime/apis/common/v1"
	"github.com/yndd/nddo-runtime/pkg/odns"
	"github.com/yndd/nddo-runtime/pkg/resource"
	srlv1alpha1 "github.com/yndd/nddp-srl3/apis/srl3/v1alpha1"
	"github.com/yndd/nddp-srl3/internal/model"
	"github.com/yndd/nddp-srl3/pkg/ygotsrl"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func InitSrl(c resource.ClientApplicator, name, platform string) abstraction.Abstractor {
	return &srlabstract{
		client:   c,
		name:     name,
		platform: platform,
		// srl device specific model information
		newSrlDeviceList: func() srlv1alpha1.IFSrl3DeviceList { return &srlv1alpha1.Srl3DeviceList{} },
		newSrlDevice:     func() srlv1alpha1.IFSrl3Device { return &srlv1alpha1.Srl3Device{} },
		m: &model.Model{
			ModelData:       make([]*gnmi.ModelData, 0),
			StructRootType:  reflect.TypeOf((*ygotsrl.Device)(nil)),
			SchemaTreeRoot:  ygotsrl.SchemaTree["Device"],
			JsonUnmarshaler: ygotsrl.Unmarshal,
			EnumData:        ygotsrl.ΛEnum,
		},
	}
}

type srlabstract struct {
	// k8s client
	client resource.ClientApplicator
	// name of the device
	name string
	// platform type
	platform string
	// generic interface for the device list
	newSrlDeviceList func() srlv1alpha1.IFSrl3DeviceList
	newSrlDevice     func() srlv1alpha1.IFSrl3Device
	// device Model
	m *model.Model
}

func (x *srlabstract) GetInterfaceName(itfceName string) (string, error) {
	// naming -> slot/mda/port/breakout ?
	if strings.HasPrefix(itfceName, "lag") {
		split := strings.Split(itfceName, "-")
		switch len(split) {
		case 2:
			nbr, err := strconv.Atoi(split[1])
			if err != nil {
				return "", err
			}
			switch x.platform {
			case "ixrd1", "ixrd2", "ixrd3", "ixrd2l", "ixrd3l":
				if nbr > 32 {
					return "", fmt.Errorf("wrong lag id cannot be bigger than 32, we got: %d", nbr)
				}
			case "ixrh1", "ixrh2":
				if nbr > 127 {
					return "", fmt.Errorf("wrong lag id cannot be bigger than 127, we got: %d", nbr)
				}
			}
		default:
			return "", fmt.Errorf("wrong lag naming, got: %s", itfceName)
		}
		itfceName = strings.Join([]string{split[0], split[1]}, "")

	}
	if strings.HasPrefix(itfceName, "int") {
		split := strings.Split(itfceName, "-")
		switch len(split) {
		case 2:
			split := strings.Split(split[1], "/")
			switch len(split) {
			case 2:
				spi := strings.Join([]string{split[0], split[1]}, "/")
				itfceName = strings.Join([]string{"ethernet", spi}, "-")
			case 3:
				spi := strings.Join([]string{split[1], split[2]}, "/")
				itfceName = strings.Join([]string{"ethernet", spi}, "-")
			default:
				return "", fmt.Errorf("wrong interface naming, got: %s", itfceName)
			}
		default:
			return "", fmt.Errorf("wrong interface naming, got: %s", itfceName)
		}
	}
	return itfceName, nil
}

func (x *srlabstract) GetSelectedItfces(ctx context.Context, mg resource.Managed, deviceName string, epgSelectors []*nddov1.EpgInfo, nodeItfceSelectors map[string]*nddov1.ItfceInfo) (map[string]itfceinfo.ItfceInfo, error) {
	// get all ndda interfaces within the oda scope
	// oda is organization, deployement, availability zone
	// and device name
	opts := odns.GetClientListOptionFromResourceName(mg.GetName())
	opts = append(opts, client.MatchingLabels{nddov1.LabelNddaDevice: deviceName})
	fmt.Printf("opts: %v\n", opts)
	nddaDeviceList := x.newSrlDeviceList()
	if err := x.client.List(ctx, nddaDeviceList, opts...); err != nil {
		return nil, err
	}

	selectedItfces := make(map[string]itfceinfo.ItfceInfo, 0)
	if err := x.getItfcesByEpgSelector(epgSelectors, nddaDeviceList, selectedItfces); err != nil {
		return nil, err
	}
	if err := x.getItfcesByItfceSelector(nodeItfceSelectors, nddaDeviceList, selectedItfces); err != nil {
		return nil, err
	}
	return selectedItfces, nil
}

func (x *srlabstract) getItfcesByEpgSelector(epgSelectors []*nddov1.EpgInfo, nddaDeviceList srlv1alpha1.IFSrl3DeviceList, selectedItfces map[string]itfceinfo.ItfceInfo) error {
	for _, d := range nddaDeviceList.GetDevices() {
		deviceConfig, err := x.getDeviceConfig(d.GetSpec().Properties.Raw)
		if err != nil {
			return err
		}
		for itfceName, i := range deviceConfig.Interface {
			fmt.Printf("getItfcesByEpgSelector: itfceepg: %s, nodename: %s, itfcename: %s\n", d.GetEndpointGroup(), d.GetDeviceName(), *i.Name)
			for _, epgSelector := range epgSelectors {
				if epgSelector.EpgName != "" && epgSelector.EpgName == d.GetEndpointGroup() {
					fmt.Printf("getItfcesByEpgSelector: %s\n", d.GetName())
					// avoid selecting lag members
					if !(i.Ethernet != nil && i.Ethernet.AggregateId != nil) {
						if _, ok := selectedItfces[itfceName]; !ok {
							selectedItfces[itfceName] = itfceinfo.NewItfceInfo(
								itfceinfo.WithItfceName(itfceName),
								itfceinfo.WithInnerVlanId(epgSelector.InnerVlanId),
								itfceinfo.WithOuterVlanId(epgSelector.OuterVlanId),
								itfceinfo.WithItfceKind(ygotndda.NddaCommon_InterfaceKind_INTERFACE),
								itfceinfo.WithIpv4Prefixes(epgSelector.Ipv4Prefixes),
								itfceinfo.WithIpv6Prefixes(epgSelector.Ipv6Prefixes),
							)
						}
					}
				}
			}
		}
	}
	return nil
}

func (x *srlabstract) getItfcesByItfceSelector(nodeItfceSelectors map[string]*nddov1.ItfceInfo, nddaDeviceList srlv1alpha1.IFSrl3DeviceList, selectedItfces map[string]itfceinfo.ItfceInfo) error {
	for _, d := range nddaDeviceList.GetDevices() {
		deviceConfig, err := x.getDeviceConfig(d.GetSpec().Properties.Raw)
		if err != nil {
			return err
		}
		for _, i := range deviceConfig.Interface {
			for deviceName, itfceInfo := range nodeItfceSelectors {
				fmt.Printf("getItfcesByItfceSelector: nodename: %s, itfcename: %s, nodename: %s\n", d.GetDeviceName(), *i.Name, deviceName)
				itfceName, err := x.GetInterfaceName(itfceInfo.ItfceName)
				if err != nil {
					return err
				}
				// avoid selecting lag members
				if !(i.Ethernet != nil && i.Ethernet.AggregateId != nil) {
					if deviceName == d.GetDeviceName() && itfceName == *i.Name {
						fmt.Printf("getItfcesByItfceSelector selected: nodename: %s, itfcename: %s, nodename: %s\n", d.GetDeviceName(), *i.Name, deviceName)
						if _, ok := selectedItfces[itfceName]; !ok {
							selectedItfces[itfceName] = itfceinfo.NewItfceInfo(
								itfceinfo.WithItfceName(itfceName),
								itfceinfo.WithInnerVlanId(itfceInfo.InnerVlanId),
								itfceinfo.WithOuterVlanId(itfceInfo.OuterVlanId),
								itfceinfo.WithItfceKind(ygotndda.NddaCommon_InterfaceKind_INTERFACE),
								itfceinfo.WithIpv4Prefixes(itfceInfo.Ipv4Prefixes),
								itfceinfo.WithIpv6Prefixes(itfceInfo.Ipv6Prefixes),
							)
						}
					}
				}
			}
		}
	}
	return nil
}

func (x *srlabstract) getDeviceConfig(config []byte) (*ygotsrl.Device, error) {
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
