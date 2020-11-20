package main

import (
	"fmt"
	"sync"
)

func main02() {
	go func() {
		fmt.Println("hello")
	}()
	// 继续执行自己的逻辑

	// 匿名函数赋值给变量
	sayHello := func() {
		fmt.Println("say hello")
	}
	go sayHello()
}
func main01() {
	go sayHello()
	// 继续执行自己的逻辑
	fmt.Println("main ...")
}

var wg sync.WaitGroup

func main() {
	// 这个例子将决定main goroutine ，直到goroutine 托管sayHello 函数为止。
	wg.Add(1)
	go sayHello()
	wg.Wait()
	fmt.Println("run end")
}
func sayHello() {
	defer wg.Done()
	fmt.Println("say hello word")
}
