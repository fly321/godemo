package main

import "fmt"

// 声明变量

//var name string
//var age int
//var isOk bool

// 批量声明
var (
	name string
	age  int
	isOk bool
)

func main() {
	name = "hello fly"
	age = 11
	isOk = true
	fmt.Printf("nmae:%s\n", name)

	// 声明变量同时赋值
	var s1 string = "20"
	fmt.Println("声明变量同时赋值" + s1)

	// 类型推导 根据值 自动推导
	var s2 = "11"
	fmt.Println("值是：" + s2)

	// 简短变量声明 只能在函数内使用 局部变量
	s3 := "heheh"
	fmt.Println("简短变量声明：" + s3)

}
