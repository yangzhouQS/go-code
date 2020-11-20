package main

import (
	"fmt"
	"time"
)

func result(c chan int) {
	ret := <-c
	fmt.Println("接受完毕", ret)
}
func main() {
	// 创建
	//var ch chan int
	// 声明的通道后需要使用`make`函数初始化之后才能使用
	ch := make(chan int, 4) // 创建缓冲区为1的通道
	//go result(ch)
	// 发送
	ch <- 100
	ch <- 200
	ch <- 300
	ch <- 400
	go func() {
		for {
			value, ok := <-ch
			if !ok {
				break
			}
			fmt.Println("接收到的数据", value)
		}
		close(ch)
	}()
	ch <- 99
	fmt.Println("发送完毕")
	time.Sleep(time.Millisecond)
	//// 关闭
	//close(ch)
}
