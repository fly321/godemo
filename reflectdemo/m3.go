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

func (s Student1) GetInfo() string {
	str := fmt.Sprintf("姓名：%v,年龄:%v,积分:%v\n", s.Name, s.Age, s.Score)
	return str
}

func (s *Student1) SetInfo(name string, age, score int) {
	s.Name = name
	s.Age = age
	s.Score = score
}

func (s Student1) Print() {
	fmt.Println("这是一个打印方法")
}

func PrintStructField1(s interface{}) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("传入的参数不是一个结构体")
		return
	}

	field0 := t.Field(0)
	fmt.Printf("%#v\n", field0)

	fmt.Printf("字段名称：%v\n", field0.Name)
	fmt.Printf("字段类型：%v\n", field0.Type)
	fmt.Printf("字段tag:%v\n", field0.Tag.Get("json"))

	field1, ok := t.FieldByName("Age")
	if ok {
		fmt.Printf("字段名称：%v\n", field1.Name)
		fmt.Printf("字段类型：%v\n", field1.Type)
		fmt.Printf("字段tag:%v\n", field1.Tag.Get("json"))
	}

	// 调用方法
	info := v.MethodByName("GetInfo").Call(nil)
	fmt.Println(info)

}

func PrintStructField(s interface{}) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("传入的参数不是一个结构体")
		return
	}

	field0 := t.Elem().Field(0)
	fmt.Printf("%#v\n", field0)

	fmt.Printf("字段名称：%v\n", field0.Name)
	fmt.Printf("字段类型：%v\n", field0.Type)
	fmt.Printf("字段tag:%v\n", field0.Tag.Get("json"))

	field1, ok := t.Elem().FieldByName("Age")
	if ok {
		fmt.Printf("字段名称：%v\n", field1.Name)
		fmt.Printf("字段类型：%v\n", field1.Type)
		fmt.Printf("字段tag:%v\n", field1.Tag.Get("json"))
	}

	// 调用方法
	info := v.MethodByName("GetInfo").Call(nil)
	fmt.Println(info)

	// 执行带参方法
	var params []reflect.Value
	params = append(params, reflect.ValueOf("lisi"))
	params = append(params, reflect.ValueOf(20))
	params = append(params, reflect.ValueOf(13))
	fmt.Println(params)
	v.MethodByName("SetInfo").Call(params)

	info2 := v.MethodByName("GetInfo").Call(nil)
	fmt.Println(info2)

}

func main() {
	s1 := Student1{
		"呼来来",
		11,
		12,
	}
	PrintStructField1(s1)
	PrintStructField(&s1)
}
