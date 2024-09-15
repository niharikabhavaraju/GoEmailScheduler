package config

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var db *gorm.DB

// Connect initializes the database connection
func Connect() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get the database connection string from environment variable
	connectionString := os.Getenv("DB_CONNECTION_STRING")

	var err error
	db, err = gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	log.Println("Database connection established")
}

// GetDB returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
