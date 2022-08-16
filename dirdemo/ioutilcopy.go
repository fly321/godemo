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

	ioutil.WriteFile("./1.txt", byteStr, 0777)

}
