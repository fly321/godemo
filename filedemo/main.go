package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 打开文件
	file, err := os.Open("./main.go")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 读取文件内容
	//fmt.Println(file)
	var temp1 = make([]byte, 128)
	/*n, err := file.Read(temp1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("读取了%d个字节\n", n)

	s := string(temp1)
	fmt.Println(s)*/

	var strSlice []byte

	for {
		n, err := file.Read(temp1)
		if err == io.EOF {
			fmt.Println("文件读取完毕")
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("读取了%d个字节\n", n)
		strSlice = append(strSlice, temp1...)
	}

	fmt.Println(string(strSlice))

}
