package sysinfo

import "runtime"

type ErrorInfo struct {
	Error string
}
type SysInfo struct {
	OS          string
	MemInfo     MemInfo
	DiskInfos   []DiskInfo
	CPUInfos    []CPUInfo
	HostInfo    HostInfo
	InterfInfos []InterfInfo
	PsItems     []PsAuxItem
	Errors      []ErrorInfo `json:",omitempty"`
}

func GetSysInfo() SysInfo {
	errs := make([]ErrorInfo, 0)
	diskInfos, err := GetDiskInfos()

	if err != nil {
		errs = append(errs, ErrorInfo{err.Error()})
	}

	mem, err := GetMemInfo()
	if err != nil {
		errs = append(errs, ErrorInfo{err.Error()})
	}

	cpuInfos, err := GetCPUInfo()
	if err != nil {
		errs = append(errs, ErrorInfo{err.Error()})
	}

	hostInfo, err := GetHostInfo()
	if err != nil {
		errs = append(errs, ErrorInfo{err.Error()})
	}

	interfInfos, err := GetInterInfos()
	if err != nil {
		errs = append(errs, ErrorInfo{err.Error()})
	}

	psItems, err := PsAuxTop(0)
	if err != nil {
		errs = append(errs, ErrorInfo{err.Error()})
	}

	return SysInfo{
		OS:          runtime.GOOS,
		MemInfo:     mem,
		DiskInfos:   diskInfos,
		CPUInfos:    cpuInfos,
		HostInfo:    hostInfo,
		InterfInfos: interfInfos,
		PsItems:     psItems,
		Errors:      errs,
	}
}
