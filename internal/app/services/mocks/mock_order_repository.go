// Code generated by MockGen. DO NOT EDIT.
// Source: internal/app/repositories/order-repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	models "khodza/rest-api/internal/app/models"
	db "khodza/rest-api/pkg/db"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockOrderRepositoryInterface is a mock of OrderRepositoryInterface interface.
type MockOrderRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockOrderRepositoryInterfaceMockRecorder
}

// MockOrderRepositoryInterfaceMockRecorder is the mock recorder for MockOrderRepositoryInterface.
type MockOrderRepositoryInterfaceMockRecorder struct {
	mock *MockOrderRepositoryInterface
}

// NewMockOrderRepositoryInterface creates a new mock instance.
func NewMockOrderRepositoryInterface(ctrl *gomock.Controller) *MockOrderRepositoryInterface {
	mock := &MockOrderRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockOrderRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderRepositoryInterface) EXPECT() *MockOrderRepositoryInterfaceMockRecorder {
	return m.recorder
}

// BeginTransaction mocks base method.
func (m *MockOrderRepositoryInterface) BeginTransaction() (db.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginTransaction")
	ret0, _ := ret[0].(db.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginTransaction indicates an expected call of BeginTransaction.
func (mr *MockOrderRepositoryInterfaceMockRecorder) BeginTransaction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginTransaction", reflect.TypeOf((*MockOrderRepositoryInterface)(nil).BeginTransaction))
}

// CreateOrder mocks base method.
func (m *MockOrderRepositoryInterface) CreateOrder(tx db.Transaction, order models.Order) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", tx, order)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrder indicates an expected call of CreateOrder.
func (mr *MockOrderRepositoryInterfaceMockRecorder) CreateOrder(tx, order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockOrderRepositoryInterface)(nil).CreateOrder), tx, order)
}

// CreateOrderItem mocks base method.
func (m *MockOrderRepositoryInterface) CreateOrderItem(tx db.Transaction, orderItem models.OrderItem) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrderItem", tx, orderItem)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrderItem indicates an expected call of CreateOrderItem.
func (mr *MockOrderRepositoryInterfaceMockRecorder) CreateOrderItem(tx, orderItem interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrderItem", reflect.TypeOf((*MockOrderRepositoryInterface)(nil).CreateOrderItem), tx, orderItem)
}

// DeleteOrder mocks base method.
func (m *MockOrderRepositoryInterface) DeleteOrder(orderID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrder", orderID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOrder indicates an expected call of DeleteOrder.
func (mr *MockOrderRepositoryInterfaceMockRecorder) DeleteOrder(orderID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrder", reflect.TypeOf((*MockOrderRepositoryInterface)(nil).DeleteOrder), orderID)
}

// DeleteOrderItems mocks base method.
func (m *MockOrderRepositoryInterface) DeleteOrderItems(tx db.Transaction, orderID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrderItems", tx, orderID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOrderItems indicates an expected call of DeleteOrderItems.
func (mr *MockOrderRepositoryInterfaceMockRecorder) DeleteOrderItems(tx, orderID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrderItems", reflect.TypeOf((*MockOrderRepositoryInterface)(nil).DeleteOrderItems), tx, orderID)
}

// GetOrder mocks base method.
func (m *MockOrderRepositoryInterface) GetOrder(orderID int) (models.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrder", orderID)
	ret0, _ := ret[0].(models.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrder indicates an expected call of GetOrder.
func (mr *MockOrderRepositoryInterfaceMockRecorder) GetOrder(orderID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrder", reflect.TypeOf((*MockOrderRepositoryInterface)(nil).GetOrder), orderID)
}

// GetOrderItems mocks base method.
func (m *MockOrderRepositoryInterface) GetOrderItems(orderID int) ([]models.OrderItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderItems", orderID)
	ret0, _ := ret[0].([]models.OrderItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderItems indicates an expected call of GetOrderItems.
func (mr *MockOrderRepositoryInterfaceMockRecorder) GetOrderItems(orderID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderItems", reflect.TypeOf((*MockOrderRepositoryInterface)(nil).GetOrderItems), orderID)
}

// GetOrders mocks base method.
func (m *MockOrderRepositoryInterface) GetOrders(status string) ([]models.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrders", status)
	ret0, _ := ret[0].([]models.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrders indicates an expected call of GetOrders.
func (mr *MockOrderRepositoryInterfaceMockRecorder) GetOrders(status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrders", reflect.TypeOf((*MockOrderRepositoryInterface)(nil).GetOrders), status)
}

// GetOrdersCount mocks base method.
func (m *MockOrderRepositoryInterface) GetOrdersCount() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrdersCount")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrdersCount indicates an expected call of GetOrdersCount.
func (mr *MockOrderRepositoryInterfaceMockRecorder) GetOrdersCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrdersCount", reflect.TypeOf((*MockOrderRepositoryInterface)(nil).GetOrdersCount))
}

// UpdateOrder mocks base method.
func (m *MockOrderRepositoryInterface) UpdateOrder(tx db.Transaction, orderID int, newOrder models.Order) (models.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrder", tx, orderID, newOrder)
	ret0, _ := ret[0].(models.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOrder indicates an expected call of UpdateOrder.
func (mr *MockOrderRepositoryInterfaceMockRecorder) UpdateOrder(tx, orderID, newOrder interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrder", reflect.TypeOf((*MockOrderRepositoryInterface)(nil).UpdateOrder), tx, orderID, newOrder)
}
