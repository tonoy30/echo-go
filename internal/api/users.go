package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tonoy30/echo-go/pkg/models"
)

func (a App) Login(c echo.Context) error {
	user, err := models.ValidateLoginRequest(c)
	if err != nil {
		return c.JSON(err.Code, err)
	}
	token, err := a.userService.Login(user)
	if err != nil {
		return c.JSON(err.Code, err)
	}
	response := &models.LoginResponse{Token: token}
	return c.JSON(http.StatusOK, response)
}
func (a App) Register(c echo.Context) error {
	user, err := models.ValidateRegisterRequest(c)
	if err != nil {
		return c.JSON(err.Code, err)
	}
	err = a.userService.CreateAccount(user)
	if err != nil {
		return c.JSON(err.Code, err)
	}
	return c.String(http.StatusCreated, "")
}
