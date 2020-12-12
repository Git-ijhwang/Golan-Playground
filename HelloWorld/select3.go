package main

import (
	"fmt"
	"time"
)

func func1(c1 chan string) {
	for i := 0; i < 10; i++ {
		c1 <- "from 1"
		time.Sleep(time.Second * 2)
	}
}

func func2(c2 chan string) {
	for i := 0; i < 10; i++ {
		c2 <- "from 2"
		time.Sleep(time.Second * 3)
	}
}

func main() {
	var c1 chan string = make(chan string)
	var c2 chan string = make(chan string)

	go func1(c1)
	go func2(c2)

	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			case <-time.After(time.Second * 5):
				fmt.Println("timeout")

			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}
