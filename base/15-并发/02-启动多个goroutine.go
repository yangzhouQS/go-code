package main

import (
	"fmt"
	"sync"
	"time"
)

func main01() {
	for i := 0; i < 10000; i++ {
		go func(i int) {
			fmt.Println("i = ", i)
		}(i)
	}
	fmt.Println("done main")
	time.Sleep(time.Millisecond)
}

var wg sync.WaitGroup

func hello03(i int) {
	wg.Done()
	fmt.Println("hello ->", i)
}
func main() {
	for i := 0; i < 100; i++ {
		go hello03(i)
		wg.Add(1) // 创建一个goroutine添加一个
	}
	fmt.Println("----main start")
	wg.Wait() // 登记的goroutine全部结束
	fmt.Println("main done")
}
