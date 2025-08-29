
package config

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	// Unset environment variables to ensure a clean test
	os.Unsetenv("SERVER_PORT")
	os.Unsetenv("COSMOS_DB_URI")
	os.Unsetenv("COSMOS_DB_NAME")
	os.Unsetenv("COSMOS_CONTAINER_NAME")
	os.Unsetenv("ENVIRONMENT")

	// Test case 1: Default values
	t.Run("DefaultValues", func(t *testing.T) {
		config, err := LoadConfig()
		assert.NoError(t, err)
		assert.Equal(t, 8080, config.ServerPort)
		assert.Equal(t, "development", config.Environment)
		assert.Equal(t, "product-db", config.DatabaseName)
		assert.Equal(t, "products", config.ContainerName)
		assert.Equal(t, "", config.CosmosDBURI)
	})

	// Test case 2: Environment variables override defaults
	t.Run("WithEnvironmentVariables", func(t *testing.T) {
		// Set environment variables
		os.Setenv("SERVER_PORT", "9090")
		os.Setenv("COSMOS_DB_URI", "test_uri")
		os.Setenv("COSMOS_DB_NAME", "test_db")
		os.Setenv("COSMOS_CONTAINER_NAME", "test_container")
		os.Setenv("ENVIRONMENT", "production")

		config, err := LoadConfig()
		assert.NoError(t, err)

		port, _ := strconv.Atoi(os.Getenv("SERVER_PORT"))
		assert.Equal(t, port, config.ServerPort)
		assert.Equal(t, "test_uri", config.CosmosDBURI)
		assert.Equal(t, "test_db", config.DatabaseName)
		assert.Equal(t, "test_container", config.ContainerName)
		assert.Equal(t, "production", config.Environment)

		// Unset environment variables after the test
		os.Unsetenv("SERVER_PORT")
		os.Unsetenv("COSMOS_DB_URI")
		os.Unsetenv("COSMOS_DB_NAME")
		os.Unsetenv("COSMOS_CONTAINER_NAME")
		os.Unsetenv("ENVIRONMENT")
	})
}
