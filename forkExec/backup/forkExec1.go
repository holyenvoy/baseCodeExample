// Package main provides ...
package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var buildAt string
var buildOn string
var gitCommit string

var verCmd = flag.Bool("version", false, "print build version")
var parent_flag = flag.Bool("P", false, "Parent Watch Mode")
var stopping bool = false
var childPid int = 0
var childPid2 int = 0

func ParentSignalMonitor() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGSTOP)
	for {
		s := <-c
		fmt.Printf("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT:
			if childPid > 0 {
				syscall.Kill(childPid, syscall.SIGQUIT)
			}
			stopping = true
			return
		case syscall.SIGHUP:
			if childPid > 0 {
				syscall.Kill(childPid, syscall.SIGHUP)
			}
		default:
			return
		}
	}
}

//register signals handler.
func SignalMonitor() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGSTOP)
	for {
		s := <-c
		fmt.Printf("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT:
			return
		case syscall.SIGHUP:
			//reload()
		default:
			return
		}
	}
}

func main() {

	flag.Parse()
	if *verCmd {
		fmt.Println("Build At: ", buildAt)
		fmt.Println("Build On: ", buildOn)
		fmt.Println("Git Commit: ", gitCommit)
		return
	}

	if *parent_flag == true {

		go ParentSignalMonitor()
		var i int
		for i = 0; i < len(os.Args); i++ {
			if os.Args[i] == "-P" {
				break
			}
		}

		var failCount int = 0
		for {

			var proc_attr syscall.ProcAttr
			var sys_attr syscall.SysProcAttr
			var waitStatus syscall.WaitStatus

			tNow := time.Now()
			stdErrFileName := fmt.Sprintf("logs/cgc.log.%04d%02d%02d%02d%02d%02d", tNow.Year(),
				tNow.Month(), tNow.Day(), tNow.Hour(), tNow.Minute(), tNow.Second())
			stdErrFile, err := os.OpenFile(stdErrFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
			if err != nil {
				fmt.Printf("os.OpenFile,err:%v", err)
			}

			sys_attr.Ptrace = false
			proc_attr.Sys = &sys_attr
			proc_attr.Files = []uintptr{uintptr(syscall.Stdin), uintptr(syscall.Stdout), stdErrFile.Fd()}
			proc_attr.Env = os.Environ()

			pid, err := syscall.ForkExec(os.Args[0], append(os.Args[0:i], os.Args[i+1:]...), &proc_attr)
			childPid = pid
			if err != nil {
				fmt.Printf("syscall.ForkExec:err:%v", err)
			}
			//fmt.Printf("syscall.ForkExec:err:%v", err)
			syscall.Wait4(pid, &waitStatus, 0, nil)

			if time.Now().Sub(tNow).Seconds() < float64(30) {
				failCount++
			} else {
				failCount = 0
			}

			if failCount > 5 {
				fmt.Printf("Parent process exited. failCount > 5")
				//quit
				return
			}
			if stopping {
				fmt.Printf("Parent process exited.")
				return
			}
		}
	} else {
		Init()
		SignalMonitor()
		Stop()
	}
}

func Init() {
	fmt.Printf("start init...\n")

}

func Stop() {
	fmt.Printf("start shutdown...\n")
}
