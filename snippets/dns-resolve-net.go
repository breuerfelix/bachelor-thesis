package main

import (
	"net"
	"fmt"
	"time"
)

var INTERVAL = 5 * time.Second

func refreshDNS(domain string) {
	for {
		ips, _ := net.LookupIP(domain)

		for _, ip := range ips {
			fmt.Printf("IP: %s\n", ip.String())
		}

		time.Sleep(INTERVAL)
	}
}

func main() {
	go refreshDNS("hivemq.fbr.ai")

	// ...
}

// Output:
// IP: 192.168.0.2
// IP: 192.168.0.3
// IP: 192.168.0.4
