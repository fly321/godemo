package main

import (
	"fmt"
	"unicode"
)

func main() {
	s1 := "hellofly哈哈嘻嘻"
	s2 := 0
	for _, v := range s1 {
		if unicode.Is(unicode.Han, v) {
			s2++
		}
	}

	fmt.Println(s2)

}
