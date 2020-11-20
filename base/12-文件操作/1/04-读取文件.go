package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main01() {
	filePath := "./b.txt"
	fp, err := os.Open(filePath)
	// 打开文件失败: 1>文件不存在2>文件权限;3>文件打开的上限
	if err != nil {
		fmt.Println("打开文件失败", err)
	}
	defer fp.Close()
	//	1.将文件整个读入内存
	content, error := ioutil.ReadAll(fp)
	if error != nil {
		fmt.Println("读取文件内容失败", error)
	}
	fmt.Println(string(content))
	fmt.Println("----")
	content2, error2 := ioutil.ReadFile(filePath)
	if error2 != nil {
		panic(error2)
	}
	fmt.Println(string(content2))
}

func main() {
	filePath := "./b.txt"
	fp, err := os.Open(filePath)
	// 打开文件失败: 1>文件不存在2>文件权限;3>文件打开的上限
	if err != nil {
		fmt.Println("打开文件失败", err)
	}
	defer fp.Close()

	// 创建切片缓冲区
	r := bufio.NewReader(fp)
	// 2.按字节读取文件内容
	chunks := make([]byte, 0)
	buf := make([]byte, 1024) //一次读取多少个字节
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		fmt.Println(string(buf[:n]))
		break
		if 0 == n {
			break
		}
		chunks = append(chunks, buf[:n]...)
	}
	fmt.Println(string(chunks))
}

// 3、按行读取
