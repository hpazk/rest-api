package user

import (
	"github.com/hpazk/rest-api/helper"
	"github.com/labstack/echo/v4"
)

type UserRoutes struct{}

func (r UserRoutes) Route() []helper.Route {
	userRepo := NewRepository(&UsersStorage{})
	userService := NewServices(userRepo)
	userHandler := NewHandler(userService)

	return []helper.Route{
		{
			Method:  echo.POST,
			Path:    "/users",
			Handler: userHandler.UserRegistration,
		},
	}
}
