package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jayeshjadhav/mobile-store/config"
	"github.com/jayeshjadhav/mobile-store/middleware"
	"github.com/jayeshjadhav/mobile-store/models"
	"github.com/jayeshjadhav/mobile-store/routes"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r := gin.Default()
	r.Use(middleware.CORSMiddleware()) //enable CORS
	// Connect to DB
	config.ConnectDatabase()

	// Auto migrate models
	config.DB.AutoMigrate(&models.User{}, &models.Product{})

	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{"message": "API running!"})
	// })
	routes.RegisterRoutes(r)
	fmt.Println("JWT Secret:", os.Getenv("JWT_SECRET"))

	r.Run() // default on localhost:8080
}
