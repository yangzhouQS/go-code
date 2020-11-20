package main

import (
	"fmt"
	"strings"
)

func main() {
	// 声明字符串
	s := "hello word"
	fmt.Println(s)
	//	多行字符串
	s1 := `
		a
a1
		c
`
	s2 := "世界你好"
	fmt.Println(s1)
	// 字符串的长度
	fmt.Println(len(s))
	fmt.Println(len(s1))

	// 字符串的拼接 + 号
	ret := fmt.Sprintf("%s%s%d", s, s2, 100)
	fmt.Println(ret)                         //hello word世界你好100
	fmt.Println(strings.Split(s, ""))        //[h e l l o   w o r d]
	fmt.Println(strings.Split("abcdef", "")) // [a b c d e f]
	fmt.Println(strings.HasPrefix(s, "he"))
	fmt.Println(strings.HasSuffix(s, "rd"))
	fmt.Println(strings.Contains(s, "llo"))

	s = "abcdaef世界"
	fmt.Println(strings.Index(s, "d"))
	fmt.Println(strings.LastIndex(s, "a"))

	//	 字符串遍历
	for i := 0; i < len(s); i++ {
		fmt.Println(string(s[i]), s[i])
	}
	for index, v := range s {
		//fmt.Println(index, v)
		fmt.Printf("%d -> %c\n", index, v)
	}
	// 字符串修改
	a1 := "北京"
	a2 := []rune(a1)
	a2[0] = '南'
	fmt.Println(string(a2))
}
