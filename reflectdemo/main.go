package main

import (
	"fmt"
	"reflect"
)

// 通过反射获取变量类型
func reflectFn(x interface{}) {
	of := reflect.TypeOf(x)
	fmt.Println(of)
}

func main() {
	a := 10
	reflectFn(a)
	b := "3333"
	reflectFn(b)
}
