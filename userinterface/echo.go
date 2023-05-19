package ui

import (
	"github.com/JY8752/go-onion-architecture-sample/registory"
	handler "github.com/JY8752/go-onion-architecture-sample/userinterface/handler/user"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ApiClient struct {
	client    *echo.Echo
	registory registory.Registory
}

func NewApiClient(registory registory.Registory) *ApiClient {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return &ApiClient{
		client:    e,
		registory: registory,
	}
}

func (a *ApiClient) RegisterRoute() {
	// user
	handler.UserHandler(a.client, a.registory)

	// todo

}

func (a *ApiClient) Start() {
	a.client.Logger.Fatal(a.client.Start(":8080"))
}
