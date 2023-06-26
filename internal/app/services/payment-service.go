package services

import (
	"fmt"
	"khodza/rest-api/internal/app/models"
	"net/http"
)

type PaymentService struct {
	orderService OrderService
}

func NewPaymentService(
	orderService OrderService,
) *PaymentService {
	return &PaymentService{
		orderService: orderService,
	}
}

func (s *PaymentService) PerformPayment(payment models.Payment) (models.Order, CustomError) {

	orderRes, errC := s.orderService.GetOrder(payment.OrderID)
	if errC.StatusCode != 0 {
		return models.Order{}, errC
	}

	if payment.RetailPrice != orderRes.Order.RetailPrice {
		return models.Order{}, CustomError{
			StatusCode: http.StatusBadRequest,
			Message:    "The provided cash less or more then needed",
			Err:        fmt.Errorf("the provided cash less or more then needed"),
		}
	}

	updatedOrder, err := s.orderService.ChangeStatus(payment.OrderID, "paid")
	if err != nil {
		return models.Order{}, CustomError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error on changing status of order ",
			Err:        err,
		}
	}
	return updatedOrder, CustomError{}
}
