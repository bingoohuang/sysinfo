package sysinfo

import "runtime"

type ErrorInfo struct {
	Error string
}
type SysInfo struct {
	OS          string
	MemInfo     *MemInfo
	DiskInfos   []DiskInfo
	CPUInfos    []CPUInfo
	HostInfo    *HostInfo
	InterfInfos []InterfInfo
	PsItems     []PsAuxItem
	Errors      []ErrorInfo `json:",omitempty"`
}

const (
	Disk   = "disk"
	Mem    = "mem"
	CPU    = "cpu"
	Host   = "host"
	Interf = "interf"
	PS     = "ps"
)

func GetSysInfo(showsMap map[string]bool) SysInfo {
	var err error

	errs := make([]ErrorInfo, 0)
	si := SysInfo{OS: runtime.GOOS}

	if _, ok := showsMap[Disk]; ok {
		if si.DiskInfos, err = GetDiskInfos(); err != nil {
			errs = append(errs, ErrorInfo{err.Error()})
		}
	}

	if _, ok := showsMap[Mem]; ok {
		if si.MemInfo, err = GetMemInfo(); err != nil {
			errs = append(errs, ErrorInfo{err.Error()})
		}
	}

	if _, ok := showsMap[CPU]; ok {
		if si.CPUInfos, err = GetCPUInfo(); err != nil {
			errs = append(errs, ErrorInfo{err.Error()})
		}
	}

	if _, ok := showsMap[Host]; ok {
		if si.HostInfo, err = GetHostInfo(); err != nil {
			errs = append(errs, ErrorInfo{err.Error()})
		}
	}

	if _, ok := showsMap[Interf]; ok {
		if si.InterfInfos, err = GetInterInfos(); err != nil {
			errs = append(errs, ErrorInfo{err.Error()})
		}
	}

	if _, ok := showsMap[PS]; ok {
		if si.PsItems, err = PsAuxTop(0); err != nil {
			errs = append(errs, ErrorInfo{err.Error()})
		}
	}

	si.Errors = errs

	return si
}
