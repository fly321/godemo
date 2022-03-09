package main

import (
	"fmt"
	"strings"
)

func main() {
	path := "d:\\ff\\cc\\1.txt"
	fmt.Println(path)

	s := "I'm ok"
	fmt.Println(s)

	s2 := `hello fly,hello xi`
	fmt.Println(s2)

	// 统计字符串长度
	fmt.Println(len(s))

	s3 := "嘿嘿"
	fmt.Println(s + s3)

	ss1 := fmt.Sprintf("%s%s", s, s3)
	fmt.Println(ss1)
	fmt.Printf("%s%s", s, s3)

	s4 := "我就在\\呼吸"

	// 分割
	ret := strings.Split(s4, "\\")
	fmt.Println(ret)

	// 包含
	fmt.Println(strings.Contains(s4, "huhu"))

	// 前缀
	fmt.Println(strings.HasPrefix(s4, "我"))

	// 后缀
	fmt.Println(strings.HasSuffix(s4, "我"))

	s5 := "abcdeb"
	fmt.Println(strings.Index(s5, "c"))
	fmt.Println(strings.LastIndex(s5, "b"))

	fmt.Println(strings.Join(ret, "+"))

}
