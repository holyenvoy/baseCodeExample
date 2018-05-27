package main

import (
	"fmt"
)

func main() {

	testString := "123 %s %s abc %s%% %s"

	//testString = fmt.Sprintf(testString, "test1", " test2 ")
	//testString = fmt.Sprintf(testString, "test3")

	testString = fmt.Sprintf(testString, "test1", " test2 ", "test3")

	testString = fmt.Sprintf(testString, "resr")

	fmt.Printf("testString :%v\n", testString)

	extra_percent := fmt.Sprintf("%v%%", 123)
	fmt.Printf("extra_percent:%v\n", extra_percent)

}
