package main

import "fmt"

func main() {
	var age = 100                  // 变量
	pAge := &age                   // age的地址
	aAgeValue := *pAge             // 地址值
	fmt.Printf("%T \n", age)       // int
	fmt.Printf("%T \n", pAge)      // *int
	fmt.Printf("%T \n", aAgeValue) // int
	// 修改前
	fmt.Println(age, pAge, *pAge)
	//修改地址对应的值
	*pAge = 99
	// 修改后
	fmt.Println(age, pAge, *pAge)

}
