package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "中国在地球上"
	//chars := []rune(str)
	fmt.Println(utf8.DecodeLastRuneInString(str))
}
