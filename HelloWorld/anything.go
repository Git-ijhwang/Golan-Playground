package main

import "fmt"

func main() {
	x := make([]int, 10, 15)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x = append(x, 123, 234, 345, 456, 567)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 432)
	fmt.Println(len(x))
	fmt.Println(cap(x))

}
