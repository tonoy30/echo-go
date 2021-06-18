package usersmock

import (
	"github.com/tonoy30/echo-go/pkg/data"
	"github.com/tonoy30/echo-go/pkg/domain"
)

var mockData = []domain.User{
	{Username: "user1", Password: "password1"},
	{Username: "user2", Password: "password2"},
}

type MockDataStock struct{}

func NewMockDataStore() data.IUserProvider {
	return &MockDataStock{}
}
func (m MockDataStock) CreateAccount(user *domain.User) error {
	mockData = append(mockData, *user)
	return nil
}
func (m MockDataStock) UsernameExists(username string) (bool, error) {
	for user := range mockData {
		if mockData[user].Username == username {
			return true, nil
		}
	}
	return false, nil
}

func (m MockDataStock) FindByUsername(username string) (*domain.User, error) {
	for user := range mockData {
		if mockData[user].Username == username {
			return &mockData[user], nil
		}
	}
	return nil, nil
}

func (m MockDataStock) FindByEmail(email string) (*domain.User, error) {
	for user := range mockData {
		if mockData[user].Email == email {
			return &mockData[user], nil
		}
	}
	return nil, nil
}
