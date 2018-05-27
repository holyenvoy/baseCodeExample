package main

import (
	"fmt"
	"time"
)

type ExtRelationEntity struct {
	FriendUid  int64 `json:"friend_uid"`
	ExtendTime int64 `json:"extend_time"`
	ExtendType int32 `json:"extend_type"`
}

func main() {
	mapResults := make(map[int]string)

	// 数组分片定义的几种方法，够能够进行append操作
	//var arrResults [][]string
	//arrResults := [][]int{}
	arrResults := make([][]int, 0)

	count := 5
	for i := 0; i < count; i++ {
		valueStr := fmt.Sprintf("this is %d\r\n", i)
		mapResults[i] = valueStr
		var tmpArr []int
		for j := 0; j < 2; j++ {
			tmpArr = append(tmpArr, j)
		}
		arrResults = append(arrResults, tmpArr)
	}
	fmt.Println(mapResults)
	fmt.Println(arrResults)

	var uidMembers []int

	for _, member := range arrResults {
		fmt.Println(member[len(member)-1])
		uidMembers = append(uidMembers, member[0])
	}
	fmt.Println(uidMembers)

	fmt.Println()

	var s = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	index := 0
	endIndex := len(s) - 1
	first := time.Now().UnixNano()

	// 删除指定元素
	var result = make([]string, 0)
	for k, v := range s {
		if v == "8" {
			result = append(result, s[index:k]...)
			index = k + 1
		} else if k == endIndex {
			result = append(result, s[index:endIndex+1]...)
		}
	}
	fmt.Printf("del ret method 1:%v\n", result)
	now := time.Now().UnixNano()
	fmt.Printf("duartion :%vns\n", now-first)

	for k, v := range s {
		if v == "8" {
			kk := k + 1
			s = append(s[:k], s[kk:]...)
		}
	}
	fmt.Printf("del ret method 2:%v\n", s)

	resultPointer := make([]*int, 0)
	for _, ret := range resultPointer {
		fmt.Printf("ret :%v\n", ret)

	}

	fmt.Printf("1 ---- s:%v\n", s)
	var s2 = []string{"11", "22", "32"}
	s = append(s2, s...)
	//fmt.Printf("2 ---- s:%v\n", s)
	//fmt.Printf("2 ---- s2:%v\n", s2)
	fmt.Printf("first s = %v  len:%v \n", s, len(s))
	s = s[1:11]
	fmt.Printf("second s = %v  len:%v \n", s, len(s))

	fmt.Printf("2 ---- s:%v\n", s)
	fmt.Printf("----s:%v ------ \n", s)
	s = s[2:4]
	fmt.Printf("2:4 ---- s:%v\n", s)

	resultEntity := make([]*ExtRelationEntity, 0)
	resultEntity = append(resultEntity, &ExtRelationEntity{})
	for _, entity := range resultEntity {
		fmt.Printf("%v %v %v\n", entity.ExtendTime, entity.ExtendType, entity.FriendUid)
	}

	fmt.Printf("get value:%v\n", 365*2%365)

}
