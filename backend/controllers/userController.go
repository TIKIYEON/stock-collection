package controllers

import (
	"StockCollection/models"
	"log"

	"github.com/gin-gonic/gin"

	"StockCollection/Initializers"
)

/*
func UserControllerRegister(router *gin.RouterGroup) {
    router.GET("/user", getUser)
    router.POST("/user", createUser)
    router.PUT("/user", updateUser)
    router.DELETE("/user", deleteUser)
}
*/


func UserCreate(c *gin.Context) {
    // Get data from request
    /*
    var newUser struct {
        UID uint `json:"uid"`
        Password string `json:"password"`
        Mail string `json:"mail"`
        PhoneNumber string `json:"phone_number"`
    }

    if err := c.BindJSON((&newUser)); err != nil {
        c.JSON(400, gin.H{"error": "Invalid request"})
        return
    }
    */
    var newUser models.User

    if err := c.BindJSON((&newUser)); err != nil {
        c.JSON(400, gin.H{"error": "Invalid request"})
        return
    }

    // Create user in database
    /*
    user := models.User{UID: newUser.UID, Password: newUser.Password, Mail: newUser.Mail, PhoneNumber: newUser.PhoneNumber}
    result := Initializers.DB.Create(&user)

    if result.Error != nil {
        log.Printf("Failed to create user: %v", result.Error)
        c.JSON(500, gin.H{"error": "Failed to create user"})
        return
    }
    */
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
