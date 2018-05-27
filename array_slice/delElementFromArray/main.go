package main

import (
	"fmt"
	"time"
)

func delElementFromArray(original []int64, element int64) []int64 {
	if len(original) <= 0 {
		return nil
	}

	for index, value := range original {
		if value == element {
			newIndex := index + 1
			original = append(original[:index], original[newIndex:]...)
			break
		}
	}

	return original
}

func main() {
	delArray, testArray := make([]int64, 0), make([]int64, 0)

	var i int64

	for i = 1; i < 60; i += 1 {
		testArray = append(testArray, i)
	}

	for i = 0; i < 15; i++ {
		delArray = append(delArray, i)
	}

	fmt.Printf("len:%v \nlen:%v\n", len(testArray), len(delArray))
	startTime := time.Now()
	for _, v := range delArray {
		testArray = delElementFromArray(testArray, v)
	}

	dis := time.Now().Sub(startTime).Seconds()

	fmt.Printf("dis:%v\n", dis)
	fmt.Println(testArray)
}
