package main

import (
	"fmt"
	"net"
)

func main() {
	networkType := "tcp"
	Service := "FTP"

	p, _ := net.LookupPort(networkType, Service)
	fmt.Printf("Port associated with %s is %d\n", Service, p)
}
