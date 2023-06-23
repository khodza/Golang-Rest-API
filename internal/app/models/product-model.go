package models

type Product struct {
	ID          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name" validate:"required"`
	Barcode     string `json:"barcode" db:"barcode" validate:"required"`
	SupplyPrice int    `json:"supply_price" db:"supply_price" validate:"required"`
	RetailPrice int    `json:"retail_price" db:"retail_price" validate:"required"`
}
