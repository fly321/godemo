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
}
