package ui

import (
	"github.com/JY8752/go-onion-architecture-sample/registory"
	"github.com/JY8752/go-onion-architecture-sample/userinterface/request"
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
	// user
	a.client.POST("/user", func(c echo.Context) error {
		r := new(request.CreateUserRequest)
		if err := c.Bind(r); err != nil {
			return err
		}

		a.registory.UserService().Create(r.Name)
		return c.String(200, "Hello World")
	})

	// todo

}

func (a *ApiClient) Start() error {
	return a.client.Start(":8080")
}
