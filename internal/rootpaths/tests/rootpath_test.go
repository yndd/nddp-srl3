package rootpaths_tests

import (
	"testing"
	"time"

	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/openconfig/goyang/pkg/yang"
	"github.com/openconfig/ygot/ygot"
	"github.com/yndd/nddp-srl3/internal/rootpaths"
	"github.com/yndd/nddp-srl3/pkg/ygotsrl"
)

// Test_rootpathsHighLevel testing rootpath calculation. taking configs and expected string representations of the rootpaths
func Test_rootpathsHighLevel(t *testing.T) {

	var testTabe = []struct {
		config    string
		rootpaths []string
	}{
		{sampleDataIf, []string{
			"elem:{name:\"interface\" key:{key:\"name\" value:\"ethernet-1/48\"}} elem:{name:\"subinterface\" key:{key:\"index\" value:\"100\"}}",
		}},
		{sampleDataLarger, []string{
			"elem:{name:\"network-instance\" key:{key:\"name\" value:\"multus-sriov2-bridged\"}}",
			"elem:{name:\"interface\" key:{key:\"name\" value:\"irb0\"}} elem:{name:\"subinterface\" key:{key:\"index\" value:\"1386\"}} elem:{name:\"anycast-gw\"}",
			"elem:{name:\"interface\" key:{key:\"name\" value:\"irb0\"}} elem:{name:\"subinterface\" key:{key:\"index\" value:\"1386\"}} elem:{name:\"ipv4\"}",
			"elem:{name:\"interface\" key:{key:\"name\" value:\"irb0\"}} elem:{name:\"subinterface\" key:{key:\"index\" value:\"1386\"}} elem:{name:\"ipv6\"} elem:{name:\"address\" key:{key:\"ip-prefix\" value:\"2a02:1800:83:7000::1/64\"}}",
			"elem:{name:\"interface\" key:{key:\"name\" value:\"irb0\"}} elem:{name:\"subinterface\" key:{key:\"index\" value:\"1386\"}} elem:{name:\"ipv6\"} elem:{name:\"neighbor-discovery\"}",
			"elem:{name:\"interface\" key:{key:\"name\" value:\"lag1\"}} elem:{name:\"subinterface\" key:{key:\"index\" value:\"200\"}}",
			"elem:{name:\"interface\" key:{key:\"name\" value:\"lag1\"}} elem:{name:\"subinterface\" key:{key:\"index\" value:\"100\"}}",
			"elem:{name:\"interface\" key:{key:\"name\" value:\"lag1\"}} elem:{name:\"subinterface\" key:{key:\"index\" value:\"102\"}}",
			"elem:{name:\"network-instance\" key:{key:\"name\" value:\"multus-ipvlan-bridged\"}}",
			"elem:{name:\"network-instance\" key:{key:\"name\" value:\"multus-routed\"}}",
			"elem:{name:\"tunnel-interface\" key:{key:\"name\" value:\"vxlan0\"}} elem:{name:\"vxlan-interface\" key:{key:\"index\" value:\"2106\"}}",
			"elem:{name:\"tunnel-interface\" key:{key:\"name\" value:\"vxlan0\"}} elem:{name:\"vxlan-interface\" key:{key:\"index\" value:\"1386\"}}",
			"elem:{name:\"tunnel-interface\" key:{key:\"name\" value:\"vxlan0\"}} elem:{name:\"vxlan-interface\" key:{key:\"index\" value:\"2105\"}}",
			"elem:{name:\"tunnel-interface\" key:{key:\"name\" value:\"vxlan0\"}} elem:{name:\"vxlan-interface\" key:{key:\"index\" value:\"2143\"}}",
			"elem:{name:\"interface\" key:{key:\"name\" value:\"ethernet-1/48\"}} elem:{name:\"subinterface\" key:{key:\"index\" value:\"100\"}}",
			"elem:{name:\"network-instance\" key:{key:\"name\" value:\"multus-sriov1-bridged\"}}",
		}},
	}

	for _, entry := range testTabe {
		// load config from variable
		rootSchema, notification, err := loadConfigIntoGnmiNotification(entry.config)
		if err != nil {
			t.Error(err)
		}
		// build hierarchy
		rce := rootpaths.ConfigElementHierarchyFromGnmiUpdate(rootSchema, notification)

		// get rootpaths
		paths := rce.GetRootPaths()

		// iterate over calculated paths making sure they can be found in the expected test results
		for _, path := range paths {
			if !pathIn(path.String(), entry.rootpaths) {
				t.Errorf("entry %s could not be found in the expected results.", path.String())
			}
		}
		// iterate over the expected results making sure they can all be found in the actual results
		for _, path := range entry.rootpaths {
			found := false
			for _, calcPath := range paths {
				if calcPath.String() == path {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("expected entry %s in results but couldn't find it.", path)
			}
		}
	}
}

// pathIn checks wether the testString is contained in the list
func pathIn(testString string, list []string) bool {
	for _, x := range list {
		if x == testString {
			return true
		}
	}
	return false
}

// loadConfigIntoGnmiNotification load a config from a string returning the rootpath and the config as *gnmi.Notification
func loadConfigIntoGnmiNotification(config string) (*yang.Entry, *gnmi.Notification, error) {
	device := &ygotsrl.Device{}
	err := ygotsrl.Unmarshal([]byte(config), device)
	if err != nil {
		return nil, nil, err
	}

	notification, err := ygot.TogNMINotifications(device, time.Now().UnixNano(), ygot.GNMINotificationsConfig{UsePathElem: true})
	if err != nil {
		return nil, nil, err
	}
	// get Root Schema
	schema, err := ygotsrl.Schema()
	if err != nil {
		return nil, nil, err
	}
	deviceSchema := schema.RootSchema()

	return deviceSchema, notification[0], nil

}

// simple config with just an interface, that should result in a rootpath representing the subinterface
var sampleDataIf = `
{
    "interface": [
        {
            "admin-state": "enable",
            "name": "ethernet-1/48",
            "subinterface": [
                {
                    "admin-state": "enable",
                    "index": 100,
                    "ipv4": {
                        "address": [
                            {
                                "ip-prefix": "100.112.10.1/31"
                            }
                        ],
                        "allow-directed-broadcast": false
                    },
                    "ipv6": {
                        "address": [
                            {
                                "ip-prefix": "2a02:1800:80:7050::1/64"
                            }
                        ]
                    },
                    "type": "routed",
                    "vlan": {
                        "encap": {
                            "single-tagged": {
                                "vlan-id": 100
                            }
                        }
                    }
                }
            ]
        }
    ]
}
`

// a config that contains multiple interfsaces / subinterfaces network instances etc.
var sampleDataLarger = `
{
    "interface": [
        {
            "admin-state": "enable",
            "name": "ethernet-1/48",
            "subinterface": [
                {
                    "admin-state": "enable",
                    "index": 100,
                    "ipv4": {
                        "address": [
                            {
                                "ip-prefix": "100.112.10.1/31"
                            }
                        ],
                        "allow-directed-broadcast": false
                    },
                    "ipv6": {
                        "address": [
                            {
                                "ip-prefix": "2a02:1800:80:7050::1/64"
                            }
                        ]
                    },
                    "type": "routed",
                    "vlan": {
                        "encap": {
                            "single-tagged": {
                                "vlan-id": 100
                            }
                        }
                    }
                }
            ]
        },
        {
            "admin-state": "enable",
            "name": "irb0",
            "subinterface": [
                {
                    "admin-state": "enable",
                    "anycast-gw": {
                        "virtual-router-id": 1
                    },
                    "index": 1386,
                    "ipv4": {
                        "address": [
                            {
                                "anycast-gw": true,
                                "ip-prefix": "100.112.3.1/24"
                            }
                        ],
                        "allow-directed-broadcast": false,
                        "arp": {
                            "duplicate-address-detection": true,
                            "evpn": {
                                "advertise": [
                                    {
                                        "admin-tag": 0,
                                        "route-type": "dynamic"
                                    }
                                ]
                            },
                            "host-route": {
                                "populate": [
                                    {
                                        "route-type": "dynamic"
                                    }
                                ]
                            },
                            "learn-unsolicited": true,
                            "timeout": 14400
                        }
                    },
                    "ipv6": {
                        "address": [
                            {
                                "ip-prefix": "2a02:1800:83:7000::1/64"
                            }
                        ],
                        "neighbor-discovery": {
                            "duplicate-address-detection": true,
                            "evpn": {
                                "advertise": [
                                    {
                                        "admin-tag": 0,
                                        "route-type": "dynamic"
                                    }
                                ]
                            },
                            "host-route": {
                                "populate": [
                                    {
                                        "route-type": "dynamic"
                                    }
                                ]
                            },
                            "learn-unsolicited": "global",
                            "reachable-time": 30,
                            "stale-time": 14400
                        }
                    }
                },
                {
                    "admin-state": "enable",
                    "index": 2105
                },
                {
                    "admin-state": "enable",
                    "index": 2106
                },
                {
                    "admin-state": "enable",
                    "index": 2143
                }
            ]
        },
        {
            "admin-state": "enable",
            "name": "lag1",
            "subinterface": [
                {
                    "admin-state": "enable",
                    "index": 100,
                    "type": "bridged",
                    "vlan": {
                        "encap": {
                            "single-tagged": {
                                "vlan-id": 100
                            }
                        }
                    }
                },
                {
                    "admin-state": "enable",
                    "index": 102,
                    "type": "bridged",
                    "vlan": {
                        "encap": {
                            "single-tagged": {
                                "vlan-id": 102
                            }
                        }
                    }
                },
                {
                    "admin-state": "enable",
                    "index": 200,
                    "type": "bridged",
                    "vlan": {
                        "encap": {
                            "single-tagged": {
                                "vlan-id": 200
                            }
                        }
                    }
                }
            ]
        }
    ],
    "network-instance": [
        {
            "admin-state": "enable",
            "bridge-table": {
                "discard-unknown-dest-mac": false,
                "mac-learning": {
                    "admin-state": "enable"
                },
                "protect-anycast-gw-mac": false
            },
            "interface": [
                {
                    "name": "irb0.2143"
                },
                {
                    "name": "lag1.100"
                }
            ],
            "name": "multus-ipvlan-bridged",
            "protocols": {
                "bgp-evpn": {
                    "bgp-instance": [
                        {
                            "admin-state": "enable",
                            "default-admin-tag": 0,
                            "ecmp": 2,
                            "encapsulation-type": "vxlan",
                            "evi": 2143,
                            "id": 1,
                            "vxlan-interface": "vxlan0.2143"
                        }
                    ]
                },
                "bgp-vpn": {
                    "bgp-instance": [
                        {
                            "id": 1,
                            "route-target": {
                                "export-rt": "target:65555:2143",
                                "import-rt": "target:65555:2143"
                            }
                        }
                    ]
                }
            },
            "type": "mac-vrf",
            "vxlan-interface": [
                {
                    "name": "vxlan0.2143"
                }
            ]
        },
        {
            "admin-state": "enable",
            "interface": [
                {
                    "name": "ethernet-1/48.100"
                },
                {
                    "name": "irb0.1386"
                }
            ],
            "ip-forwarding": {
                "receive-ipv4-check": true,
                "receive-ipv6-check": true
            },
            "name": "multus-routed",
            "protocols": {
                "bgp-evpn": {
                    "bgp-instance": [
                        {
                            "admin-state": "enable",
                            "default-admin-tag": 0,
                            "ecmp": 2,
                            "encapsulation-type": "vxlan",
                            "evi": 1386,
                            "id": 1,
                            "vxlan-interface": "vxlan0.1386"
                        }
                    ]
                },
                "bgp-vpn": {
                    "bgp-instance": [
                        {
                            "id": 1,
                            "route-target": {
                                "export-rt": "target:65555:1386",
                                "import-rt": "target:65555:1386"
                            }
                        }
                    ]
                }
            },
            "type": "ip-vrf",
            "vxlan-interface": [
                {
                    "name": "vxlan0.1386"
                }
            ]
        },
        {
            "admin-state": "enable",
            "bridge-table": {
                "discard-unknown-dest-mac": false,
                "mac-learning": {
                    "admin-state": "enable"
                },
                "protect-anycast-gw-mac": false
            },
            "interface": [
                {
                    "name": "irb0.2105"
                },
                {
                    "name": "lag1.200"
                }
            ],
            "name": "multus-sriov1-bridged",
            "protocols": {
                "bgp-evpn": {
                    "bgp-instance": [
                        {
                            "admin-state": "enable",
                            "default-admin-tag": 0,
                            "ecmp": 2,
                            "encapsulation-type": "vxlan",
                            "evi": 2105,
                            "id": 1,
                            "vxlan-interface": "vxlan0.2105"
                        }
                    ]
                },
                "bgp-vpn": {
                    "bgp-instance": [
                        {
                            "id": 1,
                            "route-target": {
                                "export-rt": "target:65555:2105",
                                "import-rt": "target:65555:2105"
                            }
                        }
                    ]
                }
            },
            "type": "mac-vrf",
            "vxlan-interface": [
                {
                    "name": "vxlan0.2105"
                }
            ]
        },
        {
            "admin-state": "enable",
            "bridge-table": {
                "discard-unknown-dest-mac": false,
                "mac-learning": {
                    "admin-state": "enable"
                },
                "protect-anycast-gw-mac": false
            },
            "interface": [
                {
                    "name": "irb0.2106"
                },
                {
                    "name": "lag1.102"
                }
            ],
            "name": "multus-sriov2-bridged",
            "protocols": {
                "bgp-evpn": {
                    "bgp-instance": [
                        {
                            "admin-state": "enable",
                            "default-admin-tag": 0,
                            "ecmp": 2,
                            "encapsulation-type": "vxlan",
                            "evi": 2106,
                            "id": 1,
                            "vxlan-interface": "vxlan0.2106"
                        }
                    ]
                },
                "bgp-vpn": {
                    "bgp-instance": [
                        {
                            "id": 1,
                            "route-target": {
                                "export-rt": "target:65555:2106",
                                "import-rt": "target:65555:2106"
                            }
                        }
                    ]
                }
            },
            "type": "mac-vrf",
            "vxlan-interface": [
                {
                    "name": "vxlan0.2106"
                }
            ]
        }
    ],
    "tunnel-interface": [
        {
            "name": "vxlan0",
            "vxlan-interface": [
                {
                    "egress": {
                        "source-ip": "use-system-ipv4-address"
                    },
                    "index": 1386,
                    "ingress": {
                        "vni": 1386
                    },
                    "type": "routed"
                },
                {
                    "egress": {
                        "source-ip": "use-system-ipv4-address"
                    },
                    "index": 2105,
                    "ingress": {
                        "vni": 2105
                    },
                    "type": "bridged"
                },
                {
                    "egress": {
                        "source-ip": "use-system-ipv4-address"
                    },
                    "index": 2106,
                    "ingress": {
                        "vni": 2106
                    },
                    "type": "bridged"
                },
                {
                    "egress": {
                        "source-ip": "use-system-ipv4-address"
                    },
                    "index": 2143,
                    "ingress": {
                        "vni": 2143
                    },
                    "type": "bridged"
                }
            ]
        }
    ]
}
`
