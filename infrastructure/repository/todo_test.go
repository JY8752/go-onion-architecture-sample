package infrastructure_test

import (
	"os"
	"testing"

	"github.com/JY8752/go-onion-architecture-sample/domain/model"
	"github.com/JY8752/go-onion-architecture-sample/domain/repository"
	db "github.com/JY8752/go-onion-architecture-sample/infrastructure"
	infrastructure "github.com/JY8752/go-onion-architecture-sample/infrastructure/repository"
	"github.com/stretchr/testify/assert"
)

var todoRep repository.TodoRepository

func TestMain(m *testing.M) {
	d := db.NewDBClient("file:testdb?mode=memory")
	todoRep = infrastructure.NewTodoRepository(d)

	code := m.Run()

	d.Client.Close() // Exitするとdeferが実行されないので
	os.Exit(code)
}

func TestCreate(t *testing.T) {
	// when
	userId := model.UserId(1)
	id, err := todoRep.Create(userId, "title", "description")
	if err != nil {
		t.Fatal(err)
	}

	// then
	todos, err := todoRep.List(userId)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, len(todos))
	assert.Equal(t, id, todos[0].Id)
	assert.Equal(t, model.Title("title"), todos[0].Title)
	assert.Equal(t, model.Description("description"), todos[0].Description)
}
