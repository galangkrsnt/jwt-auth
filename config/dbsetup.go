package config

import (
	"fmt"
	"jwt-auth/model"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func AutoMigrateDataBase() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading env file")
	}

	dbDriver := os.Getenv("DB_DRIVER")
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		dbHost, dbUser, dbPassword, dbName, dbPort)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed Connect to Database")
	} else {
		fmt.Println("Connected to ", dbDriver)
	}

	DB.AutoMigrate(&model.User{})
}
