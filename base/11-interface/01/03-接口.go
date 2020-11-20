package main

// 接口声明
type GetPerson interface {
	GetName(id int) (name string, err error)
	GetAge(id int) (age int, err error)
	GetGender(id int) (gender string, err error)
}


func main() {
}
