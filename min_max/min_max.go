package main

import "fmt"

func main() {
	x := Min(1, 3, 2, 0)
	fmt.Printf("The minimum is: %d\n", x)
	arr := []int{7, 9, 3, 5, 1}
	x = Min(arr...)
	fmt.Printf("The minimum in the array arr is: %d\n", x)
}

/*
range是go语言系统定义的一个函数。
函数的含义是在一个数组中遍历每一个值，返回该值的下标值和此处的实际值。
假如说a[0]=10，则遍历到a[0]的时候返回值为0，10两个值。
*/
func Min(a ...int) int {
	if len(a) == 0 {
		return 0
	}
	min := a[0]
	for _, v := range a {
		fmt.Printf("===== v :%d\n", v)
		if v < min {
			min = v
		}
	}
	return min
}

/*
The minimum is: 0
The minimum in the array arr is: 1
*/
