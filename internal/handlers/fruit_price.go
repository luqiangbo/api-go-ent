package handlers

import (
	"net/http"

	"api-go-ent/ent"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateFruitPrice 创建水果价格记录
func CreateFruitPrice(c *gin.Context) {
	var fruitPrice ent.FruitPrice
	if err := c.ShouldBindJSON(&fruitPrice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := c.MustGet("db").(*ent.Client)
	created, err := client.FruitPrice.Create().
		SetName(fruitPrice.Name).
		SetPrice(fruitPrice.Price).
		SetUnit(fruitPrice.Unit).
		SetRemark(fruitPrice.Remark).
		Save(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, created)
}

// GetFruitPrice 获取单个水果价格记录
func GetFruitPrice(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}

	client := c.MustGet("db").(*ent.Client)
	fruitPrice, err := client.FruitPrice.Get(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	c.JSON(http.StatusOK, fruitPrice)
}

// ListFruitPrices 获取所有水果价格记录
func ListFruitPrices(c *gin.Context) {
	client := c.MustGet("db").(*ent.Client)

	fruitPrices, err := client.FruitPrice.Query().All(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, fruitPrices)
}

// UpdateFruitPrice 更新水果价格记录
func UpdateFruitPrice(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}

	var fruitPrice ent.FruitPrice
	if err := c.ShouldBindJSON(&fruitPrice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := c.MustGet("db").(*ent.Client)
	updated, err := client.FruitPrice.UpdateOneID(id).
		SetName(fruitPrice.Name).
		SetPrice(fruitPrice.Price).
		SetUnit(fruitPrice.Unit).
		SetRemark(fruitPrice.Remark).
		Save(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updated)
}

// DeleteFruitPrice 删除水果价格记录
func DeleteFruitPrice(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}

	client := c.MustGet("db").(*ent.Client)
	err = client.FruitPrice.DeleteOneID(id).Exec(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "record deleted"})
}
