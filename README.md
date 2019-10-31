# sysinfo
a command to show the system info as full as possible

## install:

`go get github.com/bingoohuang/sysinfo/...`

## usage: 

```bash
➜  sysinfo -h
Usage of sysinfo:
  -ditto string
    	ditto mark (same as above (default "\"")
  -format string
    	display format json/table (default "table")
  -show string
    	only show specified info(host,mem,cpu,disk,interf,ps) (default "host,mem,cpu,disk,interf")
```

## demo:

```bash
➜  sysinfo
+---+----------+--------+--------------+-------+--------+----------+--------------------------------------+------------------+----------------+
| # | HOSTNAME | UPTIME | UPTIME HUMAN | PROCS | OS     | PLATFORM | HOST ID                              | PLATFORM VERSION | KERNEL VERSION |
+---+----------+--------+--------------+-------+--------+----------+--------------------------------------+------------------+----------------+
| 1 | bogon    | 114161 | 32 hours     |   415 | darwin | darwin   | 7c8bb636-e593-3ce4-8528-9bd24a688851 | 10.14.5          | 18.6.0         |
+---+----------+--------+--------------+-------+--------+----------+--------------------------------------+------------------+----------------+

+---+-------+----------+-----------------+
| # | TOTAL | FREE     | USED PERCENTAGE |
+---+-------+----------+-----------------+
| 1 | 16GiB | 770.7MiB | 59.88%          |
+---+-------+----------+-----------------+

+---+-------------+--------------+--------+------------------------------------------+-------+------+
| # | PHYSICAL ID | VENDOR ID    | FAMILY | MODEL NAME                               | CORES |  MHZ |
+---+-------------+--------------+--------+------------------------------------------+-------+------+
| 1 |             | GenuineIntel | 6      | Intel(R) Core(TM) i7-8750H CPU @ 2.20GHz |     6 | 2200 |
+---+-------------+--------------+--------+------------------------------------------+-------+------+

+---+-----------------+--------------+--------+----------+----------+----------+--------------+
| # | PATH            | DEVICE       | FSTYPE | TOTAL    | USED     | FREE     | USED PERCENT |
+---+-----------------+--------------+--------+----------+----------+----------+--------------+
| 1 | /               | /dev/disk1s1 | apfs   | 233.5GiB | 149.1GiB | 81.72GiB | 64.59%       |
| 2 | /private/var/vm | /dev/disk1s4 | "      | "        | 2GiB     | "        | 02.39%       |
+---+-----------------+--------------+--------+----------+----------+----------+--------------+

+---+----------------+-------------------+----------------------------------------------+
| # | INTERFACE NAME | HARDWARE ADDR     | ADDRS                                        |
+---+----------------+-------------------+----------------------------------------------+
| 1 | en5            | ac:de:48:00:11:22 | fe80::aede:48ff:fe00:1122/64                 |
| 2 | en0            | f0:18:98:a5:12:27 | fe80::49:dfcf:6bf2:7a96/64 192.168.162.95/24 |
| 3 | awdl0          | 46:20:fa:ed:57:2c | fe80::4420:faff:feed:572c/64                 |
| 4 | en8            | 00:e0:4c:68:1d:97 | fe80::4d:5ead:7376:4e97/64 192.168.217.47/24 |
+---+----------------+-------------------+----------------------------------------------+
```

