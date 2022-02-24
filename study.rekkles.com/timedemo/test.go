package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now() // local time
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())

	ret := time.Unix(1641575030, 0)
	fmt.Println(ret)
	fmt.Println(time.Second)

	fmt.Println(now.Add(24 * time.Hour))
	fmt.Println(ret.After(now))

	// per second run
	// timer := time.Tick(time.Second)
	// for t := range timer {
	// 	fmt.Println(t)
	// }

	// time format
	fmt.Println(now.Format("2006-1:2A3:04:05.000 AM"))

	// string to time
	timeObj, err := time.Parse("2006-1:2", "2000-01:02")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Println(timeObj)

	// Sub
	fmt.Println(now.Sub(timeObj))

	// Zone
	fmt.Println(now.UTC())

	// next day
	time.Parse("2006-01-02 15:04:05", "2022-01-19 01:45:00")
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("load location failed %v", err)
		return
	}
	timeObj2, err := time.ParseInLocation("2006-01-02 15:04:05", "2022-01-08 02:45:00", loc)
	if err != nil {
		fmt.Println("time parse failed %v", err)
		return
	}
	fmt.Println(timeObj2)
	fmt.Println(now)
	fmt.Println(timeObj2.Sub(now))

	// Sleep
	time.Sleep(5 * time.Second)

}
