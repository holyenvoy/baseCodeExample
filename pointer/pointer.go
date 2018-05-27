package main

import (
	"fmt"
)

type MsgData struct {
	TextContent   string `protobuf:"bytes,3,opt,name=text_content,json=textContent" json:"text_content"`
	SensitiveType int64  `protobuf:"varint,4,opt,name=sensitiveType" json:"sensitiveType"`
}

type SessionInfo struct {
	PeerId      int64      `protobuf:"varint,1,opt,name=peer_id,json=peerId" json:"peer_id"`
	UpdatedTime int64      `protobuf:"varint,2,opt,name=updated_time,json=updatedTime" json:"updated_time"`
	MsgData     []*MsgData `protobuf:"bytes,3,rep,name=msg_list,json=msgList" json:"msg_list"`
	Hash        string     `protobuf:"bytes,7,opt,name=hash" json:"hash"`
}

type SingleSession struct {
	PeerId  int64
	MsgData *MsgData
}

func main() {

	msgDataV1 := new(MsgData)
	msgDataV2 := new(MsgData)
	msgDatas := make([]*MsgData, 0)

	sessionInfos := make([]*SessionInfo, 0)
	singleSession := new(SingleSession)
	if singleSession == nil {
		fmt.Printf("singleSession is nil \n")
	}

	if singleSession.MsgData == nil {
		fmt.Printf("singleSession.MsgData is nil \n")
	}

	fmt.Printf("singleSession:%v\n", singleSession)
	fmt.Printf("singleSession.MsgData:%v\n", singleSession.MsgData)

	singleSession.MsgData = msgDataV2
	singleSession.MsgData.SensitiveType = 1

	fmt.Println(singleSession)

	msgDataV2.SensitiveType = 123
	msgDataV1 = msgDataV2

	msgDataV1.TextContent = "234"

	msgDatas = append(msgDatas, msgDataV1)
	msgDatas = append(msgDatas, msgDataV2)

	session := &SessionInfo{
		PeerId:  1233,
		MsgData: msgDatas,
		Hash:    "123",
	}

	sessionInfos = append(sessionInfos, session)

	fmt.Println(msgDataV1)
	fmt.Println(msgDatas)

	fmt.Println(sessionInfos)

}
