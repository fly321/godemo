package main

import "fmt"

func main() {
	// 创建管道
	ch := make(chan int, 3)

	// 给管道里面存储数据
	ch <- 10
	ch <- 21
	ch <- 32

	// 获取管道里面的内容
	a := <-ch
	fmt.Println(a) // 10
	<-ch           //21
	c := <-ch
	fmt.Println(c) // 32
	ch <- 56
	// 打印管道的长度容量
	fmt.Printf("值：%#v,类型：%#T,容量: %v, 长度 %v", ch, ch, cap(ch), len(ch))
	close(ch)
}
