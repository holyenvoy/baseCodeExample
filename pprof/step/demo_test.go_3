package main

import (
	"bufio"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func BenchmarkHi(b *testing.B) {
	b.ReportAllocs()

	req, err := http.ReadRequest(bufio.NewReader(strings.NewReader("GET / HTTP/1.0\r\n\r\n")))
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		rw := httptest.NewRecorder()
		handleHi(rw, req)
	}
}
