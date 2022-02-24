/*
Copyright 2022 NDD.

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

package srl3

import (
	"context"
	"encoding/json"
	"fmt"

	//"strings"
	"time"

	"github.com/karimra/gnmic/target"
	gnmitypes "github.com/karimra/gnmic/types"
	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/openconfig/gnmi/proto/gnmi_ext"
	"github.com/pkg/errors"
	ndrv1 "github.com/yndd/ndd-core/apis/dvr/v1"
	nddv1 "github.com/yndd/ndd-runtime/apis/common/v1"
	"github.com/yndd/ndd-runtime/pkg/event"
	"github.com/yndd/ndd-runtime/pkg/logging"
	"github.com/yndd/ndd-runtime/pkg/reconciler/managed"
	"github.com/yndd/ndd-runtime/pkg/resource"
	"github.com/yndd/ndd-runtime/pkg/utils"
	"github.com/yndd/ndd-yang/pkg/leafref"
	"github.com/yndd/ndd-yang/pkg/yentry"
	"github.com/yndd/ndd-yang/pkg/yparser"
	"github.com/yndd/ndd-yang/pkg/yresource"
	"github.com/yndd/nddp-system/pkg/gvkresource"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	cevent "sigs.k8s.io/controller-runtime/pkg/event"

	//"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/source"

	srl3v1alpha1 "github.com/yndd/nddp-srl3/apis/srl3/v1alpha1"
	"github.com/yndd/nddp-srl3/internal/shared"
)

const (
	// Errors
	errUnexpectedDevice       = "the managed resource is not a Device resource"
	errKubeUpdateFailedDevice = "cannot update Device"
	errReadDevice             = "cannot read Device"
	errCreateDevice           = "cannot create Device"
	errUpdateDevice           = "cannot update Device"
	errDeleteDevice           = "cannot delete Device"
)

// SetupDevice adds a controller that reconciles Devices.
func SetupDevice(mgr ctrl.Manager, o controller.Options, nddcopts *shared.NddControllerOptions) (string, chan cevent.GenericEvent, error) {
	//func SetupDevice(mgr ctrl.Manager, o controller.Options, nddcopts *shared.NddControllerOptions) error {

	name := managed.ControllerName(srl3v1alpha1.DeviceGroupKind)

	events := make(chan cevent.GenericEvent)

	y := initYangDevice(
		nddcopts.DeviceSchema,
	)

	r := managed.NewReconciler(mgr,
		resource.ManagedKind(srl3v1alpha1.DeviceGroupVersionKind),
		managed.WithPollInterval(nddcopts.Poll),
		managed.WithExternalConnecter(&connectorDevice{
			log:          nddcopts.Logger,
			kube:         mgr.GetClient(),
			usage:        resource.NewNetworkNodeUsageTracker(mgr.GetClient(), &ndrv1.NetworkNodeUsage{}),
			deviceSchema: nddcopts.DeviceSchema,
			nddpSchema:   nddcopts.NddpSchema,
			y:            y,
			newClientFn:  target.NewTarget,
			gnmiAddress:  nddcopts.GnmiAddress},
		),
		managed.WithValidator(&validatorDevice{
			log:          nddcopts.Logger,
			deviceSchema: nddcopts.DeviceSchema,
			y:            y},
		),
		managed.WithLogger(nddcopts.Logger.WithValues("Device-controller", name)),
		managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))))

	DeviceHandler := &EnqueueRequestForAllDevice{
		client: mgr.GetClient(),
		log:    nddcopts.Logger,
		ctx:    context.Background(),
	}

	//return ctrl.NewControllerManagedBy(mgr).
	return srl3v1alpha1.DeviceGroupKind, events, ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(o).
		For(&srl3v1alpha1.Srl3Device{}).
		Owns(&srl3v1alpha1.Srl3Device{}).
		WithEventFilter(resource.IgnoreUpdateWithoutGenerationChangePredicate()).
		/*
			Watches(
				&source.Channel{Source: events},
				&handler.EnqueueRequestForObject{},
			).
		*/
		Watches(&source.Kind{Type: &srl3v1alpha1.Srl3Device{}}, DeviceHandler).
		Watches(&source.Channel{Source: events}, DeviceHandler).
		Complete(r)
}

