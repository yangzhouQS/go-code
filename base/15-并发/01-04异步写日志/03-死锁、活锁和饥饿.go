package main

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	mu    sync.Mutex
	value int
}

var wg sync.WaitGroup

func main() {
	printSum := func(v1, v2 *value) {
		defer wg.Done()
		v1.mu.Lock()
		defer v1.mu.Unlock()


		time.Sleep(time.Second * 2)
		v1.mu.Lock()
		defer v1.mu.Unlock()
		fmt.Println("v1 = ", v1, "v2 = ", v2)
	}
	var a, b value
	wg.Add(2)
	go printSum(&a, &b) // 锁定a
	go printSum(&b, &a) // 锁定b
	wg.Wait()
}
