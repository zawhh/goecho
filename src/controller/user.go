package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/zawhh/goecho/src/model"
	"net/http"
)

func User(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	if err := c.Validate(u); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}