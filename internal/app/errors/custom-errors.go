package custom_errors

import (
	"errors"
	"strings"
)

// payment
var ErrPaymentNotEqual = errors.New("provided cash more or less that needed")

// orders errors
var ErrOrderNotFound = errors.New("order not found")
var ErrOrderItemsNotFound = errors.New("order items not found")

// products errors
var ErrProductCodeExist = errors.New("product already exist")
var ErrProductNotFound = errors.New("product not found")

// users errors
var ErrEmailExist = errors.New("email already exists")
var ErrUserNotFound = errors.New("user not found")

// validation
var ErrValidation = errors.New("validation failed")

func IsValidationErr(err string) bool {
	return strings.HasPrefix(err, ErrValidation.Error())
}
