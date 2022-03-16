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

package srl

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	//"strings"
	"time"

	"github.com/karimra/gnmic/target"
	gnmitypes "github.com/karimra/gnmic/types"
	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/openconfig/gnmi/proto/gnmi_ext"
	"github.com/pkg/errors"
	ndrv1 "github.com/yndd/ndd-core/apis/dvr/v1"
	"github.com/yndd/ndd-runtime/pkg/event"
	"github.com/yndd/ndd-runtime/pkg/logging"
	"github.com/yndd/ndd-runtime/pkg/reconciler/managed"
	"github.com/yndd/ndd-runtime/pkg/resource"
	"github.com/yndd/ndd-runtime/pkg/utils"
	//"github.com/yndd/ndd-yang/pkg/yentry"
	"github.com/yndd/ndd-yang/pkg/yparser"
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

	srlv1alpha1 "github.com/yndd/nddp-srl3/apis/srl3/v1alpha1"
	"github.com/yndd/nddp-srl3/internal/model"
	"github.com/yndd/nddp-srl3/internal/shared"
	"github.com/yndd/nddp-srl3/pkg/ygotsrl"
	"github.com/yndd/nddp-system/pkg/ygotnddp"
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

	name := managed.ControllerName(srlv1alpha1.DeviceGroupKind)

	events := make(chan cevent.GenericEvent)

	dm := &model.Model{
		StructRootType:  reflect.TypeOf((*ygotsrl.Device)(nil)),
		SchemaTreeRoot:  ygotsrl.SchemaTree["Device"],
		JsonUnmarshaler: ygotsrl.Unmarshal,
		EnumData:        ygotsrl.ΛEnum,
	}

	sm := &model.Model{
		StructRootType:  reflect.TypeOf((*ygotnddp.Device)(nil)),
		SchemaTreeRoot:  ygotnddp.SchemaTree["Device"],
		JsonUnmarshaler: ygotnddp.Unmarshal,
		EnumData:        ygotnddp.ΛEnum,
	}

	r := managed.NewReconciler(mgr,
		resource.ManagedKind(srlv1alpha1.DeviceGroupVersionKind),
		managed.WithPollInterval(nddcopts.Poll),
		managed.WithExternalConnecter(&connectorDevice{
			log:   nddcopts.Logger,
			kube:  mgr.GetClient(),
			usage: resource.NewNetworkNodeUsageTracker(mgr.GetClient(), &ndrv1.NetworkNodeUsage{}),
			//deviceSchema: nddcopts.DeviceSchema,
			//nddpSchema:   nddcopts.NddpSchema,
			deviceModel: dm,
			systemModel: sm,
			newClientFn: target.NewTarget,
			gnmiAddress: nddcopts.GnmiAddress},
		),
		managed.WithValidator(&validatorDevice{
			log: nddcopts.Logger,
			//deviceSchema: nddcopts.DeviceSchema,
			deviceModel: dm,
			systemModel: sm,
		},
		),
		managed.WithLogger(nddcopts.Logger.WithValues("Srl3Device", name)),
		managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))))

	DeviceHandler := &EnqueueRequestForAllDevice{
		client: mgr.GetClient(),
		log:    nddcopts.Logger,
		ctx:    context.Background(),
	}

	//return ctrl.NewControllerManagedBy(mgr).
	return srlv1alpha1.DeviceGroupKind, events, ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(o).
		For(&srlv1alpha1.Srl3Device{}).
		Owns(&srlv1alpha1.Srl3Device{}).
		WithEventFilter(resource.IgnoreUpdateWithoutGenerationChangePredicate()).
		/*
			Watches(
				&source.Channel{Source: events},
				&handler.EnqueueRequestForObject{},
			).
		*/
		Watches(&source.Kind{Type: &srlv1alpha1.Srl3Device{}}, DeviceHandler).
		Watches(&source.Channel{Source: events}, DeviceHandler).
		Complete(r)
}

type validatorDevice struct {
	log logging.Logger
	//deviceSchema *yentry.Entry
	deviceModel *model.Model
	systemModel *model.Model
}

