package devicecollector

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/pkg/errors"
	"github.com/yndd/ndd-yang/pkg/yparser"
	srlv1alpha1 "github.com/yndd/nddp-srl3/apis/srl3/v1alpha1"
	"github.com/yndd/nddp-srl3/internal/shared"
	systemv1alpha1 "github.com/yndd/nddp-system/apis/system/v1alpha1"
	"github.com/yndd/nddp-system/pkg/gvkresource"
	"github.com/yndd/nddp-system/pkg/ygotnddp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
)

const (
	unmanagedResource = "Unmanaged resource"
)

var (
	pbRootPath     = &gnmi.Path{}
	pbDevicePrefix = &gnmi.Path{Elem: []*gnmi.PathElem{{Name: "device"}}}
)

func (c *collector) handleSubscription(resp *gnmi.SubscribeResponse) error {
	log := c.log.WithValues("target", c.target.Config.Name, "address", c.target.Config.Address)
	//log.Debug("handle target update from device")

	switch resp.GetResponse().(type) {
	case *gnmi.SubscribeResponse_Update:
		//log.Debug("handle target update from device", "Prefix", resp.GetUpdate().GetPrefix())

		// check if the target cache exists
		crDeviceName := shared.GetCrDeviceName(c.namespace, c.target.Config.Name)

		if !c.cache.HasTarget(crDeviceName) {
			log.Debug("handle target update target not found in ygot schema cache")
			return errors.New("target cache does not exist")
		}

		if !c.cache.GetCache().GetCache().HasTarget(crDeviceName) {
			log.Debug("handle target update target not found in cache")
			return errors.New("target cache does not exist")
		}

		resourceList, err := c.cache.GetSystemResourceList(shared.GetCrSystemDeviceName(crDeviceName))
		if err != nil {
			return err
		}
		/*
			resourceList, err := c.getResourceList(crDeviceName)
			if err != nil {
				return err
			}
		*/
		//log.Debug("resourceList", "list", resourceList)

		// check read/write maps

		// handle deletes
		if err := c.handleDeletes(crDeviceName, resourceList, resp.GetUpdate().Delete); err != nil {
			log.Debug("handleDeletes", "error", err)
			return err
		}

		if err := c.handleUpdates(crDeviceName, resourceList, resp.GetUpdate().Update); err != nil {
			log.Debug("handleUpdates", "error", err)
			return err
		}

	case *gnmi.SubscribeResponse_SyncResponse:
		//log.Debug("SyncResponse")
	}

	return nil
}

func (c *collector) handleDeletes(crDeviceName string, resourceList map[string]*ygotnddp.NddpSystem_Gvk, delPaths []*gnmi.Path) error {
	if len(delPaths) > 0 {
		/*
			for _, p := range delPaths {
				c.log.Debug("handleDeletes", "crDeviceName", crDeviceName, "path", yparser.GnmiPath2XPath(p, true))
			}
		*/
		// validate deletes
		if err := c.cache.ValidateDelete(crDeviceName, delPaths, true); err != nil {
			return err
		}

		// trigger reconcile event, but group them to avoid multiple reconciliation triggers
		resourceNames := map[string]string{}
		for _, path := range delPaths {
			xpath := yparser.GnmiPath2XPath(path, true)
			resourceName, err := c.findManagedResource(xpath, resourceList)
			if err != nil {
				return err
			}

			if *resourceName != unmanagedResource {
				resourceNames[*resourceName] = ""
			}
		}
		for resourceName := range resourceNames {
			c.triggerReconcileEvent(resourceName)
		}
	}
	return nil
}

