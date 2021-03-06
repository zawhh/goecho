package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/zawhh/goecho/src/controller"
	"github.com/zawhh/goecho/src/model"
	"gopkg.in/go-playground/validator.v9"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	//e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	//	Format: "method=${method}, uri=${uri}, status=${status}\n",
	//}))

	v := validator.New()
	v.RegisterValidation("carBrand", model.CustomCarValidator)
	e.Validator = &model.CustomValidator{Validator: v}
	// Routes
	e.GET("/", controller.Hello)
	e.POST("/users", controller.User)
	e.POST("/cars", controller.Car)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
