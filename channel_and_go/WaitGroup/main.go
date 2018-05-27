package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var retryBackoffDuration = time.Millisecond * 30

func processMsgIndex(value string) bool {

	if value == "allen" {
		return false
	}
	fmt.Printf("get value:%v\n", value)
	return true
}

func main() {

	var (
		wg        sync.WaitGroup
		failCount uint32
	)

	var hashMap = []string{"12", "23", "allen", "34"}

	for _, value := range hashMap {
		wg.Add(1)
		go func(valueFunc string) {
			defer wg.Done()
			for i := 1; i <= 3; i++ {
				if processMsgIndex(valueFunc) {
					break
				}
				time.Sleep(retryBackoffDuration * time.Duration(i+1))
				fmt.Printf("sleep:%v\n", retryBackoffDuration*time.Duration(i+1))

				if i == 3 {
					atomic.AddUint32(&failCount, 1)
				}
			}
		}(value)
	}
	wg.Wait()
	if atomic.LoadUint32(&failCount) != 0 {
		fmt.Printf("failCount:%v", failCount)
	}

}
