package main

import "fmt"

func f1() int { // 5
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) { // 6
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) { // 5
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) { // 5
	defer func(x int) {
		x++
	}(x)
	return 5
}

func calc4(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
func main() {
	/*
	A 1 2
	*/
	x := 1
	y := 2
	defer calc4("AA", x, calc4("A", x, y))
	x = 10
	defer calc4("BB", x, calc4("B", x, y))
	y = 20

	//fmt.Println(f1())
	//fmt.Println(f2())
	//fmt.Println(f3())
	//fmt.Println(f4())
}
