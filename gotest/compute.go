package util

//"fmt"

func Sum(n ...int) int {
	var c int
	for _, i := range n {
		c += i
	}
	return c
}

func Abs(a int) int {
	if a > 0 {
		return a
	} else {
		return a * (-1)
	}
}
