package main

import (
	"StockCollection/Initializers"

	"StockCollection/controllers"

	"github.com/gin-gonic/gin"
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
	controllers.StockControllerRegister(&r.RouterGroup)
	controllers.StockElementsControllerRegister(&r.RouterGroup)
	controllers.PortfolioControllerRegister(&r.RouterGroup)
	r.Run() // listen and serve on port 8080
}
