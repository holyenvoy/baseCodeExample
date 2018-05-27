package main

import (
	"fmt"
)

type Base struct {
	name string
	age  int
}

func (base *Base) Set(myname string, myage int) {
	base.name = myname
	base.age = myage
}

type Derived struct {
	Base
	name string
}

func (base *Base) GetCount(x int, y int) (int, int) {
	y = 11 + x
	return x, y

}

func (base *Base) myPrint(num int) {
	num += num
	fmt.Println("defer --- myprint :", num)
}

func main() {
	b := &Derived{}

	b.Set("sina", 30)
	fmt.Println("b.name =", b.name, "\tb.Base.name =", b.Base.name)
	fmt.Println("b.age =", b.age, "\tb.Base.age =", b.Base.age)

	x, y := b.GetCount(1, 2)
	fmt.Println("x1 = ", x, " y1 = ", y)

	defer b.myPrint(x)

	x, y = b.GetCount(10, 20)
	fmt.Println("x2 = ", x, " y2 = ", y)
	defer b.myPrint(x)

}