type Device struct {
	*yresource.Resource
}

func initYangDevice(deviceSchema *yentry.Entry, opts ...yresource.Option) yresource.Handler {
	return &Device{&yresource.Resource{
		DeviceSchema: deviceSchema,
	}}

}

// GetRootPath returns the rootpath of the resource
func (r *Device) GetRootPath(mg resource.Managed) []*gnmi.Path {

	cr, ok := mg.(*srl3v1alpha1.Srl3Device)
	if !ok {
		fmt.Printf("wrong cr: %v\n", cr)
		return []*gnmi.Path{}
	}

	return []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{},
		},
	}
}

// GetParentDependency returns the parent dependency of the resource
func (r *Device) GetParentDependency(mg resource.Managed) []*leafref.LeafRef {
	rootPath := r.GetRootPath(mg)
	// if the path is not bigger than 1 element there is no parent dependency
	if len(rootPath[0].GetElem()) < 2 {
		return []*leafref.LeafRef{}
	}
	dependencyPath := r.DeviceSchema.GetParentDependency(rootPath[0], rootPath[0], "")
	// the dependency path is the rootPath except for the last element
	//dependencyPathElem := rootPath[0].GetElem()[:(len(rootPath[0].GetElem()) - 1)]
	// check for keys present, if no keys present we return an empty list
	keysPresent := false
	for _, pathElem := range dependencyPath.GetElem() {
		if len(pathElem.GetKey()) != 0 {
			keysPresent = true
		}
	}
	if !keysPresent {
		return []*leafref.LeafRef{}
	}

	// return the rootPath except the last entry
	return []*leafref.LeafRef{
		{
			RemotePath: dependencyPath,
		},
	}
}

type validatorDevice struct {
	log          logging.Logger
	deviceSchema *yentry.Entry
	y            yresource.Handler
}

func (v *validatorDevice) ValidateLeafRef(ctx context.Context, mg resource.Managed, cfg []byte) (managed.ValidateLeafRefObservation, error) {
	return managed.ValidateLeafRefObservation{
		Success:          true,
		ResolvedLeafRefs: []*leafref.ResolvedLeafRef{}}, nil
}

func (v *validatorDevice) ValidateParentDependency(ctx context.Context, mg resource.Managed, cfg []byte) (managed.ValidateParentDependencyObservation, error) {
	return managed.ValidateParentDependencyObservation{
		Success:          true,
		ResolvedLeafRefs: []*leafref.ResolvedLeafRef{}}, nil
}

// ValidateResourceIndexes validates if the indexes of a resource got changed
// if so we need to delete the original resource, because it will be dangling if we dont delete it
func (v *validatorDevice) ValidateResourceIndexes(ctx context.Context, mg resource.Managed) (managed.ValidateResourceIndexesObservation, error) {
	return managed.ValidateResourceIndexesObservation{Changed: false, ResourceIndexes: map[string]string{}}, nil
}

// A connector is expected to produce an ExternalClient when its Connect method
// is called.
type connectorDevice struct {
	log          logging.Logger
	kube         client.Client
	usage        resource.Tracker
	deviceSchema *yentry.Entry
	nddpSchema   *yentry.Entry
	y            yresource.Handler
	newClientFn  func(c *gnmitypes.TargetConfig) *target.Target
	gnmiAddress  string
}

