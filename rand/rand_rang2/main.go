package main

import (
	"math/rand"

	"fmt"
)

func main() {
	for i := 0; i < 2; i++ {
		num := RandInt(60, 90)
		fmt.Println(num)
	}

	for i := 1; i <= 10; i++ {
		fmt.Printf("origin:%v fix rand:%v\n", i, RandFixInt(i))
	}
}

func RandInt(min, max int) int {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Intn(max-min) + min
}

func RandFixInt(input int) int {
	return rand.Intn(input)
}
