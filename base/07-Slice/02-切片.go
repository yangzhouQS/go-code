package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"李白", "白居易", "杜甫", "李商隐"}
	fmt.Printf("s的数据类型 = %T", s) // []string
	fmt.Println(s, len(s), cap(s))

	//	切片遍历
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
	for index, value := range s {
		fmt.Println(index, value)
	}

	s2 := []int{1, 2, 3}
	s2 = append(s2, 4, 5, 6)
	fmt.Println(s2)

	var a = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(a)

	b := [...]int{3, 7, 8, 9, 1}
	sort.Ints(b[:]) // [1 3 7 8 9]
	fmt.Println(b)
}