// Connect produces an ExternalClient by:
// 1. Tracking that the managed resource is using a NetworkNode.
// 2. Getting the managed resource's NetworkNode with connection details
// A resource is mapped to a single target
func (c *connectorDevice) Connect(ctx context.Context, mg resource.Managed) (managed.ExternalClient, error) {
	log := c.log.WithValues("resource", mg.GetName())
	//log.Debug("Connect")

	cr, ok := mg.(*srl3v1alpha1.Srl3Device)
	if !ok {
		return nil, errors.New(errUnexpectedDevice)
	}
	if err := c.usage.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, errTrackTCUsage)
	}

	// find network node that is configured status
	nn := &ndrv1.NetworkNode{}
	if err := c.kube.Get(ctx, types.NamespacedName{Name: cr.GetNetworkNodeReference().Name}, nn); err != nil {
		return nil, errors.Wrap(err, errGetNetworkNode)
	}

	if nn.GetCondition(ndrv1.ConditionKindDeviceDriverConfigured).Status != corev1.ConditionTrue {
		return nil, errors.New(targetNotConfigured)
	}

	cfg := &gnmitypes.TargetConfig{
		Name:       "dummy",
		Address:    c.gnmiAddress,
		Username:   utils.StringPtr("admin"),
		Password:   utils.StringPtr("admin"),
		Timeout:    10 * time.Second,
		SkipVerify: utils.BoolPtr(true),
		Insecure:   utils.BoolPtr(true),
		TLSCA:      utils.StringPtr(""), //TODO TLS
		TLSCert:    utils.StringPtr(""), //TODO TLS
		TLSKey:     utils.StringPtr(""),
		Gzip:       utils.BoolPtr(false),
	}

	cl := target.NewTarget(cfg)
	if err := cl.CreateGNMIClient(ctx); err != nil {
		return nil, errors.Wrap(err, errNewClient)
	}

	tns := []string{nn.GetName()}

	return &externalDevice{client: cl, targets: tns, log: log, deviceSchema: c.deviceSchema, nddpSchema: c.nddpSchema, y: c.y}, nil
}

// An ExternalClient observes, then either creates, updates, or deletes an
// external resource to ensure it reflects the managed resource's desired state.
type externalDevice struct {
	client       *target.Target
	targets      []string
	log          logging.Logger
	deviceSchema *yentry.Entry
	nddpSchema   *yentry.Entry
	y            yresource.Handler
}

func (e *externalDevice) Close() {
	e.client.Close()
}

