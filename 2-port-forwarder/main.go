package main

import "net"
import "log"
import "io"



func handle(src net.Conn, address string) {
	dst, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatalf("unable to connect to %s: %v", address, err)
	}
	defer dst.Close()
	
	go func () {
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()
	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	address := "www.e1.ru:443"
	listener, err := net.Listen("tcp", ":443")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
			}
			go handle(conn, address)
	}
}
