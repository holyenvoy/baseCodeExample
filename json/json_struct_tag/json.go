package main

import (
	"encoding/json"
	"fmt"
)

const MSG_CODE_RTV_LEAVE_CHANNEL = 9000

type NotifyMsg struct {
	Body struct {
		Code    int    `json:"code"`
		Content string `json:"content"`
	} `json:"body"`
	Ext struct {
		//Ext string `json:"ext"`
	} `json:"ext"`
}

type ContentMsg struct {
	EffectId int    `json:"effect_id"`
	Text     string `json:"text"`
}

type NotificationContent struct {
	Text string `json:"text"`
}

type ApnsPushMsg struct {
	Aps struct {
		Alert struct {
			Title string `json:"title"`
			Body  string `json:"body"`
		} `json:"alert"`
	} `json:"aps"`
}

type ApnsPushMsgV2 struct {
	Aps struct {
		Alert struct {
			Title           string `json:"title"`
			Body            string `json:"body"`
			TextContainsUrl bool   `json:"text_contains_url"`
		} `json:"alert"`
	} `json:"aps"`
}

type MyData struct {
	Name  string  `json:"item"`
	Other float32 `json:"amount"`
}

type InstructionMsgBody struct {
	Body struct {
		Url string `json:"url"`
	} `json:"body"`
	Ext struct {
	} `json:"ext"`
}

type UserInfo struct {
	UserId       int64  `protobuf:"varint,1,opt,name=user_id,json=userId" json:"user_id"`
	UserName     string `protobuf:"bytes,2,opt,name=user_name,json=userName" json:"user_name"`
	UserNickName string `protobuf:"bytes,3,opt,name=user_nick_name,json=userNickName" json:"user_nick_name"`
	AvatarUrl    string `protobuf:"bytes,4,opt,name=avatar_url,json=avatarUrl" json:"avatar_url"`
	ArUrl        string `protobuf:"bytes,5,opt,name=ar_url,json=arUrl" json:"ar_url"`
	UserGender   uint32 `protobuf:"varint,6,opt,name=user_gender,json=userGender" json:"user_gender"`
	Status       int32  `protobuf:"varint,7,opt,name=status" json:"status"`
	Mobile       string `protobuf:"bytes,8,opt,name=mobile" json:"mobile"`
	AppVersion   string `protobuf:"bytes,9,opt,name=app_version,json=appVersion" json:"app_version"`
	Model        string `protobuf:"bytes,10,opt,name=model" json:"model"`
}

type RelationInfo struct {
	FromUserId   int64     `protobuf:"varint,1,opt,name=from_user_id,json=fromUserId" json:"from_user_id"`
	ToUserId     int64     `protobuf:"varint,2,opt,name=to_user_id,json=toUserId" json:"to_user_id"`
	FromUserInfo *UserInfo `protobuf:"bytes,3,opt,name=from_user_info,json=fromUserInfo" json:"from_user_info"`
	Reason       string    `protobuf:"bytes,5,opt,name=reason" json:"reason"`
	SourceType   string    `protobuf:"bytes,6,opt,name=source_type,json=sourceType" json:"source_type"`
	InvitedAt    int64     `protobuf:"varint,7,opt,name=invited_at,json=invitedAt" json:"invited_at"`
}

type Array struct {
	Name []string `json:"name"`
}

type RelationBody struct {
	RelationType int32  `protobuf:"varint,1,opt,name=relation_type" json:"relation_type,omitempty"`
	SourceType   int32  `protobuf:"varint,2,opt,name=source_type" json:"source_type,omitempty"`
	Reason       string `protobuf:"bytes,3,opt,name=reason" json:"reason,omitempty"`
	Ext          string `protobuf:"bytes,4,opt,name=ext" json:"ext,omitempty"`
}

func main() {

	//
	//======================================================
	// 	通过struct的json标签生成json格式,注意域的大小写
	//======================================================
	//
	var testContentMsg *ContentMsg
	contentMsg := new(ContentMsg)
	if testContentMsg == nil {
		testContentMsg = &ContentMsg{Text: "dfs"}
		fmt.Printf("contentMsg is nil :%v\n\n", testContentMsg.EffectId)

	} else {
		fmt.Printf("contentMsg:%v\n\n", contentMsg)

	}
	contentMsg.Text = "wudebao方式端口了封禁"
	contentMsgJson, err := json.Marshal(contentMsg)

	fmt.Printf("contentMsg:%v\n", string(contentMsgJson))

	msg_body := &NotifyMsg{}
	msg_body.Body.Code = MSG_CODE_RTV_LEAVE_CHANNEL
	msg_body.Body.Content = string(contentMsgJson)
	msg, err := json.Marshal(msg_body)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("\nmsg body 3 is:%v\n msg:%s\n", msg_body, msg)

	json.Unmarshal(msg, &msg_body)
	fmt.Printf("===msg_body:%v content:%v\n", msg_body, msg_body.Body.Content)

	contentNtyMsg := new(NotificationContent)
	err = json.Unmarshal([]byte(msg_body.Body.Content), &contentNtyMsg)
	fmt.Printf("Unmarshal content err:%v ,text:%v\n", err, contentNtyMsg.Text)

	var a = [5][2]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}

	var unjson [][]int
	jsonArray, err := json.Marshal(a)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("jsonArray is:%v\n", string(jsonArray))
	json.Unmarshal(jsonArray, &unjson)
	fmt.Println(unjson)

	apns := &ApnsPushMsg{}
	apns.Aps.Alert.Body = "body"
	apns.Aps.Alert.Title = "title"
	apnsJson, err := json.Marshal(apns)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("\napnsJson:%s\n", apnsJson)

	textMsg := new(ApnsPushMsgV2)
	err = json.Unmarshal(apnsJson, &textMsg)
	fmt.Printf("err:%v apnsJson:%v textMsg:%v", err, string(apnsJson), textMsg)

	instruct := &InstructionMsgBody{}
	instruct.Body.Url = "http://wiki.meitu.com/Shanliao_DBProxy/admin/session/single"

	instructJson, err := json.Marshal(instruct)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("\n instructJson:%s\n", instructJson)

	relation := &RelationInfo{
		FromUserId: 112500,
		ToUserId:   112661,
		Reason:     "Reason",
		SourceType: "SourceType",
		InvitedAt:  213,
		FromUserInfo: &UserInfo{
			UserName: "UserName",
			ArUrl:    "fdsf",
		},
	}

	relationJson, err := json.Marshal(relation)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("\n relationJson:%s\n", relationJson)

	//nameSlice := make([]string, 0)
	var nameSlice1 []string  // 0x0,json后是null
	nameSlice2 := []string{} // 0x118c130,json 后是 []

	fmt.Printf("nameSlice:%p\n", nameSlice1)
	fmt.Printf("nameSlice:%p\n", nameSlice2)

	array := Array{
		Name: nameSlice1,
	}
	arrayJson, err := json.Marshal(array)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("arrayJson:%v\n", string(arrayJson))

	relationBody := &RelationBody{
		RelationType: 1,
		SourceType:   1,
		Reason:       "34",
		Ext:          "ew",
	}

	relationJson, err = json.Marshal(relationBody)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("\n relation body Json:%s\n", relationJson)

	if err := json.Unmarshal([]byte(relationJson), relationBody); err != nil {
		fmt.Printf("Unmarshal error ")
	}

}
