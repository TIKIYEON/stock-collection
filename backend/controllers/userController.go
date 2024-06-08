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
	// router.GET("/user", CheckUserExists)
	// router.GET("/user/:user_id/portfolio", GetUserPortfolio)
}

func CheckUserExists(c *gin.Context, mail, password string) bool {
	// Get user record from the database based on the provided username
	var user models.User
	if err := Initializers.DB.Where("mail = ?", mail).First(&user).Error; err != nil {
		// User not found in the database
		return false
	}

	// Compare the password provided in the request with the stored password
	if user.Password == password {
		// Passwords match, indicating valid credentials
		return true
	}

	// Passwords don't match
	return false
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

	// Create portfolio for the user
	var newPortfolio models.Portfolio
	newPortfolio.UserID = newUser.UID
	Initializers.DB.Create(&newPortfolio)

	// Return response
	c.JSON(201, gin.H{
		"user": newUser,
	})
}

// GetUserPortfolios gets all portfolios of a user
func GetUserPortfolio(c *gin.Context) {
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
	var userWithPortfolio models.User
	if err := Initializers.DB.Where(&models.User{UID: uint(userIDUint)}).
		Preload("Portfolios.Stocks").
		First(&userWithPortfolio).Error; err != nil {
		log.Printf("Failed to find user: %v", err)
		c.JSON(500, gin.H{"error": "Failed to find user"})
		return
	}

	// Return response
	c.JSON(200, gin.H{
		"portfolios": userWithPortfolio.Portfolio,
	})
}
