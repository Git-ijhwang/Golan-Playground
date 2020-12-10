package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("no-search-file")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	fmt.Println("file opened")
}
