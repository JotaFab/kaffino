package database

import (
	"context"
	"log"

	"kaffino/internal/coffeeshop"
)

// dbInit checks if the products table exists and creates it if it doesn't.
// It also populates the table with some example products.
func (s *service) DbInit() error {
	// Create Kaffino tables
	err := s.createKaffinoTables()
	if err != nil {
		log.Println("Error creating tables:", err)
		return err
	}

	// Populate the inventory table with example products
	err = s.populateInventoryTable()
	if err != nil {
		log.Println("Error populating inventory table:", err)
		return err
	}

	return nil
}

// A function that creates the Kaffino tables.
func (s *service) createKaffinoTables() error {
	// SQL statement to create the products table
	createTableSQL := `
			CREATE TABLE IF NOT EXISTS users (
				id VARCHAR(36) PRIMARY KEY,  -- UUID for user identification
				email VARCHAR(255) UNIQUE NOT NULL, -- User's email address (unique)
				subscriber BOOLEAN DEFAULT FALSE, -- Indicates if the user is a subscriber
				username VARCHAR(255),       -- User's username
				created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
				updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
			);
						
			CREATE TABLE IF NOT EXISTS inventory (
				id VARCHAR(36) PRIMARY KEY,
				code VARCHAR(255) UNIQUE,
				images TEXT,             -- Comma-separated list of image URLs
				title VARCHAR(255) NOT NULL,
				description TEXT,
				long_description TEXT,
				discount DECIMAL(10, 2) DEFAULT 0.00,
				reviews TEXT,            -- Comma-separated list of review IDs (or review text)
				tags TEXT,               -- Comma-separated list of tags
				map_size_price TEXT,    -- JSON object for size-price mapping
				schedules TEXT,          -- Comma-separated list of schedules
				created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
				updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
				stock_quantity INTEGER DEFAULT 0,
				sizes TEXT               -- Comma-separated list of available sizes
			);

			CREATE TABLE IF NOT EXISTS orders (
				id VARCHAR(36) PRIMARY KEY,  -- UUID for order identification
				user_id VARCHAR(36) NOT NULL, -- UUID of the user who placed the order
				order_date DATETIME DEFAULT CURRENT_TIMESTAMP, -- Date and time the order was placed
				total_amount DECIMAL(10, 2) NOT NULL, -- Total amount of the order
				shipping_address TEXT,       -- Shipping address
				billing_address TEXT,        -- Billing address
				payment_method VARCHAR(255),  -- Payment method used (e.g., "Credit Card", "PayPal")
				order_status VARCHAR(255) DEFAULT 'Pending', -- Order status (e.g., "Pending", "Processing", "Shipped", "Delivered", "Cancelled")
				created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
				updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
				FOREIGN KEY (user_id) REFERENCES users(id) -- Foreign key to the users table
			);

			CREATE TABLE IF NOT EXISTS order_items (
				id VARCHAR(36) PRIMARY KEY,  -- UUID for order item identification
				order_id VARCHAR(36) NOT NULL, -- UUID of the order
				product_id VARCHAR(36) NOT NULL, -- UUID of the product
				quantity INTEGER NOT NULL,      -- Quantity of the product ordered
				price DECIMAL(10, 2) NOT NULL,   -- Price of the product at the time of the order
				created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
				updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
				FOREIGN KEY (order_id) REFERENCES orders(id), -- Foreign key to the orders table
				FOREIGN KEY (product_id) REFERENCES inventory(id) -- Foreign key to the inventory table
			);
		`

	// Execute the create table SQL statement
	_, err := s.db.Exec(createTableSQL)
	if err != nil {
		log.Println("Error creating tables:", err)
		return err
	}

	log.Println("Kaffino tables created successfully.")
	return nil
}

// A function that populates the inventory table based on our product data.
func (s *service) populateInventoryTable() error {
	// Check if the inventory table is already populated
	var tablePopulated bool
	err := s.db.QueryRow(`
		SELECT EXISTS (
			SELECT 1
			FROM inventory
		)
	`).Scan(&tablePopulated)
	if err != nil {
		log.Println("Error checking if inventory table is populated:", err)
		return err
	}

	// If the inventory table is not populated, insert the product data
	if !tablePopulated {
		log.Println("Inventory table is not populated, inserting product data...")

		// Define products
		products := []coffeeshop.Product{
			{
				Code:               "BEAN001",
				Images:             []string{"whole_bean1.jpg", "whole_bean2.jpg"},
				Discount:	 0.00,
				Title:              "Peruvian Whole Bean Coffee",
				Description:        "High-altitude Arabica beans, perfect for home roasting.",
				LongDescription:    "Experience the rich and complex flavors of our Peruvian Whole Bean Coffee. Sourced from the high-altitude regions of Peru, these Arabica beans are perfect for home roasting, allowing you to customize your coffee experience to your exact preferences.",
				Reviews:            []string{"Great taste!"},
				MapSizePrice:       map[string]float64{"Half Bag (6oz)": 9.00, "Full Bag (12oz)": 18.00},
				Schedules:          []string{"Morning", "Afternoon"},
				Tags:               []string{"coffee", "beans", "whole"},
				StockQuantity:      100,
				Sizes:              []string{"Half Bag (6oz)", "Full Bag (12oz)"},
			},
			{
				Code:               "DRINK001",
				Images:             []string{"cappuccino1.jpg", "cappuccino2.jpg"},
				Discount:           2.50,
				Title:              "Classic Cappuccino",
				Description:        "Espresso with steamed milk and foamed milk.",
				LongDescription:    "A perfectly balanced cappuccino with rich espresso, steamed milk, and a delicate layer of foamed milk. A classic choice for any coffee lover.",
				Reviews:            []string{"Perfect for a morning boost."},
				MapSizePrice:       map[string]float64{"Small": 12.00, "Medium": 22.00, "Large": 30.00},
				Schedules:          []string{"Anytime"},
				Tags:               []string{"coffee", "cappuccino", "classic"},
				StockQuantity:      50,
				Sizes:              []string{"Small", "Medium", "Large"},
			},
			{
				Code:               "BLEND002",
				Images:             []string{"signature_blend1.jpg", "signature_blend2.jpg"},
				Title:              "Kaffino Signature Blend",
				Description:        "A unique blend of Peruvian and Ethiopian beans.",
				LongDescription:    "Our signature blend combines the best of Peruvian and Ethiopian beans, creating a harmonious balance of flavors with notes of chocolate and citrus. Perfect for any time of day.",
				Discount:           5.00,
				Reviews:            []string{"My favorite blend."},
				MapSizePrice:       map[string]float64{"12oz": 15.00, "1lb": 25.00},
				Schedules:          []string{"Evening"},
				Tags:               []string{"coffee", "blend", "signature"},
				StockQuantity:      75,
				Sizes:              []string{"12oz", "1lb"},
			},
		}

		// Insert products using CreateProduct method
		for _, product := range products {
			product.NewProduct()
			err := s.CreateProduct(context.Background(), &product)
			if err != nil {
				log.Println("Error inserting product data:", err)
				return err
			}
		}

		log.Println("Inventory table populated successfully.")
	} else {
		log.Println("Inventory table is already populated.")
	}
	return nil
}