func (e *externalDevice) Observe(ctx context.Context, mg resource.Managed) (managed.ExternalObservation, error) {
	log := e.log.WithValues("Resource", mg.GetName())
	//log.Debug("Observing ...")

	cr, ok := mg.(*srl3v1alpha1.Srl3Device)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errUnexpectedDevice)
	}

	// rootpath of the resource
	rootPath := e.y.GetRootPath(cr)
	hierElements := e.deviceSchema.GetHierarchicalResourcesLocal(true, rootPath[0], &gnmi.Path{}, make([]*gnmi.Path, 0))
	//log.Debug("Observing hierElements ...", "Path", yparser.GnmiPath2XPath(rootPath[0], false), "hierElements", hierElements)

	gvkTransaction, err := gvkresource.GetGvkTransaction(mg)
	if err != nil {
		return managed.ExternalObservation{}, err
	}

	// gnmi get request
	req := &gnmi.GetRequest{
		//Prefix:   &gnmi.Path{Target: GnmiTarget, Origin: GnmiOrigin},
		Prefix:   &gnmi.Path{Target: shared.GetCrDeviceName(mg.GetNamespace(), cr.GetNetworkNodeReference().Name)},
		Path:     rootPath,
		Encoding: gnmi.Encoding_JSON,
		//Type:     gnmi.GetRequest_DataType(gnmi.GetRequest_STATE),
		Extension: []*gnmi_ext.Extension{
			{Ext: &gnmi_ext.Extension_RegisteredExt{
				RegisteredExt: &gnmi_ext.RegisteredExtension{Id: gnmi_ext.ExtensionID_EID_EXPERIMENTAL, Msg: []byte(gvkTransaction)}}},
		},
	}

	// gnmi get response
	exists := true
	resp, err := e.client.Get(ctx, req)
	if err != nil {
		if er, ok := status.FromError(err); ok {
			switch er.Code() {
			case codes.ResourceExhausted:
				// we use this to signal the device or cache is exhausted
				return managed.ExternalObservation{
					Ready:      false,
					Exhausted:  true,
					Exists:     false,
					Pending:    true,
					Failed:     true,
					HasData:    false,
					IsUpToDate: false,
				}, nil
			case codes.Unavailable:
				// we use this to signal not ready
				return managed.ExternalObservation{
					Ready:      false,
					Exists:     false,
					Pending:    true,
					Failed:     true,
					HasData:    false,
					IsUpToDate: false,
				}, nil
			case codes.NotFound:
				// the k8s resource does not exists but the data can still exist
				// if data exists it means we go from UMR -> MR
				exists = false
			case codes.AlreadyExists:
				// the system cache has the resource but the action did not complete so we should skip the next reconcilation
				// loop and wait
				return managed.ExternalObservation{
					Ready:      true,
					Exhausted:  false,
					Exists:     true,
					Pending:    true,
					Failed:     false,
					HasData:    false,
					IsUpToDate: false,
				}, nil
			case codes.FailedPrecondition:
				// the k8s resource exists but is in failed status, compare the response spec with current spec
				// if the specs are equal return observation.ResponseSuccess -> False
				// if the specs are not equal follow the regular procedure
				//log.Debug("observing when using gnmic: resource failed")
				failedObserve, err := processObserve(rootPath[0], hierElements, &cr.Spec, resp, e.deviceSchema)
				if err != nil {
					return managed.ExternalObservation{}, err
				}
				if failedObserve.upToDate {
					// there is no difference between the previous spec and the current spec, so we dont retry
					// given the previous attempt failed
					return managed.ExternalObservation{
						Ready:      true,
						Exists:     true,
						Pending:    false,
						Failed:     true,
						HasData:    false,
						IsUpToDate: false,
					}, nil
				} else {
					// this should trigger an update
					return managed.ExternalObservation{
						Ready:      true,
						Exists:     true,
						Pending:    false,
						Failed:     false,
						HasData:    true,
						IsUpToDate: false,
					}, nil
				}
			}
		}
	}

	// processObserve
	// o. marshal/unmarshal data
	// 1. check if resource exists
	// 2. remove parent hierarchical elements from spec
	// 3. remove resource hierarchical elements from gnmi response
	// 4. remove state
	// 5. transform the data in gnmi to process the delta
	// 6. find the resource delta: updates and/or deletes in gnmi
	//exists, deletes, updates, b, err := processObserve(rootPath[0], hierElements, &cr.Spec, resp, e.deviceSchema)
	//e.log.Debug("processObserve", "notification", resp.GetNotification())
	observe, err := processObserve(rootPath[0], hierElements, &cr.Spec, resp, e.deviceSchema)
	if err != nil {
		return managed.ExternalObservation{}, err
	}
	if !observe.hasData {
		// No Data exists -> Create it or Delete is complete
		//log.Debug("Observing Response:", "observe", observe, "exists", exists, "Response", resp)
		return managed.ExternalObservation{
			Ready:      true,
			Exists:     exists,
			Pending:    false,
			Failed:     false,
			HasData:    false,
			IsUpToDate: false,
		}, nil
	}
	// Data exists

	if !observe.upToDate {
		// resource is NOT up to date
		log.Debug("Observing Response: resource NOT up to date", "Observe", observe, "exists", exists, "Response", resp)
		return managed.ExternalObservation{
			Ready:      true,
			Exists:     exists,
			Pending:    false,
			Failed:     false,
			HasData:    true,
			IsUpToDate: false,
			//ResourceDeletes:  observe.deletes,
			//ResourceUpdates:  observe.updates,
		}, nil
	}
	// resource is up to date
	//log.Debug("Observing Response: resource up to date", "Observe", observe, "Response", resp)
	return managed.ExternalObservation{
		Ready:      true,
		Exists:     exists,
		Pending:    false,
		Failed:     false,
		HasData:    true,
		IsUpToDate: true,
	}, nil
}

