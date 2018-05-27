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

	//
	//
	//
	birthday, _ := time.Parse("2006-01-02", "2016-11-11")
	fmt.Printf("birthday:%v\n", birthday.Unix())
	layout := "2006-01-02 15:04:05"
	dateStr := "2015-12-14 00:00:00"
	timestamp1, _ := time.Parse(layout, "2016-09-23 20:52:49")
	timestamp2, _ := time.ParseInLocation(layout, dateStr, time.Local)
	fmt.Println(time.Local)
	fmt.Println(timestamp1, timestamp2)               //2015-12-14 00:00:00 +0000 UTC   2015-12-14 00:00:00 +0800 CST
	fmt.Println(timestamp1.Unix(), timestamp2.Unix()) //1450051200   1450022400
	fmt.Printf("%s\n", "-------------防------------")

	//
	//
	//
	now := time.Now()
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())

	fmt.Println(time.Now().Format("2006-01-02"))

	now = time.Now()
	year, month, day := now.Date()
	today_str := fmt.Sprintf("%d-%d-%d 00:00:00", year, month, day)
	today_time := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	fmt.Printf("today_str :%v \n", today_str)
	fmt.Printf("today_time :%v \n", today_time)

	str_time := time.Unix(now.Unix(), 0).Format("2006-01-02")
	fmt.Println(str_time)

	tagCreateTime, _ := time.Parse(time.RFC3339, "2016-09-23T20:52:49.429729339Z")
	tagCreateTime.In(time.Local)
	fmt.Printf("tagCreateTime :%v \n", tagCreateTime)

	// CST=UTC/GMT +8 小时
	tm := time.Unix(tagCreateTime.Unix(), 0)
	timesParseUTC := tm.Format(layout)
	fmt.Printf("timesParseUTC :%v \n", timesParseUTC)

}
