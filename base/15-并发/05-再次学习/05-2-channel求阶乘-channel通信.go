package main

import (
	"fmt"
	"sync"
)

var (
	myMap = make(map[int]int, 20)
	//cacheMap = make(map[int]int)
	wg      sync.WaitGroup
	chanMap chan map[int]int
)

func jieCheng(num int) {
	defer wg.Done()
	n := 1
	for i := 1; i <= num; i++ {
		n *= i
	}
	myMap[num] = n // fatal error: concurrent map writes
}
func main() {
	for i := 1; i <= 20; i++ {
		wg.Add(1)
		go jieCheng(i)
	}
	wg.Wait()
	for i, v := range myMap {
		fmt.Printf("%d:  %d\n", i, v)
	}
}
