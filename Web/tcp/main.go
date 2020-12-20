package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}

	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Panic(err)
		}

		go handle(conn)
	}

}

func handle(conn net.Conn) {
	/* close after 10Sec. */
	/*
		err := conn.SetDeadline(time.Now().Add(10*time.Second))
		if err != nil {
			log.Println("CONN Timeout")
		}
	*/

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := strings.ToLower(scanner.Text())
		bs := []byte(ln)
		r := rot13(bs)
		fmt.Fprintf(conn, "I heard you say: %s\n", r)
	}

	defer conn.Close()

}

func rot13(bs []byte) []byte {
	var r13 = make([]byte, len(bs))
	for i, v := range bs {
		if v <= 109 {
			r13[i] = v + 13
		} else {
			r13[i] = v - 13
		}
	}
	return r13
}
