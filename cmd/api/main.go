package main

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"

	"api-go-ent/internal/config"
	"api-go-ent/internal/database"
	"api-go-ent/internal/handlers"
	"api-go-ent/internal/middleware"

	"github.com/gin-gonic/gin"
)

// validateGinMode 验证 Gin 模式
func validateGinMode(mode string) bool {
	switch mode {
	case gin.DebugMode, gin.ReleaseMode, gin.TestMode:
		return true
	default:
		return false
	}
}

func main() {
	// 获取项目根目录
	_, b, _, _ := runtime.Caller(0)
	projectRoot := filepath.Join(filepath.Dir(b), "../..")

	// 加载配置
	cfg, err := config.LoadConfig(filepath.Join(projectRoot, "configs"))
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 验证并设置 Gin 模式
	if !validateGinMode(cfg.App.Env) {
		log.Printf("Warning: invalid gin mode '%s', defaulting to 'debug'", cfg.App.Env)
		cfg.App.Env = gin.DebugMode
	}
	gin.SetMode(cfg.App.Env)

	// 初始化数据库
	client, err := database.NewClient(cfg.Database.GetDSN())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer client.Close()

	// 创建 Gin 引擎
	r := gin.Default()

	// 注册中间件
	r.Use(middleware.ErrorHandler())
	r.Use(middleware.DatabaseMiddleware(client))

	// CORS 配置
	r.Use(middleware.CorsMiddleware(cfg.Cors))

	// 注册路由
	handlers.RegisterRoutes(r)

	// 启动服务器
	addr := fmt.Sprintf(":%d", cfg.App.Port)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
