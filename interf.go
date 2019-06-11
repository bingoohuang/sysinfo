package main

import (
	"strings"

	"github.com/shirou/gopsutil/net"
	"github.com/thoas/go-funk"
)

type InterfInfo struct {
	InterfaceName string
	HardwareAddr  string
	Addrs         string
}

func GetInterInfos() ([]InterfInfo, error) {
	// get interfaces MAC/hardware address
	interfStats, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	interfs := make([]InterfInfo, 0, len(interfStats))
	for _, interf := range interfStats {
		if interf.HardwareAddr == "" || len(interf.Addrs) == 0 {
			continue
		}

		addrs := funk.Map(interf.Addrs, func(a net.InterfaceAddr) string { return a.Addr }).([]string)
		interfs = append(interfs, InterfInfo{
			InterfaceName: interf.Name,
			HardwareAddr:  interf.HardwareAddr,
			Addrs:         strings.Join(addrs, " "),
		})
	}

	return interfs, nil
}
