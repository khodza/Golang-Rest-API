package repositories

import (
	"fmt"
	"khodza/rest-api/internal/app/models"
	"strings"

	"github.com/jmoiron/sqlx"
)

type OrderRepositoryInterface interface {
	CreateOrder(order models.Order, tx *sqlx.Tx) (int, error)
	GetOrder(orderID int) (models.Order, error)
	GetOrders(status string) ([]models.Order, error)
	GetOrdersCount() (int, error)
	UpdateOrder(orderID int, newOrder models.Order, tx *sqlx.Tx) (models.Order, error)
	DeleteOrder(orderID int) error
	CreateOrderItem(orderItem models.OrderItem, tx *sqlx.Tx) (int, error)
	GetOrderItems(orderID int) ([]models.OrderItem, error)
	DeleteOrderItems(orderID int, tx *sqlx.Tx) error
	BeginTransaction() (*sqlx.Tx, error)
}
type OrderRepository struct {
	db *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) OrderRepositoryInterface {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) CreateOrder(order models.Order, tx *sqlx.Tx) (int, error) {
	query := "INSERT INTO orders (user_id, supply_price, retail_price, status) VALUES($1, $2, $3, $4) RETURNING id"
	var createdOrder models.Order

	//if transaction
	if tx != nil {
		err := tx.Get(&createdOrder, query, order.UserID, order.SupplyPrice, order.RetailPrice, order.Status)
		if err != nil {
			return 0, err
		}
	} else {
		err := r.db.Get(&createdOrder, query, order.UserID, order.SupplyPrice, order.RetailPrice, order.Status)
		if err != nil {
			return 0, err
		}
	}

	return createdOrder.ID, nil
}

func (r *OrderRepository) GetOrder(orderID int) (models.Order, error) {
	var order models.Order
	query := "SELECT * FROM orders WHERE id = $1"
	err := r.db.Get(&order, query, orderID)
	if err != nil {
		return models.Order{}, err
	}
	return order, nil
}

func (r *OrderRepository) GetOrders(status string) ([]models.Order, error) {
	var orders []models.Order
	var err error
	query := "SELECT * FROM orders"
	if status != "" {
		query += " WHERE status = $1"
		err = r.db.Select(&orders, query, status)
	} else {
		err = r.db.Select(&orders, query)
	}
	if err != nil {
		return []models.Order{}, err
	}
	return orders, nil
}

func (r *OrderRepository) GetOrdersCount() (int, error) {
	var count int
	query := "SELECT COUNT(*) FROM orders"
	err := r.db.Get(&count, query)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *OrderRepository) UpdateOrder(orderID int, newOrder models.Order, tx *sqlx.Tx) (models.Order, error) {
	updateQuery := "UPDATE orders SET"
	params := []interface{}{}
	paramCount := 1

	if newOrder.Status != "" {
		updateQuery += fmt.Sprintf(" status = $%d,", paramCount)
		params = append(params, newOrder.Status)
		paramCount++
	}

	if newOrder.UserID != 0 {
		updateQuery += fmt.Sprintf(" user_id = $%d,", paramCount)
		params = append(params, newOrder.UserID)
		paramCount++
	}

	if newOrder.SupplyPrice != 0 {
		updateQuery += fmt.Sprintf(" supply_price = $%d,", paramCount)
		params = append(params, newOrder.SupplyPrice)
		paramCount++
	}

	if newOrder.RetailPrice != 0 {
		updateQuery += fmt.Sprintf(" retail_price = $%d,", paramCount)
		params = append(params, newOrder.RetailPrice)
		paramCount++
	}

	if len(params) == 0 {
		var updatedOrder models.Order
		getQuery := "SELECT * FROM orders WHERE id = $1"
		err := r.db.Get(&updatedOrder, getQuery, orderID)
		if err != nil {
			return models.Order{}, err
		}

		return updatedOrder, nil
	}

	// Add updated_at column update
	updateQuery += " updated_at = CURRENT_TIMESTAMP,"

	updateQuery = strings.TrimSuffix(updateQuery, ",")

	updateQuery += fmt.Sprintf(" WHERE id = $%d", paramCount)
	params = append(params, orderID)

	//if transaction
	if tx != nil {
		_, err := tx.Exec(updateQuery, params...)
		if err != nil {
			return models.Order{}, err
		}
	} else {
		_, err := r.db.Exec(updateQuery, params...)
		if err != nil {
			return models.Order{}, err
		}
	}

	var updatedOrder models.Order
	getQuery := "SELECT * FROM orders WHERE id = $1"
	err := r.db.Get(&updatedOrder, getQuery, orderID)
	if err != nil {
		return models.Order{}, err
	}

	return updatedOrder, nil
}

func (r *OrderRepository) DeleteOrder(orderID int) error {
	query := "DELETE FROM orders WHERE id = $1"
	_, err := r.db.Exec(query, orderID)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) CreateOrderItem(orderItem models.OrderItem, tx *sqlx.Tx) (int, error) {
	query := "INSERT INTO order_items (order_id, product_id, quantity) VALUES($1, $2, $3) RETURNING id"
	var createdOrderItem models.OrderItem
	if tx != nil {
		err := tx.Get(&createdOrderItem, query, orderItem.OrderID, orderItem.ProductID, orderItem.Quantity)
		if err != nil {
			return 0, err
		}
	} else {
		err := r.db.Get(&createdOrderItem, query, orderItem.OrderID, orderItem.ProductID, orderItem.Quantity)
		if err != nil {
			return 0, err
		}
	}
	return createdOrderItem.OrderID, nil
}

func (r *OrderRepository) GetOrderItems(orderID int) ([]models.OrderItem, error) {
	var orderItems []models.OrderItem
	query := "SELECT * FROM order_items WHERE order_id = $1"
	err := r.db.Select(&orderItems, query, orderID)
	if err != nil {
		return []models.OrderItem{}, err
	}
	return orderItems, nil
}

func (r *OrderRepository) DeleteOrderItems(orderID int, tx *sqlx.Tx) error {
	query := "DELETE FROM order_items WHERE order_id = $1"
	if tx != nil {
		_, err := tx.Exec(query, orderID)
		if err != nil {
			return err
		}
	} else {
		_, err := r.db.Exec(query, orderID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *OrderRepository) BeginTransaction() (*sqlx.Tx, error) {
	return r.db.Beginx()
}
