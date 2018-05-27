package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[string]string{"b": "15", "z": "123123", "x": "sdf", "a": "12", "aa": "1112", "Aa": "212"}
	mkey := make([]string, len(m))
	mvalue := make([]string, len(m))
	i := 0
	for k, v := range m {
		mkey[i] = k
		mvalue[i] = v
		i++
	}
	sort.Strings(mkey)
	sort.Strings(mvalue)
	fmt.Println(mkey)
	fmt.Println(mvalue)
}
