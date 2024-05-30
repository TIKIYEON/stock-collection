package main

import (
	"github.com/gin-gonic/gin"
    "example.com/m/v2/Initializers"
)

func main() {
    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Welcome to the Gin API",
        })
    })
    r.Run() // listen and serve on
    initializers.ConnectDatabse()
}
