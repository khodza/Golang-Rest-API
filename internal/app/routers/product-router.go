package routers

import (
	"khodza/rest-api/internal/app/handlers"

	"github.com/gin-gonic/gin"
)

func SetupProductRouter(router *gin.RouterGroup, productHandler *handlers.ProductHandler) {
	userGroup := router.Group("")
	userGroup.GET("/", productHandler.GetProducts)
	userGroup.POST("/", productHandler.CreateProduct)
	userGroup.GET("/:id", productHandler.GetProduct)
	userGroup.PATCH("/:id", productHandler.UpdateProduct)
	userGroup.DELETE("/:id", productHandler.DeleteProduct)
}
