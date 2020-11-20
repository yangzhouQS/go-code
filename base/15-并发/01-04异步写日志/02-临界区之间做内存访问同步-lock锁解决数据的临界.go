package main

import (
	"fmt"
	"sync"
)

func main() {
	var flag sync.Mutex
	var value int
	go func() {
		flag.Lock() // 在内存中独立操作变量
		value++
		// 我们的goroutine 应该独占 该内存的访问权。
		flag.Unlock()
	}()
	flag.Lock()
	if value == 0 {
		fmt.Println("value == 0")
	} else {
		fmt.Println("value = ,", value)
	}
	flag.Unlock()
}
