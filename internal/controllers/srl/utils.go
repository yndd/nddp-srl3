package srl

import (
	"encoding/json"
	"fmt"

	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/openconfig/ygot/ygot"
	"github.com/pkg/errors"
	"github.com/yndd/ndd-yang/pkg/yparser"
	"github.com/yndd/nddp-srl3/internal/rootpaths"
	"github.com/yndd/nddp-srl3/pkg/ygotsrl"
)

type observe struct {
	hasData  bool
	upToDate bool
	deletes  []*gnmi.Path
	updates  []*gnmi.Update
}

func (v *validatorDevice) getRootPaths(x *gnmi.Notification) ([]*gnmi.Path, error) {
	schema := v.deviceModel.SchemaTreeRoot
	rootConfigElement := rootpaths.ConfigElementHierarchyFromGnmiUpdate(schema, x)
	result := rootConfigElement.GetRootPaths()

	return result, nil
}

func (e *externalDevice) getGoStruct(x interface{}) (ygot.ValidatedGoStruct, error) {
	config, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return e.deviceModel.NewConfigStruct(config, false)
}

// 1. validate the repsonse to check if it contains the right # elements, data
func (e *externalDevice) processObserve(specData interface{}, resp *gnmi.GetResponse) (*observe, error) {
	// validate gnmi resp information
	if len(resp.GetNotification()) == 0 {
		return &observe{hasData: false}, nil
	}
	if len(resp.GetNotification()) > 1 {
		return &observe{hasData: false}, errors.New("processObserve invalid update length")
	}
	if len(resp.GetNotification()[0].GetUpdate()) == 0 || len(resp.GetNotification()[0].GetUpdate()) > 1 {
		return &observe{hasData: false}, errors.New("processObserve invalid update length")
	}

	// validate the response data from the cache
	x, err := yparser.GetValue(resp.GetNotification()[0].GetUpdate()[0].Val)
	if err != nil {
		return nil, errors.Wrap(err, errJSONMarshal)
	}
	switch x.(type) {
	case nil:
		// resource has no data
		return &observe{hasData: false}, nil
	}

	deletes, updates, err := e.goStructDiff(x, specData)
	if err != nil {
		return &observe{hasData: false}, err
	}
	return &observe{
		hasData:  true,
		upToDate: len(deletes) == 0 && len(updates) == 0,
		deletes:  deletes,
		updates:  updates,
	}, nil
}

func (e *externalDevice) goStructDiff(origin, spec interface{}) ([]*gnmi.Path, []*gnmi.Update, error) {
	srcConfig, err := e.getGoStruct(origin)
	if err != nil {
		return nil, nil, err
	}

	specConfig, err := e.getGoStruct(spec)
	if err != nil {
		return nil, nil, err
	}

	// skipping specValidation, this will probably result in missing leaf leafrefs
	srcConfigTmp, err := ygot.DeepCopy(srcConfig)
	if err != nil {
		return nil, nil, err
	}
	newConfig := srcConfigTmp.(*ygotsrl.Device) // Typecast

	// Merge spec into newconfig, which is right now jsut the actual config
	err = ygot.MergeStructInto(newConfig, specConfig)
	if err != nil {
		return nil, nil, err
	}

	// validate the new config
	//err = newConfig.Validate()
	//if err != nil {
	//	return &observe{hasData: false}, nil
	//}

	// create a diff of the actual compared to the to-become-new config
	actualVsSpecDiff, err := ygot.Diff(srcConfig, newConfig, &ygot.DiffPathOpt{MapToSinglePath: true})
	if err != nil {
		return nil, nil, err
	}

	return actualVsSpecDiff.GetDelete(), validateNotification(actualVsSpecDiff), nil
}

func validateNotification(n *gnmi.Notification) []*gnmi.Update {
	updates := make([]*gnmi.Update, 0)
	for _, u := range n.GetUpdate() {
		fmt.Printf("processObserve: diff update old path: %s, value: %v\n", yparser.GnmiPath2XPath(u.GetPath(), true), u.GetVal())
		// workaround since the diff can return double pathElem
		update := validateUpdate(u)
		fmt.Printf("processObserve: diff update new path: %s, value: %v\n", yparser.GnmiPath2XPath(update.GetPath(), true), update.GetVal())
		updates = append(updates, update)
	}
	return updates
}

// workaround for the dif handling
func validateUpdate(u *gnmi.Update) *gnmi.Update {
	if len(u.GetPath().GetElem()) <= 1 {
		return u
	}
	// when the 2nd last pathElem has a key and the last PathElem is an entry in the Key we should trim the last entry from the path
	// e.g. /interface[name=ethernet-1/49]/subinterface[index=1]/ipv4/address[ip-prefix=100.64.0.0/31]/ip-prefix, value: string_val:"100.64.0.0/31"
	// e.g. /interface[name=ethernet-1/49]/subinterface[index=1]/ipv4/address[ip-prefix=100.64.0.0/31]/ip-prefix, value: string_val:"100.64.0.0/31"
	if len(u.GetPath().GetElem()[len(u.GetPath().GetElem())-2].GetKey()) > 0 {
		if _, ok := u.GetPath().GetElem()[len(u.GetPath().GetElem())-2].GetKey()[u.GetPath().GetElem()[len(u.GetPath().GetElem())-1].GetName()]; ok {
			u.Path.Elem = u.Path.Elem[:len(u.GetPath().GetElem())-1]
			v := make(map[string]interface{})
			b, _ := json.Marshal(v)
			u.Val = &gnmi.TypedValue{Value: &gnmi.TypedValue_JsonVal{JsonVal: b}}
		}
	}

	return u
}
