package main

import "fmt"

func main() {
	slice := make([]int, 0, 10)
	slice = append(slice, 1)
	slice = append(slice, 2)
	slice = append(slice, 3)
	slice = append(slice, 4)
	slice = append(slice, 5)
	slice = append(slice, 6)

	fmt.Println(slice[2:]) // 跳过前2个
	fmt.Println(slice[3:]) // 跳过前2个
	fmt.Println(slice[:2]) // 只去前2个

	// 删除索引为3位置的元素
	slice = append(slice[:2], slice[3:]...)
	fmt.Println(slice)
}
