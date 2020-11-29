package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Hostname())    // loop <nil>
	fmt.Println(os.Getpagesize()) // 4096
	//fmt.Println(os.Environ()) // 返回环境变量
	fmt.Println(os.Getenv("tmp")) // D:\tools\TEMP\TMP
	fmt.Println(os.Getuid())      // -1
	fmt.Println(os.Geteuid())     // -1
}
