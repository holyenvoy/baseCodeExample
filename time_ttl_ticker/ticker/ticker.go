package main

import (
	"fmt"
	"time"
)

func main() {

	var ttl time.Duration
	var d int
	d = 1
	ttl = time.Duration(d) * time.Minute
	if sec := ttl.Seconds(); sec > 0 {
		ttlString := fmt.Sprintf("%ds", int64(ttl.Seconds()))
		fmt.Printf("sec:%v\n", sec)
		fmt.Printf("ttlString:%v\n", ttlString)
	} else {
		fmt.Printf("sec:%v\n", sec)
	}
	fmt.Printf("%s\n", "-------------------------")

	var warnTicker *time.Ticker
	go func() {
		warnTicker = time.NewTicker(time.Second)
		for t := range warnTicker.C {
			//a.processCirclesWarning(t.Unix())
			fmt.Println(t.Unix())
		}
	}()

	time.Sleep(10 * time.Minute)
}
