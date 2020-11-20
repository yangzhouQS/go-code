package main

import "fmt"

func main() {
	fmt.Println("开始")
	defer fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")
	fmt.Println("结束")
	/*
	开始
	结束
	c
	b
	a
	*/
}
