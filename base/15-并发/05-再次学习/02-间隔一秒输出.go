package main

import (
	"fmt"
	"strconv"
	"time"
)

func printHello() {
	for i := 0; i < 10; i++ {
		fmt.Printf("hello i = %s \n", strconv.Itoa(i))
		time.Sleep(time.Second * 1)
	}
}
func main() {
	go printHello()
	for i := 0; i < 10; i++ {
		fmt.Println("main hello word", i, time.Now())
		time.Sleep(time.Second)
	}
}
