package main

import "fmt"

func main() {
	//math.MaxFloat32 // float32最大值

	f1 := 1.235
	// 默认go语言中的小数都是float64
	fmt.Printf("%T\n", f1)

	f2 := float32(1.23456)
	fmt.Printf("%T\n", f2)

}
