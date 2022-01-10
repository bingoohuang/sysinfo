module github.com/bingoohuang/sysinfo

go 1.14

require (
	github.com/bingoohuang/gg v0.0.0-20220107092152-7fa91d9ab879
	github.com/bingoohuang/gou v0.0.0-20200225004418-9b3655665c46
	github.com/docker/go-units v0.4.0
	github.com/gobars/cmd v0.0.0-20191114090003-c6a602977f49
	github.com/jedib0t/go-pretty v4.3.0+incompatible
	github.com/shirou/gopsutil/v3 v3.21.10
	github.com/thoas/go-funk v0.9.0
)

replace github.com/shirou/gopsutil/v3 => ../gopsutil
