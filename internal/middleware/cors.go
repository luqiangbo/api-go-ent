package middleware

import (
	"time"

	"api-go-ent/internal/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CorsMiddleware 创建 CORS 中间件
func CorsMiddleware(cfg config.CorsConfig) gin.HandlerFunc {
	corsConfig := cors.Config{
		AllowOrigins:     cfg.AllowOrigins,
		AllowMethods:     cfg.AllowMethods,
		AllowHeaders:     cfg.AllowHeaders,
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           time.Duration(cfg.MaxAge) * time.Second,
	}
	return cors.New(corsConfig)
}
