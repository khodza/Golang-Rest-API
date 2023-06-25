package routers

import (
	"khodza/rest-api/internal/app/handlers"

	"github.com/gin-gonic/gin"
)

func SetupProductRouter(router *gin.RouterGroup, productHandler *handlers.ProductHandler) {
	productGroup := router.Group("")
	productGroup.GET("/", productHandler.GetProducts)
	productGroup.POST("/", productHandler.CreateProduct)
	productGroup.GET("/:id", productHandler.GetProduct)
	productGroup.PATCH("/:id", productHandler.UpdateProduct)
	productGroup.DELETE("/:id", productHandler.DeleteProduct)
}
