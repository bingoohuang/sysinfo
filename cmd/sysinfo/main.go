package main

import (
	"flag"
	"os"
	"strings"

	"github.com/bingoohuang/gg/pkg/codec"

	"github.com/bingoohuang/sysinfo"
)

func main() {
	show := flag.String("show", "host,mem,cpu,disk,interf", "only show specified info(host,mem,cpu,disk,interf,ps)")
	format := flag.String("format", "table", "display format json/table/markdown")
	ditto := flag.String("ditto", `"`, "ditto mark (same as above")
	flag.Parse()

	showsMap := make(map[string]bool)

	for _, p := range strings.Split(*show, ",") {
		p = strings.TrimSpace(p)
		if p != "" {
			showsMap[strings.ToLower(p)] = true
		}
	}

	switch *format {
	case "table", "markdown", "md":
		sysinfo.PrintTable(showsMap, *ditto, os.Stdout, *format)
	case "json", "js":
		fallthrough
	default:
		b := codec.Json(sysinfo.GetSysInfo(showsMap))
		_, _ = os.Stdout.Write(b)
	}
}
