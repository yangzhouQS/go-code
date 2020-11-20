package main

import (
	"fmt"
	"time"
)

func main() {
	// 当前的时间
	fmt.Println("当前时间: ", time.Now()) //当前时间:  2020-02-03 20:27:22.836858943 +0800 CST m=+0.000072078
	now := time.Now()
	// 时间类型
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	// 2020-02-03 20:30:44

	// 时间戳
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano()) // 纳秒时间戳

	/*
	const (
	Nanosecond  Duration = 1
	Microsecond          = 1000 * Nanosecond
	Millisecond          = 1000 * Microsecond
	Second               = 1000 * Millisecond
	Minute               = 60 * Second
	Hour                 = 60 * Minute
	)
	*/
	//	时间间隔计算
	fmt.Println(now.Add(time.Hour * 1))        // 1小时后
	fmt.Println(now.Add(time.Hour * 24))       // 24小时后
	fmt.Println(now.Add(time.Hour * 24 * 365)) // 1年后

	//	定时器
	fmt.Println("---定时器---")
	//timer := time.Tick(time.Second)
	//for i := range timer {
	//	fmt.Println("hello", timer, "---------", i)
	//}

	//	formatDemo时间格式化
	//formatDemo()
	fmt.Println(now.Format("2006/01/02 15:04:05 PM"))
	fmt.Println(now.Format("2006/01/02 15:04:05.000"))
	fmt.Println(now.Format("2006/01/02 15:04:05 Mon Jan"))
}

// 时间格式化
func formatDemo() {
	now := time.Now()
	// 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))
}
