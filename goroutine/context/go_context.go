package main

import (
	"fmt"
	"time"

	"golang.org/x/net/context"
)

func SendMsg(ctx context.Context, count int) {

	fmt.Printf("start send msg count:%v\n", count)

	if count == 3 {

		go func() {
			time.Sleep(2 * time.Second)

			fmt.Printf("in go routin \n")
			count += 3
			SendMsg(ctx, count)
		}()

	}

	fmt.Printf("end send msg count:%v\n", count)

}

func main() {

	SendMsg(context.Background(), 3)

	time.Sleep(5 * time.Second)

	fmt.Printf("end main\n")

}
