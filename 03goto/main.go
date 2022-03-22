package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		for j := 0; j <= 10; j++ {
			if j == 2 {
				goto gotoTag
			}
			fmt.Println(i, j)
		}
	}
	/**
	0 0
	0 1
	xxx
	*/
gotoTag:
	fmt.Println("狂喊老弟")
}
