package validators

import (
	"fmt"
	"khodza/rest-api/internal/app/models"

	"github.com/go-playground/validator/v10"
)

type UserValidator struct {
	validate *validator.Validate
}

func NewUserValidator() *UserValidator {
	return &UserValidator{
		validate: validator.New(),
	}
}

func (v *UserValidator) ValidateUser(user *models.User) error {
	err := v.validate.Struct(user)
	if err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, fmt.Sprintf("%s is %s", err.Field(), err.Tag()))
		}

		return fmt.Errorf("validation failed: %v", validationErrors)
	}

	return nil
}
