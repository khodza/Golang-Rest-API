package validators

import (
	"fmt"
	custom_errors "khodza/rest-api/internal/app/errors"
	"khodza/rest-api/internal/app/models"

	"github.com/go-playground/validator/v10"
)

type ProductValidatorInterface interface {
	ValidateProduct(product *models.Product) error
}

type ProductValidator struct {
	validate *validator.Validate
}

func NewProductValidator() ProductValidatorInterface {
	return &ProductValidator{
		validate: validator.New(),
	}
}

func (v *ProductValidator) ValidateProduct(product *models.Product) error {
	err := v.validate.Struct(product)
	if err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, fmt.Sprintf("%s is %s", err.Field(), err.Tag()))
		}

		return fmt.Errorf("%e : %v", custom_errors.ErrValidation, validationErrors)
	}

	return nil
}
