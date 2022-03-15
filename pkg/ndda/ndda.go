package ndda

import (
	"context"
	"reflect"

	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/yndd/ndd-runtime/pkg/logging"
	"github.com/yndd/nddo-runtime/pkg/resource"
	srlv1alpha1 "github.com/yndd/nddp-srl3/apis/srl3/v1alpha1"
	"github.com/yndd/nddp-srl3/internal/model"
	"github.com/yndd/nddp-srl3/pkg/ygotsrl"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func New(opts ...Option) Handler {
	s := &handler{
		newSrlDeviceList: func() srlv1alpha1.IFSrl3DeviceList { return &srlv1alpha1.Srl3DeviceList{} },
		m: &model.Model{
			ModelData:       make([]*gnmi.ModelData, 0),
			StructRootType:  reflect.TypeOf((*ygotsrl.Device)(nil)),
			SchemaTreeRoot:  ygotsrl.SchemaTree["Device"],
			JsonUnmarshaler: ygotsrl.Unmarshal,
			EnumData:        ygotsrl.ΛEnum,
		},
		ctx: context.Background(),
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (r *handler) WithLogger(log logging.Logger) {
	r.log = log
}

func (r *handler) WithClient(c client.Client) {
	r.client = resource.ClientApplicator{
		Client:     c,
		Applicator: resource.NewAPIPatchingApplicator(c),
	}
}

type handler struct {
	log logging.Logger
	// kubernetes
	client client.Client
	ctx    context.Context

	m *model.Model

	newSrlDeviceList func() srlv1alpha1.IFSrl3DeviceList
}
