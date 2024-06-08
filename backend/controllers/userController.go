package controllers

import (
	"StockCollection/models"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"

	"StockCollection/Initializers"
)

func UserControllerRegister(router *gin.RouterGroup) {
	router.POST("/user", UserCreate)
	router.GET("/user", CheckUserExists)
	router.GET("/user/:user_id/portfolios", GetUserPortfolios)
}

func CheckUserExists(c *gin.Context) {
	// Get data from request
	var user models.User

	if err := c.BindJSON((&user)); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	// Check if user exists
	var existingUser models.User
	result := Initializers.DB.Where(&models.User{Mail: user.Mail, Password: user.Password}).First(&existingUser)

	if result.Error != nil {
		log.Printf("Failed to find user: %v", result.Error)
		c.JSON(500, gin.H{"error": "Failed to find user"})
		return
	}

	// Return response
	c.JSON(200, gin.H{
		"user": existingUser,
	})

}

// UserCreate creates a new user
func UserCreate(c *gin.Context) {
	// Get data from request
	var newUser models.User

	if err := c.BindJSON((&newUser)); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	// Create user in database
	result := Initializers.DB.Create(&newUser)

	if result.Error != nil {
		log.Printf("Failed to create user: %v", result.Error)
		c.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}

	// Return response
	c.JSON(201, gin.H{
		"user": newUser,
	})
}

// GetUserPortfolios gets all portfolios of a user
func GetUserPortfolios(c *gin.Context) {
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

	// Retrieve all portfolios of the user
	/* 	var portfolios []models.Portfolio
	   	if err := Initializers.DB.Where(&models.Portfolio{UserID: uint(userIDUint)}).
	   		Preload("User").
	   		Find(&portfolios).Error; err != nil {
	   		log.Printf("Failed to find portfolios: %v", err)
	   		c.JSON(500, gin.H{"error": "Failed to find portfolios"})
	   		return
	   	} */

	// Retreive user and preload portfolios and stocks
	var userWithPortfolios models.User
	if err := Initializers.DB.Where(&models.User{UID: uint(userIDUint)}).
		Preload("Portfolios.Stocks").
		First(&userWithPortfolios).Error; err != nil {
		log.Printf("Failed to find user: %v", err)
		c.JSON(500, gin.H{"error": "Failed to find user"})
		return
	}

	// Return response
	c.JSON(200, gin.H{
		"portfolios": userWithPortfolios.Portfolios,
	})
}
