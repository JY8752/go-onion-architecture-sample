package repository

//go:generate mockgen -source=$GOFILE -destination=../../mocks/mock_$GOPACKAGE/$GOFILE

import model "github.com/JY8752/go-onion-architecture-sample/domain/model"

type UserRepository interface {
	Create(string) (model.UserId, error)
	Get(model.UserId) (model.User, error)
}
