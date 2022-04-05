package main

import "fmt"

func main() {
	// &: 获取内存地址
	// * : 根据地址取值

	n := 10
	p := &n
	fmt.Println(p)
	fmt.Printf("%T\r\n", p)

	m := *p
	fmt.Println(m)

	// new创建一个内存地址
	a1 := new(int)
	*a1 = 13
	fmt.Println(*a1)

}
