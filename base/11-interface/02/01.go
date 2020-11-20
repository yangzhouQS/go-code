package main

import (
	"fmt"
)

type Operator interface {
	add()
	sub()
}

type sum struct {
	sum1 float64
	sum2 float64
}

func (s *sum) add() {
	fmt.Println(s.sum1 + s.sum2)
}
func (s *sum) sub() {
	fmt.Println(s.sum1 - s.sum2)
}

func main() {
	sum := &sum{3, 2}
	sum.add()
	sum.sub()
}
