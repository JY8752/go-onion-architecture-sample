package common

import (
	"strconv"

	"github.com/JY8752/go-onion-architecture-sample/domain/model"
)

func GetTodoId(str string) (model.TodoId, error) {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}

	return model.TodoId(i), nil
}
