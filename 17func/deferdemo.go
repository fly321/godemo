package main

import "fmt"

func deferdemo() {
	fmt.Println("xixi")
	// 函数执行完调用
	defer fmt.Println("我是中间的可我最后执行拉")
	defer fmt.Println("hahah") // 后进先出
	fmt.Println("jajha")
}

func main() {
	deferdemo()
}
