package main

import (
	"fmt"
	//"strconv"
	"bytes"
	"strings"
)

func GetBetweenStr(str, start, end string) string {
	n := strings.Index(str, start)
	if n == -1 {
		return ""
	}
	n = n + len(start)
	str = string([]byte(str)[n:])
	m := strings.Index(str, end)
	if m == -1 {
		return ""
	}
	str = string([]byte(str)[:m])
	return str
}
func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

func RemoveDuplicate(list *[]string) []string {
	var x []string = []string{}
	for _, i := range *list {
		if len(x) == 0 {
			x = append(x, i)
		} else {
			for k, v := range x {
				if i == v {
					break
				}
				if k == len(x)-1 {
					x = append(x, i)
				}
			}
		}
	}
	return x
}

func UniqueSlice(slice *[]string) {
	found := make(map[string]bool)
	total := 0

	for i, val := range *slice {
		if _, ok := found[val]; !ok {
			found[val] = true
			(*slice)[total] = (*slice)[i]
			total++
		}
	}

	*slice = (*slice)[:total]
}

func uniqueSlice(slice *[]int64) {
	found := make(map[int64]bool)
	total := 0

	for i, val := range *slice {
		if _, ok := found[val]; !ok {
			found[val] = true
			(*slice)[total] = (*slice)[i]
			total++
		} else {
			fmt.Printf("get :%v val:%v\n", i, val)
		}
	}

	*slice = (*slice)[:total]
}

func main() {

	betweenStr := GetBetweenStr("123456abcdef1----,12343", "123", "f1--")
	fmt.Printf("betweenStr :%v\n", betweenStr)

	substr := Substr("123456abcdef1----,12343", 2, 5)
	fmt.Printf("substr :%v\n", substr)

	var duplicatsArray = []string{"11", "22", "33", "44", "44", "fdsfds", "33", "44", "55", "66", "77", "88", "44", "89"}

	duplicatsArray = RemoveDuplicate(&duplicatsArray)
	fmt.Printf("duplicatsArray :%v\n", duplicatsArray)

	duplicatsArray = []string{"11", "22", "33", "44", "44", "fdsfds", "33", "44", "55", "66", "77", "88", "44", "89"}
	UniqueSlice(&duplicatsArray)
	fmt.Printf("duplicatsArray :%v\n", duplicatsArray)

	duplicatsArrayInt := []int64{1, 2, 3, 4, 5, 1, 1, 2, 5, 6, 7, 8, 9, 9, 9, 9}
	uniqueSlice(&duplicatsArrayInt)
	fmt.Printf("duplicatsArrayInt :%v\n", duplicatsArrayInt)

	var buffer bytes.Buffer
	args := make([]interface{}, 0, 5*len(duplicatsArrayInt))
	for _, id := range duplicatsArrayInt {
		if buffer.Len() != 0 {
			buffer.WriteString(",")
		}
		buffer.WriteString("(?,?,?,?,?)")
		args = append(args, id)
	}
	fmt.Printf("args:%v\n", args)
}
