package main

import (
	"fmt"
	"strconv"
	"sync"
)

type Person struct {
	Name, Address string
	Age           int
}

var (
	ch = make(chan *Person, 11)
	wg sync.WaitGroup
)

func rand() {
	for i := 0; i < 10; i++ {
		p := &Person{Name: "tom" + strconv.Itoa(i), Age: i + 10, Address: "西安"}
		ch <- p
	}
	close(ch)
}

func main() {
	go rand()
	for person := range ch {
		fmt.Println(person)
	}
}
