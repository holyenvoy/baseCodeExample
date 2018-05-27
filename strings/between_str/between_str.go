package main

import (
	"fmt"
	//"strconv"
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

func main() {

	betweenStr := GetBetweenStr("123456abcdef1----,12343", "123", "f1--")
	fmt.Printf("betweenStr :%v\n", betweenStr)

	substr := Substr("123456abcdef1----,12343", 2, 5)
	fmt.Printf("substr :%v\n", substr)

}
