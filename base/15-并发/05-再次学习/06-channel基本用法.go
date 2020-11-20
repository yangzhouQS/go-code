package main

import "fmt"

func main() {
	// 1.定义chan类型
	var ch chan int
	var ch2 = make(chan int, 10)
	// 2.初始化管道
	ch = make(chan int, 10)

	// 3.数据的写入
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i * i
		}
		close(ch)
	}()

	// 4.循环读取ch写入ch2
	func() {
		for {
			value, ok := <-ch
			if !ok {
				break
			}
			ch2 <- value
		}
		// 5.关闭管道
		close(ch2)
	}()

	// 获取管道的长度和容量
	fmt.Println("长度和容量: ",len(ch), cap(ch))
	for v := range ch2 {
		fmt.Println(v)
	}
}
