package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	array := [5]int64{1, 2, 3, 4, 5}
	jsonArray, err := json.Marshal(array)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("jsonArray is:%v\n", string(jsonArray))

	array2 := [][]int{{1, 2}, {3, 4}}
	jsonArray, err = json.Marshal(array2)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("array2 is:%v\n", string(jsonArray))

}