```bash
# ./sysinfo 
+---+-------------+---------+--------------+-------+-------+-------------+--------------------------------------+------------------+-----------------------+
| # | HOSTNAME    |  UPTIME | UPTIME HUMAN | PROCS | OS    | PLATFORM    | HOST ID                              | PLATFORM VERSION | KERNEL VERSION        |
+---+-------------+---------+--------------+-------+-------+-------------+--------------------------------------+------------------+-----------------------+
| 1 | BJCA-device | 1023995 | 11 days      |   215 | linux | bjca-device | fbe926a2-ade0-5bb0-9e37-04d4c439354d | 3.0              | 3.10.0-862.el7.x86_64 |
+---+-------------+---------+--------------+-------+-------+-------------+--------------------------------------+------------------+-----------------------+

+---+----------+----------+-----------------+
| # | TOTAL    | FREE     | USED PERCENTAGE |
+---+----------+----------+-----------------+
| 1 | 7.572GiB | 5.552GiB | 17.04%          |
+---+----------+----------+-----------------+

+---+-------------+--------------+--------+-------------------------------------------+-------+------+
| # | PHYSICAL ID | VENDOR ID    | FAMILY | MODEL NAME                                | CORES |  MHZ |
+---+-------------+--------------+--------+-------------------------------------------+-------+------+
| 1 | 0           | GenuineIntel | 6      | Intel(R) Xeon(R) CPU E3-1230 v6 @ 3.50GHz |     8 | 3900 |
+---+-------------+--------------+--------+-------------------------------------------+-------+------+

+---+-------+--------------------------------------+--------+----------+----------+----------+--------------+
| # | PATH  | DEVICE                               | FSTYPE | TOTAL    | USED     | FREE     | USED PERCENT |
+---+-------+--------------------------------------+--------+----------+----------+----------+--------------+
| 1 | /     | /dev/mapper/centos_bjca--device-root | xfs    | 899.6GiB | 1.823GiB | 897.7GiB | 00.20%       |
| 2 | /boot | /dev/sda1                            | "      | 1014MiB  | 141.9MiB | 872.1MiB | 14.00%       |
+---+-------+--------------------------------------+--------+----------+----------+----------+--------------+

+---+----------------+-------------------+-----------------------------------------------+
| # | INTERFACE NAME | HARDWARE ADDR     | ADDRS                                         |
+---+----------------+-------------------+-----------------------------------------------+
| 1 | team0          | 04:d4:c4:39:35:4e | 192.168.221.46/24 fe80::6d4:c4ff:fe39:354e/64 |
+---+----------------+-------------------+-----------------------------------------------+
```

```bash
$ ./sysinfo 
+---+----------------------+---------+--------------+-------+-------+----------+--------------------------------------+------------------+-----------------------+
| # | HOSTNAME             |  UPTIME | UPTIME HUMAN | PROCS | OS    | PLATFORM | HOST ID                              | PLATFORM VERSION | KERNEL VERSION        |
+---+----------------------+---------+--------------+-------+-------+----------+--------------------------------------+------------------+-----------------------+
| 1 | fs01-192-168-126-182 | 6149498 | 2 months     |   422 | linux | centos   | 86f6047a-3ed8-4e0c-a486-fafd068e63ba | 7.5.1804         | 3.10.0-862.el7.x86_64 |
+---+----------------------+---------+--------------+-------+-------+----------+--------------------------------------+------------------+-----------------------+

+---+----------+----------+-----------------+
| # | TOTAL    | FREE     | USED PERCENTAGE |
+---+----------+----------+-----------------+
| 1 | 62.37GiB | 21.37GiB | 19.46%          |
+---+----------+----------+-----------------+

+---+-------------+--------------+--------+-------------------------------------------+-------+------+
| # | PHYSICAL ID | VENDOR ID    | FAMILY | MODEL NAME                                | CORES | MHZ  |
+---+-------------+--------------+--------+-------------------------------------------+-------+------+
| 1 | 0           | GenuineIntel | 6      | Intel(R) Xeon(R) CPU E5-2603 v4 @ 1.70GHz | 6     | 1700 |
| 2 | 1           | "            | "      | "                                         | "     | "    |
+---+-------------+--------------+--------+-------------------------------------------+-------+------+

+---+-------+-------------------------+--------+----------+----------+----------+--------------+
| # | PATH  | DEVICE                  | FSTYPE | TOTAL    | USED     | FREE     | USED PERCENT |
+---+-------+-------------------------+--------+----------+----------+----------+--------------+
| 1 | /     | /dev/mapper/centos-root | xfs    | 49.98GiB | 14.09GiB | 35.89GiB | 28.19%       |
| 2 | /boot | /dev/sda2               | "      | 1014MiB  | 168.9MiB | 845.1MiB | 16.66%       |
| 3 | /home | /dev/mapper/centos-home | "      | 2.098TiB | 99.54GiB | 2.001TiB | 04.63%       |
+---+-------+-------------------------+--------+----------+----------+----------+--------------+

+----+-----------------+-------------------+------------------------------------------------+
|  # | INTERFACE NAME  | HARDWARE ADDR     | ADDRS                                          |
+----+-----------------+-------------------+------------------------------------------------+
|  1 | enp2s0f0        | 60:f1:8a:45:cc:fc | 192.168.126.182/24 fe80::e5b:aa94:f0a6:f653/64 |
|  2 | virbr0          | 52:54:00:df:61:5f | 192.168.122.1/24                               |
|  3 | docker0         | 02:42:37:bc:3a:23 | 172.17.0.1/16 fe80::42:37ff:febc:3a23/64       |
|  4 | br-71e483483644 | 02:42:ed:5f:f1:2d | 172.18.0.1/16 fe80::42:edff:fe5f:f12d/64       |
|  5 | br-12a39463ab93 | 02:42:aa:e0:18:c0 | 172.19.0.1/16                                  |
|  6 | br-83a77a912773 | 02:42:e2:5b:77:96 | 172.20.0.1/16 fe80::42:e2ff:fe5b:7796/64       |
|  7 | veth1653d1a     | c6:d0:18:97:30:d9 | fe80::c4d0:18ff:fe97:30d9/64                   |
|  8 | vethae7f85b     | b6:4c:88:b6:44:0e | fe80::b44c:88ff:feb6:440e/64                   |
|  9 | veth6d46af4     | 06:ca:eb:79:3d:65 | fe80::4ca:ebff:fe79:3d65/64                    |
| 10 | veth732cbb9     | 3a:94:7e:ba:b7:4c | fe80::3894:7eff:feba:b74c/64                   |
| 11 | veth3839eb8     | d6:06:be:9c:32:47 | fe80::d406:beff:fe9c:3247/64                   |
| 12 | veth847cd7d     | 92:fa:2c:4b:b3:9e | fe80::90fa:2cff:fe4b:b39e/64                   |
| 13 | vetheab1d14     | f6:88:50:6d:2e:f4 | fe80::f488:50ff:fe6d:2ef4/64                   |
| 14 | vetha7b3a21     | 7a:f5:62:c5:a3:8f | fe80::78f5:62ff:fec5:a38f/64                   |
| 15 | vethca8a5bc     | b2:b1:a3:da:b8:ff | fe80::b0b1:a3ff:feda:b8ff/64                   |
| 16 | vethe4524f4     | 8e:f7:4b:b8:3c:6c | fe80::8cf7:4bff:feb8:3c6c/64                   |
| 17 | veth896a936     | e2:67:23:66:a3:22 | fe80::e067:23ff:fe66:a322/64                   |
+----+-----------------+-------------------+------------------------------------------------+
```
## thanks

