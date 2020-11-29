package main

import "fmt"

func main() {
	var m1 map[string]int         // panic: assignment to entry in nil map
	fmt.Println(m1 == nil)        // 内存中没有分配内存空间
	m1 = make(map[string]int, 10) // 使用make分配内存空间初始化
	m1["a"] = 100
	m1["b"] = 99
	fmt.Println(m1)

	// 判断map中是否包含指定的key
	value, ok := m1["a"]
	fmt.Println(value, ok) // 100 true
	_, isB := m1["b"]
	if isB {
		fmt.Println("b 对应的key存在的")
	} else {
		fmt.Println("不存在的")
	}

	// 删除
	delete(m1, "a")
	fmt.Println(m1) //map[b:99]

	//	map的遍历
	for k, v := range m1 {
		fmt.Println(k, v)
	}
}
