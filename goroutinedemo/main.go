package main

import (
	"fmt"
	"time"
)

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("test() 你好Golang:", i)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {

	/***

	携程 加载会未完成 到5
	main() 你好Golang: 0
	test() 你好Golang: 0
	main() 你好Golang: 1
	test() 你好Golang: 1
	main() 你好Golang: 2
	main() 你好Golang: 3
	test() 你好Golang: 2
	main() 你好Golang: 4
	main() 你好Golang: 5
	test() 你好Golang: 3
	main() 你好Golang: 6
	test() 你好Golang: 4
	main() 你好Golang: 7
	main() 你好Golang: 8
	test() 你好Golang: 5
	main() 你好Golang: 9


	*/

	go test()
	for i := 0; i < 10; i++ {
		fmt.Println("main() 你好Golang:", i)
		time.Sleep(time.Millisecond * 50)
	}
}
