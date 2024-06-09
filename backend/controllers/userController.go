package controllers

import (
	"StockCollection/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"StockCollection/Initializers"
)

func UserControllerRegister(router *gin.RouterGroup) {
	router.POST("/login", Login)
	router.POST("/register", UserCreate)
}

func CheckUserExists(mail string) bool {
	// Get user record from the database based on the provided email
	var user models.User
	if err := Initializers.DB.Where("mail = ?", mail).First(&user).Error; err != nil {
		// User not found in the database
		return false
	}
	// User exists in the database
	return true
}

func CheckUserCredentials(c *gin.Context, mail string, password string) bool {
	var user models.User
	// Check if email exists
	if err := Initializers.DB.Where("mail = ?", mail).First(&user).Error; err != nil {
		// c.JSON(400, gin.H{"error": "Email doesn't exists"})
		return false
	}
	// password matches
	if user.Password == password {
		c.JSON(201, gin.H{"user": user, "HX-Redirect": "stockies.html"})
		return true
	}
	// password doesnt match
	return false
}

func Login(c *gin.Context) {
	var loginData struct {
		Mail     string `json:"mail"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if CheckUserCredentials(c, loginData.Mail, loginData.Password) {
		return
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid email or password"})
		//c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}
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
		"user":    newUser,
		"success": true,
		"message": "User successfully created",
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
