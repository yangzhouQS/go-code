package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func (s Student) showInfo() {
	fmt.Println(s.Name, s.Age)
}
func (s Student) getName() {
	fmt.Println(s.Name)
}
func testStruct(a interface{}) {
	//typeOf := reflect.TypeOf(a)
	valueOf := reflect.ValueOf(a)
	// 值对应的类型
	vKind := valueOf.Kind()
	if vKind != reflect.Struct {
		panic("error type")
	}
	// 结构体字段个数
	fieldNum := valueOf.NumField()
	fmt.Println("字段个数:", fieldNum) // 字段个数: 2

	// 结构体方法的个数
	methodsNum := valueOf.NumMethod()
	fmt.Println("方法个数:", methodsNum) // 方法个数: 0
}
func main() {
	s := Student{"Tom", 16}
	testStruct(s)
	s.getName()
	s.showInfo()
}
