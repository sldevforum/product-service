package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yourusername/product-service/internal/database"
)

func TestGetGinEngine(t *testing.T) {
	repo := database.NewInMemoryRepository()
	router := GetGinEngine(repo)

	// Test health check endpoint
	t.Run("HealthCheck", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/health", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "UP")
	})

	// Test not found route
	t.Run("NotFound", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/not-found", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})

}

// test one to 3 