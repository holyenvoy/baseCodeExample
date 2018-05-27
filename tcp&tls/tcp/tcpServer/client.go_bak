package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

var nick string = ""

func main() {
	addr, err := net.ResolveTCPAddr("tcp", ":4040")
	checkErr(err)
	conn, err := net.DialTCP("tcp", nil, addr)
	checkErr(err)
	// 读取提示
	data := make([]byte, 1024)
	conn.Read(data)
	fmt.Println(string(data))
	// 输入昵称
	fmt.Print("输入昵称:")
	fmt.Scanf("%v", &nick)
	fmt.Println("Hello " + nick)
	conn.Write([]byte("nick|" + nick))

	go Handle(conn)

	for {
		someTex := ""
		fmt.Scanf("%v", &someTex)
		conn.Write([]byte("say|" + nick + "|" + someTex))
	}
}

const BufLength = 128

func Handle(conn net.Conn) {
	for {
		data := make([]byte, 1024)
		buf := make([]byte, BufLength)
		for {
			n, err := conn.Read(buf)
			if err != nil && err != io.EOF {
				checkErr(err)
			}
			data = append(data, buf[:n]...)
			if n != BufLength {
				break
			}
		}

		fmt.Println(string(data))
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
