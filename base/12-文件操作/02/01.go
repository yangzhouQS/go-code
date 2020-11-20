package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func read1(filePath string) {
	file, err := os.OpenFile(filePath, os.O_RDWR, 066)
	if err != nil {
		fmt.Println("文件读取失败:", err)
	} else {
		defer file.Close()
	}
	reader := bufio.NewReader(file)
	for {
		// 安装每一行的换行符读取
		str, err := reader.ReadString('\n')
		fmt.Println(string(str))
		if err == io.EOF { // 读取结束的标志
			fmt.Println("读取文件结束")
			break
		}
	}
}
func read2(filePath string) {
	// 读取文件不太大的文件
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(file))
}
func read3(filePath string)  {

}
func main() {
	const filePath = "./12-文件操作/b.txt"
	read2(filePath)
}
