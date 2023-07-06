package handlers

import (
	"khodza/rest-api/internal/app/models"
	"khodza/rest-api/internal/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type OrderHandler struct {
	orderService services.OrderServiceInterface
	logger       *zap.Logger
}

func NewOrderHandler(orderService services.OrderServiceInterface, logger *zap.Logger) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
		logger:       logger,
	}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var newOrder models.OrderReq
	if err := HandleJSONBinding(c, &newOrder, h.logger); err != nil {
		return
	}

	createdOrder, CustomError := h.orderService.CreateOrder(newOrder)
	if CustomError.StatusCode != 0 {
		SendCustomError(c, CustomError, "Failed to create order", h.logger)
		return
	}

	//logging
	LoggingResponse(c, "CreateOrder", h.logger)

	c.JSON(http.StatusOK, createdOrder)
}

func (h *OrderHandler) GetOrder(c *gin.Context) {
	orderID, err := GetId(c, h.logger)
	if err != nil {
		return
	}

	order, CustomError := h.orderService.GetOrder(orderID)
	if CustomError.StatusCode != 0 {
		SendCustomError(c, CustomError, "Failed to get order", h.logger)
		return
	}

	//logging
	LoggingResponse(c, "GetOrder", h.logger)

	c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) GetOrders(c *gin.Context) {
	orders, CustomError := h.orderService.GetOrders()
	if CustomError.StatusCode != 0 {
		SendCustomError(c, CustomError, "Failed to get orders", h.logger)
		return
	}

	//logging
	LoggingResponse(c, "GetOrders", h.logger)

	c.JSON(http.StatusOK, orders)
}

func (h *OrderHandler) GetOrderItems(c *gin.Context) {
	orderID, err := GetId(c, h.logger)
	if err != nil {
		return
	}

	orderItems, CustomError := h.orderService.GetOrderItems(orderID)
	if CustomError.StatusCode != 0 {
		SendCustomError(c, CustomError, "Failed to get items", h.logger)
		return
	}

	//logging
	LoggingResponse(c, "GetOrderItems", h.logger)

	c.JSON(http.StatusOK, orderItems)
}

func (h *OrderHandler) UpdateOrder(c *gin.Context) {
	var newOrder models.OrderReq
	orderID, err := GetId(c, h.logger)
	if err != nil {
		return
	}

	if err = HandleJSONBinding(c, &newOrder, h.logger); err != nil {
		return
	}

	updatedOrder, CustomError := h.orderService.UpdateOrder(orderID, newOrder)
	if CustomError.StatusCode != 0 {
		SendCustomError(c, CustomError, "Failed to update order", h.logger)
		return
	}

	//logging
	LoggingResponse(c, "UpdateOrder", h.logger)

	c.JSON(http.StatusOK, updatedOrder)
}

func (h *OrderHandler) DeleteOrder(c *gin.Context) {
	orderID, err := GetId(c, h.logger)
	if err != nil {
		return
	}

	CustomError := h.orderService.DeleteOrder(orderID)
	if CustomError.StatusCode != 0 {
		SendCustomError(c, CustomError, "Failed to delete order", h.logger)
		return
	}

	//logging
	LoggingResponse(c, "DeleteOrder", h.logger)

	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}

func (h *OrderHandler) GetPaidOrders(c *gin.Context) {
	orders, CustomError := h.orderService.GetPaidOrders()
	if CustomError.StatusCode != 0 {
		SendCustomError(c, CustomError, "Failed to get paid orders", h.logger)
		return
	}

	//logging
	LoggingResponse(c, "GetPaidOrders", h.logger)

	c.JSON(http.StatusOK, orders)
}
