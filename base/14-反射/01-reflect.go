package main

import (
	"fmt"
	"reflect"
)

type person struct {
	name string `json:"name"`
	age  int8   `json:"age"`
}

func getTypeOf(t interface{}) {
	fmt.Println("type->", reflect.TypeOf(t))
}

//type name和type kind
func getTypeOf2(i interface{}) {
	t := reflect.TypeOf(i)
	fmt.Printf("name:%v, kind:%v\n", t.Name(), t.Kind())
}
func getTypeOf3(i interface{}) {
	//t := reflect.ValueOf(i)
	//fmt.Println(t, t.Kind())
	v := reflect.ValueOf(i)
	k := v.Kind()
	fmt.Println(k, v)
}

func reflectSetValue2(i interface{}) {
	v := reflect.ValueOf(i)
	k := v.Kind()
	fmt.Println(k == reflect.Int64) //true
	if k == reflect.Int64 {
		v.Elem().SetInt(999)
	}
}
func main() {
	fmt.Println(reflect.TypeOf(person{}))
	var a int = 100
	var p person = person{"tom", 16}
	getTypeOf(a)  //type-> int
	getTypeOf(p)  // type-> main.person
	getTypeOf2(p) //name:person, kind:struct
	getTypeOf2(a) //name:int, kind:int
	getTypeOf3(a) //100 int
	getTypeOf3(p) //{tom 16} struct
	var b float64 = 3.1415
	getTypeOf3(b)

	// 根据反射修改值
	var c int64 = 6
	reflectSetValue2(&c)
	fmt.Println("c = ", c)
}
