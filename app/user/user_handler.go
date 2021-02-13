package user

import (
	"net/http"

	"github.com/hpazk/rest-api/helper"
	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userService Services
}

func NewHandler(userService Services) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) UserRegistration(c echo.Context) error {
	req := new(RequestUser)

	if err := c.Bind(req); err != nil {

		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "invalid request", nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	if err := c.Validate(req); err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	existEmail := h.userService.CheckExistEmail(*req)
	if existEmail != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", existEmail.Error(), nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	newUser, err := h.userService.CreateUser(*req)
	if err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	userData := UserResponseFormatter(newUser)

	response := helper.ResponseFormatter(http.StatusOK, "success", "user succesfully registered", userData)

	return c.JSON(http.StatusOK, response)
}

func (h *userHandler) UserLogin(c echo.Context) error {
	req := new(RequestUserLogin)

	if err := c.Bind(req); err != nil {

		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "invalid request", nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	if err := c.Validate(req); err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	userAuth, err := h.userService.AuthUser(*req)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusUnauthorized, "error", err.Error(), nil)

		return c.JSON(http.StatusUnauthorized, response)
	}

	userData := UserResponseFormatter(userAuth)

	response := helper.ResponseFormatter(http.StatusOK, "success", "user authenticated", userData)

	return c.JSON(http.StatusOK, response)

}
