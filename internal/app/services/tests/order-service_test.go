package service_test

import (
	"errors"
	"fmt"
	custom_errors "khodza/rest-api/internal/app/errors"
	"khodza/rest-api/internal/app/models"
	"khodza/rest-api/internal/app/services"
	"khodza/rest-api/internal/app/services/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// Unit test for CreateOrder function

func TestCreateOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	orderRepository := mocks.NewMockOrderRepositoryInterface(ctrl)
	productService := mocks.NewMockProductServiceInterface(ctrl)

	orderService := services.NewOrderService(orderRepository, productService)

	newOrder := models.OrderReq{
		UserID: 1,
		Products: []models.OrderProduct{
			{ProductID: 1, Quantity: 2},
			{ProductID: 2, Quantity: 3},
		},
	}

	// Create a mock transaction
	mockTx := mocks.NewMockTransaction(ctrl)
	mockTx.EXPECT().Commit().Return(nil)
	orderRepository.EXPECT().BeginTransaction().Return(mockTx, nil)

	// Mock the necessary repository and service calls
	productService.EXPECT().GetProduct(1).Return(models.Product{SupplyPrice: 10.0, RetailPrice: 20.0}, nil)
	productService.EXPECT().GetProduct(2).Return(models.Product{SupplyPrice: 15.0, RetailPrice: 25.0}, nil)
	orderRepository.EXPECT().CreateOrder(gomock.Any(), gomock.Any()).Return(1, nil).AnyTimes()
	orderRepository.EXPECT().CreateOrderItem(gomock.Any(), gomock.Any()).Return(1, nil).AnyTimes()

	// Call the function being tested
	result, err := orderService.CreateOrder(newOrder)

	// Assert the expected result
	expectedOrder := models.Order{
		ID:          1,
		UserID:      1,
		Status:      "pending",
		SupplyPrice: 65.0,
		RetailPrice: 115.0,
	}
	assert.NoError(t, err)
	assert.Equal(t, expectedOrder, result)
}

func TestCreateOrder_GetProductErr(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOrderRepo := mocks.NewMockOrderRepositoryInterface(ctrl)
	mockProdServ := mocks.NewMockProductServiceInterface(ctrl)
	// Create a mock transaction
	mockTx := mocks.NewMockTransaction(ctrl)
	mockTx.EXPECT().Rollback().Return(nil)
	mockOrderRepo.EXPECT().BeginTransaction().Return(mockTx, nil)

	mockProdServ.EXPECT().GetProduct(gomock.Any()).AnyTimes().Return(models.Product{}, custom_errors.ErrProductNotFound)

	service := services.NewOrderService(mockOrderRepo, mockProdServ)

	newOrder := models.OrderReq{
		UserID: 1,
		Products: []models.OrderProduct{
			{ProductID: 1, Quantity: 2},
			{ProductID: 2, Quantity: 3},
		},
	}
	order, err := service.CreateOrder(newOrder)

	assert.Error(t, err)
	assert.Equal(t, custom_errors.ErrProductNotFound, err)
	assert.Equal(t, models.Order{}, order)
}

func TestCreateOrder_CreateOrderRepoErr(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOrderRepo := mocks.NewMockOrderRepositoryInterface(ctrl)
	mockProdServ := mocks.NewMockProductServiceInterface(ctrl)
	// Create a mock transaction
	mockTx := mocks.NewMockTransaction(ctrl)
	mockTx.EXPECT().Rollback().Return(nil)
	mockOrderRepo.EXPECT().BeginTransaction().Return(mockTx, nil)

	mockProdServ.EXPECT().GetProduct(gomock.All()).AnyTimes().Return(models.Product{SupplyPrice: 10.0, RetailPrice: 20.0}, nil)
	mockOrderRepo.EXPECT().CreateOrder(mockTx, gomock.Any()).Return(0, errors.New(""))

	service := services.NewOrderService(mockOrderRepo, mockProdServ)

	newOrder := models.OrderReq{
		UserID: 1,
		Products: []models.OrderProduct{
			{ProductID: 1, Quantity: 2},
			{ProductID: 2, Quantity: 3},
		},
	}
	order, err := service.CreateOrder(newOrder)

	assert.Error(t, err)
	assert.Equal(t, errors.New(""), err)
	assert.Equal(t, models.Order{}, order)
}

func TestCreateOrder_CreateOrderItemsRepo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOrderRepo := mocks.NewMockOrderRepositoryInterface(ctrl)
	mockProdServ := mocks.NewMockProductServiceInterface(ctrl)
	// Create a mock transaction
	mockTx := mocks.NewMockTransaction(ctrl)
	mockTx.EXPECT().Rollback().Return(nil)
	// mockTx.EXPECT().Commit().Return(nil)

	mockOrderRepo.EXPECT().BeginTransaction().Return(mockTx, nil)

	mockProdServ.EXPECT().GetProduct(gomock.All()).AnyTimes().Return(models.Product{SupplyPrice: 10.0, RetailPrice: 20.0}, nil)
	mockOrderRepo.EXPECT().CreateOrder(mockTx, gomock.Any()).Return(1, nil)

	mockOrderRepo.EXPECT().CreateOrderItem(mockTx, gomock.Any()).Return(0, custom_errors.ErrCreateOrderItems).AnyTimes()

	service := services.NewOrderService(mockOrderRepo, mockProdServ)

	newOrder := models.OrderReq{
		UserID: 1,
		Products: []models.OrderProduct{
			{ProductID: 1, Quantity: 2},
			{ProductID: 2, Quantity: 3},
		},
	}
	order, err := service.CreateOrder(newOrder)

	assert.Error(t, err)
	fmt.Println(err)
	assert.Equal(t, custom_errors.ErrCreateOrderItems, err)
	assert.Equal(t, models.Order{}, order)
}

// Unit test for GetOrder function

func TestGetOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOrderRepo := mocks.NewMockOrderRepositoryInterface(ctrl)
	mockProdServ := mocks.NewMockProductServiceInterface(ctrl)

	order := models.Order{ID: 1, UserID: 2, SupplyPrice: 20.0, RetailPrice: 30.0, Status: "pending"}

	orderItems := []models.OrderItem{{ID: 1, OrderID: 1, ProductID: 1, Quantity: 2}, {ID: 2, ProductID: 2, Quantity: 1}}
	mockOrderRepo.EXPECT().GetOrder(gomock.All()).Return(order, nil)
	mockOrderRepo.EXPECT().GetOrderItems(gomock.All()).Return(orderItems, nil)

	service := services.NewOrderService(mockOrderRepo, mockProdServ)
	readyOrder, err := service.GetOrder(1)

	assert.NoError(t, err)
	assert.Equal(t, models.OrderRes{Order: order, Products: orderItems}, readyOrder)

}
