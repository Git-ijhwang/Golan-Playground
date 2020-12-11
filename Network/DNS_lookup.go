package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	Host := "www.google.com"
	serverAddress, err := net.ResolveIPAddr("ip", Host)
	if err != nil {
		fmt.Println("Couldn't Resolve")
		os.Exit(1)
	}
	fmt.Println(serverAddress.String())
}
