package main

import (
	"fmt"
)

func main() {
	num := []int{1, 2, 3, 4, 5, 56, 6}
	for index, value := range num {
		fmt.Println(index, value)
	}

	mapList := make(map[string]string)
	mapList["a"] = "A"
	mapList["b"] = "B"
	mapList["c"] = "C"
	mapList["d"] = "D"
	fmt.Println(mapList)
	// 遍历map
	for k, v := range mapList {
		fmt.Println(k, v)
	}
}
