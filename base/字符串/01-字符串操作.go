package main

import (
	"fmt"
	"strings"
)

func main2() {
	// 字符串字面量使用双引号（""）或者反引号（``）来创建。
	str := "hello golang string \"sss\" "
	str2 := `hello "word" with's
			可以换行
		`
	fmt.Println(str)  // hello golang string "sss"
	fmt.Println(str2) // hello "word" with's
}
func main3() {
	str := "hello golang string "
	str2 := `hello "word" with's`
	// 字符串拼接
	fmt.Println(str2 + str) // hello "word" with'shello golang string

	// 字符串追加
	str += "世界你好吗"
	fmt.Println(str) // hello golang string 世界你好吗

	// 字符串比价
	fmt.Println("a" < "b") // true
	fmt.Println("a" < "A") // false
}

func main4() {
	// 比较字符串

	/*
		第一个问题是，有些Unicode编码的字符可以用两个或者多个不同的字节序列来表示
		要处理非ASCII字符，这个问题也仅仅在以下情况下才存在：当我们有两个看起来一样的字符时，或者当我们从一个外部来源中读取UTF-8字节时，这个来源的码点到字节的映射是合法的UTF-8但又不同于Go语言的映射。如果这真的是一个问题，那么也可以写一个自定义的标准化函数来保证不出错。
		第二个问题是，有些情况下用户可能会希望把不同的字符看成相同的。
		第三个问题是，有些字符的排序是与语言相关的。
	*/
}

func main5() {
	//	字符和字符串

	// 字符切片转换为字符串
	str := ""
	for _, char := range []rune{'a', 'b', 'c', 'd', 'e'} {
		fmt.Println(char, string(char))
		str += string(char)
	}
	fmt.Println(str)

	// 转换为字符码点
	chars := []rune("hello")
	fmt.Println(chars) // [104 101 108 108 111]

	// 字符切片转换为字符串
	fmt.Println(string(chars)) // hello

	/*var buffer bytes.Buffer
	for {
		if piece, ok := getNextValidString(); ok {
			buffer.WriteString(piece)
		} else {
			break
		}
	}
	fmt.Print(buffer.String(),"\n")*/
}

// https://studygolang.com/articles/5769
func main() {
	str := "hello你好abc"
	fmt.Println("包含:", strings.Contains(str, "ab"))
	fmt.Println("开头:", strings.HasPrefix(str, "he"))
	fmt.Println("结尾:", strings.HasSuffix(str, "abc"))
	fmt.Println("长度:", len(str))
	fmt.Println("包含任一字符:", strings.ContainsAny(str, "def"))
	fmt.Println(strings.Contains(str, "")) // true
	fmt.Println("是否包含字符:", strings.ContainsRune(str, 'a'))
	fmt.Println("是否包含字符:", strings.ContainsRune(str, 'd'))
	fmt.Println("index首次出现索引:", strings.Index(str, "l"))
	fmt.Println("index首次出现索引:", strings.Index(str, "h")) // 0
	fmt.Println("index首次出现索引:", strings.Index(str, ""))  // 0
	fmt.Println("返回字符在字符串:", strings.IndexRune(str, 'a'))
}
