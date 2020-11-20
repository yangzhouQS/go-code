package main

import "fmt"

type MyInt int

func main() {
	// 声明int32的变量x值位10

	// 1:
	//var x int32
	//x = 10

	// 2:
	//var x int32 = 10

	// 3:
	var x = int32(10)
	fmt.Println(x)

	// 1:
	//var y MyInt
	//y = 100

	// 2:
	//var y MyInt = 100

	// 3:
	//var y = MyInt(19)

	// 4:强制类型转化
	y := MyInt(99)

	fmt.Println(y)

}
