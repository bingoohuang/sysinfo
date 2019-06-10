package main

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/table"
)

type TablePrinter struct {
	dittoMark string
}

func printTable(dittoMark string) {
	info := GetSysInfo()

	printer := TablePrinter{dittoMark: dittoMark}
	printer.tableHost(info.HostInfo)
	printer.tableMem(info.MemInfo)
	printer.tableCPUInfos(info.CPUInfos)
	printer.tableDiskInfos(info.DiskInfos)
	printer.tableInterfInfos(info.InterfInfos)
	printer.tableErrors(info.Errors)
}

func (p TablePrinter) tableErrors(errs []string) {
	if len(errs) == 0 {
		return
	}

	rows := make([]table.Row, len(errs))
	for i, c := range errs {
		rows[i] = table.Row{i + 1, c}
	}
	p.TableRender(table.Row{"#", "Error"}, rows...)
	fmt.Println()
}

func (p TablePrinter) TableRender(header table.Row, rows ...table.Row) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(header)

	if p.dittoMark != "" {
		t.AppendRows(p.DittoMarkRows(rows))
	} else {
		t.AppendRows(rows)
	}
	t.Render()
}

func (p TablePrinter) DittoMarkRows(rows []table.Row) []table.Row {
	mark := make(map[int]interface{}, 0)
	for i, row := range rows {
		for j, cell := range row {
			v, ok := mark[j]
			if ok && v != "" && v == cell {
				rows[i][j] = p.dittoMark
			} else {
				mark[j] = cell
			}
		}
	}

	return rows
}
