package middleware

import (
	"api-go-ent/ent"

	"github.com/gin-gonic/gin"
)

// DatabaseMiddleware 注入数据库客户端到上下文
func DatabaseMiddleware(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", client)
		c.Next()
	}
}
