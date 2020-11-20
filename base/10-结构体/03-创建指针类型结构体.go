package main

import "fmt"

type person struct {
	name string
	age  int8
}

func f(p *person) {
	p.age = 100
}
func main() {
	// 创建值类型的结构体
	p := new(person)
	fmt.Printf("p = %T \n", p) // p = *main.person
	p.name = "小明"
	p.age = 19
	f(p)
	fmt.Printf("p = %v \n", p)

	p2 := person{"tom", 9}
	f(&p2)
	fmt.Println(p2)

	p3 := &person{}
	fmt.Printf("p3 = %T,%v", p3, p3) // *main.person

}
