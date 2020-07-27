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

	b := make([]byte, 512)

	for {
		size, err := conn.Read(b[0:])
		if err == io.EOF {
			log.Printf("client %s disconnected", conn.RemoteAddr().String())
			break
		}
		if err != nil {
			log.Printf("unexpected error: %v", err)
			break
		}
		log.Printf("received %d bytes from %s: %s", size, conn.RemoteAddr().String(), string(b))
		log.Printf("writing data")
		if _, err := conn.Write(b[:size]); err != nil {
			log.Printf("unable to write data to client %s: %v", conn.RemoteAddr().String(), err)
			break
		}
	}
}
