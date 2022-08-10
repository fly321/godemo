package main

import (
	"fmt"
	"time"
)

func main() {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	stringChan := make(chan string, 10)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}

	// 不需要close

	for {
		select {
		case v := <-intChan:
			fmt.Printf("从int 读取%d\n", v)
			time.Sleep(time.Millisecond * 50)

		case v := <-stringChan:
			fmt.Printf("从string 读取%v\n", v)
			time.Sleep(time.Millisecond * 50)
		default:
			fmt.Println("获取完毕")
			// 退出
			return
		}

	}

	/**

	从int 读取0
	从string 读取hello0
	从int 读取1
	从string 读取hello1
	从int 读取2
	从string 读取hello2
	从string 读取hello3
	从string 读取hello4
	从int 读取3
	从int 读取4
	从int 读取5
	从int 读取6
	从int 读取7
	从int 读取8
	从int 读取9
	获取完毕


	*/

}
