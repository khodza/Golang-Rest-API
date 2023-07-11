package custom_errors

import (
	"errors"
	"strings"
)

// 400
var ErrEmailExist = errors.New("email already exists")
var ErrUserNotFound = errors.New("user not found")
var ValidationErr = "validation failed :"

func IsValidationErr(err string) bool {
	return strings.HasPrefix(err, ValidationErr)
}
