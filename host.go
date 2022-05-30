package sysinfo

import (
	"runtime"
	"time"

	"github.com/gobars/cmd"

	units "github.com/docker/go-units"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
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
	NumCPU          int
	MemAvailable    string
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

	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

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
		NumCPU:          runtime.NumCPU(),
		MemAvailable: units.BytesSize(float64(vmStat.Available)) + "/" + units.BytesSize(float64(vmStat.Total)) +
			", " + formatPercent(float64(vmStat.Available)/float64(vmStat.Total)),
	}, nil
}
