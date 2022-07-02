package main

import "fmt"

func main() {
	s1 := "big"
	byteStr := []byte(s1)
	byteStr[0] = 'P'
	fmt.Println(string(byteStr))
}
