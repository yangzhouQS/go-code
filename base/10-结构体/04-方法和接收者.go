package main

import "fmt"

type person struct {
	name string
	age  int8
}

func newPerson(name string, age int8) *person {
	return &person{
		name: name,
		age:  age,
	}
}

// 值类型的接收者
//当方法作用于值类型接收者时，Go语言会在代码运行时将接收者的值复制一份。
func (p person) SetAge2(newAge int8) {
	p.age = newAge
	fmt.Println("副本", p)
}

// 指针类型的接收者
func (p *person) setAge(newAge int8) {
	fmt.Println("--", p)
	p.age = newAge
}
func main() {
	p := person{"tom", 18}
	p.setAge(16)
	p.setAge(26)
	p.SetAge2(100)
	fmt.Println(p)
}
