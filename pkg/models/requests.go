package models

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tonoy30/echo-go/pkg/domain"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func ValidateRegisterRequest(w http.ResponseWriter, r *http.Request) (*domain.User, *Error) {
	vars := mux.Vars(r)

	if _, hasUsername := vars["username"]; !hasUsername {
		return nil, BindError()
	}
	if _, hasPassword := vars["password"]; !hasPassword {
		return nil, BindError()
	}

	registerRequest := &RegisterRequest{
		Username: vars["username"],
		Password: vars["password"],
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

func ValidateLoginRequest(w http.ResponseWriter, r *http.Request) (*domain.User, *Error) {
	vars := mux.Vars(r)
	if _, hasUsername := vars["username"]; !hasUsername {
		return nil, BindError()
	}
	if _, hasPassword := vars["password"]; !hasPassword {
		return nil, BindError()
	}
	loginRequest := &LoginRequest{
		Username: vars["username"],
		Password: vars["password"],
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
