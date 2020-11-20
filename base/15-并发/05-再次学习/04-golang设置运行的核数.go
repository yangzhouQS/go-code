package main

import (
	"fmt"
	"runtime"
)

//设置golang运行的核数
//1.8 版本以上的会自动设置
func main() {
	//设置CPU运行的核数
	//NumCPU 返回本地机器的逻辑cpu个数
	num := runtime.NumCPU()

	fmt.Println(num) // 12

	//GOMAXPROCS 设置可同时执行的最大CPU数
	runtime.GOMAXPROCS(num)
}
