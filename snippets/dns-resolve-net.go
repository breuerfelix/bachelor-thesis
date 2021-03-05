package main

import (
	"net"
	"fmt"
)

func main() {
	ips, _ := net.LookupIP("cluster.hivemq.internal")

	for _, ip := range ips {
		fmt.Printf("IP: %s\n", ip.String())
	}
}

// Output:
// IP: 192.168.0.2
// IP: 192.168.0.3
// IP: 192.168.0.4
