package main

import (
	"fmt"
	"time"
)

func main() {
	var data int
	go func() {
		data++
		fmt.Println("go data = ", data)
	}()
	time.Sleep(time.Second)
	if data == 0 {
		fmt.Println("data === 0")
	} else {
		fmt.Printf("data = %d \n", data)
	}
}
