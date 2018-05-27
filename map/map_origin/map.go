package main

import "fmt"

func main() {
	blogArticleViews := map[string]int{
		"unix":         0,
		"python":       1,
		"go":           2,
		"javascript":   3,
		"testing":      4,
		"philosophy":   5,
		"startups":     6,
		"productivity": 7,
		"hn":           8,
		"reddit":       9,
		"C++":          10,
	}
	for key, views := range blogArticleViews {
		fmt.Println("There are", views, "views for", key)
	}

	mykey := "C++11"
	value, ok := blogArticleViews[mykey]
	if ok {
		fmt.Printf("get mykey:%v\n", value)
	} else {
		fmt.Printf("can not get my key\n")
	}

	remindUids := make(map[int32][]int64)

	remindUids[1] = append(remindUids[1], 234)
	remindUids[1] = append(remindUids[1], 234)
	remindUids[2] = append(remindUids[2], 11)

	fmt.Println(remindUids)

}
