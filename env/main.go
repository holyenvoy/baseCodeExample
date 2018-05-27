package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("get env:%v\n", os.Getenv("HOST"))
}
