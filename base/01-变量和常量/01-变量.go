package main

import "fmt"

// 声明不同的变量
var name string = "李四"
var age int = 16

// 一次声明多个变量
var (
	name2 string // '' 默认之为空字符串
	age2  int    // int的默认值为0
)

func main() {
	// 变量的初始化
	name2 = "大牛"
	age2 = 15
	fmt.Println(name, name2, age, age2)
	fmt.Printf("name = %s, name2 = %s, age = %d , age2 = %d\n", name, name2, age, age2)

	//	短变量声明,类型自动推断
	a := 100
	b := 99

	c, _ := 111, 22 // _下划线代表匿名变量,不需要这个变量

	// 一次声明多个变量
	a1, a2 := 3, 4
	fmt.Println(a, b, c)
	fmt.Println(a1, a2)
	ret1, _ := fn()
	fmt.Println(ret1)
}
func fn() (int, int) {
	return 4, 5
}
