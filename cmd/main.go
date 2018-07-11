package main

import (
	"github.com/gin-gonic/gin"
	"gin-app/middleware"
	"gin-app/app"
	"gin-app/controller"
)

func main() {

	r := gin.New()
	r.Use(middleware.Filter())
	r.GET("/ping", func(context *gin.Context) {
		context.JSONP(200, gin.H{
			"message": app.Config,
		})
	})

	r.GET("/login", controller.Login)
	r.Run(":9090")
}
