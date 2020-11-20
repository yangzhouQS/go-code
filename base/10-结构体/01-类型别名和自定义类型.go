package main

import "fmt"

func main() {
	//	1.自定义类型
	type MyInt int

	// 2.类型别名
	type NewInt = int

	var a MyInt
	var b NewInt
	fmt.Printf("a %T, b %T\n", a, b)
	//	a main.MyInt, b int
}
