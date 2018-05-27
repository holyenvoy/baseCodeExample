package main

import (
	"fmt"
	"time"
)

func testValue(j int) int {

	go func(int) {
		time.Sleep(time.Second * 1)
		fmt.Printf("in go func j=%v\n", j)
	}(j)

	fmt.Printf("in test j:%v\n", j)
	return j
}

func main() {
	//尾部加括号传入参数直接执行
	sum := func(a, b int) int {
		return a + b
	}(3, 4)
	fmt.Println(sum)

	//赋值给变量时使用
	f := func(i, j int) (result int) {
		result = i + j
		return result
	}
	fmt.Println(f(1, 3))

	var j int = 5
	a := func() func() {
		var i int = 10
		return func() {
			fmt.Printf("i, j: %d, %d\n", i, j)
		}
	}()

	a()
	j *= 2
	a()

	fmt.Printf("j = %v\n", j)

	testValue(j)

	fmt.Printf("testValue j = %v\n", j)

	time.Sleep(2 * time.Second)

}
