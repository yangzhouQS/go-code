package main

import (
	"fmt"
)

func main01() {
	switch n := 12; n {
	case 0, 2, 4, 6, 8, 10, 12, 14:
		fmt.Println("偶数")
		fallthrough // 允许执行下边的case
	case 1, 3, 7, 9, 11, 13, 15:
		fmt.Println("奇数")
	default:
		fmt.Println("未知")
	}
}

func main() {
	ans := -10
	switch ans {
	case 1:
	case -1:
		fmt.Println("1 .one")
	case 2:
		fmt.Println("2 .one")
	case 3:
		fmt.Println("3 .one")
	default:
		fmt.Println("not number")
	}
}
