package coffeeshop

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         string    `json:"id"`
	Email      string    `json:"email"`
	Subscriber bool      `json:"subscriber"`
	Username   string    `json:"username"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type CartItem struct {
	ItemID   string  `json:"id"`       // Unique ID for the cart item
	Quantity int     `json:"quantity"` // Quantity of the product in the cart
	Price    float64 `json:"price"`    // Price of the product at the time it was added to the cart
	Size     string  `json:"size"`     // Size of the product
	Schedule string  `json:"schedule"` // Schedule for the product

}

type Cart struct {
	ID        string     `json:"id"`
	UserID    string     `json:"user_id"`
	Items     []CartItem `json:"items"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// Product represents a product in the coffee shop.
type Product struct {
	ID              string             `json:"id"`
	Code            string             `json:"code"`
	Images          []string           `json:"images"`
	Discount        float64            `json:"discount"`
	Title           string             `json:"title"`
	Description     string             `json:"description"`
	LongDescription string             `json:"long_description"`
	Reviews         []string           `json:"reviews"`
	MapSizePrice    map[string]float64 `json:"map_size_price"`
	Schedules       []string           `json:"shedules"`
	Tags            []string           `json:"tags"`
	CreatedAt       time.Time          `json:"created_at"`
	UpdatedAt       time.Time          `json:"updated_at"`
	StockQuantity   int                `json:"stock_quantity"`
	Sizes           []string           `json:"sizes"`
}

// NewProduct creates a new product and initializes its fields.
func (p *Product) NewProduct() {


	p.ID = uuid.New().String()
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	return
}

// UpdateProduct updates the product time updatedAt.
func (p *Product) UpdateProduct() {
	p.UpdatedAt = time.Now()
	return 
}
