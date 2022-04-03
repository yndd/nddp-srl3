package gnmiserver

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/antchfx/jsonquery"
	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/openconfig/ygot/ygot"
	"github.com/yndd/nddp-srl3/internal/model"
	"github.com/yndd/nddp-srl3/pkg/ygotsrl"
)

// Type definition for checks on the populateNotification() call results
type populateNotificationChecker func([]*gnmi.Notification, error) error

// Type definition for populateNotification Test input + resultcheck
type populateNotificationTestInput struct {
	description string
	gostruct    ygot.GoStruct
	req         *gnmi.GetRequest
	model       *model.Model
	ts          int64
	prefix      *gnmi.Path
	resultCheck []populateNotificationChecker
}

// Test_populateNotification test the populateNotification function
func Test_populateNotification(t *testing.T) {
	// -------------------------------------
	// define generic parameters
	devmodel := &model.Model{
		StructRootType:  reflect.TypeOf((*ygotsrl.Device)(nil)),
		SchemaTreeRoot:  ygotsrl.SchemaTree["Device"],
		JsonUnmarshaler: ygotsrl.Unmarshal,
		EnumData:        ygotsrl.ΛEnum,
	}
	ts := time.Now().UnixMilli()
	prefix := &gnmi.Path{}

	// -------------------------------------
	// create the test matrix
	tests := []populateNotificationTestInput{
		{
			description: "Test wildcard query with explicit * wildcard",
			gostruct:    SetupDevice1(),
			req: SetupGnmiGetReq(
				[]*gnmi.Path{
					{
						Elem: []*gnmi.PathElem{
							{
								Name: "interface",
								Key:  map[string]string{"name": "*"},
							},
						},
					},
				},
				&gnmi.Path{},
			),
			model:  devmodel,
			ts:     ts,
			prefix: prefix,
			resultCheck: []populateNotificationChecker{
				populateNotificationPrint(),
				populateNotificationCheckError(false),
				populateNotificationCheckNumUpdates(1),
				populateNotificationCheckNumElementsInUpdate(1, "/interface"),
				populateNotificationCheckNumElementsInUpdate(5, "/interface/*"),
				populateNotificationCheckNumElementsInUpdate(0, "/system"),
			},
		},
		{
			description: "Test wildcard query without explicit * wildcard",
			gostruct:    SetupDevice1(),
			req: SetupGnmiGetReq(
				[]*gnmi.Path{
					{
						Elem: []*gnmi.PathElem{
							{
								Name: "interface",
							},
						},
					},
				},
				&gnmi.Path{},
			),
			model:  devmodel,
			ts:     ts,
			prefix: prefix,
			resultCheck: []populateNotificationChecker{
				populateNotificationCheckError(false),
				populateNotificationCheckNumUpdates(1),
				populateNotificationCheckNumElementsInUpdate(1, "/interface"),
				populateNotificationCheckNumElementsInUpdate(5, "/interface/*"),
				populateNotificationCheckNumElementsInUpdate(0, "/system"),
			},
		},
		{
			description: "Test wildcard query without multiple explicit wildcard",
			gostruct:    SetupDevice1(),
			req: SetupGnmiGetReq(
				[]*gnmi.Path{
					{
						Elem: []*gnmi.PathElem{
							{
								Name: "interface",
								Key:  map[string]string{"name": "*"},
							},
							{
								Name: "subinterface",
								Key:  map[string]string{"index": "*"},
							},
							{
								Name: "ipv4",
							},
						},
					},
				},
				&gnmi.Path{},
			),
			model:  devmodel,
			ts:     ts,
			prefix: prefix,
			resultCheck: []populateNotificationChecker{
				populateNotificationCheckError(false),
				populateNotificationCheckNumUpdates(1),
				populateNotificationCheckNumElementsInUpdate(7, "//interface//subinterface//ipv4"),
			},
		},
		{
			description: "Check without any path elements",
			gostruct:    SetupDevice1(),
			req: SetupGnmiGetReq(
				[]*gnmi.Path{
					{
						Elem: []*gnmi.PathElem{},
					},
				},
				&gnmi.Path{},
			),
			model:  devmodel,
			ts:     ts,
			prefix: prefix,
			resultCheck: []populateNotificationChecker{
				populateNotificationCheckError(false),
				populateNotificationCheckNumUpdates(1),
			},
		},
	}

	// -------------------------------------
	// execute tests
	for testNo, test := range tests {
		n, err := populateNotification(test.gostruct, test.req, test.model, test.ts, test.prefix)

		// execute all the defined checks
		for _, rc := range test.resultCheck {
			checkerror := rc(n, err)
			// if the check returns an error throw it as the test result, making the test fail
			if checkerror != nil {
				t.Errorf("%s, [Test: %d -> %s]", checkerror, testNo+1, test.description)
			}
		}
	}
}

