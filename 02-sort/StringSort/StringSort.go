package main

import (
	"fmt"
	"strings"
)

// 字符串安装ASCII码排序,第一个系统比较后边的
func StringSort(arrString []string) []string {
	length := len(arrString)
	if length <= 1 {
		return arrString
	}
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			// 左边小于 右边
			if strings.Compare(arrString[j], arrString[i]) == -1 {
				arrString[j], arrString[i] = arrString[i], arrString[j]
			}
		}
	}
	return arrString
}
func main() {
	arrString := []string{"c", "z", "D", "d", "a"}
	fmt.Println(StringSort(arrString))
	fmt.Println("A", "a")
}
func main01() {
	/*
		首先比较第一个字母
			if a == b {
					return 0
			}
			if a < b {
				return -1
			}
			return +1

	*/
	fmt.Println(strings.Compare("a", "b")) // -1
	fmt.Println(strings.Compare("a", "a")) // 0
	fmt.Println(strings.Compare("c", "a")) // +1

	fmt.Println("a" < "a")
	fmt.Println("a1" < "a2")
	fmt.Println("a1a" < "a1c")

	a1 := "a"
	a2 := "a"

	fmt.Println(&a1, &a2) // 0xc0000881e0 0xc0000881f0
	fmt.Println(a2 == a1) // true
}
