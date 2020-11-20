package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, error := gorm.Open("mysql", "root:111111@(127.0.0.1)/demo?charset=utf8mb4&parseTime=True&loc=Local")
	defer db.Close()

	if error != nil {
		panic(error)
	}

	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "ok"})
	})
	router.Run(":3000")
}
