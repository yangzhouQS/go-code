package main

import (
	"fmt"
)

func main() {
	var data int
	go func() {
		data++
	}()
	//time.Sleep(time.Millisecond)
	if data == 0 {
		fmt.Printf("value = %v .\n", data)
	} else {
		fmt.Println("NO")
	}
}
