package validators

import (
	"fmt"
	custom_errors "khodza/rest-api/internal/app/errors"
	"khodza/rest-api/internal/app/models"

	"github.com/go-playground/validator/v10"
)

type UserValidatorInterface interface {
	ValidateUserCreate(user *models.User) error
	ValidateUserUpdate(user *models.User) error
}

type UserValidator struct {
	validate *validator.Validate
}

func NewUserValidator() UserValidatorInterface {
	return &UserValidator{
		validate: validator.New(),
	}
}

func (v *UserValidator) ValidateUserCreate(user *models.User) error {
	err := v.validate.Struct(user)
	if err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, fmt.Sprintf("%s is %s", err.Field(), err.Tag()))
		}

		return fmt.Errorf("%e : %v", custom_errors.ErrValidation, validationErrors)
	}

	return nil
}

func (v *UserValidator) ValidateUserUpdate(user *models.User) error {
	if user.Email == "" {
		return nil
	}
	err := v.validate.Var(user.Email, "email")
	if err != nil {
		return fmt.Errorf("%e : %s is %s", custom_errors.ErrValidation, "email", err.Error())
	}

	return nil
}
