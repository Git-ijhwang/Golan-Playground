package main

import (
	"fmt"
	"strings"
	"github.com/labstack/echo"
	"./scrapper"
)

var fileName = "jobs.csv"

//Handler function
func Handler(c echo.Context) error {
	return c.File("home.html")
}

//HandleFunc function
func HandleFunc(c echo.Context) error {
	query := strings.ToLower(scrapper.CleanString(c.FormValue("query")))
	fmt.Println(query)
	scrapper.Scrapper(query)

	return c.Attachment(fileName, query+".csv")
}

func main() {
	e := echo.New()
	e.GET("/", Handler)
	e.POST("/scrape", HandleFunc)
	e.Logger.Fatal(e.Start(":1323"))

}