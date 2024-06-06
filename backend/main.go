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
	r.POST("/register", controllers.UserCreate)
	r.Run() // listen and serve on port 8080
}

/*
CRUD links to consider:
- https://fenyuk.medium.com/golang-crud-in-rest-api-in-a-generic-way-9c395a60309e
- https://dev.to/samzhangjy/restful-crud-with-golang-for-beginners-23ia
*/
