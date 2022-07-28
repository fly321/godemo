package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string
	Age  int
}

func main() {
	s1 := Student{
		Name: "呼啦啦",
		Age:  11,
	}
	marshal, err := json.Marshal(s1)
	if err != nil {
		fmt.Errorf("json.Marshal failed, err:%v", err)
	}
	fmt.Printf("json.Marshal:%s\n", string(marshal))
	//字符串转换切片
	var s2 Student
	err1 := json.Unmarshal(marshal, &s2)
	if err1 != nil {
		fmt.Errorf("json.Unmarshal failed, err:%v", err1)
	}
	fmt.Printf("json.Unmarshal:%#v\n", s2)
}