func (e *externalDevice) Create(ctx context.Context, mg resource.Managed, ignoreTransaction bool) error {
	//log := e.log.WithValues("Resource", mg.GetName())
	//log.Debug("Creating ...")

	cr, ok := mg.(*srl3v1alpha1.Srl3Device)
	if !ok {
		return errors.New(errUnexpectedDevice)
	}

	// get the rootpath of the resource
	rootPath := e.y.GetRootPath(mg)

	// create k8s object
	// processCreate
	// 0. marshal/unmarshal data
	// 1. transform the spec data to gnmi updates
	updates, err := processCreateK8s(mg, rootPath[0], &cr.Spec, e.deviceSchema, e.nddpSchema, ignoreTransaction)
	if err != nil {
		return errors.Wrap(err, errCreateObject)
	}
	/*
		for _, update := range updates {
			log.Debug("Create Fine Grane Updates", "Path", yparser.GnmiPath2XPath(update.Path, true), "Value", update.GetVal())
		}

		if len(updates) == 0 {
			log.Debug("cannot create object since there are no updates present")
			return errors.Wrap(err, errCreateObject)
		}
	*/

	crSystemDeviceName := shared.GetCrSystemDeviceName(shared.GetCrDeviceName(mg.GetNamespace(), mg.GetNetworkNodeReference().Name))

	req := &gnmi.SetRequest{
		Prefix:  &gnmi.Path{Target: crSystemDeviceName},
		Replace: updates,
	}

	_, err = e.client.Set(ctx, req)
	if err != nil {
		return errors.Wrap(err, errCreateDevice)
	}

	return nil
}

func (e *externalDevice) Update(ctx context.Context, mg resource.Managed, obs managed.ExternalObservation) error {
	//log := e.log.WithValues("Resource", mg.GetName())
	//log.Debug("Updating ...")

	cr, ok := mg.(*srl3v1alpha1.Srl3Device)
	if !ok {
		return errors.New(errUnexpectedDevice)
	}

	// get the rootpath of the resource
	rootPath := e.y.GetRootPath(mg)

	updates, err := processUpdateK8s(mg, rootPath[0], &cr.Spec, e.deviceSchema, e.nddpSchema)
	if err != nil {
		return errors.Wrap(err, errUpdateDevice)
	}
	/*
		for _, update := range updates {
			log.Debug("Update Fine Grane Updates", "Path", yparser.GnmiPath2XPath(update.Path, true), "Value", update.GetVal())
		}
	*/

	crSystemDeviceName := shared.GetCrSystemDeviceName(shared.GetCrDeviceName(mg.GetNamespace(), mg.GetNetworkNodeReference().Name))

	req := gnmi.SetRequest{
		Prefix:  &gnmi.Path{Target: crSystemDeviceName},
		Replace: updates,
	}

	_, err = e.client.Set(ctx, &req)
	if err != nil {
		return errors.Wrap(err, errUpdateDevice)
	}

	return nil
}

