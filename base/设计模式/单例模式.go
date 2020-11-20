package main

import (
	"fmt"
	"sync"
)

/*
单例模式
分布式执行需要加锁,但是性能问题不是很好的{var wg sync.Mutex}
GO提供了Once
*/
type WebConfig struct {
	Port int
}

/*var webConfig *WebConfig
var wg sync.Mutex

func GetWebConfig() *WebConfig {
	wg.Lock()
	defer wg.Unlock()
	if webConfig == nil {
		webConfig = &WebConfig{Port: 8080}
		return webConfig
	}
	return webConfig
}*/

var webConfig *WebConfig
var once sync.Once

func GetWebConfig() *WebConfig {
	once.Do(func() {
		webConfig = &WebConfig{Port: 8080}
	})
	return webConfig
}

func main() {
	webConfig := GetWebConfig()
	fmt.Println(webConfig)
}
