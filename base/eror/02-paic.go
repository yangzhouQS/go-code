package main

import "fmt"

func test(a, b int) int {
	if b == 0 {
		panic("除数不可以为0")
		return 0
	}
	return a / b
}

func main() {
	fmt.Println(test(10, 0))
}
