
package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/yourusername/product-service/internal/database"
	"github.com/yourusername/product-service/internal/models"
)

func setupRouter() (*gin.Engine, *database.InMemoryRepository) {
	repo := database.NewInMemoryRepository()
	productHandler := NewProductHandler(repo)

	router := gin.Default()
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
	return router, repo
}

func TestProductHandlers(t *testing.T) {
	router, repo := setupRouter()

	// Test CreateProduct
	t.Run("CreateProduct", func(t *testing.T) {
		product := models.Product{Name: "Test Product", Description: "Test Description", Price: 10.0, InventoryCount: 100}
		jsonValue, _ := json.Marshal(product)
		req, _ := http.NewRequest(http.MethodPost, "/api/products", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
	})

	// Test GetProducts
	t.Run("GetProducts", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/api/products", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	// Test GetProductByID
	t.Run("GetProductByID", func(t *testing.T) {
		products, _ := repo.GetProducts(nil)
		productID := products[0].ID

		req, _ := http.NewRequest(http.MethodGet, "/api/products/"+productID, nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	// Test UpdateProduct
	t.Run("UpdateProduct", func(t *testing.T) {
		products, _ := repo.GetProducts(nil)
		product := products[0]
		product.Name = "Updated Product"

		jsonValue, _ := json.Marshal(product)
		req, _ := http.NewRequest(http.MethodPut, "/api/products/"+product.ID, bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	// Test DeleteProduct
	t.Run("DeleteProduct", func(t *testing.T) {
		products, _ := repo.GetProducts(nil)
		productID := products[0].ID

		req, _ := http.NewRequest(http.MethodDelete, "/api/products/"+productID, nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNoContent, w.Code)
	})

	// Test CheckProductAvailability
	t.Run("CheckProductAvailability", func(t *testing.T) {
		product := models.Product{Name: "Available Product", Description: "Test Description", Price: 10.0, InventoryCount: 10}
		createdProduct, _ := repo.CreateProduct(nil, product)

		req, _ := http.NewRequest(http.MethodGet, "/api/products/"+createdProduct.ID+"/availability", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "true")
	})
}
