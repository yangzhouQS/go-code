package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 创建文件
	fp, err := os.Create("./b.txt")

	if err != nil {
		fmt.Println("文件创建失败", err)
		return
	} else {
		defer fp.Close()
		fmt.Println(fp)
	}

	//	写入字符串内容
	//ret, _ := fp.WriteString("世界你好码")
	//fmt.Println(ret) // 返回的是写入文件的字节数

	// 写入字节
	buf := make([]byte, 1024)
	buf[0] = 'h'
	buf[1] = 'e'
	buf[2] = 'l'
	buf[3] = 'l'
	buf[4] = 'o'
	//n, _ := fp.Write(buf)
	//fmt.Println(n) // 1024

	//	在指定位置写入
	n, _ := fp.Seek(0, io.SeekEnd)
	ln, _ := fp.WriteAt(buf, n)
	fmt.Println(ln)
}
