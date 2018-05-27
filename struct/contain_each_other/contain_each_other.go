package main

import (
	"fmt"
)

type User struct {
	uid               int64
	conn              *MsgConn
	bLogined          bool
	BKicked           bool // 被另外登陆的一方踢下线
	bPeerClosed       bool
	BHeartBeatTimeout bool // 心跳超时
}

type MsgConn struct {
	stopChan   chan bool
	remoteAddr string // 为每个连接创建一个唯一标识符
	user       *User  // MsgConn与User一一映射
}

func main() {

	msgConn := &MsgConn{
		remoteAddr: "123",
	}

	user := &User{
		uid:  111,
		conn: msgConn,
	}

	msgConn.user = user

	fmt.Println(msgConn.user.uid)
	fmt.Println(user.conn.remoteAddr)
}
