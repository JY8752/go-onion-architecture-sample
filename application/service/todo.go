package service

import (
	"github.com/JY8752/go-onion-architecture-sample/domain/model"
	"github.com/JY8752/go-onion-architecture-sample/domain/repository"
)

type TodoService interface {
	Create(model.UserId, model.Title, model.Description) (model.TodoId, error)
	List(model.UserId) ([]model.Todo, error)
	Delete(model.TodoId) error
}

type todoService struct {
	todoRep repository.TodoRepository
}

func NewTodoService(todoRep repository.TodoRepository) TodoService {
	return &todoService{
		todoRep: todoRep,
	}
}

func (t *todoService) Create(userId model.UserId, title model.Title, description model.Description) (model.TodoId, error) {
	return t.todoRep.Create(userId, title, description)
}

func (t *todoService) List(userId model.UserId) ([]model.Todo, error) {
	return t.todoRep.List(userId)
}

func (t *todoService) Delete(todoId model.TodoId) error {
	return t.todoRep.Delete(todoId)
}
