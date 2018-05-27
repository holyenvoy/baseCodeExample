package main

import (
	"io"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hi,Allen, Hello, world!\n")
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, r.URL.Path+"Allen.Wu\n")
}

/*
HTTP 接口会有许多的 URL 和对应的 Handler。这里就要讲 net/http 的另外一个重要的概念：ServeMux。
Mux 是 multiplexor 的缩写，就是多路传输的意思（请求传过来，根据某种判断，分流到后端多个不同的地方）。
ServeMux 可以注册多了 URL 和 handler 的对应关系，并自动把请求转发到对应的 handler 进行处理
*/
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/", echoHandler)

	http.ListenAndServe(":12345", mux)
}
