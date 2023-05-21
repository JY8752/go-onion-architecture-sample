package response

import "github.com/JY8752/go-onion-architecture-sample/domain/model"

type CreateTodoResponse struct {
	Id model.TodoId `json:"id"`
}

type GetTodosResponse struct {
	Todos []model.Todo `json:"todos"`
}
