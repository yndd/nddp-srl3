package v1alpha1

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/openconfig/ygot/ygot"
	nddv1 "github.com/yndd/ndd-runtime/apis/common/v1"
	"github.com/yndd/ndd-runtime/pkg/meta"
	"github.com/yndd/nddo-intent-runtime/pkg/intent"
	"github.com/yndd/nddo-runtime/pkg/odns"
	"github.com/yndd/nddo-runtime/pkg/resource"
	srlv1alpha1 "github.com/yndd/nddp-srl3/apis/srl3/v1alpha1"
	"github.com/yndd/nddp-srl3/pkg/ygotsrl"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func InitSrl(c resource.ClientApplicator, p intent.Intent, name string) intent.Intent {
	newDeviceList := func() srlv1alpha1.IFSrl3DeviceList { return &srlv1alpha1.Srl3DeviceList{} }
	return &srlintent{
		client:        c,
		name:          name,
		parent:        p,
		newDeviceList: newDeviceList,
	}
}

type srlintent struct {
	// k8s client
	client resource.ClientApplicator
	// key
	name string
	// parent
	parent intent.Intent
	// children
	// Data
	device        *ygotsrl.Device
	newDeviceList func() srlv1alpha1.IFSrl3DeviceList
}

func (x *srlintent) GetData() interface{} {
	return x.device
}

func (x *srlintent) Deploy(ctx context.Context, mg resource.Managed, labels map[string]string) error {
	cr, err := x.buildCR(mg, x.name, labels)
	if err != nil {
		return err
	}
	return x.client.Apply(ctx, cr)
}

func (x *srlintent) Destroy(ctx context.Context, mg resource.Managed, labels map[string]string) error {
	cr, err := x.buildCR(mg, x.name, labels)
	if err != nil {
		return err
	}
	return x.client.Delete(ctx, cr)
}

func (x *srlintent) List(ctx context.Context, mg resource.Managed, resources map[string]map[string]struct{}) (map[string]map[string]struct{}, error) {
	// local CR list
	opts := []client.ListOption{
		client.MatchingLabels{srlv1alpha1.LabelNddaOwner: odns.GetOdnsResourceKindName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind))},
	}
	list := x.newDeviceList()
	if err := x.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}

	var Empty struct{}
	for _, i := range list.GetDevices() {
		if _, ok := resources[i.GetObjectKind().GroupVersionKind().Kind]; !ok {
			resources[i.GetObjectKind().GroupVersionKind().Kind] = make(map[string]struct{})
		}
		resources[i.GetObjectKind().GroupVersionKind().Kind][i.GetName()] = Empty
	}

	return resources, nil
}

func (x *srlintent) Validate(ctx context.Context, mg resource.Managed, resources map[string]map[string]struct{}) (map[string]map[string]struct{}, error) {
	// local CR validation
	resourceName := odns.GetOdnsResourceName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind),
		[]string{
			strings.ToLower(x.name)})

	if r, ok := resources[srlv1alpha1.DeviceKindKind]; ok {
		delete(r, resourceName)
	}

	return resources, nil
}

func (x *srlintent) Delete(ctx context.Context, mg resource.Managed, resources map[string]map[string]struct{}) error {
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

func (x *srlintent) buildCR(mg resource.Managed, deviceName string, labels map[string]string) (*srlv1alpha1.Srl3Device, error) {
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

	j, err := ygot.EmitJSON(x.device, &ygot.EmitJSONConfig{
		Format:         ygot.RFC7951,
		SkipValidation: true,
	})
	if err != nil {
		return nil, err
	}

	var d srlv1alpha1.Device
	if err := json.Unmarshal([]byte(j), &d); err != nil {
		return nil, err
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
			Device: &d,
		},
	}, nil
}
