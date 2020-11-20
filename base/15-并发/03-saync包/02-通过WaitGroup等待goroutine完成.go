package main

import (
	"fmt"
	"sync"
	"time"
)

func hello(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Printf("hello from %v!\n", id)
}
func main() {
	var wg sync.WaitGroup
	const numGreeters = 5
	wg.Add(numGreeters)
	for i := 0; i < numGreeters; i++ {
		go hello(&wg, i+1)
	}
	wg.Wait()
}
func main020() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("goroutine 1")
		time.Sleep(1)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("goroutine 2")
		time.Sleep(2)
	}()
	wg.Wait()
	fmt.Println("run end")
}
