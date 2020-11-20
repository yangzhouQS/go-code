package main

import "fmt"

func main() {
	arr := [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("类型 %T\n", arr)

	// for rang遍历数组
	for k, v := range arr {
		fmt.Println(k, v)
	}

	// 遍历求数组的和
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		fmt.Println(i, "->", arr[i])
	}
	fmt.Printf("sum = %[1]d \n", sum) // sum = 28

	numbers1 := []int{1, 2, 3, 4, 5, 6}
	for i := range numbers1 {
		if i == 3 {
			numbers1[i] = i
		}
	}
	fmt.Println(numbers1)
}
