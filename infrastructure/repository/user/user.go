package infrastructure

import (
	repository "github.com/JY8752/go-onion-architecture-sample/domain/repository/user"
	db "github.com/JY8752/go-onion-architecture-sample/infrastructure"
)

type userRepository struct {
	dbClient *db.DBClient
}

func NewUserRepository(db *db.DBClient) repository.UserRepository {
	return &userRepository{
		dbClient: db,
	}
}

func (u *userRepository) Create() {
	// TODO
}
