package srl

import (
	"context"
	"encoding/json"
	"reflect"
	"testing"

	"github.com/go-logr/logr"
	"github.com/openconfig/ygot/ygot"
	"github.com/yndd/ndd-runtime/pkg/logging"
	srlv1alpha1 "github.com/yndd/nddp-srl3/apis/srl3/v1alpha1"
	"github.com/yndd/nddp-srl3/internal/model"
	"github.com/yndd/nddp-srl3/pkg/ygotsrl"
	"k8s.io/apimachinery/pkg/runtime"
)

func Test_ValidateCrSpecDelete(t *testing.T) {
	v := &validatorDevice{
		log: logging.NewLogrLogger(logr.Discard()),
		deviceModel: &model.Model{
			StructRootType:  reflect.TypeOf((*ygotsrl.Device)(nil)),
			SchemaTreeRoot:  ygotsrl.SchemaTree["Device"],
			JsonUnmarshaler: ygotsrl.Unmarshal,
			EnumData:        ygotsrl.ΛEnum,
		},
	}

	runningGoStruct := getRunning(t)
	runningConfig, err := getByteArrFromYgot(runningGoStruct)
	if err != nil {
		t.Errorf("Check Error: %s", err)
	}

	specGoStruct := getSpecGoStruct(t)
	mg, err := getSrlCrd(specGoStruct)
	if err != nil {
		t.Errorf("Check Error: %s", err)
	}

	foo, err := v.ValidateCrSpecDelete(context.Background(), mg, runningConfig)
	_ = foo
	if err != nil {
		t.Error(err)
	}
}

func getSpecGoStruct(t *testing.T) *ygotsrl.Device {
	specGoStruct := &ygotsrl.Device{}
	se14 := specGoStruct.GetOrCreateInterface("ethernet-1/4")
	se14s5 := se14.GetOrCreateSubinterface(5)
	se14s5.GetOrCreateIpv4().GetOrCreateAddress("192.168.5.2/32")

	//specGoStruct.GetOrCreateNetworkInstance("special")
	soni := specGoStruct.GetOrCreateNetworkInstance("other")
	soni.GetOrCreateInterface("ethernet-1/5.7")

	return specGoStruct
}

func getRunning(t *testing.T) *ygotsrl.Device {
	runningGoStruct := &ygotsrl.Device{}
	e14 := runningGoStruct.GetOrCreateInterface("ethernet-1/4")
	e14s5 := e14.GetOrCreateSubinterface(5)
	e14s5.GetOrCreateIpv4().GetOrCreateAddress("192.168.5.1/32")
	e14s5.GetOrCreateIpv4().GetOrCreateAddress("192.168.5.2/32")

	e15 := runningGoStruct.GetOrCreateInterface("ethernet-1/5")
	e15s5 := e15.GetOrCreateSubinterface(5)
	e15s5.GetOrCreateIpv4().GetOrCreateAddress("192.168.5.5/32")
	e15s7 := e15.GetOrCreateSubinterface(7)
	e15s7.GetOrCreateIpv4().GetOrCreateAddress("192.168.7.1/32")

	sni := runningGoStruct.GetOrCreateNetworkInstance("special")
	sni.AdminState = ygotsrl.SrlNokiaCommon_AdminState_enable
	sni.GetOrCreateInterface("ethernet-1/4.5")

	oni := runningGoStruct.GetOrCreateNetworkInstance("other")
	oni.AdminState = ygotsrl.SrlNokiaCommon_AdminState_enable
	oni.GetOrCreateInterface("ethernet-1/5.7")
	oni.GetOrCreateInterface("ethernet-1/4.7")

	return runningGoStruct
}

func getByteArrFromYgot(gostruct ygot.GoStruct) ([]byte, error) {

	jsonTree, err := ygot.ConstructIETFJSON(gostruct, &ygot.RFC7951JSONConfig{})
	if err != nil {
		return nil, err
	}

	byte_config, err := json.Marshal(jsonTree)
	if err != nil {
		return nil, err
	}

	return byte_config, nil
}

func getSrlCrd(gostruct ygot.GoStruct) (*srlv1alpha1.Srl3Device, error) {
	byte_config, err := getByteArrFromYgot(gostruct)
	if err != nil {
		return nil, err
	}

	srlDevCr := &srlv1alpha1.Srl3Device{
		Spec: srlv1alpha1.DeviceSpec{
			Properties: runtime.RawExtension{
				Raw: byte_config,
			},
		},
		Status: srlv1alpha1.DeviceStatus{},
	}

	return srlDevCr, nil
}
