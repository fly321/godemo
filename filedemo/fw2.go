package main

import (
	"bufio"
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
	buffaloWriter := bufio.NewWriter(file)
	buffaloWriter.WriteString("hello world\n")
	buffaloWriter.Flush()

}
