package handlers

import (
	"fmt"
	"khodza/rest-api/internal/app/models"
	"khodza/rest-api/internal/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productService services.ProductService
}

func NewProductHandler(productService services.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	products, CustomError := h.productService.GetProducts()
	fmt.Println("HELLLLO")
	if CustomError.StatusCode != 0 {
		c.JSON(CustomError.StatusCode, gin.H{"error": CustomError.Message})
		fmt.Println(CustomError.Err.Error())
		return
	}

	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product models.Product
	if err := HandleJSONBinding(c, &product); err != nil {
		return
	}

	createdProduct, CustomError := h.productService.CreateProduct(product)

	if CustomError.StatusCode != 0 {
		c.JSON(CustomError.StatusCode, gin.H{"error": CustomError.Message})
		fmt.Println(CustomError.Err.Error())
		return
	}

	c.JSON(http.StatusCreated, createdProduct)
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	productID, err := GetId(c)
	if err != nil {
		return
	}

	product, CustomError := h.productService.GetProduct(productID)

	if CustomError.StatusCode != 0 {
		c.JSON(CustomError.StatusCode, gin.H{"message": CustomError.Message})
		fmt.Println(CustomError.Err.Error())
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	productID, err := GetId(c)
	if err != nil {
		return
	}

	var product models.Product
	if err := HandleJSONBinding(c, &product); err != nil {
		return
	}

	updatedProduct, CustomError := h.productService.UpdateProduct(productID, product)
	if CustomError.StatusCode != 0 {
		c.JSON(CustomError.StatusCode, gin.H{"message": CustomError.Message})
		fmt.Println(CustomError.Err.Error())
		return
	}

	c.JSON(http.StatusOK, updatedProduct)
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	productID, err := GetId(c)
	if err != nil {
		return
	}
	if CustomError := h.productService.DeleteProduct(productID); CustomError.StatusCode != 0 {
		c.JSON(CustomError.StatusCode, gin.H{"message": CustomError.Message})
		fmt.Println(CustomError.Err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
