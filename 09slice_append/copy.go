package main

import "fmt"

func main() {
	a1 := []int{1, 3, 5}
	a2 := a1 //引用的

	var a3 = make([]int, 3, 3)
	copy(a3, a1) // 复制的

	fmt.Println(a1, a2, a3)
	a1[0] = 100
	fmt.Println(a1, a2, a3)

	// 将a1中的3移除
	a1 = append(a1[:1], a1[2:]...)
	fmt.Println(a1)

}
