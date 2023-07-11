// Code generated by MockGen. DO NOT EDIT.
// Source: internal/app/repositories/user-repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	models "khodza/rest-api/internal/app/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserRepositoryInterface is a mock of UserRepositoryInterface interface.
type MockUserRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryInterfaceMockRecorder
}

// MockUserRepositoryInterfaceMockRecorder is the mock recorder for MockUserRepositoryInterface.
type MockUserRepositoryInterfaceMockRecorder struct {
	mock *MockUserRepositoryInterface
}

// NewMockUserRepositoryInterface creates a new mock instance.
func NewMockUserRepositoryInterface(ctrl *gomock.Controller) *MockUserRepositoryInterface {
	mock := &MockUserRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepositoryInterface) EXPECT() *MockUserRepositoryInterfaceMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserRepositoryInterface) CreateUser(user models.User) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserRepositoryInterfaceMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserRepositoryInterface)(nil).CreateUser), user)
}

// DeleteUser mocks base method.
func (m *MockUserRepositoryInterface) DeleteUser(userID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserRepositoryInterfaceMockRecorder) DeleteUser(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserRepositoryInterface)(nil).DeleteUser), userID)
}

// GetUser mocks base method.
func (m *MockUserRepositoryInterface) GetUser(userID int) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", userID)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockUserRepositoryInterfaceMockRecorder) GetUser(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserRepositoryInterface)(nil).GetUser), userID)
}

// GetUsers mocks base method.
func (m *MockUserRepositoryInterface) GetUsers() ([]models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers")
	ret0, _ := ret[0].([]models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockUserRepositoryInterfaceMockRecorder) GetUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockUserRepositoryInterface)(nil).GetUsers))
}

// UpdateUser mocks base method.
func (m *MockUserRepositoryInterface) UpdateUser(userID int, user models.User) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", userID, user)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUserRepositoryInterfaceMockRecorder) UpdateUser(userID, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserRepositoryInterface)(nil).UpdateUser), userID, user)
}