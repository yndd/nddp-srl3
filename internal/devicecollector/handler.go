package devicecollector

import (
	"fmt"
	"strings"

	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/pkg/errors"
	"github.com/yndd/ndd-yang/pkg/yparser"
	srlv1alpha1 "github.com/yndd/nddp-srl3/apis/srl3/v1alpha1"
	"github.com/yndd/nddp-srl3/internal/cache"
	"github.com/yndd/nddp-srl3/internal/shared"
	"github.com/yndd/nddp-system/pkg/gvkresource"
	"github.com/yndd/nddp-system/pkg/ygotnddp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
)

const (
	unmanagedResource = "Unmanaged resource"
)

/*
var (
	pbRootPath     = &gnmi.Path{}
	pbDevicePrefix = &gnmi.Path{Elem: []*gnmi.PathElem{{Name: "device"}}}
)
*/

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

		// lists the k8s resources that should be triggered for a reconcile event since an update happend
		// the the config belonging to the CR
		// this is used to validate if the config is up to date or not
		resourceNames := map[string]struct{}{}
		// handle deletes
		if err := c.handleDeletes(crDeviceName, resourceNames, resourceList, resp.GetUpdate().Delete); err != nil {
			log.Debug("handleDeletes", "error", err)
			return err
		}

		if err := c.handleUpdates(crDeviceName, resourceNames, resourceList, resp.GetUpdate().Update); err != nil {
			log.Debug("handleUpdates", "error", err)
			return err
		}
		for resourceName := range resourceNames {
			c.triggerReconcileEvent(resourceName)
		}

	case *gnmi.SubscribeResponse_SyncResponse:
		//log.Debug("SyncResponse")
	}

	return nil
}

// handleDeletes updates the running config to align the device cache based on the delete information.
// A reconcile event is triggered to the k8s controller if the delete path matches a managed k8s resource (MR)
func (c *collector) handleDeletes(crDeviceName string, resourceNames map[string]struct{}, resourceList map[string]*ygotnddp.NddpSystem_Gvk, delPaths []*gnmi.Path) error {
	if len(delPaths) > 0 {
		/*
			for _, p := range delPaths {
				c.log.Debug("handleDeletes", "crDeviceName", crDeviceName, "path", yparser.GnmiPath2XPath(p, true))
			}
		*/
		// validate deletes
		if err := c.cache.ValidateDelete(crDeviceName, delPaths, cache.Origin_Subscription); err != nil {
			return err
		}

		// trigger reconcile event, but group them to avoid multiple reconciliation triggers

		for _, path := range delPaths {
			xpath := yparser.GnmiPath2XPath(path, true)
			resourceName, err := c.findManagedResource(xpath, resourceList)
			if err != nil {
				return err
			}

			if *resourceName != unmanagedResource {
				resourceNames[*resourceName] = struct{}{}
			}
		}
	}
	return nil
}

// handleUpdates updates the running config to align the device cache based on the update information.
// A reconcile event is triggered to the k8s controller if the update path matches a managed k8s resource (MR)
func (c *collector) handleUpdates(crDeviceName string, resourceNames map[string]struct{}, resourceList map[string]*ygotnddp.NddpSystem_Gvk, updates []*gnmi.Update) error {
	if len(updates) > 0 {
		/*
			for _, u := range updates {
				c.log.Debug("handleUpdates", "path", yparser.GnmiPath2XPath(u.GetPath(), true), "val", u.GetVal())
			}
		*/

		// validate updates
		if err := c.cache.ValidateUpdate(crDeviceName, updates, false, true, cache.Origin_Subscription); err != nil {
			return err
		}

		// check of we need to trigger a reconcile event
		for _, u := range updates {
			xpath := yparser.GnmiPath2XPath(u.GetPath(), true)
			// check if this is a managed resource or unmanged resource
			// name == unmanagedResource is an unmanaged resource
			resourceName, err := c.findManagedResource(xpath, resourceList)
			if err != nil {
				return err
			}
			if *resourceName != unmanagedResource {
				resourceNames[*resourceName] = struct{}{}
			}
		}
	}
	return nil
}

// findManagedResource returns a the k8s resourceName as a string (using gvk convention [group, version, kind, namespace, name])
// by validation the best path match in the resourcelist of the system cache
// if no match is find unmanagedResource is returned, since this path is not managed by the k8s controller
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

// triggerReconcileEvent triggers an external event to the k8s controller with the object resource reference
// this should result in a reconcile trigger on the k8s controller for the particular cr.
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

// getObject returns the k8s object based on the gvk resource name (group, version, kind, namespace, name)
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
