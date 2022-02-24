package main

import (
	"fmt"
	"net"
	"os"
	"reflect"
	"strings"

	"github.com/openconfig/ygot/ygot"
	"github.com/yndd/nddp-srl3/internal/gnmi"
	"github.com/yndd/nddp-srl3/pkg/ygotsrl"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/openconfig/gnmi/proto/gnmi"
)

func callback(newConfig ygot.ValidatedGoStruct) error { // Apply the config to your device and return nil if success. return error if fails.		/
	// Do something ...
	fmt.Println(newConfig)
	return nil
}

const (
	yangImportDir = "/Users/henderiw/CodeProjects/yang/srlinux/21_11_2/ietf"
	yangModuleDir = "/Users/henderiw/CodeProjects/yang/srlinux/21_11_2/srl"
)

func main() {

	_, modules, err := parseYangModules([]string{yangImportDir}, []string{yangModuleDir})
	if err != nil {
		fmt.Println(err)
	}

	modelData := make([]*pb.ModelData, 0)
	for moduleName, module := range modules {
		//fmt.Printf("Module: %s, Moduledata: %v %v\n", moduleName, module.Organization, module.YangVersion)

		moduleRevision := ""
		if module.Revision != nil {
			moduleRevision = module.Revision[0].Name
		}
		moduleOrganization := "Nokia"
		if module.Organization != nil {
			moduleOrganization = module.Organization.Name
		}

		modelData = append(modelData, &pb.ModelData{
			Name:         moduleName,
			Organization: moduleOrganization,
			Version:      moduleRevision,
		})

	}

	model := &gnmi.Model{
		ModelData:       modelData,
		StructRootType:  reflect.TypeOf((*ygotsrl.Device)(nil)),
		SchemaTreeRoot:  ygotsrl.SchemaTree["Device"],
		JsonUnmarshaler: ygotsrl.Unmarshal,
		EnumData:        ygotsrl.ΛEnum,
	}

	c, err := os.ReadFile("/Users/henderiw/CodeProjects/yndd/nddp-srl3/cmd/test/srl.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	config := string(c)
	//config := string()

	g := grpc.NewServer()
	s, err := gnmi.NewServer(model, []byte(config), callback)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	pb.RegisterGNMIServer(g, s)
	reflection.Register(g)
	listen, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	g.Serve(listen)

	/*
		d, err := createDevice()
		if err != nil {
			fmt.Println(err)
		}
		j, err := ygot.EmitJSON(d, &ygot.EmitJSONConfig{
			Format: ygot.Internal,
			Indent: "  ",
			RFC7951Config: &ygot.RFC7951JSONConfig{
				AppendModuleName: true,
			},
		})
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(j)

		var x interface{}
		json.Unmarshal([]byte(j), &x)
		fmt.Println(x)
	*/

}

type itfceInfo struct {
	index uint32
	ipv4  []string
	ipv6  []string
}

func createDevice() (*ygotsrl.Device, error) {
	var err error
	d := &ygotsrl.Device{
		RoutingPolicy: &ygotsrl.SrlNokiaRoutingPolicy_RoutingPolicy{},
	}

	_, err = createInterface(d, "irb0")
	if err != nil {
		return nil, err
	}

	systemItfceInfo := &itfceInfo{
		index: 0,
		ipv4:  []string{"100.64.0.1/32"},
		ipv6:  []string{"1000:64::1/128"},
	}

	itfceSystem, err := createInterface(d, "system0")
	if err != nil {
		return nil, err
	}
	_, err = createSubInterface(itfceSystem, systemItfceInfo)
	if err != nil {
		return nil, err
	}

	ni, err := createNetworkInstance(d, "default")
	if err != nil {
		return nil, err
	}
	ni.NewInterface(strings.Join([]string{"system0", "0"}, "."))

	ni.Protocols, err = createBgpProtocol(ni, 65000)
	if err != nil {
		return nil, err
	}

	_, err = d.NewTunnelInterface("vxlan0")
	if err != nil {
		return nil, err
	}

	d.RoutingPolicy, err = createRoutingPolicy()
	if err != nil {
		return nil, err
	}
	return d, nil
}

func createBgpProtocol(ni *ygotsrl.SrlNokiaNetworkInstance_NetworkInstance, as uint32) (*ygotsrl.SrlNokiaNetworkInstance_NetworkInstance_Protocols, error) {
	p := &ygotsrl.SrlNokiaNetworkInstance_NetworkInstance_Protocols{
		Bgp: &ygotsrl.SrlNokiaNetworkInstance_NetworkInstance_Protocols_Bgp{
			AdminState:       ygotsrl.SrlNokiaCommon_AdminState_enable,
			AutonomousSystem: ygot.Uint32(as),
			RouterId:         ygot.String("100.64.0.1"),
			EbgpDefaultPolicy: &ygotsrl.SrlNokiaNetworkInstance_NetworkInstance_Protocols_Bgp_EbgpDefaultPolicy{
				ExportRejectAll: ygot.Bool(false),
				ImportRejectAll: ygot.Bool(false),
			},
			Ipv4Unicast: &ygotsrl.SrlNokiaNetworkInstance_NetworkInstance_Protocols_Bgp_Ipv4Unicast{
				AdminState: ygotsrl.SrlNokiaCommon_AdminState_enable,
				Multipath: &ygotsrl.SrlNokiaNetworkInstance_NetworkInstance_Protocols_Bgp_Ipv4Unicast_Multipath{
					AllowMultipleAs: ygot.Bool(true),
					MaxPathsLevel_1: ygot.Uint32(64),
					MaxPathsLevel_2: ygot.Uint32(64),
				},
			},
			Ipv6Unicast: &ygotsrl.SrlNokiaNetworkInstance_NetworkInstance_Protocols_Bgp_Ipv6Unicast{
				AdminState: ygotsrl.SrlNokiaCommon_AdminState_enable,
				Multipath: &ygotsrl.SrlNokiaNetworkInstance_NetworkInstance_Protocols_Bgp_Ipv6Unicast_Multipath{
					AllowMultipleAs: ygot.Bool(true),
					MaxPathsLevel_1: ygot.Uint32(64),
					MaxPathsLevel_2: ygot.Uint32(64),
				},
			},
			Evpn: &ygotsrl.SrlNokiaNetworkInstance_NetworkInstance_Protocols_Bgp_Evpn{
				AdminState: ygotsrl.SrlNokiaCommon_AdminState_enable,
			},
		},
	}
	underlayGroup, err := p.Bgp.NewGroup("underlay")
	if err != nil {
		return nil, err
	}
	underlayGroup.AdminState = ygotsrl.SrlNokiaCommon_AdminState_enable
	underlayGroup.NextHopSelf = ygot.Bool(true)
	underlayGroup.ExportPolicy = ygot.String("export-local")
	underlayGroup.Ipv4Unicast = &ygotsrl.SrlNokiaNetworkInstance_NetworkInstance_Protocols_Bgp_Group_Ipv4Unicast{
		AdminState: ygotsrl.SrlNokiaCommon_AdminState_enable,
	}
	underlayGroup.Ipv6Unicast = &ygotsrl.SrlNokiaNetworkInstance_NetworkInstance_Protocols_Bgp_Group_Ipv6Unicast{
		AdminState: ygotsrl.SrlNokiaCommon_AdminState_enable,
	}
	underlayGroup.Evpn = &ygotsrl.SrlNokiaNetworkInstance_NetworkInstance_Protocols_Bgp_Group_Evpn{
		AdminState: ygotsrl.SrlNokiaCommon_AdminState_enable,
	}

	overlayGroup, err := p.Bgp.NewGroup("overlay")
	if err != nil {
		return nil, err
	}
	overlayGroup.AdminState = ygotsrl.SrlNokiaCommon_AdminState_enable
	overlayGroup.Evpn = &ygotsrl.SrlNokiaNetworkInstance_NetworkInstance_Protocols_Bgp_Group_Evpn{
		AdminState: ygotsrl.SrlNokiaCommon_AdminState_enable,
	}
	return p, nil
}

func createNetworkInstance(d *ygotsrl.Device, niName string) (*ygotsrl.SrlNokiaNetworkInstance_NetworkInstance, error) {
	ni, err := d.NewNetworkInstance(niName)
	if err != nil {
		return nil, err
	}
	ni.Type = ygotsrl.SrlNokiaNetworkInstance_NiType_default
	ni.AdminState = ygotsrl.SrlNokiaCommon_AdminState_enable

	return ni, nil
}

func createSubInterface(i *ygotsrl.SrlNokiaInterfaces_Interface, itfceInfo *itfceInfo) (*ygotsrl.SrlNokiaInterfaces_Interface_Subinterface, error) {
	si, err := i.NewSubinterface(itfceInfo.index)
	if err != nil {
		return nil, err
	}
	si.AdminState = ygotsrl.SrlNokiaCommon_AdminState_enable
	si.Ipv4 = &ygotsrl.SrlNokiaInterfaces_Interface_Subinterface_Ipv4{}
	for _, a := range itfceInfo.ipv4 {
		_, err := si.Ipv4.NewAddress(a)
		if err != nil {
			return nil, err
		}
	}
	si.Ipv6 = &ygotsrl.SrlNokiaInterfaces_Interface_Subinterface_Ipv6{}
	for _, a := range itfceInfo.ipv6 {
		_, err := si.Ipv6.NewAddress(a)
		if err != nil {
			return nil, err
		}
	}
	return si, nil
}

func createInterface(d *ygotsrl.Device, itfceName string) (*ygotsrl.SrlNokiaInterfaces_Interface, error) {
	i, err := d.NewInterface(itfceName)
	if err != nil {
		return nil, err
	}

	i.AdminState = ygotsrl.SrlNokiaCommon_AdminState_enable

	return i, nil
}

func createRoutingPolicy() (*ygotsrl.SrlNokiaRoutingPolicy_RoutingPolicy, error) {
	rp := &ygotsrl.SrlNokiaRoutingPolicy_RoutingPolicy{}

	prefixSetIpv4, err := rp.NewPrefixSet("local-ipv42")
	if err != nil {
		return nil, err
	}
	_, err = prefixSetIpv4.NewPrefix("100.64.0.0/16", "32..32")
	if err != nil {
		return nil, err
	}

	prefixSetIpv6, err := rp.NewPrefixSet("local-ipv6")
	if err != nil {
		return nil, err
	}
	_, err = prefixSetIpv6.NewPrefix("1000:64::/48", "128..128")
	if err != nil {
		return nil, err
	}

	policy, err := rp.NewPolicy("export-local")
	if err != nil {
		return nil, err
	}

	statement10, err := policy.NewStatement(10)
	if err != nil {
		return nil, err
	}
	statement10.Match = &ygotsrl.SrlNokiaRoutingPolicy_RoutingPolicy_Policy_Statement_Match{
		PrefixSet: ygot.String("local-ipv4"),
	}
	statement10.Action = &ygotsrl.SrlNokiaRoutingPolicy_RoutingPolicy_Policy_Statement_Action{
		Accept: &ygotsrl.SrlNokiaRoutingPolicy_RoutingPolicy_Policy_Statement_Action_Accept{},
	}

	statement20, err := policy.NewStatement(20)
	if err != nil {
		return nil, err
	}
	statement20.Match = &ygotsrl.SrlNokiaRoutingPolicy_RoutingPolicy_Policy_Statement_Match{
		PrefixSet: ygot.String("local-ipv6"),
	}
	statement20.Action = &ygotsrl.SrlNokiaRoutingPolicy_RoutingPolicy_Policy_Statement_Action{
		Accept: &ygotsrl.SrlNokiaRoutingPolicy_RoutingPolicy_Policy_Statement_Action_Accept{},
	}
	return rp, nil
}
