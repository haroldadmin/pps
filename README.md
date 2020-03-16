# Parallel Port Scanner

This is a simple port scanner written in Go to find open ports on a host. Ports are scanned parallelly.

## Installation

To get the tool, make sure you have Go installed on the your system, and can access the `go` command line too.

```bash
go get github.com/haroldadmin/pps
```

This will install Parallel Port Scanner in your Path, and you should be able to access it directly.

## Usage

Run the `pps` command with the `hostname` argument.

```bash
pps --hostname 192.168.1.2

--------------------
Parallel Port Scanner
Scanning 192.168.1.2 with 1000 goroutines
Open: Port 53
--------------------
```

Only ports 1-1000 are scanned on the host.

Optionally, you can also control the parallelism of the tool with the `parallelism` argument.

```bash
pps --hostname 192.168.1.2 --parallelism 4

--------------------
Parallel Port Scanner
Scanning 192.168.1.2 with 4 goroutines
Open: Port 53
--------------------
```

## Disclaimer

If it wasn't clear from the name of the tool already (*PPs*), this project meant solely as a fun learning experience. *I do not recommend using it as a tool in your network-utilities belt.*
