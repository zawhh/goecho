package model

import "gopkg.in/go-playground/validator.v9"

type Car struct {
	Brand string `json:"brand" validate:"required,carBrand"`
	Miles int    `json:"miles" validate:"required"`
}

// Structure
func CustomCarValidator(fl validator.FieldLevel) bool {

	if fl.Field().String() == "invalid" {
		return false
	}

	return true
}
