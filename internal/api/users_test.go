package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	usersmock "github.com/tonoy30/echo-go/pkg/mocks/data/users"
	"github.com/tonoy30/echo-go/pkg/services"
	"github.com/tonoy30/echo-go/pkg/settings"
)

func TestRegisterAccount_UsernameExists(t *testing.T) {
	// test data
	testUser := `{"username": "user1", "password": "password23", ""}`

	// echo setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(testUser))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// internal setup
	settings := &settings.Settings{}
	mockData := usersmock.NewMockDataStore()
	userService := services.NewUserService(settings, mockData)

	mockApp := App{
		userService: userService,
	}
	// act
	mockApp.Register(c)
	// assert
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
