package database

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/yourusername/product-service/internal/models"
)

// ProductRepository defines the interface for product database operations
type ProductRepository interface {
	GetProducts(ctx context.Context) ([]models.Product, error)
	GetProductByID(ctx context.Context, id string) (models.Product, error)
	CreateProduct(ctx context.Context, product models.Product) (models.Product, error)
	UpdateProduct(ctx context.Context, product models.Product) error
	DeleteProduct(ctx context.Context, id string) error
	CheckProductAvailability(ctx context.Context, id string) (models.ProductAvailability, error)
}

// InMemoryRepository implements ProductRepository using in-memory storage
type InMemoryRepository struct {
	products map[string]models.Product
	mutex    sync.RWMutex
}

// NewInMemoryRepository creates a new in-memory repository
func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		products: make(map[string]models.Product),
	}
}

// GetProducts retrieves all products from the in-memory store
func (r *InMemoryRepository) GetProducts(ctx context.Context) ([]models.Product, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	
	products := make([]models.Product, 0, len(r.products))
	for _, product := range r.products {
		products = append(products, product)
	}
	
	return products, nil
}

// GetProductByID retrieves a product by its ID
func (r *InMemoryRepository) GetProductByID(ctx context.Context, id string) (models.Product, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	
	product, exists := r.products[id]
	if !exists {
		return models.Product{}, errors.New("product not found")
	}
	
	return product, nil
}

// CreateProduct creates a new product in the in-memory store
func (r *InMemoryRepository) CreateProduct(ctx context.Context, product models.Product) (models.Product, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	
	// Generate a UUID if not provided
	if product.ID == "" {
		product.ID = uuid.New().String()
	}
	
	// Check if product with the same ID already exists
	if _, exists := r.products[product.ID]; exists {
		return models.Product{}, errors.New("product with this ID already exists")
	}
	
	// Set timestamps
	now := time.Now()
	product.CreatedAt = now
	product.UpdatedAt = now
	
	// Store the product
	r.products[product.ID] = product
	
	return product, nil
}

// UpdateProduct updates an existing product
func (r *InMemoryRepository) UpdateProduct(ctx context.Context, product models.Product) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	
	// Check if product exists
	_, exists := r.products[product.ID]
	if !exists {
		return errors.New("product not found")
	}
	
	// Update timestamp
	product.UpdatedAt = time.Now()
	
	// Update the product
	r.products[product.ID] = product
	
	return nil
}

// DeleteProduct deletes a product from the in-memory store
func (r *InMemoryRepository) DeleteProduct(ctx context.Context, id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	
	// Check if product exists
	if _, exists := r.products[id]; !exists {
		return errors.New("product not found")
	}
	
	// Delete the product
	delete(r.products, id)
	
	return nil
}

// CheckProductAvailability checks if a product is available
func (r *InMemoryRepository) CheckProductAvailability(ctx context.Context, id string) (models.ProductAvailability, error) {
	product, err := r.GetProductByID(ctx, id)
	if err != nil {
		return models.ProductAvailability{}, err
	}
	
	availability := models.ProductAvailability{
		ProductID:      product.ID,
		InventoryCount: product.InventoryCount,
		IsAvailable:    product.InventoryCount > 0,
	}
	
	return availability, nil
}

// SampleProduct creates a sample product for testing
func SampleProduct(name, description string, price float64, inventory int) models.Product {
	return models.Product{
		ID:             uuid.New().String(),
		Name:           name,
		Description:    description,
		Price:          price,
		InventoryCount: inventory,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
}
