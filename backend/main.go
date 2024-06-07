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
	controllers.UserControllerRegister(&r.RouterGroup)
	//r.POST("/register", controllers.UserCreate)
	r.Run() // listen and serve on port 8080
}
