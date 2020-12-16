package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func main() {
	sages := []string{"Gandhi", "MLK", "Buddha", "Jesus", "Muhammad"}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
