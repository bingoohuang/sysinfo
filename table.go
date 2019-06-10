package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/table"
)

func printTable() {
	info := GetSysInfo()

	tableHost(info)
	fmt.Println()

	tableMem(info)
	fmt.Println()

	tableCPUInfos(info)
	fmt.Println()

	tableDiskInfos(info)
	fmt.Println()

	tableInterfInfos(info)
	fmt.Println()

	if len(info.Errors) > 0 {
		tableErrors(info)
		fmt.Println()
	}
}

func tableErrors(info SysInfo) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Error"})
	for i, c := range info.Errors {
		t.AppendRow(table.Row{i + 1, c})
	}
	t.Render()
}

func tableInterfInfos(info SysInfo) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Interface Name", "Hardware Addr", "Addrs"})
	for i, c := range info.InterfInfos {
		t.AppendRow(table.Row{i + 1, c.InterfaceName, c.HardwareAddr, strings.Join(c.Addrs, " ")})
	}
	t.Render()
}

func tableDiskInfos(info SysInfo) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Device", "Path", "Fstype", "Disk Total", "Disk Used", "Disk Free", "Disk Used Percent"})
	for i, c := range info.DiskInfos {
		t.AppendRow(table.Row{i + 1, c.Device, c.Path, c.Fstype, c.Total, c.Used, c.Free, c.UsedPercent})
	}
	t.Render()
}

func tableCPUInfos(info SysInfo) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"CPU PhysicalID", "VendorID", "Family", "Model Name", "Cores", "Mhz"})
	for _, c := range info.CPUInfos {
		t.AppendRow(table.Row{c.PhysicalID, c.VendorID, c.Family, c.ModelName, c.Cores, c.Mhz})
	}
	t.Render()
}

func tableHost(info SysInfo) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"OS",
		"Hostname", "Uptime", "Procs", "HostOS", "Platform", "HostID", "Platform Version", "Kernel Version"})
	h := info.HostInfo
	t.AppendRow(table.Row{info.OS,
		h.Hostname, h.Uptime, h.Procs, h.OS, h.Platform, h.HostID, h.PlatformVersion, h.KernelVersion})
	t.Render()
}

func tableMem(info SysInfo) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Total Mem", "Free Mem", "Mem Used Percentage"})
	m := info.MemInfo
	t.AppendRow(table.Row{m.Total, m.Free, m.UsedPercentage})
	t.Render()
}
