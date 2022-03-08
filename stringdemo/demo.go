package main

import "fmt"

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

}
