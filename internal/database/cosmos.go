package database

import (
	"context"
	"errors"
	"github.com/yourusername/product-service/internal/models"
)

// CosmosDBRepository is a placeholder for a Cosmos DB implementation
// Note: This is just a placeholder and doesn't actually connect to Cosmos DB
// For local testing, use the InMemoryRepository instead
type CosmosDBRepository struct{}

// NewCosmosDBRepository creates a new repository connected to Cosmos DB
// Note: This is just a placeholder and will return an error
func NewCosmosDBRepository(connectionString, databaseName, containerName string) (*CosmosDBRepository, error) {
	return nil, errors.New("Cosmos DB repository is not available in this build - use NewInMemoryRepository() instead")
}

// The following methods are placeholders and will all return errors
// They are included to satisfy the ProductRepository interface

func (r *CosmosDBRepository) GetProducts(ctx context.Context) ([]models.Product, error) {
	return nil, errors.New("not implemented")
}

func (r *CosmosDBRepository) GetProductByID(ctx context.Context, id string) (models.Product, error) {
	return models.Product{}, errors.New("not implemented")
}

func (r *CosmosDBRepository) CreateProduct(ctx context.Context, product models.Product) (models.Product, error) {
	return models.Product{}, errors.New("not implemented")
}

func (r *CosmosDBRepository) UpdateProduct(ctx context.Context, product models.Product) error {
	return errors.New("not implemented")
}

func (r *CosmosDBRepository) DeleteProduct(ctx context.Context, id string) error {
	return errors.New("not implemented")
}

func (r *CosmosDBRepository) CheckProductAvailability(ctx context.Context, id string) (models.ProductAvailability, error) {
	return models.ProductAvailability{}, errors.New("not implemented")
}
