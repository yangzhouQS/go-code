package main

import (
	"fmt"
)

//切片的创建和初始化
func main01() {
	//只指定长度，则默认容量和长度相等
	slice := make([]string, 5)
	//指定长度和容量，容量不能小于长度
	slice2 := make([]string, 3, 5)

	fmt.Println(len(slice), cap(slice))   //5 5
	fmt.Println(len(slice2), cap(slice2)) //3 5
}

//使用切片字面量
func main02() {
	// 长度和容量都是2
	slice := []string{"西安", "北京"}
	slice = append(slice, "西沙群岛", "可可西里")
	fmt.Println(slice)

	slice2 := make([]string, 3)
	slice2[0] = "a"
	slice2[1] = "b"
	slice2[2] = "c"

	// 长度所有越界
	//slice2[3]="d" // panic: runtime error: index out of range [3] with length 3
	fmt.Println(slice2)

	//使用索引声明切片
	//下面创建了一个长度为100的切片
	s2 := []int{99: 0}

	fmt.Println(len(s2), cap(s2)) //100 100
}

//nil和空切片
func main033() {
	var a []int
	b := make([]int, 0)
	if a == nil {
		fmt.Println("a is nil")
	} else {
		fmt.Println("a is not nil")
	}
	if b == nil {
		fmt.Println("b is nil")
	} else {
		fmt.Println("b is not nil")
	}
	/*a is nil
	b is not nil*/
}

// 声明时不做任何初始化就会创建一个nil切片
func main04() {
	var s []string

	// new 产生的是指针，需要用* 
	s2 := *new(*[]string)
	fmt.Println(s, s2) // [] <nil>

	//声明空切片

	//使用make
	s3 := make([]int, 0)
	//使用切片字面量
	s4 := []int{}
	fmt.Println(s3, s4 == nil) // [] false
}

// 切片遍历
func main044() {
	s := []int{1, 3, 5}

	for i := 0; i < len(s); i++ {
		fmt.Println(i, s[i])
	}

	for index, value := range s {
		fmt.Println(index, value)
	}
}

// 切片扩容
func main05() {
	var slice []int
	for i := 0; i < 20; i++ {
		slice = append(slice, i)
		fmt.Printf("s:%v,len:len:%d,cap:%d,&p:%p \n", slice, len(slice), cap(slice), &slice)
	}

	s := make([]string, 5)
	s = append(s, "a", "b", "c")
	fmt.Println(s)
}

// 切片复制
func main06() {
	a := []int{2, 5, 8, 9, 7}
	b := a
	fmt.Println(len(a), len(b), a, b) // 一抹一样的 //5 5 [2 5 8 9 7] [2 5 8 9 7]
	b[0] = 99
	fmt.Println(len(a), len(b), a, b) // 一抹一样的  //5 5 [99 5 8 9 7] [99 5 8 9 7]
	// 切片赋值是引用传递
}

func main() {
	a := []int{2, 5, 8, 9, 7}
	b := make([]int, 5)
	copy(b, a)
	fmt.Println(len(a), len(b), a, b) // 一抹一样的 //5 5 [2 5 8 9 7] [2 5 8 9 7]
	b[0] = 99
	fmt.Println(len(a), len(b), a, b) // 5 5 [2 5 8 9 7] [99 5 8 9 7] 不会修改源切片的
}

// 从切片删除元素
func main08() {
	// 从切片中删除元素
	a := []int{30, 31, 32, 33, 34, 35, 36, 37}
	fmt.Println(a[2:])
	fmt.Println(append(a[:2]))

	//删除索引为4
	a = append(a[:4], a[5:]...)
	fmt.Println(a) //[30 31 32 33 35 36 37]
}

func main09() {
	var a = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(a)
}
