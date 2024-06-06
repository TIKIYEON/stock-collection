package main

import (
	"StockCollection/Initializers"
	"StockCollection/models"
)

func init() {
    Initializers.LoadEnvVariables()
    Initializers.ConnectToDatabase()
}

func main() {
    Initializers.DB.AutoMigrate(&models.Stock{}, &models.StockElement{}, &models.User{})
}
