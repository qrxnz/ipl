package main

import (
	"errors"
	"net"
	"reflect"
	"testing"
)

// mockNetInfo is a mock implementation of the netInfo interface for testing.
type mockNetInfo struct {
	interfaces []net.Interface
	addrs      map[string][]net.Addr
	err        error
}

// Interfaces returns the mock list of interfaces.
func (m *mockNetInfo) Interfaces() ([]net.Interface, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.interfaces, nil
}

// InterfaceAddrs returns the mock list of addresses for a given interface.
func (m *mockNetInfo) InterfaceAddrs(iface *net.Interface) ([]net.Addr, error) {
	if addrs, ok := m.addrs[iface.Name]; ok {
		if addrs == nil {
			return nil, errors.New("mock error getting addresses")
		}
		return addrs, nil
	}
	return nil, nil
}

func TestGetIPAndMask(t *testing.T) {
	// Mock data
	ip1, ipNet1, _ := net.ParseCIDR("192.168.1.1/24")
	ipNet1.IP = ip1
	ip2, ipNet2, _ := net.ParseCIDR("10.0.0.1/8")
	ipNet2.IP = ip2

	tests := []struct {
		name     string
		mock     *mockNetInfo
		expected map[string]IPInfo
	}{
		{
			name: "successful case with two interfaces",
			mock: &mockNetInfo{
				interfaces: []net.Interface{
					{Name: "eth0"},
					{Name: "lo"},
				},
				addrs: map[string][]net.Addr{
					"eth0": {ipNet1},
					"lo":   {ipNet2},
				},
			},
			expected: map[string]IPInfo{
				"eth0": {ip: "192.168.1.1", netmask: "255.255.255.0"},
				"lo":   {ip: "10.0.0.1", netmask: "255.0.0.0"},
			},
		},
		{
			name: "interface with no addresses",
			mock: &mockNetInfo{
				interfaces: []net.Interface{
					{Name: "eth0"},
				},
				addrs: map[string][]net.Addr{
					"eth0": {},
				},
			},
			expected: map[string]IPInfo{},
		},
		{
			name: "error getting interfaces",
			mock: &mockNetInfo{
				err: errors.New("interfaces error"),
			},
			expected: map[string]IPInfo{},
		},
		{
			name: "interface with non-IPNet address",
			mock: &mockNetInfo{
				interfaces: []net.Interface{
					{Name: "eth0"},
				},
				addrs: map[string][]net.Addr{
					"eth0": {&net.IPAddr{IP: net.ParseIP("192.168.1.1")}},
				},
			},
			expected: map[string]IPInfo{},
		},
		{
			name: "interface with IPv6 address",
			mock: &mockNetInfo{
				interfaces: []net.Interface{
					{Name: "eth0"},
				},
				addrs: map[string][]net.Addr{
					"eth0": {
						&net.IPNet{
							IP:   net.ParseIP("2001:db8::1"),
							Mask: net.CIDRMask(64, 128),
						},
					},
				},
			},
			expected: map[string]IPInfo{},
		},
		{
			name: "error getting addresses for an interface",
			mock: &mockNetInfo{
				interfaces: []net.Interface{
					{Name: "eth0"},
				},
				addrs: map[string][]net.Addr{
					"eth0": nil, // This will cause an error in the mock
				},
			},
			expected: map[string]IPInfo{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := getIPAndMask(tt.mock)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("getIPAndMask() = %v, want %v", actual, tt.expected)
			}
		})
	}
}
