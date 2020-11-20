package main

import "fmt"

var global_a string = "Global A string"
var a int = 100

func test03() {
	var a int = 100
	fmt.Println("a = ", a) //a =  100
	fmt.Println(global_a)  // Global A string
}
func test04() {
	fmt.Println("a before = ", a)
	var a int = 99
	fmt.Println("a after = ", a)
	/*
	a before =  100
	a after =  99
	*/
}
func main() {
	test03()
	test04()
	fmt.Println(global_a)
}
