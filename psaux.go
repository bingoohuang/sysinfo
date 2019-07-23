package sysinfo

type PsAuxItem struct {
	User    string
	Pid     string
	Ppid    string
	CPU     string
	Mem     string
	Vsz     string
	Rss     string
	Tty     string
	Stat    string
	Start   string
	Time    string
	Command string
}
