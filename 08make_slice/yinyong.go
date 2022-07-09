package main

import "fmt"

func main() {
	// 引用切片
	s1 := []int{1, 2, 3}
	s2 := s1
	s2[0] = 100
	fmt.Println(s1)
	fmt.Println(s2)
}
