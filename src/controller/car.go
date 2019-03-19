package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/zawhh/goecho/src/model"
)

func Car(c echo.Context) error {
	u := new(model.Car)
	if err := c.Bind(u); err != nil {
		return err
	}
	if err := c.Validate(u); err != nil {
		//return echo.NewHTTPError(http.StatusBadRequest, "Bad error, "+err.Error())
		return c.JSON(http.StatusBadRequest, model.ValidationError{
			CustomMsg:     "Invalid Payload",
			ValidationMsg: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, u)
}