1. https://www.socketloop.com/tutorials/golang-get-hardware-information-such-as-disk-memory-and-cpu-usage
1. https://github.com/jedib0t/go-pretty
1. https://github.com/zcalusic/sysinfo
1. https://github.com/jaypipes/ghw
1. [goreleaser](https://goreleaser.com/)

## goreleaser usage

```bash
# clone it outside GOPATH
git clone https://github.com/goreleaser/goreleaser
cd goreleaser

# get dependencies using go modules (needs go 1.11+)
go get ./...

# build
go install

# check it works
goreleaser --version

# Run goreleaser init to create an example .goreleaser.yaml file:
goreleaser init

# test the configuration at any time
goreleaser --snapshot --skip-publish --rm-dist

#  export either a GITHUB_TOKEN or GITLAB_TOKEN environment variable, 
#  which should contain a valid GitHub token with the repo scope or GitLab token with api scope. 
#  It will be used to deploy releases to your GitHub/GitLab repository.
export GITHUB_TOKEN="YOUR_GH_TOKEN"
# or
export GITLAB_TOKEN="YOUR_GL_TOKEN"

# GoReleaser will use the latest Git tag of your repository. Create a tag and push it to GitHub:
$ git tag -a v0.1.0 -m "First release"
$ git push origin v0.1.0

$ v=1.5.3; git tag -a v$v -m "v$v"; git push origin v$v;

# run GoReleaser at the root of your repository:
goreleaser --rm-dist
```

## [how to delete a git tag locally and remote](https://gist.github.com/mobilemind/7883996)

```bash
# delete local tag v0.1.0
git tag -d v0.1.0
# delete remote tag v0.1.0 (eg, GitHub version too)
git push --delete origin v0.1.0

v=0.1.0; git tag -d v$v; git push --delete origin v$v;
```
