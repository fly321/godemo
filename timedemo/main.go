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
	// 2006 表示年份的格式 01 表示 月 02 表示日 03 表示时 04 表示分 05 表示秒
	format := now.Format("2006-01-02 03:04:05")
	fmt.Println(format)

}
