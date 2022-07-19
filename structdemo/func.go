package main

import "fmt"

type Person1 struct {
	Name string
	Age  int
	Sex  string
}

func (p Person1) printInfo() {
	fmt.Printf("测试数据%#v", p)
}

func main() {
	var p1 = Person1{
		Name: "嘻嘻",
		Age:  11,
		Sex:  "难",
	}
	p1.printInfo()
}
