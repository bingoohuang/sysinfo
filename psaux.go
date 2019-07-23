package sysinfo

type PsAuxItem struct {
	User    string
	Pid     int
	Ppid    int
	CPU     float32
	Mem     float32
	Vsz     int
	Rss     int
	Tty     string
	Stat    string
	Start   string
	Time    string
	Command string
}
