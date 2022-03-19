/*
Copyright 2021 Wim Henderickx.

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

package srl

import (
	"context"

	gapi "github.com/karimra/gnmic/api"
	"github.com/karimra/gnmic/target"
	gutils "github.com/karimra/gnmic/utils"
	ndrv1 "github.com/yndd/ndd-core/apis/dvr/v1"

	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/openconfig/gnmi/proto/gnmi_ext"
	"github.com/pkg/errors"
	nddv1 "github.com/yndd/ndd-runtime/apis/common/v1"
	"github.com/yndd/ndd-runtime/pkg/logging"
	"github.com/yndd/ndd-yang/pkg/parser"
	"github.com/yndd/ndd-yang/pkg/yparser"
	"github.com/yndd/nddp-srl3/internal/device"
)

const (
	DeviceType    = "nokia-srl"
	State         = "STATE"
	Configuration = "CONFIG"
	encoding      = "JSON_IETF"
	//errors
	errGetGnmiCapabilities     = "gnmi capabilities error"
	errGnmiCreateGetRequest    = "gnmi create get request error"
	errGnmiGet                 = "gnmi get error "
	errGnmiHandleGetResponse   = "gnmi get response error"
	errGnmiCreateSetRequest    = "gnmi create set request error"
	errGnmiSet                 = "gnmi set error "
	errGnmiCreateDeleteRequest = "gnmi create delete request error"

	//
	swVersionPath = "/platform/control[slot=A]/software-version"
	chassisPath   = "/platform/chassis"
)

func init() {
	device.Register(DeviceType, func() device.Device {
		return new(srl)
	})
}

type srl struct {
	target *target.Target
	log    logging.Logger
	parser *parser.Parser
	//deviceDetails *ndddvrv1.DeviceDetails
}

func (d *srl) Init(opts ...device.DeviceOption) error {
	for _, o := range opts {
		o(d)
	}
	return nil
}

func (d *srl) WithTarget(target *target.Target) {
	d.target = target
}

func (d *srl) WithLogging(log logging.Logger) {
	d.log = log
}

func (d *srl) WithParser(log logging.Logger) {
	d.parser = parser.NewParser(parser.WithLogger((log)))
}

func (d *srl) SupportedModels(ctx context.Context) ([]*gnmi.ModelData, error) {
	d.log.Debug("verifying capabilities ...")

	ext := new(gnmi_ext.Extension)
	resp, err := d.target.Capabilities(ctx, ext)
	if err != nil {
		return nil, errors.Wrap(err, errGetGnmiCapabilities)
	}
	//t.log.Debug("Gnmi Capability", "response", resp)

	return resp.SupportedModels, nil
}

func (d *srl) Discover(ctx context.Context) (*ndrv1.DeviceDetails, error) {
	d.log.Debug("Discover SRL details ...")
	dDetails, err := d.getDeviceDetails(ctx)
	if err != nil {
		return nil, err
	}
	d.log.Debug("SRL %s discoverd: %+v", d.target.Config.Name, *dDetails)
	return dDetails, nil
}

func (d *srl) GetConfig(ctx context.Context) (interface{}, error) {
	var err error
	var req *gnmi.GetRequest
	var rsp *gnmi.GetResponse

	req, err = gapi.NewGetRequest(
		gapi.Path(""),
		gapi.DataType("config"),
		gapi.Encoding("json_ietf"),
	)
	if err != nil {
		d.log.Debug(errGnmiCreateGetRequest, "error", err)
		return nil, errors.Wrap(err, errGnmiCreateGetRequest)
	}
	rsp, err = d.target.Get(ctx, req)
	if err != nil {
		d.log.Debug(errGnmiGet, "error", err)
		return nil, errors.Wrap(err, errGnmiGet)
	}
	//
	// expects a GetResponse with a single notification which
	// in turn has a single update.
	ns := rsp.GetNotification()
	switch len(ns) {
	case 1:
		upds := ns[0].GetUpdate()
		switch len(upds) {
		case 1:
			return yparser.GetValue(upds[0].GetVal())
		default:
			return nil, errors.New("unexpected number of updates in GetResponse Notification")
		}
	default:
		return nil, errors.New("unexpected number of Notifications in GetResponse")
	}
}

func (d *srl) GNMIGet(ctx context.Context, opts ...gapi.GNMIOption) (*gnmi.GetResponse, error) {
	var err error
	var req *gnmi.GetRequest
	var rsp *gnmi.GetResponse

	req, err = gapi.NewGetRequest(opts...)
	if err != nil {
		d.log.Debug(errGnmiCreateGetRequest, "error", err)
		return nil, errors.Wrap(err, errGnmiCreateGetRequest)
	}
	rsp, err = d.target.Get(ctx, req)
	if err != nil {
		d.log.Debug(errGnmiGet, "error", err)
		return nil, errors.Wrap(err, errGnmiGet)
	}
	return rsp, nil
}

func (d *srl) GNMISet(ctx context.Context, u []*gnmi.Update, p []*gnmi.Path) (*gnmi.SetResponse, error) {
	resp, err := d.target.Set(ctx, &gnmi.SetRequest{
		Update: u,
		Delete: p,
	})
	if err != nil {
		d.log.Debug(errGnmiSet, "error", err)
		return nil, err
	}
	//d.log.Debug("set response:", "resp", resp)
	return resp, nil
}

func (d *srl) getDeviceDetails(ctx context.Context) (*ndrv1.DeviceDetails, error) {
	req, err := gapi.NewGetRequest(
		gapi.Path(swVersionPath),
		gapi.Path(chassisPath),
		gapi.Encoding("ascii"),
		gapi.DataType("state"),
	)
	if err != nil {
		return nil, err
	}
	resp, err := d.target.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	devDetails := &ndrv1.DeviceDetails{
		Type: nddv1.DeviceTypePtr(DeviceType),
	}
	for _, notif := range resp.GetNotification() {
		for _, upd := range notif.GetUpdate() {
			p := gutils.GnmiPathToXPath(upd.GetPath(), true)
			switch p {
			case "platform/control/software-version":
				if devDetails.SwVersion == nil {
					devDetails.SwVersion = utils.StringPtr(upd.GetVal().GetStringVal())
				}
			case "platform/chassis/type":
				devDetails.Kind = utils.StringPtr(upd.GetVal().GetStringVal())
			case "platform/chassis/serial-number":
				devDetails.SerialNumber = utils.StringPtr(upd.GetVal().GetStringVal())
			case "platform/chassis/hw-mac-address":
				devDetails.MacAddress = utils.StringPtr(upd.GetVal().GetStringVal())
			}
		}
	}
	return devDetails, nil
}
