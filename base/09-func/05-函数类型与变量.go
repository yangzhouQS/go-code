package main

import "fmt"

type funcAdd func(a, b int) int

func adds(a, b int) int {
	return a + b
}

// 函数作为参数
func cal(a, b int, fn funcAdd) int {
	return fn(a, b) // 返回值位int
}
func main() {
	var c funcAdd
	c = adds
	fmt.Println(c(4, 6))

	fmt.Println(cal(6, 6, adds))
}
