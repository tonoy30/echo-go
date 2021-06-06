package services

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/tonoy30/echo-go/pkg/data"
	"github.com/tonoy30/echo-go/pkg/domain"
	"github.com/tonoy30/echo-go/pkg/models"
	"github.com/tonoy30/echo-go/pkg/settings"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	CreateAccount(user *domain.User) *models.Error
	Login(user *domain.User) (string, *models.Error)
}

type UserService struct {
	userProvider data.IUserProvider
	settings     *settings.Settings
}

func NewUserService(settings *settings.Settings, userProvider data.IUserProvider) IUserService {
	return &UserService{
		userProvider: userProvider,
		settings:     settings,
	}
}

func (u UserService) CreateAccount(user *domain.User) *models.Error {
	userExists, err := u.userProvider.UsernameExists(user.Username)
	if err != nil {
		return &models.Error{
			Code:    500,
			Name:    "SERVER_ERROR",
			Message: "something went wrong",
			Error:   err.Error(),
		}
	}
	if userExists {

		return &models.Error{
			Code:    400,
			Name:    "USERNAME_TAKEN",
			Message: "Username already exists",
		}
	}
	user.ID = primitive.NewObjectID()
	hash, err := hashPassword(user.Password)

	if err != nil {
		return &models.Error{
			Code:    500,
			Name:    "SERVER_ERROR",
			Message: "Something went wrong",
			Error:   err.Error(),
		}

	}

	user.Password = hash
	err = u.userProvider.CreateAccount(user)

	if err != nil {
		return &models.Error{
			Code:    500,
			Name:    "SERVER_ERROR",
			Message: "Something went wrong",
			Error:   err.Error(),
		}
	}

	return nil
}

func (u UserService) Login(user *domain.User) (string, *models.Error) {
	userFound, err := u.userProvider.FindByUsername(user.Username)
	if err != nil {
		return "", &models.Error{
			Code:    500,
			Name:    "SERVER_ERROR",
			Message: "Something went wrong",
			Error:   err.Error(),
		}
	}

	if userFound == nil {
		return "", &models.Error{
			Code:    400,
			Name:    "INVALID_LOGIN",
			Message: "Your username or password is invalid.",
		}
	}

	err = comparePasswordWithHash(user.Password, userFound.Password)

	if err != nil {
		return "", &models.Error{
			Code:    400,
			Name:    "INVALID_LOGIN",
			Message: "Your username or password is invalid.",
		}
	}

	token, err := u.createJwtToken(userFound.ID.Hex())

	if err != nil {
		return "", &models.Error{
			Code:    500,
			Name:    "SERVER_ERROR",
			Message: "Something went wrong",
			Error:   err.Error(),
		}
	}

	return token, nil
}

func hashPassword(password string) (string, error) {
	passwordBytes := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(passwordBytes, 12)
	if err != nil {
		return "", errors.Wrap(err, "Error creating password")
	}

	return string(hash), err
}

func comparePasswordWithHash(password string, hash string) error {
	passwordBytes := []byte(password)
	hashBytes := []byte(hash)

	err := bcrypt.CompareHashAndPassword(hashBytes, passwordBytes)

	return errors.Wrap(err, "error comparing password hash")
}

func (u UserService) createJwtToken(userId string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	expiresIn, err := strconv.ParseInt(u.settings.JWTExpires, 10, 64)

	if err != nil {
		return "", errors.Wrap(err, "Error parsing int")
	}
	expiration := time.Duration(int64(time.Minute) * expiresIn)
	claims := token.Claims.(jwt.MapClaims)

	claims["id"] = userId

	claims["exp"] = time.Now().Add(expiration).Unix()

	t, err := token.SignedString([]byte(u.settings.JWTSecret))

	if err != nil {
		return "", errors.Wrap(err, "Error signing JWT token")
	}

	return t, nil
}
