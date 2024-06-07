package controllers

import (
	"StockCollection/Initializers"
	"StockCollection/models"
	"log"

	"github.com/gin-gonic/gin"
)

func StockElementsControllerRegister(router *gin.RouterGroup) {
	router.GET("/stockelements", GetStockElementsFromStockID)
}

func GetStockElementsFromStockID(c *gin.Context) {
	var stockElements []models.StockElement

	StockID := c.Query("stock_id")

	if StockID == "" {
		c.JSON(400, gin.H{"error": "missing stock_id query parameter"})
	}

	log.Printf("Received stock_id: %v", StockID)

	result := Initializers.DB.Where("stock_id = ?", StockID).Find(&stockElements)
	//result := Initializers.DB.Where(&models.Stock{SID: stock.SID}).First(&existingStock)

	if result.Error != nil {
		log.Printf("Failed to find stocks: %v", result.Error)
		c.JSON(500, gin.H{"error": "Failed to find stocks"})
		return
	}

	if len(stockElements) == 0 {
		log.Printf("No stockelements found for stock_id: %v", StockID)
		c.JSON(404, gin.H{"error": "No stockelements found"})
		return
	}

	log.Printf("Found stockelements: %v", stockElements)

	c.JSON(200, gin.H{
		"stockelements": stockElements,
	})
}
