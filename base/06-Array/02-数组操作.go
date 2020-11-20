package main

import "fmt"

func main() {
	// 数组初始化
	var arr = [3]string{"a", "b", "c"}
	var arr2 = [...]string{"c", "d"}
	arr3 := [...]int{1, 2, 3, 4, 5}

	fmt.Println(arr, arr2, arr3)

	// 数组遍历
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	// for rang 遍历
	for _, v := range arr {
		fmt.Println(v)
	}

	// 数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值。
	testArr(arr3)
	fmt.Println("arr = ", arr3) // arr =  [1 2 3 4 5]
}

func testArr(arr [5]int) {
	arr[0] = 99
	fmt.Println(arr) // [99 2 3 4 5]
}
