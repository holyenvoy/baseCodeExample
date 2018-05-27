package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

var Client = &http.Client{
	Timeout: 3 * time.Second,
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hi,Allen, Hello, world!\n")

	//ret, err := DoBytesPost("http://192.168.1.3:33331/test/getBytes", []byte("test allen"))
	//io.WriteString(w, string(ret)+fmt.Sprintf("%v", err))

	decodeBytes, err := base64.StdEncoding.DecodeString("aGVsbG86YWxsZW5XdQ==")
	if err != nil {
		io.WriteString(w, fmt.Sprintf("decode err:", err))
	}

	resp, err := DoPostBytes("http://192.168.199.35:10101/audio_text", decodeBytes)
	if err != nil || resp == nil {
		io.WriteString(w, fmt.Sprintf("get err:", err))
	}
	body, err := ReadResponseBody(resp)
	if err != nil {
		io.WriteString(w, fmt.Sprintf("get err:", err))
	}

	io.WriteString(w, string(body)+fmt.Sprintf("%v", err))

}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":33355", nil)
}

//body提交二进制数据
/*
func DoBytesPost(url string, data []byte) ([]byte, error) {

	body := bytes.NewReader(data)
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		log.Println("http.NewRequest,[err=%s][url=%s]", err, url)
		return []byte(""), err
	}
	request.Header.Set("Connection", "Keep-Alive")
	var resp *http.Response
	resp, err = http.DefaultClient.Do(request)
	if err != nil {
		log.Println("http.Do failed,[err=%s][url=%s]", err, url)
		return []byte(""), err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("http.Do failed,[err=%s][url=%s]", err, url)
	}
	return b, err
}
*/

//body提交二进制数据
func DoPostBytes(url string, data []byte) (resp *http.Response, err error) {
	body := bytes.NewReader(data)

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Connection", "Keep-Alive")

	return Client.Do(req)
}

func ReadResponseBody(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), err
	}
	return body, err
}
