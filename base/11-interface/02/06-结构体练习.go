package main

import (
	"fmt"
)

// 定义猴子结构体
type Monkey struct {
	Name string
}

func (m Monkey) climbing() { // 方法
	fmt.Println(m.Name, " 生来会爬树...")
}

// 定义实现飞翔接口
type BirdAble interface {
	Flying()
}

// 定义实现游泳接口
type FishAble interface {
	Swimming()
}

type LittleMonkey struct { // 继承
	Monkey // 匿名字段实现继承
}

func (lm LittleMonkey) Flying() {
	fmt.Println(lm.Name, " 通过学习会飞翔...")
}

func (lm LittleMonkey) Swimming() {
	fmt.Println(lm.Name, " 通过学习会游泳...")
}

func main() {
	littleMonkey := LittleMonkey{
		Monkey{"悟空"},
	}
	littleMonkey.climbing() // 继承
	littleMonkey.Flying()   // 接口
	littleMonkey.Swimming() // 接口
}
