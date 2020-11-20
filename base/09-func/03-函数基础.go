package main

import "fmt"

func add(a, b int) int {
	return a + b
}
func add2(a, b int) (c int) {
	c = a + b
	return
}
func sum2(arg ...int) int {
	fmt.Printf("%T\n", arg) //[]int slice
	sum := 0
	for _, v := range arg {
		sum += v
	}
	return sum
}
func main() {
	fmt.Println(add(4, 5))
	fmt.Println(add2(4, 5))
	fmt.Println(sum2(1, 2, 3))
}
