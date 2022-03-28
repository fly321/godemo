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

}
