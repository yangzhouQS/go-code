package routers

import (
	. "blog-service/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	app := gin.New()

	// gin中间件
	app.Use(gin.Logger())
	app.Use(gin.Recovery())

	// 路由设计
	apiv1 := app.Group("api/v1")
	{
		// tag路由
		apiv1.GET("/tags", Tag.List)
		apiv1.POST("/tags", Tag.Create)
		apiv1.DELETE("/tags/:id", Tag.Delete)
		apiv1.PUT("/tags/:id", Tag.Update)
		apiv1.PATCH("/tags/:id/state", Tag.Update)

		// article路由
		apiv1.GET("/article")
		apiv1.GET("/article/:id")
		apiv1.POST("/article")
		apiv1.PUT("/article/:id")
		apiv1.DELETE("/article/:id")
		apiv1.PATCH("/article/:id/state")
	}
	return app
}
