package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"io"
	"strconv"
)

/*
 SHA1是由NISTNSA设计为同DSA一起使用的，它对长度小于264的输入，产生长度为160bit的散列值，因此抗穷举(brute-force)性更好。
SHA-1设计时基于和MD4相同原理,并且模仿了该算法。SHA-1是由美国标准技术局（NIST）颁布的国家标准，是一种应用最为广泛的hash函数算法，
也是目前最先进的加密技术，被政府部门和私营业主用来处理敏感的信息。而SHA-1基于MD5，MD5又基于MD4。
*/

func main() {

	//sha1
	var a, b int64
	a = 103473275
	b = 104123736
	c := (a + b)
	d := strconv.FormatInt(c, 10)
	fmt.Println(c)
	fmt.Println(d)
	h := sha1.New()
	io.WriteString(h, string(d))
	fmt.Printf("get sha1 value is: %v\n", h.Sum(nil))
	bits := binary.LittleEndian.Uint32(h.Sum(nil))
	fmt.Printf("--- byte to int64:%v\n\n", int64(bits))

	//
	//
	//hmac ,use sha1
	key := []byte("123456")
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte("aaaaaa"))
	fmt.Printf("get hmac value is:  %x\n", mac.Sum(nil))
	bits = binary.LittleEndian.Uint32(mac.Sum(nil))
	fmt.Printf("--- byte to int64:%v\n\n", int64(bits))

}
