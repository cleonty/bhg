package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const Dir = "./hello/world"

func Hello(dir string) {
	os.RemoveAll(dir)
	os.MkdirAll(dir, 0755)
}

func SetupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("[-] Ctrl+C pressed in Terminal")
		os.RemoveAll(Dir)
		os.Exit(0)
	}()
}

func main() {
	SetupCloseHandler()
	for {
		Hello(Dir)
		time.Sleep(10 * time.Second)
	}
}
