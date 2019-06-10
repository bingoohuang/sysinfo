package main

import (
	"encoding/json"
	"flag"
	"os"
)

func main() {
	format := flag.String("format", "table", "display format json/table")
	ditto := flag.String("ditto", "〃", "ditto mark (same as above")
	flag.Parse()

	switch *format {
	case "table":
		printTable(*ditto)
	case "json":
		fallthrough
	default:
		b, _ := json.MarshalIndent(GetSysInfo(), "", "    ")
		_, _ = os.Stdout.Write(b)
	}
}
