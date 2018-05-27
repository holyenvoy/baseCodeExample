package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
)

func createGzip(content []byte) {

	// Create a buffer to write our archive to.
	// Make sure to check the error on Close.
	gzipBuffer := new(bytes.Buffer)
	gzipWriter := gzip.NewWriter(gzipBuffer)

	_, err := gzipWriter.Write(content)
	if err != nil {
	}
	gzipWriter.Close()

	gzipWriter.Flush()
	// [31 139 8 0 0 9 110 136 0 255 74 76 74 78 1 4 0 0 255 255 17 205 130 237 4 0 0 0] -- has close, has fulsh
	// [31 139 8 0 0 9 110 136 0 255 74 76 74 78 1 0 0 0 255 255] -- has flush, no close
	// [31 139 8 0 0 9 110 136 0 255] -- no close. no fulsh

	fmt.Printf("gzip size: %v\n", len(gzipBuffer.Bytes()))
	fmt.Printf("gzip bytes: %v\n", (gzipBuffer.Bytes()))
	//buffer := bytes.NewBuffer(gzipBuffer.Bytes())
	buffer := bytes.NewBuffer(content)
	r, err := gzip.NewReader(buffer)
	if err != nil {
		fmt.Printf("new reader err:%v", err)
		return
	}
	defer r.Close()
	ret, _ := ioutil.ReadAll(r)
	fmt.Printf("unzip bytes: %v\n", ret)

}

func main() {

	//
	// --------------------------------------------
	// 				打印字节数
	// --------------------------------------------
	//
	arr := [...]int{11, 22, 33, 44, 55}
	for key, value := range arr {
		fmt.Printf("%v %v \n", key, value)
	}
	fmt.Println(len(arr) - 1)

	//
	// --------------------------------------------
	// 				gzip压缩解压
	// --------------------------------------------
	//
	var buffer bytes.Buffer

	// 原始字节数据
	var originalByte []byte = []byte{1, 2, 3, 4}

	// gzip 压缩
	writer := gzip.NewWriter(&buffer)
	defer writer.Close()
	writer.Write(originalByte)
	writer.Flush()
	// gzip size:20 byte: [31 139 8 0 0 9 110 136 0 255 98 100 98 102 1 0 0 0 255 255]

	fmt.Printf("gzip size:%v byte: %v\n", len(buffer.Bytes()), buffer.Bytes())

	// gzip 解压
	reader, err := gzip.NewReader(&buffer)
	defer reader.Close()
	undatas, _ := ioutil.ReadAll(reader)
	fmt.Printf("ungzip size:%v, undatas:%v err:%v\n\n", len(undatas), undatas, err)

	//
	// --------------------------------------------------------------------
	// 				gzip压缩解压. str转换为[]byte后压缩解压
	// --------------------------------------------------------------------
	//
	s1 := "abcd"
	b1 := []byte(s1)
	fmt.Printf("\noriginal byte: %v\n", b1)
	createGzip(b1)
}
