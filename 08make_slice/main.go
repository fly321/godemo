package main

import "fmt"

func main() {
	// make 创建 切片 make(数组类型,长度,容量)
	s1 := make([]int, 5, 10)
	fmt.Printf("长度:%d,容量：%d", len(s1), cap(s1))
}
