package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger 日志中间件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		start := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		end := time.Now()
		latency := end.Sub(start)

		// 请求方法
		method := c.Request.Method
		// 请求路由
		uri := c.Request.RequestURI
		// 状态码
		status := c.Writer.Status()
		// 请求IP
		clientIP := c.ClientIP()

		log.Printf("[GIN] %v | %3d | %13v | %15s | %-7s %s",
			end.Format("2006/01/02 - 15:04:05"),
			status,
			latency,
			clientIP,
			method,
			uri,
		)
	}
}

// ErrorHandler 错误处理中间件
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// 只处理内部错误
		if len(c.Errors) > 0 {
			c.JSON(500, gin.H{
				"errors": c.Errors.Errors(),
			})
		}
	}
}
