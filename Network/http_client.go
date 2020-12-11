package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	url, err := url.Parse("http://diptera.myspecies.info:80")
	if err != nil {
		panic(err)
	}

	client := &http.Client{}

	req, err := http.NewRequest("GET", url.String(), nil)

	response, err := client.Do(req)

	if response.Status != "200 OK" {
		panic(err)
	}

	fmt.Println(response.Status)
	fmt.Println(response.Header)
}
