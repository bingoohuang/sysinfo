package main

import (
	"fmt"
	"time"

	"github.com/docker/go-units"
	"github.com/jedib0t/go-pretty/table"
	"github.com/shirou/gopsutil/host"
)

type HostInfo struct {
	Hostname        string
	Uptime          uint64
	UptimeHuman     string
	Procs           uint64
	OS              string
	Platform        string
	HostID          string
	PlatformVersion string
	KernelVersion   string
}

func (p TablePrinter) tableHost(h HostInfo) {
	p.TableRender(table.Row{"Hostname", "Uptime", "Procs", "Host OS",
		"Platform", "HostID", "Platform Version", "Kernel Version"},
		table.Row{h.Hostname, fmt.Sprintf("%d(%s)", h.Uptime, h.UptimeHuman), h.Procs, h.OS,
			h.Platform, h.HostID, h.PlatformVersion, h.KernelVersion})

	fmt.Println()
}

func GetHostInfo() (HostInfo, error) {
	// host or machine kernel, uptime, platform Info
	hostStat, err := host.Info()
	if err != nil {
		return HostInfo{}, err
	}

	return HostInfo{
		Hostname:        hostStat.Hostname,
		Uptime:          hostStat.Uptime,
		UptimeHuman:     units.HumanDuration(time.Duration(hostStat.Uptime) * time.Second),
		Procs:           hostStat.Procs,
		OS:              hostStat.OS,
		Platform:        hostStat.Platform,
		HostID:          hostStat.HostID,
		PlatformVersion: hostStat.PlatformVersion,
		KernelVersion:   hostStat.KernelVersion,
	}, nil
}
