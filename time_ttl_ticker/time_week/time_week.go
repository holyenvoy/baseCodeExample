package main

import (
	"fmt"
	"time"
)

func main() {
	//时间戳
	t := time.Now()
	weekDay := t.Weekday().String()

	fmt.Println(weekDay)

	if weekDay != "Monday" {
		fmt.Printf("not this Monday day\n")
	}

	if weekDay == "Friday" {
		fmt.Printf("this is Friday day\n")
	}

}
