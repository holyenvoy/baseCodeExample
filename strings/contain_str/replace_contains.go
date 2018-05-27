package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	History := "1111"
	peerUidStr := "2222"
	if strings.Contains(History, peerUidStr) {
		History = strings.Replace(History, peerUidStr, "", -1)
	}

	fmt.Println(History)

	var test []int

	fmt.Println(test)

	peerUidStr = strings.Replace(peerUidStr, "]", "", -1)
	fmt.Println(peerUidStr)
	fmt.Printf("%s\n", "==================")

	//
	//
	//
	//

	str := strings.Replace("[6193722415328067585,6193670168825102337]", "]", "", -1)
	str = strings.Replace(str, "[", "", -1)

	data := strings.Split(str, ",")

	var MsgIds []int64
	for _, id := range data {
		msgId, _ := strconv.ParseInt(id, 10, 64)
		MsgIds = append(MsgIds, msgId)
	}

	fmt.Println(MsgIds)

}
