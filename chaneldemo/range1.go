package main

import "fmt"

func main() {
	var ch1 = make(chan int, 10)
	for i := 1; i <= 10; i++ {
		ch1 <- i
	}
	// 必须关闭不然会死锁 //fatal error: all goroutines are asleep - deadlock!
	close(ch1)
	for v := range ch1 {
		fmt.Println(v)
	}
}
