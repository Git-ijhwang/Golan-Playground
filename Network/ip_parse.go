package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Invalid IP")
		os.Exit(1)
	}

	IP := os.Args[1]

	net.ParseIP(IP)

	address := net.ParseIP(IP)

	fmt.Println(address)
	fmt.Printf("%T\n", address)

	fmt.Println(address.String())
	fmt.Printf("%T\n", address.String())
}
