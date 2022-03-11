package srl3

import (
	"context"
	"fmt"

	"github.com/yndd/nddo-runtime/pkg/resource"
)

type Schema interface {
	// methods chidren
	NewDevice(c resource.ClientApplicator, name string) Device
	GetDevices() map[string]Device
	Print()
	// methods schema
	Deploy(ctx context.Context, mg resource.Managed, labels map[string]string) error
	Destroy(ctx context.Context, mg resource.Managed, labels map[string]string) error
	List(ctx context.Context, mg resource.Managed) (map[string]map[string]struct{}, error)
	Validate(ctx context.Context, mg resource.Managed, resources map[string]map[string]struct{}) (map[string]map[string]struct{}, error)
	Delete(ctx context.Context, mg resource.Managed, resources map[string]map[string]struct{}) error
}

func NewSchema(c resource.ClientApplicator) Schema {
	return &schema{
		// k8s client
		client: c,
		// parent is nil/root
		// children
		devices: make(map[string]Device),
		// data key
	}
}

type schema struct {
	// k8s client
	client resource.ClientApplicator
	// parent is nil/root
	// children
	devices map[string]Device
	// data is nil
}

func (x *schema) NewDevice(c resource.ClientApplicator, name string) Device {
	if _, ok := x.devices[name]; !ok {
		x.devices[name] = NewDevice(c, x, name)
	}
	return x.devices[name]
}

func (x *schema) GetDevices() map[string]Device {
	return x.devices
}

func (x *schema) Print() {
	fmt.Println("schema information")
	for _, d := range x.GetDevices() {
		d.Print()
	}
}

func (x *schema) Deploy(ctx context.Context, mg resource.Managed, labels map[string]string) error {
	for deviceName, d := range x.GetDevices() {
		if err := d.Deploy(ctx, mg, deviceName, labels); err != nil {
			return err
		}
	}
	return nil
}

func (x *schema) Destroy(ctx context.Context, mg resource.Managed, labels map[string]string) error {
	for deviceName, d := range x.GetDevices() {
		if err := d.Destroy(ctx, mg, deviceName, labels); err != nil {
			return err
		}
	}
	return nil
}

func (x *schema) List(ctx context.Context, mg resource.Managed) (map[string]map[string]struct{}, error) {
	resources := make(map[string]map[string]struct{})
	for _, d := range x.GetDevices() {
		if err := d.List(ctx, mg, resources); err != nil {
			return nil, err
		}
	}
	return resources, nil
}

func (x *schema) Validate(ctx context.Context, mg resource.Managed, resources map[string]map[string]struct{}) (map[string]map[string]struct{}, error) {
	for deviceName, d := range x.GetDevices() {
		if err := d.Validate(ctx, mg, deviceName, resources); err != nil {
			return nil, err
		}
	}
	return resources, nil
}

func (x *schema) Delete(ctx context.Context, mg resource.Managed, resources map[string]map[string]struct{}) error {
	for _, d := range x.GetDevices() {
		if err := d.Delete(ctx, mg, resources); err != nil {
			return err
		}
	}
	return nil
}
