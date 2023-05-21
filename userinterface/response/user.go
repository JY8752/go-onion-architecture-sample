package response

import "github.com/JY8752/go-onion-architecture-sample/domain/model"

type CreateUserResponse struct {
	Id model.UserId `json:"id"`
}

type GetUserResponse struct {
	User model.User `json:"user"`
}
