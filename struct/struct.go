package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"time"
)

type info struct {
	Name string
	Time time.Time
}
type newlist []*info

func main() {
	l, e := getFilelist("./")
	if e != nil {
		fmt.Println(e)
	}
	sort.Sort(newlist(l)) //调用标准库的sort.Sort必须要先实现Len(),Less(),Swap() 三个方法.
	for _, v := range l {
		fmt.Println("文件名：", v.Name, "修改时间：", v.Time.Unix())
	}
}

func getFilelist(path string) ([]*info, error) {
	l, err := ioutil.ReadDir(path)
	if err != nil {
		return []*info{}, err
	}
	var list []*info
	for _, v := range l {
		list = append(list, &info{v.Name(), v.ModTime()})
	}
	return list, nil
}

func (I newlist) Len() int {
	return len(I)
}
func (I newlist) Less(i, j int) bool {
	return I[i].Time.Unix() > I[j].Time.Unix()
}
func (I newlist) Swap(i, j int) {
	I[i], I[j] = I[j], I[i]
}
