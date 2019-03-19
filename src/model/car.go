package model

type Car struct {
	Brand string `json:"brand" validate:"required,carBrand"`
	Miles int    `json:"miles" validate:"required"`
}
