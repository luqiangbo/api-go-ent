package main

import (
	"fmt"
	"log"
	"time"

	"api-go-ent/config"
	"api-go-ent/database"
	"api-go-ent/handlers"
	"api-go-ent/middleware"
	"api-go-ent/utils/logger"
	"api-go-ent/utils/validator"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化日志
	if err := logger.Setup(); err != nil {
		log.Fatalf("初始化日志失败: %v", err)
	}

	// 加载配置
	cfg, err := config.Load(".env")
	if err != nil {
		logger.Logger.Printf("Warning: %v, using default configuration", err)
		cfg = config.Get()
	}

	// 设置gin模式
	if cfg.Env.IsProduction() {
		gin.SetMode(gin.ReleaseMode)
	}

	// 初始化数据库客户端
	client := database.NewClient(cfg.Database)
	defer client.Close()

	// 设置验证器
	validator.Setup()

	// 创建 Gin 路由
	r := gin.New()

	// 使用中间件
	r.Use(gin.Recovery())
	r.Use(middleware.ErrorHandler())

	// 配置CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{cfg.Env.CorsAllowOrigins},
		AllowMethods:     []string{cfg.Env.CorsAllowMethods},
		AllowHeaders:     []string{cfg.Env.CorsAllowHeaders},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           time.Duration(cfg.Env.CorsMaxAge) * time.Second,
	}))

	// 配置日志
	r.Use(gin.LoggerWithWriter(gin.DefaultWriter, "/api/v1/fruit-prices/list"))

	// 初始化处理器
	fruitPriceHandler := handlers.NewFruitPriceHandler(client)

	// API版本分组
	v1 := r.Group("/api/v1")
	{
		// 水果价格相关路由
		fruits := v1.Group("/fruit-prices")
		{
			fruits.POST("/create", logRequest("创建水果价格"), fruitPriceHandler.Create)
			fruits.POST("/list", logRequest("获取所有水果价格"), fruitPriceHandler.GetAll)
			fruits.POST("/get", logRequest("获取单个水果价格"), fruitPriceHandler.GetByID)
			fruits.POST("/update", logRequest("更新水果价格"), fruitPriceHandler.Update)
			fruits.POST("/delete", logRequest("删除水果价格"), fruitPriceHandler.Delete)
		}
	}

	// 启动服务器
	serverAddr := fmt.Sprintf(":%d", cfg.Env.AppPort)
	logger.Logger.Printf("服务器启动在 %s", serverAddr)
	if err := r.Run(serverAddr); err != nil {
		logger.Logger.Fatalf("服务器启动失败: %v", err)
	}
}

// logRequest 记录请求日志的中间件
func logRequest(operation string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 请求开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 请求结束时间
		endTime := time.Now()

		// 记录日志
		logger.Logger.Printf(
			"[%s] %s %s %s 状态码:%d 耗时:%v",
			operation,
			c.Request.Method,
			c.Request.URL.Path,
			c.ClientIP(),
			c.Writer.Status(),
			endTime.Sub(startTime),
		)
	}
}
