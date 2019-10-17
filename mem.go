package sysinfo

import (
	units "github.com/docker/go-units"
	"github.com/shirou/gopsutil/mem"
)

type MemInfo struct {
	Total          string
	Free           string
	UsedPercentage string
}

func GetMemInfo() (*MemInfo, error) {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	return &MemInfo{
		Total:          units.BytesSize(float64(vmStat.Total)),
		Free:           units.BytesSize(float64(vmStat.Free)),
		UsedPercentage: formatPercent(vmStat.UsedPercent),
	}, nil
}
