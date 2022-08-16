package main

import "os"

func main() {
	// 单层创建目录
	os.Mkdir("./abc", 0777)
	// 多曾创建目录
	os.MkdirAll("./abc/def", 0777)
}
