package main

import (
	"fmt"
	"log"
	"os"
	"farmcaptain/controllers"
	"farmcaptain/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

func init() {
	// Load environment variables
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to the PostgreSQL database
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable", 
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"))
	db, err = gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatal("failed to connect to the database")
	}

	// Migrate database models
	db.AutoMigrate(&controllers.Crop{}, &controllers.Farmer{})
}

func main() {
	// Initialize the Gin router
	r := gin.Default()

	// Middleware: JWT Authentication
	r.Use(middlewares.JWTMiddleware())

	// Crop-related routes
	r.POST("/addCrop", controllers.AddCrop)
	r.GET("/getCrop/:id", controllers.GetCrop)
	r.GET("/aiAdvice/:cropId", controllers.GetAIAdvice)

	// Farmer login and authentication
	r.POST("/login", controllers.Login)

	// Start the server
	r.Run(":8080")
}
