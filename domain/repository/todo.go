package repository

//go:generate mockgen -source=$GOFILE -destination=../../mocks/mock_$GOPACKAGE/$GOFILE

import "github.com/JY8752/go-onion-architecture-sample/domain/model"

type TodoRepository interface {
	Create(model.UserId, model.Title, model.Description) (model.TodoId, error)
	List(model.UserId) ([]model.Todo, error)
	Delete(model.TodoId) error
}
