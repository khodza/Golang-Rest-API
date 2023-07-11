package services

import (
	"khodza/rest-api/internal/app/models"
	"khodza/rest-api/internal/app/repositories"
	"khodza/rest-api/internal/app/validators"
)

type ProductServiceInterface interface {
	CreateProduct(product models.Product) (models.Product, error)
	GetProducts() ([]models.Product, error)
	GetProduct(productID int) (models.Product, error)
	UpdateProduct(productID int, product models.Product) (models.Product, error)
	DeleteProduct(productID int) error
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

func (s *ProductService) CreateProduct(product models.Product) (models.Product, error) {
	if err := s.validator.ValidateProduct(&product); err != nil {
		return models.Product{}, err
	}

	newProduct, err := s.productRepository.CreateProduct(product)
	if err != nil {
		return models.Product{}, err
	}

	return newProduct, nil
}

func (s *ProductService) GetProducts() ([]models.Product, error) {
	products, err := s.productRepository.GetProducts()
	if err != nil {
		return []models.Product{}, err
	}
	return products, nil
}

func (s *ProductService) GetProduct(productID int) (models.Product, error) {
	product, err := s.productRepository.GetProduct(productID)
	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (s *ProductService) UpdateProduct(productID int, product models.Product) (models.Product, error) {
	updatedProduct, err := s.productRepository.UpdateProduct(productID, product)
	if err != nil {
		return models.Product{}, err
	}

	return updatedProduct, nil
}

func (s *ProductService) DeleteProduct(productID int) error {
	if err := s.productRepository.DeleteProduct(productID); err != nil {

		return err
	}

	return nil
}