// populateNotificationPrint used to debug tests, prints the result
func populateNotificationPrint() populateNotificationChecker {
	return func(notis []*gnmi.Notification, err error) error {
		if err != nil {
			fmt.Println("Error: ", err)
		}
		for _, n := range notis {
			fmt.Println("\t" + n.String())
		}
		return nil
	}
}

// populateNotificationCheckError Checks that an is raised shouldError == true or specifically not raised shoudlError == false
func populateNotificationCheckError(shouldError bool) populateNotificationChecker {
	return func(notis []*gnmi.Notification, err error) error {
		if shouldError {
			if err == nil {
				return fmt.Errorf("An error was expected, got error == nil result.")
			} else {
				return nil
			}
		}
		// we did not expect an error, so hand back err
		return err
	}
}

// populateNotificationCheckNumUpdates check that the given amount of updates is
// contained in the notis []*gnmi.Notification
func populateNotificationCheckNumUpdates(x int) populateNotificationChecker {
	return func(notis []*gnmi.Notification, err error) error {
		sum := 0
		for _, n := range notis {
			sum = sum + len(n.Update)
		}
		if sum != x {
			return fmt.Errorf("Expected %d updates, found %d", x, sum)
		}
		return nil
	}
}

// populateNotificationCheckNumElementsInUpdate queries the first notifications first update, applies the given json query
// and compares the length of the result with the given count
func populateNotificationCheckNumElementsInUpdate(count int, path string) populateNotificationChecker {
	return func(notis []*gnmi.Notification, err error) error {

		r := bytes.NewReader(notis[0].Update[0].GetVal().GetJsonVal())
		doc, err := jsonquery.Parse(r)
		if err != nil {
			return fmt.Errorf("CHECK ERROR: %s", err)
		}

		list := jsonquery.Find(doc, path)
		if len(list) != count {
			return fmt.Errorf("Expected %d items to be returned, got %d", count, len(list))
		}
		return nil
	}
}

// SetupGnmiGetReq helper to setup a GnmiGetRequest
func SetupGnmiGetReq(paths []*gnmi.Path, prefix *gnmi.Path) *gnmi.GetRequest {
	result := &gnmi.GetRequest{
		Prefix: prefix,
		Path:   paths,
	}
	return result
}

// SetupDevice1 helper to create a device
func SetupDevice1() ygot.GoStruct {
	d := &ygotsrl.Device{}

	// v4 only interface
	e11 := d.GetOrCreateInterface("ethernet-1/1")
	e11s5 := e11.GetOrCreateSubinterface(5)
	e11s5.GetOrCreateIpv4().NewAddress("192.168.0.0/24")

	// v4 and v6 on different subinterfaces
	e12 := d.GetOrCreateInterface("ethernet-1/2")
	e12s5 := e12.GetOrCreateSubinterface(5)
	e12s5.GetOrCreateIpv4().NewAddress("192.168.1.0/24")
	e12s20 := e12.GetOrCreateSubinterface(20)
	e12s20.GetOrCreateIpv6().NewAddress("2001:DB8::1/128")

	// v4 only interface single subinterface
	e13 := d.GetOrCreateInterface("ethernet-1/3")
	e13s5 := e13.GetOrCreateSubinterface(5)
	e13s5.GetOrCreateIpv4().NewAddress("192.168.2.0/24")

	// v4 only interface multiple ipv4
	e14 := d.GetOrCreateInterface("ethernet-1/4")
	e14s5 := e14.GetOrCreateSubinterface(5)
	e14s5.GetOrCreateIpv4().NewAddress("192.168.3.0/24")
	e14s6 := e14.GetOrCreateSubinterface(6)
	e14s6.GetOrCreateIpv4().NewAddress("192.168.4.0/24")
	e14s7 := e14.GetOrCreateSubinterface(7)
	e14s7.GetOrCreateIpv4().NewAddress("192.168.5.0/24")

	// v4 and v6 mixed on same subinterface
	e15 := d.GetOrCreateInterface("ethernet-1/5")
	e15s5 := e15.GetOrCreateSubinterface(5)
	e15s5.GetOrCreateIpv4().NewAddress("192.168.3.0/24")
	e15s5.GetOrCreateIpv6().NewAddress("2001:DB8::2/128")

	system := d.GetOrCreateSystem()

	banner := system.GetOrCreateBanner()
	banner.LoginBanner = ygot.String("Hello World this is the Login Banner.")
	banner.MotdBanner = ygot.String("Hello World this is the Motd Banner.")

	ni1 := d.GetOrCreateNetworkInstance("MyTinyNwInstance")
	ni1.AdminState = ygotsrl.SrlNokiaCommon_AdminState_enable
	ni1.Description = ygot.String("Just a tiny NetworkInstance")

	return d
}
