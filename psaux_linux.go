package sysinfo

import (
	"regexp"
	"strconv"

	"github.com/bingoohuang/gou/os"
	"github.com/bingoohuang/gou/str"
)

func PsAuxTop(n int) ([]PsAuxItem, error) {
	auxItems := make([]PsAuxItem, 0)
	opt := str.If(n > 0, `--sort=-pcpu|head -n `+strconv.Itoa(n), `--sort=-pid --forest`)

	re := regexp.MustCompile(`\s+`)
	err := os.ExecuteBashLiner(prefix+opt+fixedLtime, func(line string) bool {
		f := re.Split(line, 13)
		auxItems = append(auxItems, PsAuxItem{
			User: f[2], Pid: str.ParseInt(f[3]), Ppid: str.ParseInt(f[4]), CPU: str.ParseFloat32(f[5]),
			Mem: str.ParseFloat32(f[6]), Vsz: str.ParseInt(f[7]), Rss: str.ParseInt(f[8]),
			Tty: f[9], Stat: f[10], Start: f[0] + ` ` + f[1], Time: f[11], Command: f[12]})
		return true
	})

	return auxItems, err
}

const prefix = `ps axo lstart,user,pid,ppid,pcpu,pmem,vsz,rss,tname,stat,time,args --no-heading `

// nolint
const fixedLtime = `|awk '{c="date -d\""$1 FS $2 FS $3 FS $4 FS $5"\" +\047%Y-%m-%d %H:%M:%S\047"; c|getline d; close(c); $1=$2=$3=$4=$5=""; printf "%s\n",d$0 }'`

// nolint
/*
https://unix.stackexchange.com/questions/401785/ps-output-with-iso-date-format
[root@fs04-192-168-126-5 ~]# ps axo lstart,user,pid,ppid,pcpu,pmem,vsz,rss,tname,stat,time,args --no-heading
Fri Jul  5 07:12:47 2019 root         1     0  0.0  0.0 194164  7268 ?        Ss   00:00:23 /usr/lib/systemd/systemd --switched-root --system --deserialize 22
Fri Jul  5 07:12:47 2019 root         2     0  0.0  0.0      0     0 ?        S    00:00:00 [kthreadd]
Fri Jul  5 07:12:47 2019 root         3     2  0.0  0.0      0     0 ?        S    00:00:00 [ksoftirqd/0]
Fri Jul  5 07:12:47 2019 root         5     2  0.0  0.0      0     0 ?        S<   00:00:00 [kworker/0:0H]
[root@fs04-192-168-126-5 ~]# date -d "Fri Jul  5 07:12:47 2019" +"%Y-%m-%d %H:%M:%S"
2019-07-05 07:12:47
*/
