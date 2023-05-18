package infrastructure

import (
	"time"

	model "github.com/JY8752/go-onion-architecture-sample/domain/model/user"
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

func (u *userRepository) Create(name string) (model.UserId, error) {
	stmt, err := u.dbClient.Client.Prepare("INSERT INTO users (name, created_at) VALUES (?, ?)")
	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(name, time.Now())
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return model.UserId(id), nil
}
