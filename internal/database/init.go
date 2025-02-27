package database

import (
	"encoding/json"
	"log"

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

		// SQL statement to insert products into the inventory table
		insertTablesSQL := `
			INSERT INTO inventory (id, code, images, title, description, long_description, discount, reviews, map_size_price, schedules, tags, stock_quantity, sizes)
			VALUES
				($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
		`

		// Generate UUIDs for the products
		id1 := uuid.New().String()
		id2 := uuid.New().String()
		id3 := uuid.New().String()

		// Define size-price maps for the products
		mapSizePrice1 := map[string]float64{"Half Bag (6oz)": 9.00, "Full Bag (12oz)": 18.00}
		mapSizePrice2 := map[string]float64{"Small": 12.00, "Medium": 22.00, "Large": 30.00}
		mapSizePrice3 := map[string]float64{"12oz": 15.00, "1lb": 25.00}

		// Marshal the maps to JSON strings
		mapSizePrice1JSON, _ := json.Marshal(mapSizePrice1)
		mapSizePrice2JSON, _ := json.Marshal(mapSizePrice2)
		mapSizePrice3JSON, _ := json.Marshal(mapSizePrice3)

		// Define sizes for the products
		sizes1 := "Half Bag (6oz),Full Bag (12oz)"
		sizes2 := "Small,Medium,Large"
		sizes3 := "12oz,1lb"

		// Execute the insert tables SQL statement
		_, err = s.db.Exec(insertTablesSQL,
			id1, "BEAN001", "whole_bean.jpg", "Peruvian Whole Bean Coffee", "High-altitude Arabica beans, perfect for home roasting.", "Experience the rich and complex flavors of our Peruvian Whole Bean Coffee. Sourced from the high-altitude regions of Peru, these Arabica beans are perfect for home roasting, allowing you to customize your coffee experience to your exact preferences.", 0.00, "", mapSizePrice1JSON, "", "coffee,beans,whole", 100, sizes1)

		if err != nil {
			log.Println("Error inserting product data:", err)
			return err
		}

		_, err = s.db.Exec(insertTablesSQL,
			id2, "DRINK001", "cappuccino.jpg", "Classic Cappuccino", "Espresso with steamed milk and foamed milk.", "A perfectly balanced cappuccino with rich espresso, steamed milk, and a delicate layer of foamed milk. A classic choice for any coffee lover.", 2.50, "", mapSizePrice2JSON, "", "coffee,cappuccino,classic", 50, sizes2)

		if err != nil {
			log.Println("Error inserting product data:", err)
			return err
		}

		_, err = s.db.Exec(insertTablesSQL,
			id3, "BLEND002", "signature_blend.jpg", "Kaffino Signature Blend", "A unique blend of Peruvian and Ethiopian beans.", "Our signature blend combines the best of Peruvian and Ethiopian beans, creating a harmonious balance of flavors with notes of chocolate and citrus. Perfect for any time of day.", 5.00, "", mapSizePrice3JSON, "", "coffee,blend,signature", 75, sizes3)

		if err != nil {
			log.Println("Error inserting product data:", err)
			return err
		}

		log.Println("Inventory table populated successfully.")
	} else {
		log.Println("Inventory table is already populated.")
	}
	return nil
}
