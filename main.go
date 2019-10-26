package main

import (
	"github.com/kataras/iris"
	"net/http"
)

func main() {
	app := iris.New()
	app.Get("/", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"code":  http.StatusOK,
			"data": "hello world!",
		})
	})

	app.Run(iris.Addr(":9000"), iris.WithoutServerError(iris.ErrServerClosed))
}