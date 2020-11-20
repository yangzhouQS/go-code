package main

import (
	"fmt"
	"time"
)

// go怨言的并发
func calc() {
	for i := 0; i < 10; i++ {
		fmt.Println("i = ", i)
		time.Sleep(time.Second)
	}
	fmt.Println("calc exit")
}
func test() (int, int, int) {
	return 1, 2, 3
}
func main() {
	fmt.Println("main start")
	//go calc()
	//time.Sleep(10 * time.Second)
	//fmt.Println("main end")

	a, b, c := test()
	fmt.Println(a, b, c)
}
