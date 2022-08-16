package main

import "os"

func main() {
	// 单层创建目录
	//os.Mkdir("./abc", 0777)
	// 多曾创建目录
	//os.MkdirAll("./abc/def", 0777)

	//os.Remove("./abc")
	//os.RemoveAll("./abc/def")
	//os.RemoveAll("./abc")

	os.Rename("./1.txt", "./2.txt")
}
