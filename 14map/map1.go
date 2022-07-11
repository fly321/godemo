package main

import "fmt"

func sumFn(x ...int) (sum int) {
	sum = 0
	for _, v := range x {
		sum += v
	}
	return sum
}
func main() {
	num := sumFn(1, 2, 3, 4, 5, 6)
	fmt.Println(num)
}
