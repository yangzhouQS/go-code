package main

import "fmt"

type Animal struct {
	name string
}

func (a *Animal) run() {
	fmt.Printf("%v,在奔跑 \n", a.name)
}

type Dog struct {
	color string
	*Animal
}

func (d *Dog) draw() {
	fmt.Printf("%v的%v 在画画", d.color, d.name)
}

func main() {
	d := Dog{color: "黄色", Animal: &Animal{name: "阿黄"}}
	d.run()
	d.draw()
}
