package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tonoy30/echo-go/pkg/domain"
	usersmock "github.com/tonoy30/echo-go/pkg/mocks/data/users"
	"github.com/tonoy30/echo-go/pkg/settings"
)

func TestCreateAccount_UserExists(t *testing.T) {
	settings := &settings.Settings{}
	userProviderMock := &usersmock.UserProviderMock{}
	userProviderMock.UsernameExistsMock = func(username string) (bool, error) {
		return true, nil
	}
	userService := NewUserService(settings, userProviderMock)
	newUser := &domain.User{
		Username: "test",
		Password: "test",
	}

	response := userService.CreateAccount(newUser)
	assert.Equal(t, "USERNAME_TAKEN", response.Name)
}
