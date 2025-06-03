package handlers

import (
	"net/http"

	"api-go-ent/ent"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FruitPriceHandler struct {
	client *ent.Client
}

func NewFruitPriceHandler(client *ent.Client) *FruitPriceHandler {
	return &FruitPriceHandler{client: client}
}

type CreateFruitPriceRequest struct {
	Name   string  `json:"name" binding:"required"`
	Price  float64 `json:"price" binding:"required"`
	Unit   string  `json:"unit" binding:"required"`
	Remark string  `json:"remark" binding:"required"`
}

type IDRequest struct {
	ID string `json:"id" binding:"required"`
}

func (h *FruitPriceHandler) Create(c *gin.Context) {
	var req CreateFruitPriceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fruitPrice, err := h.client.FruitPrice.Create().
		SetName(req.Name).
		SetPrice(req.Price).
		SetUnit(req.Unit).
		SetRemark(req.Remark).
		Save(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, fruitPrice)
}

func (h *FruitPriceHandler) GetAll(c *gin.Context) {
	fruitPrices, err := h.client.FruitPrice.Query().All(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, fruitPrices)
}

func (h *FruitPriceHandler) GetByID(c *gin.Context) {
	var req IDRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := uuid.Parse(req.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	fruitPrice, err := h.client.FruitPrice.Get(c.Request.Context(), id)
	if err != nil {
		if ent.IsNotFound(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "FruitPrice not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, fruitPrice)
}

type UpdateFruitPriceRequest struct {
	ID     string  `json:"id" binding:"required"`
	Name   string  `json:"name" binding:"required"`
	Price  float64 `json:"price" binding:"required"`
	Unit   string  `json:"unit" binding:"required"`
	Remark string  `json:"remark" binding:"required"`
}

func (h *FruitPriceHandler) Update(c *gin.Context) {
	var req UpdateFruitPriceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := uuid.Parse(req.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	fruitPrice, err := h.client.FruitPrice.UpdateOneID(id).
		SetName(req.Name).
		SetPrice(req.Price).
		SetUnit(req.Unit).
		SetRemark(req.Remark).
		Save(c.Request.Context())
	if err != nil {
		if ent.IsNotFound(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "FruitPrice not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, fruitPrice)
}

func (h *FruitPriceHandler) Delete(c *gin.Context) {
	var req IDRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := uuid.Parse(req.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	err = h.client.FruitPrice.DeleteOneID(id).Exec(c.Request.Context())
	if err != nil {
		if ent.IsNotFound(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "FruitPrice not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "FruitPrice deleted successfully"})
}
