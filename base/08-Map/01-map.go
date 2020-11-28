package main

import (
	"fmt"
	"sort"
	"strings"
)

func main01() {
	myMap := make(map[string]int)
	myMap["math"] = 99
	myMap["chinese"] = 100
	myMap["engilsh"] = 88
	userInfo := map[string]string{"userName": "lisi", "age": "26"}
	userInfo["S"] = "bsb"
	userInfo["b"] = "bsb"
	userInfo["A"] = "bsb"
	userInfo["a"] = "bsb"
	fmt.Println(myMap)    //map[chinese:100 engilsh:88 math:99]
	fmt.Println(userInfo) //map[age:26 userName:lisi]
}

//map的遍历
func main02() {
	list := map[string]string{"lisi": "99", "xiaoming": "88", "erniu": "109"}

	// key - value
	for k, v := range list {
		fmt.Println(k, v)
	}

	// key
	for k := range list {
		fmt.Println(k)
	}
}

//新增 & 删除 & 更新 & 查询

func main03() {
	m := make(map[string]int)
	// 新增
	m["chinese"] = 66
	m["match"] = 66
	fmt.Println(m) // map[chinese:66 match:66]

	// 删除
	delete(m, "chinese")
	fmt.Println(m) // map[match:66]

	// 更新
	m["match"] = 88
	fmt.Println(m) //map[match:88]

	// 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
	// 查询
	v, ok := m["match"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("not in key in map")
	}
	fmt.Println("------------")
	i, b := m["xasxasxsa"]
	fmt.Println(i, b)
}

func main04() {
	m := make(map[string]int)
	m["b"] = 10
	m["a"] = 10
	m["A"] = 10
	m["c"] = 10
	var keys []string
	// 把key单独抽取出来，放在数组中
	for k, _ := range m {
		keys = append(keys, k)
	}
	// 进行数组的排序
	sort.Strings(keys)
	// 遍历数组就是有序的了
	for _, k := range keys {
		fmt.Println(k, m[k])
	}
}

// method, map作为参数是值传递
func main05() {
	main02()
	fmt.Println("")
	var m map[int64]int64
	m = make(map[int64]int64, 1)
	fmt.Printf("m 原始的地址是：%p\n", m)
	changeM(m)
	fmt.Printf("m 改变后地址是：%p\n", m)
	fmt.Println("m 长度是", len(m))
	fmt.Println("m 参数是", m)

}

// 改变map的函数
func changeM(m map[int64]int64) {
	fmt.Printf("m 函数开始时地址是：%p\n", m)
	var max = 5
	for i := 0; i < max; i++ {
		m[int64(i)] = 2
	}
	fmt.Printf("m 在函数返回前地址是：%p\n", m)

}
func main06() {
	fmt.Println("")
	str := "how do you do"
	s := strings.Split(str, " ")
	fmt.Println("---", s)
	m := make(map[string]int, len(s))
	for _, v := range s {
		_, ok := m[v]
		if !ok {
			m[v] = 1
		} else {
			m[v] += 1
		}
	}
	fmt.Println(m)
}

/**
返回切片中指定字符串出现的索引
*/
func index(str []string, target string) int {
	for i, v := range str {
		if v == target {
			return i
		}
	}
	return -1
}

/**
判断字符串切片是否包含指定字符串
*/
func include(str []string, s string) bool {
	return index(str, s) >= 0
}

func filter(ws []string, f func(string) bool) bool {
	for _, v := range ws {
		if f(v) {
			return true
		}
	}
	return false
}

func every(nums []int, f func(int) bool) bool {
	for _, v := range nums {
		if !f(v) {
			return false
		}
	}
	return true
}

func main() {
	strSlice := []string{"a", "b", "c", "QS"}
	fmt.Println(index(strSlice, "b"))
	fmt.Println(include(strSlice, "f"))
	ret := filter(strSlice, func(s string) bool {
		if s == "QSS" {
			return true
		} else {
			return false
		}
	})
	fmt.Println(ret)
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	ret2 := every(nums, func(i int) bool {
		if i > 1 {
			return true
		} else {
			return false
		}
	})
	fmt.Println(ret2)
}
