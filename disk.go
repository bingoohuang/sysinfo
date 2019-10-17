package sysinfo

import (
	"fmt"

	units "github.com/docker/go-units"
	"github.com/shirou/gopsutil/disk"
)

type DiskInfo struct {
	Path        string
	Device      string
	Fstype      string
	Total       string
	Used        string
	Free        string
	UsedPercent string
}

func GetDiskInfos() ([]DiskInfo, error) {
	stats, err := disk.Partitions(false)
	if err != nil {
		return nil, err
	}

	diskInfos := make([]DiskInfo, 0)

	for _, p := range stats {
		diskStat, err := disk.Usage(p.Mountpoint)
		if err != nil {
			continue
		}

		diskInfos = append(diskInfos, DiskInfo{
			Device:      p.Device,
			Path:        p.Mountpoint,
			Fstype:      p.Fstype,
			Total:       units.BytesSize(float64(diskStat.Total)),
			Used:        units.BytesSize(float64(diskStat.Used)),
			Free:        units.BytesSize(float64(diskStat.Free)),
			UsedPercent: formatPercent(diskStat.UsedPercent),
		})
	}

	return diskInfos, nil
}

func formatPercent(percent float64) string {
	up := fmt.Sprintf("%0.2f%%", percent)

	if percent < 10 {
		return "0" + up
	}

	return up
}
