package main

import (
	"container/list"
	"fmt"
	"math/rand"
)

const SZ_ARRAY = 65535

type node struct {
	hs   int
	name string
	id   int
}

var hash [SZ_ARRAY]*list.List

func (nd *node) PrintHs(hv int) int {
	nd
	return hv * 2
}

func main() {
	var nd node

	for i := 0; i < SZ_ARRAY; i++ {
		hash[i] = list.New()
	}

	fmt.Println(&hash[0])
	for i := 0; i < SZ_ARRAY*2; i++ {
		nd.name = fmt.Sprintf("%d_%s", rand.Int()%SZ_ARRAY, "test")
		nd.id = rand.Int() % SZ_ARRAY
		hash[nd.id%SZ_ARRAY].PushBack(nd)
	}

	for i := 0; i < SZ_ARRAY; i++ {
		l := hash[i]
		//fmt.Println("----------[", i, "]------------")
		for e := l.Front(); e != nil; e = e.Next() {
			t := &(e.Value.(node))
			//e.Value.(node).PrintHs(e.Value.(node).id)
			x := t.PrintHs(e.Value.(node).id)
			fmt.Println(e.Value.(node).name)
			fmt.Println(e.Value.(node).hs)

			if i == 10 {
				fmt.Println(e.Value.(node).name, "=-", t.id%SZ_ARRAY, "----", e.Value.(node).hs, "=====", e.Value.(node).id)
			}
		}

	}
}
