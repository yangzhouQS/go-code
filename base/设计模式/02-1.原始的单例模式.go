package main

import (
	"fmt"
	"sync"
)

var instance *Instance
var lock sync.Mutex

type Instance struct {
	name string
}

// 双重检查
func getInstance() *Instance {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		instance = &Instance{name: "sam"}
	}
	return instance
}
func main() {
	instance := getInstance()
	instance2 := getInstance()
	fmt.Println(instance == instance2) // true
}
