package main

import "fmt"

func write(intChan chan int, n int) {
	for i := 0; i < n; i++ {
		intChan <- i + 1
		fmt.Println("write ", i)
	}
	close(intChan)
}
func read(intChan chan int, result chan int) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println("read ", v)
		result <- v
	}
	close(result)
}
func main() {
	var ch = make(chan int, 50)
	var ch2 = make(chan int, 50)
	go write(ch, 50)
	go read(ch, ch2)
	for value := range ch2 {
		fmt.Println(value)
	}
}
