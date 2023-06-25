package services

import (
	"database/sql"
	"khodza/rest-api/internal/app/models"
	"khodza/rest-api/internal/app/repositories"
	"net/http"
)

type OrderService struct {
	orderRepository repositories.OrderRepository
	productService  ProductService
}

func NewOrderService(
	orderRepository repositories.OrderRepository,
	productService ProductService,
) *OrderService {
	return &OrderService{
		orderRepository: orderRepository,
		productService:  productService,
	}
}

func (s *OrderService) CreateOrder(newOrder models.OrderReq) (models.Order, CustomError) {
	var order models.Order
	order.UserID = newOrder.UserID
	order.Status = "pending"
	products := newOrder.Products

	var supplyPrice float64
	var retailPrice float64
	for i := 0; i < len(products); i++ {
		product, CustomError := s.productService.GetProduct(products[i].ProductID)
		if CustomError.StatusCode != 0 {
			return models.Order{}, CustomError
		}

		supplyPrice += product.SupplyPrice * float64(products[i].Quantity)
		retailPrice += product.RetailPrice * float64(products[i].Quantity)
	}
	order.SupplyPrice = supplyPrice
	order.RetailPrice = retailPrice

	//create order
	orderID, err := s.orderRepository.CreateOrder(order)
	if err != nil {
		return models.Order{}, CustomError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error on creating order",
			Err:        err,
		}
	}
	order.ID = orderID
	//create order items
	for i := 0; i < len(products); i++ {
		var orderItem models.OrderItem
		orderItem.OrderID = orderID
		orderItem.ProductID, orderItem.Quantity = products[i].ProductID, products[i].Quantity
		_, err := s.orderRepository.CreateOrderItem(orderItem)
		if err != nil {
			return models.Order{}, CustomError{
				StatusCode: http.StatusInternalServerError,
				Message:    "Error on creating order items",
				Err:        err,
			}
		}
	}

	return order, CustomError{}
}

func (s *OrderService) GetOrder(orderID int) (models.OrderRes, CustomError) {
	var readyOrder models.OrderRes
	order, err := s.orderRepository.GetOrder(orderID)
	if err != nil {

		if err == sql.ErrNoRows {
			return models.OrderRes{}, CustomError{
				StatusCode: http.StatusNotFound,
				Message:    "No order found with the given ID",
				Err:        err,
			}
		}
		return models.OrderRes{}, CustomError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error on getting order",
			Err:        err,
		}
	}
	readyOrder.Order = order
	orderItems, err := s.orderRepository.GetOrderItems(order.ID)
	if err != nil {
		return models.OrderRes{}, CustomError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error on getting order items",
			Err:        err,
		}
	}
	readyOrder.Products = orderItems
	return readyOrder, CustomError{}
}

func (s *OrderService) GetOrders() ([]models.OrderRes, CustomError) {
	var resOrders []models.OrderRes
	orders, err := s.orderRepository.GetOrders()
	if err != nil {
		return []models.OrderRes{}, CustomError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error on getting order",
		}
	}
	for i := 0; i < len(orders); i++ {
		var orderRes models.OrderRes
		orderItems, err := s.orderRepository.GetOrderItems(orders[i].ID)
		if err != nil {
			return []models.OrderRes{}, CustomError{
				StatusCode: http.StatusInternalServerError,
				Message:    "Error on getting orders items",
			}
		}
		orderRes.Products = orderItems
		orderRes.Order = orders[i]
		resOrders = append(resOrders, orderRes)
	}
	return resOrders, CustomError{}
}

func (s *OrderService) UpdateOrder(orderID int, newOrder models.OrderReq) (models.Order, CustomError) {
	var updatedOrderTemp models.Order
	updatedOrderTemp.UserID = newOrder.UserID
	products := newOrder.Products

	if len(products) != 0 {
		var supplyPrice float64
		var retailPrice float64
		for i := 0; i < len(products); i++ {
			product, CustomError := s.productService.GetProduct(products[i].ProductID)
			if CustomError.StatusCode != 0 {
				return models.Order{}, CustomError
			}

			supplyPrice += product.SupplyPrice * float64(products[i].Quantity)
			retailPrice += product.RetailPrice * float64(products[i].Quantity)
		}
		updatedOrderTemp.SupplyPrice = supplyPrice
		updatedOrderTemp.RetailPrice = retailPrice
	}

	updatedOrder, err := s.orderRepository.UpdateOrder(orderID, updatedOrderTemp)
	if err != nil {
		return models.Order{}, CustomError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error on updating order",
			Err:        err,
		}
	}

	//update order items
	if len(products) != 0 {
		//delete items
		var err error
		err = s.orderRepository.DeleteOrderItems(orderID)
		if err != nil {
			return models.Order{}, CustomError{
				StatusCode: http.StatusInternalServerError,
				Message:    "Error on updating (deleting old order items) ",
				Err:        err,
			}
		}

		for i := 0; i < len(products); i++ {
			//create items
			var orderItem models.OrderItem
			orderItem.OrderID = orderID
			orderItem.ProductID, orderItem.Quantity = products[i].ProductID, products[i].Quantity
			_, err = s.orderRepository.CreateOrderItem(orderItem)
			if err != nil {
				return models.Order{}, CustomError{
					StatusCode: http.StatusInternalServerError,
					Message:    "Error on creating order items",
					Err:        err,
				}
			}
		}
	}

	return updatedOrder, CustomError{}

}

func (s *OrderService) DeleteOrder(orderID int) CustomError {
	//delete order
	err := s.orderRepository.DeleteOrder(orderID)
	if err != nil {

		return CustomError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error on deleting order",
			Err:        err,
		}
	}

	// //delete order items

	// err = s.orderRepository.DeleteOrderItems(orderID)
	// if err != nil {
	// 	return CustomError{
	// 		StatusCode: http.StatusInternalServerError,
	// 		Message:    "Error on deleting order items",
	// 		Err:        err,
	// 	}
	// }
	return CustomError{}
}
