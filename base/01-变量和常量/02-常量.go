package main

import "fmt"

const PI float64 = 3.141592653589793
const e float64 = 2.714
const (
	a = 100
	b  // 默认和上一行的值是相同的
	c
)

// `iota`是go语言的常量计数器，只能在常量的表达式中使用。
const (
	n1 = iota
	n2
	n3 = 100
	n4 = iota //3
	n5        //4
	n6        //5
)
//const (
//	a, b = iota + 1, iota + 2 //1,2
//	c, d                      //2,3
//	e, f                      //3,4
//)

func main() {
	fmt.Println(n1, n2, n3, n4, n5, n6)
	fmt.Println(PI, e)
	fmt.Println(a, b, c)
}
