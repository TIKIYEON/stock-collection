package controllers

import (
	"StockCollection/Initializers"
	"StockCollection/models"
	"log"

	"github.com/gin-gonic/gin"
)

func StockControllerRegister(router *gin.RouterGroup) {
	router.GET("/stock", GetStock)
	router.GET("/stocks", GetStocks)
}

func GetStock(c *gin.Context) {
	var stock models.Stock

	if err := c.BindJSON((&stock)); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	var existingStock models.Stock
	result := Initializers.DB.Where(&models.Stock{SID: stock.SID}).First(&existingStock)

	if result.Error != nil {
		log.Printf("Failed to find stock: %v", result.Error)
		c.JSON(500, gin.H{"error": "Failed to find stock"})
		return
	}

	c.JSON(200, gin.H{
		"stock": existingStock,
	})
}

func GetStocks(c *gin.Context) {
	var stocks []models.Stock

	result := Initializers.DB.Find(&stocks)

	if result.Error != nil {
		log.Printf("Failed to find stocks: %v", result.Error)
		c.JSON(500, gin.H{"error": "Failed to find stocks"})
		return
	}

	c.JSON(200, gin.H{
		"stocks": stocks,
	})
}
