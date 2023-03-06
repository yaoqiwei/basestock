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

	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("15:04:05"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))

	layout := "2006-01-02 15:04:05"
	t := time.Unix(now.Unix(),0)    // 参数分别是：秒数,纳秒数
	fmt.Println(t.Format(layout))

	//根据指定时间返回 time.Time 类型
	//分别指定年，月，日，时，分，秒，纳秒，时区
	t = time.Date(2011, time.Month(3), 12, 15, 30, 20, 2, now.Location())
	fmt.Println(t.Format(layout))

	fmt.Println(time.Parse("2006-01-02 15:04:05", "2021-01-10 15:01:02"))
}
