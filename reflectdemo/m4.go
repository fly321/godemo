package main

import (
	"fmt"
	"reflect"
)

type Student12 struct {
	Name  string `json:"name1"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

func main() {
	s2 := Student12{
		"呼来来",
		11,
		12,
	}

	StructChange(&s2)

	fmt.Printf("%#v\n", s2)
	//var a int = 10
	//StructChange(&a)
}

func StructChange(s interface{}) {
	t := reflect.TypeOf(s)
	//fmt.Println(t.Kind())  // ptr

	if t.Kind() != reflect.Ptr || (t.Elem().Kind() != reflect.Struct) {
		fmt.Println("传入的参数不是一个结构体 or 指针类型")
		return
	}

	v := reflect.ValueOf(s)

	name := v.Elem().FieldByName("Name")
	name.SetString("桃几ovo")

	age := v.Elem().FieldByName("Age")
	age.SetInt(15)

	score := v.Elem().FieldByName("Score")
	score.SetInt(16)
}
