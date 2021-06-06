package api

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tonoy30/echo-go/pkg/data"
	"github.com/tonoy30/echo-go/pkg/services"
	"github.com/tonoy30/echo-go/pkg/settings"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	userService services.IUserService
	settings    *settings.Settings
	server      *echo.Echo
}

func New(settings *settings.Settings, client *mongo.Client) *App {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())

	// providers
	userProvider := data.NewUserProvider(settings, client)

	// services
	userSvc := services.NewUserService(settings, userProvider)

	return &App{
		userService: userSvc,
		settings:    settings,
		server:      e,
	}
}

func (a App) ConfigureRoutes() {
	a.server.GET("/v1/public/healthy", a.HealthCheck)
	a.server.POST("/v1/public/account/register", a.Register)
	a.server.POST("/v1/public/account/login", a.Login)
}
func (a App) Start() {
	a.ConfigureRoutes()
	log.Println("Listening on port 5050")
	a.server.Start(":5050")
}
