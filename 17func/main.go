package main

import "fmt"

func main() {
	fmt.Println(sumxy(1, 2))
	f3()
	fmt.Println(f4())
	fmt.Println(f5("xixi", 2, 3, 4, 5))
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

func f5(title string, y ...int) (string, []int) {
	return title, y
}

func sumxy(x int, y int) (ret int) {
	return x + y
}
