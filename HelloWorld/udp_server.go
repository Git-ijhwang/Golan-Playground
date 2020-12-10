package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func(ctx context.Context) { //Udp Server thread
		conn, err := net.ListenPacket("udp", "127.0.0.1:8081")

		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		fmt.Println("I'm Listening....")
		var buf [1518]byte

		for {
			select {
			default:
				n, _, err := conn.ReadFrom(buf[:])
				if err == io.EOF {
					return
				}

				if err != nil {
					log.Fatal(err)
				}

				fmt.Printf("Received msg len : %d\n", n)
				fmt.Println(string(buf[:n]))

			case <-ctx.Done():
				fmt.Println("bye")
				return
			}
		}
	}(ctx)

	conn, err := net.Dial("udp", "127.0.0.1:8081")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		msg := scanner.Text()
		if msg == "end" {
			return
		}

		fmt.Fprint(conn, msg)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

}
