package main

import (
	"fmt"
	"strings"

	"github.com/jedib0t/go-pretty/table"
	"github.com/shirou/gopsutil/net"
	"github.com/thoas/go-funk"
)

type InterfInfo struct {
	InterfaceName string
	HardwareAddr  string
	Addrs         []string
}

func (p TablePrinter) tableInterfInfos(is []InterfInfo) {
	rows := make([]table.Row, len(is))
	for i, c := range is {
		rows[i] = table.Row{i + 1, c.InterfaceName, c.HardwareAddr, strings.Join(c.Addrs, " ")}
	}

	p.TableRender(table.Row{"#", "Interface Name", "Hardware Addr", "Addrs"}, rows...)
	fmt.Println()
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

		interfs = append(interfs, InterfInfo{
			InterfaceName: interf.Name,
			HardwareAddr:  interf.HardwareAddr,
			Addrs:         funk.Map(interf.Addrs, func(a net.InterfaceAddr) string { return a.Addr }).([]string),
		})
	}

	return interfs, nil
}
