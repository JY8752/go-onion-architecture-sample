package repository

import model "github.com/JY8752/go-onion-architecture-sample/domain/model/user"

type UserRepository interface {
	Create(string) (model.UserId, error)
}
