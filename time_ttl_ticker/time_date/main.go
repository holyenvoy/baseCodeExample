package main

import (
	"fmt"
	"time"
)

func timeSubV2(t1, t2 time.Time) int {

	if t1.Location().String() != t2.Location().String() {
		fmt.Printf("Location: %v\n", t1.Location().String())
		return -1
	}

	hours := t1.Sub(t2).Hours()

	if hours <= 0 {
		fmt.Printf("hours:%v\n", hours)
		return -1
	}

	// sub hours less than 24
	if hours < 24 {
		// may same day
		t1y, t1m, t1d := t1.Date()
		t2y, t2m, t2d := t2.Date()
		isSameDay := (t1y == t2y && t1m == t2m && t1d == t2d)

		if isSameDay {

			return 0
		} else {
			return 1
		}

	} else { // equal or more than 24

		if (hours/24)-float64(int(hours/24)) == 0 { // just 24's times

			fmt.Printf("ust 24's times:%v\n", hours)

			return int(hours / 24)
		} else { // more than 24 hours

			fmt.Printf("more than 24 hours:%v\n", hours)

			return int(hours/24) + 1
		}
	}
}

func timeSub(t1, t2 time.Time) int {

	if t1.Location().String() != t2.Location().String() {
		fmt.Printf("Location: %v\n", t1.Location().String())
		return -1
	}
	hours := t1.Sub(t2).Hours()

	if hours <= 0 {
		fmt.Printf("hours:%v\n", hours)
		return 0
	}
	// sub hours less than 24
	if hours < 24 {
		// may same day
		t1y, t1m, t1d := t1.Date()
		t2y, t2m, t2d := t2.Date()
		isSameDay := (t1y == t2y && t1m == t2m && t1d == t2d)

		if isSameDay {
			return 0
		} else {
			return 1
		}

	} else { // equal or more than 24

		if (hours/24)-float64(int(hours/24)) == 0 { // just 24's times

			//fmt.Printf("just 24's times:%v\n", hours)

			return int(hours / 24)
		} else { // more than 24 hours

			//fmt.Printf("more than 24 hours:%v\n", hours)

			return int(hours / 24)
		}
	}
}

func covertTime(timestamp int64) time.Time {
	layout := "2006-01-02 15:04:05"

	tm := time.Unix(timestamp, 0)
	timeString := tm.Format(layout)

	t, _ := time.Parse(layout, timeString)

	t = t.Truncate(time.Hour * 24)

	return t
}

func main() {
	layout := "2006-1-2 15:04:05"

	now := time.Now().Unix()

	nowTime := covertTime(now)
	updateAtTime := covertTime(1499961600)

	//nowTime = nowTime.Truncate(time.Hour * 24)
	//updateAtTime = updateAtTime.Truncate(time.Hour * 24)

	//1499443200
	//1499477580
	//1499702400
	//1499706000
	fmt.Printf("nowTime[%v] - updateAtTime[%v] %v\n", nowTime, updateAtTime, timeSub(nowTime, updateAtTime))

	tm := time.Unix(1499443200, 0)
	compare := tm.Format(layout)

	fmt.Printf("layout:%v\ncompare:%v\n", layout, compare)

	// just one second
	t1, _ := time.Parse(layout, "2017-7-10 18:59:59")
	t2, _ := time.Parse(layout, compare)

	fmt.Printf("t1:%v\nt2:%v\n -- 1 ----:%v\n", t1, t2, timeSub(t1, t2))

	// just one day
	t1, _ = time.Parse(layout, "2007-01-02 12:00:00")
	t2, _ = time.Parse(layout, "2007-01-03 00:22:59")
	fmt.Printf("t1:%v\nt2:%v\n -- 2 ----:%v\n", t1, t2, timeSub(t2, t1))

	if timeSub(t2, t1) != 1 {
		fmt.Printf("just one day should return 1\n")
	} else {
		fmt.Printf("one day sub:%v\n", timeSub(t2, t1))
	}

	// more than one day
	t1, _ = time.Parse(layout, "2007-01-02 23:59:59")
	t2, _ = time.Parse(layout, "2007-01-04 00:00:01")
	if timeSub(t2, t1) != 2 {
		fmt.Printf("just one day should return 2 :%v \n", timeSub(t2, t1))
	} else {
		fmt.Printf("sub:%v\n", timeSub(t2, t1))
	}

	// just 3 day
	t1, _ = time.Parse(layout, "2007-01-02 00:00:00")
	t2, _ = time.Parse(layout, "2007-01-05 00:00:00")
	if timeSub(t2, t1) != 3 {
		fmt.Printf("just 3 day should return 3\n")
	} else {
		fmt.Printf("sub:%v\n", timeSub(t2, t1))
	}

	// different month
	t1, _ = time.Parse(layout, "2007-01-02 00:00:00")
	t2, _ = time.Parse(layout, "2007-02-02 00:00:00")
	if timeSub(t2, t1) != 31 {
		fmt.Println(timeSub(t2, t1))
		fmt.Printf("just one month:31 days should return 31\n")
	} else {
		fmt.Printf("sub:%v\n", timeSub(t2, t1))
	}

	// 29 days in 2mth
	t1, _ = time.Parse(layout, "2000-02-01 00:00:00")
	t2, _ = time.Parse(layout, "2000-03-01 00:00:00")
	if timeSub(t2, t1) != 29 {
		fmt.Println(timeSub(t2, t1))
		fmt.Printf("just one month:29 days should return 29\n")
	} else {
		fmt.Printf("sub:%v\n", timeSub(t2, t1))
	}

	fmt.Println((6293275297509017601 >> 22) / 1000)
}
