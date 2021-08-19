package main

import (
	"github.com/kataras/iris/v12"
	"log"
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
	err := app.Run(
		// Start the web server at localhost:8080
		iris.Addr(":9000"),
		// skip err server closed when CTRL/CMD+C pressed:
		iris.WithoutServerError(iris.ErrServerClosed),
		// enables faster json serialization and more:
		iris.WithOptimizations,
	)

	if err != nil {
		log.Println(err.Error())
	}
}
