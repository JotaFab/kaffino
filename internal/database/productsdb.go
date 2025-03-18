package database

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

// Example CreateProduct using sqlc generated code
func (s *service) CreateProduct(ctx context.Context, product *Product) error {
	params := CreateProductParams{
		ID:          uuid.New().String(),
		Code:        product.Code,
		Images:      product.Images,
		Title:       product.Title,
		Description: product.Description,
		CreatedAt:   sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt:   sql.NullTime{Time: time.Now(), Valid: true},
	}

	err := s.q.CreateProduct(ctx, params)
	if err != nil {
		return fmt.Errorf("error creating product: %w", err)
	}

	log.Println("Product created successfully")
	return nil
}

// Example GetProduct using sqlc generated code
func (s *service) GetProduct(ctx context.Context, id string) (*Product, error) {
	productRow, err := s.q.GetProduct(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product not found: %w", err)
		}
		return nil, fmt.Errorf("error getting product: %w", err)
	}

	product := &Product{
		ID:          productRow.ID,
		Code:        productRow.Code,
		Images:      productRow.Images,
		Title:       productRow.Title,
		Description: productRow.Description,
	}

	log.Println("Product retrieved successfully")
	return product, nil
}

// ListProducts retrieves all products from the database.
func (s *service) ListProducts(ctx context.Context) ([]*Product, error) {
	query := `
		SELECT id, images, title, description, created_at, updated_at
		FROM products
		LIMIT 10
	`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error listing products: %w", err)
	}
	defer rows.Close()

	var products []*Product
	for rows.Next() {
		product := &Product{}

		err := rows.Scan(&product.ID, &product.Images, &product.Title, &product.Description,
			&product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("error scanning product: %w", err)
		}

		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating products: %w", err)
	}

	return products, nil
}

// UpdateProduct updates a product in the database.
func (s *service) UpdateProduct(ctx context.Context, product *Product) error {
	imagesJSON, err := json.Marshal(product.Images)
	if err != nil {
		return fmt.Errorf("error marshaling images: %w", err)
	}

	query := `
		UPDATE products
		SET code = ?, images = ?, title = ?, description = ?,  updated_at = ?
		WHERE id = ?
	`

	_, err = s.db.ExecContext(ctx, query,
		product.Code, imagesJSON, product.Title, product.Description, product.UpdatedAt,
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
