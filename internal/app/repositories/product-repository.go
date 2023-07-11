package repositories

import (
	"database/sql"
	"fmt"
	custom_errors "khodza/rest-api/internal/app/errors"
	"khodza/rest-api/internal/app/models"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type ProductRepositoryInterface interface {
	GetProducts() ([]models.Product, error)
	CreateProduct(product models.Product) (models.Product, error)
	GetProduct(productID int) (models.Product, error)
	UpdateProduct(productID int, product models.Product) (models.Product, error)
	DeleteProduct(productID int) error
}
type ProductRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) ProductRepositoryInterface {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) GetProducts() ([]models.Product, error) {
	var products []models.Product
	query := "SELECT * FROM products"
	err := r.db.Select(&products, query)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) CreateProduct(product models.Product) (models.Product, error) {
	query := "INSERT INTO products (name, barcode, supply_price, retail_price, description, image_url) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, name, barcode, supply_price, retail_price, description, image_url"
	var createdProduct models.Product
	err := r.db.Get(&createdProduct, query, product.Name, product.Barcode, product.SupplyPrice, product.RetailPrice, product.Description, product.Image)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			err = custom_errors.ErrProductCodeExist
		}
		return models.Product{}, err
	}
	return createdProduct, nil
}

func (r *ProductRepository) GetProduct(productID int) (models.Product, error) {
	var product models.Product
	query := "SELECT * FROM products WHERE id = $1"
	err := r.db.Get(&product, query, productID)
	if err != nil {
		//not found
		if err == sql.ErrNoRows {
			err = custom_errors.ErrProductNotFound
		}
		return models.Product{}, err
	}
	return product, nil
}

func (r *ProductRepository) UpdateProduct(productID int, product models.Product) (models.Product, error) {
	updateQuery := "UPDATE products SET"
	params := []interface{}{}
	paramCount := 1

	if product.Name != "" {
		updateQuery += fmt.Sprintf(" name = $%d,", paramCount)
		params = append(params, product.Name)
		paramCount++
	}

	if product.Barcode != "" {
		updateQuery += fmt.Sprintf(" barcode = $%d,", paramCount)
		params = append(params, product.Barcode)
		paramCount++
	}

	if product.SupplyPrice != 0 {
		updateQuery += fmt.Sprintf(" supply_price = $%d,", paramCount)
		params = append(params, product.SupplyPrice)
		paramCount++
	}

	if product.RetailPrice != 0 {
		updateQuery += fmt.Sprintf(" retail_price = $%d,", paramCount)
		params = append(params, product.RetailPrice)
		paramCount++
	}

	if product.Description != "" {
		updateQuery += fmt.Sprintf(" description = $%d", paramCount)
		params = append(params, product.Description)
		paramCount++
	}

	if product.Image != "" {
		updateQuery += fmt.Sprintf(" image_url = $%d", paramCount)
		params = append(params, product.Image)
		paramCount++
	}

	if len(params) == 0 {
		updatedProduct, err := r.GetProduct(productID)
		if err != nil {
			return models.Product{}, err
		}
		return updatedProduct, nil
	}

	// Add updated_at column update
	updateQuery += " updated_at = CURRENT_TIMESTAMP,"

	updateQuery = strings.TrimSuffix(updateQuery, ",")

	updateQuery += fmt.Sprintf(" WHERE id = $%d", paramCount)
	params = append(params, productID)

	_, err := r.db.Exec(updateQuery, params...)
	if err != nil {
		//duplicate error
		pqErr, _ := err.(*pq.Error)
		if pqErr.Code == "23505" {
			err = custom_errors.ErrProductCodeExist
		}
		//not found
		if err == sql.ErrNoRows {
			err = custom_errors.ErrProductNotFound
		}
		return models.Product{}, err
	}
	// Retrieve the updated product from the database
	updatedProduct, err := r.GetProduct(productID)
	if err != nil {
		return models.Product{}, err
	}
	return updatedProduct, nil
}

func (r *ProductRepository) DeleteProduct(productID int) error {
	query := "DELETE FROM products WHERE id = $1"
	_, err := r.db.Exec(query, productID)
	if err != nil {
		return err
	}

	return nil
}
