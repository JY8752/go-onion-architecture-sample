package service_test

import (
	"testing"

	"github.com/JY8752/go-onion-architecture-sample/application/service"
	"github.com/JY8752/go-onion-architecture-sample/domain/model"
	"github.com/JY8752/go-onion-architecture-sample/mocks/mock_repository"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	// given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_repository.NewMockTodoRepository(ctrl)

	m.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.Any()).Return(model.TodoId(1), nil)

	ts := service.NewTodoService(m)

	// when
	result, err := ts.Create(1, model.Title("title"), model.Description("description"))
	if err != nil {
		t.Fatal(err)
	}

	// when
	assert.Equal(t, model.TodoId(1), result)
}
