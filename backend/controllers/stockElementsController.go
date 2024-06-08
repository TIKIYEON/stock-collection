package controllers

import (
	"StockCollection/Initializers"
	"StockCollection/models"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func StockElementsControllerRegister(router *gin.RouterGroup) {
	router.GET("/stockelements", GetStockElementsFromStockID)
}

func GetStockElementsFromStockID(c *gin.Context) {
	var stockElements []models.Stockelement

	StockID := c.Query("stock_id")

	stockIDUint, err := strconv.ParseUint(StockID, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "stock_id query parameter is not a number"})
	}

	log.Printf("Received stock_id: %v", StockID)

	//result := Initializers.DB.Where("stock_id = ?", uint(stockIDUint)).Find(&stockElements)
	// result := Initializers.DB.Where(&models.Stockelement{StockID: uint(stockIDUint)}).Find(&stockElements)
	result := Initializers.DB.Where(&models.Stockelement{StockID: uint(stockIDUint)}).
		Order("Date DESC").
		First(&stockElements)

	if result.Error != nil {
		log.Printf("Failed to find stocks: %v", result.Error)
		c.JSON(500, gin.H{"error": "Failed to find stocks"})
		return
	}

	log.Printf("Result: %v", result)

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
