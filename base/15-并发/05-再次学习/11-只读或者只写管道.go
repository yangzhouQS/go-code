package main

import "fmt"

func main() {
	// 1.默认情况下管道是双向的管道
	//var ch chan int // 可读可写的管道
	// 2.只写的chan
	var ch2 chan<- int
	ch2 = make(chan int, 5)
	ch2 <- 10
	ch2 <- 10
	ch2 <- 10
	fmt.Println(ch2)

	// 取数据
	//num := <-ch2 // invalid operation: <-ch2 (receive from send-only type chan<- int)
	//fmt.Println(num)

	// 3.只读的chan
	var ch3 <-chan int
	var num2 = <-ch3
	fmt.Println(num2)
}
