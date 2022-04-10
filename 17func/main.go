package main

import "fmt"

func main() {
	fmt.Println(sumxy(1, 2))
	f3()
	fmt.Println(f4())
}

/**
沒有返回值
*/
func f3() {
	fmt.Println("f3")
}

/**
返回一个int类型的值
*/
func f4() int {
	return 4
}

func sumxy(x int, y int) (ret int) {
	return x + y
}
