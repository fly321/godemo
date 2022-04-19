package main

import "fmt"

func q1(f func()) {
	fmt.Println("q1")
	f()
}

func q2(x, y int) {
	fmt.Println("q2")
	fmt.Println(x + y)
}

func q3(f func(int, int), x, y int) func() {
	return func() {
		f(x, y)
	}
}
func main() {
	tmp := q3(q2, 100, 200)
	q1(tmp)
}
