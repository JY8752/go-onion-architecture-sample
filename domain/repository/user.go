package repository

import model "github.com/JY8752/go-onion-architecture-sample/domain/model"

type UserRepository interface {
	Create(string) (model.UserId, error)
	Get(model.UserId) (model.User, error)
}
