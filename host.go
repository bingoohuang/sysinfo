package sysinfo

import (
	"time"

	"github.com/gobars/cmd"

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
	KernelArch      string
	OsRelease       string
}

// GetHostInfo ...
func GetHostInfo() (*HostInfo, error) {
	// host or machine kernel, uptime, platform Info
	hostStat, err := host.Info()
	if err != nil {
		return nil, err
	}

	osRelease := ""
	cmd.BashLiner(`egrep '^(VERSION|NAME)=' /etc/os-release`, func(line string) bool {
		if osRelease != "" {
			osRelease += " "
		}
		osRelease += line
		return true
	})

	return &HostInfo{
		Hostname:        hostStat.Hostname,
		Uptime:          hostStat.Uptime,
		UptimeHuman:     units.HumanDuration(time.Duration(hostStat.Uptime) * time.Second),
		Procs:           hostStat.Procs,
		OS:              hostStat.OS,
		OsRelease:       osRelease,
		Platform:        hostStat.Platform,
		HostID:          hostStat.HostID,
		PlatformVersion: hostStat.PlatformVersion,
		KernelArch:      hostStat.KernelArch,
		KernelVersion:   hostStat.KernelVersion,
	}, nil
}
