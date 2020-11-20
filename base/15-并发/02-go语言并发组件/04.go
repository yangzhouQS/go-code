package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched() // 让CPU 把时间片让给别人，下次某个时候继续恢复执行该goroutine
		fmt.Println(i, s)
	}
}
func main() {
	go say("hello")
	say("word")
}
