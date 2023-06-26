package routers

import (
	"khodza/rest-api/internal/app/handlers"

	"github.com/gin-gonic/gin"
)

func SetupPaymentRouter(router *gin.RouterGroup, paymentHandler *handlers.PaymentHandler) {
	orderGroup := router.Group("")
	orderGroup.POST("/", paymentHandler.PerformPayment)
}
