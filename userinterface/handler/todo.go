package handler

import (
	"log"

	"github.com/JY8752/go-onion-architecture-sample/common"
	"github.com/JY8752/go-onion-architecture-sample/domain/model"
	"github.com/JY8752/go-onion-architecture-sample/registory"
	"github.com/JY8752/go-onion-architecture-sample/userinterface/request"
	"github.com/JY8752/go-onion-architecture-sample/userinterface/response"
	"github.com/labstack/echo/v4"
)

func TodoHandler(client *echo.Echo, registory registory.Registory) {
	client.POST("/:userId/todos", func(c echo.Context) error {
		// バリデーション
		_, err := common.GetUserId(c.Param("userId"))
		if err != nil {
			return err
		}

		var r request.CreateTodoRequest
		if err := c.Bind(&r); err != nil {
			return err
		}

		id, err := registory.TodoService().Create(
			model.UserId(r.UserId),
			model.Title(r.Title),
			model.Description(r.Description),
		)
		if err != nil {
			return err
		}

		return c.JSON(200, response.CreateTodoResponse{Id: id})
	})

	client.GET("/:userId/todos", func(c echo.Context) error {
		userId, err := common.GetUserId(c.Param("userId"))
		if err != nil {
			return err
		}

		todos, err := registory.TodoService().List(userId)
		if err != nil {
			log.Printf("err: %s\n", err.Error())
			return c.JSON(404, []model.Todo{})
		}

		return c.JSON(200, response.GetTodosResponse{Todos: todos})
	})

	client.DELETE("/todos/:id", func(c echo.Context) error {
		todoId, err := common.GetTodoId(c.Param("id"))
		if err != nil {
			return err
		}

		err = registory.TodoService().Delete(todoId)
		if err != nil {
			return err
		}

		return c.NoContent(204)
	})
}
