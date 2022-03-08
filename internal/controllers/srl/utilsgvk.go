package srl

import (
	"encoding/json"

	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/openconfig/ygot/ygot"
	"github.com/pkg/errors"
	"github.com/yndd/ndd-runtime/pkg/reconciler/managed"
	"github.com/yndd/ndd-runtime/pkg/resource"
	"github.com/yndd/ndd-runtime/pkg/utils"
	"github.com/yndd/ndd-yang/pkg/yparser"
	srlv1alpha1 "github.com/yndd/nddp-srl3/apis/srl3/v1alpha1"
	"github.com/yndd/nddp-system/pkg/gvkresource"
	"github.com/yndd/nddp-system/pkg/ygotnddp"
)

func (e *externalDevice) getGvkUpate(mg resource.Managed, obs managed.ExternalObservation, action ygotnddp.E_NddpSystem_ResourceAction) ([]*gnmi.Update, error) {
	e.log.Debug("getGvkUpate", "paths", obs.Paths)

	// get gvk Name
	gvkName := gvkresource.GetGvkName(mg)

	// get spec in string format
	spec, err := getSpec(mg)
	if err != nil {
		return nil, err
	}

	updates, err := getUpdates(obs.Updates)
	if err != nil {
		return nil, err
	}

	// get nddpData from gvkname, action, paths and spec

	/*
		nddpData := &ygotnddp.Device{
			Gvk: map[string]*ygotnddp.NddpSystem_Gvk{
				gvkName: {
					Name:   ygot.String(gvkName),
					Action: action,
					Path:   getPaths(crPaths),
					Status: getStatus(action),
					Reason: ygot.String(""),
					Spec:   spec,
					Delete: getPaths(crDeletes),
					Update: updates,
				},
			},
		}
	*/

	gvkData := &ygotnddp.NddpSystem_Gvk{
		Name:   ygot.String(gvkName),
		Action: action,
		Path:   getPaths(obs.Paths),
		Status: getStatus(action),
		Reason: ygot.String(""),
		Spec:   spec,
		Delete: getPaths(obs.Deletes),
		Update: updates,
	}

	/*
		gvkJson, err := json.Marshal(gvkData)
		if err != nil {
			return nil, err
		}
	*/

	nddpJson, err := ygot.EmitJSON(gvkData, &ygot.EmitJSONConfig{
		Format: ygot.RFC7951,
	})

	//return update
	return []*gnmi.Update{
		{
			Path: &gnmi.Path{
				Elem: []*gnmi.PathElem{
					{Name: "gvk", Key: map[string]string{"name": gvkName}},
				},
			},
			Val: &gnmi.TypedValue{Value: &gnmi.TypedValue_JsonVal{JsonVal: []byte(nddpJson)}},
		},
	}, nil
}

func getSpec(mg resource.Managed) (*string, error) {
	cr, ok := mg.(*srlv1alpha1.Srl3Device)
	if !ok {
		return nil, errors.New(errUnexpectedDevice)
	}
	spec, err := json.Marshal(cr.Spec.Device)
	if err != nil {
		return nil, errors.Wrap(err, errJSONMarshal)
	}
	return utils.StringPtr(string(spec)), nil
}

func getPaths(gnmiPaths []*gnmi.Path) []string {
	paths := []string{}
	for _, p := range gnmiPaths {
		paths = append(paths, yparser.GnmiPath2XPath(p, true))
	}
	return paths
}

func getUpdates(gnmiUpdates []*gnmi.Update) (map[string]*ygotnddp.NddpSystem_Gvk_Update, error) {
	updates := map[string]*ygotnddp.NddpSystem_Gvk_Update{}
	for _, u := range gnmiUpdates {
		xpath := yparser.GnmiPath2XPath(u.GetPath(), true)
		v, err := yparser.GetValue(u.GetVal())
		if err != nil {
			return nil, err
		}
		val, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		updates[xpath] = &ygotnddp.NddpSystem_Gvk_Update{
			Path: ygot.String(xpath),
			Val:  ygot.String(string(val)),
		}
	}
	return updates, nil
}

func getStatus(action ygotnddp.E_NddpSystem_ResourceAction) ygotnddp.E_NddpSystem_ResourceStatus {
	switch action {
	case ygotnddp.NddpSystem_ResourceAction_CREATE:
		return ygotnddp.NddpSystem_ResourceStatus_CREATEPENDING
	case ygotnddp.NddpSystem_ResourceAction_DELETE:
		return ygotnddp.NddpSystem_ResourceStatus_DELETEPENDING
	case ygotnddp.NddpSystem_ResourceAction_UPDATE:
		return ygotnddp.NddpSystem_ResourceStatus_UPDATEPENDING
	}
	return ygotnddp.NddpSystem_ResourceStatus_None
}
