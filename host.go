package sysinfo

import (
	"runtime"
	"time"

	"github.com/bingoohuang/gg/pkg/goip"
	units "github.com/docker/go-units"
	"github.com/gobars/cmd"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

// HostInfo ...
type HostInfo struct {
	NumCPU          int
	MemAvailable    string
	OS              string
	KernelArch      string
	Platform        string
	CpuModel        string
	KernelVersion   string
	Ips             []string
	UptimeHuman     string
	PlatformVersion string
	Hostname        string
	HostID          string
	OsRelease       string
	Procs           uint64
	CpuMhz          float64
	Uptime          uint64
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

	_, ips := goip.MainIP()
	hi := &HostInfo{
		Ips:             ips,
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
		MemAvailable: units.BytesSize(float64(vmStat.Available)) + "/" + units.BytesSize(float64(vmStat.Total)) +
			", " + formatPercent(float64(vmStat.Available)/float64(vmStat.Total)),
		NumCPU: runtime.NumCPU(),
	}

	// cpu - get cpu number of cores and speed
	if cpuStats, _ := cpu.Info(); len(cpuStats) > 0 {
		c0 := cpuStats[0]
		hi.CpuMhz = c0.Mhz
		hi.CpuModel = c0.ModelName
	}

	return hi, nil
}
