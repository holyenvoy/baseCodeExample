package main

import (
	"fmt"
)

func main() {

	var testStr string

	if len(testStr) == 0 {
		fmt.Printf("testStr is null:%v\n", testStr)
	}

	if testStr == "" {
		fmt.Printf("testStr is null:%v\ns", testStr)

	}

}
