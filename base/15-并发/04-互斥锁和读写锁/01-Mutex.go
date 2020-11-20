package main

import (
	"fmt"
	"sync"
)

var count int
var lock sync.Mutex

func increment() {
	lock.Lock()
	defer lock.Unlock()
	count++
	fmt.Println("1  count = ", count)
}
func decrement() {
	lock.Lock()
	defer lock.Unlock()
	count--
	fmt.Println("2  count = ", count)
}
func main() {
	// 增量
	var arithmetic sync.WaitGroup
	for i:=0;i<5;i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increment()
		}()
	}
	
	// 减量
	for i:=0;i<5;i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decrement()
		}()
	}

	arithmetic.Wait()
	fmt.Println("run end")
}
