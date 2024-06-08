package main

import (
	"StockCollection/Initializers"
	"StockCollection/models"
	"log"
)

func init() {
	Initializers.LoadEnvVariables()
	Initializers.ConnectToDatabase()
}

func main() {
	err := Initializers.DB.AutoMigrate(&models.Stock{},
		&models.Stockelement{},
		&models.User{},
		&models.Portfolio{},
		&models.Portfoliostocks{})

	if err != nil {
		log.Fatalf("Failed to migrate: %v", err)
	} else {
		log.Println("Migration successful")
	}
}
