package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum(x)
	ret := sum2(x...)
	fmt.Println("Return value : ", ret)
}

func sum(x []int) {
	sum := 0
	for i, v := range x {
		fmt.Printf("i : %d, v : %d\n", i, v)
		sum += v
	}
	fmt.Println("The value of sum is ", sum)
}

func sum2(x ...int) int {
	sum := 0
	for v, i := range x {
		sum += i
		fmt.Println(v, i, sum)
	}
	return sum
}
