package main

import (
	"StockCollection/Initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	Initializers.LoadEnvVariables()
	Initializers.ConnectToDatabase()
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on
}

/*
CRUD links to consider:
- https://fenyuk.medium.com/golang-crud-in-rest-api-in-a-generic-way-9c395a60309e
- https://dev.to/samzhangjy/restful-crud-with-golang-for-beginners-23ia
*/
