package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

var ch = make(chan interface{}, 10)

func main() {
	fmt.Println(ch)

	ch <- 100
	ch <- "hello"
	ch <- "word"
	ch <- Cat{Name: "Tom", Age: 3}

	// 取值
	<-ch
	<-ch
	<-ch
	v := <-ch
	fmt.Println(v)
	fmt.Printf("%T \n", v) // main.Cat
	//fmt.Printf("%T,v.Name: %s\n", v, v.name) // v.name undefined (type interface {} is interface with no methods)
	// 类型断言
	newCat := v.(Cat)
	fmt.Printf("Name = %s, Age = %d \n", newCat.Name, newCat.Age)
}
