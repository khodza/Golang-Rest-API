package handlers

import (
	"khodza/rest-api/internal/app/models"
	"khodza/rest-api/internal/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productService services.ProductService
	utils          HandlerUtilities
}

func NewProductHandler(productService services.ProductService, utils HandlerUtilities) *ProductHandler {
	return &ProductHandler{
		productService: productService,
		utils:          utils,
	}
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	products, CustomError := h.productService.GetProducts()

	if CustomError.StatusCode != 0 {
		h.utils.SendCustomError(c, CustomError, "Failed to get products")
		return
	}

	//logging
	h.utils.LoggingResponse(c, "GetProducts")

	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product models.Product
	if err := h.utils.HandleJSONBinding(c, &product); err != nil {
		return
	}

	createdProduct, CustomError := h.productService.CreateProduct(product)

	if CustomError.StatusCode != 0 {
		h.utils.SendCustomError(c, CustomError, "Failed to create product")
		return
	}

	//logging
	h.utils.LoggingResponse(c, "CreateProduct")

	c.JSON(http.StatusCreated, createdProduct)
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	productID, err := h.utils.GetId(c)
	if err != nil {
		return
	}

	product, CustomError := h.productService.GetProduct(productID)

	if CustomError.StatusCode != 0 {
		h.utils.SendCustomError(c, CustomError, "Failed to get product")
		return
	}

	//logging
	h.utils.LoggingResponse(c, "GetProduct")

	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	productID, err := h.utils.GetId(c)
	if err != nil {
		return
	}

	var product models.Product
	if err := h.utils.HandleJSONBinding(c, &product); err != nil {
		return
	}

	updatedProduct, CustomError := h.productService.UpdateProduct(productID, product)
	if CustomError.StatusCode != 0 {
		h.utils.SendCustomError(c, CustomError, "Failed to update products")
		return
	}

	//logging
	h.utils.LoggingResponse(c, "UpdateProduct")

	c.JSON(http.StatusOK, updatedProduct)
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	productID, err := h.utils.GetId(c)
	if err != nil {
		return
	}
	if CustomError := h.productService.DeleteProduct(productID); CustomError.StatusCode != 0 {
		h.utils.SendCustomError(c, CustomError, "Failed to delete product")
		return
	}

	//logging
	h.utils.LoggingResponse(c, "DeleteProduct")

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
