package services

import (
	"database/sql"
	"khodza/rest-api/internal/app/models"
	"khodza/rest-api/internal/app/repositories"
	"net/http"
	"sync"
)

type OrderServiceInterface interface {
	CreateOrder(newOrder models.OrderReq) (models.Order, CustomError)
	GetOrder(orderID int) (models.OrderRes, CustomError)
	GetOrders() ([]models.OrderRes, CustomError)
	UpdateOrder(orderID int, newOrder models.OrderReq) (models.Order, CustomError)
	DeleteOrder(orderID int) CustomError
	ChangeStatus(orderID int, status string) (models.Order, error)
	GetPaidOrders() (models.OrderPaid, CustomError)
}
type OrderService struct {
	orderRepository repositories.OrderRepositoryInterface
	productService  ProductServiceInterface
}

func NewOrderService(
	orderRepository repositories.OrderRepositoryInterface,
	productService ProductServiceInterface,
) *OrderService {
	return &OrderService{
		orderRepository: orderRepository,
		productService:  productService,
	}
}

func (s *OrderService) CreateOrder(newOrder models.OrderReq) (models.Order, CustomError) {
	//Start transaction
	tx, err := s.orderRepository.BeginTransaction()
	if err != nil {
		return models.Order{}, CustomError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error starting transaction",
			Err:        err,
		}
	}

	//Rollback
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var order models.Order
	order.UserID = newOrder.UserID
	order.Status = "pending"
	products := newOrder.Products

	//Calculating prices
	var totalSupplyPrice float64
	var totalRetailPrice float64
	for i := 0; i < len(products); i++ {
		product, CustomError := s.productService.GetProduct(products[i].ProductID)
		if CustomError.StatusCode != 0 {
			return models.Order{}, CustomError
		}

		totalSupplyPrice += product.SupplyPrice * float64(products[i].Quantity)
		totalRetailPrice += product.RetailPrice * float64(products[i].Quantity)
	}
	order.SupplyPrice = totalSupplyPrice
	order.RetailPrice = totalSupplyPrice

	//create order
	orderID, err := s.orderRepository.CreateOrderInTransaction(order, tx)
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
		_, err := s.orderRepository.CreateOrderItemInTransaction(orderItem, tx)
		if err != nil {
			return models.Order{}, CustomError{
				StatusCode: http.StatusInternalServerError,
				Message:    "Error on creating order items",
				Err:        err,
			}
		}
	}

	if err != nil {
		return models.Order{}, CustomError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error committing transaction",
			Err:        err,
		}
	}

	return order, CustomError{}
}

type ChanelResult struct {
	Order models.OrderRes
	Error CustomError
}

func (s *OrderService) GetOrder(orderID int) (models.OrderRes, CustomError) {
	var readyOrder models.OrderRes
	var error CustomError
	//create chanel
	orderCh := make(chan models.Order, 1)
	orderItemsCh := make(chan []models.OrderItem, 1)
	var wg sync.WaitGroup

	wg.Add(2)
	//Get ORDER
	go func() {
		defer wg.Done()
		order, err := s.orderRepository.GetOrder(orderID)
		if err != nil {

			if err == sql.ErrNoRows {
				error = CustomError{
					StatusCode: http.StatusNotFound,
					Message:    "No order found with the given ID",
					Err:        err,
				}
			}
			error = CustomError{
				StatusCode: http.StatusInternalServerError,
				Message:    "Error on getting order",
				Err:        err,
			}
			return
		}
		orderCh <- order
	}()
	//Get Order Items
	go func() {
		defer wg.Done()
		orderItems, err := s.orderRepository.GetOrderItems(orderID)
		if err != nil {
			error = CustomError{
				StatusCode: http.StatusInternalServerError,
				Message:    "Error on getting order items",
				Err:        err,
			}
			return
		}
		orderItemsCh <- orderItems
	}()

	wg.Wait()
	close(orderCh)
	close(orderItemsCh)
	for order := range orderCh {
		readyOrder.Order = order
	}

	for orderItems := range orderItemsCh {
		readyOrder.Products = orderItems
	}

	return readyOrder, error
}

func (s *OrderService) GetOrders() ([]models.OrderRes, CustomError) {
	// TODO: return count also, remove order items
	var resOrders []models.OrderRes
	orders, err := s.orderRepository.GetOrders("")
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

	return CustomError{}
}

func (s *OrderService) ChangeStatus(orderID int, status string) (models.Order, error) {
	var order models.Order
	order.Status = status
	updatedOrder, err := s.orderRepository.UpdateOrder(orderID, order)
	if err != nil {
		return models.Order{}, err
	}
	return updatedOrder, nil
}

func (s *OrderService) GetPaidOrders() (models.OrderPaid, CustomError) {
	var paidOrders models.OrderPaid
	allOrders, err := s.orderRepository.GetOrders("paid")
	if err != nil {
		return models.OrderPaid{}, CustomError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error on getting order",
			Err:        err,
		}
	}

	var retailPrices float64
	var supplyPrices float64
	for i := 0; i < len(allOrders); i++ {
		retailPrices += allOrders[i].RetailPrice
		supplyPrices += allOrders[i].SupplyPrice
	}
	paidOrders.Orders = allOrders
	paidOrders.NumberOfOrders = len(allOrders)
	paidOrders.RetailPrices = retailPrices
	paidOrders.SupplyPrices = supplyPrices

	return paidOrders, CustomError{}
}
