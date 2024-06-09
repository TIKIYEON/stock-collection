package controllers

import (
	"StockCollection/Initializers"
	"StockCollection/models"
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

// PortfolioControllerRegister registers the portfolio controller
func PortfolioControllerRegister(router *gin.RouterGroup) {
	router.GET("/user/:user_id/portfolio", GetPortfolioByUserID)
	//router.POST("/user/:user_id/portfolio", CreatePortfolio)
	router.PUT("/user/:user_id/stock/:stock_id", AddStockToPortfolio)
	router.DELETE("/user/:user_id/stock/:stock_id/portfolio", RemoveStockFromPortfolio)
}

// GetPortfolio gets a portfolio
func GetPortfolio(c *gin.Context) {
	var portfolio models.Portfolio

	if err := c.BindJSON((&portfolio)); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	var existingPortfolio models.Portfolio
	result := Initializers.DB.Where(&models.Portfolio{PortfolioID: portfolio.PortfolioID}).First(&existingPortfolio)

	if result.Error != nil {
		log.Printf("Failed to find portfolio: %v", result.Error)
		c.JSON(500, gin.H{"error": "Failed to find portfolio"})
		return
	}

	c.JSON(200, gin.H{
		"portfolio": existingPortfolio,
	})
}

// GetPortfolios gets all portfolios
func GetPortfolios(c *gin.Context) {
	var portfolios []models.Portfolio

	result := Initializers.DB.Find(&portfolios)

	if result.Error != nil {
		log.Printf("Failed to find portfolios: %v", result.Error)
		c.JSON(500, gin.H{"error": "Failed to find portfolios"})
		return
	}

	c.JSON(200, gin.H{
		"portfolios": portfolios,
	})
}

func GetPortfolioByUserID(c *gin.Context) {
	var portfolio models.Portfolio

	UserID := c.Param("user_id")

	UserIDUint, err := strconv.ParseUint(UserID, 10, 64)
	if err != nil {
		log.Printf("Failed to parse user_id: %v", err)
		c.JSON(400, gin.H{"error": "user_id path parameter is not a number"})
		return
	}

	result := Initializers.DB.Where(&models.Portfolio{UserID: uint(UserIDUint)}).Preload("Stocks").First(&portfolio)

	if result.Error != nil {
		log.Printf("Failed to find portfolio given User ID")
		c.JSON(500, gin.H{"error": "Failed to find portfolio given User ID"})
		return
	}

	c.JSON(200, gin.H{
		"portfolio": portfolio,
	})
}

// CreatePortfolio creates a portfolio
func CreatePortfolio(c *gin.Context) {
	userID := c.Param("user_id")

	userIDUint, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		log.Printf("Failed to parse user_id: %v", err)
		c.JSON(400, gin.H{"error": "user_id path parameter is not a number"})
		return
	}

	// Check if the user exists
	var user models.User
	if err := Initializers.DB.Where(&models.User{UID: uint(userIDUint)}).First(&user).Error; err != nil {
		log.Printf("Failed to find user: %v", err)
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	var portfolio models.Portfolio
	if err := c.BindJSON(&portfolio); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	// Set the user_id for the portfolio
	portfolio.UserID = uint(userIDUint)

	// Save the portfolio to the database
	result := Initializers.DB.Create(&portfolio)
	if result.Error != nil {
		log.Printf("Failed to create portfolio1: %v", result.Error)
		c.JSON(500, gin.H{"error": "Failed to create portfolio"})
		return
	}

	c.JSON(201, gin.H{
		"portfolio": portfolio,
	})
}

func AddStockToPortfolio(c *gin.Context) {
	var portfolio models.Portfolio
	UserID := c.Param("user_id")

	UserIDUint, err := strconv.ParseUint(UserID, 10, 64)
	if err != nil {
		log.Printf("Failed to parse user_id: %v", err)
		c.JSON(400, gin.H{"error": "user_id path parameter is not a number"})
		return
	}

	Initializers.DB.Where(&models.Portfolio{UserID: uint(UserIDUint)}).Preload("Stocks").First(&portfolio)

	var stock models.Stock
	stockID := c.Param("stock_id")

	stockIDUint, err := strconv.ParseUint(stockID, 10, 64)
	if err != nil {
		log.Printf("Failed to parse stock_id: %v", err)
		c.JSON(400, gin.H{"error": "stock_id path parameter is not a number"})
		return
	}

	Initializers.DB.Where(&models.Stock{SID: uint(stockIDUint)}).First(&stock)

	//check if stock is already in portfolio
	for _, s := range portfolio.Stocks {
		if s.SID == stock.SID {
			c.JSON(400, gin.H{"error": "Stock is already in portfolio"})
			return
		}
	}

	if err := Initializers.DB.Model(&portfolio).Association("Stocks").Append(&stock); err != nil {
		fmt.Printf("Error associating stock with portfolio: %v\n", err)
	} else {
		fmt.Println("Successfully associated stock with portfolio.")
	}

	c.JSON(200, gin.H{
		"portfolio": portfolio,
	})
}

func RemoveStockFromPortfolio(c *gin.Context) {
	var portfolio models.Portfolio
	UserID := c.Param("user_id")

	UserIDUint, err := strconv.ParseUint(UserID, 10, 64)
	if err != nil {
		log.Printf("Failed to parse user_id: %v", err)
		c.JSON(400, gin.H{"error": "user_id path parameter is not a number"})
		return
	}

	Initializers.DB.Where(&models.Portfolio{UserID: uint(UserIDUint)}).Preload("Stocks").First(&portfolio)

	var stock models.Stock
	stockID := c.Param("stock_id")

	stockIDUint, err := strconv.ParseUint(stockID, 10, 64)
	if err != nil {
		log.Printf("Failed to parse stock_id: %v", err)
		c.JSON(400, gin.H{"error": "stock_id path parameter is not a number"})
		return
	}

	Initializers.DB.Where(&models.Stock{SID: uint(stockIDUint)}).First(&stock)

	//check if stock is in portfolio
	found := false
	for _, s := range portfolio.Stocks {
		if s.SID == stock.SID {
			found = true
			break
		}
	}

	if !found {
		c.JSON(400, gin.H{"error": "Stock is not in portfolio"})
		return
	}

	if err := Initializers.DB.Model(&portfolio).Association("Stocks").Delete(&stock); err != nil {
		fmt.Println("Error deleting stock from portfolio:", err)
		c.JSON(500, gin.H{"error": "Failed to delete stock from portfolio"})
		return
	} else {
		fmt.Println("Successfully deleted stock from portfolio.")
	}

	c.JSON(200, gin.H{
		"portfolio": portfolio,
	})
}
