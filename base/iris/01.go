package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context" //需要单独引入
)

func main() {
	app := iris.New()
	app.Get("/", func(ctx context.Context) {})
	app.Run(iris.Addr(":8080"))
}
