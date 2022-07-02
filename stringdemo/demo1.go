package main

import "strings"

func main() {
	// 查找字符串是否存在
	s := "Hello, world"
	contains := strings.Contains(s, "world")
	println(contains)
}
