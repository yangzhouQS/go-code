package main

import "fmt"

func test() string {
	return "hello word"
}
func main01() {
	s := test()
	fmt.Println(s)
}

//类型简写
func sum(a, b int) (plus int) {
	plus = a + b
	return
}

//可变参数
//可变参数的类型是对应的切片类型
func plus(a int, args ...int) (sum int) {
	fmt.Printf("%T \n", args) // []int
	sum = a
	if len(args) > 0 {
		for _, v := range args {
			sum = sum + v
		}
	}
	return
}
func main() {
	plus(1, 2, 3, 4, 5)
}
func main02() {
	fmt.Println(sum(1, 5))
	fmt.Println(plus(5, 6, 8, 10, 2))
}

//返回值
//多返回值
func test2() (string, int) {
	return "hello word", 100
}

//返回值命名
//函数定义时可以给返回值命名，并在函数体中直接使用这些变量，最后通过return关键字返回。
func calc() (a int, sum int) {
	sum = 100
	a = 99
	return
}
func main111() {
	//多返回值
	a, b := test2()
	fmt.Println(a, b)   // hello word 100
	fmt.Println(calc()) // 99 100
}
