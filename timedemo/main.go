package main

import (
	"fmt"
	"time"
)

func main() {
	// 打印当前日期
	now := time.Now()
	fmt.Println(now)
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("%d-%02d-%02d %d:%d:%d\n", year, month, day, hour, minute, second)
	// 2006 表示年份的格式 01 表示 月 02 表示日 03 表示时【12小时制】 04 表示分 05 表示秒
	// 2006 表示年份的格式 01 表示 月 02 表示日 15 表示时【24小时制】 04 表示分 05 表示秒
	format := now.Format("2006-01-02 03:04:05")
	format1 := now.Format("2006-01-02 15:04:05")
	fmt.Println(format)
	fmt.Println(format1)

	// 获取当前时间戳
	unix := now.Unix()
	fmt.Println(unix)

	// 获取当前时间戳的纳秒
	nanosecond := now.UnixNano()
	fmt.Println(nanosecond)

	// 获取当前日期对象
	now1 := time.Unix(unix, 0)
	fmt.Println(now1.Format("2006-01-02 15:04:05"))

	// 日期字符串 转时间戳
	location, err := time.ParseInLocation("2006-01-02 15:04:05", "2019-01-01 00:00:00", time.Local)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(location.Unix())

	// 增加一小时
	now2 := now.Add(time.Hour)
	fmt.Println(now2.Format("2006-01-02 15:04:05"))
}
