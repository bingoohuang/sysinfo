package main

import "runtime"

type SysInfo struct {
	OS          string
	MemInfo     MemInfo
	DiskInfos   []DiskInfo
	CPUInfos    []CPUInfo
	HostInfo    HostInfo
	InterfInfos []InterfInfo
	Errors      []string `json:",omitempty"`
}

func GetSysInfo() SysInfo {
	errs := make([]string, 0)
	diskInfos, err := GetDiskInfos()
	if err != nil {
		errs = append(errs, err.Error())
	}

	mem, err := GetMemInfo()
	if err != nil {
		errs = append(errs, err.Error())
	}

	cpuInfos, err := GetCPUInfo()
	if err != nil {
		errs = append(errs, err.Error())
	}

	hostInfo, err := GetHostInfo()
	if err != nil {
		errs = append(errs, err.Error())
	}

	interfInfos, err := GetInterInfos()
	if err != nil {
		errs = append(errs, err.Error())
	}

	return SysInfo{
		OS:          runtime.GOOS,
		MemInfo:     mem,
		DiskInfos:   diskInfos,
		CPUInfos:    cpuInfos,
		HostInfo:    hostInfo,
		InterfInfos: interfInfos,
		Errors:      errs,
	}
}
