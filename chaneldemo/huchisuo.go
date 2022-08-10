package main

import (
	"fmt"
	"sync"
	"time"
)

var count int = 0
var wg1 sync.WaitGroup

// 声明互斥锁
var mutex sync.Mutex

func test() {
	mutex.Lock()
	defer wg1.Done()
	defer mutex.Unlock()
	count++
	fmt.Println("the count is : ", count)
	time.Sleep(time.Millisecond)
}

func main() {
	for r := 0; r < 100; r++ {
		wg1.Add(1)
		go test()
	}
	wg1.Wait()
}
