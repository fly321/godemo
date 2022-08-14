package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	byteStr, err := ioutil.ReadFile("./main2.go")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(byteStr))

}
