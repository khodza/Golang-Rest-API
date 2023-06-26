package handlers

import (
	"fmt"
	"khodza/rest-api/internal/app/models"
	"khodza/rest-api/internal/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	paymentService services.PaymentService
	utils          HandlerUtilities
}

func NewPaymentHandler(paymentService services.PaymentService, utils HandlerUtilities) *PaymentHandler {
	return &PaymentHandler{
		paymentService: paymentService,
		utils:          utils,
	}
}

func (h *PaymentHandler) PerformPayment(c *gin.Context) {
	var payment models.Payment
	err := h.utils.HandleJSONBinding(c, &payment)
	if err != nil {
		return
	}
	fmt.Println(payment)
	updatedOrder, CustomError := h.paymentService.PerformPayment(payment)
	if CustomError.StatusCode != 0 {
		h.utils.SendCustomError(c, CustomError, "Error on performing payment")
		return
	}

	//logging

	h.utils.LoggingResponse(c, "PerformPayment")

	c.JSON(http.StatusAccepted, updatedOrder)
}
