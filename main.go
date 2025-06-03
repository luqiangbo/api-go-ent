package main

import (
	"log"

	"api-go-ent/config"
	"api-go-ent/database"
	"api-go-ent/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.DefaultConfig

	// 初始化数据库客户端
	client := database.NewClient(cfg.Database)
	defer client.Close()

	// 创建 Gin 路由
	r := gin.Default()

	// 初始化处理器
	fruitPriceHandler := handlers.NewFruitPriceHandler(client)

	// 注册路由 - 全部使用POST方法
	r.POST("/fruit-prices/create", fruitPriceHandler.Create)
	r.POST("/fruit-prices/list", fruitPriceHandler.GetAll)
	r.POST("/fruit-prices/get", fruitPriceHandler.GetByID)
	r.POST("/fruit-prices/update", fruitPriceHandler.Update)
	r.POST("/fruit-prices/delete", fruitPriceHandler.Delete)

	// 启动服务器
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
