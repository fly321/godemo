package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	byteStr, err := ioutil.ReadFile("../filedemo/1.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile("./1.txt", byteStr, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("文件复制成功")

}
