package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	keyParts := strings.Split("5:3:1234:567", ":")
	if len(keyParts) < 4 {
		fmt.Printf("keyParts:%v", keyParts)
	}

	userId, err := strconv.ParseInt(keyParts[3], 10, 64)
	if err != nil {
		fmt.Printf("err :%v", err)
	}

	fmt.Println(userId)

	fmt.Println(112530 % 8)

}
