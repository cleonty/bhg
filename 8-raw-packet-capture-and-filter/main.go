package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var (
	snaplen  = int32(1600)
	promisc  = false
	timeout  = pcap.BlockForever
	filter   = "tcp and dst port 8544"
	devFound = false
)

func main() {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Panicln(err)
	}
	var wg sync.WaitGroup
	wg.Add(len(devices))
	for _, device := range devices {
		log.Printf("Device: %s", device.Name)
		for _, address := range device.Addresses {
			log.Printf(" IP: %s  Netmask: %s", address.IP, address.Netmask)
		}
		go func(iface string) {
			defer wg.Done()
			handle, err := pcap.OpenLive(iface, snaplen, promisc, timeout)
			if err != nil {
				log.Panicln(err)
			}
			defer handle.Close()
			if err := handle.SetBPFFilter(filter); err != nil {
				log.Panicln(err)
			}
			source := gopacket.NewPacketSource(handle, handle.LinkType())
			for packet := range source.Packets() {
				fmt.Printf("%s: %s\n", iface, packet)
			}
		}(device.Name)
	}
	wg.Wait()
}
