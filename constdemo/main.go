package main

import "fmt"

// 常量
//const pi float64 = 3.1415926
const PI = 3.1415926

const (
	STATUSOK = 200
	NOTFOUND = 404
)

// 批量声明常量 不声明 = 就和上面的值一样 P2 100 , P3 100
const (
	P1 = 100
	P2
	P3
)

// iota a0 0 a1 1 a2 2 a3 3 自动累加
const (
	A0 = iota
	A1
	A2
	A3
)

const (
	a, b = iota + 1, iota + 2 //1,2
	c, d                      //2,3
	e, f                      //3,4
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	fmt.Println(PI)                //3.1415926
	fmt.Println(STATUSOK)          // 200
	fmt.Println(P2)                // 100
	fmt.Println(P3)                // 100
	fmt.Println(A0 + A1 + A2 + A3) // 6
}
