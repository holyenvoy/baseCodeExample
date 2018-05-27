package main

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	h := md5.New()
	io.WriteString(h, "1")
	// c4ca4238a0b923820dcc509a6f75849b
	num := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Printf("\nget md5 num:%v\n", num)

	md5Ctx := md5.New()
	md5Ctx.Write([]byte("123456"))
	cipherStr := md5Ctx.Sum(nil)
	fmt.Printf("get md5 ctx value: %v\n", cipherStr)
	fmt.Printf("get md5 ctx=>hex.EncodeToString value: %v\n", hex.EncodeToString(cipherStr))
	bits := binary.LittleEndian.Uint32(cipherStr)
	fmt.Printf("--- byte to int64:%v\n", int64(bits))

}
