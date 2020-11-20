package main

import (
	"fmt"
	"os"
)

// 新增/查看/删除学生
type Student struct {
	id   int64
	name string
}
type ClassRoom struct {
	id        int64
	className string
}

func NewStudent(id int64, name string) *Student {
	return &Student{
		name: name,
		id:   id,
	}
}
func printList() {
	fmt.Println("你选择了1 查看所有学生")
	for k, v := range studentList {
		fmt.Printf("学生编号: %d, 姓名: %s \n", k, v.name)
	}
}
func stdentCreate() {
	var (
		name string
		id   int64
	)
	fmt.Println("你选择了2 创建学生")
	fmt.Print("请输入学号:")
	fmt.Scan(&id)
	fmt.Print("请输入姓名:")
	fmt.Scan(&name)

	stu := NewStudent(id, name)
	studentList[id] = stu
}
func stdentDelete() {
	fmt.Println("你选择了3 删除学生")
	var id int64
	fmt.Println("请输入学生id编号:")
	fmt.Scan(&id)
	id = int64(id)
	_, isId := studentList[id]
	if isId {
		delete(studentList, id)
		fmt.Println("删除成功拉~~~")
	} else {
		fmt.Println("输入的id不存在,请从新操作")
	}
}

var studentList = make(map[int64]*Student, 50)

func main() {
	var input string
	for {
		fmt.Println(`1>查看学生
2>新增学生
3>删除学生
4>退出
`)
		fmt.Print("请输入你的操作选项:")
		fmt.Scanln(&input)
		switch input {
		case "1":
			printList()
		case "2":
			stdentCreate()
		case "3":
			stdentDelete()
		case "4":
			os.Exit(1)

		default:
			fmt.Println("输入错误")
		}
	}
}
