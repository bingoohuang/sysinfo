package main

import "github.com/shirou/gopsutil/host"

type HostInfo struct {
	Hostname        string
	Uptime          uint64
	Procs           uint64
	OS              string
	Platform        string
	HostID          string
	PlatformVersion string
	KernelVersion   string
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
		Procs:           hostStat.Procs,
		OS:              hostStat.OS,
		Platform:        hostStat.Platform,
		HostID:          hostStat.HostID,
		PlatformVersion: hostStat.PlatformVersion,
		KernelVersion:   hostStat.KernelVersion,
	}, nil
}
