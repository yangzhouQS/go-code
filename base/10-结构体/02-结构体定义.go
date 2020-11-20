package main

import "fmt"

// 结构体定义
type Person struct {
	name string
	a, b int
	age  uint8
}

func main() {
	// 结构体实例化
	p := Person{"李四", 1, 2, 16}
	fmt.Println(p)
	var p2 Person
	p2.name = "TOM"

	p2.a = 1
	p2.b = 2
	p2.age = 18
	fmt.Println(p2)
	// 按照字段名称初始化
	p3 := Person{a: 2, b: 3, age: 19, name: "大牛"}
	fmt.Println(p3)

	var p4 Person = Person{name: "xiaoli", a: 100, b: 66, age: 68}
	fmt.Println(p4)
	fmt.Printf("p3 %+v\n",p3) // p3 {name:大牛 a:2 b:3 age:19}
	fmt.Printf("p4 %+v\n",p4) // p4 {name:xiaoli a:100 b:66 age:68}
}
