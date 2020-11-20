package main

import (
	"fmt"
	"time"
)

func hello02() {
	fmt.Println("hello gouroutine")
}
func main() {
	go hello02()
	fmt.Println("main func")

	// 先执行main的goroutine, 结束后全部的goroutine全部结束
	// main执行结束等待一会儿
	time.Sleep(time.Millisecond)
}
