
-- name: GetProduct :one
SELECT * FROM products WHERE id = ?;

-- name: ListProducts :many
SELECT * FROM products LIMIT ?;

-- name: CreateProduct :exec
INSERT INTO products (id, code, images, title, description, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: UpdateProduct :exec
UPDATE products SET code = ?, images = ?, title = ?, description = ?, updated_at = ? WHERE id = ?;

-- name: DeleteProduct :exec
DELETE FROM products WHERE id = ?;

-- name: GetTag :one
SELECT * FROM tags WHERE id = ?;

-- name: GetTagByName :one
SELECT * FROM tags WHERE name = ?;

-- name: CreateTag :exec
INSERT INTO tags (id, name) VALUES (?, ?);

-- name: CreateProductTag :exec
INSERT INTO product_tags (product_id, tag_id) VALUES (?, ?);

-- name: GetProductTags :many
SELECT t.name
FROM tags t
JOIN product_tags pt ON t.id = pt.tag_id
WHERE pt.product_id = ?;
