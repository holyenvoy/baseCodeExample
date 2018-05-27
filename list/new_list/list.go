package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	front := l.PushFront(5)
	l.InsertAfter(2, l.InsertAfter(3, front))

	l.PushBack(1)
	// fmt.Println(el.Prev().Value)

	for el := l.Front(); el != nil; el = el.Next() {
		fmt.Println(el.Value)
	}

	fmt.Println("====")

	for el := l.Back(); el != nil; el = el.Prev() {
		if el.Value.(int) > 4 {
			l.InsertAfter(4, el)
		}
	}

	for el := l.Front(); el != nil; el = el.Next() {
		fmt.Println(el.Value)
	}
}
