package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	Name string `gorm:"type:varchar(128)"` // string默认长度为255, 使用这种tag重设。
	Age  int
}

func main() {
	db, error := gorm.Open("mysql", "root:111111@(127.0.0.1)/demo?charset=utf8mb4&parseTime=True&loc=Local")
	defer db.Close()

	if error != nil {
		panic(error)
	}

	db.AutoMigrate(&User{})

	//user := User{"王麻子", 26}
	//db.Create(&user)
	var user User
	fmt.Println(db.First(&user))

	fmt.Println("run ....")
}
