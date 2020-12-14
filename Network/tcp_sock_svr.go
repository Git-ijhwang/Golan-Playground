package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const (
	connHost = "localhost"
	connPort = "8080"
	connType = "tcp"
)

func main() {
	fmt.Println("Starting " + connType + "server on " + connHost + ":" + connPort)

	/* Open the Listen Fd */
	listenFd, err := net.Listen(connType, connHost+":"+connPort)
	if err != nil {
		fmt.Println("Error listening:", err, Error())
		os.Exit(1)
	}

	defer listenFd.Close()

	/* Accept Connections from Clients */
	for {
		c, err := listenFd.Accept()
		if err != nil {
			fmt.Println("Error connection:", error.Error())
			return
		}

		fmt.Println("Client connected.")
		fmt.Println("Client " + c.RemoteAddr().String() + " Connected")

		go handleConnection(c)
	}
}

func handleConnection(conn net.Conn) {
	buffer, err := bufio.NewReader(conn).ReadBytes('\n')

	if err != nil {
		fmt.Println("Client left.")
		conn.Close()
		return
	}

	log.Println("Client message:", string(buffer[:len(buffer)-1]))
	conn.Write(buffer)
	handleConnection(conn)
}
