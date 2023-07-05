package services

import (
	"database/sql"
	"khodza/rest-api/internal/app/models"
	"khodza/rest-api/internal/app/repositories"
	"khodza/rest-api/internal/app/validators"
	"net/http"

	"github.com/lib/pq"
)

type ProductServiceInterface interface {
	CreateProduct(product models.Product) (models.Product, CustomError)
	GetProducts() ([]models.Product, CustomError)
	GetProduct(productID int) (models.Product, CustomError)
	UpdateProduct(productID int, product models.Product) (models.Product, CustomError)
	DeleteProduct(productID int) CustomError
}
type ProductService struct {
	productRepository repositories.ProductRepositoryInterface
	validator         validators.ProductValidatorInterface
}

func NewProductService(productRepository repositories.ProductRepositoryInterface, productValidator validators.ProductValidatorInterface) *ProductService {
	return &ProductService{
		productRepository: productRepository,
		validator:         productValidator,
	}
}

func (s *ProductService) CreateProduct(product models.Product) (models.Product, CustomError) {
	if err := s.validator.ValidateProduct(&product); err != nil {
		return models.Product{}, CustomError{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Err:        err,
		}
	}

	newProduct, err := s.productRepository.CreateProduct(product)
	if err != nil {
		pqErr, _ := err.(*pq.Error)
		if pqErr.Code == "23505" {
			return models.Product{}, CustomError{
				StatusCode: http.StatusBadRequest,
				Message:    "Product with the same barcode already exists",
				Err:        err,
			}
		}

		return models.Product{}, CustomError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to create product",
			Err:        err,
		}
	}

	return newProduct, CustomError{}
}

func (s *ProductService) GetProducts() ([]models.Product, CustomError) {
	products, err := s.productRepository.GetProducts()
	if err != nil {
		return []models.Product{}, CustomError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to get products",
			Err:        err,
		}
	}
	return products, CustomError{}
}

func (s *ProductService) GetProduct(productID int) (models.Product, CustomError) {
	product, err := s.productRepository.GetProduct(productID)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Product{}, CustomError{
				StatusCode: http.StatusNotFound,
				Message:    "No product found with the given ID",
				Err:        err,
			}
		}

		return models.Product{}, CustomError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Something went wrong while getting product",
			Err:        err,
		}
	}

	return product, CustomError{}
}

func (s *ProductService) UpdateProduct(productID int, product models.Product) (models.Product, CustomError) {
	// if err := s.validator.ValidateProduct(&product); err != nil {
	// 	return models.Product{}, CustomError{
	// 		StatusCode: http.StatusBadRequest,
	// 		Message:    err.Error(),
	// 		Err:        err,
	// 	}
	// }

	updatedProduct, err := s.productRepository.UpdateProduct(productID, product)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Product{}, CustomError{
				StatusCode: http.StatusNotFound,
				Message:    "No product found with the given ID",
				Err:        err,
			}
		}

		pqErr, _ := err.(*pq.Error)
		if pqErr.Code == "23505" {
			return models.Product{}, CustomError{
				StatusCode: http.StatusBadRequest,
				Message:    "Product with the same barcode already exists",
				Err:        err,
			}
		}

		return models.Product{}, CustomError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Something went wrong while updating product",
			Err:        err,
		}
	}

	return updatedProduct, CustomError{}
}

func (s *ProductService) DeleteProduct(productID int) CustomError {
	if err := s.productRepository.DeleteProduct(productID); err != nil {
		if err == sql.ErrNoRows {
			return CustomError{
				StatusCode: http.StatusNotFound,
				Message:    "No product found with the given ID",
				Err:        err,
			}
		}

		return CustomError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Something went wrong while deleting product",
			Err:        err,
		}
	}

	return CustomError{}
}
