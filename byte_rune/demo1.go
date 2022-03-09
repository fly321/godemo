package main

import "fmt"

func main() {
	age := 19

	if age > 18 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

	// 作用域
	// cc 变量此时只在if判断语句中生效
	if cc := 19; cc > 18 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

	//fmt.Println(cc) 这里是找不到cc的 作用域问题

	// 5-9
	var i = 5
	/*	for ; i < 10; i++ {
		fmt.Println(i)
	}*/

	// 5-9
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// 无限循环
	/*for {
		// todo...
	}*/

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

}
