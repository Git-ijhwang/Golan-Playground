package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	//connHost = "localhost"
	//connPort = "8080"
	connType = "tcp"
)

func main() {
	ip := os.Args[1]
	port := os.Args[2]
	fmt.Println("Starting " + connType + "server on " + ip + ":" + port)

	conn, err := net.Dial(connType, ip+":"+port)
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Text to send:")
		input, _ := reader.ReadString('\n')
		//log.Print("Server relay:", input)

		conn.Write([]byte(input))
	}
}
