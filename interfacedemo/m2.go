package main

import "fmt"

type NullInterface interface {
}

func main() {
	var a NullInterface
	var s = "hello"
	a = s
	fmt.Println(a)
}
