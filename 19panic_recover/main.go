package main

import "fmt"

func t1() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("释放连接")
	}()
	panic("我是模拟的错误")
	// 不会执行
	fmt.Println("xixixi")
}

func t2() {
	fmt.Println("cc")
}

func main() {
	t1()
	t2()
}
