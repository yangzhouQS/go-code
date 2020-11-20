package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var test1 = func(v1, v2 int) int {
	return v1 + v2
}
var tesst2 = func(v1, v2 int, v3 string) string {
	return fmt.Sprintf("%s - %s - %s", strconv.Itoa(v1), strconv.Itoa(v2), v3)
}
var bridge = func(funcName interface{}, args ...interface{}) {
	typeOf := reflect.TypeOf(funcName)
	fmt.Println(typeOf)
	fmt.Println(typeOf.Name())
	fmt.Println(typeOf.Kind())
	fmt.Println(typeOf.Align())
	fmt.Println("valueOf")
	valueOf := reflect.ValueOf(funcName)
	fmt.Println(valueOf)
	newTest1 := valueOf.MethodByName("test1")
	newTest1.Call(newTest1)
}

func main() {
	bridge("test1")
}
