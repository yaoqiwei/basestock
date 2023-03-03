package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Unix(), now.UnixNano(), now.Nanosecond())

	year, month, day := now.Date()
	fmt.Println(year, month, day)

	fmt.Printf("year:%d, month:%d, day:%d\n", year, month, day)
	// 年
	fmt.Println(now.Year())
	// 月
	fmt.Println(now.Month())
	// 日
	fmt.Println(now.Day())
	// 时分秒
	hour, minute, second := now.Clock()
	fmt.Printf("hour:%d, minute:%d, second:%d\n", hour, minute, second)
	// 时
	fmt.Println(now.Hour())
	// 分
	fmt.Println(now.Minute())
	// 秒
	fmt.Println(now.Second())
	// 返回星期
	fmt.Println(now.Weekday())
	// 返回一年中对应的第几天
	fmt.Println(now.YearDay())
	// 返回时区
	fmt.Println(now.Location())
	// 返回一年中第几天
	fmt.Println(now.YearDay())
}
