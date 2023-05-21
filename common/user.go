package common

import (
	"strconv"

	"github.com/JY8752/go-onion-architecture-sample/domain/model"
)

func GetUserId(str string) (model.UserId, error) {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}

	return model.UserId(i), nil
}
