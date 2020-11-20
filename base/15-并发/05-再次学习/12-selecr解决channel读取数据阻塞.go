package main

import "fmt"

func main() {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i + 1
	}
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello , " + fmt.Sprintf("%d", i+1)
	}
	for {
		select {
		case v := <-intChan:
			fmt.Println("intChan = ",v)
		case s := <-stringChan:
			fmt.Println(s)
		default:
			fmt.Println("都取不到了")
			return
		}
	}
}
