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
	controllers.StockControllerRegister(&r.RouterGroup)
	controllers.StockElementsControllerRegister(&r.RouterGroup)
	r.Run() // listen and serve on port 8080
}
