package main

import "fmt"

func main() {
	// 创建map
	scoreMap := make(map[string]int)
	scoreMap["数学"] = 99
	scoreMap["语文"] = 99
	scoreMap["英语"] = 99
	scoreMap["计算机"] = 99

	fmt.Println(scoreMap)

	// 创建map时直接初始化
	map2 := map[string]int{"a": 100}
	fmt.Println(map2)

	// 判断某个键是否存在
	v, ok := map2["b"]
	fmt.Println(v, ok)
	if value, ok := map2["a"]; ok {
		fmt.Println(value)
	}

	// map添加新元素
	map2["b"] = 99
	fmt.Println(map2) // map[a:100 b:99]

	// 修改
	map2["a"] = 999
	fmt.Println(map2) // map[a:999 b:99]

	// 删除
	delete(map2, "b")
	fmt.Println(map2) // map[a:999]
	map2["b"] = 999
	map2["c"] = 32
	map2["d"] = 23

	// 遍历
	for v, k := range map2 {
		fmt.Println(v, k)
	}
}
