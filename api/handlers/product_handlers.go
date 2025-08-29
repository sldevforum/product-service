package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/product-service/internal/database"
	"github.com/yourusername/product-service/internal/models"
)

// ProductHandler handles HTTP requests related to products
type ProductHandler struct {
	repo database.ProductRepository
}

// NewProductHandler creates a new product handler
func NewProductHandler(repo database.ProductRepository) *ProductHandler {
	return &ProductHandler{
		repo: repo,
	}
}

// GetProducts godoc
// @Summary Get all products
// @Description Get a list of all products
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {array} models.Product
// @Failure 500 {object} map[string]interface{}
// @Router /api/products [get]
func (h *ProductHandler) GetProducts(c *gin.Context) {
	ctx := context.Background()
	products, err := h.repo.GetProducts(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, products)
}

// GetProductByID godoc
// @Summary Get product by ID
// @Description Get a product by its ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} models.Product
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/products/{id} [get]
func (h *ProductHandler) GetProductByID(c *gin.Context) {
	id := c.Param("id")
	
	ctx := context.Background()
	product, err := h.repo.GetProductByID(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	
	c.JSON(http.StatusOK, product)
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product with the provided information
// @Tags products
// @Accept json
// @Produce json
// @Param product body models.Product true "Product information"
// @Success 201 {object} models.Product
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/products [post]
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	ctx := context.Background()
	createdProduct, err := h.repo.CreateProduct(ctx, product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusCreated, createdProduct)
}

// UpdateProduct godoc
// @Summary Update a product
// @Description Update a product with the provided information
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param product body models.Product true "Product information"
// @Success 200 {object} models.Product
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/products/{id} [put]
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// Ensure ID in URL matches ID in body
	product.ID = id
	
	ctx := context.Background()
	
	// Check if product exists
	_, err := h.repo.GetProductByID(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	
	// Update product
	if err := h.repo.UpdateProduct(ctx, product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, product)
}

// DeleteProduct godoc
// @Summary Delete a product
// @Description Delete a product by its ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 204 {object} nil
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/products/{id} [delete]
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	
	ctx := context.Background()
	
	// Check if product exists
	_, err := h.repo.GetProductByID(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	
	// Delete product
	if err := h.repo.DeleteProduct(ctx, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.Status(http.StatusNoContent)
}

// CheckProductAvailability godoc
// @Summary Check product availability
// @Description Check if a product is available (has inventory)
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} models.ProductAvailability
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/products/{id}/availability [get]
func (h *ProductHandler) CheckProductAvailability(c *gin.Context) {
	id := c.Param("id")
	
	ctx := context.Background()
	availability, err := h.repo.CheckProductAvailability(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	
	c.JSON(http.StatusOK, availability)
}