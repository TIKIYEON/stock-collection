package main

import (
	"StockCollection/Initializers"

	"StockCollection/controllers"

	"github.com/gin-gonic/gin"

	"net/http"
)

func init() {
	Initializers.LoadEnvVariables()
	Initializers.ConnectToDatabase()
}

func main() {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
        c.Writer.Header().Set("Access-Control-Max-Age", "86400")
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(200)
            return
        }
        c.Next()
    })

	r.Static("/static", "./frontend")

	controllers.UserControllerRegister(&r.RouterGroup)
	// controllers.CheckUserExists(&r.RouterGroup)
	controllers.StockControllerRegister(&r.RouterGroup)
	controllers.StockElementsControllerRegister(&r.RouterGroup)
	controllers.PortfolioControllerRegister(&r.RouterGroup)

	r.POST("/login", loginHandler)



	r.Run() // listen and serve on port 8080
}

func loginHandler(c *gin.Context) {
    // Retrieve username and password from the request
    mail := c.PostForm("mail")
    password := c.PostForm("password")

    // Check if user exists and if the provided password matches
    if controllers.CheckUserExists(c, mail, password) {
        // Authentication successful
        c.String(http.StatusOK, "Login successful!")
    } else {
        // Authentication failed
        c.String(http.StatusUnauthorized, "Invalid username or password")
    }
}