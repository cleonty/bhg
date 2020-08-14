package main

import (
	"fmt"

	"github.com/miekg/dns"
)

func main() {
	var msg dns.Msg
	fqdn := dns.Fqdn("66.ru")
	fmt.Printf("FQDN %s\n", fqdn)
	msg.SetQuestion(fqdn, dns.TypeA)
	in, err := dns.Exchange(&msg, "8.8.8.8:53")
	if err != nil {
		panic((err))
	}
	fmt.Printf("%v\n", in)
	if len(in.Answer) < 1 {
		fmt.Println("No records")
		return
	}
	for _, answer := range in.Answer {
		if a, ok := answer.(*dns.A); ok {
			fmt.Println(a.A)
		}
	}
}
