package main

import "fmt"

// int

func main() {

	// 十进制
	var i1 = 10
	fmt.Printf("%d\n", i1)
	fmt.Printf("%b\n", i1) // 10 to 2
	fmt.Printf("%o\n", i1) // 10 to 8
	fmt.Printf("%x\n", i1) // 10 to 16

	// 八进制
	i2 := 077
	fmt.Printf("%d\n", i2)

	// 16进制
	i3 := 0xffcc33
	fmt.Printf("%d\n", i3)

	// 查看变量的类型
	i4 := int8(9)
	fmt.Printf("%T\n", i4)

}
