package main

import (
	"fmt"
)

var num int = 0

func Count(ch chan int) {
	ch <- 1
	num++
	fmt.Println("counting...., num:", num)

}

func main() {
	chs := make([]chan int, 10)

	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i])
	}

	for _, ch := range chs {
		<-ch
	}
}
