package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "你好"
	fmt.Printf("长度是:%d\n", len(s))
	fmt.Printf("字符数量是:%d", utf8.RuneCountInString(s))

}
