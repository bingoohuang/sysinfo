package sysinfo

import (
	"encoding/json"
	"fmt"
	"github.com/bingoohuang/gou/reflec"
	"github.com/jedib0t/go-pretty/v6/table"
	"io"
	"os"
	"reflect"
	"regexp"
)

// TablePrinter ...
type TablePrinter struct {
	dittoMark string
	Out       io.Writer
	Format    string
}

// PrintTable ...
func PrintTable(showsMap map[string]bool, dittoMark string, out io.Writer, format string) {
	info := GetSysInfo(showsMap)

	p := TablePrinter{dittoMark: dittoMark, Out: out, Format: format}

	if p.Out == nil {
		p.Out = os.Stdout
	}

	p.table(info.HostInfo)
	p.table(info.MemInfo)
	p.table(info.CPUInfos)
	p.table(info.CPU)
	p.table(info.DiskInfos)
	p.table(info.InterfInfos)
	p.table(info.PsItems)
	p.table(info.Errors)
}

func (p TablePrinter) table(value interface{}) {
	header := make(table.Row, 0)
	rows := make([]table.Row, 0)

	header = append(header, "#")

	v := reflect.ValueOf(value)
	if v.IsNil() {
		return
	}

	switch v.Kind() {
	case reflect.Ptr:
		v = v.Elem()
		fallthrough
	case reflect.Struct:
		fields := reflec.CachedStructFields(v.Type(), "header")
		createHeader(fields, &header)
		createRow(fields, 0, v, &rows)
	case reflect.Slice:
		if v.Len() == 0 {
			return
		}

		fields := reflec.CachedStructFields(v.Type().Elem(), "header")
		createHeader(fields, &header)

		for i := 0; i < v.Len(); i++ {
			createRow(fields, i, v.Index(i), &rows)
		}
	default:
		return
	}

	p.tableRender(header, rows...)
	fmt.Println()
}

func createRow(fields []reflec.StructField, rowIndex int, v reflect.Value, rows *[]table.Row) {
	row := make(table.Row, 0)
	row = append(row, rowIndex+1)

	for _, f := range fields {
		if !IsCapital(f.Name) {
			continue
		}

		val := v.Field(f.Index).Interface()
		if f.Kind == reflect.Struct {
			j, _ := json.Marshal(val)
			val = string(j)
		}

		row = append(row, val)
	}

	*rows = append(*rows, row)
}

func createHeader(fields []reflec.StructField, header *table.Row) {
	for _, f := range fields {
		if IsCapital(f.Name) {
			*header = append(*header, BlankCamel(f.Name))
		}
	}
}

func IsCapital(s string) bool {
	if len(s) == 0 {
		return false
	}
	c := s[0]
	return c >= 'A' && c <= 'Z'
}

func (p TablePrinter) tableRender(header table.Row, rows ...table.Row) {
	t := table.NewWriter()
	t.SetOutputMirror(p.Out)
	t.AppendHeader(header)

	if p.dittoMark != "" {
		t.AppendRows(p.dittoMarkRows(rows))
	} else {
		t.AppendRows(rows)
	}

	switch p.Format {
	case "markdown":
		t.RenderMarkdown()
	default:
		t.Render()
	}
}

func (p TablePrinter) dittoMarkRows(rows []table.Row) []table.Row {
	mark := make(map[int]interface{})

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

// BlankCamel ...
func BlankCamel(str string) string {
	blank := regexp.MustCompile("(.)([A-Z][a-z]+)").ReplaceAllString(str, "${1} ${2}")
	return regexp.MustCompile("([a-z0-9])([A-Z])").ReplaceAllString(blank, "${1} ${2}")
}
