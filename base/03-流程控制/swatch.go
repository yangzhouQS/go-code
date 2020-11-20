package main

import "fmt"

func main() {
	switch n:=12;n {
	case 0,2,4,6,8,10,12,14:
		fmt.Println("偶数")
		fallthrough // 允许执行下边的case
	case 1,3,7,9,11,13,15:
		fmt.Println("奇数")
	default:
		fmt.Println("未知")
	}
}
