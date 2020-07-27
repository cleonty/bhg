package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
)

func main() {
	port := 20080
	if os.Getenv("PORT") != "" {
		var err error
		port, err = strconv.Atoi(os.Getenv("PORT"))
		if err != nil {
			log.Fatalf("PORT env var is not integer: %s")
		}
	}
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("unable to listen on port %d", port)
	}
	log.Printf("Listening on port %d", port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("unable to accept connection: %v", err)
		}
		log.Printf("received connection %s", conn.RemoteAddr().String())
		go echo(conn)
	}
}

func echo(conn net.Conn) {
	defer conn.Close()

	for {
		if _, err := io.Copy(conn, conn); err != nil {
			log.Printf("unable to copy data: %v", err)
			break
		}
	}
}
