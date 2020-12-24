package main

import (
	"container/list"
	"fmt"
	"math/rand"
)

const SZ_ARRAY = 10

type node struct {
	hs   int
	name string
	id   int
}

var hash [SZ_ARRAY]*list.List

func PrintHs() {
	for i := 0; i < SZ_ARRAY; i++ {
		fmt.Println("#", i)
		l := hash[i]
		/* loop for inner linked list */
		for e := l.Front(); e != nil; e = e.Next() {
			fmt.Println("\t Name: ", e.Value.(node).name)
			fmt.Println("\t ID: ", e.Value.(node).id)
			fmt.Println("\t Hash Key: ", e.Value.(node).hs, "\n")
			//,  t.id%SZ_ARRAY,  e.Value.(node).hs,  e.Value.(node).id)
		}
	}
}

func insert(key int) int {
	var nd node

	var h_key int = key % SZ_ARRAY

	l := hash[h_key]

	e := find(key)
	if e != nil {
		fmt.Println("Errror")
		return -1
	}

	nd.name = fmt.Sprintf("%d_%s", h_key, "test")
	nd.id = key
	nd.hs = h_key
	l.PushFront(nd)
	return h_key
}

func del(key int) {

	var h_key int = key % SZ_ARRAY

	l := hash[h_key]

	e := find(key)
	if e != nil {
		l.Remove(e)
		fmt.Println("Delete Success")
	}
}

func find(key int) *list.Element {
	var h_key int = key % SZ_ARRAY

	l := hash[h_key]

	if l.Len() <= 0 {
		return nil
	}

	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value.(node).id == key {
			fmt.Println("Found it the node!")
			return e
		}
	}
	return nil
}

var input int

func main() {

	/* List in the hash Initialization */
	for i := 0; i < SZ_ARRAY; i++ {
		hash[i] = list.New()
	}

	/* push the data into hash */
	for i := 0; i < SZ_ARRAY*3; i++ {
		insert(rand.Int())
	}

	PrintHs()

	/* Test for Delete */
	fmt.Scanf("%d", &input)
	fmt.Println(input)
	del(input)

	PrintHs()
}
