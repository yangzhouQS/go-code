package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <-chan interface{}
	var wg sync.WaitGroup
	noop := func() {
		wg.Done()
		<-c // 永远不会退出的goroutine ，这样就可以在内存中保留一 段时间用于测算
	}

	const numGoroutines = 1e6 // 定义了要创建的goroutine 的数量
	wg.Add(numGoroutines)
	before := memConsumed() // 测算在创建goroutine 之前消耗的内存总量。
	for i := numGoroutines; i > 0; i-- {
		go noop()
	}
	wg.Wait()

	after := memConsumed() // 测算在创建goroutine 之后消耗的内存总量。
	fmt.Printf("%.3fkb", float64(after-before)/numGoroutines/1000)
}
