package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

func main() {
	var x interface{}
	var a = 10
	x = a
	fmt.Println(x, a)
	fmt.Println("--------")

	// 待检测的断言
	value, ok := x.(float64)
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("transfer fail")
	}
	fmt.Printf("%T, v = %v", x, x) // int, v = 10
}
func main01() {
	var a interface{}
	var point = Point{1, 2}
	a = point
	fmt.Println(point, a)
	var b Point
	// b = a // cannot use a (type interface {}) as type Point in assignment: need type assertion
	b = a.(Point) // {1 2} // 类型断言,将a转换为Point

	fmt.Println(b)
}
