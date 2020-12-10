package main

import "fmt"
import "math/rand"

func swap(p *int, q *int) {
	t := *p
	*p = *q
	*q = t
}

func sort(letter [10]int) {
	//fmt.Println(letter)

}

func sort_print(p [100]int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", p[i])
	}
	fmt.Printf("\n")
}

func main() {
	//var letter [10]int = [5, 3, 1, 2, 9, 8, 7, 6, 4 ]
	//letter := []int{5, 3, 1, 2, 9, 8, 7, 6, 4 }
	var ptr *int
	var value int

	value = 10
	ptr = &value
	fmt.Println(*ptr)
	fmt.Println(ptr)
	fmt.Println("=-======")
	var letter [100]int
	for i := 0; i < 100; i++ {
		letter[i] = rand.Int() % 30
	}
	sort_print(letter)

	for i := 0; i < len(letter); i++ {
		for j := i; j < len(letter)-1; j++ {
			if letter[i] > letter[j+1] {
				swap(&letter[i], &letter[j+1])
			}
		}
	}
	sort_print(letter)
	//fmt.Println(letter)
}