func (c *collector) handleUpdates(crDeviceName string, resourceList map[string]*ygotnddp.NddpSystem_Gvk, updates []*gnmi.Update) error {
	if len(updates) > 0 {
		/*
			for _, u := range updates {
				c.log.Debug("handleUpdates", "path", yparser.GnmiPath2XPath(u.GetPath(), true), "val", u.GetVal())
			}
		*/

		// validate updates
		if err := c.cache.ValidateUpdate(crDeviceName, updates, false, true, true); err != nil {
			return err
		}

		// check of we need to trigger a reconcile event
		resourceNames := map[string]string{}
		for _, u := range updates {
			xpath := yparser.GnmiPath2XPath(u.GetPath(), true)
			// check if this is a managed resource or unmanged resource
			// name == unmanagedResource is an unmanaged resource
			resourceName, err := c.findManagedResource(xpath, resourceList)
			if err != nil {
				return err
			}
			if *resourceName != unmanagedResource {
				resourceNames[*resourceName] = ""
			}

			/*
				walkInternalGoStruct := true
				if !walkInternalGoStruct {
					// we create a json blob and merge this in the gostruct
					fullPath := u.GetPath()
					val := u.GetVal()
					json, err := generateJson(u)
					if err != nil {
						c.log.Debug("generate json error", "err", err)
						return err
					}
					c.log.Debug("subscription config update", "path", yparser.GnmiPath2XPath(fullPath, true), "val", val)
					//c.log.Debug("subscription config update", "path", yparser.GnmiPath2XPath(fullPath, true), "json", string(json))
					// create newGostruct which will not be validated
					newGoStruct, err := m.NewConfigStruct(json, false)
					if err != nil {
						c.log.Debug("generate new gostruct", "err", err)
						return err
					}
					// merge the newGostruct with the current config
					currGoStruct := c.cache.GetValidatedGoStruct(crDeviceName)
					if err := ygot.MergeStructInto(currGoStruct, newGoStruct); err != nil {
						c.log.Debug("merge  gostructs", "err", err)
						return err
					}
					// validate the merged config
					if err := currGoStruct.Validate(); err != nil {
						c.log.Debug("validate new gostructs", "error", err)
						return err
					}
					// since validation is successfull we can set the newGostruct as the new valid config
					c.cache.UpdateValidatedGoStruct(crDeviceName, newGoStruct)
				}
				}*/
		}
		for resourceName := range resourceNames {
			c.triggerReconcileEvent(resourceName)
		}
	}
	return nil
}

func (c *collector) findManagedResource(xpath string, resourceList map[string]*ygotnddp.NddpSystem_Gvk) (*string, error) {
	matchedResourceName := unmanagedResource
	matchedResourcePath := ""
	for resourceName, r := range resourceList {
		for _, path := range r.Path {
			if strings.Contains(xpath, path) {
				// if there is a better match we use the better match
				if len(path) > len(matchedResourcePath) {
					matchedResourcePath = path
					matchedResourceName = resourceName
				}
			}
		}
	}
	return &matchedResourceName, nil
}

func (c *collector) getResourceList(crDeviceName string) ([]*systemv1alpha1.Gvk, error) {
	crSystemDeviceName := shared.GetCrSystemDeviceName(crDeviceName)

	rl, err := c.cache.GetCache().GetJson(crSystemDeviceName,
		&gnmi.Path{Target: crSystemDeviceName},
		&gnmi.Path{Elem: []*gnmi.PathElem{{Name: "gvk"}}},
		c.nddpSchema)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return gvkresource.GetResourceList(rl)
}

func (c *collector) triggerReconcileEvent(resourceName string) error {

	gvk, err := gvkresource.String2Gvk(resourceName)
	if err != nil {
		return err
	}
	kindgroup := strings.Join([]string{gvk.Kind, gvk.Group}, ".")

	object := getObject(gvk)

	//c.log.Debug("triggerReconcileEvent", "kindgroup", kindgroup, "gvk", gvk, "object", object)

	if eventCh, ok := c.eventChs[kindgroup]; ok {
		c.log.Debug("triggerReconcileEvent with channel lookup", "kindgroup", kindgroup, "gvk", gvk, "object", object)
		eventCh <- event.GenericEvent{
			Object: object,
		}
	}
	return nil
}

func getObject(gvk *gvkresource.GVK) client.Object {
	switch gvk.Kind {
	case "Srl3Device":
		return &srlv1alpha1.Srl3Device{
			ObjectMeta: metav1.ObjectMeta{Name: gvk.Name, Namespace: gvk.NameSpace},
		}
	default:
		fmt.Printf("getObject not found gvk: %v\n", *gvk)
		return nil
	}
}

func generateJson(u *gnmi.Update) ([]byte, error) {
	var err error
	var d interface{}
	path := yparser.DeepCopyGnmiPath(u.GetPath())
	if d, err = addData(d, path.GetElem(), u.GetVal()); err != nil {
		return nil, err
	}
	return json.Marshal(d)

}

func addData(d interface{}, pe []*gnmi.PathElem, val *gnmi.TypedValue) (interface{}, error) {
	var err error
	if len(pe) == 0 {
		return nil, nil
	}
	e := pe[0].GetName()
	k := pe[0].GetKey()
	//fmt.Printf("addData, Len: %d, Elem: %s, Key: %v, Data: %v\n", len(elems), e, k, d)
	if len(pe)-1 == 0 {
		// last element
		if len(k) == 0 {
			// last element with container
			d, err = addContainerValue(d, e, val)
			return d, err
		} else {
			// last element with list
			// not sure if this will ever exist
			d, err = addListValue(d, e, k, val)
			return d, err
		}
	} else {
		if len(k) == 0 {
			// not last element -> container
			d, err = addContainer(d, e, pe, val)
			return d, err
		} else {
			// not last element -> list + keys
			d, err = addList(d, e, k, pe, val)
			return d, err
		}
	}
}

