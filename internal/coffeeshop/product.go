package coffeeshop

import (
	"time"

	"github.com/google/uuid"
)

// Product represents a product in the coffee shop.
type Product struct {
	ID                 string             `json:"id"`
	Code               string             `json:"code"`
	Images             []string           `json:"images"`
	Discount           float64            `json:"discount"`
	Title              string             `json:"title"`
	Description        string             `json:"description"`
	LongDescription    string             `json:"long_description"`
	DiscountPercentage float64            `json:"discount_percentage"`
	Reviews            []string           `json:"reviews"`
	MapSizePrice       map[string]float64 `json:"map_size_price"`
	Schedules          []string           `json:"shedules"`
	Tags               []string           `json:"tags"`
	CreatedAt          time.Time          `json:"created_at"`
	UpdatedAt          time.Time          `json:"updated_at"`
	StockQuantity      int                `json:"stock_quantity"`
	Sizes              []string           `json:"sizes"`
}

// NewProduct creates a new product and initializes its fields.
func NewProduct() *Product {
	return &Product{
		ID:        uuid.New().String(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// UpdateProduct updates the product time updatedAt.
func (p *Product) UpdateProduct() *Product {
	p.UpdatedAt = time.Now()
	return p
}

