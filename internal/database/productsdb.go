package database

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"kaffino/internal/coffeeshop"
)

// CreateProduct creates a new product in the database.
func (s *service) CreateProduct(ctx context.Context, product *coffeeshop.Product) error {
	product.NewProduct()
	imagesJSON, err := json.Marshal(product.Images)
	if err != nil {
		return fmt.Errorf("error marshaling images: %w", err)
	}

	tagsJSON, err := json.Marshal(product.Tags)
	if err != nil {
		return fmt.Errorf("error marshaling tags: %w", err)
	}

	query := `
		INSERT INTO products (id, code, images, title, description, tags, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err = s.db.ExecContext(ctx, query,
		product.ID, product.Code, imagesJSON, product.Title, product.Description,
		tagsJSON, product.CreatedAt, product.UpdatedAt)

	if err != nil {
		return fmt.Errorf("error creating product: %w", err)
	}

	log.Println("Product created successfully")
	return nil
}

// GetProduct retrieves a product from the database by ID.
func (s *service) GetProduct(ctx context.Context, id string) (*coffeeshop.Product, error) {
	query := `
		SELECT id, code, images, title, description, tags, created_at, updated_at
		FROM products
		WHERE id = ?
	`

	row := s.db.QueryRowContext(ctx, query, id)

	product := &coffeeshop.Product{}
	var imagesJSON, tagsJSON []byte

	err := row.Scan(&product.ID, &product.Code, &imagesJSON, &product.Title, &product.Description,
		&tagsJSON, &product.CreatedAt, &product.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product not found: %w", err)
		}
		return nil, fmt.Errorf("error getting product: %w", err)
	}

	if err := json.Unmarshal(imagesJSON, &product.Images); err != nil {
		return nil, fmt.Errorf("error unmarshaling images: %w", err)
	}

	if err := json.Unmarshal(tagsJSON, &product.Tags); err != nil {
		return nil, fmt.Errorf("error unmarshaling tags: %w", err)
	}

	log.Println("Product retrieved successfully")
	return product, nil
}

// ListProducts retrieves all products from the database.
func (s *service) ListProducts(ctx context.Context) ([]*coffeeshop.Product, error) {
	query := `
		SELECT id, code, images, title, description, tags, created_at, updated_at
		FROM products
		LIMIT 10
	`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error listing products: %w", err)
	}
	defer rows.Close()

	var products []*coffeeshop.Product
	for rows.Next() {
		product := &coffeeshop.Product{}
		var imagesJSON, tagsJSON []byte

		err := rows.Scan(&product.ID, &product.Code, &imagesJSON, &product.Title, &product.Description,
			&tagsJSON, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("error scanning product: %w", err)
		}

		if err := json.Unmarshal(imagesJSON, &product.Images); err != nil {
			return nil, fmt.Errorf("error unmarshaling images: %w", err)
		}

		if err := json.Unmarshal(tagsJSON, &product.Tags); err != nil {
			return nil, fmt.Errorf("error unmarshaling tags: %w", err)
		}

		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating products: %w", err)
	}

	return products, nil
}

// UpdateProduct updates a product in the database.
func (s *service) UpdateProduct(ctx context.Context, product *coffeeshop.Product) error {
	imagesJSON, err := json.Marshal(product.Images)
	if err != nil {
		return fmt.Errorf("error marshaling images: %w", err)
	}

	tagsJSON, err := json.Marshal(product.Tags)
	if err != nil {
		return fmt.Errorf("error marshaling tags: %w", err)
	}

	query := `
		UPDATE products
		SET code = ?, images = ?, title = ?, description = ?, tags = ?, updated_at = ?
		WHERE id = ?
	`

	_, err = s.db.ExecContext(ctx, query,
		product.Code, imagesJSON, product.Title, product.Description,
		tagsJSON, product.UpdatedAt,
		product.ID)

	if err != nil {
		return fmt.Errorf("error updating product: %w", err)
	}

	return nil
}

// DeleteProduct deletes a product from the database by ID.
func (s *service) DeleteProduct(ctx context.Context, id string) error {
	query := `
		DELETE FROM products
		WHERE id = ?
	`

	_, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("error deleting product: %w", err)
	}

	return nil
}
