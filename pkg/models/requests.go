package models

import (
	"github.com/labstack/echo/v4"
	"github.com/tonoy30/echo-go/pkg/domain"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func ValidateRegisterRequest(c echo.Context) (*domain.User, *Error) {
	registerRequest := new(RegisterRequest)
	if err := c.Bind(registerRequest); err != nil {
		return nil, BindError()
	}
	var validationErrors []string

	if len(registerRequest.Password) < 8 {
		validationErrors = append(validationErrors, "password must be min 8 characters long")
	}
	if len(registerRequest.Username) < 3 {
		validationErrors = append(validationErrors, "username must be min 3 characters long")
	}
	if len(validationErrors) > 0 {
		return nil, ValidationError(validationErrors)
	}

	return &domain.User{
		Username: registerRequest.Username,
		Password: registerRequest.Password,
	}, nil
}

func ValidateLoginRequest(c echo.Context) (*domain.User, *Error) {
	loginRequest := new(LoginRequest)

	if err := c.Bind(loginRequest); err != nil {
		return nil, BindError()
	}
	var validationErrors []string

	if len(loginRequest.Password) < 8 {
		validationErrors = append(validationErrors, "password must be min 8 characters long")
	}
	if len(loginRequest.Username) < 3 {
		validationErrors = append(validationErrors, "username must be min 3 characters long")
	}
	if len(validationErrors) > 0 {
		return nil, ValidationError(validationErrors)
	}

	return &domain.User{
		Username: loginRequest.Username,
		Password: loginRequest.Password,
	}, nil
}
