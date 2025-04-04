// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package database

import (
	"database/sql"
)

type Inventory struct {
	ID        string
	ProductID string
	Stock     int64
	Sizes     sql.NullString
	Price     float64
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}

type Order struct {
	ID              string
	UserID          string
	OrderDate       sql.NullTime
	TotalAmount     float64
	ShippingAddress sql.NullString
	BillingAddress  sql.NullString
	PaymentMethod   sql.NullString
	OrderStatus     sql.NullString
	CreatedAt       sql.NullTime
	UpdatedAt       sql.NullTime
}

type OrderItem struct {
	ID        string
	OrderID   string
	ProductID string
	Quantity  int64
	Price     float64
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}

type Product struct {
	ID          string
	Code        string
	Images      sql.NullString
	Title       string
	Description sql.NullString
	CreatedAt   sql.NullTime
	UpdatedAt   sql.NullTime
}

type ProductTag struct {
	ProductID string
	TagID     string
}

type Review struct {
	ID        string
	ProductID string
	UserID    string
	Rating    int64
	Comment   sql.NullString
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}

type Tag struct {
	ID        string
	Name      string
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}

type User struct {
	ID         string
	Email      string
	Subscriber sql.NullBool
	Username   sql.NullString
	CreatedAt  sql.NullTime
	UpdatedAt  sql.NullTime
}
