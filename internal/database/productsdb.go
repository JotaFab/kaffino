package database

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"kaffino/internal/coffeeshop"
)

// CreateProduct creates a new product in the database.
func (s *service) CreateProduct(ctx context.Context, product *coffeeshop.Product) error {
	imagesJSON, err := json.Marshal(product.Images)
	if err != nil {
		return fmt.Errorf("error marshaling images: %w", err)
	}

	reviewsJSON, err := json.Marshal(product.Reviews)
	if err != nil {
		return fmt.Errorf("error marshaling reviews: %w", err)
	}

	schedulesJSON, err := json.Marshal(product.Schedules)
	if err != nil {
		return fmt.Errorf("error marshaling schedules: %w", err)
	}

	tagsJSON, err := json.Marshal(product.Tags)
	if err != nil {
		return fmt.Errorf("error marshaling tags: %w", err)
	}

	mapSizePriceJSON, err := json.Marshal(product.MapSizePrice)
	if err != nil {
		return fmt.Errorf("error marshaling map_size_price: %w", err)
	}

	sizesJSON, err := json.Marshal(product.Sizes)
	if err != nil {
		return fmt.Errorf("error marshaling sizes: %w", err)
	}

	query := `
		INSERT INTO products (id, code, images, discount, title, description, long_description, discount_percentage, reviews, map_size_price, schedules, tags, created_at, updated_at, stock_quantity, sizes)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err = s.db.ExecContext(ctx, query,
		product.ID, product.Code, imagesJSON, product.Discount, product.Title, product.Description,
		product.LongDescription, product.DiscountPercentage, reviewsJSON, mapSizePriceJSON,
		schedulesJSON, tagsJSON, product.CreatedAt, product.UpdatedAt, product.StockQuantity, sizesJSON)

	if err != nil {
		return fmt.Errorf("error creating product: %w", err)
	}

	return nil
}

// GetProduct retrieves a product from the database by ID.
func (s *service) GetProduct(ctx context.Context, id string) (*coffeeshop.Product, error) {
	query := `
		SELECT id, code, images, discount, title, description, long_description, discount_percentage, reviews, map_size_price, schedules, tags, created_at, updated_at, stock_quantity, sizes
		FROM products
		WHERE id = ?
	`

	row := s.db.QueryRowContext(ctx, query, id)

	product := &coffeeshop.Product{}
	var imagesJSON, reviewsJSON, schedulesJSON, tagsJSON, mapSizePriceJSON, sizesJSON []byte

	err := row.Scan(&product.ID, &product.Code, &imagesJSON, &product.Discount, &product.Title, &product.Description,
		&product.LongDescription, &product.DiscountPercentage, &reviewsJSON, &mapSizePriceJSON,
		&schedulesJSON, &tagsJSON, &product.CreatedAt, &product.UpdatedAt, &product.StockQuantity, &sizesJSON)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product not found: %w", err)
		}
		return nil, fmt.Errorf("error getting product: %w", err)
	}

	if err := json.Unmarshal(imagesJSON, &product.Images); err != nil {
		return nil, fmt.Errorf("error unmarshaling images: %w", err)
	}

	if err := json.Unmarshal(reviewsJSON, &product.Reviews); err != nil {
		return nil, fmt.Errorf("error unmarshaling reviews: %w", err)
	}

	if err := json.Unmarshal(schedulesJSON, &product.Schedules); err != nil {
		return nil, fmt.Errorf("error unmarshaling schedules: %w", err)
	}

	if err := json.Unmarshal(tagsJSON, &product.Tags); err != nil {
		return nil, fmt.Errorf("error unmarshaling tags: %w", err)
	}

	if err := json.Unmarshal(mapSizePriceJSON, &product.MapSizePrice); err != nil {
		return nil, fmt.Errorf("error unmarshaling mapSizePrice: %w", err)
	}

	if err := json.Unmarshal(sizesJSON, &product.Sizes); err != nil {
		return nil, fmt.Errorf("error unmarshaling sizes: %w", err)
	}

	return product, nil
}

// ListProducts retrieves all products from the database.
func (s *service) ListProducts(ctx context.Context) ([]*coffeeshop.Product, error) {
	query := `
		SELECT id, code, images, discount, title, description, long_description, discount_percentage, reviews, map_size_price, schedules, tags, created_at, updated_at, stock_quantity, sizes
		FROM products
	`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error listing products: %w", err)
	}
	defer rows.Close()

	var products []*coffeeshop.Product
	for rows.Next() {
		product := &coffeeshop.Product{}
		var imagesJSON, reviewsJSON, schedulesJSON, tagsJSON, mapSizePriceJSON, sizesJSON []byte

		err := rows.Scan(&product.ID, &product.Code, &imagesJSON, &product.Discount, &product.Title, &product.Description,
			&product.LongDescription, &product.DiscountPercentage, &reviewsJSON, &mapSizePriceJSON,
			&schedulesJSON, &tagsJSON, &product.CreatedAt, &product.UpdatedAt, &product.StockQuantity, &sizesJSON)
		if err != nil {
			return nil, fmt.Errorf("error scanning product: %w", err)
		}

		if err := json.Unmarshal(imagesJSON, &product.Images); err != nil {
			return nil, fmt.Errorf("error unmarshaling images: %w", err)
		}

		if err := json.Unmarshal(reviewsJSON, &product.Reviews); err != nil {
			return nil, fmt.Errorf("error unmarshaling reviews: %w", err)
		}

		if err := json.Unmarshal(schedulesJSON, &product.Schedules); err != nil {
			return nil, fmt.Errorf("error unmarshaling schedules: %w", err)
		}

		if err := json.Unmarshal(tagsJSON, &product.Tags); err != nil {
			return nil, fmt.Errorf("error unmarshaling tags: %w", err)
		}

		if err := json.Unmarshal(mapSizePriceJSON, &product.MapSizePrice); err != nil {
			return nil, fmt.Errorf("error unmarshaling mapSizePrice: %w", err)
		}

		if err := json.Unmarshal(sizesJSON, &product.Sizes); err != nil {
			return nil, fmt.Errorf("error unmarshaling sizes: %w", err)
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

	reviewsJSON, err := json.Marshal(product.Reviews)
	if err != nil {
		return fmt.Errorf("error marshaling reviews: %w", err)
	}

	schedulesJSON, err := json.Marshal(product.Schedules)
	if err != nil {
		return fmt.Errorf("error marshaling schedules: %w", err)
	}

	tagsJSON, err := json.Marshal(product.Tags)
	if err != nil {
		return fmt.Errorf("error marshaling tags: %w", err)
	}

	mapSizePriceJSON, err := json.Marshal(product.MapSizePrice)
	if err != nil {
		return fmt.Errorf("error marshaling map_size_price: %w", err)
	}

	sizesJSON, err := json.Marshal(product.Sizes)
	if err != nil {
		return fmt.Errorf("error marshaling sizes: %w", err)
	}

	query := `
		UPDATE products
		SET code = ?, images = ?, discount = ?, title = ?, description = ?, long_description = ?, discount_percentage = ?, reviews = ?, map_size_price = ?, schedules = ?, tags = ?, updated_at = ?, stock_quantity = ?, sizes = ?
		WHERE id = ?
	`

	_, err = s.db.ExecContext(ctx, query,
		product.Code, imagesJSON, product.Discount, product.Title, product.Description,
		product.LongDescription, product.DiscountPercentage, reviewsJSON, mapSizePriceJSON,
		schedulesJSON, tagsJSON, product.UpdatedAt, product.StockQuantity, sizesJSON,
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
