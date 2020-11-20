package main

import (
	"fmt"
	"reflect"
)

func reflectTest(arg interface{}) {
	// 1.获取type类型
	typeOf := reflect.TypeOf(arg) // int
	fmt.Println(typeOf)

	// 2.获取value类型
	valueOf := reflect.ValueOf(arg)
	fmt.Println(valueOf) // 12
	fmt.Println(valueOf) // 12

	// valueOf转换为接口类型
	iV := valueOf.Interface()
	fmt.Printf("类型%T ,值为: %v \n", iV, iV)
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func reflectTestStruct(p Person) {
	typeOf := reflect.TypeOf(p)
	// 类型和种类 ==>底层的类型
	fmt.Println(typeOf.Name(), typeOf.Kind()) // Person struct
	valueOf := reflect.ValueOf(p)
	// 转换为接口类型
	newValue := valueOf.Interface()
	//类型main.Person,值:{Tom 16}
	fmt.Printf("类型%T,值:%v\n", newValue, newValue)
	// 断言为Person
	p2, ok := newValue.(Person)
	if ok {
		fmt.Println(p2.Name, p2.Age) // Tom 16
	}
}
func main() {
	//reflectTest(12)

	// 结构体反射
	p := Person{"Tom", 16}
	fmt.Println(p)
	reflectTestStruct(p)

	fmt.Println("-------")
	fmt.Println(reflect.Int8)
	fmt.Println(reflect.String)
}
