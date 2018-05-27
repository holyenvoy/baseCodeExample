package main

import (
	"fmt"
	"time"
)

var callTimes int
var count int
var Num int

func callSelf(count int) {
	count++

	fmt.Printf("count:%v\n", count)
	if count <= 8 {
		callTimes++
		if callTimes > 3 {
			return
		}
		go func() {
			time.Sleep(3 * time.Second)
			fmt.Printf("call self:%v\n", count)
			callSelf(count)
		}()
	}
}

func main() {

	callSelf(1)

	callSelf(1)

	callSelf(3)

	time.Sleep(1 * time.Second)

	fmt.Printf("call self over %v\n", callTimes)
	time.Sleep(16 * time.Second)
}
