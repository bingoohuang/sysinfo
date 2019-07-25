package sysinfo

import (
	"regexp"

	"github.com/bingoohuang/cmd"
	"github.com/bingoohuang/gou/str"
	"github.com/docker/go-units"
)

type PsAuxItem struct {
	User    string
	Pid     int
	Ppid    int
	CPU     float32
	Mem     float32
	Vsz     string
	Rss     string
	Tty     string
	Stat    string
	Start   string
	Time    string
	Command string
}

func PsAuxTop(n int) ([]PsAuxItem, error) {
	auxItems := make([]PsAuxItem, 0)
	re := regexp.MustCompile(`\s+`)
	_, status := cmd.BashLiner(PasAuxShell(n, false), func(line string) bool {
		f := re.Split(line, 13)
		auxItems = append(auxItems, PsAuxItem{
			User:    f[2],
			Pid:     str.ParseInt(f[3]),
			Ppid:    str.ParseInt(f[4]),
			CPU:     str.ParseFloat32(f[5]),
			Mem:     str.ParseFloat32(f[6]),
			Vsz:     units.BytesSize(str.ParseFloat64(f[7])),
			Rss:     units.BytesSize(str.ParseFloat64(f[8])),
			Tty:     f[9],
			Stat:    f[10],
			Start:   f[0] + ` ` + f[1],
			Time:    f[11],
			Command: f[12]})
		return true
	})

	return auxItems, status.Error
}

func PasAuxShell(topN int, heading bool) string {
	return prefix + str.If(heading, "", noheading) + psAuxTopOpt(topN) + fixedLtime
}
