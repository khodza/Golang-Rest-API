package routers

import (
	"khodza/rest-api/internal/app/handlers"

	"github.com/gin-gonic/gin"
)

func ConnectRoutersToHandlers(router *gin.Engine, handlersMap map[string]interface{}) {
	for route, handler := range handlersMap {
		routeGroup := router.Group("/" + route)
		switch route {
		case "users":
			userHandler := handler.(*handlers.UserHandler)
			SetupUserRouter(routeGroup, userHandler)
		case "products":
			productHandler := handler.(*handlers.ProductHandler)
			SetupProductRouter(routeGroup, productHandler)
		case "orders":
			orderHandler := handler.(*handlers.OrderHandler)
			SetupOrderRouter(routeGroup, orderHandler)
		case "payments":
			paymentHandler := handler.(*handlers.PaymentHandler)
			SetupPaymentRouter(routeGroup, paymentHandler)
		}
	}
}
