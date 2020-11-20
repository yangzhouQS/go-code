package main

import (
	"fmt"
	"sync"
)

type Single struct {
	data int
}

var singleton *Single
var once sync.Once

func getInterFace() *Single {
	once.Do(func() { // 智能信号,时时刻刻只可以运行一个
		singleton = &Single{
			data: 10,
		}
	})
	return singleton
}
func main() {
	face := getInterFace()
	fmt.Println(face)
}
