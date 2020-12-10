package main

import (
	"fmt"
)
import "math/rand"

func swap(p *int, q *int) {
	t := *p
	*p = *q
	*q = t
}

//func sort(letter [100]int) {
func sort(letter *[10]int) {
	//
	//	fmt.Println(*letter)
	//	//fmt.Println(letter+1)
	//	for i, v := range letter {
	//		fmt.Println(i)
	//		fmt.Println(v)
	//	}

	for i := 0; i < len(letter); i++ {
		for j := i; j < len(letter)-1; j++ {
			if letter[i] > letter[j+1] {
				swap(&letter[i], &letter[j+1])
			}
		}
	}
}

func sort_print(p [10]int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", p[i])
	}
	fmt.Printf("\n")
}

func main() {
	var letter [10]int
	for i := 0; i < len(letter); i++ {
		letter[i] = rand.Int() % 10
	}

	sort_print(letter)
	sort(&letter)
	sort_print(letter)
}
