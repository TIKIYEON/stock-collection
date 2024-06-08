package controllers

import (
	"StockCollection/Initializers"
	"StockCollection/models"
	"log"

	"github.com/gin-gonic/gin"
)

func PortfolioStockControllerRegister(router *gin.RouterGroup) {
	router.POST("/portfoliostock", CreatePortfolioStock)
}

func CreatePortfolioStock(c *gin.Context) {
	var portfolioStock models.Portfoliostocks

	if err := c.BindJSON(&portfolioStock); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	// Save the portfolio-stock association to the database
	result := Initializers.DB.Create(&portfolioStock)
	if result.Error != nil {
		log.Printf("Failed to create portfolio-stock association: %v", result.Error)
		c.JSON(500, gin.H{"error": "Failed to create portfolio-stock association"})
		return
	}

	c.JSON(201, gin.H{
		"portfolio_stock": portfolioStock,
	})
}
