package sysinfo

import (
	"sort"

	"github.com/shirou/gopsutil/v3/cpu"
)

// CPUInfo ...
type CPUInfo struct {
	PhysicalID string
	VendorID   string
	Family     string
	ModelName  string
	Cores      int
	Mhz        float64
}

// CPUInfoKey ...
type CPUInfoKey struct {
	PhysicalID string
	VendorID   string
	Family     string
	ModelName  string
	Mhz        float64
}

/**
GetCPUInfo 获得CPU信息
https://www.cnblogs.com/emanlee/p/3587571.html

# 总核数 = 物理CPU个数 X 每颗物理CPU的核数
# 总逻辑CPU数 = 物理CPU个数 X 每颗物理CPU的核数 X 超线程数

# 查看物理CPU个数
cat /proc/cpuinfo| grep "physical id"| sort| uniq| wc -l

# 查看每个物理CPU中core的个数(即核数)
cat /proc/cpuinfo| grep "cpu cores"| uniq

# 查看逻辑CPU的个数
cat /proc/cpuinfo| grep "processor"| wc -l
复制代码

查看CPU信息（型号）
cat /proc/cpuinfo | grep name | cut -f2 -d: | uniq -c
*/

// GetCPUInfo ...
func GetCPUInfo() ([]CPUInfo, error) {
	// cpu - get cpu number of cores and speed
	cpuStats, err := cpu.Info()
	if err != nil {
		return nil, err
	}

	cpuInfosMap := map[CPUInfoKey]int{}

	physicalMap := make(map[string]int)

	for _, cpuStat := range cpuStats {
		if cpuStat.PhysicalID != "" {
			if _, exists := physicalMap[cpuStat.PhysicalID]; exists {
				physicalMap[cpuStat.PhysicalID] += int(cpuStat.Cores)
				continue
			}
		}

		key := CPUInfoKey{
			PhysicalID: cpuStat.PhysicalID,
			VendorID:   cpuStat.VendorID,
			Family:     cpuStat.Family,
			ModelName:  cpuStat.ModelName,
			Mhz:        cpuStat.Mhz,
		}
		if _, ok := cpuInfosMap[key]; ok {
			cpuInfosMap[key] += int(cpuStat.Cores)
		} else {
			cpuInfosMap[key] = int(cpuStat.Cores)
		}

		if cpuStat.PhysicalID != "" {
			physicalMap[cpuStat.PhysicalID] = int(cpuStat.Cores)
		}
	}

	var cpuInfos []CPUInfo
	for k, v := range cpuInfosMap {
		cpuInfos = append(cpuInfos, CPUInfo{
			PhysicalID: k.PhysicalID,
			VendorID:   k.VendorID,
			Family:     k.Family,
			ModelName:  k.ModelName,
			Mhz:        k.Mhz,
			Cores:      v,
		})
	}
	sort.Slice(cpuInfos, func(i, j int) bool {
		return cpuInfos[i].PhysicalID < cpuInfos[j].PhysicalID
	})

	for i, c := range cpuInfos {
		if c.PhysicalID != "" {
			cpuInfos[i].Cores = physicalMap[c.PhysicalID]
		}
	}

	return cpuInfos, err
}
