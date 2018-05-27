package main

import (
	"fmt"
	"strings"
)

func main() {
	content := "你好hello 你猜DFworld, 聊Qq"

	content = strings.ToUpper(content)
	fmt.Println(content)

	content = strings.ToLower(content)
	fmt.Println(content)
}
