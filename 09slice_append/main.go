package main

import "fmt"

func main() {
	// 追加数组
	s2 := []string{"笑哈哈", "哈哈度"}
	// s2[] = “jhaja" 错误写法
	s2 = append(s2, "叽里呱啦")

	fmt.Println(s2)
}
