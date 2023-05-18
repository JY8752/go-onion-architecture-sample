package service

import repository "github.com/JY8752/go-onion-architecture-sample/domain/repository/user"

type UserService interface {
	Create()
}

type userService struct {
	userRep repository.UserRepository
}

func NewUserService(userRep repository.UserRepository) UserService {
	return &userService{
		userRep: userRep,
	}
}

func (u *userService) Create() {
	// TODO
}
