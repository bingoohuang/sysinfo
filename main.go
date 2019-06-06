package main

import (
	"encoding/json"
	"flag"
	"os"
)

func main() {
	format := flag.String("format", "table", "display format json/table")
	flag.Parse()

	switch *format {
	case "table":
		printTable()
	case "json":
		fallthrough
	default:
		b, _ := json.MarshalIndent(GetSysInfo(), "", "    ")
		_, _ = os.Stdout.Write(b)
	}
}
