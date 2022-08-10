package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func fn1(ch chan int) {
	for i := 1; i <= 100; i++ {
		ch <- i
		fmt.Printf("正在写入：%v\n\n", i)
		time.Sleep(time.Millisecond * 50)
	}
	close(ch)
	wg.Done()
}

func fn2(ch chan int) {
	for v := range ch {
		time.Sleep(time.Millisecond * 50)
		fmt.Printf("正在读取：%v\n\r", v)
	}
	wg.Done()

}
func main() {
	/**
	正在写入：1

	正在写入：2

	正在读取：1
	正在写入：3

	正在读取：2
	正在读取：3
	正在写入：4

	正在写入：5

	正在读取：4
	正在读取：5
	正在写入：6

	正在写入：7

	正在读取：6
	正在读取：7
	正在写入：8

	正在写入：9

	正在读取：8
	正在读取：9
	正在写入：10

	正在读取：10
	退出

	*/
	var ch = make(chan int, 10)
	wg.Add(1)
	go fn1(ch)
	wg.Add(1)
	go fn2(ch)
	wg.Wait()
	fmt.Println("退出")
}
