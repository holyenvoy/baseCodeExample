package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type Card struct {
	Name      string
	Age       int
	Addresses []*Address
}

func main() {
	//
	//======================================================
	// 	通过struct生成json格式
	//======================================================
	//
	pa := &Address{"private", "Shanghai", "China"}
	pu := &Address{"work", "Beijing", "China"}
	c := Card{"Xin", 32, []*Address{pa, pu}}

	js, _ := json.Marshal(c)
	fmt.Printf("Json: %s\n\n", string(js))

	req := new(Card)
	json.Unmarshal(js, req)
	fmt.Println(req)

	var test []*Address
	test = append(test, pu)
	test = append(test, pa)
	js, _ = json.Marshal(test)
	fmt.Printf("Json array: %s\n\n", string(js))

}