func addContainer(d interface{}, e string, elems []*gnmi.PathElem, val *gnmi.TypedValue) (interface{}, error) {
	var err error
	// initialize the data
	//fmt.Printf("addContainer: %v pathElem: %s val: %v\n", elems, e, val)
	if reflect.TypeOf((d)) == nil {
		d = make(map[string]interface{})
	}
	switch dd := d.(type) {
	case map[string]interface{}:
		// add the container
		dd[e], err = addData(dd[e], elems[1:], val)
		return d, err
	default:
		return nil, errors.New("addListLastValue JSON unexpected data structure")
	}
	//}

}

func addList(d interface{}, e string, k map[string]string, elems []*gnmi.PathElem, val *gnmi.TypedValue) (interface{}, error) {
	var err error
	//fmt.Printf("addList pathElem: %s, key: %v d: %v\n", e, k, d)
	// lean approach -> since we know the query should return paths that match the original query we can assume we match the path

	// initialize the data
	if reflect.TypeOf((d)) == nil {
		d = make(map[string]interface{})
	}
	switch dd := d.(type) {
	case map[string]interface{}:
		// initialize the data
		if _, ok := dd[e]; !ok {
			dd[e] = make([]interface{}, 0)
		}
		switch l := dd[e].(type) {
		case []interface{}:
			// check if the list entry exists
			for i, le := range l {
				// initialize the data
				if reflect.TypeOf((le)) == nil {
					le = make(map[string]interface{})
				}
				found := true
				switch dd := le.(type) {
				case map[string]interface{}:
					for keyName, keyValue := range k {
						if dd[keyName] != keyValue {
							found = false
						}
					}
					if found {
						// augment the list
						l[i], err = addData(dd, elems[1:], val)
						if err != nil {
							return nil, err
						}
						return d, err
					}
				}
			}
			// list entry not found, add a list entry
			de := make(map[string]interface{})
			for keyName, keyValue := range k {
				de[keyName] = keyValue
			}
			// augment the list
			x, err := addData(de, elems[1:], val)
			if err != nil {
				return nil, err
			}
			// add the list entry to the list
			dd[e] = append(l, x)
			return d, nil
		default:
			return nil, errors.New("list last value JSON unexpected data structure")
		}
	default:
		return nil, errors.New("list last value JSON unexpected data structure")
	}
}

func addContainerValue(d interface{}, e string, val *gnmi.TypedValue) (interface{}, error) {
	var err error
	// check if the data was initialized
	if reflect.TypeOf((d)) == nil {
		d = make(map[string]interface{})
	}
	switch dd := d.(type) {
	case map[string]interface{}:
		// add the value to the element
		dd[e], err = yparser.GetValue(val)
		return d, err
	default:
		// we should never end up here
		return nil, errors.New("container last value JSON unexpected data structure")
	}
}

func addListValue(d interface{}, e string, k map[string]string, val *gnmi.TypedValue) (interface{}, error) {
	// initialize the data
	if reflect.TypeOf((d)) == nil {
		d = make(map[string]interface{})
	}
	switch dd := d.(type) {
	case map[string]interface{}:
		// initialize the data
		if _, ok := dd[e]; !ok {
			dd[e] = make([]interface{}, 0)
		}
		// create a container and initialize with keyNames/keyValues and value
		de := make(map[string]interface{})
		// add value
		v, err := yparser.GetValue(val)
		if err != nil {
			return nil, err
		}
		switch vv := v.(type) {
		case map[string]interface{}:
			for k, v := range vv {
				de[k] = v
			}
		default:
		}

		// add keyNames/keyValues
		for keyName, keyValue := range k {
			de[keyName] = keyValue
		}
		// add the data to the list
		switch l := dd[e].(type) {
		case []interface{}:
			dd[e] = append(l, de)
		default:
			return nil, errors.New("list last value JSON unexpected data structure")
		}
	}
	return d, nil
}

func cleanPath(path *gnmi.Path) *gnmi.Path {
	// clean the path for now to remove the module information from the pathElem
	p := yparser.DeepCopyGnmiPath(path)
	for _, pe := range p.GetElem() {
		pe.Name = strings.Split(pe.Name, ":")[len(strings.Split(pe.Name, ":"))-1]
		keys := make(map[string]string)
		for k, v := range pe.GetKey() {
			if strings.Contains(v, "::") {
				keys[strings.Split(k, ":")[len(strings.Split(k, ":"))-1]] = v
			} else {
				keys[strings.Split(k, ":")[len(strings.Split(k, ":"))-1]] = strings.Split(v, ":")[len(strings.Split(v, ":"))-1]
			}
		}
		pe.Key = keys
	}
	return p
}
