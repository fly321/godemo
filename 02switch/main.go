package main

import "fmt"

func main() {

	switch n := 2; n {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	default:
		fmt.Println(3)
	}

}
