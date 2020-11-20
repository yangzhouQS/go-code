package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("pi = %.8f\n", math.Pi)
	a := float64(0.002222)

	// %T打印变量的类型
	fmt.Printf("a = %T\n", a) // flot64
	// %v打印变量的值
	fmt.Printf("a = %v\n", a)
}
