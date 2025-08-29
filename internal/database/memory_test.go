
package database

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yourusername/product-service/internal/models"
)

func TestInMemoryRepository(t *testing.T) {
	repo := NewInMemoryRepository()
	ctx := context.Background()

	// Test CreateProduct
	t.Run("CreateProduct", func(t *testing.T) {
		product := models.Product{Name: "Test Product", Description: "Test Description", Price: 10.0, InventoryCount: 100}
		createdProduct, err := repo.CreateProduct(ctx, product)

		assert.NoError(t, err)
		assert.NotEmpty(t, createdProduct.ID)
		assert.Equal(t, product.Name, createdProduct.Name)
	})

	// Test GetProducts
	t.Run("GetProducts", func(t *testing.T) {
		products, err := repo.GetProducts(ctx)
		assert.NoError(t, err)
		assert.Len(t, products, 1)
	})

	// Test GetProductByID
	t.Run("GetProductByID", func(t *testing.T) {
		products, _ := repo.GetProducts(ctx)
		productID := products[0].ID

		product, err := repo.GetProductByID(ctx, productID)
		assert.NoError(t, err)
		assert.Equal(t, productID, product.ID)
	})

	// Test UpdateProduct
	t.Run("UpdateProduct", func(t *testing.T) {
		products, _ := repo.GetProducts(ctx)
		product := products[0]
		product.Name = "Updated Product"

		err := repo.UpdateProduct(ctx, product)
		assert.NoError(t, err)

		updatedProduct, _ := repo.GetProductByID(ctx, product.ID)
		assert.Equal(t, "Updated Product", updatedProduct.Name)
	})

	// Test DeleteProduct
	t.Run("DeleteProduct", func(t *testing.T) {
		products, _ := repo.GetProducts(ctx)
		productID := products[0].ID

		err := repo.DeleteProduct(ctx, productID)
		assert.NoError(t, err)

		_, err = repo.GetProductByID(ctx, productID)
		assert.Error(t, err)
	})

	// Test CheckProductAvailability
	t.Run("CheckProductAvailability", func(t *testing.T) {
		product := models.Product{Name: "Available Product", Description: "Test Description", Price: 10.0, InventoryCount: 10}
		createdProduct, _ := repo.CreateProduct(ctx, product)

		availability, err := repo.CheckProductAvailability(ctx, createdProduct.ID)
		assert.NoError(t, err)
		assert.True(t, availability.IsAvailable)
		assert.Equal(t, 10, availability.InventoryCount)

		// Test for unavailable product
		unavailableProduct := models.Product{Name: "Unavailable Product", Description: "Test Description", Price: 10.0, InventoryCount: 0}
		createdUnavailableProduct, _ := repo.CreateProduct(ctx, unavailableProduct)

		availability, err = repo.CheckProductAvailability(ctx, createdUnavailableProduct.ID)
		assert.NoError(t, err)
		assert.False(t, availability.IsAvailable)
		assert.Equal(t, 0, availability.InventoryCount)
	})
}
