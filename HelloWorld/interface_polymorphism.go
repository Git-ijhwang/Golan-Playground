package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

//Keyword Identifier Type
type human interface {
	speak()
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, "- the agent speak")
}

func (s person) speak() {
	fmt.Println("I am", s.first, s.last, "- the person speak")
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I called human", h.(person).first, h.(person).last)
	case secretAgent:
		fmt.Println("I called human", h.(secretAgent).first, h.(secretAgent).last)
	}
}

func main() {
	sa1 := secretAgent{
		person: person{
			"James",
			"Bone",
		},
		ltk: true,
	}
	sa2 := secretAgent{
		person: person{
			"Miss",
			"MoneyPenny",
		},
		ltk: true,
	}

	p1 := person{
		first: "Dr.",
		last:  "Yes",
	}

	//Example of Interface
	fmt.Println(sa1)
	sa1.speak()
	p1.speak()

	//Example of Polymorphism
	bar(sa1)
	bar(sa2)
	bar(p1)
}
