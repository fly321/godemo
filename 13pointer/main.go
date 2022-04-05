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

}
