package middleware

import (
	"github.com/gin-gonic/gin"
	"io"
	"fmt"
)

func Filter () gin.HandlerFunc {
	return FilterWithWriter(gin.DefaultWriter)
}

func FilterWithWriter(out io.Writer) gin.HandlerFunc{
	return func(c *gin.Context) {
		ip := c.ClientIP()
		nonce := c.Query("nonce")
		timestamp := c.Query("timestamp")
		sign := c.Query("sign")
		fmt.Println(ip, nonce, timestamp, sign)

		//TODO 做一些安全性验证 防重放 限制访问频率的工作
	}
}

