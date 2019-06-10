package main

import (
	"fmt"

	"github.com/jedib0t/go-pretty/table"

	"github.com/docker/go-units"
	"github.com/shirou/gopsutil/mem"
)

type MemInfo struct {
	Total          string
	Free           string
	UsedPercentage string
}

func (p TablePrinter) tableMem(m MemInfo) {
	p.TableRender(table.Row{"Total Mem", "Free Mem", "Mem Used Percentage"},
		table.Row{m.Total, m.Free, m.UsedPercentage})
	fmt.Println()
}

func GetMemInfo() (MemInfo, error) {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return MemInfo{}, err
	}

	return MemInfo{
		Total:          units.BytesSize(float64(vmStat.Total)),
		Free:           units.BytesSize(float64(vmStat.Free)),
		UsedPercentage: fmt.Sprintf("%.2f%%", vmStat.UsedPercent),
	}, nil
}