func printRootPaths(crName string, crRootPaths []string) {
	fmt.Println("++++++++++++++++++++++++++++++++++++++")
	fmt.Printf("rootPaths for cr: %s\n", crName)
	for _, p := range crRootPaths {
		fmt.Printf("  rootPath: %s\n", p)
	}
	fmt.Println("++++++++++++++++++++++++++++++++++++++")
}

func printHierPaths(crName string, crHierPaths map[string][]string) {
	fmt.Println("++++++++++++++++++++++++++++++++++++++")
	fmt.Printf("hierPaths for cr: %s\n", crName)
	for p, hierPaths := range crHierPaths {
		for _, hierPath := range hierPaths {
			fmt.Printf("  rootPath: %s hierPath: %s\n", p, hierPath)
		}
	}
	fmt.Println("++++++++++++++++++++++++++++++++++++++")

}

// ValidateResourceIndexes validates if the indexes of a resource got changed
// if so we need to delete the original resource, because it will be dangling if we dont delete it
func (v *validatorDevice) ValidateRootPaths(ctx context.Context, mg resource.Managed, resourceList map[string]*ygotnddp.NddpSystem_Gvk) (managed.ValidateRootPathsObservation, error) {
	log := v.log.WithValues("Resource", mg.GetName())
	log.Debug("ValidateRootPaths ...")

	cr, ok := mg.(*srlv1alpha1.Srl3Device)
	if !ok {
		return managed.ValidateRootPathsObservation{}, errors.New(errUnexpectedDevice)
	}

	for resourceName := range resourceList {
		log.Debug("ValidateRootPaths resourceList", "resourceName", resourceName)
	}

	crRootPaths, err := v.getRootPaths(cr.Spec.Device)
	if err != nil {
		return managed.ValidateRootPathsObservation{}, err
	}

	rootPaths := []string{}
	for _, crRootPath := range crRootPaths {
		//log.Debug("ValidateRootPaths rootPaths", "path", yparser.GnmiPath2XPath(crRootPath, true))
		rootPaths = append(rootPaths, yparser.GnmiPath2XPath(crRootPath, true))
	}

	hierPaths, err := getHierPaths(mg, crRootPaths, resourceList)
	if err != nil {
		return managed.ValidateRootPathsObservation{}, err
	}
	//log.Debug("ValidateRootPaths", "hierPaths", hierPaths)

	hierRootPaths := map[string][]string{}
	for rootPath, crRootPaths := range hierPaths {
		for _, crRootPath := range crRootPaths {
			//log.Debug("findPaths", "path", yparser.GnmiPath2XPath(crRootPath, true))
			if strings.HasPrefix(yparser.GnmiPath2XPath(crRootPath, true), "/routing-policy") {
				break
			}
			hierRootPaths[rootPath] = append(hierRootPaths[rootPath], yparser.GnmiPath2XPath(crRootPath, true))
		}
	}

	printRootPaths(mg.GetName(), rootPaths)
	printHierPaths(mg.GetName(), hierRootPaths)

	return managed.ValidateRootPathsObservation{
		Changed:     false,
		RootPaths:   rootPaths,
		HierPaths:   hierRootPaths,
		DeletePaths: []*gnmi.Path{}, // TODO find delta
	}, nil
}

// A connector is expected to produce an ExternalClient when its Connect method
// is called.
type connectorDevice struct {
	log   logging.Logger
	kube  client.Client
	usage resource.Tracker
	//deviceSchema *yentry.Entry
	//nddpSchema   *yentry.Entry
	deviceModel *model.Model
	systemModel *model.Model
	//y            yresource.Handler
	newClientFn func(c *gnmitypes.TargetConfig) *target.Target
	gnmiAddress string
}

// Connect produces an ExternalClient by:
// 1. Tracking that the managed resource is using a NetworkNode.
// 2. Getting the managed resource's NetworkNode with connection details
// A resource is mapped to a single target
func (c *connectorDevice) Connect(ctx context.Context, mg resource.Managed) (managed.ExternalClient, error) {
	log := c.log.WithValues("resource", mg.GetName())
	//log.Debug("Connect")

	cr, ok := mg.(*srlv1alpha1.Srl3Device)
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
		Name:       cr.GetNetworkNodeReference().Name,
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

	//return &externalDevice{client: cl, targets: tns, log: log, deviceSchema: c.deviceSchema, nddpSchema: c.nddpSchema, deviceModel: c.deviceModel, systemModel: c.systemModel}, nil
	return &externalDevice{client: cl, targets: tns, log: log, deviceModel: c.deviceModel, systemModel: c.systemModel}, nil
}

