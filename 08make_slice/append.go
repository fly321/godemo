package main

import "fmt"

func main() {
	// 创建切片
	s1 := make([]int, 3, 5)
	fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))

	// 切片追加元素
	s1 = append(s1, 1, 2, 3)
	fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
}
