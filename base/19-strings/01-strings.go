package main

import (
	"fmt"
	"strings"
)

func main() {
	var str = " hello word Abc deF "

	fmt.Println("------包含检测-----")
	fmt.Println(strings.Contains(str, "abc"))     // false
	fmt.Println(strings.ContainsRune(str, 'w'))   // 包含字符
	fmt.Println(strings.ContainsAny(str, "adef")) // 包含字符串任一字符

	fmt.Println("---大小写格式化--")
	fmt.Println(strings.ToUpper(str)) // 大写
	fmt.Println(strings.ToLower(str)) // 小写
	fmt.Println(strings.ToTitle(str)) // 全部大写

	fmt.Println("-----前缀匹配检测---")
	fmt.Println(strings.HasPrefix(str, " hello")) // 前缀检测
	fmt.Println(strings.HasSuffix(str, "def "))   // false

	fmt.Println(strings.Count(str, "o"))
	fmt.Println("------返回索引---")
	fmt.Println(strings.Index(str, "Abc"))    // 字符串第一次出现的索引
	fmt.Println(strings.IndexRune(str, 'e'))  // unicode码值 第一次出现的索引
	fmt.Println(strings.IndexByte(str, 'e'))  // 字符
	fmt.Println(strings.IndexAny(str, "dac")) // 任一utf-8码值

	fmt.Println("-----重复")
	fmt.Println(strings.Repeat("abc", 2)) // 字符串重复只指定次数

	fmt.Println("-----比较大小")
	fmt.Println(strings.Compare("a", "b")) // -1
	fmt.Println(strings.Compare("a", "a")) // 0
	fmt.Println(strings.Compare("b", "a")) // 1

	fmt.Println("---去除前后空格---")
	fmt.Println(strings.Trim(" abc ", " "))
	fmt.Println(strings.TrimSpace(" abc "))

	fmt.Println("----切片拼接为字符串---")
	sSlice := []string{"a", "b", "c", "d", "e"}
	fmt.Println(strings.Join(sSlice, "##")) // a##b##c
	str = "abcd"
	retSplit := strings.Split(str, "")
	fmt.Printf("%T ,%v \n", retSplit, retSplit)  // []string ,[a b c d]
	fmt.Println(strings.SplitN("abcdef", "", 2)) // 分割为几段

}
