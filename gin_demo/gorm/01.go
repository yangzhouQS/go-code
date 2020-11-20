package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

type UserInfo struct {
	ID   uint
	Name string
	Age  uint
}

func main() {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "hello gin",
		})
	})
	// ? 传递参数
	r.GET("/user", func(ctx *gin.Context) {
		query := ctx.Query("id")
		defaultQuery := ctx.DefaultQuery("name", "admin")
		fmt.Println(query)
		ctx.JSON(http.StatusOK, gin.H{
			"defaultQuery": defaultQuery,
			"message":      "ok",
			"query":        query,
		})
	})
	// 位置参数
	r.GET("/user/:id", func(ctx *gin.Context) {
		params := ctx.Param("id")
		fmt.Println(params)
		ctx.JSON(http.StatusOK, gin.H{
			"params": params,
			"message":      "ok",
		})
	})
	//
	r.Run(":8200")
}
