package main

import (
	"fmt"
	"net"
)

func main() {
	ipAndMask := getIPAndMask()

	for interfaceName, info := range ipAndMask {
		fmt.Printf("[%s] %s/%s\n", interfaceName, info.ip, info.netmask)
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