func (e *externalDevice) Delete(ctx context.Context, mg resource.Managed) error {
	//log := e.log.WithValues("Resource", mg.GetName())
	//log.Debug("Deleting ...")

	// get the rootpath of the resource
	rootPath := e.y.GetRootPath(mg)

	updates, err := processDeleteK8sResource(mg, rootPath[0], e.nddpSchema)
	if err != nil {
		return errors.Wrap(err, errDeleteDevice)
	}
	/*
		for _, update := range updates {
			log.Debug("Delete Fine Grane Updates", "Path", yparser.GnmiPath2XPath(update.Path, true), "Value", update.GetVal())
		}
	*/

	crSystemDeviceName := shared.GetCrSystemDeviceName(shared.GetCrDeviceName(mg.GetNamespace(), mg.GetNetworkNodeReference().Name))

	req := gnmi.SetRequest{
		Prefix:  &gnmi.Path{Target: crSystemDeviceName},
		Replace: updates,
	}

	_, err = e.client.Set(ctx, &req)
	if err != nil {
		return errors.Wrap(err, errDeleteDevice)
	}

	return nil
}

func (e *externalDevice) GetTarget() []string {
	return e.targets
}

func (e *externalDevice) GetConfig(ctx context.Context, mg resource.Managed) ([]byte, error) {
	//e.log.Debug("Get Config ...")
	req := &gnmi.GetRequest{
		Prefix: &gnmi.Path{Target: shared.GetCrDeviceName(mg.GetNamespace(), mg.GetNetworkNodeReference().Name)},
		Path: []*gnmi.Path{
			{
				Elem: []*gnmi.PathElem{},
			},
		},
		Encoding: gnmi.Encoding_JSON,
		//Type:     gnmi.GetRequest_DataType(gnmi.GetRequest_CONFIG),
	}

	resp, err := e.client.Get(ctx, req)
	if err != nil {
		return make([]byte, 0), errors.Wrap(err, errGetConfig)
	}

	if len(resp.GetNotification()) != 0 {
		if len(resp.GetNotification()[0].GetUpdate()) != 0 {
			x2, err := yparser.GetValue(resp.GetNotification()[0].GetUpdate()[0].Val)
			if err != nil {
				return make([]byte, 0), errors.Wrap(err, errGetConfig)
			}

			data, err := json.Marshal(x2)
			if err != nil {
				return make([]byte, 0), errors.Wrap(err, errJSONMarshal)
			}
			return data, nil
		}
	}
	//e.log.Debug("Get Config Empty response")
	return nil, nil
}

func (e *externalDevice) GetResourceName(ctx context.Context, mg resource.Managed, path *gnmi.Path) (string, error) {
	//e.log.Debug("Get GetResourceName ...", "remotePath", yparser.GnmiPath2XPath(path, true))
	crSystemDeviceName := shared.GetCrSystemDeviceName(shared.GetCrDeviceName(mg.GetNamespace(), mg.GetNetworkNodeReference().Name))

	// gnmi get request
	req := &gnmi.GetRequest{
		Prefix:   &gnmi.Path{Target: crSystemDeviceName},
		Path:     []*gnmi.Path{path},
		Encoding: gnmi.Encoding_JSON,
		Extension: []*gnmi_ext.Extension{
			{Ext: &gnmi_ext.Extension_RegisteredExt{
				RegisteredExt: &gnmi_ext.RegisteredExtension{Id: gnmi_ext.ExtensionID_EID_EXPERIMENTAL, Msg: []byte(gvkresource.Operation_GetResourceNameFromPath)}}},
		},
	}

	// gnmi get response
	resp, err := e.client.Get(ctx, req)
	if err != nil {
		return "", errors.Wrap(err, errGetResourceName)
	}

	x2, err := yparser.GetValue(resp.GetNotification()[0].GetUpdate()[0].Val)
	if err != nil {
		return "", errors.Wrap(err, errJSONMarshal)
	}

	d, err := json.Marshal(x2)
	if err != nil {
		return "", errors.Wrap(err, errJSONMarshal)
	}

	var resourceName nddv1.ResourceName
	if err := json.Unmarshal(d, &resourceName); err != nil {
		return "", errors.Wrap(err, errJSONUnMarshal)
	}

	//e.log.Debug("Get ResourceName Response", "remotePath", yparser.GnmiPath2XPath(path, true), "ResourceName", resourceName)

	return resourceName.Name, nil
}
