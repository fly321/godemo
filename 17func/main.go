package main

import "fmt"

func main() {
	fmt.Println(sumxy(1, 2))
}

func sumxy(x int, y int) (ret int) {
	return x + y
}
