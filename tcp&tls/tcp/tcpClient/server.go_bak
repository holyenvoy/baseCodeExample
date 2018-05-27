package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", ":4040")
	checkErr(err)
	listen, err := net.ListenTCP("tcp", addr)
	checkErr(err)
	fmt.Println("Start server...")
	for {
		conn, err := listen.Accept()
		checkErr(err)
		go Handle(conn) // 每次建立一个连接就放到单独的线程内做处理
	}
}

const BufLength = 128

var users map[string]net.Conn = make(map[string]net.Conn, 10)

func Handle(conn net.Conn) {
	conn.Write([]byte("欢迎加入2B聊天组~"))
	for {
		data := make([]byte, 0) //此处做一个输入缓冲以免数据过长读取到不完整的数据
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

		cmd := strings.Split(string(data), "|")
		fmt.Println("命令:", cmd)

		switch cmd[0] {
		case "nick":
			fmt.Println("注册名称:" + cmd[1])
			users[cmd[1]] = conn
		case "say":
			for k, v := range users {
				if k != cmd[1] {
					fmt.Println("给" + k + "发送消息:" + cmd[2])
					v.Write([]byte(cmd[1] + ":[" + cmd[2] + "]"))
				}
			}
		}
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
