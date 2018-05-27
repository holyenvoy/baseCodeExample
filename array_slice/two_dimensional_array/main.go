package main

import "fmt"

func getUids(uids ...int) {
	fmt.Println(uids)
}

func main() {
	/* 数组 - 5 行 2 列*/
	var a = [5][2]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}
	var i, j int

	/* 输出数组元素 */
	for i = 0; i < 5; i++ {
		for j = 0; j < 2; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}

	for i, num := range a {
		fmt.Printf("%v %v %v\n", i, num[0], num[1])
	}

}