// An ExternalClient observes, then either creates, updates, or deletes an
// external resource to ensure it reflects the managed resource's desired state.
type externalDevice struct {
	client  *target.Target
	targets []string
	log     logging.Logger
	//deviceSchema *yentry.Entry
	//nddpSchema   *yentry.Entry
	deviceModel *model.Model
	systemModel *model.Model
}

func (e *externalDevice) Close() {
	e.client.Close()
}

func (e *externalDevice) Observe(ctx context.Context, mg resource.Managed) (managed.ExternalObservation, error) {
	log := e.log.WithValues("Resource", mg.GetName())
	log.Debug("Observing ...")

	cr, ok := mg.(*srlv1alpha1.Srl3Device)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errUnexpectedDevice)
	}

	// TODO get paths dynamically

	gvkName := gvkresource.GetGvkName(mg)

	crDeviceName := "ygot." + shared.GetCrDeviceName(mg.GetNamespace(), cr.GetNetworkNodeReference().Name)

	// gnmi get request
	req := &gnmi.GetRequest{
		Prefix:   &gnmi.Path{Target: crDeviceName},
		Path:     []*gnmi.Path{{}},
		Encoding: gnmi.Encoding_JSON,
		Extension: []*gnmi_ext.Extension{
			{Ext: &gnmi_ext.Extension_RegisteredExt{
				RegisteredExt: &gnmi_ext.RegisteredExtension{Id: gnmi_ext.ExtensionID_EID_EXPERIMENTAL, Msg: []byte(gvkName)}}},
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
				// TODO
				return managed.ExternalObservation{
					Ready:      true,
					Exists:     true,
					Pending:    false,
					Failed:     true,
					HasData:    false,
					IsUpToDate: false,
				}, nil
				/*
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
				*/
			}
		}
	}

	observe, err := e.processObserve(mg.GetRootPaths(), mg.GetHierPaths(), *cr.Spec.Device, resp)
	if err != nil {
		log.Debug("Observe Response", "error", err)
		if strings.Contains(err.Error(), "not found") {
			return managed.ExternalObservation{
				Ready:      true,
				Exists:     false, // we set exists to false to ensure the resource is recreated
				Pending:    false,
				Failed:     false,
				HasData:    false,
				IsUpToDate: false,
				Deletes:    []*gnmi.Path{},
				Updates:    []*gnmi.Update{},
			}, nil
		}
		return managed.ExternalObservation{}, err
	}
	log.Debug("Observe Response", "observe", observe)

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
			Deletes:    observe.deletes,
			Updates:    observe.updates,
		}, nil
	}
	// Data exists

	if !observe.upToDate {
		// resource is NOT up to date
		log.Debug("Observing Response: resource NOT up to date", "Observe", observe, "exists", exists)
		return managed.ExternalObservation{
			Ready:      true,
			Exists:     exists,
			Pending:    false,
			Failed:     false,
			HasData:    true,
			IsUpToDate: false,
			Deletes:    observe.deletes,
			Updates:    observe.updates,
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
		Deletes:    observe.deletes,
		Updates:    observe.updates,
	}, nil
	/*
		return managed.ExternalObservation{
			Ready:      true,
			Exists:     exists,
			Pending:    false,
			Failed:     false,
			HasData:    true,
			IsUpToDate: true,
		}, nil
	*/
}

