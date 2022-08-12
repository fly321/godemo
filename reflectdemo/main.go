package main

import (
	"fmt"
	"reflect"
)

type myInt int
type n1 struct {
	Name string
	Age  int
}

// 通过反射获取变量类型
func reflectFn(x interface{}) {
	of := reflect.TypeOf(x)
	fmt.Printf("类型:%v,名称：%v,种类：%v\r\n", of, of.Name(), of.Kind())
}

func main() {
	a := 10
	reflectFn(a)
	b := "3333"
	reflectFn(b)

	var m myInt = 1
	reflectFn(m)

	m1 := n1{
		Name: "xixi",
		Age:  11,
	}
	reflectFn(m1)

	reflectFn(&m)

	m2 := [3]int{1, 2, 3}
	m3 := []int{11, 22, 33}
	reflectFn(m2)
	reflectFn(m3)

}
