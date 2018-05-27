package main

import (
	"io"
	"net/http"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hi,Allen, Hello, world!\n")
}

/*
每次写 Handler 的时候，都要定义一个类型，然后编写对应的 ServeHTTP 方法，
这个步骤对于所有 Handler 都是一样的。重复的工作总是可以抽象出来，net/http 也
这么做了，它提供了 http.HandleFunc 方法，允许直接把特定类型的函数作为 handler

HandleFunc registers the handler function for the given pattern in the DefaultServeMux.
*/

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":12345", nil)
}
