package main

import "fmt"

func main() {
	var m1 map[string]int
	m1 = make(map[string]int, 10)

	m1["hehe"] = 100
	m1["huha"] = 200

	fmt.Println(m1)

	v, ok := m1["jhah"]
	if !ok {
		fmt.Println("key不存在")
	} else {
		fmt.Println(v)
	}
}
