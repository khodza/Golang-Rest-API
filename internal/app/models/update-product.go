package models

type UpdateProduct struct {
	ID          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Barcode     string `json:"barcode" db:"barcode"`
	SupplyPrice int    `json:"supply_price" db:"supply_price"`
	RetailPrice int    `json:"retail_price" db:"retail_price"`
}
