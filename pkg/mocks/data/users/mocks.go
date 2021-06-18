package usersmock

import (
	"github.com/tonoy30/echo-go/pkg/domain"
)

type UserProviderMock struct {
	CreateAccountMock  func(user *domain.User) error
	UsernameExistsMock func(username string) (bool, error)
	FindByUsernameMock func(username string) (*domain.User, error)
	FindByEmailMock    func(email string) (*domain.User, error)
}

func (u UserProviderMock) CreateAccount(user *domain.User) error {
	return u.CreateAccountMock(user)
}
func (u UserProviderMock) UsernameExists(username string) (bool, error) {
	return u.UsernameExistsMock(username)
}
func (u UserProviderMock) FindByUsername(username string) (*domain.User, error) {
	return u.FindByUsernameMock(username)
}
func (u UserProviderMock) FindByEmail(email string) (*domain.User, error) {
	return u.FindByEmailMock(email)
}
