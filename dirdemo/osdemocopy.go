package main

import (
	"io"
	"os"
)

func copyfile(srcFileName string, dstFileName string) (err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	dstFile, err := os.OpenFile(dstFileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	var tempSlice = make([]byte, 1024)

	for {
		n1, err := srcFile.Read(tempSlice)

		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if _, err := dstFile.Write(tempSlice[:n1]); err != nil {
			return err
		}

	}

	return nil
}

func main() {
	copyfile("../filedemo/1.txt", "./2.txt")
}
