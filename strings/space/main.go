package main

import (
	"fmt"
	"strings"
)

func main() {

	someString := "one    two   three four 1 1  34   45"

	words := strings.Fields(someString)

	fmt.Println(words, len(words)) // [one two three four] 4

	fmt.Printf("\nThe output is:%v\n", strings.Replace(someString, " ", "", -1))

}
