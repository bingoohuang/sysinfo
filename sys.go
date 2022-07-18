package sysinfo

import (
	"runtime"

	"github.com/klauspost/cpuid/v2"
)

// ErrorInfo ...
type ErrorInfo struct {
	Error string
}

// SysInfo ...
type SysInfo struct {
	OS        string     `json:",omitempty"`
	MemInfo   *MemInfo   `json:",omitempty"`
	DiskInfos []DiskInfo `json:",omitempty"`
	// CPUInfos    []CPUInfo      `json:",omitempty"`
	CPUInfo     *cpuid.CPUInfo `json:",omitempty"`
	HostInfo    *HostInfo      `json:",omitempty"`
	InterfInfos []InterfInfo   `json:",omitempty"`
	PsItems     []PsAuxItem    `json:",omitempty"`
	Errors      []ErrorInfo    `json:",omitempty"`
}

// GetSysInfo ...
func GetSysInfo(showsMap map[string]bool) SysInfo {
	var err error

	errs := make([]ErrorInfo, 0)
	si := SysInfo{OS: runtime.GOOS}

	if _, ok := showsMap[("disk")]; ok {
		if si.DiskInfos, err = GetDiskInfos(); err != nil {
			errs = append(errs, ErrorInfo{err.Error()})
		}
	}

	if _, ok := showsMap[("mem")]; ok {
		if si.MemInfo, err = GetMemInfo(); err != nil {
			errs = append(errs, ErrorInfo{err.Error()})
		}
	}

	if _, ok := showsMap["cpu"]; ok {
		//if si.CPUInfos, err = GetCPUInfo(); err != nil {
		//	errs = append(errs, ErrorInfo{err.Error()})
		//}

		si.CPUInfo = &cpuid.CPU
	}

	if _, ok := showsMap[("host")]; ok {
		if si.HostInfo, err = GetHostInfo(); err != nil {
			errs = append(errs, ErrorInfo{err.Error()})
		}
	}

	if _, ok := showsMap[("interf")]; ok {
		if si.InterfInfos, err = GetInterInfos(); err != nil {
			errs = append(errs, ErrorInfo{err.Error()})
		}
	}

	if _, ok := showsMap[("ps")]; ok {
		if si.PsItems, err = PsAuxTop(0); err != nil {
			errs = append(errs, ErrorInfo{err.Error()})
		}
	}

	si.Errors = errs

	return si
}
