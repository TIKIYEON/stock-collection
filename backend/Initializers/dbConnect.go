package Initializers

import (
	"StockCollection/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func ConnectToDatabase() {
	var err error
	dsn := os.Getenv("DB")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		panic("Failed to connect to database!")
	}

	errs := DB.AutoMigrate(&models.Stock{},
		&models.Stockelement{},
		&models.User{},
		&models.Portfolio{})

	if errs != nil {
		log.Fatalf("Failed to migrate: %v", err)
	} else {
		log.Println("Migration successful")
	}
}
