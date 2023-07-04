package handlers

import (
	"khodza/rest-api/internal/app/models"
	"khodza/rest-api/internal/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PaymentHandler struct {
	paymentService services.PaymentService
	logger         *zap.Logger
}

func NewPaymentHandler(paymentService services.PaymentService, logger *zap.Logger) *PaymentHandler {
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

	updatedOrder, CustomError := h.paymentService.PerformPayment(payment)
	if CustomError.StatusCode != 0 {
		SendCustomError(c, CustomError, "Error on performing payment", h.logger)
		return
	}

	//logging

	LoggingResponse(c, "PerformPayment", h.logger)

	c.JSON(http.StatusAccepted, updatedOrder)
}
