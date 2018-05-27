package main

// 一般建议,同一个package下,只有一个init函数, init在main函数之前执行,用来做一些初始化工作

import (
	"fmt"
)

// the first init function in this go source file
func init() {
	fmt.Println("do in init1")
}

// the second init function in this go source file
func init() {
	fmt.Println("do in init2")
}

type FromUserRole struct {
	a int
}

func main() {
	fmt.Println("do in main")

	var a, b, c int
	var Token int
	userRole := &FromUserRole{a: 1}
	c = 1
	if a == 0 && b > 0 {
		fmt.Printf("a:%v b:%v\n", a, b)
	} else if c != 0 {
		fmt.Println("else")
	}

	Token = 1
	if Token > 0 && userRole != nil {
		fmt.Println("123")
	}

}
