package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func printHello2(i int) {
	defer wg.Done()
	//fmt.Println("print hello!", i)
}
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	startTime := time.Now()
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go printHello2(i)
	}
	// 主线程等待所有登记的goroutine都结束
	defer wg.Wait()
	// 计算程序的运行时间
	endTime := time.Since(startTime)
	fmt.Println(endTime)
}
