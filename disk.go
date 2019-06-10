package main

import (
	"fmt"

	"github.com/jedib0t/go-pretty/table"

	"github.com/docker/go-units"
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

func (p TablePrinter) tableDiskInfos(ds []DiskInfo) {
	rows := make([]table.Row, len(ds))
	for i, c := range ds {
		rows[i] = table.Row{i + 1, c.Device, c.Path, c.Fstype, c.Total, c.Used, c.Free, c.UsedPercent}
	}

	p.TableRender(table.Row{"#", "Device", "Path", "Fstype", "Disk Total", "Disk Used", "Disk Free", "Disk Used"}, rows...)
	fmt.Println()
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
			UsedPercent: fmt.Sprintf("%.2f%%", diskStat.UsedPercent),
		})
	}

	return diskInfos, nil
}
