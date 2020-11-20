package main

import "fmt"

func A() {
	fmt.Println("A")
}
func B(x int) {
	defer func() {
		// recover不会让错误阻止程序崩溃 , 后面的代码继续运行
		err := recover() //出现错误的时候才会处理
		if err != nil { //判断进行提示
			fmt.Println(err, "数组索引可能越界")
		}
	}()
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr[x], "B")
}
func C() {
	fmt.Println("c")
}
func main() {
	A()
	B(10)
	C()
}
