package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "普通的路由"})
	})

	// 路由组, 模块化使用
	userGroup := router.Group("/user")
	{
		userGroup.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "获取User"})
		})
		userGroup.POST("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "创建User"})
		})
		userGroup.DELETE("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "删除User"})
		})
		userGroup.PUT("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "更新User"})
		})
	}
	router.Run(":3000")
}
