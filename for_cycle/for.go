package main

import (
	"fmt"
	"time"
)

func main() {
	var count, Num int
	for {
		Num++
		if Num >= 3 {
			break
		}
		for {
			now := time.Now().UnixNano() / 1e6
			fmt.Println(now)

			time.Sleep(time.Second * 1)
			count++
			if count == 2 {
				continue
			}
			if count >= 3 {
				break
			}

		}
	}

}
