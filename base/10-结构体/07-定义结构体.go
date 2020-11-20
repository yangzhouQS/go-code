package main

import "fmt"

// 定义结构体
type UserInfo struct {
	Name   string
	Age    int
	Gender string
}

// 基本类型字段
type student3 struct {
	id    int
	name  string
	age   int
	class string
}

// 预留字段
type people struct {
	name string
	age  int
	_    string // 预留字段
}

// 结构体作为字段
type date struct {
	year  int
	month int
	day   int
}
type student2 struct {
	name     string
	age      int
	birthday date
}

type student struct {
	id    int // int类型的默认值为 0
	name  string
	age   int
	class string
}
type rectangle struct {
	width  int
	height int
}

func (rec rectangle) area() int {
	return rec.width * rec.height
}

// 面向对象的结构体
func main() {
	rec := rectangle{10, 20}
	fmt.Println(rec.area())

}

// 匿名字段
func main6() {
	type people struct {
		name string
		age  int
	}
	type teacher struct {
		people     // 匿名字段
		department string
	}
	// 匿名字段初始化
	teach := teacher{
		people{name: "华罗庚", age: 86},
		"教研部",
	}
	// 匿名字段的访问
	fmt.Println(teach)
	fmt.Println(teach.people)
	fmt.Println(teach.department)

}

// 嵌入式结构体
func main5() {
	type Person struct {
		name     string
		birthday struct {
			year  int
			month int
			date  int
		}
	}
	fmt.Println(Person{}) //{ {0 0 0}}

	// 结构体变量
	var user = struct {
		name string
		age  int
	}{"sam", 26}
	fmt.Println(user) // {sam 26}
}

// 结构体对象初始化
func main4() {
	// 1.字段全部初始化
	stu := student{101, "sam", 22, "计算机科学"}
	stu2 := &student{102, "sam", 22, "s"}
	fmt.Println(stu)
	fmt.Println(stu2)

	// 2.部分初始化
	stu3 := student{name: "小白"}
	stu4 := &student{name: "xx"}
	fmt.Println(stu3)
	fmt.Println(stu4)
}

// 结构体对象
func main3() {
	var stu *student
	// 等价于
	stu2 := new(student)
	fmt.Println(stu, stu2) // <nil> &{0  0 }
}

func main01() {
	// 结构体变量
	var stu student                             // 变量名称 变量的类型为结构体
	stu2 := student{01, "sam", 22, "桥梁工程2022级"} // 类型推断

	// 字段访问
	stu.name = "李四"
	stu.age = 26
	stu.class = "计算机01班"
	fmt.Println(stu)
	fmt.Println(stu2)
}
