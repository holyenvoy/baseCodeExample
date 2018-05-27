package main

import (
	"io"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":443")
	if err != nil {
		log.Fatal(err)
	}

	io.WriteString(conn, "Hello Server\n")
	conn.Close()
}
