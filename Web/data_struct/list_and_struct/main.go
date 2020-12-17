package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type user struct {
	Name  string
	Motto string
	Admin bool
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	u1 := user{
		Name:  "Buddha",
		Motto: "blah",
		Admin: false,
	}

	u2 := user{
		Name:  "Jesus",
		Motto: "blah, blah",
		Admin: true,
	}

	u3 := user{
		Name:  "",
		Motto: "Nobody",
		Admin: true,
	}

	users := []user{u1, u2, u3}
	//
	//xs := []string{"zero", "one", "two", "three", "four", "five" }
	//
	//data := struct {
	//	Words []string
	//	Lname string
	//}{
	//	xs,
	//	"Injun",
	//}

	//err := tpl.Execute(os.Stdout, data)
	err := tpl.Execute(os.Stdout, users)
	if err != nil {
		log.Fatal(err)
	}
}
