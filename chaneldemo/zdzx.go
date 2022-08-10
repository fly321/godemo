package main

func main() {
	// 2. 管道声明为只写
	ch1 := make(chan<- int, 2)
	ch1 <- 10
	// chaneldemo\zdzx.go:7:4: invalid operation: cannot receive from send-only channel ch1 (variable of type chan<- int)
	//<-ch1

	// 只读
	ch2 := make(<-chan int, 2)
	/**
	chaneldemo\zdzx.go:12:2: invalid operation: cannot send to receive-only channel ch2 (variable of type <-chan int)
	*/
	//ch2 <- 33
	<-ch2

}
