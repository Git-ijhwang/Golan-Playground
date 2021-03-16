package main

import "fmt"

func test(row rune) {
	fmt.Println(row)
	fmt.Println(string(row))
}
func main() {
	x := make([]int, 10, 15)
	fmt.Println("array X : ", x)
	fmt.Println("length of X : ", len(x))
	fmt.Println("Cap of X: ", cap(x))
	x = append(x, 123, 234, 345, 456, 567)

	fmt.Println("Result of Append : ", x)
	fmt.Println("len of X : ", len(x))
	fmt.Println("Cap of X : ", cap(x))

	x = append(x, 432)
	fmt.Println("len of X after append : ", len(x))
	fmt.Println("Cap of X : ", cap(x))
	fmt.Println( fmt.Sprintf("Cap of X : %%"))
	n := 5
	c := 'A'
	fmt.Printf("%c\n", c)
	c = c+1
	fmt.Printf("%c\n", c)
	c = c+rune(n*2)
	fmt.Printf("%c\n", c)
	test(c)
	test(c+1)

}
