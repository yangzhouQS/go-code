package main

import "fmt"

func main() {
	//	1.&根据值取地址
	// 2.*根据地址取值
	n := 100
	fmt.Println(&n)
	fmt.Printf("%T\n", &n) // *int int类型的指针变量
	p := &n
	// 对地址取值
	fmt.Println(*p)

	//var a *int // int类型的指针变量
	// 指针的初始化位nil
	//*a = 100   // panic: runtime error: invalid memory address or nil pointer dereference
	//fmt.Println(a)

	a := new(int)
	*a = 100
	fmt.Println("a = ",a)
}
