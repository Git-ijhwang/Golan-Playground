package main

import (
	"container/list"
	"fmt"
	//"go/types"
)

type node struct {
	name string
	id   int
}

var hash [10]*list.List

func main() {
	var nd node

	for i := 0; i < 10; i++ {
		hash[i] = list.New()
	}

	fmt.Println(&hash[0])
	for i := 0; i < 25; i++ {
		nd.name = fmt.Sprintf("%s%d", "test", i)
		nd.id = i
		hash[nd.id%10].PushBack(nd)
		//t.PushBack(nd)
	}

	for i := 0; i < 10; i++ {
		l := hash[i]
		fmt.Println("----------", l, "------------")
		for e := l.Front(); e != nil; e = e.Next() {
			fmt.Println(e.Value.(node).name)
		}
	}
}
