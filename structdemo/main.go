package main

import "fmt"

type Person struct {
	name string
	age  int
	sex  string
}

func main() {
	// 第一种方式创建
	var p1 Person
	p1.name = "张三"
	p1.age = 11
	p1.sex = "男"
	fmt.Printf("值：%v,类型:%T", p1, p1)
	fmt.Printf("值：%#v,类型:%T", p1, p1)

	// 第二种方式创建
	var p2 = new(Person)
	p2.age = 11
	p2.name = "回函三"
	p2.sex = "难"

	fmt.Printf("值：%v,类型:%T", p2, p2)
	fmt.Printf("值：%#v,类型:%T", p2, p2)

}
