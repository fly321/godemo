package main

import "fmt"

func f1() {
	fmt.Println("嘻嘻")
}

func f2() int {
	return 10
}

func f33(x func() int) int {
	a := x()
	return a
}

var f9 = func(x, y int) {
	fmt.Println(x + y)
}

func main() {
	a := f1
	fmt.Printf("%T\r\n", a)
	b := f2
	fmt.Printf("%T\r\n", b)
	c := f33(f2)
	fmt.Println(c)

	f9(100, 200)

	ret := f99(200)
	fmt.Println(ret(100))

}

/**
函数闭包传值
*/
func f99(x int) func(int) int {
	return func(i int) int {
		return x + i
	}
}
