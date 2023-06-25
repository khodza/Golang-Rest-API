package models

import "time"

type Order struct {
	ID          int       `json:"id" db:"id"`
	UserID      int       `json:"user_id" db:"user_id"`
	SupplyPrice float64   `json:"supply_price" db:"supply_price"`
	RetailPrice float64   `json:"retail_price" db:"retail_price"`
	Status      string    `json:"status" db:"status"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`
}

type OrderItem struct {
	ID        int       `json:"id" db:"id"`
	OrderID   int       `json:"order_number" db:"order_id"`
	ProductID int       `json:"product_id" db:"product_id"`
	Quantity  int       `json:"quantity" db:"quantity"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

type OrderProduct struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type OrderReq struct {
	UserID   int            `json:"user_id" db:"user_id"`
	Products []OrderProduct `json:"products" db:"-"`
}

type OrderRes struct {
	Order    Order       `json:"order"`
	Products []OrderItem `json:"products"`
}
