package database

import (
	"fmt"
	"log"
	"os"
	"post_management/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func InitDB() {
	var err error

	// Load environment variables from .env file
	_ = godotenv.Load()

	// Load database config from environment variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Construct the DSN (Data Source Name)
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbname, password)

	// Connect to the database
	DB, err = gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
		return
	}

	// Automatically migrate the schema
	DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
	fmt.Println("Database connection established successfully")
}
