package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// 1.声明hero类型
type Hero struct {
	Name string
	Age  int
}

// 2.创建hero类型的切片
type HeroList []Hero

// 3.实现interface接口的三个方法
/*type Interface interface {
	Len()
	Less()
	Swap( )
}*/
func (h HeroList) Len() int {
	return len(h)
}

// 排序标准
func (h HeroList) Less(i, j int) bool {
	return h[j].Age > h[i].Age
}

// 排序标准成立的时候排序规则
func (h HeroList) Swap(i, j int) {
	// temp := h[i]
	// h[i] = h[j]
	// h[j] = temp
	h[i], h[j] = h[j], h[i]
}

func main() {
	var heroSlice HeroList
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("小英雄 -- %d", i+1),
			Age:  rand.Intn(100),
		}
		heroSlice = append(heroSlice, hero)
	}
	fmt.Println("-----------------排序前-------")
	for _, hero := range heroSlice {
		fmt.Println(hero)
	}
	fmt.Println("--------------排序后--------")
	sort.Sort(heroSlice)
	for _, hero := range heroSlice {
		fmt.Println(hero)
	}
}
