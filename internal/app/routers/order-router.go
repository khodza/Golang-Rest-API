package routers

import (
	"khodza/rest-api/internal/app/handlers"

	"github.com/gin-gonic/gin"
)

func SetupOrderRouter(router *gin.RouterGroup, orderHandler *handlers.OrderHandler) {
	orderGroup := router.Group("")
	orderGroup.GET("/", orderHandler.GetOrders)
	orderGroup.POST("/", orderHandler.CreateOrder)
	orderGroup.GET("/paid", orderHandler.GetPaidOrders)
	orderGroup.GET("/:id", orderHandler.GetOrder)
	orderGroup.PATCH("/:id", orderHandler.UpdateOrder)
	orderGroup.DELETE("/:id", orderHandler.DeleteOrder)
}
