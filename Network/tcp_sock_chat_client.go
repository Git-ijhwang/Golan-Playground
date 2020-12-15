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

	go send_task(conn, reader)
	go recv_task(conn)
	for {
		//fmt.Print("Text to send:")
		//input, _ := reader.ReadString('\n')

		//conn.Write([]byte(input))

		//n, _ := conn.Read(buf[0:])
		//fmt.Print("Received Message from Server:")
		//fmt.Println(string(buf[0:n]))
	}
}

func recv_task(conn net.Conn) {
	var buf [512]byte
	for {
		//fmt.Print("Text to send:")
		//input, _ := reader.ReadString('\n')

		//conn.Write([]byte(input))

		n, _ := conn.Read(buf[0:])
		fmt.Print("Received Message from Server:")
		fmt.Println(string(buf[0:n]))
	}
}
func send_task(conn net.Conn, reader *bufio.Reader) {
	for {
		fmt.Print("Text to send:")
		input, _ := reader.ReadString('\n')
		conn.Write([]byte(input))
	}

}
