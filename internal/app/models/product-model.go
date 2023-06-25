package models

import "time"

type Product struct {
	ID          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name" validate:"required"`
	Description string    `json:"description" db:"description"`
	Image       string    `json:"image_url" db:"image_url"`
	Barcode     string    `json:"barcode" db:"barcode" validate:"required"`
	SupplyPrice float64   `json:"supply_price" db:"supply_price" validate:"required"`
	RetailPrice float64   `json:"retail_price" db:"retail_price" validate:"required"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
