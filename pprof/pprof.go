package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"
	"sync"
	"time"
)

var cpuProfile = flag.String("cpuProfile", "", "write cpu profile to file")

func startCPUProfile() {
	if *cpuProfile != "" {
		f, err := os.Create(*cpuProfile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can not create cpu profile output file: %s",
				err)
			return
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			fmt.Fprintf(os.Stderr, "Can not start cpu profile: %s", err)
			f.Close()
			return
		}
	}
}

func stopCPUProfile() {
	if *cpuProfile != "" {
		pprof.StopCPUProfile() // 把记录的概要信息写到已指定的文件
	}
}

func main() {
	flag.Parse()

	startCPUProfile()

	//这里实现了远程获取pprof数据的接口
	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()

	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go work(&wg)
	}
	time.Sleep(3 * time.Second)
	defer stopCPUProfile()

	wg.Wait()
	// Wait to see the global run queue deplete.
	time.Sleep(3 * time.Second)

}

func work(wg *sync.WaitGroup) {
	time.Sleep(time.Second)

	var counter int
	for i := 0; i < 1e10; i++ {
		time.Sleep(time.Millisecond * 100)
		counter++
	}
	wg.Done()
}
