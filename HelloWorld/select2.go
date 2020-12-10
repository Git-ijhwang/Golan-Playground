package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- 42
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			ch2 <- 420
			time.Sleep(100 * time.Millisecond)
		}
	}()

	for i := 0; true; i++ {
		select {
		case val := <-ch1:
			fmt.Printf("got %d from ch1\n", val)
		case val := <-ch2:
			fmt.Printf("got %d from ch2\n", val)
		}
		time.Sleep(10 * time.Millisecond)
	}

	fmt.Println("----------------------------------")
}
