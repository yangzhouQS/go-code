package main

import (
	"01-ArrayList/ArrayList"
	"01-ArrayList/Queue"
	"01-ArrayList/StackArray"
	"fmt"
)

func main01() {
	fmt.Println("ArrayList")

	list := ArrayList.NewArrayList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	fmt.Println(list)
}
func main02() {
	fmt.Println("ArrayList")

	list := ArrayList.NewArrayList()
	list.Append("a")
	list.Append("b")
	list.Append("c")
	fmt.Println(list, list.Size())
	fmt.Println(list.String())
}

func main03() {
	fmt.Println("ArrayList")
	// 必须实现ArrayList的所有方法
	var list ArrayList.List = ArrayList.NewArrayList()
	list.Append("a")
	list.Append("b")
	list.Append("c")
	fmt.Println(list, list.Size())
	//list.Delete(1)
	list.Insert(1, "hello")
	fmt.Println(list, list.Size())
	fmt.Println(list.String())
}

func main04() {
	fmt.Println("ArrayList")
	// 必须实现ArrayList的所有方法
	var list ArrayList.List = ArrayList.NewArrayList()
	list.Append("a")
	list.Append("b")
	list.Append("c")
	fmt.Println(list, list.Size())
	for i := 1; i < 5; i++ {
		list.Insert(1, fmt.Sprint("hello", i))
		fmt.Println(list, list.Size())
	}
	fmt.Println(list, list.Size())
}

func main054() {
	fmt.Println("ArrayList")
	// 必须实现ArrayList的所有方法
	var list ArrayList.List = ArrayList.NewArrayList()
	list.Append("a")
	list.Append("b")
	list.Append("c")
	fmt.Println(list, list.Size())
	for i := 1; i < 10; i++ {
		list.Insert(1, fmt.Sprint("hello", i))
		fmt.Println(list, list.Size())
	}
	fmt.Println(list, list.Size())
}

func main05() {
	fmt.Println("ArrayList")
	// 必须实现ArrayList的所有方法
	var list ArrayList.List = ArrayList.NewArrayList()
	list.Append("a")
	list.Append("b")
	list.Append("c")
	fmt.Println(list, list.Size())
	for i := 1; i < 10; i++ {
		list.Insert(1, fmt.Sprint("hello", i))
		fmt.Println(list, list.Size())
	}

	for it := list.Iterator(); it.HasNext(); {
		item, err := it.Next()
		if err != nil {
			panic(err)
		}
		fmt.Println(item)
	}
	fmt.Println(list, list.Size())
}

func main06() {
	stack := StackArray.NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	fmt.Println(stack.Size())
	fmt.Println(stack.Pop())
	fmt.Println(stack)
}

// 数组栈
func main07() {
	stack := ArrayList.NewArrayListStackX()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	for it := stack.Iterator; it.HasNext(); {
		item, _ := it.Next()
		fmt.Println(item)
	}
	fmt.Println(stack.Size())
	fmt.Println(stack.Pop())
	fmt.Println(stack)
}

func add(num int) int {
	sum := 0
	if num == 0 {
		return 0
	} else {
		sum = num + add(num-1)
	}
	return sum
}

// 栈实现递归
func main08() {
	stack := StackArray.NewStack()
	stack.Push(5)
	sum := 0
	for !stack.IsEmpty() {
		value := stack.Pop()
		if value == 0 {
			sum += 0
		} else {
			sum += value.(int)
			stack.Push((value.(int) - 1))
		}
	}
	fmt.Println(sum)
	fmt.Println(add(5))
}

// 斐波那契数列
func fib(num int) int {
	if num == 1 || num == 2 {
		return 1
	} else {
		return fib(num-1) + fib(num-2)
	}
}

// 1 1 2 3 5 8 13
func main09() {
	stack := StackArray.NewStack()
	stack.Push(7)
	sum := 0
	for !stack.IsEmpty() {
		value := stack.Pop()
		if value == 1 || value == 2 {
			sum += 1
		} else {
			stack.Push(value.(int) - 1)
			stack.Push(value.(int) - 2)
		}
	}
	fmt.Println(sum)
	fmt.Println(fib(7))
}

func main10(){
	queue := Queue.NewQueue()
	queue.Enqueue("A")
	queue.Enqueue("B")
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
}
