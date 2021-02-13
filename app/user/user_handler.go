package user

import (
	"net/http"

	"github.com/hpazk/rest-api/auth"
	"github.com/hpazk/rest-api/helper"
	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userService Services
	authService auth.AuthService
}

func NewHandler(userService Services, authService auth.AuthService) *userHandler {
	return &userHandler{userService, authService}
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

	auth_token, err := h.authService.GetAccessToken(newUser.ID)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusInternalServerError, "error", err.Error(), nil)

		return c.JSON(http.StatusInternalServerError, response)
	}

	userData := UserResponseFormatter(newUser, auth_token)

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

	auth_token, err := h.authService.GetAccessToken(userAuth.ID)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusInternalServerError, "error", err.Error(), nil)

		return c.JSON(http.StatusInternalServerError, response)
	}

	userData := UserResponseFormatter(userAuth, auth_token)

	response := helper.ResponseFormatter(http.StatusOK, "success", "user authenticated", userData)

	return c.JSON(http.StatusOK, response)

}

func (h *userHandler) SecretResource(c echo.Context) error {
	response := helper.M{"message": "this is secret route"}

	return c.JSON(http.StatusOK, response)
}
