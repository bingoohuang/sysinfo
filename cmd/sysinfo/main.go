package main

import (
	"flag"
	"os"

	"github.com/bingoohuang/gou/enc"
	"github.com/bingoohuang/sysinfo"
)

func main() {
	format := flag.String("format", "table", "display format json/table")
	ditto := flag.String("ditto", `"`, "ditto mark (same as above")
	flag.Parse()

	switch *format {
	case "table":
		sysinfo.PrintTable(*ditto)
	case "json":
		fallthrough
	default:
		b := enc.JSONPretty(sysinfo.GetSysInfo())
		_, _ = os.Stdout.WriteString(b)
	}
}
