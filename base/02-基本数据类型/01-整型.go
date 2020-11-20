package main

import "fmt"

func main() {
	//	10进制
	var a int = 100
	fmt.Printf("a = %d\n", a)
	// 8进制
	var b int = 012
	fmt.Printf("b = %o \n", b)

	// 十六进制  以0x开头
	var c int = 0xff
	fmt.Printf("c = %x\n", c)

	fmt.Println(0.2-0.1)
}
