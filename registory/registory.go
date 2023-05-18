package registory

import (
	service "github.com/JY8752/go-onion-architecture-sample/application/service/user"
	repository "github.com/JY8752/go-onion-architecture-sample/domain/repository/user"
	db "github.com/JY8752/go-onion-architecture-sample/infrastructure"
	infrastructure "github.com/JY8752/go-onion-architecture-sample/infrastructure/repository/user"
)

type Registory interface {
	UserRep() repository.UserRepository
	UserService() service.UserService
}

type registory struct {
	dbClient *db.DBClient
}

func NewRegistory(db *db.DBClient) Registory {
	return &registory{
		dbClient: db,
	}
}

func (r *registory) UserRep() repository.UserRepository {
	return infrastructure.NewUserRepository(r.dbClient)
}

func (r *registory) UserService() service.UserService {
	return service.NewUserService(r.UserRep())
}
