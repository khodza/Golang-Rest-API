package service_test

import (
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
