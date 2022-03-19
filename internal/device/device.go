/*
Copyright 2021 NDD.

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

package device

import (
	"context"

	gapi "github.com/karimra/gnmic/api"
	"github.com/karimra/gnmic/target"
	"github.com/openconfig/gnmi/proto/gnmi"
	ndddvrv1 "github.com/yndd/ndd-core/apis/dvr/v1"
	nddv1 "github.com/yndd/ndd-runtime/apis/common/v1"
	"github.com/yndd/ndd-runtime/pkg/logging"
)

type Device interface {
	// Init initializes the device
	Init(...DeviceOption) error
	// WithTarget, initializes the device target
	WithTarget(target *target.Target)
	// WithLogging initializes the device logging
	WithLogging(log logging.Logger)
	// Discover, discovers the device and its respective data
	Discover(ctx context.Context) (*ndddvrv1.DeviceDetails, error)
	// retrieve device supported models using gNMI capabilities RPC
	SupportedModels(ctx context.Context) ([]*gnmi.ModelData, error)
	// GetConfig, gets the config from the device
	GetConfig(ctx context.Context) (interface{}, error)
	// Get, gets the gnmi path from the tree
	GNMIGet(ctx context.Context, opts ...gapi.GNMIOption) (*gnmi.GetResponse, error)
	// Set creates a single transaction for updates and deletes
	GNMISet(ctx context.Context, updates []*gnmi.Update, deletes []*gnmi.Path) (*gnmi.SetResponse, error)
}

var Devices = map[nddv1.DeviceType]Initializer{}

type Initializer func() Device

func Register(name nddv1.DeviceType, initFn Initializer) {
	Devices[name] = initFn
}

type DeviceOption func(Device)

func WithTarget(target *target.Target) DeviceOption {
	return func(d Device) {
		d.WithTarget(target)
	}
}

func WithLogging(log logging.Logger) DeviceOption {
	return func(d Device) {
		d.WithLogging(log)
	}
}
