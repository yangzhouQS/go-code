package main

import "fmt"

func funcA() {
	fmt.Println("func A")
}

func funcB() {
	defer func() {
		fmt.Println("释放数据库连接")

		error := recover()
		if error != nil {
			fmt.Println(error, "--")
			//	后面的代码回继续执行的
		}
	}()
	panic("panic in B")
}

func funcC() {
	fmt.Println("func C")
}
func mainx() {
	funcA()
	funcB()
	funcC()
	//var s string
	//fmt.Scanln(&s)
	//fmt.Println(s)

	var (
		name    string
		age     int
		married bool
	)
	fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
}

func Increase() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

func main() {
	in := Increase()
	fmt.Println(in())
	fmt.Println(in())
}
/*
输出：
1
2
*/
