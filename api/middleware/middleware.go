package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger is a middleware function that logs the request method, path and how long it took to process
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		
		// Process request
		c.Next()
		
		// Log request details
		latency := time.Since(start)
		log.Printf("| %3d | %13v | %s | %s |",
			c.Writer.Status(),
			latency,
			c.Request.Method,
			c.Request.URL.Path,
		)
	}
}

// HealthCheck provides a simple health check endpoint
func HealthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "UP",
			"service": "product-service",
			"time": time.Now().Format(time.RFC3339),
		})
	}
}