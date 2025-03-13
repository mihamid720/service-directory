package main

import (
	"log"
	"services-api/internal/database"
	"services-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database connection
	database.InitDB()

	// Create Gin router
	router := gin.Default()

	// Define routes
	router.GET("/services", handlers.ListServices)
	router.GET("/services/:id", handlers.GetService)

	// Start server
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
