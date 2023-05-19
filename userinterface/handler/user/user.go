package handler

import (
	"strconv"

	model "github.com/JY8752/go-onion-architecture-sample/domain/model/user"
	"github.com/JY8752/go-onion-architecture-sample/registory"
	"github.com/JY8752/go-onion-architecture-sample/userinterface/request"
	"github.com/labstack/echo/v4"
)

func UserHandler(client *echo.Echo, registory registory.Registory) {
	client.POST("/users", func(c echo.Context) error {
		var r request.CreateUserRequest
		if err := c.Bind(&r); err != nil {
			return err
		}

		id, err := registory.UserService().Create(r.Name)
		if err != nil {
			return err
		}

		return c.JSON(200, id)
	})

	client.GET("/user/:id", func(c echo.Context) error {
		i, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return err
		}

		id := model.UserId(i)

		user, err := registory.UserService().Get(id)
		if err != nil {
			return err
		}

		return c.JSON(200, user)
	})
}