func (e *externalDevice) Create(ctx context.Context, mg resource.Managed, obs managed.ExternalObservation) error {
	log := e.log.WithValues("Resource", mg.GetName())
	log.Debug("Creating ...")

	updates, err := e.getGvkUpate(mg, obs, ygotnddp.NddpSystem_ResourceAction_CREATE)
	if err != nil {
		return errors.Wrap(err, errCreateDevice)
	}

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
	log := e.log.WithValues("Resource", mg.GetName())
	log.Debug("Updating ...")

	/*
		cr, ok := mg.(*srlv1alpha1.Srl3Device)
		if !ok {
			return errors.New(errUnexpectedDevice)
		}

		crPaths, err := e.findPaths(cr.Spec.Device)
		if err != nil {
			return err
		}
	*/

	// create k8s object
	// processCreate
	// 0. marshal/unmarshal data
	// 1. transform the spec data to gnmi updates
	/*
		updates, err := e.processK8s(mg, crPaths, systemv1alpha1.E_GvkAction_Update)
		if err != nil {
			return errors.Wrap(err, errUpdateDevice)
		}
	*/
	updates, err := e.getGvkUpate(mg, obs, ygotnddp.NddpSystem_ResourceAction_UPDATE)
	if err != nil {
		return errors.Wrap(err, errCreateDevice)
	}

	crSystemDeviceName := shared.GetCrSystemDeviceName(shared.GetCrDeviceName(mg.GetNamespace(), mg.GetNetworkNodeReference().Name))

	req := &gnmi.SetRequest{
		Prefix:  &gnmi.Path{Target: crSystemDeviceName},
		Replace: updates,
	}

	_, err = e.client.Set(ctx, req)
	if err != nil {
		return errors.Wrap(err, errUpdateDevice)
	}

	return nil
}

func (e *externalDevice) Delete(ctx context.Context, mg resource.Managed, obs managed.ExternalObservation) error {
	log := e.log.WithValues("Resource", mg.GetName())
	log.Debug("Deleting ...", "obs", obs)

	/*
		cr, ok := mg.(*srlv1alpha1.Srl3Device)
		if !ok {
			return errors.New(errUnexpectedDevice)
		}

		crPaths, err := e.findPaths(cr.Spec.Device)
		if err != nil {
			return err
		}
	*/

	// create k8s object
	// processCreate
	// 0. marshal/unmarshal data
	// 1. transform the spec data to gnmi updates
	/*
		updates, err := e.processK8s(mg, crPaths, systemv1alpha1.E_GvkAction_Delete)
		if err != nil {
			return errors.Wrap(err, errDeleteDevice)
		}
	*/
	updates, err := e.getGvkUpate(mg, obs, ygotnddp.NddpSystem_ResourceAction_DELETE)
	if err != nil {
		return errors.Wrap(err, errCreateDevice)
	}

	crSystemDeviceName := shared.GetCrSystemDeviceName(shared.GetCrDeviceName(mg.GetNamespace(), mg.GetNetworkNodeReference().Name))

	req := &gnmi.SetRequest{
		Prefix:  &gnmi.Path{Target: crSystemDeviceName},
		Replace: updates,
	}

	_, err = e.client.Set(ctx, req)
	if err != nil {
		return errors.Wrap(err, errDeleteDevice)
	}

	return nil
}

func (e *externalDevice) GetResourceList(ctx context.Context, mg resource.Managed) (map[string]*ygotnddp.NddpSystem_Gvk, error) {
	// get system device list
	crSystemDeviceName := "ygot." + shared.GetCrSystemDeviceName(shared.GetCrDeviceName(mg.GetNamespace(), mg.GetNetworkNodeReference().Name))
	// gnmi get request
	reqSystemCache := &gnmi.GetRequest{
		Prefix:   &gnmi.Path{Target: crSystemDeviceName},
		Path:     []*gnmi.Path{{}},
		Encoding: gnmi.Encoding_JSON,
	}

	// gnmi get response
	resp, err := e.client.Get(ctx, reqSystemCache)
	if err != nil {
		return nil, err
	}
	var systemCache interface{}
	if len(resp.GetNotification()) == 0 {
		return nil, nil
	}
	if len(resp.GetNotification()) != 0 && len(resp.GetNotification()[0].GetUpdate()) != 0 {
		// get value from gnmi get response
		systemCache, err = yparser.GetValue(resp.GetNotification()[0].GetUpdate()[0].Val)
		if err != nil {
			return nil, errors.Wrap(err, errJSONMarshal)
		}

		switch systemCache.(type) {
		case nil:
			// resource has no data
			return nil, nil
		}
	}

	systemData, err := json.Marshal(systemCache)
	if err != nil {
		return nil, err
	}
	goStruct, err := e.systemModel.NewConfigStruct(systemData, true)
	if err != nil {
		return nil, err
	}
	nddpDevice, ok := goStruct.(*ygotnddp.Device)
	if !ok {
		return nil, errors.New("wrong object nddp")
	}

	return nddpDevice.Gvk, nil
}
