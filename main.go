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

func main() {
	ipAndMask := getIPAndMask()

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

func getIPAndMask() map[string]IPInfo {
	ipMaskList := make(map[string]IPInfo)

	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error getting network interfaces:", err)
		return ipMaskList
	}

	for _, iface := range interfaces {
		addrs, err := iface.Addrs()
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
