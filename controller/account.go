package controller

import (
	"github.com/gin-gonic/gin"
	"gin-app/app"
)

func Login (context *gin.Context) {
	context.JSONP(200, gin.H{
		"message": app.Config,
	})
}