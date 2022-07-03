package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 字符串转数字
	s := "110"
	s1, _ := strconv.ParseInt(s, 10, 64)
	fmt.Printf("%v,%T\n", s1, s1)

	// 字符串转浮点数
	s2, _ := strconv.ParseFloat(s+".111", 64)
	fmt.Printf("%v,%T\n", s2, s2)

	// 数字转字符串
	s3 := fmt.Sprintf("%d", s1)
	fmt.Printf("%v,%T\n", s3, s3)
}
