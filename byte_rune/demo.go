package main

import "fmt"

func main() {
	s := "hello流沙河"

	n := len(s)
	fmt.Println(n)

	/*for i := 0; i < n; i++ {
		//fmt.Println(s[i])
		fmt.Printf("%c\n", s[i])
	}*/

	for _, c := range s {
		//fmt.Println(c)
		fmt.Printf("%c\n", c)
	}

	// 字符串修改
	s2 := "陈儿"       // 分解
	s3 := []rune(s2) // 字符串强转rune切片
	s3[0] = '红'      // 替换单个 就是得用'' 是单个字符

	fmt.Println(string(s3))

	s4 := "嘻"
	s5 := '嘻'
	//s4:string,s5:int32
	fmt.Printf("s4:%T,s5:%T\n", s4, s5)

}
