package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/yourusername/product-service/api/handlers"
	"github.com/yourusername/product-service/api/middleware"
	_ "github.com/yourusername/product-service/docs"
	"github.com/yourusername/product-service/internal/config"
	"github.com/yourusername/product-service/internal/database"
)

// @title Product Service API
// @version 1.0
// @description API for managing products
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.example.com/support
// @contact.email support@example.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	
	// Initialize in-memory repository for testing
	repo := database.NewInMemoryRepository()
	
	// Add some sample products for testing
	ctx := context.Background()
	repo.CreateProduct(ctx, database.SampleProduct("Laptop", "High-performance laptop", 1299.99, 10))
	repo.CreateProduct(ctx, database.SampleProduct("Smartphone", "Latest smartphone model", 899.99, 15))
	repo.CreateProduct(ctx, database.SampleProduct("Headphones", "Noise-cancelling headphones", 249.99, 20))
	
	// Initialize handlers
	productHandler := handlers.NewProductHandler(repo)
	
	// Set up Gin router
	router := gin.Default()
	
	// Add middleware
	router.Use(middleware.Logger())
	
	// Health check endpoint
	router.GET("/health", middleware.HealthCheck())
	
	// API routes
	api := router.Group("/api")
	{
		products := api.Group("/products")
		{
			products.GET("", productHandler.GetProducts)
			products.GET("/:id", productHandler.GetProductByID)
			products.POST("", productHandler.CreateProduct)
			products.PUT("/:id", productHandler.UpdateProduct)
			products.DELETE("/:id", productHandler.DeleteProduct)
			products.GET("/:id/availability", productHandler.CheckProductAvailability)
		}
	}
	
	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	// Start server
	serverAddr := fmt.Sprintf(":%d", cfg.ServerPort)
	log.Printf("Starting server on %s", serverAddr)
	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// GetGinEngine configures and returns a new Gin engine
func GetGinEngine(repo database.ProductRepository) *gin.Engine {
	// Initialize handlers
	productHandler := handlers.NewProductHandler(repo)
	
	// Set up Gin router
	router := gin.Default()
	
	// Add middleware
	router.Use(middleware.Logger())
	
	// Health check endpoint
	router.GET("/health", middleware.HealthCheck())
	
	// API routes
	api := router.Group("/api")
	{
		products := api.Group("/products")
		{
			products.GET("", productHandler.GetProducts)
			products.GET("/:id", productHandler.GetProductByID)
			products.POST("", productHandler.CreateProduct)
			products.PUT("/:id", productHandler.UpdateProduct)
			products.DELETE("/:id", productHandler.DeleteProduct)
			products.GET("/:id/availability", productHandler.CheckProductAvailability)
		}
	}
	
	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	return router
}