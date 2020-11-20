package main

import "fmt"

// 接口定义的抽象方法 ,不包含函数体
type operation interface {
	add()
	sub()
}
type sum struct {
	num1 float64
	num2 float64
}

// 接口的实现
func (s *sum) add() {
	s.num1 = 100
	fmt.Println(s.num1 + s.num2)
}
func (s sum) sub() {
	fmt.Println(s.num2 - s.num1)
}

func main() {
	s := sum{4, 5}
	// 调用继承的方法
	s.add()
	s.sub()
	fmt.Println(s) //{100 5}
}
