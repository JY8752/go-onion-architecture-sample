package service

import (
	model "github.com/JY8752/go-onion-architecture-sample/domain/model/user"
	repository "github.com/JY8752/go-onion-architecture-sample/domain/repository/user"
)

type UserService interface {
	Create(string) (model.UserId, error)
}

type userService struct {
	userRep repository.UserRepository
}

func NewUserService(userRep repository.UserRepository) UserService {
	return &userService{
		userRep: userRep,
	}
}

func (u *userService) Create(name string) (model.UserId, error) {
	return u.userRep.Create(name)
}