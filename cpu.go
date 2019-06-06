package main

import (
	"github.com/shirou/gopsutil/cpu"
)

type CPUInfo struct {
	VendorID    string
	Family      string
	ModelName   string
	Cores       int32
	IndexNumber int32
	Speed       float64
}

func GetCPUInfo() ([]CPUInfo, error) {
	// cpu - get CPU number of cores and speed
	cpuStats, err := cpu.Info()
	if err != nil {
		return nil, err
	}

	cpuInfos := make([]CPUInfo, len(cpuStats))

	for i, cpuStat := range cpuStats {
		cpuInfos[i] = CPUInfo{
			IndexNumber: cpuStat.CPU,
			VendorID:    cpuStat.VendorID,
			Family:      cpuStat.Family,
			ModelName:   cpuStat.ModelName,
			Cores:       cpuStat.Cores,
			Speed:       cpuStat.Mhz,
		}
	}

	return cpuInfos, err
}
