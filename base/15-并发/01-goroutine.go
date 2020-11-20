package main

import (
	"fmt"
	"sync"
	"time"
)

//启动多个goroutine
var wg sync.WaitGroup

func test() {
	fmt.Println("go run goroutine")
}
func hello(i int) {
	wg.Done() // goroutine结束 -1
	fmt.Println("hello - ", i)
}

// main也是一个goroutine
func main() {
	// 开启goroutine执行函数
	go test()
	time.Sleep(time.Millisecond)
	fmt.Println("done")
	for i := 0; i > 10000; i++ {
		wg.Add(1) // 启动一个goroutine就+1
		go hello(i)
	}
	wg.Wait() // 等待所有的goroutine结束
}
