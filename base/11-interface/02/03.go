package main

import (
	"fmt"
)

type AInterfaceA interface {
	test01()
	test02()
}
type AInterfaceB interface {
	test01()
	test03()
}
type AInterfaceC interface {
	AInterfaceA
	AInterfaceB
}

type Stu struct {
}

func (s *Stu) test01() {
	fmt.Println("test1")
}

func (s *Stu) test02() {
	fmt.Println("test2")
}

func (s *Stu) test03() {
	fmt.Println("test3")
}
func main() {
	stu := Stu{}
	var a AInterfaceA = &stu
	var b AInterfaceB = &stu
	var c AInterfaceC = &stu

	fmt.Println(a, b, c)
	c.test01()
	c.test02()
	c.test03()
}
func main01() {
	stu := Stu{}
	var a AInterfaceA = &stu
	var b AInterfaceB = &stu
	fmt.Println(a, b)
	a.test01()
	a.test02()
	b.test01()
	b.test03()
}
