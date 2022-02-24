package gnmi

import (
	"sync"

	"github.com/openconfig/ygot/ygot"

	pb "github.com/openconfig/gnmi/proto/gnmi"
)

type ConfigCallback func(ygot.ValidatedGoStruct) error

var (
	pbRootPath         = &pb.Path{}
	supportedEncodings = []pb.Encoding{pb.Encoding_JSON, pb.Encoding_JSON_IETF}
)

type Server struct {
	pb.UnimplementedGNMIServer

	model    *Model
	callback ConfigCallback

	config ygot.ValidatedGoStruct
	mu     sync.RWMutex // mu is the RW lock to protect the access to config
}

// NewServer creates an instance of Server with given json config.
func NewServer(model *Model, config []byte, callback ConfigCallback) (*Server, error) {
	rootStruct, err := model.NewConfigStruct(config)
	if err != nil {
		return nil, err
	}
	s := &Server{
		model:    model,
		config:   rootStruct,
		callback: callback,
	}
	if config != nil && s.callback != nil {
		if err := s.callback(rootStruct); err != nil {
			return nil, err
		}
	}
	return s, nil
}
