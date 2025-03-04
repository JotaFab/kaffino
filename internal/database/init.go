package database

import (
	"context"
	"fmt"
	"log"

	"kaffino/internal/coffeeshop"

	"github.com/google/uuid"
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

	// Populate the products table with example products
	err = s.populateproductsTable()
	if err != nil {
		log.Println("Error populating products table:", err)
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
						
			CREATE TABLE IF NOT EXISTS products (
				id VARCHAR(36) PRIMARY KEY,
				code VARCHAR(255) UNIQUE,
				images TEXT,             -- Comma-separated list of image URLs
				title VARCHAR(255) NOT NULL,
				description TEXT,
				created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
				updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			);
				
			CREATE TABLE IF NOT EXISTS inventory (
				id VARCHAR(36) PRIMARY KEY,
				product_id VARCHAR(36) NOT NULL,
				stock INTEGER NOT NULL DEFAULT 0,
				sizes TEXT,               -- Comma-separated list of available sizes
				price DECIMAL(10,2) NOT NULL DEFAULT 0.0,
				created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
				updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
				FOREIGN KEY (product_id) REFERENCES products(id)
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
				FOREIGN KEY (product_id) REFERENCES products(id) -- Foreign key to the products table
			);

			CREATE TABLE IF NOT EXISTS reviews (
				id VARCHAR(36) PRIMARY KEY,
				product_id VARCHAR(36) NOT NULL,
				user_id VARCHAR(36) NOT NULL,
				rating INTEGER NOT NULL,
				comment TEXT,
				created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
				updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
				FOREIGN KEY (product_id) REFERENCES products(id),
				FOREIGN KEY (user_id) REFERENCES users(id)
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

// A function that populates the products table based on our product data.
func (s *service) populateproductsTable() error {
	// Check if the products table is already populated
	var tablePopulated bool
	err := s.db.QueryRow(`
		SELECT EXISTS (
			SELECT 1
			FROM products
		)
	`).Scan(&tablePopulated)
	if err != nil {
		log.Println("Error checking if products table is populated:", err)
		return err
	}

	// If the products table is not populated, insert the product data
	if !tablePopulated {
		log.Println("products table is not populated, inserting product data...")

		// Define products
		products := []coffeeshop.Product{
			{
				Code:        "BEAN001",
				Images:      []string{"whole_bean1.jpg", "whole_bean2.jpg"},
				Title:       "Peruvian Whole Bean Coffee",
				Description: "High-altitude Arabica beans, perfect for home roasting.",
				Tags:        []string{"coffee", "beans", "whole"},
			},
			{
				Code:        "DRINK001",
				Images:      []string{"cappuccino1.jpg", "cappuccino2.jpg"},
				Title:       "Classic Cappuccino",
				Description: "Espresso with steamed milk and foamed milk.",
				Tags:        []string{"coffee", "cappuccino", "classic"},
			},
			{
				Code:        "BLEND002",
				Images:      []string{"signature_blend1.jpg", "signature_blend2.jpg"},
				Title:       "Kaffino Signature Blend",
				Description: "A unique blend of Peruvian and Ethiopian beans.",
				Tags:        []string{"coffee", "blend", "signature"},
			},
			{
				Code:        "ACC001",
				Images:      []string{"french_press1.jpg", "french_press2.jpg"},
				Title:       "French Press",
				Description: "Classic coffee brewing device.",
				Tags:        []string{"coffee", "french press", "accessories"},
			},
			{
				Code:        "GRIND001",
				Images:      []string{"coffee_grinder1.jpg", "coffee_grinder2.jpg"},
				Title:       "Coffee Grinder",
				Description: "Electric coffee grinder for home use.",
				Tags:        []string{"coffee", "grinder", "accessories"},
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

			// Create inventory for the product
			inventory := coffeeshop.Inventory{
				ProductID: product.ID,
				Stock:     100,    // Example stock
				Size:      "12oz", // Example size
				Price:     15.00,  // Example price
			}

			inventoryID := uuid.New().String()
			_, err = s.db.Exec(`
				INSERT INTO inventory (id, product_id, stock, sizes, price)
				VALUES (?, ?, ?, ?, ?)
			`, inventoryID, inventory.ProductID, inventory.Stock, inventory.Size, inventory.Price)

			if err != nil {
				log.Println("Error inserting inventory data:", err)
				return err
			}
			fmt.Printf("Inserted inventory for product %s\n", product.Title)
		}

		log.Println("products table populated successfully.")
	} else {
		log.Println("products table is already populated.")
	}
	return nil
}
