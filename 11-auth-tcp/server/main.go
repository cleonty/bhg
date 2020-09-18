package main

import (
	"bufio"
	"crypto/tls"
	"crypto/x509"
	"io"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	var (
		err        error
		clientCert []byte
		serverCert tls.Certificate
		pool       *x509.CertPool
		tlsConf    *tls.Config
	)
	log.SetFlags(log.Lshortfile)

	if clientCert, err = ioutil.ReadFile("../client/clientCrt.pem"); err != nil {
		log.Fatalln(err)
	}
	if serverCert, err = tls.LoadX509KeyPair("serverCrt.pem", "serverKey.pem"); err != nil {
		log.Fatalln(err)
	}
	pool = x509.NewCertPool()
	pool.AppendCertsFromPEM(clientCert)
	tlsConf = &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientCAs:    pool,
		ClientAuth:   tls.RequireAndVerifyClientCert,
	}
	tlsConf.BuildNameToCertificate()

	ln, err := tls.Listen("tcp", "localhost:443", tlsConf)
	if err != nil {
		log.Println(err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	for {
		msg, err := r.ReadString('\n')
		if err == io.EOF {
			log.Println("client disconnected")
			return
		}
		if err != nil {
			log.Println(err)
			return
		}

		log.Println("client:", msg)

		n, err := conn.Write([]byte("world\n"))
		if err != nil {
			log.Println(n, err)
			return
		}
	}
}
