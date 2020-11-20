package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		如果文件不存在不会自动创建
		通过open函数打开的文件只能读取, 不能写入
	*/
	// fp, err := os.Open("D:/code/goproject/src/base/12-文件操作/b.txt")
	fp, err := os.Open("./a.txt")
	defer fp.Close()
	if err != nil {
		fmt.Println("打开文件失败", err)
	} else {
		fmt.Println(fp)
	}
}
