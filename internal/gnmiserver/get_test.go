package gnmiserver

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/yndd/nddp-srl3/internal/tests/mocks"
)

func Test_handleGet(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	gnmiServer := getServer(mockCtrl)

	request := composeGnmiRequest(
		[]*gnmi.Path{
			{Elem: []*gnmi.PathElem{{Name: "interface", Key: map[string]string{"name": "*"}}}},
			{Elem: []*gnmi.PathElem{}},
		},
		gnmi.Encoding_JSON_IETF,
		"ygot.system",
	)

	gnmiServer.HandleGet(request)

}

func composePath(elems []*gnmi.PathElem) *gnmi.Path {
	return &gnmi.Path{Elem: elems}
}

func getServer(mockCtrl *gomock.Controller) *GnmiServerImpl {
	result := &GnmiServerImpl{
		cache: mocks.NewMockCache(mockCtrl),
	}
	return result
}

func composeGnmiRequest(paths []*gnmi.Path, encoding gnmi.Encoding, target string) *gnmi.GetRequest {
	request := &gnmi.GetRequest{
		Prefix:   &gnmi.Path{Target: target},
		Path:     paths,
		Encoding: encoding,
	}
	return request
}
