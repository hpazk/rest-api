package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/hpazk/rest-api/app/user"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	fmt.Println("running...")

	e := echo.New()

	e.Validator = &CustomValidator{validator: validator.New()}

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.Pre(middleware.RemoveTrailingSlash())

	// e.GET("/users", func(c echo.Context) error {
	// 	response := M{"message": "success"}

	// 	return c.JSON(http.StatusOK, response)
	// })
	userRepo := user.NewRepository(&user.UsersStorage{})
	userService := user.NewServices(userRepo)
	userHandler := user.NewHandler(userService)

	e.POST("/users", userHandler.UserRegistration)

	e.Logger.Fatal(e.Start(":8080"))
}
