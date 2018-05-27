package main

import (
	"time"
)

func main() {
	message_chan := make(chan int, 1)

	go func() {
		time.Sleep(time.Second * 3)
		println("start recv...")
		println(<-message_chan)
		println(<-message_chan)
		println(<-message_chan)
		println("finish recv...")
	}()

	println("start send 10...")
	message_chan <- 10

	println("start send 20...")
	message_chan <- 20

	println("start send 30...")
	message_chan <- 30

	println("finish send...")

	time.Sleep(time.Second * 3)
	close(message_chan)
}
