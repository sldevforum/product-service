package models

import (
	"time"
)

// Product represents the product entity in the system
// @Description Product information
type Product struct {
	ID           string    `json:"id"`
	Name         string    `json:"name" binding:"required"`
	Description  string    `json:"description" binding:"required"`
	Price        float64   `json:"price" binding:"required,gt=0"`
	InventoryCount int     `json:"inventoryCount" binding:"required,gte=0"`
	CreatedAt    time.Time `json:"createdAt,omitempty"`
	UpdatedAt    time.Time `json:"updatedAt,omitempty"`
}

// ProductAvailability represents the product availability information
// @Description Product availability information
type ProductAvailability struct {
	ProductID      string `json:"productId"`
	IsAvailable    bool   `json:"isAvailable"`
	InventoryCount int    `json:"inventoryCount"`
}