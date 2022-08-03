package main

import "fmt"

type Nu1 interface{}

func main() {
	map1 := make(map[string]Nu1)
	map1["Name"] = "11"
	map1["age"] = 33
	fmt.Println(map1)

	var m2 Nu1
	m2 = "123"
	v, ok := m2.(string)
	fmt.Println(v, ok)

}
