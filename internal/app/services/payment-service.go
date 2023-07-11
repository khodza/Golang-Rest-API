package services

import (
	custom_errors "khodza/rest-api/internal/app/errors"
	"khodza/rest-api/internal/app/models"
)

type PaymentServiceInterface interface {
	PerformPayment(payment models.Payment) (models.Order, error)
}
type PaymentService struct {
	orderService OrderServiceInterface
}

func NewPaymentService(
	orderService OrderServiceInterface,
) PaymentServiceInterface {
	return &PaymentService{
		orderService: orderService,
	}
}

func (s *PaymentService) PerformPayment(payment models.Payment) (models.Order, error) {

	orderRes, err := s.orderService.GetOrder(payment.OrderID)
	if err != nil {
		return models.Order{}, err
	}

	if payment.RetailPrice != orderRes.Order.RetailPrice {
		return models.Order{}, custom_errors.ErrPaymentNotEqual
	}

	updatedOrder, err := s.orderService.ChangeStatus(payment.OrderID, "paid")
	if err != nil {
		return models.Order{}, err
	}
	return updatedOrder, err
}
