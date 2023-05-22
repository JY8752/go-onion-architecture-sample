package service_test

import (
	"testing"

	"github.com/JY8752/go-onion-architecture-sample/application/service"
	"github.com/JY8752/go-onion-architecture-sample/domain/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// UserRepo is a mock version of UserRepo interface
type UserRepoMock struct {
	mock.Mock
}

// Create is a mock function
func (m *UserRepoMock) Create(name string) (model.UserId, error) {
	args := m.Called(name)
	return args.Get(0).(model.UserId), args.Error(1)
}

func (m *UserRepoMock) Get(id model.UserId) (model.User, error) {
	args := m.Called(id)
	return args.Get(0).(model.User), args.Error(1)
}

func TestCreateUser(t *testing.T) {
	// given
	m := new(UserRepoMock)
	m.On("Create", "user1").Return(model.UserId(1), nil)

	userService := service.NewUserService(m)

	// when
	result, err := userService.Create("user1")
	if err != nil {
		t.Fatal(err)
	}

	// then
	assert.Equal(t, model.UserId(1), result)
}
