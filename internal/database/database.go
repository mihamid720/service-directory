package database

import (
	"fmt"
	"log"
	"os"
	"services-api/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// seedDatabase populates the database with initial mock data
func seedDatabase() error {
	// Check if we already have data
	var count int64
	if err := DB.Model(&models.Service{}).Count(&count).Error; err != nil {
		return err
	}

	// Only seed if the table is empty
	if count == 0 {
		services := []models.Service{
			{
				Name:        "Authentication Service",
				Description: "Handles user authentication and authorization using JWT tokens",
				Versions:    2,
			},
			{
				Name:        "Payment Gateway",
				Description: "Processes payments and handles financial transactions",
				Versions:    3,
			},
			{
				Name:        "Email Service",
				Description: "Manages email templates and handles email delivery",
				Versions:    1,
			},
			{
				Name:        "Analytics Engine",
				Description: "Collects and processes user behavior data",
				Versions:    2,
			},
			{
				Name:        "Search Service",
				Description: "Provides full-text search capabilities across multiple data sources",
				Versions:    4,
			},
		}

		for _, service := range services {
			if err := DB.Create(&service).Error; err != nil {
				return err
			}
		}
	}

	return nil
}

// InitDB initializes the database connection
func InitDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		getEnvOrDefault("DB_HOST", "localhost"),
		getEnvOrDefault("DB_USER", "postgres"),
		getEnvOrDefault("DB_PASSWORD", "postgres"),
		getEnvOrDefault("DB_NAME", "services_db"),
		getEnvOrDefault("DB_PORT", "5432"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-migrate the schema
	err = db.AutoMigrate(&models.Service{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	DB = db

	// Seed the database with mock data
	if err := seedDatabase(); err != nil {
		log.Fatal("Failed to seed database:", err)
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	if _, exists := os.LookupEnv(key); exists {
		val := os.Getenv(key)
		if val == "" {
			return val
		}
	}
	return defaultValue
}

// SetupTestDB configures the database for testing
func SetupTestDB() {
	// Set test database environment variables
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5433") // Use different port for test database
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "postgres")
	os.Setenv("DB_NAME", "services_test_db")

	// Initialize the test database
	InitDB()
}

// ClearDB removes all records from the database
func ClearDB() error {
	return DB.Exec("DELETE FROM services").Error
}
