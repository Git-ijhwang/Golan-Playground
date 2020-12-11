package main

import (
	"fmt"
	"net/http"
)

func main() {

	url := "http://diptera.myspecies.info:80"

	response, err := http.Head(url)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.Status, " ", response.StatusCode)
	for i, j := range response.Header {
		fmt.Println(i, " : ", j)
	}
}
