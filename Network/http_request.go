package main

import (
	"fmt"
	"io/ioutil"
	"net"
)

func main() {
	address := "www.google.com:80"
	connection, err := net.Dial("tcp", address)

	if err != nil {
		panic(err)
	}

	defer connection.Close()

	req := "HEAD / HTTP/1.0\r\n\r\n"

	_, err = connection.Write([]byte(req))
	if err != nil {
		panic(err)
	}

	response, err := ioutil.ReadAll(connection)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(response))

}
