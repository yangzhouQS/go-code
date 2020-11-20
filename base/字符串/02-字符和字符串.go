package main

import "fmt"

func main() {
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
