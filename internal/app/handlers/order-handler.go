package handlers

import (
	"khodza/rest-api/internal/app/models"
	"khodza/rest-api/internal/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderService services.OrderService
	utils        HandlerUtilities
}

func NewOrderHandler(orderService services.OrderService, utils HandlerUtilities) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
		utils:        utils,
	}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var newOrder models.OrderReq
	if err := h.utils.HandleJSONBinding(c, &newOrder); err != nil {
		return
	}

	createdOrder, CustomError := h.orderService.CreateOrder(newOrder)
	if CustomError.StatusCode != 0 {
		h.utils.SendCustomError(c, CustomError, "Failed to create order")
		return
	}

	//logging
	h.utils.LoggingResponse(c, "CreateOrder")

	c.JSON(http.StatusOK, createdOrder)
}

func (h *OrderHandler) GetOrder(c *gin.Context) {
	orderID, err := h.utils.GetId(c)
	if err != nil {
		return
	}

	order, CustomError := h.orderService.GetOrder(orderID)
	if CustomError.StatusCode != 0 {
		h.utils.SendCustomError(c, CustomError, "Failed to get order")
		return
	}

	//logging
	h.utils.LoggingResponse(c, "GetOrder")

	c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) GetOrders(c *gin.Context) {
	orders, CustomError := h.orderService.GetOrders()
	if CustomError.StatusCode != 0 {
		h.utils.SendCustomError(c, CustomError, "Failed to get orders")
		return
	}

	//logging
	h.utils.LoggingResponse(c, "GetOrders")

	c.JSON(http.StatusOK, orders)
}

func (h *OrderHandler) UpdateOrder(c *gin.Context) {
	var newOrder models.OrderReq
	orderID, err := h.utils.GetId(c)
	if err != nil {
		return
	}

	if err = h.utils.HandleJSONBinding(c, &newOrder); err != nil {
		return
	}

	updatedOrder, CustomError := h.orderService.UpdateOrder(orderID, newOrder)
	if CustomError.StatusCode != 0 {
		h.utils.SendCustomError(c, CustomError, "Failed to update order")
		return
	}

	//logging
	h.utils.LoggingResponse(c, "UpdateOrder")

	c.JSON(http.StatusOK, updatedOrder)
}

func (h *OrderHandler) DeleteOrder(c *gin.Context) {
	orderID, err := h.utils.GetId(c)
	if err != nil {
		return
	}

	CustomError := h.orderService.DeleteOrder(orderID)
	if CustomError.StatusCode != 0 {
		h.utils.SendCustomError(c, CustomError, "Failed to delete order")
		return
	}

	//logging
	h.utils.LoggingResponse(c, "DeleteOrder")

	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}
