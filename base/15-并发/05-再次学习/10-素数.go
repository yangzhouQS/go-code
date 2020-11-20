package main

import (
	"fmt"
	"time"
)

// intChan存放1-8000
func putNum(intChan chan int) {
	for i := 0; i < 80000; i++ {
		intChan <- i
	}
	close(intChan)
}
func printNum(intChan, primeChan chan int, exitChan chan bool) {
	for {
		value, ok := <-intChan
		if !ok {
			break
		}
		flag := true
		for i := 2; i <= value; i++ {
			if i%value == 0 { // 改value不是素数
				flag = false
				break
			}
		}
		if !flag {
			primeChan <- value
		}
	}
	fmt.Println("取不到数据退出!!!")
	exitChan <- true
}
func main() {
	startTime := time.Now()
	ch1 := make(chan int, 1000)
	primeChan := make(chan int, 2000)
	exitChan := make(chan bool, 4)
	go putNum(ch1)
	for i := 0; i < 4; i++ {
		go printNum(ch1, primeChan, exitChan)
	}

	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		// 4个结果取完就退出
		close(primeChan)
	}()

	// 取出primeChan
	for {
		_, ok := <-primeChan
		if !ok {
			break
		}
		//fmt.Println("素数 ,", v)
	}
	fmt.Println("main线程结束退出")
	fmt.Println("程序运行时间: ", time.Since(startTime))
}
