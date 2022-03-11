package srl3

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/openconfig/ygot/ygot"
	nddv1 "github.com/yndd/ndd-runtime/apis/common/v1"
	"github.com/yndd/ndd-runtime/pkg/meta"
	"github.com/yndd/nddo-runtime/pkg/odns"
	"github.com/yndd/nddo-runtime/pkg/resource"
	srlv1alpha1 "github.com/yndd/nddp-srl3/apis/srl3/v1alpha1"
	"github.com/yndd/nddp-srl3/pkg/ygotsrl"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	errDeleteDevice = "cannot delete device"
	errGetDevice    = "cannot get device"
)

type Device interface {
	// methods data
	Get() *ygotsrl.Device
	// methods schema
	Print() error
	Deploy(ctx context.Context, mg resource.Managed, deviceName string, labels map[string]string) error
	Destroy(ctx context.Context, mg resource.Managed, deviceName string, labels map[string]string) error
	List(ctx context.Context, mg resource.Managed, resources map[string]map[string]struct{}) error
	Validate(ctx context.Context, mg resource.Managed, deviceName string, resources map[string]map[string]struct{}) error
	Delete(ctx context.Context, mg resource.Managed, resources map[string]map[string]struct{}) error
	//ListResourcesByTransaction(ctx context.Context, cr srlv1alpha1.IFSrlTransaction, resources map[string]map[string]map[string]interface{}) error
}

func NewDevice(c resource.ClientApplicator, p Schema, name string) Device {
	newDeviceList := func() srlv1alpha1.IFSrl3DeviceList { return &srlv1alpha1.Srl3DeviceList{} }
	return &device{
		// k8s client
		client: c,
		// key
		name: name,
		// parent
		parent: p,
		// data
		device:        &ygotsrl.Device{},
		newDeviceList: newDeviceList,
	}
}

type device struct {
	// k8s client
	client resource.ClientApplicator
	// key
	name string
	// parent
	parent Schema
	// children
	// Data
	device        *ygotsrl.Device
	newDeviceList func() srlv1alpha1.IFSrl3DeviceList
}

func (x *device) Get() *ygotsrl.Device {
	return x.device
}

func (x *device) Print() error {
	deviceJsonConfig, err := ygot.EmitJSON(x.device, &ygot.EmitJSONConfig{
		Format:        ygot.RFC7951,
		RFC7951Config: &ygot.RFC7951JSONConfig{},
	})
	if err != nil {
		return err
	}
	fmt.Printf("Device: %s \n config: %s\n", x.name, deviceJsonConfig)
	return nil
}

func (x *device) Deploy(ctx context.Context, mg resource.Managed, deviceName string, labels map[string]string) error {
	cr, err := x.buildCR(mg, deviceName, labels)
	if err != nil {
		return err
	}
	return x.client.Apply(ctx, cr)
}

func (x *device) Destroy(ctx context.Context, mg resource.Managed, deviceName string, labels map[string]string) error {
	cr, err := x.buildCR(mg, deviceName, labels)
	if err != nil {
		return err
	}
	return x.client.Delete(ctx, cr)
}

func (x *device) List(ctx context.Context, mg resource.Managed, resources map[string]map[string]struct{}) error {
	// local CR list
	opts := []client.ListOption{
		client.MatchingLabels{srlv1alpha1.LabelNddaOwner: odns.GetOdnsResourceKindName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind))},
	}
	list := x.newDeviceList()
	if err := x.client.List(ctx, list, opts...); err != nil {
		return err
	}

	var Empty struct{}
	for _, i := range list.GetDevices() {
		if _, ok := resources[i.GetObjectKind().GroupVersionKind().Kind]; !ok {
			resources[i.GetObjectKind().GroupVersionKind().Kind] = make(map[string]struct{})
		}
		resources[i.GetObjectKind().GroupVersionKind().Kind][i.GetName()] = Empty
	}

	return nil
}

func (x *device) Validate(ctx context.Context, mg resource.Managed, deviceName string, resources map[string]map[string]struct{}) error {
	// local CR validation
	resourceName := odns.GetOdnsResourceName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind),
		[]string{
			strings.ToLower(deviceName)})

	if r, ok := resources[srlv1alpha1.DeviceKindKind]; ok {
		delete(r, resourceName)
	}

	return nil
}

func (x *device) Delete(ctx context.Context, mg resource.Managed, resources map[string]map[string]struct{}) error {
	// local CR deletion
	if res, ok := resources[srlv1alpha1.DeviceKindKind]; ok {
		for resName := range res {
			o := &srlv1alpha1.Srl3Device{
				ObjectMeta: metav1.ObjectMeta{
					Name:      resName,
					Namespace: mg.GetNamespace(),
				},
			}
			if err := x.client.Delete(ctx, o); err != nil {
				return err
			}
		}
	}
	return nil
}

func (x *device) buildCR(mg resource.Managed, deviceName string, labels map[string]string) (*srlv1alpha1.Srl3Device, error) {
	resourceName := odns.GetOdnsResourceName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind),
		[]string{
			//strings.ToLower(x.name),
			strings.ToLower(deviceName)})

	labels[srlv1alpha1.LabelNddaDeploymentPolicy] = string(mg.GetDeploymentPolicy())
	labels[srlv1alpha1.LabelNddaOwner] = odns.GetOdnsResourceKindName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind))
	labels[srlv1alpha1.LabelNddaOwnerGeneration] = strconv.Itoa(int(mg.GetGeneration()))
	labels[srlv1alpha1.LabelNddaDevice] = deviceName
	//labels[srlv1alpha1.LabelNddaItfce] = itfceName

	namespace := mg.GetNamespace()
	if namespace == "" {
		namespace = "default"
	}

	var d interface{}
	d, err := ygot.ConstructIETFJSON(x.device, &ygot.RFC7951JSONConfig{})
	if err != nil {
		return nil, err
	}
	device, ok := d.(srlv1alpha1.Device)
	if !ok {
		return nil, errors.New("wrong device object")
	}

	return &srlv1alpha1.Srl3Device{
		ObjectMeta: metav1.ObjectMeta{
			Name:            resourceName,
			Namespace:       namespace,
			Labels:          labels,
			OwnerReferences: []metav1.OwnerReference{meta.AsController(meta.TypedReferenceTo(mg, mg.GetObjectKind().GroupVersionKind()))},
		},
		Spec: srlv1alpha1.DeviceSpec{
			ResourceSpec: nddv1.ResourceSpec{
				NetworkNodeReference: &nddv1.Reference{
					Name: deviceName,
				},
			},
			Device: &device,
		},
	}, nil
}
