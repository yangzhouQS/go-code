package main

import (
	"fmt"
	"sync"
)

var counter int = 0
var lock sync.Mutex

func Count(lock *sync.Mutex) {
	lock.Lock()
	defer lock.Unlock()
	counter++
	fmt.Println("counter = ", counter)
}

func main() {
	lock := &sync.Mutex{}
	for i := 0; i < 5; i++ {
		go Count(lock)
	}
	fmt.Println("run end")
}
