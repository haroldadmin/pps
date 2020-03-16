package main

import (
	"flag"
	"fmt"
	"sync"
	"time"

	"github.com/haroldadmin/pps/scanner"
)

const numberOfPorts int = 1000

type config struct {
	hostname    string
	timeout     time.Duration
	parallelism int
}

func main() {
	fmt.Println("--------------------")
	fmt.Println("Parallel Port Scanner")

	config := parseFlags()

	fmt.Printf("Scanning %s with %d goroutines\n", config.hostname, config.parallelism)
	defer fmt.Println("--------------------")

	portsChan := make(chan int, numberOfPorts)
	resultsChan := make(chan *scanner.PortScanResult, numberOfPorts)

	var wg sync.WaitGroup

	for port := 1; port <= numberOfPorts; port++ {
		portsChan <- port
		wg.Add(1)
	}

	for i := 0; i < config.parallelism; i++ {
		go func() {
			for {
				port, ok := <-portsChan
				if !ok {
					return
				}
				resultsChan <- runScan(config.hostname, port, config.timeout)
				wg.Done()
			}
		}()
	}

	wg.Wait()

	isAnyPortOpen := false

	for i := 0; i < numberOfPorts; i++ {
		result := <-resultsChan
		if result.IsOpen {
			isAnyPortOpen = true
			fmt.Printf("Open: Port %d\n", result.Port)
		}
	}

	if !isAnyPortOpen {
		fmt.Println("No open ports found.")
	}

}

func parseFlags() config {
	hostname := flag.String("hostname", "localhost", "The Host to scan for open ports")
	timeout := flag.Int("timeout", 1, "The timeout duration in seconds for a port scan")
	parallelism := flag.Int("parallelism", numberOfPorts, "The number of ports to be scanned simultaneously")

	flag.Parse()

	return config{
		hostname:    *hostname,
		timeout:     (time.Duration(*timeout) * time.Second),
		parallelism: *parallelism,
	}
}

func runScan(hostname string, port int, timeout time.Duration) *scanner.PortScanResult {
	result := scanner.ScanTCP(hostname, port, timeout)
	return &result
}
