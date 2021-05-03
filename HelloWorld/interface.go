package main

import (
	"fmt"
	"math"
)

type Square struct {
	Length float64
}

func (s *Square) getType() {
	fmt.Println("I'm Square");
}

func (s *Square) Area() float64 {
	return s.Length * s.Length
}

type Circle struct {
	Radius float64
}
func (s *Circle) getType() {
	fmt.Println("I'm Circle");
}


func (c * Circle) Area() float64 {
	return math.Pi * c.Radius *c.Radius
}

type action interface {
	Area() float64
	getType()
}


func sumAreas(shapes []Shape) float64 {
	total := 0.0

	for _, shape := range shapes  {
		total += shape.Area()
	}

	return total
}

type Shape interface {
	Area() float64
}

// func doit (a action)  {
// 	a.getType()
// 	fmt.Println(a.Area())
// }

func main() {
	var s *Square
	s = &Square{20}
	c := &Circle{10}

	var intf action
	intf = s
	intf.getType()
	fmt.Println("	area : ", intf.Area())
	intfb := c
	intfb.getType()
	fmt.Println("	area: ", intfb.Area())


	// ret := doit(c)
	// fmt.Println(s.Area())
	// fmt.Println(c.Area())

	
	shapes := []Shape{s, c}
	sa := sumAreas(shapes)
	fmt.Println("Total Area is ", sa);
}



