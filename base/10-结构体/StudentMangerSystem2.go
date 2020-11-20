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
type StudentList struct {
	Map map[int64]*Student
}

func newStudent(id int64, name string) *Student {
	return &Student{
		name: name,
		id:   id,
	}
}
func (s *StudentList) studentPrint() {
	fmt.Println("你选择了1 查看所有学生")
	//for k, v := range s {
	//	fmt.Printf("学生编号: %d, 学生姓名: %s \n", k, v)
	//}
	fmt.Println(s)
}
func (s *StudentList) studentAdd() {
	var (
		name string
		id   int64
	)
	fmt.Println("你选择了2 创建学生")
	fmt.Print("请输入学号:")
	fmt.Scan(&id)
	fmt.Print("请输入姓名:")
	fmt.Scan(&name)

	stu := newStudent(id, name)
	fmt.Println(s, stu)
}

func main() {
	s := &StudentList{}
	fmt.Println(s, Student{})
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
			s.studentPrint()
		case "2":
			s.studentAdd()
		case "3":
		case "4":
			os.Exit(1)

		default:
			fmt.Println("输入错误")
		}
	}
}
