package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func StatConst() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// 开始时间
		statrtTime := time.Now()

		// 设置上下文信息
		ctx.Set("userName", "李四")

		// 程序继续往下执行
		ctx.Next()

		fmt.Println("StatConst")
		endTime := time.Since(statrtTime)
		fmt.Println("程序路由执行时间:", endTime)
	}
}
func main() {
	router := gin.Default()

	// 注册中间件
	router.Use(StatConst())

	router.GET("/", func(ctx *gin.Context) {
		// 获取在上下文设置的信息
		userName, ok := ctx.Get("userName")
		if ok {
			fmt.Println(userName)
		}
		fmt.Println()
		ctx.JSON(http.StatusOK, gin.H{"message": "get index"})
	})

	router.Run(":3000")
}
