package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"sort"
)

const nPorts = 1024

type ScanResult struct {
	port   int
	isOpen bool
}

func main() {
	host := "scanme.nmap.org"
	if len(os.Args) > 1 {
		host = os.Args[1]
	}
	log.Printf("scanning %s\n", host)
	ports := make(chan int, 100)
	results := make(chan ScanResult)
	var openPorts []int

	for i := 0; i < cap(ports); i++ {
		go worker(host, ports, results)
	}
	go func() {
		for port := 1; port <= nPorts; port++ {
			ports <- port
		}
	}()
	for port := 1; port <= nPorts; port++ {
		result := <-results
		if result.isOpen {
			openPorts = append(openPorts, result.port)
		}
	}
	close(ports)
	close(results)
	sort.Ints(openPorts)
	for _, port := range openPorts {
		log.Printf("port %3d open\n", port)
	}
}

func worker(host string, ports <-chan int, results chan<- ScanResult) {
	for port := range ports {
		isOpen := scan(host, port)
		result := ScanResult{
			port:   port,
			isOpen: isOpen,
		}
		results <- result
	}
}

func scan(host string, port int) bool {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}
