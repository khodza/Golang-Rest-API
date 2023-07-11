package handlers

import (
	"khodza/rest-api/internal/app/models"
	"khodza/rest-api/internal/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PaymentHandler struct {
	paymentService services.PaymentServiceInterface
	logger         *zap.Logger
}

func NewPaymentHandler(paymentService services.PaymentServiceInterface, logger *zap.Logger) *PaymentHandler {
	return &PaymentHandler{
		paymentService: paymentService,
		logger:         logger,
	}
}

func (h *PaymentHandler) PerformPayment(c *gin.Context) {
	var payment models.Payment
	err := HandleJSONBinding(c, &payment, h.logger)
	if err != nil {
		return
	}

	updatedOrder, err := h.paymentService.PerformPayment(payment)
	if err != nil {
		c.Error(err)
		return
	}

	//logging

	LoggingResponse(c, "PerformPayment", h.logger)

	c.JSON(http.StatusAccepted, updatedOrder)
}
