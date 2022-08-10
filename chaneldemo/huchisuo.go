package main

import (
	"fmt"
	"sync"
	"time"
)

var count int = 0
var wg1 sync.WaitGroup

func test() {
	defer wg1.Done()
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
