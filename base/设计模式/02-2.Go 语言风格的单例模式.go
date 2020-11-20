package main

import (
	"fmt"
	"sync"
)

type Instance struct {
	name string
}

var (
	instance *Instance
	once     sync.Once
)

func GetInstance() *Instance {
	once.Do(func() {
		instance = &Instance{name: "sam"}
	})
	return instance
}
func main() {
	instance := GetInstance()
	instance2 := GetInstance()
	fmt.Println(instance == instance2) // true
}
