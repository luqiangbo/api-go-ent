package handlers

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册所有路由
func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			fruitPrices := v1.Group("/fruit-prices")
			{
				fruitPrices.POST("/list", ListFruitPrices)
				fruitPrices.POST("/create", CreateFruitPrice)
				fruitPrices.POST("/detail", GetFruitPrice)
				fruitPrices.POST("/update", UpdateFruitPrice)
				fruitPrices.POST("/delete", DeleteFruitPrice)
			}
		}
	}
}
