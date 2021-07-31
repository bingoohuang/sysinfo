package sysinfo

import (
	units "github.com/docker/go-units"
	"github.com/shirou/gopsutil/v3/mem"
)

// MemInfo ...
type MemInfo struct {
	Total          string
	Free           string
	Available      string
	UsedPercentage string
}

// GetMemInfo ...
func GetMemInfo() (*MemInfo, error) {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	return &MemInfo{
		Total:          units.BytesSize(float64(vmStat.Total)),
		Free:           units.BytesSize(float64(vmStat.Free)),
		Available:      units.BytesSize(float64(vmStat.Available)),
		UsedPercentage: formatPercent(vmStat.UsedPercent),
	}, nil
}
