package scanner

import (
	"fmt"
	"net"
	"time"
)

// PortScanResult contains information about the Port Scan
type PortScanResult struct {
	Port   int
	IsOpen bool
}

// Scan takes takes in a protocol, hostname, port and a timeout duration, and attempts to dial the assembled network address.
// If the dialup is successful, the port is considered open. If unsuccessful, the port is considered closed.
func Scan(protocol, hostname string, port int, timeout time.Duration) PortScanResult {
	addr := fmt.Sprintf("%s:%d", hostname, port)
	conn, err := net.DialTimeout(protocol, addr, timeout)

	info := PortScanResult{
		Port: port,
	}

	if err != nil {
		info.IsOpen = false
	} else {
		info.IsOpen = true
		defer conn.Close()
	}

	return info
}

// ScanTCP checks if the given port is open on the hostname, using the TCP protocol
func ScanTCP(hostname string, port int, timeout time.Duration) PortScanResult {
	return Scan("tcp", hostname, port, timeout)
}
