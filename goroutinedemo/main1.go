package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func test1() {
	for i := 0; i < 10; i++ {
		fmt.Println("test() 你好Golang:", i)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done()
}

func main() {
	wg.Add(1)
	go test1()
	wg.Add(1)
	go test1()
	for i := 0; i < 10; i++ {
		fmt.Println("main() 你好Golang:", i)
		time.Sleep(time.Millisecond * 50)
	}
	wg.Wait()
	fmt.Println("主线程退出")
}
