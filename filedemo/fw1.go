package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("./1.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0777)
	defer file.Close()
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}

	n, err := file.Write([]byte("hello world1\n"))

	n1, _ := file.WriteString("hello,fly \n")
	if err != nil {
		return
	}

	if err != nil {
		fmt.Printf("write file failed, err:%v\n", err)
		return
	}

	fmt.Printf("write %d bytes\n", n)
	fmt.Printf("write %d bytes\n", n1)

}
