package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
)

/*
HMAC是密钥相关的哈希运算消息认证码（Hash-based Message Authentication Code）,HMAC运算利用哈希算法，以一个密钥和一个消息为输入，
生成一个消息摘要作为输出。HMAC是需要一个密钥的。所以，HMAC_SHA1也是需要一个密钥的，而SHA1不需要。
*/

func main() {
	//hmac ,use sha1
	key := []byte("123456")
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte("aaaaaa"))
	fmt.Printf("get hmac value is:  %x\n", mac.Sum(nil))
	bits := binary.LittleEndian.Uint32(mac.Sum(nil))
	fmt.Printf("--- byte to int64:%v\n\n", int64(bits))
}
