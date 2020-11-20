package main

import (
	"02-sort/SelectSort"
	"fmt"
)

func main() {
	arr := []int{7, 1, 42, 8, 9, 10, 15}
	fmt.Println(arr)
	//fmt.Println(SelectSort.SelectSortMax(arr))
	sortArr := SelectSort.SelectSort(arr)
	fmt.Println(sortArr)
}
