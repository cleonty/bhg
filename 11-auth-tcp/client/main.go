package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile)
	var (
		err        error
		cert       tls.Certificate
		serverCert []byte
		pool       *x509.CertPool
		tlsConf    *tls.Config
	)
	if cert, err = tls.LoadX509KeyPair("clientCrt.pem", "clientKey.pem"); err != nil {
		log.Fatalln(err)
	}
	if serverCert, err = ioutil.ReadFile("../server/serverCrt.pem"); err != nil {
		log.Fatalln(err)
	}
	pool = x509.NewCertPool()
	pool.AppendCertsFromPEM(serverCert)
	tlsConf = &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      pool,
	}
	tlsConf.BuildNameToCertificate()

	conn, err := tls.Dial("tcp", "localhost:443", tlsConf)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	n, err := conn.Write([]byte("hello\n"))
	if err != nil {
		log.Println(n, err)
		return
	}

	buf := make([]byte, 100)
	n, err = conn.Read(buf)
	if err != nil {
		log.Println(n, err)
		return
	}

	println(string(buf[:n]))
}
