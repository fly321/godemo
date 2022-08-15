package main

import (
	"io/ioutil"
)

func main() {
	ioutil.WriteFile("./1.txt", []byte("hello world\n"), 0777)
}
