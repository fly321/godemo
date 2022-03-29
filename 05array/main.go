package main

import "fmt"

func main() {
	var arr1 [3]bool // [true,false,true]
	var arr2 [4]bool // [true,false,true,false]

	fmt.Printf("%T,%T \r\n", arr1, arr2)
	// 默认 false
	fmt.Println(arr1, arr2)

	// 初始化数据
	arr1 = [3]bool{true, false, true}
	arr2 = [4]bool{true, false, true, false}

	// 不知道数量写 [...]
	a2 := [...]int{1, 2, 3}
	fmt.Println(arr1, arr2, a2)

	// 初始化
	a3 := [5]int{0: 1, 4: 4}
	fmt.Println("a3:", a3)

	// 多维数组
	// [[1,2],[2,3],[3,4]]

	var a11 [3][2]int
	a11 = [3][2]int{
		{1, 2},
		{2, 3},
		{3, 4},
	}
	fmt.Println(a11)

}
