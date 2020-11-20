package main

import (
	"fmt"
)

type USB interface {
	Start()
	Stop()
}
type Phone struct {
	name string
}

// Phone实现USB接口的方法
func (p *Phone) Start() {
	fmt.Println(p.name, "开始工作")
}
func (p *Phone) Stop() {
	fmt.Println(p.name, "结束工作")
}

// 上传和下载的接口
type UploadAndDownload interface {
	Upload()   // 上传
	Download() // 下载
}
type Computer struct {
	name string
}

func (c *Computer) Start() {
	fmt.Println(c.name, "开始工作")
}

func (c *Computer) Stop() {
	fmt.Println(c.name, "结束工作")
}

func (c *Computer) Download() {
	fmt.Println(c.name, "开始下载")
}

func (c *Computer) Upload() {
	fmt.Println(c.name, "开始上传")
}

func main() {
	p := Phone{"苹果手机"}
	p.Start()
	p.Stop()
	/*
		1.接口不可以直接实例化
		2.接口可以指向实现了该类型的自定义变量
		3.接口中的方法只有方法体
		4.结构体实现了接口的所有方法, 描述为自定义类型实现了接口
		5.自定义了实现了接口的某个方法,就可以将该自定义类型的实例赋值给接口变量
		6.只要是自定义类型就可以实现接口,不仅限于结构体类型
		7.一个自定义类型可以实现多个接口
		8.接口中不可以有变量
		8.
	*/
	c := Computer{"MAC笔记本"}
	var computer Computer = c
	computer.Start()
	computer.Stop()
	var downloadAndUpload UploadAndDownload = &Computer{"华硕笔记本"}
	downloadAndUpload.Upload()
	downloadAndUpload.Download()
}
