package main

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type M map[string]interface{}

type User struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

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
	e.POST("/users", func(c echo.Context) error {
		user := new(User)

		if err := c.Bind(user); err != nil {
			response := M{"message": "error"}

			return c.JSON(http.StatusBadRequest, response)
		}

		if err := c.Validate(user); err != nil {
			response := M{
				"message": "error",
				"errors":  err.Error(),
			}

			return c.JSON(http.StatusBadRequest, response)
		}

		response := M{
			"message": "success",
			"data":    user,
		}

		return c.JSON(http.StatusOK, response)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
