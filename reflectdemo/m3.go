package main

import (
	"fmt"
	"reflect"
)

type Student1 struct {
	Name  string `json:"name1"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

func (s Student1) getInfo() string {
	str := fmt.Sprintf("姓名：%v,年龄:%v,积分:%v\n", s.Name, s.Age, s.Score)
	return str
}

func (s Student1) setInfo(name string, age, score int) {
	s.Name = name
	s.Age = age
	s.Score = score
}

func (s Student1) Print() {
	fmt.Println("这是一个打印方法")
}

func PrintStructField(s interface{}) {
	t := reflect.TypeOf(s)

	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("传入的参数不是一个结构体")
	}

	field0 := t.Field(0)
	fmt.Printf("%#v\n", field0)

	fmt.Printf("字段名称：%v\n", field0.Name)
	fmt.Printf("字段类型：%v\n", field0.Type)
	fmt.Printf("字段tag:%v\n", field0.Tag.Get("json"))

}

func main() {
	s1 := Student1{
		"呼来来",
		11,
		12,
	}
	PrintStructField(s1)
}
