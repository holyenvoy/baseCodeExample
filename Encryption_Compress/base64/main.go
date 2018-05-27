package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {

	//input := []byte("\"{\"body\":{\"content\":{\"reason\":\"test....\",\"circle_name\":\"\\u975e\\u6cd5\\u5708\\u5b50\\u540d\\uff0c\\u8bf7\\u4fee\\u6539\"}},\"ext\":[]}\"")

	input := []byte("hello:allenWu")
	// 演示base64编码
	encodeString := base64.StdEncoding.EncodeToString(input)
	fmt.Printf("encodeString:%v\n", encodeString)

	// 对上面的编码结果进行base64解码
	decodeBytes, err := base64.StdEncoding.DecodeString("/rk5+QABAAAAFwABAAIAAQAAAWNh4Qc1AAAAAAAAAAEAAAAAAAAAATU0ZDRkZmVkYTdiODg0ZGIAAQoBMRIS57uZ5oiR6K6y5Liq56yR6K+d")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("\ndecodeBytes:%v\n\n", string(decodeBytes))

	fmt.Println()

	// 如果要用在url中，需要使用URLEncoding
	//uEnc := base64.URLEncoding.EncodeToString([]byte(input))
	//fmt.Println(uEnc)
	/*
		uDec, err := base64.URLEncoding.DecodeString("H4sIAAAAAAAAAxWNUQqDMBBE77LfIgHdRr1KV4rZbKBgE9AstEju3nU+HgNvYC4IJf5guYBLrpLrXQ/ZzpJhgSpn7S3QAb8P3uWVt4+YIJ09CumDI5Kid5MxoDOOLpKm5Jh0CsmTjknuJQ4ztNaBfO3jubY/i0spDXoAAAA=")
		if err != nil {
			fmt.Printf("decode err:%v ret:%v\n", err, string(uDec))
			return
		}
		fmt.Println(string(uDec))
	*/

	decodeBytes, err = base64.StdEncoding.DecodeString("H4sIAAAAAAAAAxWNUQqDMBBE77LfIgHdRr1KV4rZbKBgE9AstEju3nU+HgNvYC4IJf5guYBLrpLrXQ/ZzpJhgSpn7S3QAb8P3uWVt4+YIJ09CumDI5Kid5MxoDOOLpKm5Jh0CsmTjknuJQ4ztNaBfO3jubY/i0spDXoAAAA=")
	if err != nil {
		fmt.Printf("decode err:%v ret:%v\n", err, string(decodeBytes))
		return
	}
	fmt.Printf("decodeBytes spscail :%v\n", string(decodeBytes))

	fmt.Println()

}
