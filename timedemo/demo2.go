package main

import (
	"fmt"
	"time"
)

func main() {
	// 定时器 每隔一秒
	ticker := time.NewTicker(time.Second)
	i := 5
	for item := range ticker.C {
		i--
		fmt.Println(item.Format("2006-01-02 15:04:05"))
		if i == 0 {
			// 关闭定时器
			ticker.Stop()
			break
		}
	}
	fmt.Println("over")
}
