package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	for _, salutation := range []string{"hello", "word", "good day"} {
		wg.Add(1)
		go func(salutation string) {
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation)
	}
	wg.Wait()
}
func main0001() {
	for _, salutation := range []string{"hello", "word", "good day"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(salutation)
		}()
	}
	wg.Wait()
}
func main001() {
	salutation := "hello"
	wg.Add(1)
	go func() {
		defer wg.Done()
		salutation = "welcome"
	}()
	fmt.Println("salutation", salutation) // hello
	wg.Wait()
	fmt.Println("salutation", salutation) // welcome
}
