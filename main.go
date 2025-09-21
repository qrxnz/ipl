package main

import (
	"fmt"
	"net"

	"github.com/charmbracelet/lipgloss"
)

var (
	interfaceStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("63"))  // Purple
	ipStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("86"))   // Aqua
	separatorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240")) // Gray
)

// netInfo provides an interface for network-related operations, to allow for mocking in tests.
type netInfo interface {
	Interfaces() ([]net.Interface, error)
	InterfaceAddrs(iface *net.Interface) ([]net.Addr, error)
}

// realNetInfo is a concrete implementation of netInfo that uses the net package.
type realNetInfo struct{}

// Interfaces returns the list of system's network interfaces.
func (r *realNetInfo) Interfaces() ([]net.Interface, error) {
	return net.Interfaces()
}

// InterfaceAddrs returns the list of addresses for a given interface.
func (r *realNetInfo) InterfaceAddrs(iface *net.Interface) ([]net.Addr, error) {
	return iface.Addrs()
}

func main() {
	netinfo := &realNetInfo{}
	ipAndMask := getIPAndMask(netinfo)

	for interfaceName, info := range ipAndMask {
		fmt.Printf("[%s] %s%s%s\n",
			interfaceStyle.Render(interfaceName),
			ipStyle.Render(info.ip),
			separatorStyle.Render("/"),
			ipStyle.Render(info.netmask),
		)
	}
}

type IPInfo struct {
	ip      string
	netmask string
}

func getIPAndMask(n netInfo) map[string]IPInfo {
	ipMaskList := make(map[string]IPInfo)

	interfaces, err := n.Interfaces()
	if err != nil {
		fmt.Println("Error getting network interfaces:", err)
		return ipMaskList
	}

	for _, iface := range interfaces {
		addrs, err := n.InterfaceAddrs(&iface)
		if err != nil {
			fmt.Println("Error getting addresses for interface", iface.Name, ":", err)
			continue
		}

		for _, addr := range addrs {
			if ipNet, ok := addr.(*net.IPNet); ok && ipNet.IP.To4() != nil { // Check for IPv4
				ipMaskList[iface.Name] = IPInfo{
					ip:      ipNet.IP.String(),
					netmask: net.IP(ipNet.Mask).String(),
				}
			}
		}
	}

	return ipMaskList
}
