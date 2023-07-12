package services

import (
	"khodza/rest-api/internal/app/models"
	"khodza/rest-api/internal/app/repositories"
	"sync"
)

type OrderServiceInterface interface {
	CreateOrder(newOrder models.OrderReq) (models.Order, error)
	GetOrder(orderID int) (models.OrderRes, error)
	GetOrders() (models.AllOrdersRes, error)
	UpdateOrder(orderID int, newOrder models.OrderReq) (models.Order, error)
	DeleteOrder(orderID int) error
	ChangeStatus(orderID int, status string) (models.Order, error)
	GetPaidOrders() (models.OrderPaid, error)
	GetOrderItems(orderId int) ([]models.OrderItem, error)
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

func (s *OrderService) CreateOrder(newOrder models.OrderReq) (models.Order, error) {
	//Start transaction
	tx, err := s.orderRepository.BeginTransaction()
	if err != nil {
		return models.Order{}, err
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
		var product models.Product
		product, err = s.productService.GetProduct(products[i].ProductID)
		if err != nil {
			return models.Order{}, err
		}

		totalSupplyPrice += product.SupplyPrice * float64(products[i].Quantity)
		totalRetailPrice += product.RetailPrice * float64(products[i].Quantity)
	}
	order.SupplyPrice = totalSupplyPrice
	order.RetailPrice = totalRetailPrice

	//create order
	orderID, err := s.orderRepository.CreateOrder(tx, order)
	if err != nil {
		return models.Order{}, err
	}
	order.ID = orderID
	//create order items

	for i := 0; i < len(products); i++ {
		var orderItem models.OrderItem
		orderItem.OrderID = orderID
		orderItem.ProductID, orderItem.Quantity = products[i].ProductID, products[i].Quantity
		_, err = s.orderRepository.CreateOrderItem(tx, orderItem)
		if err != nil {
			return models.Order{}, err
		}
	}

	return order, nil
}

func (s *OrderService) GetOrder(orderID int) (models.OrderRes, error) {
	var readyOrder models.OrderRes
	var errorAny error
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
			errorAny = err
			return
		}
		orderCh <- order
	}()
	//Get Order Items
	go func() {
		defer wg.Done()
		orderItems, err := s.orderRepository.GetOrderItems(orderID)
		if err != nil {
			errorAny = err
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
	if errorAny != nil {
		return models.OrderRes{}, errorAny
	}

	return readyOrder, nil
}

func (s *OrderService) GetOrders() (models.AllOrdersRes, error) {
	var allOrderRes models.AllOrdersRes
	orders, err := s.orderRepository.GetOrders("")
	if err != nil {
		return models.AllOrdersRes{}, err
	}

	orderCount, err := s.orderRepository.GetOrdersCount()
	if err != nil {
		return models.AllOrdersRes{}, err
	}

	allOrderRes.Orders = orders
	allOrderRes.Count = orderCount

	return allOrderRes, nil
}

func (s *OrderService) GetOrderItems(orderId int) ([]models.OrderItem, error) {

	orderItems, err := s.orderRepository.GetOrderItems(orderId)
	if err != nil {
		return []models.OrderItem{}, err
	}
	return orderItems, err

}

func (s *OrderService) UpdateOrder(orderID int, newOrder models.OrderReq) (models.Order, error) {
	tx, err := s.orderRepository.BeginTransaction()

	if err != nil {
		return models.Order{}, err
	}

	//Rollback
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var updatedOrderTemp models.Order
	updatedOrderTemp.UserID = newOrder.UserID
	products := newOrder.Products

	//Calculate prices
	if len(products) != 0 {
		var totalSupplyPrice float64
		var totalRetailPrice float64
		for i := 0; i < len(products); i++ {
			product, err := s.productService.GetProduct(products[i].ProductID)
			if err != nil {
				return models.Order{}, err
			}

			totalSupplyPrice += product.SupplyPrice * float64(products[i].Quantity)
			totalRetailPrice += product.RetailPrice * float64(products[i].Quantity)
		}
		updatedOrderTemp.SupplyPrice = totalSupplyPrice
		updatedOrderTemp.RetailPrice = totalRetailPrice
	}

	//Updating user
	updatedOrder, err := s.orderRepository.UpdateOrder(tx, orderID, updatedOrderTemp)
	if err != nil {
		return models.Order{}, err
	}

	//Update order items
	if len(products) != 0 {
		//delete items
		var err error
		err = s.orderRepository.DeleteOrderItems(tx, orderID)
		if err != nil {
			return models.Order{}, err
		}

		for i := 0; i < len(products); i++ {
			//create items
			var orderItem models.OrderItem
			orderItem.OrderID = orderID
			orderItem.ProductID, orderItem.Quantity = products[i].ProductID, products[i].Quantity
			_, err = s.orderRepository.CreateOrderItem(tx, orderItem)
			if err != nil {
				return models.Order{}, err
			}
		}
	}

	return updatedOrder, err

}

func (s *OrderService) DeleteOrder(orderID int) error {
	//delete order
	err := s.orderRepository.DeleteOrder(orderID)
	if err != nil {
		return err
	}

	return nil
}

func (s *OrderService) ChangeStatus(orderID int, status string) (models.Order, error) {
	var order models.Order
	order.Status = status
	updatedOrder, err := s.orderRepository.UpdateOrder(nil, orderID, order)
	if err != nil {
		return models.Order{}, err
	}
	return updatedOrder, nil
}

func (s *OrderService) GetPaidOrders() (models.OrderPaid, error) {
	var paidOrders models.OrderPaid
	allOrders, err := s.orderRepository.GetOrders("paid")
	if err != nil {
		return models.OrderPaid{}, err
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

	return paidOrders, err
}
