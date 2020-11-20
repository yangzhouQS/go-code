package main

import (
	"fmt"
	"sync"
)

var chs chan int // 创建
var wg sync.WaitGroup

func printChs(ch chan int) {
	defer wg.Done()
	for {
		x, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("current ch = ", x)
	}
	close(ch)
}
func f1(ch chan int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch <- i
	}
}

func main() {
	fmt.Println(chs)
	// 通道初始化
	chs = make(chan int, 100)
	fmt.Println(chs)
	fmt.Println(len(chs), cap(chs)) // 0 10

	wg.Add(2)
	go f1(chs)
	go printChs(chs)

	wg.Wait()
	fmt.Println("run end")
}
