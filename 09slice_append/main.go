package main

import "fmt"

func main() {
	// 追加数组
	s2 := []string{"笑哈哈", "哈哈度", "对嘟嘟"}
	fmt.Printf("s2:%v,len:%d,cap:%d\r\n", s2, len(s2), cap(s2))
	// s2[] = “jhaja" 错误写法
	s2 = append(s2, "叽里呱啦")
	fmt.Printf("s2:%v,len:%d,cap:%d\r\n", s2, len(s2), cap(s2))

	s2 = append(s2, "叽里1")
	fmt.Printf("s2:%v,len:%d,cap:%d\r\n", s2, len(s2), cap(s2))

	/**
	s2:[笑哈哈 哈哈度 对嘟嘟],len:3,cap:3
	s2:[笑哈哈 哈哈度 对嘟嘟 叽里呱啦],len:4,cap:6
	s2:[笑哈哈 哈哈度 对嘟嘟 叽里呱啦 叽里1],len:5,cap:6
	*/

}
