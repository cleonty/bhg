package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for port := 0; port < 65000; port++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			address := fmt.Sprintf("127.0.0.1:%d", port)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				fmt.Printf("port %3d closed or filtered\n", port)
				return
			}
			fmt.Printf("connection to %s successfull\n", address)
			conn.Close()
		}(port)
	}
	wg.Wait()
}
