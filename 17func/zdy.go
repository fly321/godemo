package main

import "fmt"

// 自定义类型
type calc func(int, int) int

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func main() {
	var c calc
	c = add
	fmt.Printf("c类型是%T\n", c)
}
