package main

import (
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"runtime/pprof"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	pid      int
	progname string
)

func init() {
	pid = os.Getpid()
	paths := strings.Split(os.Args[0], "/")
	paths = strings.Split(paths[len(paths)-1], string(os.PathSeparator))
	progname = paths[len(paths)-1]
	runtime.MemProfileRate = 1
}
func saveHeapProfile() {
	fmt.Printf("saveHeapProfile \n")
	runtime.GC()
	f, err := os.Create(fmt.Sprintf("prof/heap_%s_%d_%s.prof", progname, pid, time.Now().Format("2006_01_02_03_04_05")))
	if err != nil {
		return
	}
	defer f.Close()
	pprof.Lookup("heap").WriteTo(f, 1)
}
func waitForSignal() os.Signal {
	signalChan := make(chan os.Signal, 1)
	defer close(signalChan)
	signal.Notify(signalChan, os.Kill, os.Interrupt)
	s := <-signalChan
	signal.Stop(signalChan)
	return s
}
func connect(source string) *sql.DB {
	db, err := sql.Open("mysql", source)
	if err != nil {
		return nil
	}
	if err := db.Ping(); err != nil {
		return nil
	}
	return db
}

type User struct {
	uid       int
	name      string
	nick      string
	forbidden int
	cid       int
}

func query(db *sql.DB, name string, id int, dataChan chan *User) {
	for {
		time.Sleep(time.Millisecond)
		user := &User{
			cid:  id,
			name: name,
		}
		err := db.QueryRow("SELECT nickname, uid, forbidden FROM users WHERE login_name = ?", name).Scan(&user.nick, &user.uid, &user.forbidden)
		if err != nil {
			continue
		}
		dataChan <- user
	}
}
func main() {
	defer saveHeapProfile()
	db := connect("allen:123456@tcp(localhost:3306)/meituTestDB?charset=utf8")
	if db == nil {
		return
	}
	userChan := make(chan *User, 100)
	for i := 0; i < 100; i++ {
		go query(db, "Alex", i+1, userChan)
	}
	allUsers := make([]*User, 1<<12)
	go func() {
		for user := range userChan {
			fmt.Printf("routine[%d] get user %+v\n", user.cid, user)
			allUsers = append(allUsers, user)
		}
	}()
	time.Sleep(10 * time.Second)
	saveHeapProfile()
	return

	s := waitForSignal()
	fmt.Printf("signal got: %v, all users: %d\n", s, len(allUsers))
}
