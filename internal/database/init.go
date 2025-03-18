package database

import (
	"context"
	"fmt"
	"log"

	"database/sql"

	"github.com/google/uuid"
)

// dbInit checks if the products table exists and creates it if it doesn't.
// It also populates the table with some example products.
func (s *service) DbInit() error {

	// Populate the products table with example products
	err := s.populateproductsTable()
	if err != nil {
		log.Println("Error populating products table:", err)
		return err
	}

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
		products := []Product{
			{
				Code: sql.NullString{
					String: "BEAN001",
					Valid:  true, // Set Valid to true since you have a value
				},
				Images: sql.NullString{
					String: "whole_bean1.jpg, whole_bean2.jpg",
					Valid:  true,
				},
				Title: "Peruvian Whole Bean Coffee",
				Description: sql.NullString{
					String: "High-altitude Arabica beans, perfect for home roasting.",
					Valid:  true,
				},
			},
			{
				Code: sql.NullString{
					String: "DRINK001",
					Valid:  true,
				},
				Images: sql.NullString{
					String: "cappuccino1.jpg, cappuccino2.jpg",
					Valid:  true,
				},
				Title: "Classic Cappuccino",
				Description: sql.NullString{
					String: "Espresso with steamed milk and foamed milk.",
					Valid:  true},
			},
			{
				Code: sql.NullString{
					String: "BLEND002",
					Valid:  true,
				},
				Images: sql.NullString{
					String: "signature_blend1.jpg, signature_blend2.jpg",
					Valid:  true,
				},
				Title: "Kaffino Signature Blend",
				Description: sql.NullString{
					String: "A unique blend of Peruvian and Ethiopian beans.",
					Valid:  true},
			},
			{
				Code: sql.NullString{
					String: "ACC001",
					Valid:  true,
				},
				Images: sql.NullString{
					String: "french_press1.jpg, french_press2.jpg",
					Valid:  true,
				},
				Title: "French Press",
				Description: sql.NullString{
					String: "Classic coffee brewing device.",
					Valid:  true},
			},
			{
				Code: sql.NullString{
					String: "GRIND001",
					Valid:  true,
				},
				Images: sql.NullString{
					String: "coffee_grinder1.jpg, coffee_grinder2.jpg",
					Valid:  true,
				},
				Title: "Coffee Grinder",
				Description: sql.NullString{
					String: "Electric coffee grinder for home use.",
					Valid:  true,
				},
			},
		}

		// Insert products using CreateProduct method
		for _, product := range products {
			err := s.CreateProduct(context.Background(), &product)
			if err != nil {
				log.Println("Error inserting product data:", err)
				return err
			}

			// Create inventory for the product
			inventory := Inventory{
				ProductID: product.ID,
				Stock:     100,   // Example stock
				Price:     15.00, // Example price
			}

			inventoryID := uuid.New().String()
			_, err = s.db.Exec(`
				INSERT INTO inventory (id, product_id, stock, sizes, price)
				VALUES (?, ?, ?, ?, ?)
			`, inventoryID, inventory.ProductID, inventory.Stock, inventory.Sizes, inventory.Price)

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
