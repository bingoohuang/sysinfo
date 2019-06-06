package main

import (
	"fmt"

	"github.com/docker/go-units"
	"github.com/shirou/gopsutil/disk"
)

type DiskInfo struct {
	Path   string
	Device string
	Fstype string
	//DriveType   string
	//Vendor      string
	//Partitions  string
	Total       string
	Used        string
	Free        string
	UsedPercent string
}

//func placeholder(s string) string {
//	if s == "" {
//		return "N/A"
//	}
//	return s
//}

//
//func GetDiskInfos2() ([]DiskInfo, error) {
//	block, err := ghw.Block()
//	if err != nil {
//		return nil, err
//	}
//
//	diskInfos := make([]DiskInfo, len(block.Disks))
//	f := func(p *ghw.Partition) string {
//		return p.Name + " " + units.BytesSize(float64(p.SizeBytes)) + " " +
//			placeholder(p.Label) + " " + placeholder(p.MountPoint)
//	}
//	for i, disk := range block.Disks {
//		diskInfos[i] = DiskInfo{
//			Name:       disk.Name,
//			DriveType:  disk.DriveType.String(),
//			Vendor:     disk.Vendor,
//			Partitions: strings.Join(funk.Map(disk.Partitions, f).([]string), "\n"),
//			Total:      units.BytesSize(float64(disk.SizeBytes)),
//		}
//
//	}
//
//	return diskInfos, nil
//}

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
