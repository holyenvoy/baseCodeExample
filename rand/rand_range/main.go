package main

import (
	"fmt"
	"math/rand"
	"time"
)

//生成若干个不重复的随机数
func RandomTestBase() {

	circle := 10
	countNum := 10
	count := 0

	//测试5次
	for i := 0; i < circle; i++ {
		nums := generateRandomNumber(0, 100, countNum)
		fmt.Println(nums)
		for _, num := range nums {
			if num <= 50 {
				count++
			}
		}
	}

	fmt.Printf("count=%v\n", count)
}

//生成count个[start,end)结束的不重复的随机数
func generateRandomNumber(start int, end int, count int) []int {
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}

	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn((end - start)) + start

		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}

		if !exist {
			nums = append(nums, num)
		}
	}

	return nums
}

func main() {

	fmt.Println(generateRandomNumber(0, 100, 1))

	RandomTestBase()
}
