package handlers

import (
	"khodza/rest-api/internal/app/models"
	"khodza/rest-api/internal/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ProductHandler struct {
	productService services.ProductService
	logger         *zap.Logger
}

func NewProductHandler(productService services.ProductService, logger *zap.Logger) *ProductHandler {
	return &ProductHandler{
		productService: productService,
		logger:         logger,
	}
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	products, CustomError := h.productService.GetProducts()

	if CustomError.StatusCode != 0 {
		SendCustomError(c, CustomError, "Failed to get products", h.logger)
		return
	}

	//logging
	LoggingResponse(c, "GetProducts", h.logger)

	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product models.Product
	if err := HandleJSONBinding(c, &product, h.logger); err != nil {
		return
	}

	createdProduct, CustomError := h.productService.CreateProduct(product)

	if CustomError.StatusCode != 0 {
		SendCustomError(c, CustomError, "Failed to create product", h.logger)
		return
	}

	//logging
	LoggingResponse(c, "CreateProduct", h.logger)

	c.JSON(http.StatusCreated, createdProduct)
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	productID, err := GetId(c, h.logger)
	if err != nil {
		return
	}

	product, CustomError := h.productService.GetProduct(productID)

	if CustomError.StatusCode != 0 {
		SendCustomError(c, CustomError, "Failed to get product", h.logger)
		return
	}

	//logging
	LoggingResponse(c, "GetProduct", h.logger)

	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	productID, err := GetId(c, h.logger)
	if err != nil {
		return
	}

	var product models.Product
	if err := HandleJSONBinding(c, &product, h.logger); err != nil {
		return
	}

	updatedProduct, CustomError := h.productService.UpdateProduct(productID, product)
	if CustomError.StatusCode != 0 {
		SendCustomError(c, CustomError, "Failed to update products", h.logger)
		return
	}

	//logging
	LoggingResponse(c, "UpdateProduct", h.logger)

	c.JSON(http.StatusOK, updatedProduct)
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	productID, err := GetId(c, h.logger)
	if err != nil {
		return
	}
	if CustomError := h.productService.DeleteProduct(productID); CustomError.StatusCode != 0 {
		SendCustomError(c, CustomError, "Failed to delete product", h.logger)
		return
	}

	//logging
	LoggingResponse(c, "DeleteProduct", h.logger)

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
