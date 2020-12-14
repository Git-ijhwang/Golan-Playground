package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	//ip := "192.168.69.59

	ip := "127.0.0.1"
	port := "8080"

	address := ip + ":" + port

	client_conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("[+] Unable to connect")
		panic(err)
	}

	server_ip := client_conn.RemoteAddr().String()

	fmt.Println("[+] Server IP is ", server_ip)
	fmt.Println("[+] Enter message to send :")

	input_reader := bufio.NewReader(os.Stdin)

	input_msg, err := input_reader.ReadString('\n')
	if err != nil {
		fmt.Println("[+] Error occured")
		panic(err)
	}

	client_conn.Write([]byte(input_msg))

}
