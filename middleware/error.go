package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorResponse 统一的错误响应结构
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// ErrorHandler 统一错误处理中间件
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// 只处理已经设置的错误
		if len(c.Errors) > 0 {
			err := c.Errors.Last()

			// 可以根据错误类型返回不同的状态码
			var statusCode int
			switch err.Type {
			case gin.ErrorTypeBind:
				statusCode = http.StatusBadRequest
			case gin.ErrorTypePrivate:
				statusCode = http.StatusInternalServerError
			default:
				statusCode = http.StatusInternalServerError
			}

			c.JSON(statusCode, ErrorResponse{
				Code:    statusCode,
				Message: err.Error(),
			})
			return
		}
	}
}
