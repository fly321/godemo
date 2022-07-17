package main

import "fmt"

type Person struct {
	name string
	age  int
	sex  string
}

func main() {
	var p1 Person
	p1.name = "张三"
	p1.age = 11
	p1.sex = "男"
	fmt.Printf("值：%v,类型:%T", p1, p1)
	fmt.Printf("值：%#v,类型:%T", p1, p1)
}
