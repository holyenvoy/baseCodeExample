package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Info
}

type Info struct {
	Age int
}

func (user User) Print(name string) {
	fmt.Println("reflect Print()", name)
}

func Reflect(inter interface{}) {
	t := reflect.TypeOf(inter)                        //从接口中获取结构的类型
	v := reflect.ValueOf(inter)                       //从接口中获取结构的值
	if t.Kind() == reflect.Ptr && v.Elem().CanSet() { //传入的是指针,可以修改
		v = v.Elem()

		if f := v.Kind(); f == reflect.Struct { //如果字段属性是结构体
			if x := v.FieldByName("Age"); x.IsValid() {
				x.SetInt(888)
			}
		}

		if f := v.FieldByName("Name"); f.Kind() == reflect.String && f.IsValid() {
			f.SetString("haha")
		}

		if f := v.FieldByName("Id"); f.Kind() == reflect.Int && f.IsValid() {
			f.SetInt(99)
		}

		if f := v.MethodByName("Print"); f.IsValid() {
			args := []reflect.Value{reflect.ValueOf("Just Test For Print")}
			f.Call(args)
		}
	} else {
		fmt.Println("not pointer:", inter, "\n")
	}
}

func main() {
	user := User{1, "Allen.Wu", Info{23}}
	fmt.Println("1. input object but not pointer")
	Reflect(user)

	fmt.Println("2. input object and is pointer")
	Reflect(&user)

	fmt.Println("3. has changed the pointer:", user)

}
