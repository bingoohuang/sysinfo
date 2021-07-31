package sysinfo

import (
	"time"

	units "github.com/docker/go-units"
	"github.com/shirou/gopsutil/v3/host"
)

// HostInfo ...
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

// GetHostInfo ...
func GetHostInfo() (*HostInfo, error) {
	// host or machine kernel, uptime, platform Info
	hostStat, err := host.Info()
	if err != nil {
		return nil, err
	}

	return &HostInfo{
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
