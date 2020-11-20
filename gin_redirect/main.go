package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "ok index"})
	})

	// HTTP重定向
	router.GET("/a", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"router": "/a"})
	})
	router.GET("/b", func(ctx *gin.Context) {
		// 路由重定向
		ctx.Request.URL.Path = "/a"
		router.HandleContext(ctx)

		// http重定向
		//ctx.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})
	router.Run(":3000")
}
