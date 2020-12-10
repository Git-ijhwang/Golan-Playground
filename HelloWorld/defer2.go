package main

import "fmt"

func cleanup(name string) {
	fmt.Printf("Cleaning up %s\n", name)
}

func worker() {
	defer cleanup("A")
}

func main() {
	worker()
}
