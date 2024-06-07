package controllers

import (
	"StockCollection/models"
	"log"

	"github.com/gin-gonic/gin"

	"StockCollection/Initializers"
)

func UserControllerRegister(router *gin.RouterGroup) {
	router.POST("/user", UserCreate)
	router.GET("/user", CheckUserExists)
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
