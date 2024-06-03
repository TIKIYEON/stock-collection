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
