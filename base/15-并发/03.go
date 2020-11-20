package main

import (
	"fmt"
	"runtime"
	"time"
)

func hello2(i int) {
	//fmt.Println("i  = ", i)
}
func main() {
	for i := 0; i < 100; i++ {
		go hello2(i)
	}
	fmt.Println(runtime.NumCPU())
	fmt.Println("runtime.NumGoroutine", runtime.NumGoroutine())
	time.Sleep(time.Second)
}
