package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	sage := map[string]string{
		"test1": "TEST1",
		"test2": "TEST2",
		"test3": "TEST3",
		"test4": "TEST4",
		"test5": "TEST5",
	}

	err := tpl.Execute(os.Stdout, sage)
	if err != nil {
		log.Fatal(err)
	}
}
