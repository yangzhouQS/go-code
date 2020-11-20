package main

import (
	"fmt"
	"sort"
)

func main() {
	sortInt := []int{-1, -2, 100, 20}
	sort.Ints(sortInt)
	fmt.Println(sortInt)
}
