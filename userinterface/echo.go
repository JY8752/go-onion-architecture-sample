package ui

import (
	"github.com/JY8752/go-onion-architecture-sample/registory"
	"github.com/labstack/echo/v4"
)

type ApiClient struct {
	client    *echo.Echo
	registory registory.Registory
}

func NewApiClient(registory registory.Registory) *ApiClient {
	return &ApiClient{
		client:    echo.New(),
		registory: registory,
	}
}

func (a *ApiClient) RegisterRoute() {
	a.client.POST("/user", func(c echo.Context) error {
		a.registory.UserService().Create()
		return c.String(200, "Hello World")
	})
}

func (a *ApiClient) Start() error {
	return a.client.Start(":8080")
}
