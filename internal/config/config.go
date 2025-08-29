package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds all configuration for the service
type Config struct {
	ServerPort   int
	CosmosDBURI  string
	DatabaseName string
	ContainerName string
	Environment  string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	// Load .env file if it exists
	godotenv.Load()
	
	// Set default values
	config := &Config{
		ServerPort:   8080,
		Environment:  "development",
		DatabaseName: "product-db",
		ContainerName: "products",
	}
	
	// Override with environment variables if set
	if port := os.Getenv("SERVER_PORT"); port != "" {
		if p, err := strconv.Atoi(port); err == nil {
			config.ServerPort = p
		}
	}
	
	if uri := os.Getenv("COSMOS_DB_URI"); uri != "" {
		config.CosmosDBURI = uri
	} else {
		// In a production environment, we would return an error here
		// For development, we'll just log a warning
		fmt.Println("Warning: COSMOS_DB_URI not set")
	}
	
	if dbName := os.Getenv("COSMOS_DB_NAME"); dbName != "" {
		config.DatabaseName = dbName
	}
	
	if containerName := os.Getenv("COSMOS_CONTAINER_NAME"); containerName != "" {
		config.ContainerName = containerName
	}
	
	if env := os.Getenv("ENVIRONMENT"); env != "" {
		config.Environment = env
	}
	
	return config, nil
}