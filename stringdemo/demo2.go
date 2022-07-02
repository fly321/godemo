package main

import "fmt"

func main() {
	s1 := "big"
	byteStr := []byte(s1)
	byteStr[0] = 'P'
	fmt.Println(string(byteStr))

	// 字符串有汉子转换
	s2 := "中国"
	runeStr := []rune(s2)
	runeStr[0] = '大'
	fmt.Println(string(runeStr))

}
