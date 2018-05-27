package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"reflect"
	//"runtime"
	"database/sql"
	"encoding/json"
	"os"
	"time"

	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
)

var mySQLUrlPattern = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4"

//插入demo
func insert() {
	db, err := sql.Open("mysql", "root:@/test?charset=utf8")
	checkErr(err)

	stmt, err := db.Prepare(`INSERT user (user_name,user_age,user_sex) values (?,?,?)`)
	checkErr(err)
	res, err := stmt.Exec("tony", 20, 1)
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
}

//查询demo
func query() []int {

	dbUrl := fmt.Sprintf(mySQLUrlPattern, "app", "@x#", "192.168.129.221", 3722, "testim")
	db, err := sql.Open("mysql", dbUrl)
	checkErr(err)

	rows, err := db.Query("select uid from users_recent_login")
	checkErr(err)

	uids := make([]int, 0)

	//普通demo
	for rows.Next() {
		var userId int
		rows.Columns()
		err = rows.Scan(&userId)
		checkErr(err)
		uids = append(uids, userId)

		//fmt.Println(userId)
	}

	return uids

	/*
		//字典类型
		//构造scanArgs、values两个数组，scanArgs的每个值指向values相应值的地址
		columns, _ := rows.Columns()
		scanArgs := make([]interface{}, len(columns))
		values := make([]interface{}, len(columns))
		for i := range values {
			scanArgs[i] = &values[i]
		}

		for rows.Next() {
			//将行数据保存到record字典
			err = rows.Scan(scanArgs...)
			record := make(map[string]string)
			for i, col := range values {
				if col != nil {
					record[columns[i]] = string(col.([]byte))
				}
			}
			fmt.Println(record)
		}

	*/
}

//更新数据
func update() {
	db, err := sql.Open("mysql", "root:@/test?charset=utf8")
	checkErr(err)

	stmt, err := db.Prepare(`UPDATE user SET user_age=?,user_sex=? WHERE user_id=?`)
	checkErr(err)
	res, err := stmt.Exec(21, 2, 1)
	checkErr(err)
	num, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(num)
}

//删除数据
func remove() {
	db, err := sql.Open("mysql", "root:@/test?charset=utf8")
	checkErr(err)

	stmt, err := db.Prepare(`DELETE FROM user WHERE user_id=?`)
	checkErr(err)
	res, err := stmt.Exec(1)
	checkErr(err)
	num, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(num)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// 重写生成连接池方法
func newPool(REDIS_DB int) *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000, // max number of connections
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "127.0.0.1:6379")
			if err != nil {
				panic(err.Error())
			}
			// 选择db
			c.Do("SELECT", REDIS_DB)
			return c, err
		},
	}
}

// 生成连接池
var pool = newPool(0)

func redisServer(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	// 从连接池里面获得一个连接
	c := pool.Get()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()
	dbkey := "netgame:info"
	if ok, err := redis.Bool(c.Do("LPUSH", dbkey, "yangzetao")); ok {
	} else {
		log.Print(err)
	}
	msg := fmt.Sprintf("用时：%s", time.Now().Sub(startTime))
	io.WriteString(w, msg+"\n\n")
}

type OnlineDuration struct {
	ActiveStatus bool   `json:"active_status"`
	Duration     int64  `json:"duration"`
	Timestamp    int64  `json:"timestamp"`
	Date         string `json:"date"`
	HasGain      bool   `json:"has_gain"`
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func startRedis(uids []int, filename string) {

	fmt.Printf("star redis :%v", len(uids))
	var f *os.File
	var err error
	var allDuration int64
	var onlineUids []int
	if checkFileIsExist(filename) { //如果文件存在
		f, err = os.OpenFile(filename, os.O_APPEND|os.O_RDWR, 0666) //打开文件
		fmt.Println("file is not exist")
		check(err)
	} else {
		f, err = os.Create(filename) //创建文件
		check(err)

	}

	// 从连接池里面获得一个连接
	c := pool.Get()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()

	nowDate := time.Unix(time.Now().Unix(), 0).Format("2006-01-02")

	for _, uid := range uids {
		dbkey := fmt.Sprintf("uonlined_%d", uid)
		result, err := redis.String(c.Do("GET", dbkey))

		onlineDuration := new(OnlineDuration)

		if err = json.Unmarshal([]byte(result), onlineDuration); err != nil {
			fmt.Errorf("get err:%v", err)
		}

		lastDate := onlineDuration.Date

		if nowDate == lastDate && (onlineDuration.Duration/60 > 0) {
			//fmt.Printf("uid :%v Duration: %v \n", uid, onlineDuration.Duration/60)

			onlineUids = append(onlineUids, uid)
			allDuration = allDuration + onlineDuration.Duration/60

			wireteString := fmt.Sprintf("uid :%v Duration: %v\r\n", uid, onlineDuration.Duration/60)
			_, err := io.WriteString(f, wireteString) //写入文件(字符串)
			check(err)
		}
	}

	averageStr := fmt.Sprintf("\n\nallDuration:%v uids:%v, average:%v\n", allDuration, len(uids), allDuration/int64(len(uids)))
	io.WriteString(f, averageStr)
	averageStr = fmt.Sprintf("\n\nallDuration:%v onlineUids:%v, average:%v\n", allDuration, len(onlineUids), allDuration/int64(len(onlineUids)))
	io.WriteString(f, averageStr)
	io.WriteString(f, "------over------")

}

func getProp(d interface{}, label string) (interface{}, bool) {
	switch reflect.TypeOf(d).Kind() {
	case reflect.Struct:
		v := reflect.ValueOf(d).FieldByName(label)
		return v.Interface(), true
	}
	return nil, false
}

func startHashRedis() {

	// 从连接池里面获得一个连接
	c := pool.Get()

	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()

	onlineDuration := new(OnlineDuration)

	onlineDuration.ActiveStatus = true
	onlineDuration.Timestamp = 123

	keyPattern := "hashKey"
	reply, err := c.Do("HSET", keyPattern, "allen1", onlineDuration)
	reply, err = c.Do("HSET", keyPattern, "allen2", onlineDuration)
	fmt.Printf("hset reply:%v, err:%v\n", reply, err)

	result := make(map[string]*OnlineDuration)
	reply, err = c.Do("HGETALL", keyPattern)

	fmt.Printf("hget all reply:%v, err:%v result:%v\n", reply, err, result)

}

func main() {
	/*
		uids := query()

		nowDate := time.Unix(time.Now().Unix(), 0).Format("2006-01-02")

		filename := fmt.Sprintf("/home/wdb/allen/task/online_%v", nowDate)

		startRedis(uids, filename)
	*/

	startHashRedis()
}
