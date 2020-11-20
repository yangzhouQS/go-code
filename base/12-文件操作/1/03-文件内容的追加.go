package main

import (
	"fmt"
	"os"
)

func main() {
	fp, err := os.OpenFile("./a.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0)
	defer fp.Close()
	if err != nil {
		fmt.Println("打开文件失败", err)
		return
	}
	fp.WriteString("hello")
}
