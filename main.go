package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type M map[string]interface{}

func main() {
	fmt.Println("running...")

	e := echo.New()

	e.GET("/users", func(c echo.Context) error {
		response := M{"message": "success"}

		return c.JSON(http.StatusOK, response)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
