package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	server_ip := ""
	port := "8080"
	server_address := server_ip + ":" + port

	/* Listen */
	listner, err := net.Listen("tcp", server_address)
	if err != nil {
		fmt.Println("[+] Unable to listen")
		panic(err)
	}

	/* Accept */
	server_conn, err := listner.Accept()
	if err != nil {
		fmt.Println("[+] Unable to accept")
		panic(err)
	}
	client_ip := server_conn.RemoteAddr().String()

	fmt.Println("[+] Client IP is ", client_ip)

	conn_reader := bufio.NewReader(server_conn)
	client_msg, err := conn_reader.ReadString('\n')
	if err != nil {
		fmt.Println("[+] Error occured")
		panic(err)
	}

	fmt.Println("[+] client sent a message : ", client_msg)

}
