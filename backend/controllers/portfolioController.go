package controllers

import (
	"StockCollection/Initializers"
	"StockCollection/models"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

// PortfolioControllerRegister registers the portfolio controller
func PortfolioControllerRegister(router *gin.RouterGroup) {
	router.GET("/portfolio", GetPortfolio)
	router.GET("/portfolios", GetPortfolios)
	router.POST("/user/:user_id/portfolio", CreatePortfolio)
}

// GetPortfolio gets a portfolio
func GetPortfolio(c *gin.Context) {
	var portfolio models.Portfolio

	if err := c.BindJSON((&portfolio)); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	var existingPortfolio models.Portfolio
	result := Initializers.DB.Where(&models.Portfolio{PID: portfolio.PID}).First(&existingPortfolio)

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

	// Load the related User entity
	if err := Initializers.DB.Model(&portfolio).Association("User").Find(&portfolio.User); err != nil {
		log.Printf("Failed to load related user: %v", err)
		c.JSON(500, gin.H{"error": "Failed to load related user"})
		return
	}

	c.JSON(201, gin.H{
		"portfolio": portfolio,
	})
}
