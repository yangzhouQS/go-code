package main

import (
	"fmt"
)

type Person struct {
	id   int
	name string
	age  int
}

func (p Person) speak() {
	fmt.Printf("id = %d, name = %s, age = %d \n", p.id, p.name, p.age)
}

func (p Person) calc() {
	sum := 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Printf("%s 0...100的和为 %d \n", p.name, sum)
}

// 传入参数
func (p Person) calc2(n int) {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	fmt.Printf("%s 0...%d的和为 %d \n", p.name, n, sum)
}

// 传入参数 有返回值
func (p Person) calc3(n int) int {
	sum := 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Printf("%s 0...%d的和为 %d \n", p.name, n, sum)
	return sum
}
func (p Person) getSum(a, b int) (sum int) {
	sum = 0
	sum = a + b
	return
}
func main() {
	p := Person{1, "sam", 16}
	var p2 Person
	p2 = Person{}
	p3 := Person{}
	p3.name = "xxx"
	p4 := Person{4, "sam", 16}
	fmt.Println(p, p2, p3, p4)
	p.speak()
	p.calc()
	p.calc2(100)
	fmt.Println(p.calc3(100))
	fmt.Println(p.getSum(10, 20))

}
