package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

// 匿名字段处理
type Manager struct {
	User
	title string
}

func main() {
	m := Manager{User: User{1, "Corwien", 18}, title: "123456"}
	t := reflect.TypeOf(m)

	fmt.Println("%#v\n", t.Field(1))
	fmt.Println("%#v\n", t.FieldByIndex([]int{0, 1}))
}
