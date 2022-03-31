package main

import "fmt"

func main() {
	var s1 []int // 定义一个存放int的切片
	var s2 []string

	fmt.Println(s1, s2)
	fmt.Println(s1 == nil) // true
	fmt.Println(s2 == nil) //true

	// 初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"嘻嘻", "哈哈", "上官"}

	fmt.Println(s1, s2)
	fmt.Println(s1 == nil) // false
	fmt.Println(s2 == nil) // false

	// 长度和容量
	fmt.Println("s1:len:%d,cap:%d", len(s1), cap(s1))
	fmt.Println("s2:len:%d,cap:%d", len(s2), cap(s2))

	// 由数组得到切片
	a2 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(a2[0:2]) // 左包含右不包含
}
