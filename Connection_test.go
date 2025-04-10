package gonetworkmanager

import (
	"reflect"
	"testing"

	"github.com/kubeovn/dbus/v5"
)

func TestDecodeSettings(t *testing.T) {
	settings := map[string]map[string]dbus.Variant{
		"ipv4": {
			"address-data": dbus.MakeVariant([]map[string]dbus.Variant{
				{
					"address": dbus.MakeVariant("192.168.1.156"),
					"prefix":  dbus.MakeVariant(24),
				},
			}),
			"dns-search": dbus.MakeVariant([]string{}),
			"method":     dbus.MakeVariant("manual"),
			"route-data": dbus.MakeVariant([]map[string]dbus.Variant{}),
			"routes":     dbus.MakeVariant([][]uint32{}),
			"addresses": dbus.MakeVariant([][]uint32{
				{
					2617354432,
					24,
					16885952,
				},
			}),
			"gateway":      dbus.MakeVariant("192.168.1.1"),
			"route-metric": dbus.MakeVariant(100),
			"dhcp-timeout": dbus.MakeVariant(45),
		},
		"ipv6": {
			"addr-gen-mode": dbus.MakeVariant(3),
			"address-data":  dbus.MakeVariant([]map[string]dbus.Variant{}),
			"routes":        dbus.MakeVariant([][]interface{}{}),
			"dns-search":    dbus.MakeVariant([]string{}),
			"method":        dbus.MakeVariant("auto"),
			"route-data":    dbus.MakeVariant([]map[string]dbus.Variant{}),
			"dhcp-timeout":  dbus.MakeVariant(45),
			"route-metric":  dbus.MakeVariant(100),
			"addresses":     dbus.MakeVariant([][]interface{}{}),
		},
		"proxy": {},
		"connection": {
			"uuid":                 dbus.MakeVariant("390e5c2b-7312-415e-80e6-7b94a5c24fc3"),
			"autoconnect-priority": dbus.MakeVariant(1),
			"autoconnect-retries":  dbus.MakeVariant(0),
			"id":                   dbus.MakeVariant("main"),
			"interface-name":       dbus.MakeVariant("eth0"),
			"permissions":          dbus.MakeVariant([]string{}),
			"timestamp":            dbus.MakeVariant(1669049774),
			"type":                 dbus.MakeVariant("802-3-ethernet"),
		},
		"802-3-ethernet": {
			"auto-negotiate":        dbus.MakeVariant(false),
			"mac-address-blacklist": dbus.MakeVariant([]string{}),
			"s390-options":          dbus.MakeVariant(map[string]string{}),
		},
	}

	result := decodeSettings(settings)

	expected := ConnectionSettings{
		"ipv4": {
			"address-data": []map[string]interface{}{
				{
					"address": "192.168.1.156",
					"prefix":  24,
				},
			},
			"dns-search": []string{},
			"method":     "manual",
			"route-data": []map[string]interface{}(nil),
			"routes":     [][]uint32{},
			"addresses": [][]uint32{
				{
					2617354432,
					24,
					16885952,
				},
			},
			"gateway":      "192.168.1.1",
			"route-metric": 100,
			"dhcp-timeout": 45,
		},
		"ipv6": {
			"addr-gen-mode": 3,
			"address-data":  []map[string]interface{}(nil),
			"routes":        [][]interface{}{},
			"dns-search":    []string{},
			"method":        "auto",
			"route-data":    []map[string]interface{}(nil),
			"dhcp-timeout":  45,
			"route-metric":  100,
			"addresses":     [][]interface{}{},
		},
		"proxy": {},
		"connection": {
			"uuid":                 "390e5c2b-7312-415e-80e6-7b94a5c24fc3",
			"autoconnect-priority": 1,
			"autoconnect-retries":  0,
			"id":                   "main",
			"interface-name":       "eth0",
			"permissions":          []string{},
			"timestamp":            1669049774,
			"type":                 "802-3-ethernet",
		},
		"802-3-ethernet": {
			"auto-negotiate":        false,
			"mac-address-blacklist": []string{},
			"s390-options":          map[string]string{},
		},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("failed: \nexpected: %#v\nresult  : %#v", expected, result)
	}
}

func TestDecode(t *testing.T) {
	ipSettings := map[string]dbus.Variant{
		"address-data": dbus.MakeVariant([]map[string]dbus.Variant{
			{
				"address": dbus.MakeVariant("192.168.1.156"),
				"prefix":  dbus.MakeVariant(24),
			},
		}),
		"dns-search": dbus.MakeVariant([]string{}),
		"method":     dbus.MakeVariant("manual"),
		"route-data": dbus.MakeVariant([]map[string]dbus.Variant{}),
		"routes":     dbus.MakeVariant([][]uint32{}),
		"addresses": dbus.MakeVariant([][]uint32{
			{
				2617354432,
				24,
				16885952,
			},
		}),
		"gateway":      dbus.MakeVariant("192.168.1.1"),
		"route-metric": dbus.MakeVariant(100),
		"dhcp-timeout": dbus.MakeVariant(45),
	}

	result := decode(ipSettings)

	expected := map[string]interface{}{
		"address-data": []map[string]interface{}{
			{
				"address": "192.168.1.156",
				"prefix":  24,
			},
		},
		"dns-search": []string{},
		"method":     "manual",
		"route-data": []map[string]interface{}(nil),
		"routes":     [][]uint32{},
		"addresses": [][]uint32{
			{
				2617354432,
				24,
				16885952,
			},
		},
		"gateway":      "192.168.1.1",
		"route-metric": 100,
		"dhcp-timeout": 45,
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("failed: \nexpected: %#v\nresult  : %#v", expected, result)
	}
}
